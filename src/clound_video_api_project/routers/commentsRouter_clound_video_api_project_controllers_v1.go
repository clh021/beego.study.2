package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["clound_video_api_project/controllers/v1:FrameController"] = append(beego.GlobalControllerRouter["clound_video_api_project/controllers/v1:FrameController"],
		beego.ControllerComments{
			"Get",
			`/*.*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["clound_video_api_project/controllers/v1:FrameController"] = append(beego.GlobalControllerRouter["clound_video_api_project/controllers/v1:FrameController"],
		beego.ControllerComments{
			"GetFrame",
			`/*.*/[0-9].[0-9]/[0-9].[0-9]`,
			[]string{"get"},
			nil})

}
