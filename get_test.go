package stash

import (
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get", func() {
	It("should return nil when the item does not exist and no default value is provided", func() {
		Expect(Get(context.Background(), "key")).To(BeNil())
	})

	It("should return the value set to it", func() {
		ctx := context.Background()
		ctx = Set(ctx, "key", "bob")
		value, ok := Get(ctx, "key").(string)
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal("bob"))
	})

	It("should return the value from the default value if it does not exist", func() {
		ctx := context.Background()
		value, ok := Get(ctx, "key", func() interface{} {
			return "bob"
		}).(string)
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal("bob"))
	})
})
