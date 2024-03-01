package customerrors

import "github.com/gin-gonic/gin"

type CustomErrorInterface interface {
	GetStatus() int
	Serialize() []ErrorFormat
}

type ErrorFormat struct {
	Message string  `json:"message"`
	Field   *string `json:"field,omitempty"`
}

type CustomError struct {
	Status int
}

func (e *CustomError) GetStatus() int {
	return e.Status
}

func Send(ctx *gin.Context, err CustomErrorInterface) {
	ctx.JSON(err.GetStatus(), gin.H{"errors": err.Serialize()})
}
