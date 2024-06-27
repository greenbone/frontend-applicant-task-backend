package vulnerabilitiesController

import (
	"applicant-backend/src/services/vulnerability_service"
	"github.com/gin-gonic/gin"
)

type VulnerabilityController struct {
	service *vulnerability_service.VulnerabilityService
}

func NewVulnerabilityController(service *vulnerability_service.VulnerabilityService, router gin.IRouter) *VulnerabilityController {
	controller := &VulnerabilityController{
		service: service,
	}
	controller.registerRoutes(router)
	return controller
}

func (v *VulnerabilityController) registerRoutes(router gin.IRouter) {
	group := router.Group("/vulnerabilities")

	group.GET("/", v.GetVulnerabilities)
}

func (v *VulnerabilityController) GetVulnerabilities(context *gin.Context) {
	vulnerabilities := v.service.GetVulnerabilities()
	context.JSON(200, vulnerabilities)
}
