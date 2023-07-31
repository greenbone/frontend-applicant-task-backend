package devices

import (
	"applicant-backend/src/services/devices_service"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	service *devices_service.DeviceService
}

func NewDeviceController(service *devices_service.DeviceService) *DeviceController {
	return &DeviceController{service: service}
}

func (d *DeviceController) GetDevices(context *gin.Context) {
	devices := d.service.GetDevices()
	context.JSON(200, devices)
}

func (d *DeviceController) GetDeviceById(context *gin.Context) {
	device, err := d.service.GetDeviceById(context.Param("oid"))
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(200, device)
}

func (d *DeviceController) GetVulnerabilitiesForDevice(context *gin.Context) {
	oid := context.Param("oid")
	vulns := d.service.GetVulnerabilitiesForDevice(oid)
	context.JSON(200, vulns)
}
