package respository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func dbConn() (db *gorm.DB) {

	dbDriver := "mysql"    // Database driver
	dbUser := "root"       // Mysql username
	dbPass := "111111"     // Mysql password
	dbName := "golangidev" // Mysql schema

	// Realize the connection with mysql driver
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8")

	// If error stop the application
	if err != nil {
		panic(err.Error())
		log.Println("Db can not connect")
		fmt.Println("Db can not connect")
	}

	// Return db object to be used by another functions
	return db
}

type BaseRepo interface {
	QueryForList(sql string) (res []interface{})
	QueryForOne(sql string) (res interface{})
	GetTotalUsers(sql string) (res interface{})
	GetTotalUsersByPaging(current_page int, per_page_num int) (res interface{})
	SearchUsersByPaging(current_page int, per_page_num int) (res interface{})
	GetUserById(id int) (res interface{})
	Save(obj interface{}) (result int64)
	Update(obj interface{}) (result int64)
	Delete(obj interface{}) (result int64)
	Remove(obj interface{}) (result int64)
}
