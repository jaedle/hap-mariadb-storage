package dummy_test

import (
	"github.com/jaedle/hap-mariadb-storage/dummy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dummy", func() {
	It("returns true", func() {
		Expect(dummy.Dummy()).To(BeTrue())
	})
})
