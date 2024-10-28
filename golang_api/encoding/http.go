package encoding

import (
	"encoding/json"
	"golang_api/lib"
	"net/http"

	"github.com/go-playground/validator"
)

/*
ScanBody сканирует тело входящего http запроса и декодирует его строки во внутреннюю структуру Search.
Перед этим производится проверка тела запроса на соответствие полей структуры.
*/
func ScanBody(r *http.Request) (*lib.Search, error) {
	var s lib.Search
	validate := validator.New()

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
