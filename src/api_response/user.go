package api_response

import "time"

type Contacts struct {
	ID        int     `json:"id"`
	Hp        *string `json:"hp"`
	Instagram *string `json:"instagram"`
	Twitter   *string `json:"twitter"`
	Facebook  *string `json:"facebook"`
	Tiktok    *string `json:"tiktok"`
	Biography *string `json:"biography"`
}

type User struct {
	ID          int       `json:"id"`
	UserName    string    `json:"userName"`
	Email       string    `json:"email"`
	Image       string    `json:"image"`
	DisplayName string    `json:"displayName"`
	Birthday    time.Time `json:"birthday"`
	Sex         string    `json:"sex"`
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
	Name string `json:"realname"`
	Job  Job    `json:"job"`
}

type Job struct {
	ID             int       `json:"id"`
	Jobname        string    `json:"jobname"`
	DateOfFirstJob time.Time `json:"dateoffirstjob"`
}
