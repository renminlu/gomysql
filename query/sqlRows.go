package query

import (
	// "database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetRows(db *sqlx.DB, sql string) (*map[int]map[string]string, error) {
	results := make(map[int]map[string]string)

	//①：查询操作
	//Db.Query(sql)：最原始的sql操作
	//Db.Query(&struct, sql)：会按照结构体属性类型 进行数据类型转换
	sqlRows, err := db.Query(sql)
	if err != nil {
		return &results, err
	}
	defer sqlRows.Close() //必须放在err后面，否则发送err时不可close
	//①：获取字段列表
	fields, _ := sqlRows.Columns()
	// fmt.Println(fields)
	//

	//values用于和scans进行映射（断言）
	values := make([][]byte, len(fields)) //可以处理字符、数字、null等
	//values := make([]string, len(fields))//不能处理【null值的字段】
	// values := make([](sql.NullString), len(fields))  //报错：sql.NullString未定义

	//scans：一维数组
	scans := make([]interface{}, len(fields))
	for i := range values {
		scans[i] = &values[i]
		//sqlRows.Scan参数只能时interface类型的
		//interface不支持range迭代，所以在此处把【scans和values】映射（断言）
	}

	i := 0
	//Next循环迭代
	for sqlRows.Next() {

		//②：获取当前行的字段值列表：索引型的一维数组
		//sqlRows.Scan每Next一次：查询出来的不定长一行值放到scans中
		//因为【scans[i]= &values[i]】,也就是每行都放在values里
		if err := sqlRows.Scan(scans...); err != nil {
			return &results, err
		}

		//③：【键值对】对应合并
		row := make(map[string]string) //每行数据【key=>val】
		for k, v := range values {
			row[fields[k]] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	return &results, nil
}
