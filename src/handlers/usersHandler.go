package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/repositories"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/schemas"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/security"
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

	hashedPassword, _ := security.HashPassword(request.Password)

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	if err := repositories.NewUserRepository(db).CreateUser(user); err != nil {
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

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
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

	idQuery, _ := strconv.ParseUint(id, 10, 64)
	userID := ctx.MustGet("userID").(uint64)
	if userID != idQuery {
		logger.Error("Unauthorized")
		sendError(ctx, 401, "Unauthorized")
		return
	}

	request := UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
		return
	}

	user := schemas.User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		logger.Error("Failed to find user", err)
		sendError(ctx, 404, err.Error())
		return
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password, _ = security.HashPassword(request.Password)

	if err := db.Save(&user).Error; err != nil {
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

	idQuery, _ := strconv.ParseUint(id, 10, 64)
	userID := ctx.MustGet("userID").(uint64)
	if userID != idQuery {
		logger.Error("Unauthorized")
		sendError(ctx, 401, "Unauthorized")
	}

	if err := db.Where("id = ?", id).Delete(&schemas.User{}).Error; err != nil {
		logger.Error("Failed to delete user", err)
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, 204, nil)
}
