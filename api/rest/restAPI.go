package rest

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/resource/langEN"
	"github.com/hifat/hifat-blog-api/internal/routes"
	"github.com/hifat/hifat-blog-api/internal/utils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func API() {
	/* ---------------------------- Validator config ---------------------------- */
	var trans ut.Translator

	binding.Validator.Engine().(*validator.Validate).RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		if _, ok := langEN.Validate[name]; !ok {
			return name
		}

		return langEN.Validate[name]
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}

	utils.Trans = trans

	/* --------------------------- Running API server --------------------------- */
	router := gin.Default()
	api := router.Group("/api")

	routes.AuthRoute(api)

	router.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
