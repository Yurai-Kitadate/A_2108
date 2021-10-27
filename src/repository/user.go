package repository

import (
	"fmt"

	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryError struct {
	s string
}

func (e *UserRepositoryError) Error() string {
	return e.s
}

func (user_repository UserRepository) GetUserByID(userID int) (domain.DBUser, error) {
	db := user_repository.db

	user := domain.DBUser{}
	err := db.Preload("Contacts").
		Preload("Creator").Preload("Job").
		Preload("Address").First(&user, userID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBUser{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		return domain.DBUser{}, &UserRepositoryError{"Other Error"}
	}
	return user, err
}

func (user_repository UserRepository) PostUser(user domain.User, PassWord string) (int, error) {
	db := user_repository.db
	place := domain.DBPlace{
		ID:         user.Place.ID,
		Area:       user.Place.Area,
		Prefecture: user.Place.Prefecture,
		City:       user.Place.City,
		Name:       user.Place.Name,
	}
	contacts := domain.DBContacts{
		ID:        0,
		HomePage:  user.Contacts.Hp,
		Instagram: user.Contacts.Instagram,
		Twitter:   user.Contacts.Twitter,
		Facebook:  user.Contacts.Facebook,
		TikTok:    user.Contacts.Tiktok,
		Biography: user.Contacts.Biography, // NULL にならない NULLable

	}
	var plan []domain.DBPlan

	user_db := domain.DBUser{
		ID:          0,
		UserName:    user.UserName,
		Email:       user.Email,
		Password:    PassWord,
		Place:       place,
		Plans:       plan,
		Contacts:    contacts,
		Image:       user.Image,
		DisplayName: user.DisplayName,
		DateOfBirth: user.DateOfBirth,
		Sex:         user.Sex,
	}
	err := db.Create(&user_db).Error
	return user.ID, err
}

func (user_repository UserRepository) DeleteUserByUserID(userID int) error {
	db := user_repository.db
	err := db.Delete(&domain.DBUser{}, userID).Error
	return err
}

func (user_repository UserRepository) GetUserByCreatorID(creatorID int) (domain.DBUser, error) {
	db := user_repository.db
	user := domain.DBUser{}

	err := db.Preload("Contacts").Preload("Creator").Preload("Job").
		Preload("Address").First(&user, "db_users.creator_id = ?", creatorID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBUser{}, &UserRepositoryError{"Not creator"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return user, err
}

func (user_repository UserRepository) GetContactsByUserID(userID int) (domain.DBContacts, error) {
	db := user_repository.db
	contacts := domain.DBContacts{}

	err := db.First(&contacts).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBContacts{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		return domain.DBContacts{}, &UserRepositoryError{"Other Error"}
	}

	return domain.DBContacts{}, nil
}

func (user_repository UserRepository) PostCreatorByUserID(creator api_response.Creator, userID int) (int, error) {
	db := user_repository.db
	user, err := user_repository.GetUserByID(userID)
	if err != nil {
		if err.Error() == "Not creator" {
			return -1, err
		} else {
			return -1, &UserRepositoryError{"Other Error"}
		}
	}

	job := domain.DBJob{
		ID:             0,
		JobName:        creator.Job.Jobname,
		DateOfFirstJob: creator.Job.DateOfFirstJob,
	}

	creator_db := domain.DBCreator{
		ID:       0,
		RealName: creator.Name,
		Job:      job,
	}

	user.Creator = &creator_db
	err2 := db.Save(&user).Error
	return user.Creator.ID, err2
}

func (user_repository UserRepository) DeleteCreatorByCreatorID(creatorID int) error {
	db := user_repository.db
	err := db.Delete(&domain.Creator{}, creatorID).Error
	if err == gorm.ErrRecordNotFound {
		return &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return err
}

func (user_repository UserRepository) GetJobByCreatorID(creatorID int) (domain.DBJob, error) {
	db := user_repository.db
	job := domain.DBJob{}

	err := db.First(&job, creatorID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBJob{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return job, nil
}

func (user_repository UserRepository) GetPlaceByCreatorID(creatorID int) (domain.DBCreator, error) {
	db := user_repository.db
	place := domain.DBCreator{}

	err := db.First(place, creatorID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBCreator{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return place, nil
}
