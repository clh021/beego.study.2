// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	v1 "clound_video_api_project/controllers/v1"
	v2 "clound_video_api_project/controllers/v2"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(
			&v1.FrameController{},
		),
	)
	ns2 := beego.NewNamespace("/v2",
		beego.NSInclude(
			&v2.GPSFrameController{},
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)
}
