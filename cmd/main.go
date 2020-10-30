package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/my-Sakura/SMService/controller"
	"github.com/my-Sakura/SMService/utils"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/mysql")
	if err != nil {
		panic(err)
	}

	appCode := os.Getenv("AppCode")

	config := utils.NewConfig(appCode)

	s := controller.NewSMSController(db, config)
	r.Use(utils.Cors())
	s.RegistRouter(r)

	r.Run(":8000")
}
