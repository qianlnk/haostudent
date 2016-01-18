/************************************************
File: 		router.go
Version:		0.0.1
Decription:	all Api router
Author:		Qianno.Xie
Email:		qianno.xie@appcaochs.com
History:		2015.12.07 created by Qianno.Xie
************************************************/
package router

import (
	"github.com/go-martini/martini"
	//"haostudent/config"
	"haostudent/handlers/user"
	"haostudent/handlers/common"
	//"html/template"
	//"net/http"
	//"fmt"
)


func NewRouter() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Static("template"))
	m.Get("/", common.LoginHandler)
	m.Get("/ping",common.PingHandler)
	/*m.Group("/user", func(r martini.Router){
		r.Post("/login", user.LoginHandler)
		r.Get("/admin", user.AdminHandler)
	})*/
	m.Post("/userlogin", user.LoginHandler)
	m.Post("/userregister",user.RegisterHandler)
	m.Get("/user", user.UserHandler)
	return m
}
