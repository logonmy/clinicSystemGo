package controller

import (
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

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

/**
 * 格式化数据库返回的数组数据
 */
func FormatSQLRowsToMapArray(rows *sqlx.Rows) []map[string]interface{} {
	var results []map[string]interface{}
	for rows.Next() {
		m := make(map[string]interface{})
		rows.MapScan(m)
		results = append(results, m)
	}
	return results
}

/**
 * 格式化数据库返回的数组数据
 */
func FormatSQLRowToMap(row *sqlx.Row) map[string]interface{} {
	result := make(map[string]interface{})
	row.MapScan(result)
	return result
}
