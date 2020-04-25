package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/favians/golang_starter/api/models"
	"github.com/favians/golang_starter/bootstrap"
	"github.com/favians/golang_starter/modules/notification"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thedevsaddam/govalidator"
)

func GetReport(c echo.Context) error {
	model := models.Report{}

	rp, err := strconv.Atoi(c.QueryParam("rp"))
	page, err := strconv.Atoi(c.QueryParam("p"))
	kode := c.QueryParam("kode")
	// rs_id, _ := strconv.Atoi(c.QueryParam("rumah_sakit_id"))
	kondisi := c.QueryParam("kondisi")
	suhu := c.QueryParam("suhu")
	demam := c.QueryParam("demam")
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

	result, err := model.GetList(page, rp, orderby, sort, &models.ReportFilterable{kode, kondisi, suhu, demam})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func GetReportById(c echo.Context) error {
	model := models.Report{}

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

func AddReport(c echo.Context) error {
	report := models.Report{}
	pasien := models.Pasien{}

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"kode":           []string{"required"},
		"rumah_sakit_id": []string{"required"},
		"longitude":      []string{"required"},
		"latitude":       []string{"required"},
		"kondisi":        []string{"required"},
		"suhu":           []string{"required"},
		"demam":          []string{"required"},
	}

	vld := ValidateRequest(c, rules, &report)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}
	pasien.FindByCode(report.Kode)

	pLatitude, _ := strconv.ParseFloat(pasien.Latitude, 64)
	pLongitude, _ := strconv.ParseFloat(pasien.Longitude, 64)
	rLatitude, _ := strconv.ParseFloat(report.Latitude, 64)
	rLongitude, _ := strconv.ParseFloat(report.Longitude, 64)

	checkDistance := distance(pLatitude, pLongitude, rLatitude, rLongitude, "K")

	if checkDistance > 1.0 {
		log.Println("sending alert Email")
		SendEmail(pasien.Email)
	}

	log.Println(checkDistance)
	log.Println(pasien)

	result, err := report.Create()
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new Report")
	}

	return c.JSON(http.StatusCreated, result)
}

func EditReport(c echo.Context) error {
	model := models.Report{}

	id, err := strconv.Atoi(c.QueryParam("id"))

	defer c.Request().Body.Close()

	rules := govalidator.MapData{
		"kode":           []string{},
		"rumah_sakit_id": []string{},
		"longitude":      []string{},
		"latitude":       []string{},
		"kondisi":        []string{},
		"suhu":           []string{},
		"demam":          []string{},
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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update Report")
	}

	return c.JSON(http.StatusOK, model)
}

func DeleteReport(c echo.Context) error {
	model := models.Report{}

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
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete Report")
	}

	return c.JSON(http.StatusOK, model)
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func SendEmail(email string) error {

	EmailPayload := notification.EmailPayload{
		To:          email,
		Subject:     "Covid 19 Absensi",
		Message:     "KAMU SUDAH KELUAR JALUR, TOLONG JANGAN PERGI PERGI DULU BRO",
		MessageType: "html",
	}

	res, err := notification.SendEmail(&bootstrap.App.Hedwig, EmailPayload)
	if err != nil {
		Logging(res)
		Logging(err)

		return err
	}
	log.Println(res)

	return nil
}

func Logging(content interface{}) {
	bootstrap.App.Log.Logger.
		WithFields(log.Fields{
			"logCode": "notification_process",
			"error":   content,
		}).
		Error("notification process")
}
