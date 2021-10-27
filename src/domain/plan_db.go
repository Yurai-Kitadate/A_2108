package domain

import "time"

type DBPlan struct {
	ID          int
	Title       string
	Description string
	Image       string
	CreatorID   int           // DBUser.ID
	Conditions  []DBCondition `gorm:"foreignKey:PlanID"`
	Days        []DBDay       `gorm:"foreignKey:PlanID"`
}

type DBCondition struct {
	ID              int
	PlanID          int          // DBPlan.ID
	Seasons         []DBSeason   `gorm:"foreignKey:ConditionID"`
	TimeSpans       []DBTimeSpan `gorm:"foreignKey:ConditionID"`
	Categories      []DBCategory `gorm:"foreignKey:ConditionID"`
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
	Seasons     []DBSeason `gorm:"foreignKey:SeasonDefinitionID"`
}

type DBTimeSpan struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	TimeSpanDefinitionID int // DBTimeSpanDefinition.ID
}

type DBTimeSpanDefinition struct {
	ID          int
	Description string
	TimeSpan    []DBTimeSpan `gorm:"foreignKey:TimeSpanDefinitionID"`
}

type DBCategory struct {
	ID                   int
	ConditionID          int // DBCondition.ID
	CategoryDefinitionID int // DBCategoryDefinition.ID
}

type DBCategoryDefinition struct {
	ID          int
	Description string
	Category    []DBCategory `gorm:"foreignKey:CategoryDefinitionID"`
}

type DBDay struct {
	ID        int
	PlanID    int // DBPlan.ID
	NthDay    int
	Headings  []DBHeading  `gorm:"foreignKey:DayID"`
	Schedules []DBSchedule `gorm:"foreignKey:DayID"`
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
