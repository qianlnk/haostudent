/************************************************
File: 		common.go 
Version:		0.0.1
Description:	common handler
Author:		Qianno.Xie
Email:		qianno.xie@appcoachs.com
History:		2015.12.09 created by Qianno.Xie
************************************************/
package common

import(
	"haostudent/htmlskip"
	"net/http"
	"reflect"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	login := &htmlskip.LoginView{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName("LoginSkip")
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func PingHandler(w http.ResponseWriter, r *http.Request){
	ping := &htmlskip.PingTest{}
	controller := reflect.ValueOf(ping)
	method := controller.MethodByName("PingSkip")
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}