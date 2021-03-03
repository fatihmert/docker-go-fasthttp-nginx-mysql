package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var MYSQL_USER string
var MYSQL_PASSWORD string
var MYSQL_DATABASE string

var DB *sql.DB

func prepareDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
}

func initDatabase() {
	var err error
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", MYSQL_USER, MYSQL_PASSWORD, "127.0.0.1", "3306", MYSQL_DATABASE)

	DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("Failed connect db")
	}

	fmt.Println("Connection Opened to db")
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello from docker!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi, %s!\n", ctx.UserValue("name"))
}

func main() {
	log.SetOutput(os.Stdout)

	prepareDotEnv()
	initDatabase()

	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/db", func(ctx *fasthttp.RequestCtx){
		id := 1
		var col string
		sqlStatement := `SELECT start FROM test WHERE id=$1`
		row := DB.QueryRow(sqlStatement, id)
		err := row.Scan(&col)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(ctx, "Zero rows found")
			} else {
				fmt.Fprintf(ctx, "err: %s", err)
			}
		}
	});

	log.Fatal(fasthttp.ListenAndServe(":3000", r.Handler))
}
