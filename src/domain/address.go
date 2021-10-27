package domain

type Place struct {
	ID         int    `json:"id,omitempty"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
}

type Address struct {
	ID          int    `json:"id"`
	Area        string `json:"area"`
	Prefecture  string `json:"prefecture"`
	City        string `json:"city"`
	Description string `json:"description"`
}
