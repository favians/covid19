package handlers

import (
	"net/http"
	"strconv"

	"github.com/favians/golang_starter/api/models"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

func GetAdmins(c echo.Context) error {
	model := models.Admin{}

	rp, err := strconv.Atoi(c.QueryParam("rp"))
	page, err := strconv.Atoi(c.QueryParam("p"))
	name := c.QueryParam("name")
	instansi := c.QueryParam("instansi")
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"rp":   []string{"numeric"},
		"page": []string{"numeric"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := model.GetList(page, rp, orderby, sort, &models.AdminFilterable{name, instansi})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func GetAdminById(c echo.Context) error {
	model := models.Admin{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"id": []string{"numeric"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := model.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func AddAdmin(c echo.Context) error {
	model := models.Admin{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"name":     []string{"required"},
		"username": []string{"required"},
		"password": []string{"required"},
		"instansi": []string{"required"},
	}

	vld := ValidateRequest(c, rules, &model)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := model.Create()
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new Admin")
	}

	return c.JSON(http.StatusCreated, result)
}

func EditAdmin(c echo.Context) error {
	model := models.Admin{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"name":     []string{},
		"username": []string{},
		"password": []string{},
		"instansi": []string{"required"},
	}

	_, err = model.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	vld := ValidateRequest(c, rules, &model)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	c.Bind(&model)

	err = model.Update()
	if err != nil {
		log.Printf("FAILED TO UPDATE: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update Admin")
	}

	return c.JSON(http.StatusOK, model)
}

func DeleteAdmin(c echo.Context) error {
	model := models.Admin{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"id": []string{"required", "numeric"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	_, err = model.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&model)

	err = model.Delete()
	if err != nil {
		log.Printf("FAILED TO DELETE: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete Admin")
	}

	return c.JSON(http.StatusOK, model)
}
