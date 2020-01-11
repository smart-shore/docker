package controllers

import (
	"github.com/astaxie/beego"
	"smart-shore.club/beego-demo/models"
	"smart-shore.club/beego-demo/utils"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) URLMapping() {
	u.Mapping("QryUserInfo", u.QryUserInfo)
	u.Mapping("QryUserForPage", u.QryUserForPage)
	u.Mapping("AddUserInfo", u.AddUser)
}

// @router /users/ [get]
func (u *UserController) QryUserInfo() {
	user := models.User{}
	user.Id = 1111
	user.Name = "张小小"
	user.Address = "山东济南"
	user.Age = 20
	u.Data["json"] = utils.Success(user)
	u.ServeJSON()
}

//  @router /users/:pageNo/:pageSize  [get]
func (u *UserController) QryUserForPage() {
	pageNo, _ := u.GetInt(":pageNo")
	pageSize, _ := u.GetInt(":pageSize")
	rest := make(map[string]int)
	rest["pageNo"] = pageNo
	rest["pageSize"] = pageSize
	u.Data["json"] = &rest
	u.ServeJSON()
}

// @router /users [post]
func (u *UserController) AddUser() {
	user := models.User{}
	utils.ConvertBytesToJson(u.Ctx.Input.RequestBody,&user)
	u.Data["json"] = utils.Success(&user)
	u.ServeJSON()
}
