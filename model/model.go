package model

import (
	"clinicSystemGo/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/**
 * DB
 */
var (
	DB     *sqlx.DB
	Config *config.Config
)

func init() {
	Config = config.New()
	var err error
	DB, err = sqlx.Connect("postgres", Config.Postgres.Connect)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
  fmt.Println(Config.Postgres.Connect + " connet success ! ! !")
	DB.SetMaxIdleConns(Config.Postgres.MaxIdle)
	DB.SetMaxOpenConns(Config.Postgres.MaxOpen)
}
