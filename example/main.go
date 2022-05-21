package main

import (
	"context"
	"database/sql"
	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jaedle/hap-mariadb-store/storelib"
	"log"
)

func main() {
	// establish connection to database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
	if err != nil {
		panic(err)
	}

	// ensure connection is established
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// instantiate store with
	store := storelib.New(storelib.Configuration{
		Db:    db,
		Table: "hap-example-table",
	})

	// ensure table is instantiated
	if err := store.Init(); err != nil {
		panic(err)
	}

	// create dummy server containing a bridge to test pairing
	server, err := hap.NewServer(store, accessory.NewBridge(accessory.Info{
		Name: "bridge",
	}).A)
	if err != nil {
		panic(err)
	}

	log.Fatal(server.ListenAndServe(context.Background()))

}
