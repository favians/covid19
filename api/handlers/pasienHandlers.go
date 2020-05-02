package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/favians/golang_starter/api/models"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

func GetPasien(c echo.Context) error {
	model := models.Pasien{}

	rp, err := strconv.Atoi(c.QueryParam("rp"))
	page, err := strconv.Atoi(c.QueryParam("p"))
	nama := c.QueryParam("nama")
	jk := c.QueryParam("jk")
	kode := c.QueryParam("kode")
	status := c.QueryParam("status")
	// rs_id, _ := strconv.Atoi(c.QueryParam("rumah_sakit_id"))
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

	result, err := model.GetList(page, rp, orderby, sort, &models.PasienFilterable{nama, jk, kode, status})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetPasienById(c echo.Context) error {
	model := models.Pasien{}

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

func AddPasien(c echo.Context) error {
	model := models.Pasien{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"nama":      []string{"required"},
		"no_hp":     []string{"required"},
		"ttl":       []string{},
		"jk":        []string{},
		"status":    []string{},
		"email":     []string{},
		"longitude": []string{},
		"latitude":  []string{},
	}

	vld := ValidateRequest(c, rules, &model)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	admin := c.Get("admin").(models.Admin)
	log.Println(admin)

	model.Kode = strconv.FormatInt(time.Now().Unix(), 10)
	model.RumahSakitID = admin.RumahSakitID
	model.AdminID = admin.ID
	model.Status = strings.ToUpper(model.Status)

	result, err := model.Create()
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new pasien")
	}

	return c.JSON(http.StatusCreated, result)
}

func EditPasien(c echo.Context) error {
	model := models.Pasien{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"nama":           []string{},
		"no_hp":          []string{},
		"rumah_sakit_id": []string{},
		"ttl":            []string{},
		"jk":             []string{},
		"status":         []string{},
		"email":          []string{},
		"longitude":      []string{},
		"latitude":       []string{},
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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update pasien")
	}

	return c.JSON(http.StatusOK, model)
}

func DeletePasien(c echo.Context) error {
	model := models.Pasien{}

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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete pasien")
	}

	return c.JSON(http.StatusOK, model)
}
