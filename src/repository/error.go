package repository

type RepositoryError string

func (e RepositoryError) Error() string {
	return string(e)
}

const (
	NotFoundError = RepositoryError("Can't find correspond item")
)
