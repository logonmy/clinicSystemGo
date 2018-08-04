package main

import (
	"clinicSystemGo/controller"
	"clinicSystemGo/lib/mission"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()

	// jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("0.9434990896465933"), nil
	// 	},
	// 	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// 	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// 	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	// 	SigningMethod: jwt.SigningMethodHS256,
	// })

	app.Logger().SetLevel("debug")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	// 设置静态文件
	app.StaticServe("./uploads", "/uploads")

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

	// app.Handle("POST", "/personnel/login", controller.PersonnelLogin)
	// app.Use(jwtHandler.Serve)

	//定时任务
	mission.StartMission()

	file := app.Party("/file", crs).AllowMethods(iris.MethodOptions)
	{
		file.Post("/upload", controller.FileUpload)
	}

	clinic := app.Party("/clinic", crs).AllowMethods(iris.MethodOptions)
	{
		clinic.Post("/add", controller.ClinicAdd)
		clinic.Post("/list", controller.ClinicList)
		clinic.Post("/update/status", controller.ClinicUpdateStatus)
		clinic.Post("/update", controller.ClinicUpdate)
		clinic.Post("/getByID", controller.ClinicGetByID)
		clinic.Post("/code", controller.GetClinicCode)
		clinic.Post("/test", controller.GetClinicCode)
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
		personnel.Post("/FunMenusByPersonnel", controller.FunMenusByPersonnel)
		personnel.Post("/update", controller.PersonnelUpdate)
		personnel.Post("/delete", controller.PersonnelDelete)
		personnel.Post("/PersonnelWithAuthorizationList", controller.PersonnelWithAuthorizationList)
		personnel.Post("/PersonnelAuthorizationAllocation", controller.PersonnelAuthorizationAllocation)
		personnel.Post("/PersonnelWithUsername", controller.PersonnelWithUsername)
		personnel.Post("/UpdatePersonnelStatus", controller.UpdatePersonnelStatus)
		personnel.Post("/UpdatePersonnelUsername", controller.UpdatePersonnelUsername)
		personnel.Post("/PersonnelRoles", controller.PersonnelRoles)
		personnel.Post("/PersonnelDepartmentList", controller.PersonnelDepartmentList)
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
		doctorVisitSchedule.Post("/RemoveScheduleByID", controller.RemoveScheduleByID)
	}

	patient := app.Party("/patient", crs).AllowMethods(iris.MethodOptions)
	{
		patient.Post("/create", controller.PatientAdd)
		patient.Post("/list", controller.PatientList)
		patient.Post("/getById", controller.PatientGetByID)
		patient.Post("/update", controller.PatientUpdate)
		patient.Post("/getByCertNo", controller.PatientGetByCertNo)
		patient.Post("/getByKeyword", controller.PatientsGetByKeyword)
		patient.Post("/MemberPateintList", controller.MemberPateintList)
		patient.Post("/PersonalMedicalRecordUpsert", controller.PersonalMedicalRecordUpsert)
		patient.Post("/PersonalMedicalRecord", controller.PersonalMedicalRecord)
		patient.Post("/GetLastBodySign", controller.GetLastBodySign)
		patient.Post("/UpsertPatientHeight", controller.UpsertPatientHeight)
		patient.Post("/UpsertPatientWeight", controller.UpsertPatientWeight)
		patient.Post("/UpsertPatientBmi", controller.UpsertPatientBmi)
		patient.Post("/UpsertPatientBloodType", controller.UpsertPatientBloodType)
		patient.Post("/UpsertPatientRhBloodType", controller.UpsertPatientRhBloodType)
		patient.Post("/UpsertPatientTemperature", controller.UpsertPatientTemperature)
		patient.Post("/UpsertPatientBreathe", controller.UpsertPatientBreathe)
		patient.Post("/UpsertPatientPulse", controller.UpsertPatientPulse)
		patient.Post("/UpsertPatientBloodPressure", controller.UpsertPatientBloodPressure)
		patient.Post("/UpsertPatientVision", controller.UpsertPatientVision)
		patient.Post("/UpsertPatientBloodSugar", controller.UpsertPatientBloodSugar)
		patient.Post("/UpsertPatientOxygenSaturation", controller.UpsertPatientOxygenSaturation)
		patient.Post("/PatientHeightList", controller.PatientHeightList)
		patient.Post("/PatientWeightList", controller.PatientWeightList)
		patient.Post("/PatientBmiList", controller.PatientBmiList)
		patient.Post("/PatientBloodTypeList", controller.PatientBloodTypeList)
		patient.Post("/PatientRhBloodTypeList", controller.PatientRhBloodTypeList)
		patient.Post("/PatientTemperatureList", controller.PatientTemperatureList)
		patient.Post("/PatientBreatheList", controller.PatientBreatheList)
		patient.Post("/PatientPulseList", controller.PatientPulseList)
		patient.Post("/PatientBloodPressureList", controller.PatientBloodPressureList)
		patient.Post("/PatientVisionList", controller.PatientVisionList)
		patient.Post("/PatientBloodSugarList", controller.PatientBloodSugarList)
		patient.Post("/PatientOxygenSaturationList", controller.PatientOxygenSaturationList)
		patient.Post("/PatientBloodSugarListByDate", controller.PatientBloodSugarListByDate)
	}

	triage := app.Party("/triage", crs).AllowMethods(iris.MethodOptions)
	{
		triage.Post("/register", controller.TriageRegister)
		triage.Post("/TriagePatientDetail", controller.TriagePatientDetail)
		triage.Post("/patientlist", controller.TriagePatientList)
		triage.Post("/RecptionPatientList", controller.RecptionPatientList)
		triage.Post("/getById", controller.PatientGetByID)
		triage.Post("/personnelList", controller.TriagePersonnelList)
		triage.Post("/completeBodySign", controller.TriageCompleteBodySign)
		triage.Post("/completePreMedicalRecord", controller.TriageCompletePreMedicalRecord)
		triage.Post("/completePreDiagnosis", controller.TriageCompletePreDiagnosis)
		triage.Post("/GetHealthRecord", controller.GetHealthRecord)
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
		triage.Post("/PrescriptionChinesePatientDelete", controller.PrescriptionChinesePatientDelete)
		triage.Post("/PrescriptionChinesePatientGet", controller.PrescriptionChinesePatientGet)
		triage.Post("/PrescriptionChinesePatientList", controller.PrescriptionChinesePatientList)
		triage.Post("/ExaminationPatientCreate", controller.ExaminationPatientCreate)
		triage.Post("/ExaminationPatientGet", controller.ExaminationPatientGet)
		triage.Post("/OtherCostPatientCreate", controller.OtherCostPatientCreate)
		triage.Post("/OtherCostPatientGet", controller.OtherCostPatientGet)
		triage.Post("/MaterialPatientCreate", controller.MaterialPatientCreate)
		triage.Post("/MaterialPatientGet", controller.MaterialPatientGet)
		triage.Post("/ReceiveRecord", controller.ReceiveRecord)
		triage.Post("/DiagnosisTreatmentPatientCreate", controller.DiagnosisTreatmentPatientCreate)
		triage.Post("/TriagePatientVisitDetail", controller.TriagePatientVisitDetail)
		triage.Post("/TriagePatientRecord", controller.TriagePatientRecord)
		triage.Post("/QuickReception", controller.QuickReception)
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
		// charge.Post("/unPay/create", controller.ChargeUnPayCreate)
		// charge.Post("/unPay/delete", controller.ChargeUnPayDelete)
		// 创建支付订单
		charge.Post("/payment/create", controller.ChargePaymentCreate)
		// 获取支付状态
		charge.Post("/payment/query", controller.ChargePaymentQuery)

		// 支付到账通知
		charge.Post("/payment/notice", controller.ChargeNotice)
		charge.Post("/paid/list", controller.ChargePaidList)
		// 查询有待缴费的就诊记录
		charge.Post("/traigePatient/unpay", controller.GetUnChargeTraigePatients)
		// 查询待缴费的项目列表
		charge.Post("/unPay/list", controller.ChargeUnPayList)
		charge.Post("/paid/list", controller.ChargePaidList)
		// 查询已缴费的就诊记录
		charge.Post("/traigePatient/paid", controller.GetPaidTraigePatients)
		// 获取交易流水日报表
		charge.Post("/business/transaction", controller.BusinessTransaction)
		// 获取分析类报表
		charge.Post("/business/transaction/analysis", controller.BusinessTransactionAnalysis)
		// 获取交易流水月报表
		charge.Post("/business/transaction/month", controller.BusinessTransactionMonth)
		// 获取交易详情
		charge.Post("/business/transaction/detail", controller.BusinessTransactionDetail)
		// 获取挂账交易详情
		charge.Post("/business/transaction/credit", controller.BusinessTransactionCredit)
		// 患者支付列表
		charge.Post("/PatientChargeList", controller.PatientChargeList)

	}

	//挂账
	onCredit := app.Party("/onCredit", crs).AllowMethods(iris.MethodOptions)
	{
		onCredit.Post("/traigePatient/list", controller.OnCreditTraigePatient)
		onCredit.Post("/list", controller.OnCreditList)
		// onCredit.Post("/create", controller.OnCreditCreate)
		// onCredit.Post("/registtion/list", controller.OnCreditRegisttionList)
		// onCredit.Post("/registtion/detail", controller.OnCreditRegisttionDetail)
		onCredit.Post("/repay", controller.OnCreditRepay)
	}

	appointment := app.Party("/appointment", crs).AllowMethods(iris.MethodOptions)
	{
		appointment.Post("/create", controller.AppointmentCreate)
	}

	drug := app.Party("/clinic_drug", crs).AllowMethods(iris.MethodOptions)
	{
		drug.Post("/ClinicDrugCreate", controller.ClinicDrugCreate)
		drug.Post("/ClinicDrugUpdate", controller.ClinicDrugUpdate)
		drug.Post("/ClinicDrugOnOff", controller.ClinicDrugOnOff)
		drug.Post("/ClinicDrugList", controller.ClinicDrugList)
		drug.Post("/ClinicDrugStock", controller.ClinicDrugListWithStock)
		drug.Post("/ClinicDrugDetail", controller.ClinicDrugDetail)
		drug.Post("/ClinicDrugBatchSetting", controller.ClinicDrugBatchSetting)
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
		drug.Post("/DrugStockList", controller.DrugStockList)
		drug.Post("/PrescriptionWesternPatientModelCreate", controller.PrescriptionWesternPatientModelCreate)
		drug.Post("/PrescriptionWesternPatientModelList", controller.PrescriptionWesternPatientModelList)
		drug.Post("/PrescriptionWesternPersonalPatientModelList", controller.PrescriptionWesternPersonalPatientModelList)
		drug.Post("/PrescriptionWesternPatientModelDetail", controller.PrescriptionWesternPatientModelDetail)
		drug.Post("/PrescriptionWesternPatientModelUpdate", controller.PrescriptionWesternPatientModelUpdate)
		drug.Post("/PrescriptionWesternPatientModelDelete", controller.PrescriptionWesternPatientModelDelete)
		drug.Post("/PrescriptionChinesePatientModelCreate", controller.PrescriptionChinesePatientModelCreate)
		drug.Post("/PrescriptionChinesePatientModelList", controller.PrescriptionChinesePatientModelList)
		drug.Post("/PrescriptionChinesePersonalPatientModelList", controller.PrescriptionChinesePersonalPatientModelList)
		drug.Post("/PrescriptionChinesePatientModelDetail", controller.PrescriptionChinesePatientModelDetail)
		drug.Post("/PrescriptionChinesePatientModelUpdate", controller.PrescriptionChinesePatientModelUpdate)
		drug.Post("/PrescriptionChinesePatientModelDelete", controller.PrescriptionChinesePatientModelDelete)
	}

	role := app.Party("/role", crs).AllowMethods(iris.MethodOptions)
	{
		role.Post("/create", controller.RoleCreate)
		role.Post("/update", controller.RoleUpdate)
		role.Post("/delete", controller.RoleDelete)
		role.Post("/listByClinicID", controller.RoleList)
		role.Post("/roleDetail", controller.RoleDetail)
		role.Post("/RoleFunctionUnset", controller.RoleFunctionUnset)
		role.Post("/RoleAllocation", controller.RoleAllocation)
		role.Post("/PersonnelsByRole", controller.PersonnelsByRole)
	}

	business := app.Party("/business", crs).AllowMethods(iris.MethodOptions)
	{
		business.Post("/menubar/create", controller.MenubarCreate)
		business.Post("/menubar/list", controller.MenubarList)
		business.Post("/menubar/list/clinicUnset", controller.MenubarListByClinicID)
		business.Post("/clinic/assign", controller.BusinessAssign)
		business.Post("/clinic/menubar", controller.MenuGetByClinicID)
	}

	admin := app.Party("/admin", crs).AllowMethods(iris.MethodOptions)
	{
		admin.Post("/login", controller.AdminLogin)
		admin.Post("/create", controller.AdminCreate)
		admin.Post("/list", controller.AdminList)
		admin.Post("/update", controller.AdminUpdate)
		admin.Post("/getByID", controller.AdminGetByID)
		admin.Post("/onOff", controller.AdminOnOff)
	}

	diagnosis := app.Party("/diagnosis", crs).AllowMethods(iris.MethodOptions)
	{
		diagnosis.Post("/create", controller.DiagnosisCreate)
	}

	medicalRecord := app.Party("/medicalRecord", crs).AllowMethods(iris.MethodOptions)
	{
		medicalRecord.Post("/upsert", controller.MedicalRecordCreate)
		medicalRecord.Post("/renew", controller.MedicalRecordRenew)
		medicalRecord.Post("/MedicalRecordRenewUpdate", controller.MedicalRecordRenewUpdate)
		medicalRecord.Post("/MedicalRecordRenewDelete", controller.MedicalRecordRenewDelete)
		medicalRecord.Post("/findByTriageId", controller.MedicalRecordFindByTriageID)
		medicalRecord.Post("/model/create", controller.MedicalRecordModelCreate)
		medicalRecord.Post("/model/update", controller.MedicalRecordModelUpdate)
		medicalRecord.Post("/model/delete", controller.MedicalRecordModelDelete)
		medicalRecord.Post("/listByPid", controller.MedicalRecordListByPID)
		medicalRecord.Post("/model/list", controller.MedicalRecordModelList)
		medicalRecord.Post("/model/listByOperation", controller.MedicalRecordModelListByOperation)
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
		examination.Post("/ExaminationPersonalPatientModelList", controller.ExaminationPersonalPatientModelList)
		examination.Post("/ExaminationPatientModelDetail", controller.ExaminationPatientModelDetail)
		examination.Post("/ExaminationPatientModelUpdate", controller.ExaminationPatientModelUpdate)
		examination.Post("/ExaminationPatientModelDelete", controller.ExaminationPatientModelDelete)
		examination.Post("/ExaminationReportModelCreate", controller.ExaminationReportModelCreate)
		examination.Post("/ExaminationReportModelList", controller.ExaminationReportModelList)
		examination.Post("/ExaminationReportModelDetail", controller.ExaminationReportModelDetail)
		examination.Post("/ExaminationReportModelUpdate", controller.ExaminationReportModelUpdate)
		examination.Post("/ExaminationReportModelDelete", controller.ExaminationReportModelDelete)
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
		treatment.Post("/TreatmentPersonalPatientModelList", controller.TreatmentPersonalPatientModelList)
		treatment.Post("/TreatmentPatientModelDetail", controller.TreatmentPatientModelDetail)
		treatment.Post("/TreatmentPatientModelUpdate", controller.TreatmentPatientModelUpdate)
		treatment.Post("/TreatmentPatientModelDelete", controller.TreatmentPatientModelDelete)
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
		material.Post("/MaterialStockList", controller.MaterialStockList)
	}

	laboratory := app.Party("/laboratory", crs).AllowMethods(iris.MethodOptions)
	{
		laboratory.Post("/create", controller.LaboratoryCreate)
		laboratory.Post("/list", controller.LaboratoryList)
		laboratory.Post("/detail", controller.LaboratoryDetail)
		laboratory.Post("/update", controller.LaboratoryUpdate)
		laboratory.Post("/onOff", controller.LaboratoryOnOff)
		laboratory.Post("/association", controller.LaboratoryAssociation)
		laboratory.Post("/associationList", controller.AssociationList)
		laboratory.Post("/item/create", controller.LaboratoryItemCreate)
		laboratory.Post("/item/detail", controller.LaboratoryItemDetail)
		laboratory.Post("/item/update", controller.LaboratoryItemUpdate)
		laboratory.Post("/item/onOff", controller.LaboratoryItemStatus)
		laboratory.Post("/item/list", controller.LaboratoryItemList)
		laboratory.Post("/LaboratoryPatientModelCreate", controller.LaboratoryPatientModelCreate)
		laboratory.Post("/LaboratoryPatientModelList", controller.LaboratoryPatientModelList)
		laboratory.Post("/LaboratoryPersonalPatientModelList", controller.LaboratoryPersonalPatientModelList)
		laboratory.Post("/LaboratoryPatientModelDetail", controller.LaboratoryPatientModelDetail)
		laboratory.Post("/LaboratoryPatientModelUpdate", controller.LaboratoryPatientModelUpdate)
		laboratory.Post("/LaboratoryPatientModelDelete", controller.LaboratoryPatientModelDelete)
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
		dictionaries.Post("/ManuFactoryList", controller.ManuFactoryList)
		dictionaries.Post("/Laboratorys", controller.Laboratorys)
		dictionaries.Post("/Examinations", controller.Examinations)
		dictionaries.Post("/LaboratoryItems", controller.LaboratoryItems)
		dictionaries.Post("/Drugs", controller.Drugs)
		dictionaries.Post("/SupplierList", controller.SupplierList)
		dictionaries.Post("/InstockWayList", controller.InstockWayList)
		dictionaries.Post("/OutstockWayList", controller.OutstockWayList)
		dictionaries.Post("/DiagnosisList", controller.DiagnosisList)
	}

	dataImport := app.Party("/dataImport", crs).AllowMethods(iris.MethodOptions)
	{
		dataImport.Post("/ImportFrequency", controller.ImportFrequency)
		dataImport.Post("/ImportDoseUnit", controller.ImportDoseUnit)
		dataImport.Post("/ImportDoseForm", controller.ImportDoseForm)
		dataImport.Post("/ImportManuFactory", controller.ImportManuFactory)
		dataImport.Post("/ImportrRouteAdministration", controller.ImportrRouteAdministration)
		dataImport.Post("/ImportrLaboratorySample", controller.ImportrLaboratorySample)
		dataImport.Post("/ImportLaboratory", controller.ImportLaboratory)
		dataImport.Post("/ImportLaboratoryItem", controller.ImportLaboratoryItem)
		dataImport.Post("/ImportExamination", controller.ImportExamination)
		// dataImport.Post("/ImportDrugType", controller.ImportDrugType)
		dataImport.Post("/ImportDrug", controller.ImportDrug)
		dataImport.Post("/ImportDrugClass", controller.ImportDrugClass)
		dataImport.Post("/ImportDiagnosis", controller.ImportDiagnosis)
	}

	chiefComplaint := app.Party("/chiefComplaint", crs).AllowMethods(iris.MethodOptions)
	{
		chiefComplaint.Post("/list", controller.ChiefComplaintList)
	}

	drugDelivery := app.Party("/drugDelivery", crs).AllowMethods(iris.MethodOptions)
	{
		drugDelivery.Post("/traige/waiting/list", controller.DrugDeliveryWaiting)
		drugDelivery.Post("/traige/issued/list", controller.DrugDeliveryIssued)
		drugDelivery.Post("/traige/refund/list", controller.DrugDeliveryRefund)
		drugDelivery.Post("/list", controller.DrugDeliveryList)
		drugDelivery.Post("/record/create", controller.DrugDeliveryRecordCreate)
		drugDelivery.Post("/record/refund", controller.DrugDeliveryRecordRefund)
		drugDelivery.Post("/record/list", controller.DrugDeliveryRecordList)
		drugDelivery.Post("/record/detail", controller.DrugDeliveryRecordDetail)
		drugDelivery.Post("/record/refund/list", controller.DrugDeliveryRecordRefundList)
		drugDelivery.Post("/record/refund/detail", controller.DrugDeliveryRecordRefundDetail)
		drugDelivery.Post("/DrugDeliveryStockList", controller.DrugDeliveryStockList)
	}

	examinationTriage := app.Party("/examinationTriage", crs).AllowMethods(iris.MethodOptions)
	{
		examinationTriage.Post("/ExaminationTriageList", controller.ExaminationTriageList)
		examinationTriage.Post("/ExaminationTriageWaiting", controller.ExaminationTriageWaiting)
		examinationTriage.Post("/ExaminationTriageChecked", controller.ExaminationTriageChecked)
		examinationTriage.Post("/ExaminationTriageChecking", controller.ExaminationTriageChecking)
		examinationTriage.Post("/ExaminationTriageRecordCreate", controller.ExaminationTriageRecordCreate)
		examinationTriage.Post("/ExaminationTriageRecordList", controller.ExaminationTriageRecordList)
		examinationTriage.Post("/ExaminationTriageUpdate", controller.ExaminationTriageUpdate)
		examinationTriage.Post("/ExaminationTriagePatientRecordList", controller.ExaminationTriagePatientRecordList)
	}

	laboratoryTriage := app.Party("/laboratoryTriage", crs).AllowMethods(iris.MethodOptions)
	{
		laboratoryTriage.Post("/LaboratoryTriageList", controller.LaboratoryTriageList)
		laboratoryTriage.Post("/LaboratoryTriageDetail", controller.LaboratoryTriageDetail)
		laboratoryTriage.Post("/LaboratoryTriageWaiting", controller.LaboratoryTriageWaiting)
		laboratoryTriage.Post("/LaboratoryTriageChecked", controller.LaboratoryTriageChecked)
		laboratoryTriage.Post("/LaboratoryTriageChecking", controller.LaboratoryTriageChecking)
		laboratoryTriage.Post("/LaboratoryTriageRecordCreate", controller.LaboratoryTriageRecordCreate)
		laboratoryTriage.Post("/LaboratoryTriageRecordList", controller.LaboratoryTriageRecordList)
		laboratoryTriage.Post("/LaboratoryTriageRecordDetail", controller.LaboratoryTriageRecordDetail)
		laboratoryTriage.Post("/LaboratoryTriageUpdate", controller.LaboratoryTriageUpdate)
		laboratoryTriage.Post("/LaboratoryTriagePatientRecordList", controller.LaboratoryTriagePatientRecordList)
	}

	treatmentTriage := app.Party("/treatmentTriage", crs).AllowMethods(iris.MethodOptions)
	{
		treatmentTriage.Post("/TreatmentTriageList", controller.TreatmentTriageList)
		treatmentTriage.Post("/TreatmentTriageWaiting", controller.TreatmentTriageWaiting)
		treatmentTriage.Post("/TreatmentTriageChecked", controller.TreatmentTriageChecked)
		treatmentTriage.Post("/TreatmentTriageChecking", controller.TreatmentTriageChecking)
		treatmentTriage.Post("/TreatmentTriageRecordCreate", controller.TreatmentTriageRecordCreate)
		treatmentTriage.Post("/TreatmentTriageRecordUpdate", controller.TreatmentTriageRecordUpdate)
		treatmentTriage.Post("/TreatmentTriageRecordList", controller.TreatmentTriageRecordList)
		treatmentTriage.Post("/TreatmentTriageUpdate", controller.TreatmentTriageUpdate)
		treatmentTriage.Post("/TreatmentPatientModelDelete", controller.TreatmentPatientModelDelete)
		treatmentTriage.Post("/TreatmentTriagePatientRecordList", controller.TreatmentTriagePatientRecordList)
	}

	hcPay := app.Party("/hcPay", crs).AllowMethods(iris.MethodOptions)
	{
		hcPay.Post("/CreateHcOrder", controller.CreateHcOrderWeb)
		hcPay.Post("/QueryHcOrder", controller.QueryHcOrderWeb)
		hcPay.Post("/QueryHcOrderTest", controller.QueryHcOrderTest)
		hcPay.Post("/HcRefund", controller.HcRefundWeb)
		hcPay.Post("/QueryHcRefund", controller.QueryHcRefundWeb)
		hcPay.Post("/HcOrderClose", controller.HcOrderCloseWeb)
		hcPay.Post("/FaceToFace", controller.FaceToFaceWeb)
		hcPay.Post("/FaceToFaceCancel", controller.FaceToFaceCancelWeb)
		hcPay.Post("/DownloadBill", controller.DownloadBillWeb)
	}

	drugRetail := app.Party("/drugRetail", crs).AllowMethods(iris.MethodOptions)
	{
		drugRetail.Post("/createOrder", controller.CreateDrugRetailOrder)
		drugRetail.Post("/createPaymentOrder", controller.CreateDrugRetailPaymentOrder)
		drugRetail.Post("/payment/status", controller.DrugRetailPaymentStatus)
		drugRetail.Post("/list", controller.DrugRetailList)
		drugRetail.Post("/detail", controller.DrugRetailDetail)
		drugRetail.Post("/refund", controller.DrugRetailRefund)
	}

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
