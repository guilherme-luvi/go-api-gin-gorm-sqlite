package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/schemas"
)

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

// ListOpenings represents the request to list all openings
func ListOpenings(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Error("Failed to list openings", err)
		sendError(ctx, 500, "Failed to list openings")
		return
	}

	sendSuccess(ctx, 200, openings)
}

// GetOpening represents the request to get an opening
func GetOpeningById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Error("Invalid id")
		sendError(ctx, 400, "Invalid id")
	}

	opening := schemas.Opening{}

	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.Error("Opening register not found", err)
		sendError(ctx, 404, "Opening register not found")
		return
	}

	sendSuccess(ctx, 200, opening)
}

// UpdateOpening represents the request to update an opening
func UpdateOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Error("Invalid id")
		sendError(ctx, 400, "Invalid id")
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.Error("Opening register not found", err)
		sendError(ctx, 404, "Opening register not found")
		return
	}

	opening.Role = request.Role
	opening.Company = request.Company
	opening.Location = request.Location
	opening.Remote = *request.Remote
	opening.Link = request.Link
	opening.Salary = request.Salary

	if err := db.Save(&opening).Error; err != nil {
		logger.Error("Failed to update opening register", err)
		sendError(ctx, 500, "Failed to update opening resgister")
		return
	}

	sendSuccess(ctx, 200, opening)
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
		sendError(ctx, 404, "Opening register not found")
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Error("Failed to delete opening register", err)
		sendError(ctx, 500, "Failed to delete opening resgister")
		return
	}

	sendSuccess(ctx, 204, "Record deleted")
}
