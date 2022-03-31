package helpers

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Service interface {
	ValidateInput(document gojsonschema.JSONLoader, validationModel string) (errors []Error)
}

type HelperService struct{}

func NewHelperService() Service {
	return &HelperService{}
}

func (hlp *HelperService) ValidateInput(document gojsonschema.JSONLoader, validationModel string) (errors []Error) {

	v := gojsonschema.NewReferenceLoader(
		fmt.Sprintf("file://%s", "./schema/json/"+validationModel+".json"),
	)
	_, err := v.LoadJSON()
	if err != nil {
		return []Error{
			NewError(ErrBadRequest, "unable to load validation schema"),
		}
	}
	result, err := gojsonschema.Validate(v, document)
	if err != nil {
		return []Error{
			NewError(ErrBadRequest, "unable to validate schema"),
		}
	}

	if result.Valid() {
		return nil
	}
	// msg := fmt.Sprintf("%v validations errors.\n", len(result.Errors()))
	msg := "validations errors"
	for i, desc := range result.Errors() {
		msg += fmt.Sprintf("%v: %s\n", i, desc)
		er := NewError(ErrValidationError, fmt.Sprintf("%s ", desc))
		errors = append(errors, er)
	}
	return errors
}

func ErrorResponse(err interface{}) (convertedMap map[string][]Error) {
	errorResponse := []Error{}
	switch e := err.(type) {
	case Error:
		errorResponse = []Error{e}
	case []Error:
		errorResponse = e
	default:
		errorResponse[0] = NewError(ErrUnknown, fmt.Sprintf("%v", err))
	}
	convertedMap = make(map[string][]Error, 1)
	convertedMap["errors"] = errorResponse
	return convertedMap
}
