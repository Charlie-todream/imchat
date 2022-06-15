package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strings"
	"time"
	"tobepower/chat/pkg/response"
)

type Attach struct {
}

func (ctr *Attach) UploadLocal(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		response.MessageError(c, err, "上传文件失败")
	}
	log.Println(file.Filename)
	ofilename := file.Filename
	//todo 创建一个新文件d
	suffix := ".png"

	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := c.Request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	filename := fmt.Sprintf("%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)

	dstfile := "./mnt/" + filename

	if err := c.SaveUploadedFile(file, dstfile); err != nil {
		response.MessageError(c, err, "上传失败")
		return
	}
	url := "/mnt/" + filename

	response.MessageSucess(c, "上传成功", url)

}
