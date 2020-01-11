package routers

import (
	"github.com/astaxie/beego"
	"smart-shore.club/beego-demo/controllers"
)

func init()  {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Include(&controllers.UserController{})
}
