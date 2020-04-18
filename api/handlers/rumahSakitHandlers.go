package handlers

import (
	"net/http"
	"strconv"

	"github.com/favians/golang_starter/api/models"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

func GetRumahSakit(c echo.Context) error {
	model := models.RumahSakit{}

	rp, err := strconv.Atoi(c.QueryParam("rp"))
	page, err := strconv.Atoi(c.QueryParam("p"))
	nama := c.QueryParam("nama")
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

	result, err := model.GetList(page, rp, orderby, sort, &models.RumahSakitFilterable{nama})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetRumahSakitById(c echo.Context) error {
	model := models.RumahSakit{}

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

func AddRumahSakit(c echo.Context) error {
	model := models.RumahSakit{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"nama":  []string{"required"},
		"lower": []string{"required"},
		"upper": []string{"required"},
		"start": []string{"required"},
		"stop":  []string{"required"},
	}

	vld := ValidateRequest(c, rules, &model)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	result, err := model.Create()
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new Rumah Sakit")
	}

	return c.JSON(http.StatusCreated, result)
}

func EditRumahSakit(c echo.Context) error {
	model := models.RumahSakit{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"nama":  []string{},
		"lower": []string{},
		"upper": []string{},
		"start": []string{},
		"stop":  []string{},
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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update Rumah Sakit")
	}

	return c.JSON(http.StatusOK, model)
}

func DeleteRumahSakit(c echo.Context) error {
	model := models.RumahSakit{}

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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete Rumah Sakit")
	}

	return c.JSON(http.StatusOK, model)
}
