package sql_test

import (
	"testing"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func TestMysql(t *testing.T) {

	stmt, err := db.Prepare("insert into user values(?,?)")
	checkErr(err)

	stmt.Exec(1, "test")
	//log.Println(rs.LastInsertId())
	checkErr(err)

}
