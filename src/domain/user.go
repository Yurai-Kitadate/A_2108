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
