package stash

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStash(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stash Suite")
}
