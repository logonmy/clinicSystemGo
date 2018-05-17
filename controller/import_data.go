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
	for index, row := range xlFile.Sheets[0].Rows {
		laboratoryItmeSets := []string{"name", "en_name", "unit_name", "clinical_significance", "data_type", "is_special"}
		var laboratoryItemValues []string
		if index == 0 || index == 1 {
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
			laboratoryItemValues = append(laboratoryItemValues, "'"+cell.String()+"'")
		}
		laboratorySetStr := strings.Join(laboratoryItmeSets, ",")
		laboratoryValueStr := strings.Join(laboratoryItemValues, ",")

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
