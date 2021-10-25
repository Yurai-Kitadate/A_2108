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
type Contacts struct {
	ID        int    `json:"id"`
	Hp        string `json:"hp"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Tiktok    string `json:"tiktok"`
	Biography string `json:"biography"`
}
type User struct {
	ID          int       `json:"id"`
	UserName    string    `json:"userName"`
	Image       string    `json:"image"`
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName"`
	Birthday    time.Time `json:"birthday"`
	Sex         string    `json:"sex"`
	Contacts    Contacts  `json:"contacts"`
	Creator     Creator   `json:"creator"`
	Job         Job       `json:"job"`
	Address     Address   `json:"address"`
}
