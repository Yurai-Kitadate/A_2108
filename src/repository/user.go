package repository

import (
	"fmt"

	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
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
			res.Creator = nil
			return res, nil
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

	user, err := user_repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Creator != nil {
		err = user_repository.DeleteCreatorByCreatorID(user.Creator.ID)
		if err != nil {
			return err
		}
	}
	db.Delete(&domain.DBContacts{}, user.Contacts.ID)
	db.Delete(&domain.DBPlace{}, user.Place.ID)
	db.Delete(&domain.User{}, userID)

	return err
}

func (user_repository UserRepository) GetUserByCreatorID(creatorID int) (domain.User, error) {
	db := user_repository.db

	creator := domain.DBCreator{}
	err := db.First(&creator).Error
	if err == gorm.ErrRecordNotFound {
		return domain.User{}, &UserRepositoryError{"Not creator"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}

	user, err2 := user_repository.GetUserByID(creator.UserID)
	if err2 == gorm.ErrRecordNotFound {
		return domain.User{}, &UserRepositoryError{"Not creator"}
	} else if err2 != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return user, err2
}

func (ur UserRepository) GetUserByUserName(userName string) (domain.User, error) {
	return ur.getUserBy_("user_name", userName)
}

func (user_repository UserRepository) GetUserByEmail(email string) (domain.User, error) {
	return user_repository.getUserBy_("e_mail", email)
}

// Warning:
// 	fieldに定数以外を使うな! SQLインジェクションを引き起こすので.
func (ur UserRepository) getUserBy_(field string, content string) (domain.User, error) {
	db := ur.db

	db_user := domain.DBUser{}
	query := fmt.Sprintf("%s = ?", field)
	err := db.Where(query, content).First(&db_user).Error
	if err == gorm.ErrRecordNotFound {
		return domain.User{}, &UserRepositoryError{"Not creator"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}

	user, err := ur.GetUserByID(db_user.ID)
	if err == gorm.ErrRecordNotFound {
		return domain.User{}, &UserRepositoryError{"Not creator"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return user, err
}

func (user_repository UserRepository) GetContactsByUserID(userID int) (domain.Contacts, error) {
	db := user_repository.db
	db_contacts := domain.DBContacts{}

	err := db.First(&db_contacts).Error
	if err == gorm.ErrRecordNotFound {
		return domain.Contacts{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
		return domain.Contacts{}, &UserRepositoryError{"Other Error"}
	}

	contacts := domain.Contacts{
		ID:        db_contacts.ID,
		Hp:        db_contacts.HomePage,
		Instagram: db_contacts.Instagram,
		Twitter:   db_contacts.Twitter,
		Facebook:  db_contacts.Facebook,
		Tiktok:    db_contacts.TikTok,
		Biography: db_contacts.Biography,
	}

	return contacts, nil
}

func (user_repository UserRepository) PostCreatorByUserID(creator domain.Creator, userID int) (int, error) {
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
	job := domain.DBJob{}
	err := db.Where("creator_id = ?", creatorID).First(&job).Error
	if err == gorm.ErrRecordNotFound {
		return &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}

	err = db.Delete(&job).Error
	if err == gorm.ErrRecordNotFound {
		return &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}

	err = db.Delete(&domain.Creator{}, creatorID).Error
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

	err := db.Where("creator_id = ?", creatorID).First(&job).Error
	if err == gorm.ErrRecordNotFound {
		return domain.DBJob{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return job, nil
}

func (user_repository UserRepository) GetPlaceByUserID(userID int) (domain.Place, error) {
	db := user_repository.db
	place := domain.Place{}

	err := db.Where("user_id = ?", userID).First(&place).Error
	if err == gorm.ErrRecordNotFound {
		return domain.Place{}, &UserRepositoryError{"Record Not Found"}
	} else if err != nil {
		fmt.Printf("DB Error: %v\n", err)
	}
	return place, nil
}

// ここら辺オーバーヘッドやばいので後で修正します.
func (ur UserRepository) GetIsUniqueEmail(email string) (bool, error) {
	_, err := ur.GetUserByEmail(email)
	if err.Error() == "Record Not Found" {
		return true, nil
	} else if err != nil {
		return false, err
	}
	return false, nil
}

func (ur UserRepository) GetIsUniqueUserName(username string) (bool, error) {
	_, err := ur.GetUserByUserName(username)
	if err.Error() == "Record Not Found" {
		return true, nil
	} else if err != nil {
		return false, err
	}
	return false, nil
}

func (ur UserRepository) PutUser(user domain.User) error {
	err := ur.DeleteUserByUserID(user.ID)
	if err != nil {
		return err
	}
	_, err = ur.PostUser(user)
	if err != nil {
		return err
	}
	return nil
}
