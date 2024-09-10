package handler

// ハンドラーで共通の処理をまとめる

import (
	"errors"
	"fmt"
	"log/slog"
	"upsider-base/shared"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, &shared.ValidationError{}):
		ctx.JSON(400, gin.H{"message": err.Error()})
	case errors.Is(err, &shared.ArgumentError{}):
		ctx.JSON(400, gin.H{"message": err.Error()})
	case errors.Is(err, &shared.UnauthorizedError{}):
		ctx.JSON(401, gin.H{"message": err.Error()})
	case errors.Is(err, &shared.NotFoundError{}):
		ctx.JSON(404, gin.H{"message": err.Error()})
	case errors.Is(err, &shared.ConflictError{}):
		ctx.JSON(409, gin.H{"message": err.Error()})
	default:
		// 意図しないエラーはクライアントに見せるべきではないのでロギングのみにする
		slog.Error(fmt.Sprintf("unexpected error: %+v", err))
		ctx.JSON(500, gin.H{"message": "internal server error"})
	}
}
