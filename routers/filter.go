package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"edroity.com/tool/mindoc/conf"
	"edroity.com/tool/mindoc/models"
	"encoding/json"
)

func init()  {
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session(conf.LoginSessionName).(models.Member)

		if !ok {
			if ctx.Input.IsAjax() {
				jsonData := make(map[string]interface{},3)

				jsonData["errcode"] = 403
				jsonData["message"] = "请登录后再操作"
				returnJSON, _ := json.Marshal(jsonData)

				ctx.ResponseWriter.Write(returnJSON)
			}else{
				ctx.Redirect(302, beego.URLFor("AccountController.Login"))
			}
		}
	}
	beego.InsertFilter("/manager",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/manager/*",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/setting",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/setting/*",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/book",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/book/*",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/api/*",beego.BeforeRouter,FilterUser)

	var FinishRouter = func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Add("MinDoc-Version",conf.VERSION)
		ctx.ResponseWriter.Header().Add("MinDoc-Site","http://www.iminho.me")
	}

	beego.InsertFilter("/*",beego.BeforeRouter ,FinishRouter, false)
}
