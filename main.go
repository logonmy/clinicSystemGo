package main

import (
	"clinicSystemGo/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/lib/pq"
)

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

	clinic := app.Party("/clinic")
	{
		clinic.Post("/add", controller.ClinicAdd)
		clinic.Post("/detailByCode", controller.GetClinicByCode)
		clinic.Post("/list", controller.ClinicList)
		// clinic.Post("/update", controller.ClinicUpdte)
	}

	department := app.Party("/department")
	{
		department.Post("/create", controller.DepartmentCreate)
		department.Post("/list", controller.DepartmentList)
		department.Post("/delete", controller.DepartmentDelete)
		department.Post("/update", controller.DepartmentUpdate)
	}

	personnel := app.Party("/personnel")
	{
		personnel.Post("/login", controller.PersonnelLogin)
		personnel.Post("/create", controller.PersonnelCreate)
		personnel.Post("/getById", controller.PersonnelGetByID)
		personnel.Post("/list", controller.PersonnelList)
	}

	patient := app.Party("/patient")
	{
		patient.Post("/create", controller.PatientAdd)
		patient.Post("/list", controller.PatientList)
		patient.Post("/getById", controller.PatientGetByID)
	}

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
