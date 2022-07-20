package v1

import (
	"api_gateway/api/models"
	"api_gateway/genproto/company_service"
	"api_gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCompany godoc
// @ID Create-Company
// @Router /v1/company [POST]
// @Summary create company
// @Description create company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.CreateCompany true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCompany(c *gin.Context) {
	var company company_service.CreateCompany

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Create(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "OK", resp)
}

// GetAllCompany godoc
// @ID Get-All-Companies
// @Router /v1/company [get]
// @Summary get company all
// @Description get company all
// @Tags company
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=company_service.GetAllCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCompanies(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CompanyService().GetAll(
		c.Request.Context(),
		&company_service.GetAllCompanyRequest{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Name:   c.Query("name"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while getting companies", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get Company godoc
// @ID get-company
// @Router /v1/company/{company_id} [GET]
// @Summary get company
// @Description Get Company
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCompany(c *gin.Context) {
	var company company_service.Company
	company_id := c.Param("company_id")

	if !util.IsValidUUID(company_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "company id is not valid", errors.New("company id is not valid"))
		return
	}

	resp, err := h.services.CompanyService().Get(
		c.Request.Context(),
		&company_service.CompanyId{Id: company_id},
	)

	if !handleError(h.log, c, err, "error while getting company") {
		return
	}

	err = ProtoToStruct(&company, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", &company)
}

// UpdateCompany godoc
// @ID update_company
// @Router /v1/company/{company_id} [PUT]
// @Summary Update Company
// @Description Update Company by ID
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Param company body company_service.CreateCompany true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.CompanyResult} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCompany(c *gin.Context) {
	var status models.Status
	var company company_service.CreateCompany

	company_id := c.Param("company_id")

	if !util.IsValidUUID(company_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "company id is not valid", errors.New("company id is not valid"))
		return
	}

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Update(
		context.Background(),
		&company_service.Company{
			Id:   company_id,
			Name: company.Name,
		},
	)

	if !handleError(h.log, c, err, "error while updating company") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", "Updated")
}

// Delete Company godoc
// @ID delete-company
// @Router /v1/company/{company_id} [DELETE]
// @Summary delete company
// @Description Delete Company
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCompany(c *gin.Context) {
	var company company_service.Company
	company_id := c.Param("company_id")

	if !util.IsValidUUID(company_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "company id is not valid", errors.New("company id is not valid"))
		return
	}

	resp, err := h.services.CompanyService().Delete(
		context.Background(),
		&company_service.CompanyId{
			Id: company_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting company") {
		return
	}

	err = ProtoToStruct(&company, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", "deleted")
}
