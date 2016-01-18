
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, e := sql.Open("mysql", "root:qianno@(192.168.99.100:3306)/lnktest?charset=utf8")
	if e != nil {
		print("ERROR")
		return
	}
	println("Conn DB OK")
	r, e := db.Exec("insert into emp(name, sex, class) values('lib','man','class9')")
	if e != nil {
		fmt.Print("error:%v\n", e)
	}
	fmt.Println("r=", r)
	rows, e := db.Query("select name,id from emp")
	if e != nil {
		fmt.Print("error:%v\n", e)
		return
	}
	if rows == nil {
		print("Rows is nil")
		return
	}
	for rows.Next() {
		var u, p string
		err := rows.Scan(&u, &p)
		if err != nil {
			print("Row error!")
			return
		}
		fmt.Println(u, " ", p)
	}
}