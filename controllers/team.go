package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/sushi86/timemanagement/models"
	"io/ioutil"
	"strconv"
)

func (h *Handler) GetTeams(ctx echo.Context) (err error) {
	db := h.DB
	var team []models.Team
	db.Find(&team)
	return ctx.JSON(200, team)
}

func (h *Handler) GetTeam(ctx echo.Context) (err error) {
	db := h.DB
	id, _ := strconv.Atoi(ctx.Param("id"))

	var team models.Team

	db.Find(&team, id)

	return ctx.JSON(200, team)
}

func (h *Handler) CreateTeam(ctx echo.Context) (err error) {
	db := h.DB

	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	var team models.Team
	json.Unmarshal(bodyBytes, &team)

	db.Save(&team)

	return ctx.JSON(200, team)
}

func (h *Handler) UpdateTeam(ctx echo.Context) (err error) {
	db := h.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}
	var teamNew models.Team
	json.Unmarshal(bodyBytes, &teamNew)

	res := db.Model(&models.Team{}).Where("id = ?", id).Update("name", teamNew.Name)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, ctx.JSON(200, &Response{Message: "Update Ok", Error: "", Status: "Ok"}))
}

func (h *Handler) DeleteTeam(ctx echo.Context) (err error) {
	db := h.DB

	id, _ := strconv.Atoi(ctx.Param("id"))
	var team models.Team
	db.Find(&team, id)
	res := db.Delete(&team)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, &Response{Message: strconv.Itoa(int(res.RowsAffected)) + " rows affected", Error: "", Status: "Ok"})
}
