package sql_test

import "testing"
import "github.com/astaxie/beedb"

func TestBeeDB(t *testing.T) {
	beedb.OnDebug = true
	orm := beedb.New(db)
	orm.Save(User{Id: 3, Name: "smark"})

}

type User struct {
	Id   int `PK`
	Name string
}
