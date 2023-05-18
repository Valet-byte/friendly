package model

type UserSettings struct {
	Id                   string `json:"-"`
	UserId               string `json:"user_id"`
	IsCloseProfile       bool   `json:"is_close_profile"`
	MessageFromStrangers bool   `json:"message_from_strangers"`
	IsHideMode           bool   `json:"is_hide_mode"`
}
