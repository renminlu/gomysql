package table

import (
	// "fmt"
	// "database/sql"
	"github.com/renminlu/gomysql/query"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
 * db 数据库连接对象
 * table 表名
 */
func New(db *sqlx.DB, table string) (*TableStruct, error) {
	ts := TableStruct{}
	rows, err := query.GetRows(db, "desc "+table)
	if err != nil {
		return &ts, err
	}
	structs := make(map[string]map[string]string, 5)
	for _, v := range *rows {

		field := make(map[string]string, 5)
		field["Type"] = v["Type"]
		field["Null"] = v["Null"]
		field["Key"] = v["Key"]
		field["Default"] = v["Default"]
		field["Extra"] = v["Extra"]

		if v["Key"] == "PRI" && v["Extra"] == "auto_increment" {
			ts.primary = v["Field"]
		}
		// fmt.Println("table.New调试", field)
		structs[v["Field"]] = field
	}
	ts.structs = structs
	return &ts, nil
}
