package main

import (
	"applicant-backend/src/controller/devices"
	vulnerabilities2 "applicant-backend/src/controller/vulnerabilities"
	"applicant-backend/src/entities"
	"applicant-backend/src/helper"
	devices2 "applicant-backend/src/routes/devices"
	"applicant-backend/src/routes/vulnerabilities"
	"applicant-backend/src/services/devices_service"
	"applicant-backend/src/services/vulnerability_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	Config, err := helper.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	devicesService := devices_service.NewDeviceService()
	vulnerabilityService := vulnerability_service.NewVulnerabilityService()
	vulnerabilityController := vulnerabilities2.NewVulnerabilityController(vulnerabilityService)
	deviceController := devices.NewDeviceController(devicesService)

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

	err = r.Run("0.0.0.0:" + Config.Port)
	if err != nil {
		log.Fatalf("Encountered Gin error: %d", err)
		return
	}
}
