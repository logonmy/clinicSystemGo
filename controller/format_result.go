package controller

import (
	"github.com/jmoiron/sqlx"
)

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
