package validator

import (
	"log"

	validator "github.com/go-playground/validator/v10"
)

func Struct(req interface{}) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Printf("req: { %+v} | error: %s", req, err)
		return err
	}

	return nil
}
