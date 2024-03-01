package customerrors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalServerError struct {
	*CustomError
	Message string
}

func (e *InternalServerError) Serialize() []ErrorFormat {
	return []ErrorFormat{{Message: e.Message}}
}

func SendInternalServerError(ctx *gin.Context, err error) {
	iserr := &InternalServerError{
		&CustomError{Status: http.StatusBadRequest},
		"internal server error",
	}

	fmt.Println(err)

	Send(ctx, iserr)
}
