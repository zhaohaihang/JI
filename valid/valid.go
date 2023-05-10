package valid

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("emailorphone", EmailOrPhone)
	}
}

var EmailOrPhone validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	matchedEmail, err1 := regexp.MatchString(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`, value)
	matchedPhone, err2 := regexp.MatchString(`^1[3456789]\d{9}$`, value)
	return (err1 == nil && matchedEmail) || (err2 == nil && matchedPhone)
}
