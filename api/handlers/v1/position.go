package v1

import (
	"api_gateway/api/models"
	pb "api_gateway/genproto/position_service"
	"api_gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Position godoc
// @ID create-position
// @Router /v1/position [POST]
// @Summary Create Position
// @Description Create Position
// @Tags position
// @Accept json
// @Produce json
// @Param position body models.CreatePositionModel true "position"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreatePosition(c *gin.Context) {
	var position models.CreatePositionModel
	var pos_attributes []*pb.PositionAttribute
	err := c.BindJSON(&position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	for _, val := range position.PositionAttributes {
		pos_attributes = append(pos_attributes, &pb.PositionAttribute{
			AttributeId: val.AttributeId,
			Value:       val.Value,
		})
	}
	resp, err := h.services.PositionService().Create(
		context.Background(),
		&pb.CreatePositionRequest{
			Name:               position.Name,
			ProfessionId:       position.ProfessionId,
			CompanyId:          position.CompanyId,
			PositionAttributes: pos_attributes,
		},
	)

	if !handleError(h.log, c, err, "error while creating position") {
		return
	}

	err = ProtoToStruct(&position, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Id)
}

// GetAllPosition godoc
// @ID get-position-all
// @Router /v1/position [GET]
// @Summary get position all
// @Description get all position
// @Tags position
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param profession_id query string false "profession_id"
// @Param company_id query string false "company_id"
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllPositionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllPositions(c *gin.Context) {

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}
	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionService().GetAll(
		c.Request.Context(),
		&pb.GetAllPositionRequest{
			Offset:       int64(offset),
			Limit:        int64(limit),
			Name:         c.Query("name"),
			ProfessionId: c.Query("profession_id"),
			CompanyId:    c.Query("company_id"),
		},
	)
	if !handleError(h.log, c, err, "error while getting all positions") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// GetPosition godoc
// @ID get-position-id
// @Router /v1/position/{position_id} [GET]
// @Summary Get Position
// @Description Get Position
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Success 200 {object} models.ResponseModel{data=position_service.GetPositionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetPosition(c *gin.Context) {
	position_id := c.Param("position_id")

	if !util.IsValidUUID(position_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid position id", errors.New("position id is not valid"))
		return
	}

	resp, err := h.services.PositionService().Get(
		context.Background(),
		&pb.PositionId{
			Id: position_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// UpdatePosition godoc
// @ID update-position
// @Router /v1/position/{position_id} [PUT]
// @Summary Update Position
// @Description Update Position by ID
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Param position body position_service.CreatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionStatus} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdatePosition(c *gin.Context) {
	var (
		position pb.GetPositionResponse
	)

	position_id := c.Param("position_id")

	if !util.IsValidUUID(position_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid position id", errors.New("position id is not valid"))
		return
	}

	err := c.BindJSON(&position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().Update(
		context.Background(),
		&pb.GetPositionResponse{
			Id:                 position_id,
			Name:               position.Name,
			ProfessionId:       position.ProfessionId,
			CompanyId:          position.CompanyId,
			PositionAttributes: position.PositionAttributes,
		},
	)

	if !handleError(h.log, c, err, "error while getting position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// Delete Position godoc
// @ID delete-position
// @Router /v1/position/{position_id} [DELETE]
// @Summary Delete Position
// @Description Delete Position by given ID
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionStatus} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeletePosition(c *gin.Context) {

	position_id := c.Param("position_id")

	if !util.IsValidUUID(position_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid position id", errors.New("position id is not valid"))
		return
	}

	resp, err := h.services.PositionService().Delete(
		context.Background(),
		&pb.PositionId{
			Id: position_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting position") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Status)
}

/*
{
	"company_id": "96d9afd3-1221-4783-9228-836d4e132fe0",
	"name": "STRING",
	"position_attributes": [
	  {

		"attribute_id": "b5d52075-ccd9-4e34-b3e8-2fc850765b3b",
		"value": "bitta"
	  }
	],
	"profession_id": "6d3bca61-02df-40ea-aaa5-019c87ea1216"
  } */
