package customerrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FieldError struct {
	*CustomError
	Errors []ErrorFormat
}

func (e *FieldError) Serialize() []ErrorFormat {
	return e.Errors
}

func SendFieldError(ctx *gin.Context, errs []ErrorFormat) {
	fielderr := &FieldError{
		&CustomError{Status: http.StatusBadRequest},
		errs,
	}

	Send(ctx, fielderr)
}
