package customerrors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ValidationError struct {
	*CustomError
	Err error
}

func (e *ValidationError) Serialize() []ErrorFormat {
	errs := []ErrorFormat{}
	for _, fieldErr := range e.Err.(validator.ValidationErrors) {
		field := fieldErr.Field()
		errs = append(errs, ErrorFormat{Message: SetMessage(fieldErr.Tag(), field), Field: &field})
	}

	return errs
}

func SetMessage(tag, field string) string {
	return fmt.Sprintf("%s %s", cases.Title(language.English, cases.Compact).String(field), MessageForTag(tag))

}

func MessageForTag(tag string) string {
	switch tag {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "must be greater than or equal to 0"
	case "max":
		return "must be less than or equal to 100"
	default:
		return "is invalid"
	}
}

func SendValidationError(ctx *gin.Context, err error) {
	valErr := &ValidationError{
		&CustomError{Status: http.StatusBadRequest},
		err,
	}

	Send(ctx, valErr)
}
