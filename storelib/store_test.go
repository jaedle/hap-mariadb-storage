package storelib_test

import (
	"database/sql"
	"github.com/brutella/hap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jaedle/hap-mariadb-storage/storelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const testDatasource = "root:password@tcp(localhost:3307)/database"

var db *sql.DB

var _ = Describe("Store", func() {

	BeforeEach(func() {
		var err error
		db, err = sql.Open("mysql", testDatasource)
		Expect(err).NotTo(HaveOccurred())

		Expect(db.Ping()).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if db != nil {
			_ = db.Close()
		}
	})

	It("is assignable to hap store", func() {
		var store hap.Store = storelib.New(nil, "hap")
		Expect(store).NotTo(BeNil())
	})

	It("initialises", func() {
		store := storelib.New(db, "table")
		Expect(store.Init()).NotTo(HaveOccurred())
	})

	It("initialises", func() {
		store := storelib.New(db, "table")
		Expect(store.Init()).NotTo(HaveOccurred())

		Expect(store.Set("a-key", []byte("asdf"))).NotTo(HaveOccurred())

		_, err := store.Get("a-key")

		Expect(err).NotTo(HaveOccurred())
	})

})
