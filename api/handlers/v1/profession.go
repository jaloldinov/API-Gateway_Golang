package v1

import (
	"context"
	"errors"
	"net/http"

	"api_gateway/api/models"
	"api_gateway/genproto/position_service"
	"api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// Profession godoc
// @ID create-professions
// @Router /v1/profession [POST]
// @Summary create profession
// @Description create profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body position_service.CreateProfessionRequest true "profession"
// @Success 200 {object} models.ResponseModel{data=position_service.Profession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfession(c *gin.Context) {
	var profession position_service.CreateProfessionRequest

	if err := c.BindJSON(&profession); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessionService().Create(c.Request.Context(), &profession)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllProfession godoc
// @ID get-profession
// @Router /v1/profession [GET]
// @Summary get profession all
// @Description get profession
// @Tags profession
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfession(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.ProfessionService().GetAll(
		c.Request.Context(),
		&position_service.GetAllProfessionRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all professions", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-ProfessionByID godoc
// @ID get-profession-byID
// @Router /v1/profession/{profession_id} [GET]
// @Summary get profession by ID
// @Description get profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Success 200 {object} models.ResponseModel{data=models.Profession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProfession(c *gin.Context) {
	var profession models.Profession
	profession_id := c.Param("profession_id")

	if !util.IsValidUUID(profession_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "profession id is not valid", errors.New("profession id is not valid"))
		return
	}

	resp, err := h.services.ProfessionService().Get(
		context.Background(),
		&position_service.ProfessionId{
			Id: profession_id,
		},
	)

	err = ProtoToStruct(&profession, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	if !handleError(h.log, c, err, "error while getting profession") {
		return
	}

	err = ProtoToStruct(&profession, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", profession)

}

// Update Profession godoc
// @ID update_profession
// @Router /v1/profession/{profession_id} [PUT]
// @Summary Update Profession
// @Description Update Profession by ID
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Param profession body models.CreateProfession true "profession"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateProfession(c *gin.Context) {
	var status models.Status
	var profession models.CreateProfession

	profession_id := c.Param("profession_id")

	if !util.IsValidUUID(profession_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid profession id", errors.New("profession id is not valid"))
		return
	}

	err := c.BindJSON(&profession)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	resp, err := h.services.ProfessionService().Update(
		context.Background(),
		&position_service.Profession{
			Id:   profession_id,
			Name: profession.Name,
		},
	)

	if !handleError(h.log, c, err, "error while getting profession") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Updated")
}

/// Delete Profession godoc
// @ID delete-profession
// @Router /v1/profession/{profession_id} [DELETE]
// @Summary delete profession
// @Description Delete Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteProfession(c *gin.Context) {

	var profession models.Profession
	profession_id := c.Param("profession_id")

	if !util.IsValidUUID(profession_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "profession id is not valid", errors.New("profession id is not valid"))
		return
	}

	resp, err := h.services.ProfessionService().Delete(
		context.Background(),
		&position_service.ProfessionId{
			Id: profession_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting profession") {
		return
	}

	err = ProtoToStruct(&profession, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
