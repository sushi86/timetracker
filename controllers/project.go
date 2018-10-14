package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/sushi86/timemanagement/models"
	"io/ioutil"
	"strconv"
)

func (h *Handler) GetProjects(ctx echo.Context) (err error) {
	db := h.DB
	var projects []models.Project
	db.Find(&projects)
	return ctx.JSON(200, projects)
}

func (h *Handler) GetProject(ctx echo.Context) (err error) {
	db := h.DB
	id, _ := strconv.Atoi(ctx.Param("id"))

	var project models.Project

	db.Find(&project, id)

	return ctx.JSON(200, project)
}

func (h *Handler) CreateProject(ctx echo.Context) (err error) {
	db := h.DB
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	var project models.Project
	json.Unmarshal(bodyBytes, &project)

	db.Save(&project)

	return ctx.JSON(200, project)
}

func (h *Handler) UpdateProject(ctx echo.Context) (err error) {
	db := h.DB
	id, _ := strconv.Atoi(ctx.Param("id"))
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}
	var projectNew models.Project
	json.Unmarshal(bodyBytes, &projectNew)

	res := db.Model(&models.Project{}).Where("id = ?", id).Update("name", projectNew.Name)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, ctx.JSON(200, &Response{Message: "Update Ok", Error: "", Status: "Ok"}))
}

func (h *Handler) DeleteProject(ctx echo.Context) (err error) {
	db := h.DB
	id, _ := strconv.Atoi(ctx.Param("id"))
	var project models.Project
	db.Find(&project, id)
	res := db.Delete(&project)

	if res.Error != nil {
		return ctx.JSON(500, "Something went wrong")
	}

	return ctx.JSON(200, &Response{Message: strconv.Itoa(int(res.RowsAffected)) + " rows affected", Error: "", Status: "Ok"})
}
