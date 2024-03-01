package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/customerrors"
)

type CreateTicketRequest struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

func (r *Routes) CreateTicket(ctx *gin.Context) {
	var req CreateTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		customerrors.SendValidationError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "ticket created"})
}
