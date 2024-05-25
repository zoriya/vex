package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zoriya/vex"
)

type JWTResponse struct {
	Token string `json:"token" validate:"required"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @Tags         Auth
// @Summary      Log in
// @Produce      json
// @Param        DTO body LoginDto true "Login Form"
// @Success      200	{object} JWTResponse
// @Success      403	{object} string
// @Router       /login [post]
func (h *Handler) Login(c echo.Context) error {
	var req LoginDto
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&req); err != nil {
		return err
	}

	user := h.users.GetByEmail(req.Email)
	if user == nil {
		return echo.NewHTTPError(403, "Invalid email")
	}
	if !h.users.CheckPassword(req.Password, user.Password) {
		return echo.NewHTTPError(403, "Invalid password")
	}
	return h.CreateToken(c, user)
}

func (h *Handler) CreateToken(c echo.Context, user *vex.User) error {
	claims := &jwt.RegisteredClaims{
		Subject: user.Id.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, JWTResponse{
		Token: t,
	})
}

// @Tags         Auth
// @Summary      Create an account
// @Produce      json
// @Param        DTO body RegisterDto true "Signup Form"
// @Success      200	{object} JWTResponse
// @Success      403	{object} string
// @Success      409	{object} string
// @Router       /register [post]
func (h *Handler) Register(c echo.Context) error {
	var req RegisterDto
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&req); err != nil {
		return err
	}

	user, err := h.users.Create(req.Name, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(409, "Email already taken")
	}
	return h.CreateToken(c, &user)
}

// @Tags         Auth
// @Summary      Get info about the authed user
// @Produce      json
// @Success      200    {object} vex.User
// @Router       /me [get]
// @Security JWT
func (h *Handler) GetMe(c echo.Context) error {
	id, err := GetCurrentUserId(c)
	if err != nil {
		return err
	}
	user := h.users.GetById(id)
	if user == nil {
		return echo.NewHTTPError(500, "Internal server error")
	}
	return c.JSON(200, user)
}

func GetCurrentUserId(c echo.Context) (uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return uuid.UUID{}, echo.NewHTTPError(401, "Unauthorized")
	}
	sub, err := user.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(403, "Could not retrive subject")
	}
	ret, err := uuid.Parse(sub)
	if err != nil {
		return uuid.UUID{}, echo.NewHTTPError(403, "Invalid id")
	}
	return ret, nil
}

func (h *Handler) RegisterLoginRoutes(e *echo.Echo, r *echo.Group) {
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
	r.GET("/me", h.GetMe)
}
