package domain

import "time"

type Plans []Plan

type Plan struct {
	PlanId      int         `json:"planID"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	CreatedAt   time.Time   `json:"createdAt"`
	CreatorUser interface{} `json:"creatorUser"` // User | MaskedUser
	Days        Days        `json:"days"`
	Conditions  Conditions  `json:"conditions"`
}

type Headings []struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}

type Schedule []struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	StartTime       int      `json:"startTime"`
	EndTime         int      `json:"endTime"`
	Address         *Address `json:"address"`
	HpLink          *string  `json:"hpLink"`
	ReservationLink *string  `json:"reservationLink"`
	Order           int      `json:"order"`
}

type Days []struct {
	ID       int      `json:"id"`
	NthDay   int      `json:"nthDay"`
	Headings Headings `json:"headings"`
	Schedule Schedule `json:"schedule"`
}

type Season struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type TimeSpan struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Category struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Conditions struct {
	ID       int        `json:"id"`
	Place    []Place    `json:"place"`
	Season   []Season   `json:"season"`
	TimeSpan []TimeSpan `json:"timeSpan"`
	Category []Category `json:"category"`
}
