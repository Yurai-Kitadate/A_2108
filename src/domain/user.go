package domain

import "time"

type Contacts struct {
	ID        int     `json:"id"`
	Hp        *string `json:"hp"`
	Instagram *string `json:"instagram"`
	Twitter   *string `json:"twitter"`
	Facebook  *string `json:"facebook"`
	Tiktok    *string `json:"tiktok"`
	Biography *string `json:"biography"` // nullにならないnullable
}

type User struct {
	ID          int       `json:"id"`
	UserName    string    `json:"userName"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	Image       string    `json:"image"`
	DisplayName string    `json:"displayName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Sex         int       `json:"sex"`
	Contacts    Contacts  `json:"contacts"`
	Creator     *Creator  `json:"creator"`
	Place       Place     `json:"address"`
}
type MaskedUser struct {
	ID          int      `json:"id"`
	UserName    string   `json:"userName"`
	Image       string   `json:"image"`
	DisplayName string   `json:"displayName"`
	Contacts    Contacts `json:"contacts"`
	Creator     *Creator `json:"creator"`
}

type Creator struct {
	ID   int    `json:"id"`
	Name string `json:"realName"`
	Job  Job    `json:"job"`
}

type Job struct {
	ID             int       `json:"id"`
	Jobname        string    `json:"jobName"`
	DateOfFirstJob time.Time `json:"dateOfFirstJob"`
}
