package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

type ValidatorType map[string]interface{}

var validate *validator.Validate

func Validator(err error) validator.ValidationErrorsTranslations {
	if _, ok := err.(validator.ValidationErrors); !ok {
		return map[string]string{
			"message": "unable to validate",
		}
	}

	return err.(validator.ValidationErrors).Translate(Trans)
}

// func Validator(form interface{}) (fields ValidatorType, err error) {
// 	validate = validator.New()
// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		// Get value from json tag
// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 		if name == "-" {
// 			return ""
// 		}

// 		return name
// 	})

// 	err = validate.Struct(form)
// 	if err != nil {
// 		// Check error is actually from validator
// 		if _, ok := err.(*validator.InvalidValidationError); !ok {
// 			fields = make(ValidatorType)

// 			for _, err := range err.(validator.ValidationErrors) {
// 				msg := rscLangEN.Validation[err.ActualTag()]
// 				if msg == nil {
// 					msg = "verification failed"
// 				}

// 				fieldName := err.Field()
// 				msg = strings.Replace(fmt.Sprintf("%v", msg), ":attribute", fieldName, -1)
// 				fields[fieldName] = msg
// 			}

// 			return fields, err.(validator.ValidationErrors)
// 		}

// 		return nil, err
// 	}

// 	return nil, nil
// }
