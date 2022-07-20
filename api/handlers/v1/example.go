package v1

import (
	"api_gateway/pkg/logger"

	"api_gateway/api/models"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @ID ping
// @Router /v1/ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Failure 500 {object} models.ResponseModel{error=string}
func (h *handlerV1) Ping(c *gin.Context) {
	c.JSON(200, models.ResponseModel{
		Code:    200,
		Message: "ok",
		Data:    "pong",
	})
}

// GetConfig godoc
// @ID get-config
// @Router /config [GET]
// @Summary gets project config
// @Description shows config of the project only on the development phase
// @Tags config
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Success 200 {object} models.ResponseModel{data=config.Config} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetConfig(c *gin.Context) {
	h.log.Info("get config", logger.String("environment", h.cfg.Environment))
	switch h.cfg.Environment {
	case "development":
		h.handleSuccessResponse(c, 200, "ok", h.cfg)
		return
	case "staging":
		h.handleSuccessResponse(c, 200, h.cfg.Environment, nil)
		return
	case "production":
		h.handleSuccessResponse(c, 200, "private data", nil)
		return
	}

	h.handleErrorResponse(c, 400, "wrong environment", h.cfg.Environment)
}
