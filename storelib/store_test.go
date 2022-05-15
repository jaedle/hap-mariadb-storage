package storelib_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jaedle/hap-mariadb-storage/storelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("specific tests for maria-db-storage", func() {

	var db *sql.DB

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

	It("reinitialize", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())
	})

	It("reinitializes", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())

		Expect(store.Init()).NotTo(HaveOccurred())
	})
})
