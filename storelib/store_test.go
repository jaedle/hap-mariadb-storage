package storelib_test

import (
	"github.com/brutella/hap"
	"github.com/jaedle/hap-mariadb-storage/storelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	It("is assignable to hap store", func() {
		var store hap.Store = storelib.New()
		Expect(store).NotTo(BeNil())
	})
})
