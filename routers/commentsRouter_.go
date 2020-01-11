package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/users`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "QryUserInfo",
            Router: `/users/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["smart-shore.club/beego-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "QryUserForPage",
            Router: `/users/:pageNo/:pageSize`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
