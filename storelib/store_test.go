package storelib_test

import (
	"database/sql"
	"github.com/brutella/hap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jaedle/hap-mariadb-storage/storelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const aKey = "a-key"
const anotherKey = "another-key"

const aStringValue = "value"
const anotherStringValue = "another-value"

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
		var store hap.Store = storelib.New(nil, uuid.NewString())
		Expect(store).NotTo(BeNil())
	})

	It("initialises", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())
	})

	It("can reinitalise", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())

		Expect(store.Init()).NotTo(HaveOccurred())
	})

	It("persists string", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())

		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal([]byte(aStringValue)))
	})

	It("overwrites previous value", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())
		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		Expect(store.Set(aKey, []byte(anotherStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal([]byte(anotherStringValue)))
	})

	It("returns error if key not found", func() {
		store := storelib.New(db, uuid.NewString())
		Expect(store.Init()).NotTo(HaveOccurred())
		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(anotherKey)

		Expect(err).To(HaveOccurred())
		Expect(val).To(BeNil())
	})

})
