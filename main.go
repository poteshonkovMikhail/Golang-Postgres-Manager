package main

import (
	"example.com/m/db"
	//"example.com/m/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4"
)

func main() {

	r := gin.Default()
	r.POST("/connectDB", db.ConnectDB)
	r.POST("/workDB", db.DatabaseHandler)
	r.POST("/create-database", db.CreateDatabase)
	r.POST("/sql-request", db.ExecuteQuery)
	r.POST("/db-info", db.InterfaceGetDbInformation)
	//r.GET("/items", handlers.GetItems)    // Получение списка элементов
	r.GET("/initDBlist", db.InitDBlist)
	r.Run(":8080") // Запускаем сервер на порту 8080
}
