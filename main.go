package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE clinic(
	code 				varchar(20)			PRIMARY KEY     NOT NULL,
	name        varchar(40)			NOT NULL,
	responsible_person        varchar(40)			NOT NULL,
	area        varchar(40),
	status        boolean			NOT NULL		DEFAULT true,
	create_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP
);

CREATE TABLE admin(
	username        varchar(20)			NOT NULL,
 	phone        varchar(11)			NOT NULL,
	password        varchar(40)			NOT NULL,
	clinic_code        varchar(40)		NOT NULL  references clinic(code),
	status        boolean			NOT NULL		DEFAULT true,
	create_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP,
 	update_time   timestamp				NOT NULL		DEFAULT LOCALTIMESTAMP,
 	PRIMARY KEY (username, clinic_code)
);`

// type Person struct {
// 	FirstName string `db:"first_name"`
// 	LastName  string `db:"last_name"`
// 	Email     string
// }

// type Place struct {
// 	Country string
// 	City    sql.NullString
// 	TelCode int
// }

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/initDb", func(ctx iris.Context) {
		db, err := sqlx.Connect("postgres", "user=kangcha dbname=mydb sslmode=disable password=123456")
		if err != nil {
			log.Fatalln(err)
		}
		db.MustExec(schema)
		ctx.Writef("ok %d")
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
