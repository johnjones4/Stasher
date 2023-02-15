package core

import "net/url"

type StructuredDataProperty struct {
	Type       []string                 `json:"type"`
	String     string                   `json:"str,omitempty"`
	Int        int                      `json:"int,omitempty"`
	Float      float64                  `json:"float,omitempty"`
	Bool       bool                     `json:"bool,omitempty"`
	ID         string                   `json:"id,omitempty"`
	Properties []StructuredDataProperty `json:"properties,omitempty"`
}

type Item struct {
	URL         *url.URL
	ContentType string
	Body        []byte
	Name        string
}

type Processor func(req *Item) error

type RuntimeContext struct {
	Processors []Processor
	StashKey   string
	Env        Env
}

type Env struct {
	StashKey            string `env:"STASH_KEY"`
	TelegramSecretToken string `env:"TELEGRAM_SECRET_TOKEN"`
	TelegramAllowedId   string `env:"TELEGRAM_ALLOWED_ID"`
	TelegramAPIToken    string `env:"TELEGRAM_API_TOKEN"`
	DataDir             string `env:"DATA_DIR"`
	HttpHost            string `env:"HTTP_HOST"`
}
