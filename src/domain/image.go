package domain

type DBImage struct {
	ID          int
	OwnerUserID int
	URL         string
}

type Image struct {
	ID  int
	URL string
}
