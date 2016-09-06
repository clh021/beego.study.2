package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"clound_video_api_project/models"
	"clound_video_api_project/video_logic"
)

// Operations about object
type FrameController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @router /*.* [get]
func (this *FrameController) Get() {

	path := this.Ctx.Input.Param(":path")
	ext := this.Ctx.Input.Param(":ext")
	filename := path + "." + ext
	fmt.Println(filename)
	//不是直接退出
	if strings.Compare(ext, "mp4") != 0 {
		this.Abort("403")
	}
	//判断文件是否存在
	if !models.Video_Exist(filename) {
		this.Abort("404")
	}

	//判断文件缩略图是否存在
	thumb_filename := filename + ".jpg"
	thumb_full_file_path := models.Video_Thumb_Dir + thumb_filename
	if models.Video_Thumb_Exist(thumb_filename) {
		this.ServerImage(thumb_full_file_path)
		return
	}

	//抽取视频第一帧
	result := video_logic.Video_thum(models.Video_Dir+filename, thumb_full_file_path)
	fmt.Println(result)

	if result == 1 {
		this.ServerImage(thumb_full_file_path)
		return
	}
	this.Abort("403")
}

// @Title Get
// @Description find object by objectid
// @router /*.*/[0-9].[0-9]/[0-9].[0-9] [get]
func (this *FrameController) GetFrame() {
	this.Ctx.WriteString("aaaaaaaaaaaaaa")
}

func (this *FrameController) ServerImage(file string) {
	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, file)
}
