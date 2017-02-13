package utils

import (
	"errors"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

// StringToDate 日付に変換する
func StringToDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// GetAppPath アプリケーションパスを取得する
func GetAppPath() string {
	_, f, _, _ := runtime.Caller(1)
	p, _ := filepath.Abs(filepath.Dir(filepath.Join(f, ".."+string(filepath.Separator))))

	return p
}

// GetArrayCombile 配列をkeyと結合させる
func GetArrayCombile(k []string, v []string) (map[string]string, error) {
	m := map[string]string{}

	if len(k) != len(v) {
		return m, errors.New("Both parameters should have an equal number of elements")
	}

	for index, key := range k {
		m[key] = v[index]
	}

	return m, nil
}

// DbValueToMap Value型をマップに変換する
func DbValueToMap(e reflect.Value) map[string]interface{} {
	r := make(map[string]interface{})
	size := e.NumField()

	for i := 0; i < size; i++ {
		name := e.Type().Field(i).Name
		if e.Type().Field(i).Tag.Get("json") != "" {
			name = e.Type().Field(i).Tag.Get("json")
		}
		r[name] = e.Field(i).Interface()
	}

	return r
}

// DbStructToMap DB構造体からマップに変換する
func DbStructToMap(s interface{}) map[string]interface{} {
	return DbValueToMap(reflect.ValueOf(s).Elem())
}

// DbStructListToMapList DB構造体リストからマップリストに変換する
func DbStructListToMapList(s interface{}) (r []map[string]interface{}) {
	size := reflect.ValueOf(s).Len()

	for i := 0; i < size; i++ {
		r = append(r, DbValueToMap(reflect.ValueOf(s).Index(i)))
	}

	return r
}

// ValueToMap Value型をマップに変換する
func ValueToMap(e reflect.Value) map[string]interface{} {
	r := make(map[string]interface{})
	size := e.NumField()

	for i := 0; i < size; i++ {
		r[e.Type().Field(i).Name] = e.Field(i).Interface()
	}

	return r
}

// StructToMap 構造体からマップに変換する
func StructToMap(s interface{}) map[string]interface{} {
	return ValueToMap(reflect.ValueOf(s).Elem())
}

// StructListToMapList 構造体リストからマップリストに変換する
func StructListToMapList(s interface{}) (r []map[string]interface{}) {
	size := reflect.ValueOf(s).Len()

	for i := 0; i < size; i++ {
		r = append(r, ValueToMap(reflect.ValueOf(s).Index(i)))
	}

	return r
}
