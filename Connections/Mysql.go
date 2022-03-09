package connections

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQLDatabase() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := "1234"
	Name := "system_backend_go"
	connection, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+Name+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return connection
}
