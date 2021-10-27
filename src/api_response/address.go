package api_response

type Place struct {
	ID         int    `json:"id,omitempty"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	Name       string `json:"name"`
}

type Address struct {
	ID       int    `json:"id,omitempty"`
	PlusCode string `json:"plusCode"`
}
