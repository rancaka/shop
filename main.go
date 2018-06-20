package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rancaka/shop/config"
	"github.com/rancaka/shop/db"
	"github.com/rancaka/shop/src/handlers"
)

func main() {

	conf := flag.String("conf", "/etc/shop/production.json", "Config file")
	flag.Parse()

	cfg, err := config.New(*conf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully Get config from %v!\n", *conf)

	sqlxDB, err := db.NewSQL(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlxDB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully connected to DB!\n")

	redisConn, err := db.NewRedis(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully connected to Redis!\n")

	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, err := os.Create("/var/log/shop/shop.gin.log")
		if err != nil {
			log.Fatalln(err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	handler := handlers.New(cfg, sqlxDB, redisConn)

	r := gin.Default()
	r.GET("/", handler.Index)
	r.Run()
}
