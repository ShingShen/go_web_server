package operator

import (
	"encoding/json"
	"fmt"
	"os"
	"server/utils/sqloperator"
)

func IfElse(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func LoadJson(jsonType interface{}, filename string) (interface{}, error) {
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()
	if err != nil {
		return jsonType, err
	}
	jsonParser := json.NewDecoder(jsonFile)
	err = jsonParser.Decode(&jsonType)
	return jsonType, err
}

func CreatingDataList(res sqloperator.ISqlRows) ([]byte, error) {
	columns, err := res.Columns()
	if err != nil {
		fmt.Printf("Failed to query table's column, err: %v\n", err)
		return nil, err
	}
	fmt.Println("table's columns: ", columns)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for res.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		res.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			value := values[i]
			b, ok := value.([]byte)
			if ok {
				v = string(b)
			} else {
				v = value
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	tableJsonData, _ := json.Marshal(tableData)
	// fmt.Println("tableJsonData: ", string(tableJsonData))
	return tableJsonData, nil
}
