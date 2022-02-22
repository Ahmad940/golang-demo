package validator

import (
	"demo/app/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
func ValidateStruct(obj interface{}) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

