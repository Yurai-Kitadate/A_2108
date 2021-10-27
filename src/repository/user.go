package repository

import (
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "Record Not Found"
}

func (user_repository UserRepository) GetUserByID(userID int, authority bool) (domain.DBUser, error) {
	db := user_repository.db

	{
		user := domain.DBUser{}
		err := db.Preload("Contacts").
			Preload("Creator").Preload("Job").
			Preload("Address").First(&user, userID).Error
		switch err {
		case gorm.ErrRecordNotFound:
			return domain.DBUser{}, &NotFoundError{}
		case nil:
		default:
			return domain.DBUser{}, err
		}
	}

	return domain.DBUser{}, nil
}

func (user_repository UserRepository) PostUser(user api_response.User) (int, error) {
	return -1, nil
}

func (user_repository UserRepository) DeleteUserByUserID(userID int) error {
	return nil
}

func (user_repository UserRepository) GetUserByCreatorID(creatorID int) (domain.DBUser, error) {
	return domain.DBUser{}, nil
}

func (user_repository UserRepository) GetContactsByUserID(userID int) (domain.DBContacts, error) {
	return domain.DBContacts{}, nil
}

func (user_repository UserRepository) PostCreatorByUserID(creator api_response.Creator, userID int) (int, error) {
	return -1, nil
}

func (user_repository UserRepository) DeleteCreatorByCreatorID(creatorID int) error {
	return nil
}

func (user_repository UserRepository) GetJobByCreatorID(creatorID int) (domain.DBJob, error) {
	return domain.DBJob{}, nil
}

func (user_repository UserRepository) GetPlaceByCreatorID(creatorID int) (domain.DBCreator, error) {
	return domain.DBCreator{}, nil
}
