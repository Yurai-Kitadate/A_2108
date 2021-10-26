package domain

import "time"

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
