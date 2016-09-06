package main

import (
	_ "clound_video_api_project/docs"
	_ "clound_video_api_project/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/video_image", "static_image")
	beego.Run()
}
