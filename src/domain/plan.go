package domain

import "time"

type DBPlan struct {
	ID          int
	Title       string
	Description string
	Image       string
	CreatorID   int // DBUser.ID
}

func (DBPlan) TableName() string {
	return "PLAN"
}

type DBCondition struct {
	ID              int
	PlanID          int // DBPlan.ID
	EstimatedCharge int
}

func (DBCondition) TableName() string {
	return "CONDITION"
}

type DBSeason struct {
	ID                 int
	ConditionID        int // DBCondition.ID
	SeasonDefinitionID int // DBSeasonDefinition.ID
}

func (DBSeason) TableName() string {
	return "SEASON"
}

type DBSeasonDefinition struct {
	ID          int
	Description string
}

func (DBSeasonDefinition) TableName() string {
	return "SEASONDEFINITION"
}

type DBTimeSpan struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	TimeSpanDefinitionID int // DBTimeSpanDefinition.ID
}

func (DBTimeSpan) TableName() string {
	return "TIMESPAN"
}

type DBTimeSpanDefinition struct {
	ID          int
	Description string
}

func (DBTimeSpanDefinition) TableName() string {
	return "TIMESPANDEFINITION"
}

type DBCategory struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	CategoryDefinitionID int // DBCategoryDefinition.ID
}

func (DBCategory) TableName() string {
	return "CATEGORY"
}

type DBCategoryDefinition struct {
	ID          int
	Description string
}

func (DBCategoryDefinition) TableName() string {
	return "CATEGORYDEFINITION"
}

type DBDay struct {
	ID     int
	PlanID int // DBPlan.ID
	NthDay int
}

func (DBDay) TableName() string {
	return "DAY"
}

type DBHeading struct {
	ID          int
	DayID       int // DBDay.ID
	HeadingText string
	Order       int
}

func (DBHeading) TableName() string {
	return "HEADING"
}

type DBSchedule struct {
	ID              int
	DayID           int // DBDay.ID
	Title           string
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	Place           int
	HPLink          string
	ReservationLink string
	Order           int
}

func (DBSchedule) TableName() string {
	return "SCHEDULE"
}

type Headings []struct {
	ID    int    `json:"id,omitempty"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}
type Place struct {
	ID         int    `json:"id,omitempty"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
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
