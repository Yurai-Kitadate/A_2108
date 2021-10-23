package domain

type DBPlan struct {
	ID          int
	Title       string
	Description string
	Image       string
	CreatorID   string // CreatorのID
}
