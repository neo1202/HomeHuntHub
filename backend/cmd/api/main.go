package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	// DB     *sql.DB
	DB repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line. flag可以讓你指定你的db位置, DSN就是指出postgres之位置
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	//初始化這個struct時,把conn傳過去,那邊收到的是conn的位置
	//所以可以直接conn.Close, 他們使用的是同一個conn  ->defer conn.Close() 
	defer app.DB.Connection().Close() //確認釋放資源當main結束 //Connection是我們自訂的函式

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
