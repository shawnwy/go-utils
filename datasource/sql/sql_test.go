package sql

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	mysqlURI := "root:sBL2y7Uuxqyi@tcp(10.10.10.201:53306)/dg_system?charset=utf8&parseTime=True&loc=Local"
	db := NewMySQL(mysqlURI, NewGormCFG(),
		WithMaxIdleConns(10),
		WithMaxOpenConns(100))
	rows, err := db.Exec("show tables").Rows()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for rows.Next() {
		var x interface{}
		err := rows.Scan(&x)
		if err != nil {
			return
		}
		fmt.Println(x)
	}
}
