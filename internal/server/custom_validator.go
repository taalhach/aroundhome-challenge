package server

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (this *CustomValidator) Init() error {

	// uri params validator
	this.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("param"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// query string validator
	this.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// json field validator
	this.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return nil
}

func (this *CustomValidator) Validate(i interface{}) error {
	return this.validator.Struct(i)
}
