package validator

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

const (
	phoneRegexString = "^1[0-9]{10}$"
)

var (
	phoneRegex = regexp.MustCompile(phoneRegexString)
)

// Validate overide
type Validate = validator.Validate

// New a validator
func New() *Validate {
	vd := validator.New()
	vd.RegisterValidation("phone", isPhone)
	return vd
}

func isPhone(f1 validator.FieldLevel) bool {
	return phoneRegex.MatchString(f1.Field().String())
}
