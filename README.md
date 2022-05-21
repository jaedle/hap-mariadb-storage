# HAP-MariaDB-Store

This is a store adapter using [MariaDB](https://mariadb.org/) for [hap](https://github.com/brutella/hap). It also supports [MySQL](https://www.mysql.com/).

## Supported Databases

The following databases are supported and performed integration tests against.

### MariaDB

- 10.3
- 10.4
- 10.5
- 10.6
- 10.7

### Mysql

- 5.6
- 5.7
- 8.0


## Usage

1. Add the library to your project dependencies: `go get -u github.com/jaedle/hap-mariadb-store`.
2. Configure your [database handle](https://pkg.go.dev/database/sql#DB).
3. Instantiate a new store instance
4. Initialize the store (creates database table on first run)

As the store only uses the database handle, please make sure to add any needed configuration before passing it to the store.

### Example

```go
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
```