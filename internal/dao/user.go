package dao

import (
	"gin_api_frame/internal/model"
	"gin_api_frame/pkg/database"

	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(db *database.Database) *UserDao {
	return &UserDao{
		DB: db.Mysql,
	}
}

func (userDao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	err = userDao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

func (userDao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return userDao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

func (userDao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = userDao.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = userDao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (userDao *UserDao) CreateUser(user *model.User) error {
	return userDao.DB.Model(&model.User{}).Create(&user).Error
}

func (userDao *UserDao) UpdateLastLoginById(uId uint, loginTime int64) (err error) {
	return userDao.DB.Model(&model.User{}).Select("last_login").Where("id=?", uId).Updates(&model.User{LastLogin: loginTime}).Error
}

func (userDao *UserDao) UpdateUserAvatarById(uId uint, path string) (err error) {
	return userDao.DB.Model(&model.User{}).Select("avatar").Where("id=?", uId).Updates(&model.User{Avatar: path}).Error
}

