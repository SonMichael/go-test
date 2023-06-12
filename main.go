package main

import (
	"fmt"
	"go-test/module/item/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	host, _ := os.LookupEnv("DB_HOST")
	fmt.Println(host)
	dsn := "root:@tcp(127.0.0.1:4307)/lixibox"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/items", transport.HandleCreateItem(db))
	r.Run()
}