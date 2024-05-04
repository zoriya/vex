package vex

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Email    string
	Password []byte
}

type UserService struct {
	database *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{database: db}
}

func (s UserService) GetById(id uuid.UUID) *User {
	var user User
	err := s.database.Get(&user, "select u.* from users as u where u.id = $1", id)
	if err != nil {
		return nil
	}
	return &user
}

func (s UserService) GetByEmail(email string) *User {
	var user User
	err := s.database.Get(&user, "select u.* from users as u where u.email = $1", email)
	if err != nil {
		return nil
	}
	return &user
}

func (s UserService) CheckPassword(password string, reference []byte) bool {
	return bcrypt.CompareHashAndPassword(reference, []byte(password)) == nil
}

func (s UserService) Create(name string, email string, password string) (User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	user := User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: pass,
	}
	_, err = s.database.NamedExec(
		`insert into users (id, name, email, password)
		values (:id, :name, :email, :password)`,
		user,
	)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
