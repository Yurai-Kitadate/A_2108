package domain

type Place struct {
	ID         int    `json:"id"`
	Area       int    `json:"area"`
	Prefecture int    `json:"prefecture"`
	City       int    `json:"city"`
	Name       string `json:"name"`
}

type Address struct {
	ID       int    `json:"id"`
	PlusCode string `json:"plusCode"`
}
