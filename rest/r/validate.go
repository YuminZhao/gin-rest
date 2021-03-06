package r

import (
	"errors"
	"gin-rest/config"
	"gin-rest/controllers/validate"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

var val validator.Validate
var translator ut.Translator

type errorType struct {
	Err  error
	Data map[string]string
}

func init() {
	zh := zh.New()
	en := en.New()

	uni := ut.New(zh, en)
	valid := validator.New()

	trans, _ := uni.GetTranslator(config.Server.Language)

	switch config.Server.Language {
	case "en":
		enTranslation.RegisterDefaultTranslations(valid, trans)
	default:
		zhTranslation.RegisterDefaultTranslations(valid, trans)
	}
	validate.Register(valid, trans)
	val = *valid
	translator = trans
}

func Validate(c *gin.Context, params interface{}) errorType {
	c.ShouldBind(params)
	err := val.Struct(params)
	var sliceErr, tran string
	errData := make(map[string]string)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		s := reflect.TypeOf(params).Elem()
		for _, e := range errs {
			tran = e.Translate(translator)
			for i := 0; i < s.NumField(); i++ {
				if s.Field(i).Name == e.StructField() {
					t := s.Field(i).Tag
					l := t.Get("label")
					f := t.Get("form")
					m := t.Get("message")
					if m != "" {
						messages, oks := validate.Message[m]
						if oks {
							message, ok := messages[e.Tag()]
							if ok {
								tran = message
							}
						}
					}
					tran = strings.Replace(tran, s.Field(i).Name, l, -1)
					errData[f] = tran
				}
			}
			if sliceErr == "" {
				sliceErr = tran
			}
		}
		return errorType{
			Err:  errors.New(sliceErr),
			Data: errData,
		}
	}
	return errorType{
		Err: nil,
	}
}
