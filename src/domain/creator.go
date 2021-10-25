package domain

type Creator struct {
	ID          int    `json:"id,omitempty"`
	Image       string `json:"image,omitempty"`
	DisplayName string `json:"displayName"`
	Job         *Job   `json:"job,omitempty"`
}
