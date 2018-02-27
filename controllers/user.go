package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"echoplus/models"
)

func GetUser(c echo.Context) error {
	// get params
	userId := c.Param("id")
	c.Logger().Printf(userId)
	user := models.GetUserById()
	// return user
	return c.JSON(http.StatusOK, user)
}