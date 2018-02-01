package database

import "fmt"

type ConnectionConfig struct {
	Dialect  string
	Filepath string //sqlite only
	Host     string
	User     string
	Pass     string
	Dbname   string
}

func (cc ConnectionConfig) URI() (uri string) {

	if cc.Dialect == MYSQL {
		uri = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", cc.User, cc.Pass, cc.Host, cc.Dbname)
	} else if cc.Dialect == POSTGRES {
		uri = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cc.Host, cc.User, cc.Dbname, cc.Pass)
	} else if cc.Dialect == SQLITE {
		uri = cc.Filepath
	} else if cc.Dialect == MSSQL {
		uri = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", cc.User, cc.Pass, cc.Host, cc.Dbname)
	}

	return
}
