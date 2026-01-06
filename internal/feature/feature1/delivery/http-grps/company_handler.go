package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/core/domain"
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
	var input struct {
		Name   string `json:"name"`
		UserID uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	company, err := h.companyUC.CreateCompany(
		c.Request.Context(),
		&domain.Company{
			Name:   input.Name,
			UserID: int(input.UserID),
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
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

	c.JSON(http.StatusOK, companies)
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
