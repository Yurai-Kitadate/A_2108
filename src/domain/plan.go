package domain

import "time"

type Plan struct {
	PlanId      int         `json:"planId,omitempty"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Creator     *Creator    `json:"creator,omitempty"`
	Days        Days        `json:"days,omitempty"`
	Conditions  *Conditions `json:"conditions,omitempty"`
}
type Plans []Plan

type Headings []struct {
	ID    int    `json:"id,omitempty"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}

type Schedule []struct {
	ID              int       `json:"id,omitempty"`
	Description     string    `json:"description"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"EndTime"`
	Place           Place     `json:"place"`
	HpLink          string    `json:"hpLink"`
	ReservationLink string    `json:"reservationLink"`
	Order           int       `json:"order"`
}
type Days []struct {
	Headings Headings `json:"headings"`
	Schedule Schedule `json:"schedule"`
}
type Season []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type TimeSpan []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type Category []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type Conditions struct {
	ID       int      `json:"id,omitempty"`
	Season   Season   `json:"season"`
	TimeSpan TimeSpan `json:"timeSpan"`
	Category Category `json:"category"`
}
