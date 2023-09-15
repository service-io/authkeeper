// Package main
// @author tabuyos
// @since 2023/8/29
// @description main
package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"metis/helper/database"
	"metis/old/generated/helper"
	"metis/old/generated/repository"
	"strconv"
)

func main() {
	strcase.ConfigureAcronym("ID", "id")
	strcase.ConfigureAcronym("id", "ID")

	camel := strcase.ToCamel("a_id")

	fmt.Println(camel)

	database.InitDB()
	options := []*helper.Option{
		{Table: "sys_log", Entity: "SysLog", Variable: "eto", Variables: "ets"},
		{Table: "sys_oss", Entity: "SysOss", Variable: "eto", Variables: "ets"},
		{Table: "common_area", Entity: "CommonArea", Variable: "eto", Variables: "ets"},
	}
	for _, option := range options {
		columns := getColumns(option.Table)
		option.Cols = columns
		repository.New(option).RenderAuto()
	}
}

func getColumns(table string) []helper.Column {
	columns := make([]helper.Column, 0)
	db := database.FetchDB()
	rows, err := db.Query(
		fmt.Sprintf(
			`SELECT
		COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT,CHARACTER_MAXIMUM_LENGTH,COLUMN_TYPE,NUMERIC_PRECISION,COLUMN_KEY,COLUMN_DEFAULT
		FROM information_schema.COLUMNS
		WHERE table_schema = DATABASE()  AND TABLE_NAME = '%s'`, table,
		),
	)

	if err != nil {
		log.Printf("table rows is nil with table:%s error: %v \n", table, err)
		return columns
	}

	if rows == nil {
		log.Printf("rows is nil with table:%s \n", table)
		return columns
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var maxLength, numberPrecision []byte
		var t = ""

		col := helper.Column{}
		err = rows.Scan(
			&col.ColumnName, &col.Type, &t, &col.TableName, &col.ColumnComment, &maxLength, &col.ColumnType, &numberPrecision,
			&col.ColumnKey, &col.Default,
		)
		col.Nullable = t
		if maxLength != nil {
			col.MaxLength = Byte2Int64(maxLength)
		}
		if numberPrecision != nil {
			col.NumberPrecision = Byte2Int64(numberPrecision)
		}
		if err != nil {
			log.Println(err.Error())
			continue
		}
		columns = append(columns, col)
	}

	return columns
}

func Byte2Int64(data []byte) int64 {
	var str string
	var ret int64 = 0
	for i := 0; i < len(data); i++ {
		str += string(data[i])
	}
	ret, _ = strconv.ParseInt(str, 10, 64)
	return ret
}
