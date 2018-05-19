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
		clinic.Post("/getByID", controller.ClinicGetByID)
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

	// visitType := app.Party("/visitType", crs).AllowMethods(iris.MethodOptions)
	// {
	// 	visitType.Post("/create", controller.VisitTypeCreate)
	// 	visitType.Post("/list", controller.VisitTypeList)
	// }

	doctorVisitScheduleMode := app.Party("/doctorVisitScheduleMode", crs).AllowMethods(iris.MethodOptions)
	{
		doctorVisitScheduleMode.Post("/create", controller.DoctorVisitScheduleModeAdd)
	}

	doctorVisitSchedule := app.Party("/doctorVisitSchedule", crs).AllowMethods(iris.MethodOptions)
	{
		doctorVisitSchedule.Post("/create", controller.DoctorVistScheduleCreate)
		doctorVisitSchedule.Post("/list", controller.DoctorVistScheduleList)
		doctorVisitSchedule.Post("/departments", controller.SchelueDepartments)
		doctorVisitSchedule.Post("/doctors", controller.SchelueDoctors)
		doctorVisitSchedule.Post("/DoctorsWithSchedule", controller.DoctorsWithSchedule)
		doctorVisitSchedule.Post("/CopyScheduleByDate", controller.CopyScheduleByDate)
		doctorVisitSchedule.Post("/OpenScheduleByDate", controller.OpenScheduleByDate)
		doctorVisitSchedule.Post("/CreateOneSchedule", controller.CreateOneSchedule)
		doctorVisitSchedule.Post("/DeleteOneUnOpenScheduleByID", controller.DeleteOneUnOpenScheduleByID)
		doctorVisitSchedule.Post("/StopScheduleByID", controller.StopScheduleByID)
	}

	patient := app.Party("/patient", crs).AllowMethods(iris.MethodOptions)
	{
		patient.Post("/create", controller.PatientAdd)
		patient.Post("/list", controller.PatientList)
		patient.Post("/getById", controller.PatientGetByID)
		patient.Post("/update", controller.PatientUpdate)
		patient.Post("/getByCertNo", controller.PatientGetByCertNo)
		patient.Post("/getByKeyword", controller.PatientsGetByKeyword)
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
		triage.Post("/reception", controller.TriageReception)
		triage.Post("/complete", controller.TriageComplete)
		triage.Post("/AppointmentsByDate", controller.AppointmentsByDate)
		triage.Post("/TreatmentPatientCreate", controller.TreatmentPatientCreate)
		triage.Post("/TreatmentPatientGet", controller.TreatmentPatientGet)
		triage.Post("/LaboratoryPatientCreate", controller.LaboratoryPatientCreate)
		triage.Post("/LaboratoryPatientGet", controller.LaboratoryPatientGet)
		triage.Post("/PrescriptionWesternPatientCreate", controller.PrescriptionWesternPatientCreate)
		triage.Post("/PrescriptionWesternPatientGet", controller.PrescriptionWesternPatientGet)
		triage.Post("/PrescriptionWesternPatientList", controller.PrescriptionWesternPatientList)
		triage.Post("/PrescriptionChinesePatientCreate", controller.PrescriptionChinesePatientCreate)
		triage.Post("/PrescriptionChinesePatientGet", controller.PrescriptionChinesePatientGet)
		triage.Post("/PrescriptionChinesePatientList", controller.PrescriptionChinesePatientList)
		triage.Post("/ExaminationPatientCreate", controller.ExaminationPatientCreate)
		triage.Post("/ExaminationPatientGet", controller.ExaminationPatientGet)
		triage.Post("/OtherCostPatientCreate", controller.OtherCostPatientCreate)
		triage.Post("/OtherCostPatientGet", controller.OtherCostPatientGet)
		triage.Post("/MaterialPatientCreate", controller.MaterialPatientCreate)
		triage.Post("/MaterialPatientGet", controller.MaterialPatientGet)
		triage.Post("/ReceiveRecord", controller.ReceiveRecord)
	}

	diagnosisTreatment := app.Party("/diagnosisTreatment", crs).AllowMethods(iris.MethodOptions)
	{
		diagnosisTreatment.Post("/create", controller.DiagnosisTreatmentCreate)
		diagnosisTreatment.Post("/update", controller.DiagnosisTreatmentUpdate)
		diagnosisTreatment.Post("/onOff", controller.DiagnosisTreatmentOnOff)
		diagnosisTreatment.Post("/list", controller.DiagnosisTreatmentList)
		diagnosisTreatment.Post("/detail", controller.DiagnosisTreatmentDetail)
	}

	charge := app.Party("/charge", crs).AllowMethods(iris.MethodOptions)
	{
		charge.Post("/type/init", controller.ChargeTypeInit)
		charge.Post("/type/create", controller.ChargeTypeCreate)
		// 创建代缴费项目
		charge.Post("/unPay/create", controller.ChargeUnPayCreate)
		charge.Post("/unPay/delete", controller.ChargeUnPayDelete)
		charge.Post("/unPay/list", controller.ChargeUnPayList)

		charge.Post("/pay", controller.ChargePay)
		charge.Post("/paid/list", controller.ChargePaidList)

		charge.Post("/traigePatient/unpay", controller.GetUnChargeTraigePatients)

	}

	//挂账
	onCredit := app.Party("/onCredit", crs).AllowMethods(iris.MethodOptions)
	{
		onCredit.Post("/create", controller.OnCreditCreate)
		onCredit.Post("/registtion/list", controller.OnCreditRegisttionList)
		onCredit.Post("/registtion/detail", controller.OnCreditRegisttionDetail)
		onCredit.Post("/repay", controller.OnCreditRepay)
	}

	appointment := app.Party("/appointment", crs).AllowMethods(iris.MethodOptions)
	{
		appointment.Post("/create", controller.AppointmentCreate)
	}

	drug := app.Party("/drug", crs).AllowMethods(iris.MethodOptions)
	{
		drug.Post("/create", controller.DrugAdd)
		drug.Post("/update", controller.DrugUpdate)
		drug.Post("/list", controller.DrugList)
		drug.Post("/detail", controller.DrugDetail)
		drug.Post("/search", controller.DrugSearch)
		drug.Post("/batchSetting", controller.BatchSetting)
		drug.Post("/instock", controller.DrugInstock)
		drug.Post("/instockRecord", controller.DrugInstockRecord)
		drug.Post("/instockRecordDetail", controller.DrugInstockRecordDetail)
		drug.Post("/instockUpdate", controller.DrugInstockUpdate)
		drug.Post("/instockCheck", controller.DrugInstockCheck)
		drug.Post("/instockDelete", controller.DrugInstockRecordDelete)
		drug.Post("/outstock", controller.DrugOutstock)
		drug.Post("/outstockRecord", controller.DrugOutstockRecord)
		drug.Post("/outstockRecordDetail", controller.DrugOutstockRecordDetail)
		drug.Post("/outstockUpdate", controller.DrugOutstockUpdate)
		drug.Post("/outstockCheck", controller.DrugOutstockCheck)
		drug.Post("/outstockDelete", controller.DrugOutstockRecordDelete)
		drug.Post("/PrescriptionWesternPatientModelCreate", controller.PrescriptionWesternPatientModelCreate)
		drug.Post("/PrescriptionWesternPatientModelList", controller.PrescriptionWesternPatientModelList)
		drug.Post("/PrescriptionWesternPatientModelDetail", controller.PrescriptionWesternPatientModelDetail)
		drug.Post("/PrescriptionWesternPatientModelUpdate", controller.PrescriptionWesternPatientModelUpdate)
		drug.Post("/PrescriptionChinesePatientModelCreate", controller.PrescriptionChinesePatientModelCreate)
		drug.Post("/PrescriptionChinesePatientModelList", controller.PrescriptionChinesePatientModelList)
		drug.Post("/PrescriptionChinesePatientModelDetail", controller.PrescriptionChinesePatientModelDetail)
		drug.Post("/PrescriptionChinesePatientModelUpdate", controller.PrescriptionChinesePatientModelUpdate)
	}

	role := app.Party("/role", crs).AllowMethods(iris.MethodOptions)
	{
		role.Post("/create", controller.RoleCreate)
		role.Post("/update", controller.RoleUpdate)
		role.Post("/listByClinicID", controller.RoleList)
		role.Post("/roleDetail", controller.RoleDetail)
	}

	business := app.Party("/business", crs).AllowMethods(iris.MethodOptions)
	{
		business.Post("/menubar/create", controller.MenubarCreate)
		business.Post("/menubar/list", controller.MenubarList)
		business.Post("/clinic/assign", controller.BusinessAssign)
		business.Post("/clinic/menubar", controller.MenuGetByClinicID)
		business.Post("/admin/create", controller.AdminCreate)
		business.Post("/admin/list", controller.AdminList)
		business.Post("/admin/update", controller.AdminUpdate)
		business.Post("/admin/getByID", controller.AdminGetByID)
	}

	diagnosis := app.Party("/diagnosis", crs).AllowMethods(iris.MethodOptions)
	{
		diagnosis.Post("/create", controller.DiagnosisCreate)
	}

	medicalRecord := app.Party("/medicalRecord", crs).AllowMethods(iris.MethodOptions)
	{
		medicalRecord.Post("/upsert", controller.MedicalRecordCreate)
		medicalRecord.Post("/findByTriageId", controller.MedicalRecordFindByTriageID)
		medicalRecord.Post("/model/create", controller.MedicalRecordModelCreate)
		medicalRecord.Post("/listByPid", controller.MedicalRecordListByPID)
		medicalRecord.Post("/model/list", controller.MedicalRecordModelList)
	}

	examination := app.Party("/examination", crs).AllowMethods(iris.MethodOptions)
	{
		examination.Post("/create", controller.ExaminationCreate)
		examination.Post("/update", controller.ExaminationUpdate)
		examination.Post("/onOff", controller.ExaminationOnOff)
		examination.Post("/list", controller.ExaminationList)
		examination.Post("/detail", controller.ExaminationDetail)
		examination.Post("/ExaminationPatientModelCreate", controller.ExaminationPatientModelCreate)
		examination.Post("/ExaminationPatientModelList", controller.ExaminationPatientModelList)
		examination.Post("/ExaminationPatientModelDetail", controller.ExaminationPatientModelDetail)
		examination.Post("/ExaminationPatientModelUpdate", controller.ExaminationPatientModelUpdate)
	}

	treatment := app.Party("/treatment", crs).AllowMethods(iris.MethodOptions)
	{
		treatment.Post("/create", controller.TreatmentCreate)
		treatment.Post("/update", controller.TreatmentUpdate)
		treatment.Post("/onOff", controller.TreatmentOnOff)
		treatment.Post("/list", controller.TreatmentList)
		treatment.Post("/detail", controller.TreatmentDetail)
		treatment.Post("/TreatmentPatientModelCreate", controller.TreatmentPatientModelCreate)
		treatment.Post("/TreatmentPatientModelList", controller.TreatmentPatientModelList)
		treatment.Post("/TreatmentPatientModelDetail", controller.TreatmentPatientModelDetail)
		treatment.Post("/TreatmentPatientModelUpdate", controller.TreatmentPatientModelUpdate)
	}

	otherCost := app.Party("/otherCost", crs).AllowMethods(iris.MethodOptions)
	{
		otherCost.Post("/create", controller.OtherCostCreate)
		otherCost.Post("/update", controller.OtherCostUpdate)
		otherCost.Post("/onOff", controller.OtherCostOnOff)
		otherCost.Post("/list", controller.OtherCostList)
		otherCost.Post("/detail", controller.OtherCostDetail)
	}

	material := app.Party("/material", crs).AllowMethods(iris.MethodOptions)
	{
		material.Post("/create", controller.MaterialCreate)
		material.Post("/update", controller.MaterialUpdate)
		material.Post("/onOff", controller.MaterialOnOff)
		material.Post("/list", controller.MaterialList)
		material.Post("/detail", controller.MaterialDetail)
		material.Post("/instock", controller.MaterialInstock)
		material.Post("/instockRecord", controller.MaterialInstockRecord)
		material.Post("/instockRecordDetail", controller.MaterialInstockRecordDetail)
		material.Post("/instockUpdate", controller.MaterialInstockUpdate)
		material.Post("/instockCheck", controller.MaterialInstockCheck)
		material.Post("/instockDelete", controller.MaterialInstockRecordDelete)
		material.Post("/outstock", controller.MaterialOutstock)
		material.Post("/outstockRecord", controller.MaterialOutstockRecord)
		material.Post("/outstockRecordDetail", controller.MaterialOutstockRecordDetail)
		material.Post("/outstockUpdate", controller.MaterialOutstockUpdate)
		material.Post("/outstockCheck", controller.MaterialOutstockCheck)
		material.Post("/outstockDelete", controller.MaterialOutstockRecordDelete)
	}

	laboratory := app.Party("/laboratory", crs).AllowMethods(iris.MethodOptions)
	{
		laboratory.Post("/create", controller.LaboratoryCreate)
		laboratory.Post("/list", controller.LaboratoryList)
		laboratory.Post("/detail", controller.LaboratoryDetail)
		laboratory.Post("/update", controller.LaboratoryUpdate)
		laboratory.Post("/association", controller.LaboratoryAssociation)
		laboratory.Post("/associationList", controller.AssociationList)
		laboratory.Post("/item/create", controller.LaboratoryItemCreate)
		laboratory.Post("/item/detail", controller.LaboratoryItemDetail)
		laboratory.Post("/item/update", controller.LaboratoryItemUpdate)
		laboratory.Post("/item/onOff", controller.LaboratoryItemStatus)
		laboratory.Post("/item/list", controller.LaboratoryItemList)
		laboratory.Post("/item/searchByName", controller.LaboratoryItemSearch)
		laboratory.Post("/LaboratoryPatientModelCreate", controller.LaboratoryPatientModelCreate)
		laboratory.Post("/LaboratoryPatientModelList", controller.LaboratoryPatientModelList)
		laboratory.Post("/LaboratoryPatientModelDetail", controller.LaboratoryPatientModelDetail)
		laboratory.Post("/LaboratoryPatientModelUpdate", controller.LaboratoryPatientModelUpdate)
	}

	dictionaries := app.Party("/dictionaries", crs).AllowMethods(iris.MethodOptions)
	{
		dictionaries.Post("/DoseUnitList", controller.DoseUnitList)
		dictionaries.Post("/DoseFormList", controller.DoseFormList)
		dictionaries.Post("/DrugClassList", controller.DrugClassList)
		dictionaries.Post("/DrugTypeList", controller.DrugTypeList)
		dictionaries.Post("/DrugPrintList", controller.DrugPrintList)
		dictionaries.Post("/ExaminationOrganList", controller.ExaminationOrganList)
		dictionaries.Post("/FrequencyList", controller.FrequencyList)
		dictionaries.Post("/RouteAdministrationList", controller.RouteAdministrationList)
		dictionaries.Post("/LaboratorySampleList", controller.LaboratorySampleList)
		dictionaries.Post("/CuvetteColorList", controller.CuvetteColorList)
	}

	dataImport := app.Party("/dataImport", crs).AllowMethods(iris.MethodOptions)
	{
		dataImport.Post("/ImportLaboratory", controller.ImportLaboratory)
		dataImport.Post("/ImportLaboratoryItem", controller.ImportLaboratoryItem)
	}

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
