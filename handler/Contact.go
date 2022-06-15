package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tobepower/chat/common"
	"tobepower/chat/domain/models"
	"tobepower/chat/domain/repository"
	"tobepower/chat/handler/request"
	request2 "tobepower/chat/handler/request"
	"tobepower/chat/pkg/database"
	"tobepower/chat/pkg/response"
)

type Contact struct {
}

// 好友列表
func (ctr *Contact) LoadFriend(c *gin.Context) {
	request := request2.LoadFrindRequest{}
	common.Bind(c, &request)
	contactRepository := repository.NewContactRepository(database.DB)
	users, err := contactRepository.FindFrindsById(request.Userid)
	if err != nil {

		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "列表查询失败",
			"rows": nil,
		})
		return
	}

	response.JSON(c, gin.H{

		"code": 0,
		"msg":  "成功",
		"rows": users,
	})

}

//  加载群
func (ctr *Contact) LoadCommunity(c *gin.Context) {
	request := request2.LoadFrindRequest{}
	common.Bind(c, &request)
	contactRepository := repository.NewContactRepository(database.DB)
	communitys, err := contactRepository.FindCommunitysById(request.Userid)
	if err != nil {

		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "列表查询失败",
			"rows": nil,
		})
		return
	}
	response.JSON(c, gin.H{

		"code": 0,
		"msg":  "成功",
		"rows": communitys,
	})

}

// 添加好友
func (ctr *Contact) Addfriend(c *gin.Context) {

	request := request2.AddFrindRequest{}
	common.Bind(c, &request)
	contactRepository := repository.NewContactRepository(database.DB)
	contact := contactRepository.HasFrind(request.Userid, request.Dstid)

	if contact.Id > 0 {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "该用户已经被添加过了",
			"rows": nil,
		})
		return
	}

	contactOne := models.Contact{
		Ownerid:  request.Userid,
		Dstobj:   request.Dstid,
		Cate:     models.CONCAT_CATE_USER,
		Createat: time.Now(),
	}

	contactTwo := models.Contact{
		Ownerid:  request.Dstid,
		Dstobj:   request.Userid,
		Cate:     models.CONCAT_CATE_USER,
		Createat: time.Now(),
	}

	err := contactRepository.InsertFrind(&contactOne, &contactTwo)
	if err != nil {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "添加失败",
		})
		return
	}
	response.JSON(c, gin.H{
		"code": 200,
		"msg":  "添加好友成功",
	})

}

// 添加社群
func (ctr *Contact) JoinCommunity(c *gin.Context) {

	request := request2.AddFrindRequest{}
	common.Bind(c, &request)
	contactRepository := repository.NewContactRepository(database.DB)

	// 是否加入过社群
	contact := contactRepository.HasCommunities(request.Userid, request.Dstid)
	fmt.Println(contact.Id)
	if contact.Id > 0 {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "该社群经被添加过了",
			"rows": nil,
		})
		return
	}
	insertContact := models.Contact{
		Ownerid: request.Userid,
		Dstobj:  request.Dstid,
		Cate:    models.CONCAT_CATE_COMUNITY,
	}

	err := contactRepository.JoinCommunity(&insertContact)
	// 当有新的用户加群进入聊天组map中
	AddGroupId(request.Userid, request.Dstid)
	if err != nil {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "添加失败",
		})
		return
	}
	response.JSON(c, gin.H{
		"code": 200,
		"msg":  "添加社群成功",
	})
}

// 创建社群的
func (ctr *Contact) Createcom(c *gin.Context) {
	c.HTML(http.StatusOK, "chat/createcom.html", gin.H{
		"title": "Posts",
	})
}

func (ctr *Contact) Createcommunity(c *gin.Context) {
	requstForm := request2.CreateCommunityRequest{}
	common.Bind(c, &requstForm)

	if ok := request.Validate(c, &requstForm, request.CreateCommunity); !ok {
		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "表單验证失败",
		})
		return
	}
	insertData := &models.Community{
		Name:    requstForm.Name,
		Ownerid: requstForm.Ownerid,
		Icon:    requstForm.Icon,
		Cate:    requstForm.Cate,
		Memo:    requstForm.Memo,
	}

	contactRepository := repository.NewContactRepository(database.DB)
	dstid, err := contactRepository.CreateCommunity(insertData)
	if err != nil {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "失败",
		})

	}
	insertContact := models.Contact{
		Ownerid: requstForm.Ownerid,
		Dstobj:  dstid,
		Cate:    models.CONCAT_CATE_COMUNITY,
	}

	err = contactRepository.JoinCommunity(&insertContact)
	if err != nil {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "添加失败",
		})
		return
	}
	response.JSON(c, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}
