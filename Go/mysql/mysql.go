package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	user = "root"
	pass = ""
	host = "127.0.0.1"
	port = 3306
	db   = ""
)

func init() {
	name, _ := os.Hostname()
	if name == "hyena" {
		pass = ""
	}
}

func ConnectByGorm() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)
	return gorm.Open("mysql", dns)
}

func ConnectByRaw() (*sql.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)
	return sql.Open("mysql", dns)
}
