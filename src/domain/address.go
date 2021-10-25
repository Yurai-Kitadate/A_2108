package domain

type DBAddress struct {
	ID         int
	Area       int // 大まかな地域，四国や東北等
	Prefecture int
	City       int
	Name       string // 名前
}

type Address struct {
	ID          int    `json:"id"`
	Area        string `json:"area"`
	Prefecture  string `json:"prefecture"`
	City        string `json:"city"`
	Description string `json:"description"`
}

func (DBAddress) TableName() string {
	return "ADDRESS"
}
