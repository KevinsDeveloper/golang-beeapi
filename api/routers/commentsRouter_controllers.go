package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego/api/controllers:MainController"] = append(beego.GlobalControllerRouter["beego/api/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:MainController"] = append(beego.GlobalControllerRouter["beego/api/controllers:MainController"],
		beego.ControllerComments{
			Method: "PostLists",
			Router: `/lists`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
