package main

import (
	"./graphql_util"
	"./logger"
	"./mysql_util"
	"net/http"
)

func main() {
	if !logger.Prepare() {
		return
	}
	defer logger.FpLog.Close()

	err := mysql_util.Prepare()
	if err != nil {
		return
	}
	defer mysql_util.Db.Close()

	http.Handle("/graphql", graphql_util.Graphql_handler)

	logger.Logger.Println("Server is running on port 8001")
	err = http.ListenAndServe(":8001", nil)
	if err != nil {
		logger.Logger.Println("Failed to prepare http server!")
	}
}