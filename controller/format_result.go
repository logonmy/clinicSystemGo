package controller

import (
	"database/sql"

	"github.com/fatih/structs"

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
 * 格式化数据库返回的数据
 */
func FormatSQLRowToMapArray(rows *sql.Rows) []map[string]interface{} {
	var results []map[string]interface{}
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		results = append(results, m)
	}
	return results
}
