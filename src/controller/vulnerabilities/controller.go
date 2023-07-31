package vulnerabilities

import (
	"applicant-backend/src/services/vulnerability_service"
	"github.com/gin-gonic/gin"
)

type VulnerabilityController struct {
	service *vulnerability_service.VulnerabilityService
}

func NewVulnerabilityController(service *vulnerability_service.VulnerabilityService) *VulnerabilityController {
	return &VulnerabilityController{service: service}
}

func (d *VulnerabilityController) GetVulnerabilities(context *gin.Context) {
	devices := d.service.GetVulnerabilities()
	context.JSON(200, devices)
}
