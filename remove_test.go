package stash

import (
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Remove", func() {
	It("should remove the key provided", func() {
		ctx := Set(context.Background(), "test", "blah")
		Expect(Has(ctx, "test")).To(BeTrue())

		ctx = Remove(ctx, "test")
		Expect(Has(ctx, "test")).To(BeFalse())
	})
})
