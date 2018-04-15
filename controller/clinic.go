package controller

import (
	"clinicSystemGo/model"
	"fmt"
	"crypto/md5"
	"github.com/kataras/iris"
	"encoding/hex"
)

/**
 * 获取诊所
 */
func GetClinicByCode(ctx iris.Context) {
	code := ctx.PostValue("code")
	if code == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	clinic := model.Clinic{}
	err := model.DB.Get(&clinic, "SELECT * FROM clinic WHERE code=$1", code)

	if err != nil {
			fmt.Println("err ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
	}
	ctx.JSON(iris.Map{"code": "200", "data": clinic})
}

func ClinicList(ctx iris.Context) {

	clinic := []model.Clinic{}
	err := model.DB.Select(&clinic, "SELECT * FROM clinic ")

	if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
	}
	ctx.JSON(iris.Map{"code": "200", "data": clinic})
}

func ClinicUpdte(ctx iris.Context) {



}

/**
* 新建诊所
*/
func ClinicAdd(ctx iris.Context) {
	code := ctx.PostValue("code")
	name := ctx.PostValue("name")
	responsible_person := ctx.PostValue("responsible_person")
	area := ctx.PostValue("area")
	status := ctx.PostValue("status")

	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	if (code == "" || name == "" || responsible_person == "" || area == "" || status == "" || username == "" || password == "") {
		ctx.JSON(iris.Map{"code": "-1", "msg":"缺少参数" })
		return
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	passwordMd5 := hex.EncodeToString(md5Ctx.Sum(nil))
	
	cmap := map[string]interface{}{
		"code": code ,
		"name": name,
		"responsible_person": responsible_person,
		"area": area,
		"status" : status,
	}

	amap := map[string]interface{}{
		"code": "10000",
		"name": "超级管理员",
		"username": username ,
		"password": passwordMd5 ,
		"clinic_code": code,
	  "is_clinic_admin": true,
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		fmt.Println("err ===", err.Error())
		ctx.JSON(iris.Map{"code": "-1", "msg":err })
		return
	}

	_, err = tx.NamedExec(`INSERT INTO clinic(
		code, name, responsible_person, area, status)
		VALUES (:code, :name, :responsible_person, :area, :status)`,cmap)

	if err != nil {
		fmt.Println("err ===", err.Error())
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg":err.Error() })
		return
	}

	_, err = tx.NamedExec(`INSERT INTO personnel(
		code, name, username, password, clinic_code, is_clinic_admin)
		VALUES (:code, :name, :username, :password, :clinic_code, :is_clinic_admin)`,amap)
	if err != nil {
		  fmt.Println("err ===", err.Error())
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg":err.Error() })
			return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg":err.Error() })
		return
  }
	ctx.JSON(iris.Map{"code": "200", "data":nil})
}