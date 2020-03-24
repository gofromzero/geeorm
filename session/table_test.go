// session
// @Time : 2020/3/24 10:14
// @Author : Jeffery Sun
// @File : table_test
// @Software: GoLand

package session

import (
	_ "github.com/gofromzero/geeorm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
