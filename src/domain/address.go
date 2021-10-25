package domain

type Address struct {
	ID          int    `json:"id"`
	Area        string `json:"area"`
	Prefecture  string `json:"prefecture"`
	City        string `json:"city"`
	Description string `json:"description"`
}
