package encoding

import (
	"encoding/json"
	"fmt"
	"golang_api/lib"
	"net/http"
)

// ScanBody возвращает экземпляр структуры "Рейс", сканированный из тела HTTP запроса.
// В случае неудачного декодирования ScanBody вернет error с описанием причины.
func ScanBody(r *http.Request) (lib.Search, error) {
	search := lib.Search{}
	if err := json.NewDecoder(r.Body).Decode(&search); err != nil {
		return search, fmt.Errorf("Ошибка при чтении тела http запроса.\n%v", err)
	}
	return search, nil
}
