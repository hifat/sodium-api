package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type ValidatorType map[string]interface{}

var validate *validator.Validate

func Validator(form interface{}) (fields ValidatorType, err error) {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err = validate.Struct(form)
	if err != nil {
		// Check error is actually from validator
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			fields = make(ValidatorType)

			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("1" + err.ActualTag())
				// err.Param() <--- value of options such as max=10 will return 10
				fieldName := err.Field()
				fields[fieldName] = err.ActualTag()
			}

			return fields, err.(validator.ValidationErrors)
		}

		return nil, err
	}

	return nil, nil
}
