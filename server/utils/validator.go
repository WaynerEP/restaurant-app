package utils

import (
	"github.com/WaynerEP/restaurant-app/server/models/common/response"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	CustomValidate *validator.Validate
	Trans          ut.Translator
)

func Verify(valStruct interface{}) response.ValidationErrors {
	if err := CustomValidate.Struct(valStruct); err != nil {
		validationErrors := make(response.ValidationErrors)

		for _, validationErr := range err.(validator.ValidationErrors) {
			fieldName := validationErr.Field()
			errorMessage := validationErr.Translate(Trans)
			if _, exists := validationErrors[fieldName]; !exists {
				validationErrors[fieldName] = response.ValidationError{
					Type:    validationErr.ActualTag(),
					Message: errorMessage,
				}
			}
		}
		return validationErrors
	}
	return nil
}
