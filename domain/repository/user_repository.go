package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"tobepower/chat/domain/models"
)

type IUserRepository interface {
	Create(user *models.User) error
	FindUser(phone string) (*models.User, error)
	UserCreate(user *models.User) (int64, error)
	UpdateToken(token string, id int64) error
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqlDb: db,
	}
}

func (u *UserRepository) Create(user *models.User) error {
	user.Createat = time.Now()
	return u.mysqlDb.Create(user).Error
}

func (u *UserRepository) FindUser(phone string) (*models.User, error) {
	ruser := &models.User{}
	return ruser, u.mysqlDb.Where("mobile = ?", phone).First(ruser).Error
}

func (u *UserRepository) UpdateToken(token string, id int64) error {
	User := &models.User{}
	return u.mysqlDb.Model(User).Where("id = ?", id).Update("token", token).Error
}

func (u *UserRepository) UserCreate(user *models.User) (int64, error) {
	user.Createat = time.Now()
	db := u.mysqlDb.FirstOrCreate(user, models.User{Mobile: user.Mobile, Password: user.Password, Avatar: user.Avatar, Sex: user.Sex, Nickname: user.Nickname, Salt: user.Salt, Online: user.Online, Token: user.Token, Memo: user.Memo, Createat: user.Createat})
	if db.Error != nil {
		return 0, db.Error
	}

	if db.RowsAffected == 0 {
		return 0, errors.New("插入用户失败")
	}

	return user.Id, nil
}
