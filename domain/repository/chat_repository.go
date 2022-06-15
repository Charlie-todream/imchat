package repository

import (
	"gorm.io/gorm"
	"tobepower/chat/domain/models"
)

type IChatRepository interface {
	CheckToken(userId int64, token string) bool
	SearchComunityIds(userId int64) []int64
}
type ChatRepository struct {
	mysqlDb *gorm.DB
}

func NewChatRepository(db *gorm.DB) IChatRepository {
	return &ChatRepository{
		mysqlDb: db,
	}
}

func (u *ChatRepository) CheckToken(userId int64, token string) bool {
	user := models.User{}
	u.mysqlDb.Where("id= ?", userId).Find(&user)
	return user.Token == token
}

func (u *ChatRepository) SearchComunityIds(userId int64) (comIds []int64) {
	conconts := make([]models.Contact, 0)
	comIds = make([]int64, 0)
	u.mysqlDb.Where("ownerid = ? and cate = ?", userId, models.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}
