package repository

import (
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) GetUserByID(id int) (domain.User, error) {
	res := domain.User{}
	db := ur.db

	{
		user := domain.DBUser{}
		err := db.First(&user, id).Error
		switch err {
		case gorm.ErrRecordNotFound:
			return domain.User{}, NotFoundError
		case nil:
		default:
			return domain.User{}, err
		}
		if err != nil {
			return res, NotFoundError
		}
		res.ID = id
		res.UserName = user.UserName
		res.Image = "TODO" // 対応領域ができていない
		res.Email = user.Email
		res.DisplayName = user.DisplayName
		res.Birthday = user.Birthday
		res.Sex = user.Sex
	}

	contacts, err := ur.GetContactsByUserID(id)
	if err != nil {
		return domain.User{}, err
	}
	res.Contacts = contacts

	creator, err := ur.GetCreatorByUserID(id)
	if err != nil {
		return domain.User{}, err
	}
	res.Creator = creator

	job, err := ur.GetJobByUserID(id)
	if err != nil {
		return domain.User{}, err
	}
	res.Job = job

	return res, nil
}

func (ur *UserRepository) GetUserByCreatorID(id int) (domain.User, error) {

	creator := domain.DBCreator{}
	db := ur.db

	err := db.Where("UserID = ?", id).First(&creator).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return domain.User{}, NotFoundError
	case nil:
	default:
		return domain.User{}, err
	}
	return ur.GetUserByID(creator.UserID)
}

func (ur *UserRepository) GetContactsByUserID(id int) (domain.Contacts, error) {
	res := domain.Contacts{}
	contacts := domain.DBContacts{}
	db := ur.db

	err := db.Where("UserID = ?", id).First(&contacts).Error
	switch err {
	case gorm.ErrRecordNotFound:
	case nil:
		res = domain.Contacts{
			ID:        contacts.ID,
			Hp:        contacts.HomePage,
			Instagram: contacts.Instagram,
			Twitter:   contacts.Twitter,
			Facebook:  contacts.Facebook,
			Tiktok:    contacts.TikTok,
			Biography: contacts.Biography,
		}
	default:
		return domain.Contacts{}, err
	}
	return res, nil
}

func (ur *UserRepository) GetCreatorByUserID(id int) (domain.Creator, error) {
	res := domain.Creator{}
	creator := domain.DBCreator{}
	db := ur.db

	err := db.Where("UserID = ?", id).First(&creator).Error
	switch err {
	case gorm.ErrRecordNotFound:
	case nil:
		// TODO
	default:
		return domain.Creator{}, err
	}
	return res, nil
}

func (ur *UserRepository) GetJobByUserID(id int) (domain.Job, error) {
	res := domain.Job{}
	job := domain.DBJob{}
	db := ur.db

	err := db.Where("UserID = ?", id).First(&job).Error
	switch err {
	case gorm.ErrRecordNotFound:
	case nil:
		res = domain.Job{
			ID:             job.ID,
			Jobname:        job.JobName,
			Dateoffirstjob: job.DateOfFirstJob,
		}
	default:
		return domain.Job{}, err
	}
	return res, nil
}
