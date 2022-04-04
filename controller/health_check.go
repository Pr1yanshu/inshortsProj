package controller

import (
	"github.com/gin-gonic/gin"
	"inshortsProj/logger"
	"net/http"
)

// @Summary health check
// @Schemes
// @Description do health check
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} GetHealthStatus
// @Router /api/covid/healthCheck [get]
func GetHealthStatus(ctx *gin.Context) {
	header := ctx.Request.Header
	logger.ErrorLog.Error(header)
	ctx.JSON(http.StatusOK, "Server is up and running !!!")
}
