package validator

import (
	"net/http"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

func ValidateRequest(req *http.Request, input interface{}) (err error) {
	err = validate.Struct(input)
	return err
}

func Init() {
	validate = validator.New()
	_ = validate.RegisterValidation("positive", func(fl validator.FieldLevel) bool {
		val, _ := strconv.Atoi(fl.Field().String())
		return val > 0
	})
}
