/*package main
import(
	"fmt"
	"os"
	config "github.com/spf13/viper"
)

var conf *config.Viper

func main(){
	conf = config.New()
	conf.SetConfigName("config")
	conf.SetConfigType("toml")
	conf.AddConfigPath(".")
	fmt.Printf("conf:%+v\n", conf)
	err := conf.ReadInConfig()
	if err != nil{
		fmt.Printf("err:%+v", err)
		os.Exit(1)
	}
	
	fmt.Printf("conf.mysql:%+v", conf.GetString("database.mysql"))
}*/
// server
package main

import (
    "net/http"
//    "net/http"
//    "net/http"
    "fmt"
	"html/template"
    "github.com/go-martini/martini"
)

func main() {
    m := martini.Classic()
	m.Use(martini.Static("public"))
    //1 最基础的方式
    m.Get("/", func(w  http.ResponseWriter, r *http.Request){
		t,_ := template.ParseFiles("public/hello.html")
		t.Execute(w, nil)
    })
    
    //2 返回状态,并返回状态码,如200
    //m.Get("/",func()(int,string){
     //   return 405,"状态码测试"
    //})
    
    //3 请求中带有参数 /hello/:name :不能少
    m.Get("/hello/:name",func(params martini.Params)string{
        return "weclome ==&gt; " + params["name"]
    })
    
    //4 使用多个Handler
    // 注意.当多个Handler组合处理的时候,除了最后一个Handler之外,其它Handler不能有返回.换句话说,遇到return就己经终止处理流程,后面的Handler不再执行
    m.Get("/auth",validate,func()string{
        return "恭喜你,通过验证"
    },validate)
    //5 martini中的services,可以理解为注入到Handler参数列表中的对象,可以 全局映射 或者在 请求 的时候映射
    //5.1 global mapping
    person := &Person{Name:"milo"}
    m.Map(person)
    m.Get("/person",personHandler)
    
    //5.2 request mapping 只会绑定在当前这个请求中.其它请求不受影响
    m.Get("/req",func(c martini.Context){
        s := &Student{Name:"req"}
        c.Map(s)
        fmt.Println(s.Name)
    })
    
    //5.3 services 映射到接口中.暂时没想到示例
    
    //6 处理静态文件.martini默认情况下是会把当前应用的根目录下的public当做虚拟目录--官方是这样说的.根据本人实践是放在和main方法入口同级目录下如我的是src/main/server.go,则
    //public应该是src/main/public 假设我里面放了个hello.html,则在浏览器中输入 localhost:3000/hello.html访问
    
    //6.1 指定新的静态文件夹,访问方式不变
    //m.Use(martini.Static("web"))
    
    //7 中间件,这里的中间件有点像Java的拦截器
    m.Use(func(){
        fmt.Println("加入了中间件1.....")
    })
    m.Use(func(){
        fmt.Println("加入了中间件2 ....")
    })
    
    //8 Route groups 路由组
    m.Group("/libaray",func(r martini.Router){
        r.Get("/:id",GetBook)    //http://localhost:3000/libaray/123
        r.Post("/add",AddBook)    //http://localhost:3000/libaray/add
        r.Delete("/del/:id",DelBook) //http://localhost:3000/libaray/del/123
        r.Put("/update/:id",UpdateBook) //http://localhost:3000/libaray/update/122
        
})
    
    m.Use(martini.Static("assets"))
//    m.Run()//默认运行在3000端口
    m.RunOnAddr(":9000")//改变监听的端口  listening on :9000 (development)
}

//有返回值的处理器
//func validate()string {
//    return "pass"
//}

//没有返回值的处理器
func validate(){
    fmt.Println("验证通过")
}


type Person struct{
    Name string
}

type Student struct{
    Name string
}
//获取注入Person的值
func personHandler(person *Person) string{
    fmt.Println(person)
    return person.Name
}

//使用martini.Params获得参数
func GetBook(params martini.Params){
    fmt.Println("正在根据"+params["id"]+"返回golang编程")
}

//post 里面的参数要通过http.Request获取
func AddBook(req * http.Request){
    fmt.Println("正在把"+req.FormValue("name")+"存入数据库")
}

func DelBook(params martini.Params){
    fmt.Println("正在删除"+params["id"]+"编程")
}

func UpdateBook(params martini.Params){
    fmt.Println("正在修改"+params["id"]+"编程")
}