package model

import "time"

type Event struct {
	Id            string
	UserId        string
	Name          string
	Address       string
	PhotoLink     string
	PeopleMaxSize int
	PeopleSize    int
	PeopleMinSize int
	FriendSize    int
	StartTime     time.Time
	EndTime       time.Time
	AccessType    string
}

const (
	AllUsers          = "all_users"
	FriendlyOnly      = "friendly_only"
	ByInvite          = "by_invite"
	ByInviteOrRequest = "by_invite_or_request"
)
