package coreDao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"github.com/uuidcode/coreutil"
)

func ConnectDB() *sqlx.DB {
	url := "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8&parseTime=true"
	db, err := sqlx.Connect("mysql", url)
	coreutil.CheckErr(err)

	db.MapperFunc(strcase.ToSnake)

	return db
}
