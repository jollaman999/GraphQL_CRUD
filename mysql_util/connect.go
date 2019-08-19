package mysql_util

import (
	"../config"
	"../logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Prepare() (error) {
	var err error = nil
	Db, err = sql.Open("mysql", config.Mysql_Id + ":" + config.Mysql_Password + "@tcp(" +
		config.Mysql_Address + ":" + config.Mysql_Port + ")/" + config.Mysql_Database + "?parseTime=true")
	if err != nil {
		logger.Logger.Println(err)
		return err
	} else {
		logger.Logger.Println("db is connected")
	}

	err = Db.Ping()
	if err != nil {
		logger.Logger.Println(err.Error())
		return err
	}

	return nil
}