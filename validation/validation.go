package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
)

var (
	validate *validator.Validate
)

func SetupValidation(value interface{}) []error {
	validate = validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	_ = validate.RegisterValidation("ISO8601date", func(fl validator.FieldLevel) bool {
		ISO8601DateRegexString := "^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?$"
		ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)

		return ISO8601DateRegex.MatchString(fl.Field().String())
	})

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(value)
	errs := TranslateError(err, trans)

	return errs
}

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
