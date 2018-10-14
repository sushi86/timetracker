package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Login(ctx echo.Context) (err error) {

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username == "jon" && password == "snow" {

		token := jwt.New(jwt.SigningMethodES512)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("foobar"))
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	} else {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"status": "Wrong user credentials",
		})
	}
}
