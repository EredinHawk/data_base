package lib

/*
Search структура определяет параметры для поиска авиабилетов в БД
Эти параметры декодитруются из тела HTTP запроса и в дальнейшем используются для
формирования SQL запросов.
*/
type Search struct {
	City_depart string `json:"city_depart" validate:"required"`
	City_dest   string `json:"city_dest" validate:"required"`
	Date        string `json:"date" validate:"required"`
}

// Flights структура определяет контейнер для результирующего набора данных sql запроса.
type Flights struct {
	F_number       string `json:"number"`
	F_time         string `json:"time"`
	Al_name        string `json:"airline"`
	Ap_Name_depart string `json:"airport_depart"`
	Ap_Name_dest   string `json:"airport_dest"`
	City_depart    string `json:"city_depart"`
	City_dest      string `json:"city_dest"`
	F_cost         string `json:"cost"`
}
