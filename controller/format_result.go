package controller

import (
	"time"

	"github.com/jmoiron/sqlx"
)

//Menu 子菜单
type Menu struct {
	FunctionmenuID interface{} `json:"functionmenu_id"`
	MenuName       interface{} `json:"menu_name"`
	MenuURL        interface{} `json:"menu_url"`
}

//Funtionmenus 菜单
type Funtionmenus struct {
	ParentID       interface{} `json:"parent_id"`
	ParentName     interface{} `json:"parent_name"`
	ParentURL      interface{} `json:"parent_url"`
	ChildrensMenus []Menu      `json:"childrens_menus"`
}

//References 检验项参考值
type References struct {
	ReferenceSex   interface{} `json:"reference_sex"`
	ReferenceMax   interface{} `json:"reference_max"`
	ReferenceMin   interface{} `json:"reference_min"`
	ReferenceValue interface{} `json:"reference_value"`
	IsPregnancy    interface{} `json:"isPregnancy"`
	StomachStatus  interface{} `json:"stomach_status"`
}

//LaboratoryItem 检验项
type LaboratoryItem struct {
	ClinicLaboratoryItemID interface{}  `json:"clinic_laboratory_item_id"`
	Name                   interface{}  `json:"name"`
	EnName                 interface{}  `json:"en_name"`
	UnitName               interface{}  `json:"unit_name"`
	Status                 interface{}  `json:"status"`
	IsSpecial              interface{}  `json:"is_special"`
	DataType               interface{}  `json:"data_type"`
	InstrumentCode         interface{}  `json:"instrument_code"`
	IsDelivery             interface{}  `json:"is_delivery"`
	DefaultResult          interface{}  `json:"default_result"`
	References             []References `json:"references"`
}

//PrescriptionModel 处方模板
type PrescriptionModel struct {
	ModelName               interface{}             `json:"model_name"`
	PrescriptionModelID     interface{}             `json:"prescription_patient_model_id"`
	OperationName           interface{}             `json:"operation_name"`
	IsCommon                interface{}             `json:"is_common"`
	RouteAdministrationName interface{}             `json:"route_administration_name"`
	EffDay                  interface{}             `json:"eff_day"`
	Amount                  interface{}             `json:"amount"`
	FrequencyName           interface{}             `json:"frequency_name"`
	FetchAddress            interface{}             `json:"fetch_address"`
	MedicineIllustration    interface{}             `json:"medicine_illustration"`
	CreatedTime             interface{}             `json:"created_time"`
	UpdatedTime             interface{}             `json:"updated_time"`
	Items                   []PrescriptionModelItem `json:"items"`
}

//PrescriptionModelItem 处方模板item
type PrescriptionModelItem struct {
	ClinicDrugID            interface{} `json:"clinic_drug_id"`
	DrugType                interface{} `json:"type"`
	DrugName                interface{} `json:"drug_name"`
	Specification           interface{} `json:"specification"`
	StockAmount             interface{} `json:"stock_amount"`
	OnceDose                interface{} `json:"once_dose"`
	OnceDoseUnitName        interface{} `json:"once_dose_unit_name"`
	RouteAdministrationName interface{} `json:"route_administration_name"`
	FrequencyName           interface{} `json:"frequency_name"`
	EffDay                  interface{} `json:"eff_day"`
	Amount                  interface{} `json:"amount"`
	PackingUnitName         interface{} `json:"packing_unit_name"`
	FetchAddress            interface{} `json:"fetch_address"`
	Illustration            interface{} `json:"illustration"`
	SpecialIllustration     interface{} `json:"special_illustration"`
}

// FormatSQLRowsToMapArray 格式化数据库返回的数组数据
func FormatSQLRowsToMapArray(rows *sqlx.Rows) []map[string]interface{} {
	var results []map[string]interface{}
	for rows.Next() {
		m := make(map[string]interface{})
		rows.MapScan(m)
		results = append(results, m)
	}
	return results
}

// FormatSQLRowToMap 格式化数据库返回的数组数据
func FormatSQLRowToMap(row *sqlx.Row) map[string]interface{} {
	result := make(map[string]interface{})
	row.MapScan(result)
	return result
}

//FormatLaboratoryItem 格式化检验项目
func FormatLaboratoryItem(results []map[string]interface{}) []LaboratoryItem {
	var laboratoryItems []LaboratoryItem
	for _, v := range results {
		clinicLaboratoryItemID := v["clinic_laboratory_item_id"]
		name := v["name"]
		instrumentCode := v["instrument_code"]
		isDelivery := v["is_delivery"]
		enName := v["en_name"]
		unitName := v["unit_name"]
		status := v["status"]
		isSpecial := v["is_special"]
		dataType := v["data_type"]
		defaultResult := v["default_result"]
		referenceSex := v["reference_sex"]
		referenceMax := v["reference_max"]
		referenceMin := v["reference_min"]
		referenceValue := v["reference_value"]
		isPregnancy := v["is_pregnancy"]
		stomachStatus := v["stomach_status"]
		has := false
		reference := References{
			ReferenceSex:   referenceSex,
			ReferenceMax:   referenceMax,
			ReferenceMin:   referenceMin,
			ReferenceValue: referenceValue,
			IsPregnancy:    isPregnancy,
			StomachStatus:  stomachStatus,
		}
		for k, vRes := range laboratoryItems {
			vlaboratoryItemID := vRes.ClinicLaboratoryItemID
			vreferences := vRes.References
			if vlaboratoryItemID == clinicLaboratoryItemID {
				laboratoryItems[k].References = append(vreferences, reference)
				has = true
			}
		}
		if !has {
			var references []References
			references = append(references, reference)

			laboratoryItem := LaboratoryItem{
				ClinicLaboratoryItemID: clinicLaboratoryItemID,
				Name:           name,
				EnName:         enName,
				UnitName:       unitName,
				Status:         status,
				InstrumentCode: instrumentCode,
				IsDelivery:     isDelivery,
				DataType:       dataType,
				DefaultResult:  defaultResult,
				IsSpecial:      isSpecial,
				References:     references,
			}
			laboratoryItems = append(laboratoryItems, laboratoryItem)
		}
	}
	return laboratoryItems
}

// FormatFuntionmenus 格式化菜单功能项
func FormatFuntionmenus(functionMenus []map[string]interface{}) []Funtionmenus {
	var menus []Funtionmenus
	for _, v := range functionMenus {
		childenURL := v["menu_url"]
		childenName := v["menu_name"]
		functionmenuID := v["functionmenu_id"]
		parentID := v["parent_id"]
		parentURL := v["parent_url"]
		parentName := v["parent_name"]
		has := false
		children := Menu{
			FunctionmenuID: functionmenuID,
			MenuName:       childenName,
			MenuURL:        childenURL,
		}
		for k, menu := range menus {
			parentMenuID := menu.ParentID
			childrenMenus := menu.ChildrensMenus
			if parentID == parentMenuID {
				menus[k].ChildrensMenus = append(childrenMenus, children)
				has = true
			}
		}
		if !has {
			var childrens []Menu
			childrens = append(childrens, children)
			functionMenu := Funtionmenus{
				ParentID:       parentID,
				ParentName:     parentName,
				ParentURL:      parentURL,
				ChildrensMenus: childrens,
			}
			menus = append(menus, functionMenu)
		}
	}
	return menus
}

// FormatPayOrderSn 格式化单号
func FormatPayOrderSn(clinicTriagePatientID string, chargeProjectTypeID string) string {
	var orderSn string
	orderSn = time.Now().Format("20060102") + chargeProjectTypeID + clinicTriagePatientID
	return orderSn
}

// FormatPrescriptionModel 格式化处方模板
func FormatPrescriptionModel(prescriptionModel []map[string]interface{}) []PrescriptionModel {
	var models []PrescriptionModel
	for _, v := range prescriptionModel {
		modelName := v["model_name"]
		prescriptionModelID := v["prescription_patient_model_id"]
		operationName := v["operation_name"]
		isCommon := v["is_common"]
		createdTime := v["created_time"]
		updatedTime := v["updated_time"]
		inforRouteAdministrationName := v["info_route_administration_name"]
		infoEffDay := v["info_eff_day"]
		infoAmount := v["info_amount"]
		infoFrequencyName := v["info_frequency_name"]
		infoFetchAddress := v["info_fetch_address"]
		medicineIllustration := v["medicine_illustration"]

		has := false

		clinicDrugID := v["clinic_drug_id"]
		drugType := v["type"]
		drugName := v["drug_name"]
		specification := v["specification"]
		stockAmount := v["stock_amount"]
		onceDose := v["once_dose"]
		onceDoseUnitName := v["once_dose_unit_name"]
		routeAdministrationName := v["route_administration_name"]
		frequencyName := v["frequency_name"]
		effDay := v["eff_day"]
		amount := v["amount"]
		packingUnitName := v["packing_unit_name"]
		fetchAddress := v["fetch_address"]
		illustration := v["illustration"]
		specialIllustration := v["special_illustration"]

		item := PrescriptionModelItem{
			ClinicDrugID:            clinicDrugID,
			DrugType:                drugType,
			DrugName:                drugName,
			Specification:           specification,
			StockAmount:             stockAmount,
			OnceDose:                onceDose,
			OnceDoseUnitName:        onceDoseUnitName,
			RouteAdministrationName: routeAdministrationName,
			FrequencyName:           frequencyName,
			EffDay:                  effDay,
			Amount:                  amount,
			PackingUnitName:         packingUnitName,
			FetchAddress:            fetchAddress,
			Illustration:            illustration,
			SpecialIllustration:     specialIllustration,
		}
		for k, pModel := range models {
			dprescriptionModelID := pModel.PrescriptionModelID
			items := pModel.Items
			if prescriptionModelID == dprescriptionModelID {
				models[k].Items = append(items, item)
				has = true
			}
		}
		if !has {
			var items []PrescriptionModelItem
			items = append(items, item)
			pmodel := PrescriptionModel{
				ModelName:               modelName,
				PrescriptionModelID:     prescriptionModelID,
				OperationName:           operationName,
				IsCommon:                isCommon,
				CreatedTime:             createdTime,
				RouteAdministrationName: inforRouteAdministrationName,
				EffDay:                  infoEffDay,
				Amount:                  infoAmount,
				FrequencyName:           infoFrequencyName,
				FetchAddress:            infoFetchAddress,
				MedicineIllustration:    medicineIllustration,
				UpdatedTime:             updatedTime,
				Items:                   items,
			}
			models = append(models, pmodel)
		}
	}
	return models
}
