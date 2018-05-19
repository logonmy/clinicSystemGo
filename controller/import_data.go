package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"regexp"
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
		laboratoryItmeSets := []string{"name", "en_name", "is_special", "data_type", "unit_name", "clinical_significance"}
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
		if name == "" {
			count++
			continue
		}
		_, ok := countMap[name]
		if ok {
			fmt.Println("=&=&=&=&=", name)
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
			if indexc == 4 {
				laboratoryItemReferenceValues = append(laboratoryItemReferenceValues, "'"+cell.String()+"'")
				if cell.String() == "" {
					laboratoryItemValues = append(laboratoryItemValues, "true")
					laboratoryItemValues = append(laboratoryItemValues, "2")
				} else {
					laboratoryItemValues = append(laboratoryItemValues, "false")
					if m, _ := regexp.MatchString(".*性.*", cell.String()); m {
						laboratoryItemValues = append(laboratoryItemValues, "1")
					} else {
						laboratoryItemValues = append(laboratoryItemValues, "2")
					}
				}
			} else if indexc == 5 {
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

//ImportFrequency 导入用药频率
func ImportFrequency(ctx iris.Context) {
	excelFileName := "frequency.xlsx"
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
		frequencySets := []string{"name", "py_code", "define_code", "print_code", "week_day_flag", "update_flag", "delete_flag",
			"weight", "in_out_flag", "medical_bill_code", "input_frequency", "times", "days", "code"}
		var frequencyValues []string
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
		lrow := model.DB.QueryRowx("select id from frequency where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		frequency := FormatSQLRowToMap(lrow)
		_, lok := frequency["id"]
		if lok {
			continue
		}
		for indexf, cell := range row.Cells {
			if indexf == 10 {
				continue
			}
			if cell.String() == "" {
				frequencyValues = append(frequencyValues, `NULL`)
			} else {
				if indexf == 0 || indexf == 1 || indexf == 2 || indexf == 3 || indexf == 11 || indexf == 14 {
					frequencyValues = append(frequencyValues, "'"+cell.String()+"'")
				} else {
					frequencyValues = append(frequencyValues, cell.String())
				}
			}
		}
		frequencySetStr := strings.Join(frequencySets, ",")
		frequencyValueStr := strings.Join(frequencyValues, ",")

		frequencyInsertSQL := "insert into frequency (" + frequencySetStr + ") values (" + frequencyValueStr + ")"
		fmt.Println("frequencyInsertSQL ===", frequencyInsertSQL)
		_, err = tx.Exec(frequencyInsertSQL)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

//ImportDoseUnit 导入单位
func ImportDoseUnit(ctx iris.Context) {
	excelFileName := "doseUnit.xlsx"
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
		doseUnitSets := []string{"code", "change_flag", "name", "d_code", "py_code", "deleted_flag"}
		var doseUnitValues []string
		if index == 0 || index == 1 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[2].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from dose_unit where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		doseUnit := FormatSQLRowToMap(lrow)
		_, lok := doseUnit["id"]
		if lok {
			continue
		}
		for indexf, cell := range row.Cells {
			if cell.String() == "" {
				doseUnitValues = append(doseUnitValues, `NULL`)
			} else {
				if indexf == 0 || indexf == 2 || indexf == 3 || indexf == 4 {
					doseUnitValues = append(doseUnitValues, "'"+cell.String()+"'")
				} else {
					doseUnitValues = append(doseUnitValues, cell.String())
				}
			}
		}
		doseUnitSetStr := strings.Join(doseUnitSets, ",")
		doseUnitValueStr := strings.Join(doseUnitValues, ",")

		doseUnitInsertSQL := "insert into dose_unit (" + doseUnitSetStr + ") values (" + doseUnitValueStr + ")"
		fmt.Println("doseUnitInsertSQL ===", doseUnitInsertSQL)
		_, err = tx.Exec(doseUnitInsertSQL)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

//ImportDoseForm 导入剂型
func ImportDoseForm(ctx iris.Context) {
	excelFileName := "doseForm.xlsx"
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
		doseFormSets := []string{"py_code", "deleted_flag", "name", "d_code", "code"}
		var doseFormValues []string
		if index == 0 || index == 1 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[2].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from dose_form where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		doseForm := FormatSQLRowToMap(lrow)
		_, lok := doseForm["id"]
		if lok {
			continue
		}
		for indexf, cell := range row.Cells {
			if indexf > 4 {
				break
			}
			if cell.String() == "" {
				doseFormValues = append(doseFormValues, `NULL`)
			} else {
				if indexf == 0 || indexf == 2 || indexf == 3 || indexf == 4 {
					doseFormValues = append(doseFormValues, "'"+cell.String()+"'")
				} else {
					doseFormValues = append(doseFormValues, cell.String())
				}
			}
		}
		doseFormSetStr := strings.Join(doseFormSets, ",")
		doseFormValueStr := strings.Join(doseFormValues, ",")

		doseFormInsertSQL := "insert into dose_form (" + doseFormSetStr + ") values (" + doseFormValueStr + ")"
		fmt.Println("doseFormInsertSQL ===", doseFormInsertSQL)
		_, err = tx.Exec(doseFormInsertSQL)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

//ImportDrugType 导入药品类型
func ImportDrugType(ctx iris.Context) {
	excelFileName := "drugType.xlsx"
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
		drugTypeSets := []string{"deleted_flag", "d_code", "name", "py_code", "code"}
		var drugTypeValues []string
		if index == 0 || index == 1 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[2].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from drug_type where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		drugType := FormatSQLRowToMap(lrow)
		_, lok := drugType["id"]
		if lok {
			continue
		}
		for indexf, cell := range row.Cells {
			if indexf > 4 {
				break
			}
			if cell.String() == "" {
				drugTypeValues = append(drugTypeValues, `NULL`)
			} else {
				if indexf == 1 || indexf == 2 || indexf == 3 || indexf == 4 {
					drugTypeValues = append(drugTypeValues, "'"+cell.String()+"'")
				} else {
					drugTypeValues = append(drugTypeValues, cell.String())
				}
			}
		}
		drugTypeSetStr := strings.Join(drugTypeSets, ",")
		drugTypeValueStr := strings.Join(drugTypeValues, ",")

		drugTypeInsertSQL := "insert into drug_type (" + drugTypeSetStr + ") values (" + drugTypeValueStr + ")"
		fmt.Println("drugTypeInsertSQL ===", drugTypeInsertSQL)
		_, err = tx.Exec(drugTypeInsertSQL)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

//ImportManuFactory 导入生产产商
func ImportManuFactory(ctx iris.Context) {
	excelFileName := "manu_factory.xlsx"
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
		manuFactorySets := []string{"code", "name", "abbr_name", "py_code", "d_code", "deleted_flag"}
		if index == 0 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[1].String()
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from manu_factory where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		manuFactory := FormatSQLRowToMap(lrow)
		_, lok := manuFactory["id"]
		if lok {
			continue
		}
		code := row.Cells[0].String()
		abbrName := row.Cells[2].String()
		pyCode := row.Cells[6].String()
		dCode := row.Cells[7].String()
		deletedFlag := row.Cells[10].String()
		if deletedFlag == "" {
			deletedFlag = "0"
		}
		manuFactorySetStr := strings.Join(manuFactorySets, ",")

		manuFactoryInsertSQL := "insert into manu_factory (" + manuFactorySetStr + ") values ($1,$2,$3,$4,$5,$6)"
		_, err = tx.Exec(manuFactoryInsertSQL, code, name, abbrName, pyCode, dCode, deletedFlag)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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

//ImportrRouteAdministration 导入用药途径
func ImportrRouteAdministration(ctx iris.Context) {
	excelFileName := "routeOfAdministration.xlsx"
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
		routeAdministrationSets := []string{"input_type", "weight", "is_print", "deleted_flag", "print_name", "code",
			"type_code", "d_code", "py_code", "name"}
		var routeAdministrationValues []string
		if index == 0 || index == 1 {
			continue
		}
		if count > 5 {
			break
		}
		name := row.Cells[9].String()
		fmt.Println("name", name)
		if name == "" {
			count++
			continue
		}
		lrow := model.DB.QueryRowx("select id from route_administration where name=$1 limit 1", name)
		if lrow == nil {
			continue
		}
		routeAdministration := FormatSQLRowToMap(lrow)
		_, lok := routeAdministration["id"]
		if lok {
			continue
		}
		for indexf, cell := range row.Cells {
			if cell.String() == "" {
				routeAdministrationValues = append(routeAdministrationValues, `NULL`)
			} else {
				if indexf == 1 || indexf == 2 || indexf == 3 {
					routeAdministrationValues = append(routeAdministrationValues, cell.String())
				} else {
					routeAdministrationValues = append(routeAdministrationValues, "'"+cell.String()+"'")
				}
			}
		}
		routeAdministrationSetStr := strings.Join(routeAdministrationSets, ",")
		routeAdministrationValueStr := strings.Join(routeAdministrationValues, ",")

		routeAdministrationInsertSQL := "insert into route_administration (" + routeAdministrationSetStr + ") values (" + routeAdministrationValueStr + ")"
		fmt.Println("routeAdministrationInsertSQL ===", routeAdministrationInsertSQL)
		_, err = tx.Exec(routeAdministrationInsertSQL)
		if err != nil {
			fmt.Println("err ===", err)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
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
