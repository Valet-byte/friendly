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
	AccessType    EventAccessType
}

type EventAccessType string

const (
	AllUsers          = EventAccessType("all_users")
	FriendlyOnly      = EventAccessType("friendly_only")
	ByInvite          = EventAccessType("by_invite")
	ByInviteOrRequest = EventAccessType("by_invite_or_request")
)
