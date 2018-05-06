package controller

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

//Menu 子菜单
type Menu struct {
	FunctionmenuID string `json:"functionmenu_id"`
	MenuName       string `json:"menu_name"`
	MenuURL        string `json:"menu_url"`
}

//Funtionmenus 菜单
type Funtionmenus struct {
	ParentID       string `json:"parent_id"`
	ParentName     string `json:"parent_name"`
	ParentURL      string `json:"parent_url"`
	ChildrensMenus []Menu `json:"childrens_menus"`
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
	ClinicLaboratoryItemID int64        `json:"clinic_laboratory_item_id"`
	LaboratoryItemID       int64        `json:"laboratory_item_id"`
	Name                   string       `json:"name"`
	EnName                 interface{}  `json:"en_name"`
	Unit                   interface{}  `json:"unit"`
	Status                 bool         `json:"status"`
	IsSpecial              bool         `json:"is_special"`
	DataType               int64        `json:"data_type"`
	InstrumentCode         interface{}  `json:"instrument_code"`
	IsDelivery             interface{}  `json:"is_delivery"`
	DefaultResult          interface{}  `json:"default_result"`
	References             []References `json:"references"`
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
		laboratoryItemID := v["laboratory_item_id"]
		name := v["name"]
		instrumentCode := v["instrument_code"]
		isDelivery := v["is_delivery"]
		enName := v["en_name"]
		unit := v["unit"]
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
		for k, vRes := range laboratoryItems {
			vlaboratoryItemID := vRes.LaboratoryItemID
			vreferences := vRes.References
			if vlaboratoryItemID == laboratoryItemID.(int64) {
				reference := References{
					ReferenceSex:   referenceSex,
					ReferenceMax:   referenceMax,
					ReferenceMin:   referenceMin,
					ReferenceValue: referenceValue,
					IsPregnancy:    isPregnancy,
					StomachStatus:  stomachStatus,
				}
				laboratoryItems[k].References = append(vreferences, reference)
				has = true
			}
		}
		if !has {
			var references []References
			reference := References{
				ReferenceSex:   referenceSex,
				ReferenceMax:   referenceMax,
				ReferenceMin:   referenceMin,
				IsPregnancy:    isPregnancy,
				ReferenceValue: referenceValue,
				StomachStatus:  stomachStatus,
			}
			references = append(references, reference)

			laboratoryItem := LaboratoryItem{
				ClinicLaboratoryItemID: clinicLaboratoryItemID.(int64),
				LaboratoryItemID:       laboratoryItemID.(int64),
				Name:                   name.(string),
				EnName:                 enName,
				Unit:                   unit,
				Status:                 status.(bool),
				InstrumentCode:         instrumentCode,
				IsDelivery:             isDelivery,
				DataType:               dataType.(int64),
				DefaultResult:          defaultResult,
				IsSpecial:              isSpecial.(bool),
				References:             references,
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
		for k, menu := range menus {
			parentMenuID := menu.ParentID
			childrenMenus := menu.ChildrensMenus
			if strconv.FormatInt(parentID.(int64), 10) == parentMenuID {
				childrens := Menu{
					FunctionmenuID: strconv.FormatInt(functionmenuID.(int64), 10),
					MenuName:       childenName.(string),
					MenuURL:        childenURL.(string),
				}
				menus[k].ChildrensMenus = append(childrenMenus, childrens)
				has = true
			}
		}
		if !has {
			var childrens []Menu
			children := Menu{
				FunctionmenuID: strconv.FormatInt(functionmenuID.(int64), 10),
				MenuName:       childenName.(string),
				MenuURL:        childenURL.(string),
			}
			childrens = append(childrens, children)

			functionMenu := Funtionmenus{
				ParentID:       strconv.FormatInt(parentID.(int64), 10),
				ParentName:     parentName.(string),
				ParentURL:      parentURL.(string),
				ChildrensMenus: childrens,
			}
			menus = append(menus, functionMenu)
		}
	}
	return menus
}
