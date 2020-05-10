package utils

import (
	"testing"
)

type SelectResult struct {
	A int `db:"result1"`
	B int `db:"result2"`
}

func Test_Taging(t *testing.T) {
	a := SelectResult{}
	err := GetDB().Get(&a, "select 1 as result1, 2 as result2;")
	if err != nil {
		t.Error("db select error")
	}

	if a.A != 1 {
		t.Error("result1 error")
	}

	if a.B != 2 {
		t.Error("result2 error")
	}
}

func Test_JSON(t *testing.T) {
	conn := connectionJSON{}
	err := conn.read()
	if err != nil {
		t.Error(err)
	}
}
