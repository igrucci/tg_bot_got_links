package telegram

// тут все типы с которыми будет работать клиент

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"` // в резалте лежат апдейты
}
type Update struct {
	ID      int              `json:"update_id"` // из доки тг
	Message *IncomingMessage `json:"message"`
}
type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}
type From struct {
	Username string `json:"username"`
}
type Chat struct {
	ID int `json:"id"`
}
