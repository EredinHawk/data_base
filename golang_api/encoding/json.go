package encoding

import (
	"encoding/json"
	"fmt"
	"golang_api/lib"

	"github.com/go-playground/validator"
)

// EncodingJSON маршалит экземпляр Flights в формат JSON
func EncodingJSON(obj []lib.Flights) ([]byte, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func validJSON(j string, s lib.Search) (bool, error) {
	validate := validator.New()
	err := validate.StructExcept(s, j)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return false, fmt.Errorf(errs.Error())
	}
	return true, nil
}
