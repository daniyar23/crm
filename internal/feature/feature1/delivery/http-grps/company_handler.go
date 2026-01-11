package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/core/domain"
	"github.com/daniyar23/crm/internal/feature/feature1/delivery/dto"
	"github.com/daniyar23/crm/internal/feature/feature1/usecase"
)

type CompanyHandler struct {
	companyUC *usecase.CompanyUseCase
}

func NewCompanyHandler(companyUC *usecase.CompanyUseCase) *CompanyHandler {
	return &CompanyHandler{companyUC: companyUC}
}

func (h *CompanyHandler) RegisterRoutes(r *gin.RouterGroup) {
	companies := r.Group("/companies")
	companies.POST("", h.CreateCompany)
	companies.GET("/user/:userId", h.GetByUser)
	companies.DELETE("/:id", h.DeleteCompany)
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req dto.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	company, err := h.companyUC.CreateCompany(
		c.Request.Context(),
		&domain.Company{
			Name:   req.Name,
			UserID: req.UserID,
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.CompanyResponse{
		ID:     company.ID,
		Name:   company.Name,
		UserID: company.UserID,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *CompanyHandler) GetByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	companies, err := h.companyUC.GetCompaniesByUser(
		c.Request.Context(),
		uint(userID),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]dto.CompanyResponse, 0, len(companies))
	for _, cpy := range companies {
		resp = append(resp, dto.CompanyResponse{
			ID:     cpy.ID,
			Name:   cpy.Name,
			UserID: cpy.UserID,
		})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company id"})
		return
	}

	if err := h.companyUC.DeleteCompany(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "company deleted"})
}
