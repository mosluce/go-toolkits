package toolkits

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    //
	_ "github.com/jinzhu/gorm/dialects/mysql"    //
	_ "github.com/jinzhu/gorm/dialects/postgres" //
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //
	"github.com/mosluce/go-toolkits/database"
)

func Open(config database.ConnectionConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(config.Dialect, config.URI())
	return
}
