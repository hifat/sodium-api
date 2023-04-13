package utils

import (
	"encoding/json"
	"log"
	"os"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

type Validator struct{}

func (Validator) Validate(err error) validator.ValidationErrorsTranslations {
	if _, ok := err.(validator.ValidationErrors); !ok {
		log.Println(err.Error())
		return map[string]string{
			"message": "unable to validate",
		}
	}

	objErr := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		objErr[e.StructField()] = e.Translate(Trans)
	}

	if os.Getenv("APP_MODE") == "production" {
		jsonBytes, err := json.MarshalIndent(objErr, "", "  ")
		if err != nil {
			return map[string]string{
				"message": "validator error marshalling to JSON",
			}
		}

		log.Println(string(jsonBytes))
		return map[string]string{
			"message": "bad request",
		}
	}

	return objErr
}
