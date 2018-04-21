package main

import (
	"clinicSystemGo/controller"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

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

	clinic := app.Party("/clinic", crs).AllowMethods(iris.MethodOptions)
	{
		clinic.Post("/add", controller.ClinicAdd)
		clinic.Post("/list", controller.ClinicList)
		clinic.Post("/update", controller.ClinicUpdate)
		// clinic.Post("/update", controller.ClinicUpdte)
	}

	department := app.Party("/department", crs).AllowMethods(iris.MethodOptions)
	{
		department.Post("/create", controller.DepartmentCreate)
		department.Post("/list", controller.DepartmentList)
		department.Post("/delete", controller.DepartmentDelete)
		department.Post("/update", controller.DepartmentUpdate)
	}

	personnel := app.Party("/personnel", crs).AllowMethods(iris.MethodOptions)
	{
		personnel.Post("/login", controller.PersonnelLogin)
		personnel.Post("/create", controller.PersonnelCreate)
		personnel.Post("/getById", controller.PersonnelGetByID)
		personnel.Post("/list", controller.PersonnelList)
		personnel.Post("/update", controller.PersonnelUpdate)
	}

	visitType := app.Party("/visitType", crs).AllowMethods(iris.MethodOptions)
	{
		visitType.Post("/create", controller.VisitTypeCreate)
		visitType.Post("/list", controller.VisitTypeList)
	}

	doctorVisitScheduleMode := app.Party("/doctorVisitScheduleMode", crs).AllowMethods(iris.MethodOptions)
	{
		doctorVisitScheduleMode.Post("/create", controller.DoctorVisitScheduleModeAdd)
	}

	doctorVisitSchedule := app.Party("/doctorVisitSchedule", crs).AllowMethods(iris.MethodOptions)
	{
		doctorVisitSchedule.Post("/create", controller.DoctorVistScheduleCreate)
	}

	patient := app.Party("/patient", crs).AllowMethods(iris.MethodOptions)
	{
		patient.Post("/create", controller.PatientAdd)
		patient.Post("/list", controller.PatientList)
		patient.Post("/getById", controller.PatientGetByID)
		patient.Post("/update", controller.PatientUpdate)
	}

	triage := app.Party("/triage", crs).AllowMethods(iris.MethodOptions)
	{
		triage.Post("/register", controller.TriageRegister)
		triage.Post("/patientlist", controller.TriagePatientList)
		triage.Post("/getById", controller.PatientGetByID)
		triage.Post("/personnelList", controller.TriagePersonnelList)
		triage.Post("/completeBodySign", controller.TriageCompleteBodySign)
		triage.Post("/completePreMedicalRecord", controller.TriageCompletePreMedicalRecord)
		triage.Post("/completePreDiagnosis", controller.TriageCompletePreDiagnosis)
		triage.Post("/chooseDoctor", controller.PersonnelChoose)
	}

	chargeProject := app.Party("/chargeProject", crs).AllowMethods(iris.MethodOptions)
	{
		chargeProject.Post("/type/init", controller.ChargeTypeInit)
		chargeProject.Post("/type/create", controller.ChargeTypeCreate)

		chargeProject.Post("/treatment/create", controller.ChargeProjectTreatmentCreate)
	}

	chargeUnPay := app.Party("/chargeUnPay", crs).AllowMethods(iris.MethodOptions)
	{
		// 创建代缴费项目
		chargeUnPay.Post("/create", controller.ChargeUnPayCreate)
		chargeUnPay.Post("/delete", controller.ChargeUnPayDelete)
		chargeUnPay.Post("/list", controller.ChargeUnPayList)
	}

	appointment := app.Party("/appointment", crs).AllowMethods(iris.MethodOptions)
	{
		appointment.Post("/create", controller.AppointmentCreate)
		appointment.Post("/list", controller.AppointmentList)
	}

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
