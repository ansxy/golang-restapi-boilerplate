package custom_validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	custom_error "github.com/ansxy/golang-boilerplate-gin/pkg/error"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator(v *validator.Validate) Validator {
	return Validator{
		validate: v,
	}
}

func (v *Validator) ValidateStruct(r *http.Request, values interface{}) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, values)

	if err != nil {
		fmt.Printf("Error Parse Body %v", err.Error())
		return err
	}

	err = v.validate.Struct(values)

	if err != nil {
		var missingField []string

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			missingField = append(missingField, field)

			msg := fmt.Sprintf("Fields %v are required", strings.Join(missingField, field))

			err := custom_error.SetCustomeError(&custom_error.ErrorContext{
				Code:    constant.DefaultValidateErrorCode,
				Message: msg,
			})
			return err

		}
	}

	return nil

}
