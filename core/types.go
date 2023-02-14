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
}
