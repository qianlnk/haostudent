package mysql

import (
	"fmt"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"haostudent/config"
	"golang.org/x/net/context"
)

type Handler func(context.Context, *sqlx.DB, interface{}) error

var (
	once sync.Once
	db *sqlx.DB
)

func Invoke(ctx context.Context, handle Handler, dest interface{}) error {
	return handle(ctx, db, dest)
}

func init() {
	once.Do(initializing)
}

func initializing() {
	conn, err := sqlx.Connect("mysql", config.GetConfig().GetString("database.mysql"))
	if err != nil {
		fmt.Printf("mysql connect err: %+v\n", err)
	}
	db = conn
}

func GetDB() *sqlx.DB {
	return db
}