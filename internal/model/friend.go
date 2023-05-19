package model

type Friends struct {
	UserUid1 string
	UserUid2 string
	Type     string
}

const (
	Default = "default"
	Request = "request"
)
