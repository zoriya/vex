package main

import (
	huh "github.com/charmbracelet/huh"
)

type Auth struct {
	loginForm    *huh.Form
	registerForm *huh.Form
	jwt          *string
}

func getLoginForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Email").
				Key("email"),
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
				Key("email"),
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
