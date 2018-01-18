package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    //
	_ "github.com/jinzhu/gorm/dialects/mysql"    //
	_ "github.com/jinzhu/gorm/dialects/postgres" //
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //
)

type ConnectionConfig struct {
	Dialect  string
	Filepath string //sqlite only
	Host     string
	User     string
	Pass     string
	Dbname   string
}

func (cc ConnectionConfig) URI() (uri string) {

	if cc.Dialect == "mysql" {
		uri = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", cc.User, cc.Pass, cc.Host, cc.Dbname)
	} else if cc.Dialect == "postgres" {
		uri = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cc.Host, cc.User, cc.Dbname, cc.Pass)
	} else if cc.Dialect == "sqlite3" {
		uri = cc.Filepath
	} else if cc.Dialect == "mssql" {
		uri = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", cc.User, cc.Pass, cc.Host, cc.Dbname)
	}

	return
}

func Open(config ConnectionConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(config.Dialect, config.URI())
	return
}
