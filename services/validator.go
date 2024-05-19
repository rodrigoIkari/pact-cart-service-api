package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/rodrigoikari/pact-cart-service-api/controllers/models"
)

func Validate(c any) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var element models.ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors

}
