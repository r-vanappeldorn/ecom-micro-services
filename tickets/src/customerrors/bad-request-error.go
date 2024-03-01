package customerrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BadRequestError struct {
	*CustomError
	Message string
}

func (e *BadRequestError) Serialize() []ErrorFormat {
	return []ErrorFormat{{Message: e.Message}}
}

func SendBadRequestError(ctx *gin.Context, message string) {
	badreqerr := &BadRequestError{
		&CustomError{Status: http.StatusBadRequest},
		message,
	}

	Send(ctx, badreqerr)
}
