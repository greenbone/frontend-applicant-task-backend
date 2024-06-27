package main

import (
	"applicant-backend/src/controller/devicesController"
	"applicant-backend/src/controller/vulnerabilitiesController"
	"applicant-backend/src/helper"
	"applicant-backend/src/internal"
	"applicant-backend/src/services/devices_service"
	"applicant-backend/src/services/vulnerability_service"
	"log"
)

func main() {
	Config, err := helper.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := internal.InitializeRouter()
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
