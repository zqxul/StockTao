package dao

import (
	"fmt"
	"reflect"

	"stock.tao/util"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

const (
	dsn = "root:root@tcp(127.0.0.1:3306)/doc?charset=utf8mb4&parseTime=True&loc=Local"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Printf("Create db engine error:%v\n", err)
		panic(err)
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)
}

// SQLCondition ==> sql condition inteface
type SQLCondition interface {
	Condition() map[string]interface{}
}

// Condition => build the SQL condition, the property of model must be the Pointer Type
func Condition(model SQLCondition) map[string]interface{} {
	c := make(map[string]interface{})
	t, v := reflect.TypeOf(model), reflect.Indirect(reflect.ValueOf(model))
	if t.NumField() == 0 || v.NumField() == 0 {
		return c
	}
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}
		f := t.Field(i)
		if column, ok := f.Tag.Lookup("column"); ok {
			c[column] = util.GetValue(v.Field(i))
		}
	}
	return c
}
