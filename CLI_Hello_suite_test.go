package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCLIHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CLIHello Suite")
}
