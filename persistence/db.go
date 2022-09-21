package persistence

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/kafugen/ocwcentral/env"
)

func NewDB(env *env.EnvConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&tls=true",
		env.MysqlUser,
		env.MysqlPassowrd,
		env.MysqlAddr,
		env.MysqlDBName,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
