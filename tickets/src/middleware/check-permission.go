package middleware

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"ticketing.io/src/customerrors"
	"ticketing.io/src/helpers"
	"ticketing.io/src/models"
)

func CheckPermission(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ticketId := ctx.Param("id")

		ticket := &models.Tickets{}
		ticket.FindByID(db, ticketId)

		user := helpers.GetUser(ctx)

		if ticket.UserID != user.ID {
			customerrors.SendPermissionError(ctx)

			return
		}

		ctx.Next()
	}
}
