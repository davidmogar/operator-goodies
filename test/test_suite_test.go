package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test utils Test Suite")
}

var _ = BeforeSuite(func() {
	By("tearing down the test environment")
})
