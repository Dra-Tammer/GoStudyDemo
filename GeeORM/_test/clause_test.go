package _test

import (
	Clause "geeorm/clause"
	"reflect"
	"testing"
)

func testSelect(t *testing.T) {
	var clause Clause.Clause
	clause.Set(Clause.LIMIT, 3)
	clause.Set(Clause.SELECT, "User", []string{"*"})
	clause.Set(Clause.WHERE, "Name = ?", "Tom")
	clause.Set(Clause.ORDERBY, "Age ASC")
	sql, vars := clause.Build(Clause.SELECT, Clause.WHERE, Clause.ORDERBY, Clause.LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}
