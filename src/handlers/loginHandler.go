package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/auth"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/schemas"
)

func Login(ctx *gin.Context) {
	request := LoginRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Invalid request: %v", err)
		sendError(ctx, 400, err.Error())
	}

	user := schemas.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Where(&user).First(&user).Error; err != nil {
		logger.Error("Failed to find user", err)
		sendError(ctx, 404, err.Error())
		return
	}

	if err := auth.ComparePassword(user.Password, request.Password); err != nil {
		logger.Error("Failed to authenticate user", err)
		sendError(ctx, 401, err.Error())
		return
	}

	// generate token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		logger.Error("Failed to generate token", err)
		sendError(ctx, 500, err.Error())
		return
	}

	sendSuccess(ctx, 200, token)
}
