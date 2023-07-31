package main

import (
	"applicant-backend/src/controller/devices"
	vulnerabilities2 "applicant-backend/src/controller/vulnerabilities"
	"applicant-backend/src/entities"
	devices2 "applicant-backend/src/routes/devices"
	"applicant-backend/src/routes/vulnerabilities"
	"applicant-backend/src/services/devices_service"
	"applicant-backend/src/services/vulnerability_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	devicesService := devices_service.NewDeviceService()
	vulnerabilityService := vulnerability_service.NewVulnerabilityService()
	vulnerabilityController := vulnerabilities2.NewVulnerabilityController(vulnerabilityService)
	deviceController := devices.NewDeviceController(devicesService)
	deviceRouter := devices2.NewDeviceRouter(deviceController)
	deviceRouter.RegisterRoutes()
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, entities.Error{Message: "method not allowed"})
	})
	r.Use(func(context *gin.Context) {
		context.Next()
		var error string
		for _, err := range context.Errors {
			// log, handle, etc.
			error = err.Error()
		}
		if error != "" {
			context.AbortWithStatusJSON(http.StatusInternalServerError, entities.Error{Message: error})
		}

	})
	vulnerabilityRouter := vulnerabilities.NewVulnerabilityRouter(r, vulnerabilityController)
	deviceRouterGin := devices2.NewGinDeviceRouter(r, deviceController)
	deviceRouterGin.RegisterRoutes()
	vulnerabilityRouter.RegisterRoutes()

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Encountered Gin error: %d", err)
		return
	}
}
