package controller

import (
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
