package valid

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("emailorphone", emailOrPhone)
		v.RegisterValidation("phone", phone)
		v.RegisterValidation("latitude", latitude)
		v.RegisterValidation("longitude", longitude)

	}
}

var emailOrPhone validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	matchedEmail, err1 := regexp.MatchString(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`, value)
	matchedPhone, err2 := regexp.MatchString(`^1[3456789]\d{9}$`, value)
	return (err1 == nil && matchedEmail) || (err2 == nil && matchedPhone)
}

var phone validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	matched, err := regexp.MatchString(`^1[3456789]\d{9}$`, value)
	return err == nil && matched
}

var latitude validator.Func = func(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()
	return lat >= -90 && lat <= 90
}

var longitude validator.Func = func(fl validator.FieldLevel) bool {
	lng := fl.Field().Float()
	return lng >= -180 && lng <= 180
}
