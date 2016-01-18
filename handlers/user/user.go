package user

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"haostudent/session"
	"haostudent/htmlskip"
	"haostudent/database/mysql"
	"golang.org/x/net/context"
	"github.com/jmoiron/sqlx"
	"reflect"
	//"net/url"
	//"strings"
	//"time"
)

func LoginHandler(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	username := req.PostForm["username"][0]
	password := req.PostForm["password"][0]
	mysession,_ := session.GetStore().Get(req, "session-name")
	mysession.Values["username"] = username
	mysession.Values["password"] = password
	mysession.Save(req, w)
	
	executor := func(ctx context.Context, db *sqlx.DB, dest interface{}) error {
		err := db.Get(&dest, "select 1 from user where email = ? and passed = ?", username, password)
		if err != nil {
			fmt.Println("err=", err)
			return err
		}
		return nil
	}

	var email string
	err := mysql.Invoke(context.TODO(), executor, &email)
	if err != nil{
		fmt.Println("err=", err)
	}
	fmt.Printf("email=%s\n", email)
	
	fmt.Printf("u=%s, p=%s, req=%+v, mysession=%+v\n",username, password, req,mysession)
	fmt.Printf("value = %s %s %s\n", mysession.Values["username"], mysession.Values["password"])
	//登录成功 跳转到用户管理后台
	http.Redirect(w, req, "user", http.StatusFound)
	delete(mysession.Values, "username")
	mysession.Save(req, w)
	fmt.Fprintf(w, "your name is " + username)
}

func UserHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("user start...")
	mysession, err := session.GetStore().Get(r, "session-name")
	fmt.Printf("mysession=%v\n err=%v\n values=%s\n", mysession, err, mysession.Values["username"])
	if err != nil || mysession.Values["username"] == nil{
		fmt.Println("skip...")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	
	user := &htmlskip.UserView{}
	controller := reflect.ValueOf(user)
	method := controller.MethodByName("UserSkip")
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(mysession.Values["username"])
	method.Call([]reflect.Value{responseValue, requestValue, userValue})	
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	email := r.PostForm["Email"][0]
	passwd := r.PostForm["password"][0]
	repasswd := r.PostForm["repassword"][0]
	
	fmt.Printf("Email=%s passwd=%s repasswd=%s\n", email, passwd, repasswd)
	fmt.Fprintf(w, "Email=%s passwd=%s repasswd=%s\n", email, passwd, repasswd)
}
