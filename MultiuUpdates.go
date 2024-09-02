package goo

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

// query 是where条件，变量使用? 占位符， args 是where条件参数
type MUWhere struct {
	Query string // where语句包含占位符
	Args  []any  // where语句占位符对应的参数
}

type MUClauses struct {
	DB            *gorm.DB // 数据库连接
	Table         string   // 表名
	Where         *MUWhere // 追加的where条件
	Chunk         int      // 按块大小拆分
	CaseColumn    string   // 作为条件的列名
	UpdateColumns []string // 更新列
}

/*
UPDATE table_name
SET ip = CASE id

	WHEN 83 THEN '10.215.14.216'
	WHEN 82 THEN '10.215.14.215'

END, test1 = CASE id

	WHEN 83 THEN 24,
	WHEN 82 THEN 6

END
WHERE id IN (83, 82)
*/
func MultiUpdates[T any](cla *MUClauses, datas []T) error {
	if cla.CaseColumn == "" {
		return errors.New("CaseColumn is empty")
	}

	if len(cla.UpdateColumns) <= 0 {
		return errors.New("updateColumns is empty")
	}

	if cla.Table == "" {
		return errors.New("table is empty")
	}

	if cla.DB == nil {
		return errors.New("db is nil")
	}

	if cla.Chunk <= 0 {
		cla.Chunk = 1000
	}

	if len(datas) <= 0 {
		return nil
	}

	datasChunk := ArrayChunk(datas, cla.Chunk)

	var err error

	for _, datas := range datasChunk {

		switch {
		case IsMap(datas[0]):
			err = multiUpdateMaps(cla, datas)
		case IsStruct(datas[0]):
			err = multiUpdateStructs(cla, datas)
		default:
			err = errors.New("data type must be map[string]any or struct")
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func multiUpdateMaps[T any](cla *MUClauses, datas []T) error {

	var updateColumnsStr = make(map[string][]string)
	var updateColumnsArgs = make(map[string][]any)
	var args []any
	var CaseColumnValues []any
DataLoop:
	for _, data := range datas {
		if !IsMap(data) {
			continue
		}

		mapData := AnyConvert2T[map[string]any](data, nil)
		if mapData == nil {
			return errors.New("map's data type must be map[string]any and not nil")
		}

		//判断where在不在
		if !IsSet(mapData, cla.CaseColumn) {
			continue
		}

		//判断where是否要加引号
		whereValue := mapData[cla.CaseColumn]
		// wherePlaceholder := getFormatPlaceholder(whereValue)

		//判断update列在不在
		for _, updateColumn := range cla.UpdateColumns {
			if !IsSet(mapData, updateColumn) {
				continue DataLoop
			}

			if !IsSet(updateColumnsStr, updateColumn) {
				updateColumnsStr[updateColumn] = []string{}
				updateColumnsArgs[updateColumn] = []any{}
			}
			// WHEN 83 THEN '10.215.14.216'
			//判断是否要加引号
			updateValue := mapData[updateColumn]
			// updatePlaceholder := getFormatPlaceholder(updateValue)

			format := "WHEN ? THEN ?"

			updateColumnsStr[updateColumn] = append(updateColumnsStr[updateColumn], format)
			updateColumnsArgs[updateColumn] = append(updateColumnsArgs[updateColumn], whereValue, updateValue)
		}

		CaseColumnValues = append(CaseColumnValues, whereValue)
	}

	CaseColumnValues = ArrayUnique(CaseColumnValues)

	//如果没有where列，则不执行
	if len(CaseColumnValues) == 0 {
		return errors.New("where column len is zero")
	}

	SQL := "UPDATE `" + cla.Table + "` SET"
	updateLens := len(updateColumnsStr)
	index := 0
	for updateColumn, updateValues := range updateColumnsStr {
		index++
		if len(updateValues) == 0 {
			continue
		}
		/*
			ip = CASE id

				WHEN 83 THEN '10.215.14.216'
				WHEN 82 THEN '10.215.14.215'

			END,
		*/

		SQL += fmt.Sprintf(` %s = CASE %s

		`, updateColumn, cla.CaseColumn)

		for _, updateValue := range updateValues {
			SQL += fmt.Sprintf(`%s
		`, updateValue)
		}

		args = append(args, updateColumnsArgs[updateColumn]...)

		SQL += `END`

		if index < updateLens {
			SQL += ","
		}
	}

	whereStr := ""
	for _, CaseColumnValue := range CaseColumnValues {
		whereStr += fmt.Sprintf(",%s", CaseColumnValue)
	}

	SQL += fmt.Sprintf(` 
	WHERE %s IN (?)`, cla.CaseColumn)

	var CaseColumnValuesAny any = CaseColumnValues
	args = append(args, CaseColumnValuesAny)

	if cla.Where != nil && cla.Where.Query != "" {
		SQL += fmt.Sprintf(" AND %s", cla.Where.Query)
		args = append(args, cla.Where.Args...)
	}

	// logger.Info("multiUpdateMaps SQL:%s; args:%v", strings.Replace(SQL, "\n", " ", -1), args)

	//执行sql
	return cla.DB.Exec(SQL, args...).Error
}

// 将struct转成map, 请求multiUpdateMaps
func multiUpdateStructs[T any](cla *MUClauses, datas []T) error {

	var mapDatas []map[string]any

	for _, structData := range datas {
		val := reflect.ValueOf(structData)

		//指针类型取值
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		t := val.Type()

		if t.Kind() != reflect.Struct {
			return fmt.Errorf("expected a struct, got %s", t.Kind())
		}

		var allColumnExists = make(map[string]bool)
		allColumnExists[cla.CaseColumn] = false

		for _, updateColumn := range cla.UpdateColumns {
			allColumnExists[updateColumn] = false
		}

		var mapData = make(map[string]any)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			tag := field.Tag

			tagColumn, ok := parseGormColumnTag(tag)
			if !ok {
				//如果column不存在，则使用json tag
				tagColumn = tag.Get("json")
			}

			if tagColumn == "" {
				continue
			}

			if _, ok := allColumnExists[tagColumn]; !ok {
				continue
			}
			//赋值 & 标记
			allColumnExists[tagColumn] = true
			mapData[tagColumn] = val.FieldByName(field.Name).Interface()

			if mapData[tagColumn] == nil {
				return fmt.Errorf("column value is nil:%s; data:%+v", tagColumn, structData)
			}
		}

		for col, exists := range allColumnExists {
			if !exists {
				return fmt.Errorf("struct type data must set tag `gorm:column` or `json` same as mysql column;column not exists:%s; data:%+v", col, structData)
			}
		}

		mapDatas = append(mapDatas, mapData)
	}

	return multiUpdateMaps(cla, mapDatas)
}

// parseGormColumnTag 解析 gorm 标签中的 "column" 属性
func parseGormColumnTag(tag reflect.StructTag) (columnName string, hasColumn bool) {
	if tag == "" {
		return "", false
	}

	tags := reflect.StructTag(tag)
	for _, part := range strings.Split(tags.Get("gorm"), ";") {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:"), true
		}
	}
	return "", false
}
