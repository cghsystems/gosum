package matcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSummod(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Matcher")
}
