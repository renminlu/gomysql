package table

import (
	// "fmt"
	"strings"
)

type TableStruct struct {
	//以[字段名]为第一维key
	//以[Field Type Null Key Default Extra]为第二维key
	structs map[string]map[string]string
	primary string
}

//判断：某个字段是否存在
func (t *TableStruct) IsField(field string) bool {
	//通过一个必然存在的属性：检测是否存在
	_, ok := t.structs[field]["Type"]
	return ok
}

//获取：表主键字段
func (t *TableStruct) GetTablePrimarry() string {
	return t.primary
}

//获取：某个字段：数据类型
func (t *TableStruct) GetFieldType(field string) string {
	return t.structs[field]["Type"]
}

//获取：某个字段：是否可以为空
func (t *TableStruct) FieldCanNull(field string) bool {
	// fmt.Println("GetFieldCanNull调试：", t.structs[field]["Null"])
	//Null=NO：该字段不可为空；否则可以为空
	if t.structs[field]["Null"] == "NO" {
		return false
	}
	if t.structs[field]["Null"] == "YES" {
		return true
	}
	return true
}

//获取：某个字段：索引属性
func (t *TableStruct) GetFieldKey(field string) string {
	return t.structs[field]["Key"]
}

//获取：某个字段：默认值
func (t *TableStruct) GetFieldDefault(field string) string {
	return t.structs[field]["Default"]
}

//获取：某个字段：Extra
func (t *TableStruct) GetFieldExtra(field string) string {
	return t.structs[field]["Extra"]
}

//判断：某个字段：是否是字符串系列
func (t *TableStruct) FieldTypeIsString(field string) bool {
	typeN := t.structs[field]["Type"]
	strList := [3]string{"char", "text", "json"}
	for _, v := range strList {
		if strings.Contains(typeN, v) {
			return true
		}
	}
	return false
}
