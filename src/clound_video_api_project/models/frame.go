package models

import (
	"os"

	"github.com/astaxie/beego"
)

var (
	Video_Dir       string
	Video_Thumb_Dir string
)

func init() {
	Video_Dir = beego.AppConfig.String("Video_Dir")
	Video_Thumb_Dir = beego.AppConfig.String("Video_Thumb_Dir")
}

// 检查视频文件是否存在
//存在则返回 true，否则返回 false
func Video_Exist(filename string) bool {
	file_full_path := Video_Dir + filename
	_, err := os.Stat(file_full_path)
	return err == nil || os.IsExist(err)
}

// 检查视频缩略图文件是否存在
//存在则返回 true，否则返回 false
func Video_Thumb_Exist(filename string) bool {
	file_full_path := Video_Thumb_Dir + filename
	_, err := os.Stat(file_full_path)
	return err == nil || os.IsExist(err)
}
