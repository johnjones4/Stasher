package telegram

type Message struct {
	Text string `json:"text"`
}

type User struct {
	Id int `json:"id"`
}

type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type MessageEntity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type IncomingMessage struct {
	Message
	MessageId int             `json:"message_id"`
	Chat      Chat            `json:"chat"`
	From      User            `json:"from"`
	Entities  []MessageEntity `json:"entities"`
	Date      int             `json:"date"`
}

type Update struct {
	UpdateId int             `json:"update_id"`
	Message  IncomingMessage `json:"message"`
}

type OutgoingMessage struct {
	Message
	ChatId int `json:"chat_id"`
}
