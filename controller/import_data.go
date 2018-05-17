package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"strings"

	"github.com/kataras/iris"
	"github.com/tealeg/xlsx"
)

//ImportLaboratory 导入检验
func ImportLaboratory(ctx iris.Context) {
	excelFileName := "laboratory.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}
	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("导入失败===", err)
		tx.Rollback()
		return
	}
	count := 0
	for index, row := range xlFile.Sheets[1].Rows {
		laboratorySets := []string{"name", "remark", "time_report"}
		var laboratoryValues []string
		if index == 0 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[0].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from laboratory where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		laboratory := FormatSQLRowToMap(lrow)
		_, lok := laboratory["id"]
		if lok {
			continue
		}
		for _, cell := range row.Cells {
			laboratoryValues = append(laboratoryValues, "'"+cell.String()+"'")
		}
		laboratorySetStr := strings.Join(laboratorySets, ",")
		laboratoryValueStr := strings.Join(laboratoryValues, ",")

		laboratoryInsertSQL := "insert into laboratory (" + laboratorySetStr + ") values (" + laboratoryValueStr + ") RETURNING id;"
		fmt.Println("laboratoryInsertSQL ===", laboratoryInsertSQL)
		var laboratoryID string
		err = tx.QueryRow(laboratoryInsertSQL).Scan(&laboratoryID)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		fmt.Println("laboratoryID====", laboratoryID)
		clinicLaboratoryInsertSQL := "insert into clinic_laboratory (clinic_id,price,laboratory_id) values (1,0,$1)"

		_, err2 := tx.Exec(clinicLaboratoryInsertSQL, laboratoryID)
		if err2 != nil {
			fmt.Println(" err2====", err2)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}
	}
	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//ImportLaboratoryItem 导入检验项目
func ImportLaboratoryItem(ctx iris.Context) {
	excelFileName := "laboratory_item.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}
	tx, err := model.DB.Begin()
	if err != nil {
		fmt.Println("导入失败===", err)
		tx.Rollback()
		return
	}
	count := 0
	countMap := map[string]string{
		"a": "a",
	}
	for index, row := range xlFile.Sheets[0].Rows {
		laboratoryItmeSets := []string{"name", "en_name", "unit_name", "clinical_significance"}
		laboratoryItmeReferenceSets := []string{"reference_min", "reference_max", "reference_sex", "laboratory_item_id"}
		var laboratoryItemValues []string
		var laboratoryItemReferenceValues []string
		if index == 0 || index == 1 {
			continue
		}
		if count > 2 {
			break
		}
		name := row.Cells[1].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		_, ok := countMap[name]
		if ok {
			fmt.Println("====", name)
			continue
		}
		countMap[name] = name
		lrow := model.DB.QueryRowx("select id from laboratory_item where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		laboratoryItem := FormatSQLRowToMap(lrow)
		_, lok := laboratoryItem["id"]
		if lok {
			continue
		}
		for indexc, cell := range row.Cells {
			if indexc == 0 || indexc == 3 {
				continue
			}
			if indexc == 4 || indexc == 5 {
				laboratoryItemReferenceValues = append(laboratoryItemReferenceValues, "'"+cell.String()+"'")
			} else {
				laboratoryItemValues = append(laboratoryItemValues, "'"+cell.String()+"'")
			}
		}
		laboratorySetStr := strings.Join(laboratoryItmeSets, ",")
		laboratoryValueStr := strings.Join(laboratoryItemValues, ",")

		laboratoryInsertSQL := "insert into laboratory_item (" + laboratorySetStr + ") values (" + laboratoryValueStr + ") RETURNING id;"
		fmt.Println("laboratoryInsertSQL ===", laboratoryInsertSQL)
		var laboratoryItemID string
		errq := tx.QueryRow(laboratoryInsertSQL).Scan(&laboratoryItemID)
		if errq != nil {
			fmt.Println("errq ===", errq)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errq.Error()})
			return
		}
		fmt.Println("laboratoryItemID====", laboratoryItemID)

		laboratoryItemReferenceValues = append(laboratoryItemReferenceValues, "'通用'", laboratoryItemID)
		referenceSetStr := strings.Join(laboratoryItmeReferenceSets, ",")
		referenceValueStr := strings.Join(laboratoryItemReferenceValues, ",")

		clinicLaboratoryInsertSQL := "insert into clinic_laboratory_item (clinic_id,laboratory_item_id) values (1,$1)"
		laboratoryItemReferenceSQL := "insert into laboratory_item_reference (" + referenceSetStr + ") values (" + referenceValueStr + ")"

		fmt.Println("clinicLaboratoryInsertSQL====", clinicLaboratoryInsertSQL)
		fmt.Println("laboratoryItemReferenceSQL====", laboratoryItemReferenceSQL)

		_, err2 := tx.Exec(laboratoryItemReferenceSQL)
		if err2 != nil {
			fmt.Println(" err2====", err2)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
			return
		}

		_, err1 := tx.Exec(clinicLaboratoryInsertSQL, laboratoryItemID)
		if err1 != nil {
			fmt.Println(" err1====", err1)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
			return
		}
	}
	err3 := tx.Commit()
	if err3 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": err3.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})
}
