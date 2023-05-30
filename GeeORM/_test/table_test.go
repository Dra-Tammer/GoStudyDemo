package _test

import (
	GeeORM "geeorm"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	engine, _ := GeeORM.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
