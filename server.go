package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Type             string `json:"type"`
	SerialNumber     string `json:"serial_number"`
	RegistrationDate string `json:"registration_date"`
	TransferDate     string `json:"transfer_date"`
	DepartmentCode   string `json:"department_code"`
	UsageStatus      string `json:"usage_status"`
	OS               string `json:"os"`
	IPAddress        string `json:"ip_address"`
	Notation         string `json:"notation"`
	SequenceNumber   string `json:"sequence_number"`
}

var Db *gorm.DB
var err error

func init() {
	dsn := "root:mysql@tcp(db:3306)/itam?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&Device{})
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.GET("/devices", getDevices)
	e.GET("/devices/:id", getDevice)
	e.POST("/devices", addDevice)
	e.PUT("/devices/:id", updateDevice)
	e.DELETE("/devices/:id", deleteDevice)

	e.Logger.Fatal(e.Start(":8080"))
}

func getDevices(c echo.Context) error {
	var devices []Device
	Db.Find(&devices)
	return c.JSON(http.StatusOK, devices)
}

func getDevice(c echo.Context) error {
	var device Device
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "id is not integer")
	}
	Db.First(&device, idint)
	return c.JSON(http.StatusOK, device)
}

func addDevice(c echo.Context) error {
	device := new(Device)
	if err := c.Bind(device); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	Db.Create(&device)
	return c.JSON(http.StatusOK, device)
}

func updateDevice(c echo.Context) error {
	var device Device
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "id is not integer")
	}
	Db.First(&device, idint)
	if err := c.Bind(&device); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	Db.Save(&device)
	return c.JSON(http.StatusOK, device)
}

func deleteDevice(c echo.Context) error {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "id is not integer")
	}
	Db.Delete(&Device{}, idint)
	return c.String(http.StatusOK, "deleted")
}
