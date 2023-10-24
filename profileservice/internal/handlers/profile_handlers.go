package handlers

import (
	"net/http"
	"profileservice/internal/models"
	"profileservice/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: service}
}

// Create a new student or professor.
// @Summary Create a new student or professor
// @Description Create a new student or professor with email, password, and user type
// @Accept  json
// @Produce  json
// @Param input body models.CreateRequest true "User creation info"
// @Success 200 {object} models.ResponseModel
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create [post]
func (h *ProfileHandler) Create(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		UserType string `json:"userType" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.profileService.Create(input.Email, input.Password, input.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Read a student or professor by their ID.
// @Summary Get a student or professor by ID
// @Description Get a student or professor by their ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param userType query string true "User Type (student or professor)"
// @Success 200 {object} models.ResponseModel
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /read/{id} [get]
func (h *ProfileHandler) Read(c *gin.Context) {
	id := c.Param("id")
	userType := c.DefaultQuery("userType", "")

	// Parse the ID to uint
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.profileService.Read(uint(idUint), userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update user information.
// @Summary Update user information
// @Description Update user information with email, password, and user type
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body models.UpdateRequest true "User update info"
// @Success 200 {object} models.ResponseModel
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update/{id} [put]
func (h *ProfileHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Parse the ID to uint
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.UpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.profileService.Update(uint(idUint), input.Email, input.Password, input.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Delete a student or professor by their ID.
// @Summary Delete a student or professor by ID
// @Description Delete a student or professor by their ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param userType query string true "User Type (student or professor)"
// @Success 200 {object} models.ResponseModel
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /delete/{id} [delete]
func (h *ProfileHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	userType := c.DefaultQuery("userType", "")

	// Parse the ID to uint
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.profileService.Delete(uint(idUint), userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
