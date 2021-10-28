package domain

import "time"

type DBPlan struct {
	ID          int
	Title       string
	Description string
	Image       string
	CreatedAt   time.Time
	UserID      int // DBUser.ID
}

type DBCondition struct {
	ID              int
	PlanID          int // DBPlan.ID
	EstimatedCharge int
}

type DBSeason struct {
	ID                 int
	ConditionID        int // DBCondition.ID
	SeasonDefinitionID int // DBSeasonDefinition.ID
}

type DBSeasonDefinition struct {
	ID          int
	Description string
}

type DBTimeSpan struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	TimeSpanDefinitionID int // DBTimeSpanDefinition.ID
}

type DBTimeSpanDefinition struct {
	ID          int
	Description string
}

type DBCategory struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	CategoryDefinitionID int // DBCategoryDefinition.ID
}

type DBCategoryDefinition struct {
	ID          int
	Description string
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
	DayID           int
	Title           string
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	Place           int
	HPLink          string
	ReservationLink string
	Order           int
}
