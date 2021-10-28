package domain

type DBImage struct {
	ID          int
	OwnerUserID int
	URL         string
}

type Image struct {
	ID          int    `json:"id"`
	OwnerUserID int    `json:"ownerUserID"`
	URL         string `json:"url"`
}
