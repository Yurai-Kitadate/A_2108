package domain

import "time"

var (
	Male   = "male"
	Female = "female"
	Other  = "other"
)

type DBUser struct {
	ID          int
	UserName    string
	Email       string
	Password    string
	DisplayName string
	Birthday    time.Time
	Sex         string
}

type DBContacts struct {
	ID        int
	HomePage  string
	Instagram string
	Twitter   string
	Facebook  string
	TikTok    string
	Biography string
}
