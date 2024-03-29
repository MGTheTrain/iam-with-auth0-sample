package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpController struct{}

// @Summary Public Endpoint
// @Description This route is always accessible.
// @ID public
// @Produce json
// @Success 200 {object} string
// @Security ApiKeyAuth
// @Router /public [get]
func (pc *HttpController) PublicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from a public endpoint! You don't need to be authenticated to see this."})
}

// @Summary Authenticated Endpoint
// @Description Hello from a private endpoint! You need to be authenticated to see this.
// @ID auth
// @Produce json
// @Success 200 {object} string
// @Failure 401 {object} string "Unauthorized. Please provide a valid API key."
// @Security ApiKeyAuth
// @Router /auth [get]
func (ac *HttpController) AuthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from a private endpoint! You need to be authenticated to see this."})
}

// @Summary RBAC restricted Endpoint
// @Description Hello from a private RBAC restricted endpoint! You need the read:messages scope to see this.
// @ID rbac restricted
// @Produce json
// @Success 200 {object} string
// @Failure 401 {object} string "Unauthorized. Please provide a valid API key."
// @Failure 403 {object} string "Permission denied."
// @Security ApiKeyAuth
// @Router /rbac [get]
func (rc *HttpController) RBACEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from a private RBAC restricted endpoint! You need the read:messages scope to see this."})
}

func RegisterPrivateApiEndpoints(router *gin.Engine) {
	httpController := new(HttpController)
	router.GET("/api/v1/sas/auth", httpController.AuthEndpoint)
}

func RegisterPrivateApiEndpointsWithRBAC(router *gin.Engine) {
	httpController := new(HttpController)
	router.GET("/api/v1/sas/rbac", httpController.RBACEndpoint)
}

func RegisterPublicApiEndpoints(router *gin.Engine) {
	httpController := new(HttpController)
	router.GET("/api/v1/sas/public", httpController.PublicEndpoint)
}
