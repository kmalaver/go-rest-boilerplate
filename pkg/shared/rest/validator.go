package rest

import (
	"errors"
	"reflect"
	"rest/pkg/shared/translate"
	"strings"

	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var Validator = NewValidator()

type ValidatorWrapper struct {
	validator *validator.Validate
}

func NewValidator() *ValidatorWrapper {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// translations
	en_translator, _ := translate.Translator.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, en_translator)

	return &ValidatorWrapper{
		validator: validate,
	}
}

func (v *ValidatorWrapper) Validate(obj interface{}) error {
	err := v.validator.Struct(obj)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		tranlator, _ := translate.Translator.GetTranslator("en")
		errMsg := errors.New(validationErrors[0].Translate(tranlator))
		return errMsg
	}
	return nil
}
