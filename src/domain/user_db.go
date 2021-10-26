package domain

import "time"

const (
	Male = iota
	Female
	Other
)

type DBUser struct {
	ID          int
	UserName    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Password    string
	Places      DBPlace `gorm:"foreignKey:PlaceID"`
	Plans       DBPlan  `gorm:"foreignKey:PlanID"`
	Image       string
	DisplayName string
	DateOfBirth time.Time
	Sex         int
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

type DBCreator struct {
	ID       int
	UserID   int      // DBUserのID
	Job      DBJob    `gorm:"foreignKey:CreatorID"`
	Plans    []DBPlan `gorm:"foreignKey:CreatorID"`
	RealName string
}

type DBJob struct {
	ID             int
	CreatorID      int // DBUserのID
	JobName        string
	DateOfFirstJob time.Time
}
