package main

import (
	"fmt"
	"verbme-api/db"
	"verbme-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Closing database connection...")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Diary{}, &model.Word{})
}
