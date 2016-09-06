package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["clound_video_api_project/controllers/v2:GPSFrameController"] = append(beego.GlobalControllerRouter["clound_video_api_project/controllers/v2:GPSFrameController"],
		beego.ControllerComments{
			"Get",
			`/*.*`,
			[]string{"get"},
			nil})

}
