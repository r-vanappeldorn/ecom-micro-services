package customerrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionError struct {
	*CustomError
	Message string
}

func (e *PermissionError) Serialize() []ErrorFormat {
	return []ErrorFormat{{Message: e.Message}}
}

func SendPermissionError(ctx *gin.Context) {
	iserr := &PermissionError{
		&CustomError{Status: http.StatusBadRequest},
		"Permission denied",
	}

	Send(ctx, iserr)
}
