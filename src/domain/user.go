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
	UserID    int // DBUserのID
	HomePage  string
	Instagram string
	Twitter   string
	Facebook  string
	TikTok    string
	Biography string
}

type DBJob struct {
	ID             int
	UserID         int // DBUserのID
	JobName        string
	DateOfFirstJob time.Time
}
type DBCreator struct {
	ID       int
	UserID   int // DBUserのID
	RealName string
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
