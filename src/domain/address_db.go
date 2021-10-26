package domain

type DBAddress struct {
	ID         int
	ScheduleID int
	PlusCode   string
}

type DBPlace struct {
	ID         int
	Area       int // 大まかな地域，四国や東北等
	Prefecture int
	City       int
	Name       string // 名前
}
