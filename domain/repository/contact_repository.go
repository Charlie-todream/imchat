package repository

import (
	"gorm.io/gorm"
	"time"
	"tobepower/chat/domain/models"
)

type IContactRepository interface {
	FindFrindsById(userid int64) ([]models.User, error)
	FindCommunitysById(userid int64) ([]models.Community, error)
	HasFrind(ownerid int64, dstid int64) *models.Contact
	HasCommunities(ownerid int64, dstid int64) *models.Contact
	InsertFrind(contactOne *models.Contact, contactTwo *models.Contact) error
	JoinCommunity(contact *models.Contact) error
	CreateCommunity(communities *models.Community) (int64, error)
}
type ContactRepository struct {
	mysqlDb *gorm.DB
}

func NewContactRepository(db *gorm.DB) IContactRepository {
	return &ContactRepository{
		mysqlDb: db,
	}
}

func (u *ContactRepository) FindFrindsById(userid int64) ([]models.User, error) {
	conconts := make([]models.Contact, 0)
	objIds := make([]int64, 0)
	u.mysqlDb.Where("ownerid = ? and cate = ?", userid, models.CONCAT_CATE_USER).Find(&conconts)

	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]models.User, 0)
	if len(objIds) == 0 {
		return coms, nil
	}

	return coms, u.mysqlDb.Find(&coms, objIds).Error

}

func (u *ContactRepository) FindCommunitysById(userid int64) ([]models.Community, error) {

	conconts := make([]models.Contact, 0)
	objIds := make([]int64, 0)
	u.mysqlDb.Where("ownerid = ? and cate = ?", userid, models.CONCAT_CATE_COMUNITY).Find(&conconts)

	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}

	coms := make([]models.Community, 0)
	if len(objIds) == 0 {
		return coms, nil
	}

	return coms, u.mysqlDb.Find(&coms, objIds).Error
}

// 查看是否好友存在
func (u *ContactRepository) HasFrind(ownerid int64, dstid int64) *models.Contact {
	contact := &models.Contact{}
	u.mysqlDb.Where("ownerid = ? and dstobj = ? and cate = ?", ownerid, dstid, models.CONCAT_CATE_USER).First(contact)
	return contact
}

// 查看是否加入社群
func (u *ContactRepository) HasCommunities(ownerid int64, dstid int64) *models.Contact {
	contact := &models.Contact{}
	u.mysqlDb.Where("ownerid = ? and dstobj = ? and cate = ?", ownerid, dstid, models.CONCAT_CATE_COMUNITY).First(contact)
	return contact
}

func (u *ContactRepository) InsertFrind(contactOne *models.Contact, contactTwo *models.Contact) error {

	//return contact,u.mysqlDb.Create(&contact).
	u.mysqlDb.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(contactOne).Error; err != nil {
			return err
		}

		if err := tx.Create(contactTwo).Error; err != nil {
			return err
		}
		return nil
	})

	return nil

}

func (u ContactRepository) JoinCommunity(contact *models.Contact) error {
	contact.Createat = time.Now()
	return u.mysqlDb.Create(contact).Error

}

func (u ContactRepository) CreateCommunity(communities *models.Community) (int64, error) {
	communities.Createat = time.Now()
	return communities.Id, u.mysqlDb.Create(communities).Error
}
