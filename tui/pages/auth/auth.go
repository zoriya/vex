package auth

import (
	"github.com/badoux/checkmail"
	huh "github.com/charmbracelet/huh"
)

type Model struct {
	LoginForm    *huh.Form
	RegisterForm *huh.Form
	Jwt          *string
}

func getLoginForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Email").
				Key("email").Validate(
				func(s string) error {
					err := checkmail.ValidateFormat(s)
					if err != nil {
						return err
					}
					return nil
				}),
			huh.NewInput().
				Title("Password").
				Key("password").
				Password(true),
		)).WithWidth(40)
}

func getRegisterForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Email").
				Key("email").Validate(
				func(s string) error {
					err := checkmail.ValidateFormat(s)
					if err != nil {
						return err
					}
					return nil
				}),
			huh.NewInput().
				Title("Username").
				Key("username"),
			huh.NewInput().
				Title("Password").
				Key("password").
				Password(true),
			huh.NewInput().
				Title("Repeat Password").
				Key("password_repeat").
				Password(true),
		)).WithWidth(40)
}

func New() Model {
	return Model{RegisterForm: getRegisterForm(), LoginForm: getLoginForm(), Jwt: new(string)}

}
