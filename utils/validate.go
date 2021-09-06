package utils

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func passwordValidate() func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		password, ok := fl.Field().Interface().(string)
		if ok {
			regexPassword := regexp.MustCompile("(^[^a-z]+$)|(^[^0-9]+$)|(^[^A-Z]+$)")
			ok = !regexPassword.MatchString(password)
			if ok {
				return len(password) >= 8
			}
		}

		return false
	}
}
func NewValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", passwordValidate())
	}
}
