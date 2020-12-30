package dao

import (
	"fmt"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stock.tao/util"
)

const (
	dsn = "root:root@tcp(127.0.0.1:3306)/doc?charset=utf8mb4&parseTime=True&loc=Local"
)

// Connect database source
func Connect() *gorm.DB {
	fmt.Printf("Open datasource %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Open %s error:%v\n", dsn, err)
		panic(err.Error())
	}
	return db
}

// Connect database source
func Connect(dsn string) *gorm.DB {
	fmt.Printf("Open datasource %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Open %s error:%v\n", dsn, err)
		panic(err.Error())
	}
	return db
}

// Condition => build the SQL condition, the property of model must be the Pointer Type
func Condition(model interface{}) map[string]interface{} {
	c := make(map[string]interface{})
	t, v := reflect.TypeOf(model), reflect.ValueOf(model)
	if t.NumField() == 0 || v.NumField() == 0 {
		return c
	}
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}
		f := t.Field(i)
		if info, ok := f.Tag.Lookup("gorm"); ok {
			if colName, ok := util.GetTagPart(info, "column"); ok {
				c[colName] = util.GetValue(v.Field(i))
			}
		}
	}
	return c
}
