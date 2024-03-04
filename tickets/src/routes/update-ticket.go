package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/customerrors"
	"ticketing.io/src/helpers"
	"ticketing.io/src/models"
)

type UpdateTicket struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (r *Routes) UpdateTicket(ctx *gin.Context) {
	var req UpdateTicket
	if err := ctx.ShouldBindJSON(&req); err != nil {
		customerrors.SendValidationError(ctx, err)
		return
	}

	if req.Price < 0 {
		customerrors.SendFieldError(ctx, []customerrors.ErrorFormat{{Message: "price must be greater than 0", Field: helpers.PString("price")}})
		return
	}

	// Get ticket
	id := ctx.Param("id")

	ticket := &models.Tickets{}
	ticket.FindByID(r.DB, id)

	if req.Price > 0 {
		ticket.Price = req.Price
	}

	if req.Title != "" {
		ticket.Title = req.Title
	}

	if req.Description != "" {
		ticket.Description = req.Description
	}

	updatedTicket, err := ticket.Update(r.DB)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	// Update ticket
	ctx.JSON(200, gin.H{"ticket": updatedTicket})
}
