package htmlskip

import (
	"fmt"
	"haostudent/config"
	"html/template"
	"net/http"
)
type LoginView struct {
}

type PingTest struct {
}

type UserView struct{
}

type RegisterView struct{
}

type User struct{
	UserName string
}

func (this *LoginView)LoginSkip(w http.ResponseWriter, r *http.Request){
	fmt.Printf("index View start\n")
	t,_ := template.ParseFiles(config.GetConfig().GetString("haostudent.user.login"))
	t.Execute(w, nil)
	fmt.Printf("index View end\n")
}

func (this *PingTest)PingSkip(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "OK")
}

func (this *UserView)UserSkip(w http.ResponseWriter, r *http.Request, user string){
	t, _ := template.ParseFiles(config.GetConfig().GetString("haostudent.user.index"))
	t.Execute(w, &User{user})
}

func (this *RegisterView)RegisterSkip(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles(config.GetConfig().GetString("haostudent.user.register"))
	t.Execute(w, nil)
}