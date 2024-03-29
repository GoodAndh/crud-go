package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

func NewDb() *sql.DB {
	datasource := "root:r23password@tcp/glg_restful"
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}


var (
	Store = sessions.NewCookieStore([]byte("sadbhiwqe9120u3bwhkqsbd13901283123=-9=`92`wjigdbjaswg"))
)
