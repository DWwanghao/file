package mysql
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"os"
)

var db *sql.DB
func init() {
	fmt.Println(1)
	db, _ = sql.Open("mysql", "root:123456@tcp(192.168.2.2:3306)/fileserver?charset=utf8")

	fmt.Println(2)
	db.SetMaxOpenConns(1000)
	fmt.Println(3)
	err := db.Ping()
	if err != nil {
		fmt.Println("failed connect"+err.Error())
		os.Exit(1)
	}

}
func DBConn() *sql.DB  {
	return db
}
