package initialize

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/newit-hieutm/go-backend/global"
	"github.com/newit-hieutm/go-backend/pkg/sqlc/godev"
)

func InitMysql() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=%t&loc=%s",
		global.ConfigGlobal.Db.User,
		global.ConfigGlobal.Db.Password,
		global.ConfigGlobal.Db.Host,
		global.ConfigGlobal.Db.Port,
		global.ConfigGlobal.Db.Database,
		global.ConfigGlobal.Db.Charset,
		global.ConfigGlobal.Db.ParseTime,
		global.ConfigGlobal.Db.Loc)

	fmt.Println("dsn", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connect to Mysql successfully")

	global.Dbtx = db
	global.Queries = godev.New(global.Dbtx)

	fmt.Println("Queries", global.Queries)

	GetAuthorsEagerLoadBooksRow()
}

func GetAuthorsEagerLoadBooksRow() {
	authors, err := global.Queries.GetAuthorsEagerLoadBooks(context.Background())

	if err != nil {
		panic(err)
	}

	for _, value := range(authors){
		fmt.Println(value)
	}
}
