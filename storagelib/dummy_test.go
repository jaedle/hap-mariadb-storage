package storagelib_test

import (
	"github.com/jaedle/hap-mariadb-storage/storagelib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dummy", func() {
	It("returns true", func() {
		Expect(storagelib.Dummy()).To(BeTrue())
	})
})
