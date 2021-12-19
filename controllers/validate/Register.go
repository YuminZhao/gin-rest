package validate

import (
	"gin-rest/config"
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Register(validate *validator.Validate, translator ut.Translator) {
	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.Match("^[a-z]{6,30}$", []byte(fl.Field().String()))
		return matched
	})
	switch config.Server.Language {
	case "en":
	default:
		validate.RegisterTranslation("username", translator, func(ut ut.Translator) error {
			return ut.Add("username", "{0}格式不正确", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("username", fe.Field())
			return t
		})
	}
}
