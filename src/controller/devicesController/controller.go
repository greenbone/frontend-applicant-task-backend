package devicesController

import (
	"applicant-backend/src/services/devices_service"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	service *devices_service.DeviceService
}

func NewDeviceController(service *devices_service.DeviceService, router gin.IRouter) *DeviceController {
	controller := &DeviceController{
		service: service,
	}
	controller.registerRoutes(router)
	return controller
}

func (d *DeviceController) registerRoutes(router gin.IRouter) {
	group := router.Group("/devices")

	group.GET("/", d.GetDevices)
	group.GET("/:oid", d.GetDeviceById)
	group.GET("/:oid/vulnerabilities", d.GetVulnerabilitiesForDevice)
}

func (d *DeviceController) GetDevices(context *gin.Context) {
	devices := d.service.GetDevices()
	context.JSON(200, devices)
}

func (d *DeviceController) GetDeviceById(context *gin.Context) {
	oid := context.Param("oid")
	device, err := d.service.GetDeviceById(oid)
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(200, device)
}

func (d *DeviceController) GetVulnerabilitiesForDevice(context *gin.Context) {
	oid := context.Param("oid")
	vulns := d.service.GetVulnerabilitiesForDevice(oid)
	context.JSON(200, vulns)
}
