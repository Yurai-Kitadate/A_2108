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
	DateOfBirth time.Time
	Sex         string
}

func (DBUser) TableName() string {
	return "USER"
}

type DBContacts struct {
	ID        int
	UserID    int // DBUserのID
	HomePage  string
	Instagram string
	Twitter   string
	Facebook  string
	TikTok    string
	Biography string
}

func (DBContacts) TableName() string {
	return "CONTACTS"
}

type DBJob struct {
	ID             int
	UserID         int // DBUserのID
	JobName        string
	DateOfFirstJob time.Time
}

func (DBJob) TableName() string {
	return "JOB"
}

type DBCreator struct {
	ID       int
	UserID   int // DBUserのID
	RealName string
}

func (DBCreator) TableName() string {
	return "CREATOR"
}
