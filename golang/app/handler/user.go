package handler

import (
	"fmt"
	"net/http"

	"app/database"
	"app/model"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []model.User{}
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func GetMe(c echo.Context) error {
	fmt.Println("uid")
	uid := userIDFromToken(c)
	user := model.User{}
	user.Id = uid
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Take(&user)

	return c.JSON(http.StatusOK, user)
}