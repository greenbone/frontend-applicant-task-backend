package devices

import (
	"applicant-backend/src/controller/devices"
	"github.com/gin-gonic/gin"
)

type GinDeviceRouter struct {
	controller *devices.DeviceController
	Rg         *gin.RouterGroup
}

func NewGinDeviceRouter(engine *gin.Engine, controller *devices.DeviceController) *GinDeviceRouter {
	return &GinDeviceRouter{Rg: engine.Group("/devices"), controller: controller}
}

func (r *GinDeviceRouter) RegisterRoutes() {
	r.Rg.GET("/", r.controller.GetDevices)
	r.Rg.GET("/:oid", r.controller.GetDeviceById)
	r.Rg.GET("/:oid/vulnerabilities", r.controller.GetVulnerabilitiesForDevice)
}
