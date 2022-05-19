package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

const waitBetweenPings = 500 * time.Millisecond
const maxRetries = 100
const testDatasource = "root:password@tcp(localhost:3307)/database"

func main() {
	fmt.Println("Waiting for database to startup")

	retries := 0
	for {
		if pingDatabase() == nil {
			return
		}
		retries++

		if retries > maxRetries {
			fmt.Println("database did not get healthy, exiting")
			os.Exit(1)
		}

		time.Sleep(waitBetweenPings)
	}

}

func pingDatabase() error {
	var err error
	db, err := sql.Open("mysql", testDatasource)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return db.PingContext(ctx)
}
