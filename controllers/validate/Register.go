package validate

import (
	"fmt"
	"gin-rest/config"
	"gin-rest/rest/m"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Register(validate *validator.Validate, translator ut.Translator) {

	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		var count int64
		value := fmt.Sprintf("%v", fl.Field())
		param := strings.Split(fl.Param(), "%")
		table := param[0]
		field := fl.FieldName()
		if len(param) > 1 {
			field = param[1]
		}
		err := m.DB.Table(table).Where(field+" = ?", value).Count(&count).Error
		if err != nil {
			return false
		}
		if count > 0 {
			return false
		}
		return true
	})

	validate.RegisterValidation("exists", func(fl validator.FieldLevel) bool {
		var count int64
		value := fmt.Sprintf("%v", fl.Field())
		param := strings.Split(fl.Param(), "%")
		table := param[0]
		field := fl.FieldName()
		if len(param) > 1 {
			field = param[1]
		}
		err := m.DB.Table(table).Where(field+" = ?", value).Count(&count).Error
		if err != nil {
			return false
		}
		if count > 0 {
			return true
		}
		return false
	})

	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.Match("^[a-z]{3,16}$", []byte(fl.Field().String()))
		return matched
	})
	switch config.Server.Language {
	case "en":
	default:

		validate.RegisterTranslation("unique", translator, func(ut ut.Translator) error {
			return ut.Add("unique", "{0}已存在", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("unique", fe.Field())
			return t
		})

		validate.RegisterTranslation("exists", translator, func(ut ut.Translator) error {
			return ut.Add("exists", "{0}不存在", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("exists", fe.Field())
			return t
		})

		validate.RegisterTranslation("username", translator, func(ut ut.Translator) error {
			return ut.Add("username", "{0}格式不正确", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("username", fe.Field())
			return t
		})
	}
}
