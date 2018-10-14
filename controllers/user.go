package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/sushi86/timemanagement/models"
	"io/ioutil"
	"strconv"
)

func (h *Handler) GetUser(ctx echo.Context) (err error) {
	db := h.DB
	var user []models.User
	db.Find(&user)
	return ctx.JSON(200, user)
}

func (h *Handler) CreateUser(ctx echo.Context) (err error) {
	db := h.DB
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	var user models.User
	json.Unmarshal(bodyBytes, &user)

	db.Save(&user)

	return ctx.JSON(200, user)
}

func (h *Handler) UpdateUser(ctx echo.Context) (err error) {
	db := h.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}
	var userNew models.User
	json.Unmarshal(bodyBytes, &userNew)

	res := db.Model(&models.User{}).Where("id = ?", id).Update("name", userNew.Name)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, ctx.JSON(200, &Response{Message: "Update Ok", Error: "", Status: "Ok"}))
}

func (h *Handler) DeleteUser(ctx echo.Context) (err error) {
	db := h.DB
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user models.User
	db.Find(&user, id)
	res := db.Delete(&user)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, &Response{Message: strconv.Itoa(int(res.RowsAffected)) + " rows affected", Error: "", Status: "Ok"});
}
