package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/daniyar23/crm/internal/domain"
	"github.com/daniyar23/crm/internal/services"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func (h *CompanyHandler) RegisterRoutes(r *gin.RouterGroup) {
	companies := r.Group("/companies")

	companies.POST("", h.CreateCompany)
	companies.GET("", h.GetAllCompanies)
	companies.GET("/:id", h.GetCompanyByID)
	companies.DELETE("/:id", h.DeleteCompany)
}

func NewCompanyHandler(service *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: service}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var company domain.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	id, err := h.companyService.CreateCompany(c.Request.Context(), &company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid company id",
		})
		return
	}

	company, err := h.companyService.GetCompanyByID(
		c.Request.Context(),
		uint(id64),
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, company)
}
func (h *CompanyHandler) GetAllCompanies(c *gin.Context) {
	companies, err := h.companyService.GetAllCompanies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companies)
}
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid company id",
		})
		return
	}

	err = h.companyService.DeleteCompany(c.Request.Context(), uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "company deleted successfully",
	})
}
