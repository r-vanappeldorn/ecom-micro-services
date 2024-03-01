package helpers

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func SetUser(ctx *gin.Context, user *User) {
	ctx.Set("user", user)
}

func GetUser(ctx *gin.Context) *User {
	return ctx.MustGet("user").(*User)
}
