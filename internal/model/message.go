package model

type Message struct {
	From        string
	To          string
	Body        string
	MessageType string
}

const (
	TEXT     = "text"
	PICTURE  = "picture"
	DOCUMENT = "document"
)
