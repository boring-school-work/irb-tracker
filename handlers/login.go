package handlers

import (
	"net/http"
	"strings"

	"github.com/boring-school-work/irb-tracker/actions"
	"github.com/labstack/echo/v4"
)

type loginData struct {
	Email  string
	Passwd string
	Error  string
}

// LoginView renders the login page
func LoginView(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// LoginUser logs in a user
func LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	passwd := c.FormValue("passwd")

	_, err := actions.LoginUser(email, passwd)
	if err != nil {
		if strings.Contains(err.Error(), "email") || strings.Contains(err.Error(), "password") {
			data := loginData{
				Email:  email,
				Passwd: passwd,
				Error:  "Incorrect email or password",
			}

			return c.Render(http.StatusBadRequest, "login", data)
		} else {
			return err
		}
	}

	return c.Render(http.StatusOK, "dashboard", nil)

	// TODO: if email and passwd is correct
	// 1. create a session
	// 2. store username and role in session (role -> student/faculty)
	// 3. redirect to /dashboard
}
