package domain

import "time"

type TextData struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

type CredData struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

type CardData struct {
	ID         int       `json:"id"`
	CardNumber string    `json:"card_number"`
	ExpDate    time.Time `json:"exp_date"`
	CVC        string    `json:"cvc"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Metadata   string    `json:"metadata"`
}

type BlobData struct {
	ID       int    `json:"id"`
	Data     []byte `json:"data"`
	Metadata string `json:"metadata"`
}
