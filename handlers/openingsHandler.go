package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/config"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}

// Openings CRUD:

// CreateOpening represents the request to create an opening
func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Error("Failed to create opening register", err)
		sendError(ctx, 500, "Failed to create opening resgister")
		return
	}

	sendSuccess(ctx, 201, opening)
}

// GetOpening represents the request to get an opening
func GetOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET /api/v1/opening",
	})
}

// UpdateOpening represents the request to update an opening
func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT /api/v1/opening",
	})
}

// DeleteOpening represents the request to delete an opening
func DeleteOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Error("Invalid id")
		sendError(ctx, 400, "Invalid id")
	}

	opening := schemas.Opening{}

	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.Error("Opening register not found", err)
		sendError(ctx, 404, "Opening not found")
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Error("Failed to delete opening register", err)
		sendError(ctx, 500, "Failed to delete opening resgister")
		return
	}

	sendSuccess(ctx, 204, "Record deleted")
}

// ListOpenings represents the request to list all openings
func ListOpenings(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET /api/v1/openings",
	})
}
