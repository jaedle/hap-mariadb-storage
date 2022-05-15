package storelib_test

import (
	"database/sql"
	"github.com/brutella/hap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jaedle/hap-mariadb-storage/storelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

const aKey = "a-key"
const anotherKey = "another-key"

const aStringValue = "value"
const anotherStringValue = "another-value"

const testDatasource = "root:password@tcp(localhost:3307)/database"

const fiveMegabytes = 5 * 1024 * 1024

// test own store implementation side by side with original (fs-store) implementation
// using the store interface

var _ = Describe("StoreComparison", func() {

	DescribeTable("persists string", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal([]byte(aStringValue)))
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("persists string", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()
		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		Expect(store.Set(aKey, []byte(anotherStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal([]byte(anotherStringValue)))
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("returns error if key not found", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set(aKey, []byte(aStringValue))).NotTo(HaveOccurred())

		val, err := store.Get(anotherKey)

		Expect(err).To(HaveOccurred())
		Expect(val).To(BeNil())
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("saves big content", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set(aKey, binaryContent(fiveMegabytes))).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal(binaryContent(fiveMegabytes)))
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("deletes key", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set(aKey, binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Delete(aKey)).NotTo(HaveOccurred())

		val, err := store.Get(aKey)

		Expect(err).To(HaveOccurred())
		Expect(val).To(BeNil())
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("fails to delete non existing key", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set(aKey, binaryContent(1))).NotTo(HaveOccurred())

		Expect(store.Delete(anotherKey)).To(HaveOccurred())
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("lists keys bey suffix", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set("test.suffix", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set(".suffix", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set(".suffix-1", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set("1.suffix", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set("very-long.suffix", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set("not.sufix.not", binaryContent(1))).NotTo(HaveOccurred())
		Expect(store.Set("unrelated", binaryContent(1))).NotTo(HaveOccurred())

		Expect(store.KeysWithSuffix(".suffix")).To(ConsistOf([]string{
			"test.suffix", ".suffix", "1.suffix", "very-long.suffix",
		}))
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

	DescribeTable("returns empty list on no keys", func(ts testStore) {
		defer ts.Cleanup()
		Expect(ts.Init()).NotTo(HaveOccurred())
		store := ts.Storage()

		Expect(store.Set("test.suffix", binaryContent(1))).NotTo(HaveOccurred())

		Expect(store.KeysWithSuffix(".unknown")).To(BeNil())
	},
		Entry("for mariadb-store", &mariaDbStore{}),
		Entry("for fs-store", &fsStore{}),
	)

})

func binaryContent(size int) []byte {
	var result []byte

	for i := 0; i < size; i++ {
		result = append(result, 0x27)
	}

	return result
}

type testStore interface {
	Init() error
	Cleanup()
	Storage() hap.Store
}

type mariaDbStore struct {
	db      *sql.DB
	storage *storelib.MariaDbStore
}

func (m *mariaDbStore) Init() error {
	var err error
	db, err := sql.Open("mysql", testDatasource)
	if err != nil {
		return err
	}
	m.db = db

	if err := db.Ping(); err != nil {
		return err
	}

	store := storelib.New(db, uuid.NewString())
	if err := store.Init(); err != nil {
		return err
	}

	m.storage = store

	return nil
}

func (m *mariaDbStore) Cleanup() {
	if m.db != nil {
		_ = m.db.Close()
	}
}

func (m *mariaDbStore) Storage() hap.Store {
	return m.storage
}

type fsStore struct {
	dir     string
	storage hap.Store
}

func (f *fsStore) Init() error {
	dir, err := ioutil.TempDir(os.TempDir(), "fs-store")
	if err != nil {
		return err
	}

	f.dir = dir
	f.storage = hap.NewFsStore(f.dir)

	return nil
}

func (f *fsStore) Cleanup() {
	if f.dir != "" {
		_ = os.RemoveAll(f.dir)
	}
}

func (f *fsStore) Storage() hap.Store {
	return f.storage
}
