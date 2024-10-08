package handler

import (
	"upsider-coding-test/usecase"

	"github.com/gin-gonic/gin"
)

type (
	UserHandler interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}
	userHandler struct {
		userUsecase usecase.UserUsecase
	}
	UserCreateParams struct {
		Username    string `json:"name" binding:"required"`
		Email       string `json:"email" binding:"required"`
		RawPassword string `json:"password" binding:"required"`
		CompanyID   string `json:"company_id" binding:"required"`
	}
	UserLoginParams struct {
		Email       string `json:"email" binding:"required"`
		RawPassword string `json:"password" binding:"required"`
	}
)

func (h *userHandler) Register(ctx *gin.Context) {
	var params UserCreateParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, err := h.userUsecase.Register(&usecase.RegisterInput{
		Username:    params.Username,
		Email:       params.Email,
		RawPassword: params.RawPassword,
		CompanyID:   params.CompanyID,
	})
	if err == nil {
		ctx.JSON(200, gin.H{
			"id": user.ID().String(),
		})
		return
	}
	handleError(ctx, err)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var params UserLoginParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	token, err := h.userUsecase.Login(&usecase.LoginInput{
		Email:       params.Email,
		RawPassword: params.RawPassword,
	})
	if err == nil {
		ctx.JSON(200, gin.H{"token": token.String()})
		return
	}
	handleError(ctx, err)
}
