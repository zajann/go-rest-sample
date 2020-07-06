package controllers

import (
	"go-rest-test/internal/db"
	"go-rest-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := db.InsertUser(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := db.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func GetAllUser(c echo.Context) error {
	users, err := db.GetAllUser()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = id
	if err := db.UpdateUserByID(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := db.DeleteUserByID(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
