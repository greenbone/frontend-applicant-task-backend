package vulnerabilities

import (
	"applicant-backend/src/controller/vulnerabilities"
	"github.com/gin-gonic/gin"
)

type VulnerabilityRouter struct {
	controller *vulnerabilities.VulnerabilityController
	Rg         *gin.RouterGroup
}

func NewVulnerabilityRouter(engine *gin.Engine, controller *vulnerabilities.VulnerabilityController) *VulnerabilityRouter {
	return &VulnerabilityRouter{Rg: engine.Group("/vulnerabilities"), controller: controller}
}

func (r *VulnerabilityRouter) RegisterRoutes() {
	r.Rg.GET("/", r.controller.GetVulnerabilities)
}
