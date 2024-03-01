package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/customerrors"
	"ticketing.io/src/helpers"
	"ticketing.io/src/models"
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

	if req.Price < 0 {
		customerrors.SendFieldError(ctx, []customerrors.ErrorFormat{{Message: "price must be greater than 0", Field: helpers.PString("price")}})
		return
	}

	user := helpers.GetUser(ctx)
	ticket, err := models.NewTicket(req.Title, user.ID, req.Price).Create(r.DB)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"ticket": ticket})
}
