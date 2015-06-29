package stash

import (
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Has", func() {
	It("should return false when it can not find the value", func() {
		ctx := context.Background()
		Expect(Has(ctx, "test")).To(BeFalse())
	})

	It("should return true when it can find the value", func() {
		ctx := context.Background()
		ctx = Set(ctx, "test", "blob")
		Expect(Has(ctx, "test")).To(BeTrue())
	})
})
