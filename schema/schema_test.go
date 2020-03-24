// schema
// @Time : 2020/3/24 9:57
// @Author : Jeffery Sun
// @File : schema_test
// @Software: GoLand

package schema

import (
	"github.com/gofromzero/geeorm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("Failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
