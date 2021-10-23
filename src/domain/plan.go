package domain

import "time"

type DBPlan struct {
	ID          int
	Title       string
	Description string
	Image       string
	CreatorID   int // DBUser.ID
}

type DBCondition struct {
	ID     int
	PlanID int // DBPlan.ID
}

type DBSeason struct {
	ID                 int
	ConditionID        int // DBCondition.ID
	SeasonDefinitionID int
}

type DBSeasonDefinition struct {
	ID          int
	Description string
}

type DBTimeSpan struct {
	ID          int
	ConditionID int // DBCondition.ID
	// TODO
}

type DBCategory struct {
	ID          int
	ConditionID int // DBCondition.ID
	// TODO
}

type DBDay struct {
	ID     int
	PlanID int // DBPlan.ID
	NthDay int
}

type DBHeading struct {
	ID          int
	DayID       int // DBDay.ID
	HeadingText string
	Order       int
}

type DBSchedule struct {
	ID              int
	DayID           int // DBDay.ID
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	Place           int
	HPLink          string
	ReservationLink string
	Order           int
}
