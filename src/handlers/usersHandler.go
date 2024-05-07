package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/schemas"
)

// Users CRUD:

// CreateUser creates a new user
func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Error("Failed to create user", err)
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, 201, user)
}

// GetUserById retrieves a user by its ID
func GetUserById(ctx *gin.Context) {
	id := ctx.Query("id")

	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		logger.Error("Failed to find user", err)
		sendError(ctx, 404, err.Error())
		return
	}

	sendSuccess(ctx, 200, user)
}

// UpdateUser updates a user
func UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Error("Invalid id")
		sendError(ctx, 400, "Invalid id")
		return
	}

	request := UpdateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		logger.Error("Failed to update user", err)
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, 200, user)
}

// DeleteUser deletes a user
func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Error("Invalid id")
		sendError(ctx, 400, "Invalid id")
		return
	}

	user := schemas.User{}

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		logger.Error("Failed to delete user", err)
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, 204, nil)
}
