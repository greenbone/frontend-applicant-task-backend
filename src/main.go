package main

import (
	"applicant-backend/src/controller/devicesController"
	"applicant-backend/src/controller/vulnerabilitiesController"
	"applicant-backend/src/entities"
	"applicant-backend/src/helper"
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

	vulnerabilityService := vulnerability_service.NewVulnerabilityService()
	devicesService := devices_service.NewDeviceService(vulnerabilityService)

	_ = vulnerabilitiesController.NewVulnerabilityController(vulnerabilityService, r)
	_ = devicesController.NewDeviceController(devicesService, r)

	err = r.Run("0.0.0.0:" + Config.Port)
	if err != nil {
		log.Fatalf("Encountered Gin error: %d", err)
		return
	}
}
