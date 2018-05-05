package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//References 检验项参考值
type References struct {
	ReferenceSex  interface{} `json:"reference_sex"`
	ReferenceMax  interface{} `json:"reference_max"`
	ReferenceMin  interface{} `json:"reference_min"`
	IsPregnancy   interface{} `json:"isPregnancy"`
	StomachStatus interface{} `json:"stomach_status"`
}

//LaboratoryItem 检验项
type LaboratoryItem struct {
	ClinicLaboratoryItemID int64        `json:"clinic_laboratory_item_id"`
	LaboratoryItemID       int64        `json:"laboratory_item_id"`
	Name                   string       `json:"name"`
	EnName                 interface{}  `json:"en_name"`
	Unit                   interface{}  `json:"unit"`
	Status                 bool         `json:"status"`
	References             []References `json:"references"`
}

//LaboratoryCreate 检验医嘱创建
func LaboratoryCreate(ctx iris.Context) {

}

//LaboratoryItemCreate 检验项目创建
func LaboratoryItemCreate(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	name := ctx.PostValue("name")
	enName := ctx.PostValue("en_name")
	instrumentCode := ctx.PostValue("instrument_code")
	unit := ctx.PostValue("unit")
	clinicalSignificance := ctx.PostValue("clinical_significance")
	dataType := ctx.PostValue("data_type")

	isSpecial := ctx.PostValue("is_special")
	referenceMax := ctx.PostValue("reference_max")
	referenceMin := ctx.PostValue("reference_min")
	referenceValue := ctx.PostValue("reference_value")
	items := ctx.PostValue("items")

	status := ctx.PostValue("status")
	isDelivery := ctx.PostValue("is_delivery")

	if clinicID == "" || name == "" || dataType == "" || isSpecial == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	row := model.DB.QueryRowx("select id from clinic where id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "新增失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)
	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "诊所数据错误"})
		return
	}

	laboratoryItemSets := []string{"name", "data_type", "is_special"}
	laboratoryItemValues := []string{"'" + name + "'", dataType, isSpecial}

	var itemReferenceSets []string
	var itemReferenceValues []string

	clinicLaboratoryItemSets := []string{"clinic_id"}
	clinicLaboratoryItemValues := []string{clinicID}

	if enName != "" {
		laboratoryItemSets = append(laboratoryItemSets, "en_name")
		laboratoryItemValues = append(laboratoryItemValues, "'"+enName+"'")
	}
	if instrumentCode != "" {
		laboratoryItemSets = append(laboratoryItemSets, "instrument_code")
		laboratoryItemValues = append(laboratoryItemValues, "'"+instrumentCode+"'")
	}
	if unit != "" {
		laboratoryItemSets = append(laboratoryItemSets, "unit")
		laboratoryItemValues = append(laboratoryItemValues, "'"+unit+"'")
	}
	if clinicalSignificance != "" {
		laboratoryItemSets = append(laboratoryItemSets, "clinical_significance")
		laboratoryItemValues = append(laboratoryItemValues, "'"+clinicalSignificance+"'")
	}

	if status != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "status")
		clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, status)
	}
	if isDelivery != "" {
		clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "is_delivery")
		clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, isDelivery)
	}

	laboratoryItemSetStr := strings.Join(laboratoryItemSets, ",")
	laboratoryItemValueStr := strings.Join(laboratoryItemValues, ",")

	laboratoryItemInsertSQL := "insert into laboratory_item (" + laboratoryItemSetStr + ") values (" + laboratoryItemValueStr + ") RETURNING id;"
	fmt.Println("laboratoryItemInsertSQL==", laboratoryItemInsertSQL)

	tx, err := model.DB.Begin()
	var laboratoryItemID string
	err = tx.QueryRow(laboratoryItemInsertSQL).Scan(&laboratoryItemID)
	if err != nil {
		fmt.Println("err ===", err)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": err})
		return
	}
	fmt.Println("laboratoryItemID====", laboratoryItemID)

	itemReferenceSets = append(itemReferenceSets, "laboratory_item_id")
	if isSpecial == "false" {
		itemReferenceValues = append(itemReferenceValues, laboratoryItemID)
		if referenceMax != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_max")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMax+"'")
		}
		if referenceMin != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_min")
			itemReferenceValues = append(itemReferenceValues, "'"+referenceMin+"'")
		}
		if referenceValue != "" {
			itemReferenceSets = append(itemReferenceSets, "reference_value")
			itemReferenceValues = append(itemReferenceValues, referenceValue)
		}
	} else if isSpecial == "true" && items != "" {
		itemReferenceSets = append(itemReferenceSets, "reference_sex", "age_max", "age_min", "reference_max", "reference_min", "stomach_status", "is_pregnancy")
		var results []map[string]string
		reErr := json.Unmarshal([]byte(items), &results)
		if reErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": reErr.Error()})
			return
		}
		for _, v := range results {
			var s []string
			s = append(s, laboratoryItemID)
			referenceSex := v["reference_sex"]
			ageMax := v["age_max"]
			ageMin := v["age_min"]
			referenceMax := v["reference_max"]
			referenceMin := v["reference_min"]
			stomachStatus := v["stomach_status"]
			isPregnancy := v["is_pregnancy"]
			if referenceSex != "" {
				s = append(s, "'"+referenceSex+"'")
			} else {
				s = append(s, `null`)
			}
			if ageMax != "" {
				s = append(s, ageMax)
			} else {
				s = append(s, `null`)
			}
			if ageMin != "" {
				s = append(s, ageMin)
			} else {
				s = append(s, `null`)
			}
			if referenceMax != "" {
				s = append(s, "'"+referenceMax+"'")
			} else {
				s = append(s, `null`)
			}
			if referenceMin != "" {
				s = append(s, "'"+referenceMin+"'")
			} else {
				s = append(s, `null`)
			}
			if stomachStatus != "" {
				s = append(s, "'"+stomachStatus+"'")
			} else {
				s = append(s, `null`)
			}
			if isPregnancy != "" {
				s = append(s, isPregnancy)
			} else {
				s = append(s, `null`)
			}
			str := strings.Join(s, ",")
			str = "(" + str + ")"
			itemReferenceValues = append(itemReferenceValues, str)
		}
	} else {
		ctx.JSON(iris.Map{"code": "1", "msg": "参考值是否特殊数据格式错误"})
		return
	}

	itemReferenceSetStr := strings.Join(itemReferenceSets, ",")
	itemReferenceValueStr := strings.Join(itemReferenceValues, ",")

	clinicLaboratoryItemSets = append(clinicLaboratoryItemSets, "laboratory_item_id")
	clinicLaboratoryItemValues = append(clinicLaboratoryItemValues, laboratoryItemID)

	clinicLaboratoryItemSetStr := strings.Join(clinicLaboratoryItemSets, ",")
	clinicLaboratoryItemValueStr := strings.Join(clinicLaboratoryItemValues, ",")

	itemReferenceInsertSQL := "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values (" + itemReferenceValueStr + ")"
	if isSpecial == "true" {
		itemReferenceInsertSQL = "insert into laboratory_item_reference (" + itemReferenceSetStr + ") values " + itemReferenceValueStr
	}
	fmt.Println("itemReferenceInsertSQL==", itemReferenceInsertSQL)

	clinicLaboratoryItemInsertSQL := "insert into clinic_laboratory_item (" + clinicLaboratoryItemSetStr + ") values (" + clinicLaboratoryItemValueStr + ")"
	fmt.Println("clinicLaboratoryItemInsertSQL==", clinicLaboratoryItemInsertSQL)

	_, err1 := tx.Exec(itemReferenceInsertSQL)
	if err1 != nil {
		fmt.Println(" err1====", err1)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	_, err2 := tx.Exec(clinicLaboratoryItemInsertSQL)
	if err2 != nil {
		fmt.Println(" err2====", err2)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItemID})
}

//LaboratoryItemList 检验项目列表
func LaboratoryItemList(ctx iris.Context) {
	clinicID := ctx.PostValue("clinic_id")
	keyword := ctx.PostValue("name")
	status := ctx.PostValue("status")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if clinicID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}
	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	row := model.DB.QueryRowx("select id from storehouse where clinic_id=$1 limit 1", clinicID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "查询失败"})
		return
	}
	clinic := FormatSQLRowToMap(row)

	_, ok := clinic["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "所在诊所不存在"})
		return
	}

	countSQL := `select count(cli.id) as total from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		where cli.clinic_id=$1`
	selectSQL := `select li.id as laboratory_item_id,cli.id as clinic_laboratory_item_id,li.name,li.en_name,li.unit,
		lir.reference_sex,lir.stomach_status,lir.is_pregnancy,lir.reference_max,lir.reference_min,cli.status
		from clinic_laboratory_item cli
		left join laboratory_item li on cli.laboratory_item_id = li.id
		left join laboratory_item_reference lir on lir.laboratory_item_id = li.id
		where cli.clinic_id=$1`

	if keyword != "" {
		countSQL += " and li.name ~'" + keyword + "'"
		selectSQL += " and li.name ~'" + keyword + "'"
	}
	if status != "" {
		countSQL += " and cli.status=" + status
		selectSQL += " and cli.status=" + status
	}

	fmt.Println("countSQL===", countSQL)
	fmt.Println("selectSQL===", selectSQL)
	total := model.DB.QueryRowx(countSQL, clinicID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	var results []map[string]interface{}
	rows, _ := model.DB.Queryx(selectSQL+" offset $2 limit $3", clinicID, offset, limit)
	results = FormatSQLRowsToMapArray(rows)

	var laboratoryItems []LaboratoryItem
	for _, v := range results {
		clinicLaboratoryItemID := v["clinic_laboratory_item_id"]
		laboratoryItemID := v["laboratory_item_id"]
		name := v["name"]
		enName := v["en_name"]
		unit := v["unit"]
		status := v["status"]
		referenceSex := v["reference_sex"]
		referenceMax := v["reference_max"]
		referenceMin := v["reference_min"]
		isPregnancy := v["is_pregnancy"]
		stomachStatus := v["stomach_status"]
		has := false
		for k, vRes := range laboratoryItems {
			vlaboratoryItemID := vRes.LaboratoryItemID
			vreferences := vRes.References
			if vlaboratoryItemID == laboratoryItemID.(int64) {
				reference := References{
					ReferenceSex:  referenceSex,
					ReferenceMax:  referenceMax,
					ReferenceMin:  referenceMin,
					IsPregnancy:   isPregnancy,
					StomachStatus: stomachStatus,
				}
				laboratoryItems[k].References = append(vreferences, reference)
				has = true
			}
		}
		if !has {
			var references []References
			reference := References{
				ReferenceSex:  referenceSex,
				ReferenceMax:  referenceMax,
				ReferenceMin:  referenceMin,
				IsPregnancy:   isPregnancy,
				StomachStatus: stomachStatus,
			}
			references = append(references, reference)

			laboratoryItem := LaboratoryItem{
				ClinicLaboratoryItemID: clinicLaboratoryItemID.(int64),
				LaboratoryItemID:       laboratoryItemID.(int64),
				Name:                   name.(string),
				EnName:                 enName,
				Unit:                   unit,
				Status:                 status.(bool),
				References:             references,
			}
			laboratoryItems = append(laboratoryItems, laboratoryItem)
		}
	}

	ctx.JSON(iris.Map{"code": "200", "data": laboratoryItems, "page_info": pageInfo})
}
