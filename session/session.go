/************************************************
Flie:		session.go
Version:		0.0.1
Description:	session manage
Author:		Qainno.Xie
Email:		qianno.xie@appcoachs.com
History:		2015.12.08 created by Qianno.Xie
************************************************/
package session

import(
	"sync"
	//"github.com/martini-contrib/sessions"
	"github.com/gorilla/sessions"
)

var (
	once sync.Once
	store *sessions.CookieStore
	mysession *sessions.Session
)

func init(){
	once.Do(initilaizing)
}

func initilaizing(){
	store = sessions.NewCookieStore([]byte("HAO-STUDENT-QIANNO-XIE"))
}

func GetStore() *sessions.CookieStore{
	return store
}