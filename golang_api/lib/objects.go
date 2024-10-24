package lib

// Парметры поиска авиабилетов
type Search struct {
	City_depart string `json:"city_depart"`
	City_dest   string `json:"city_dest"`
	Date        string `json:"date"`
}

// Авиабилеты
type Flights struct {
	F_id         int    `json:"id"`
	F_number     string `json:"number"`
	F_airline_id int    `json:"airline_id"`
	F_time       string `json:"time"`
	F_cost       string `json:"cost"`
}
