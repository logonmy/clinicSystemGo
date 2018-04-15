package controller

import "github.com/fatih/structs"

/**
 * 返回结果
 */
type APIJSON struct {
	Code int
	Msg  string
	Data interface{}
}

/**
 * 格式化结果
 */
func FormatResult(apijson APIJSON) map[string]interface{} {
	obj := structs.Map(apijson)
	return obj
}
