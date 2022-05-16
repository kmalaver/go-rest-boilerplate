package translate

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

var Translator *ut.UniversalTranslator

func init() {
	en := en.New()

	Translator = ut.New(
		en, // default

		// list of languages
		en,
	)
}
