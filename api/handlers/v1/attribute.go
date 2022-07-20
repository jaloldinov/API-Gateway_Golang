package v1

import (
	"api_gateway/api/models"
	"api_gateway/genproto/position_service"
	"api_gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profession godoc
// @ID create-profession
// @Router /v1/attribute [POST]
// @Summary create profession
// @Description create attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.CreateAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateAttribute(c *gin.Context) {
	var attribute position_service.CreateAttributeRequest

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().Create(c.Request.Context(), &attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating Attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "Created", resp)
}

// Get All Attribute godoc
// @ID get-all-attribute
// @Router /v1/attribute [GET]
// @Summary get all attribute
// @Description Get All Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllAttributes(c *gin.Context) {

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.AttributeService().GetAll(
		context.Background(),
		&position_service.GetAllAttributeRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all attributes") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get Attribute godoc
// @ID get-attribute
// @Router /v1/attribute/{attribute_id} [GET]
// @Summary get attribute
// @Description Get Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAttribute(c *gin.Context) {
	attribute_id := c.Param("attribute_id")

	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute id is not valid", errors.New("attribute id is not valid"))
		return
	}

	resp, err := h.services.AttributeService().Get(
		context.Background(),
		&position_service.AttributeId{
			Id: attribute_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Attribute godoc
// @ID update_attribute
// @Router /v1/attribute/{attribute_id} [PUT]
// @Summary Update Attribute
// @Description Update Attribute by ID
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Param attribute body position_service.CreateAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateAttribute(c *gin.Context) {
	var status models.Status
	var attribute position_service.CreateAttributeRequest

	attribute_id := c.Param("attribute_id")

	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid attribute id", errors.New("attribute id is not valid"))
		return
	}

	err := c.BindJSON(&attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().Update(
		context.Background(),
		&position_service.Attribute{
			Id:            attribute_id,
			Name:          attribute.Name,
			AttributeType: attribute.AttributeType,
		},
	)

	if !handleError(h.log, c, err, "error while getting attribute") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", "Updated")
}

/// Delete Attribute godoc
// @ID delete-attribute
// @Router /v1/attribute/{attribute_id} [DELETE]
// @Summary delete attribute
// @Description Delete Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteAttribute(c *gin.Context) {

	attribute_id := c.Param("attribute_id")

	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute id is not valid", errors.New("attribute id is not valid"))
		return
	}

	_, err := h.services.AttributeService().Delete(
		context.Background(),
		&position_service.AttributeId{
			Id: attribute_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
