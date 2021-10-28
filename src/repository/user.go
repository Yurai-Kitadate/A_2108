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

func (user_repository UserRepository) GetUserByID(userID int) (domain.User, error) {
	db := user_repository.db
	res := domain.User{}

	user := domain.DBUser{}
	{
		err := db.First(&user).Error
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, &UserRepositoryError{"Record Not Found"}
		} else if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return domain.User{}, &UserRepositoryError{"Other Error"}
		}
		res.ID = user.ID
		res.UserName = user.UserName
		res.Email = user.Email
		res.Password = user.Password
		res.Image = user.Image
		res.DisplayName = user.DisplayName
		res.DateOfBirth = user.DateOfBirth
		res.Sex = user.Sex
	}

	{
		contacts := domain.Contacts{}
		db_contacts := domain.DBContacts{}
		err := db.Where("user_id = ?", res.ID).First(&db_contacts).Error
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, &UserRepositoryError{"Record Not Found"}
		} else if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return domain.User{}, &UserRepositoryError{"Other Error"}
		}
		contacts.ID = db_contacts.ID
		contacts.Hp = db_contacts.HomePage
		contacts.Instagram = db_contacts.Instagram
		contacts.Twitter = db_contacts.Twitter
		contacts.Facebook = db_contacts.Facebook
		contacts.Tiktok = db_contacts.TikTok
		contacts.Biography = db_contacts.Biography
		res.Contacts = contacts
	}

	{
		creator := domain.Creator{}
		db_creator := domain.DBCreator{}
		err := db.Where("user_id = ?", res.ID).First(&db_creator).Error
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, &UserRepositoryError{"Not Creator"}
		} else if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return domain.User{}, &UserRepositoryError{"Other Error"}
		}
		creator.ID = db_creator.ID
		creator.Name = db_creator.RealName
		res.Creator = &creator
	}

	{
		job := domain.Job{}
		db_job := domain.DBJob{}

		err := db.Where("creator_id = ?", res.Creator.ID).First(&db_job).Error
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, &UserRepositoryError{"Record Not Found"}
		} else if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return domain.User{}, &UserRepositoryError{"Other Error"}
		}
		job.ID = db_job.ID
		job.Jobname = db_job.JobName
		job.DateOfFirstJob = db_job.DateOfFirstJob
		res.Creator.Job = job
	}
	return res, nil
}

func (user_repository UserRepository) PostUser(user domain.User) (int, error) {
	db := user_repository.db

	err := db.Transaction(func(tx *gorm.DB) error {
		user_db := domain.DBUser{
			UserName:    user.UserName,
			Email:       user.Email,
			Password:    user.Password,
			Image:       user.Image,
			DisplayName: user.DisplayName,
			DateOfBirth: user.DateOfBirth,
			Sex:         user.Sex,
		}
		err := db.Create(&user_db).Error
		if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return &UserRepositoryError{"Other Error"}
		}

		contacts := domain.DBContacts{
			UserID:    user_db.ID,
			HomePage:  user.Contacts.Hp,
			Instagram: user.Contacts.Instagram,
			Twitter:   user.Contacts.Twitter,
			Facebook:  user.Contacts.Facebook,
			TikTok:    user.Contacts.Tiktok,
			Biography: user.Contacts.Biography, // NULL にならない NULLable
		}
		err = db.Create(&contacts).Error
		if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return &UserRepositoryError{"Other Error"}
		}

		place := domain.DBPlace{
			Area:       user.Place.Area,
			Prefecture: user.Place.Prefecture,
			City:       user.Place.City,
			Name:       user.Place.Name,
		}
		err = db.Create(&place).Error
		if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			return &UserRepositoryError{"Other Error"}
		}
		return nil
	})

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

	user, _ := user_repository.GetUserByID(userID)
	creator_db := domain.DBCreator{
		UserID:   user.ID,
		RealName: user.Creator.Name,
	}

	err := db.Create(&creator_db).Error
	if err != nil {
		if err.Error() == "Not creator" {
			return -1, err
		} else {
			return -1, &UserRepositoryError{"Other Error"}
		}
	}

	job := domain.DBJob{
		CreatorID:      user.Creator.Job.ID,
		JobName:        user.Creator.Job.Jobname,
		DateOfFirstJob: user.Creator.Job.DateOfFirstJob,
	}
	err = db.Create(&job).Error
	if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		return 0, &UserRepositoryError{"Other Error"}
	}

	return user.ID, nil
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
