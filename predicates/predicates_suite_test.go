package predicates

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPredicates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Predicates Test Suite")
}

var _ = BeforeSuite(func() {
	By("tearing down the test environment")
})
