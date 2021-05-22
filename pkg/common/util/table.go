package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func ToTable(object interface{}) (map[string]interface{}, error) {
	table := make(map[string]interface{})
	t := reflect.TypeOf(object)
	v := reflect.ValueOf(object)
	if t.Kind().String() != "struct" {
		return nil, errors.New("only apply to struct")
	}
	for i := 0; i < t.NumField(); i++ {
		segs := strings.Split(t.Field(i).Tag.Get("table"), ",")
		if len(segs) == 0 {
			continue
		}
		tagName := segs[0]
		if v.Field(i).IsNil() {
			if StringInSlice("nullable", segs) {
				table[tagName] = nil
			}
			continue
		}
		table[tagName] = v.Field(i).Interface()
	}
	return table, nil
}

func ToColumns(object interface{}, exclude map[string]interface{}) []string {
	columns := make([]string, 0)
	t := reflect.TypeOf(object)
	if t.Kind().String() != "struct" {
		panic(errors.New("only apply to struct"))
	}
	for i := 0; i < t.NumField(); i++ {
		segs := strings.Split(t.Field(i).Tag.Get("table"), ",")
		if len(segs) == 0 || segs[0] == "" {
			continue
		}
		tagName := segs[0]
		if _, ok := exclude[tagName]; ok {
			continue
		}
		columns = append(columns, tagName)
	}
	return columns
}

func ToValues(object interface{}, exclude map[string]interface{}) []interface{} {
	values := make([]interface{}, 0)
	t := reflect.TypeOf(object)
	if t.Kind().String() != "struct" {
		panic(errors.New("only apply to struct"))
	}
	v := reflect.ValueOf(object)
	for i := 0; i < t.NumField(); i++ {
		segs := strings.Split(t.Field(i).Tag.Get("table"), ",")
		if len(segs) == 0 || segs[0] == "" {
			continue
		}
		tagName := segs[0]
		if _, ok := exclude[tagName]; ok {
			continue
		}
		values = append(values, v.Field(i).Interface())
	}
	return values
}

func TableTagFieldMap(obj interface{}) map[string]string {
	res := make(map[string]string)
	t := reflect.TypeOf(obj)
	if t.Kind().String() != "struct" {
		panic(errors.New("only apply to struct"))
	}
	for i := 0; i < t.NumField(); i++ {
		segs := strings.Split(t.Field(i).Tag.Get("table"), ",")
		if len(segs) == 0 {
			continue
		}
		tagName := segs[0]
		if tagName == "" || tagName == "-" {
			continue
		}
		res[tagName] = t.Field(i).Name
	}
	return res
}

func ToScanAll(columnNames []string, obj interface{}, tagNameMap map[string]string) ([]interface{}, error) {
	pointers := make([]interface{}, len(columnNames))
	//t := reflect.TypeOf(obj)
	//if t.Kind().String() != "struct" {
	//	panic(errors.New("only apply to struct"))
	//}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("obj not ptr"))
	}
	v = v.Elem()
	for i, colName := range columnNames {
		fieldName := colName
		if name, ok := tagNameMap[colName]; ok {
			fieldName = name
		}
		fieldV := v.FieldByName(fieldName)
		if !fieldV.IsValid() {
			return nil, fmt.Errorf("field %s not valid for obj: %v", colName, obj)
		}
		pointers[i] = fieldV.Addr().Interface()
	}
	return pointers, nil
}
