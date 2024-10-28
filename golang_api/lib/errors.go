package lib

const (
	// Ошибка обработчика GetFlights
	ScanBodyError     = "GetFlights: ошибка при сканировании тела http запроса\n%v"
	ConnectionBDError = "GetFlights: ошибка при подключении к БД\n%v"
	QueryBDError      = "GetFlights: ошибка при выполнении sql запроса\n%v"
	MarshalError      = "GetFlights: ошибка при кодировании JSON объекта\n%v"
)
