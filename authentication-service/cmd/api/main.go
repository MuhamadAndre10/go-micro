package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "80"


// var count int64 

type Config struct {
	DB *sql.DB
    Models data.Models
}

func main() {
    log.Println("Starting authentication service...")

    // TODO: connect to DB
    db := openDB()

    // todo: set up config
    app := Config{
        DB: db,
        Models: data.New(db),
    }

    srv := &http.Server{
        Addr: fmt.Sprintf(":%s", webPort),
        Handler: app.routes(),
    }

    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

func openDB() *sql.DB {   
    dsn := os.Getenv("DSN")
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatal("failed to open database connection")
    }

    if err = db.Ping(); err != nil {
        log.Fatal("failed to verify database connection")
    }

    return db
}

// func connectToDB() *sql.DB {
//     dsn := os.Getenv("DSN")

//     for {
//         connection, err := openDB(dsn)
//         if err != nil {
//             log.Println("Postgres not yet ready")
//             count++
//         } else {
//             log.Println("connected to Postgres")
//             return connection
//         }

//         if count > 10 {
//             log.Println(err)
//             return nil
//         }

//         log.Println("Backing off for two second")
//         time.Sleep(2 * time.Second)
//         continue
//     }
// }
