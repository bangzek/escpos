package escpos_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEscpos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Escpos Suite")
}
