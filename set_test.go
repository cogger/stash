package stash_test

import (
	"fmt"

	. "github.com/cogger/stash"
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type name struct {
	First string
	Last  string
}

var tests = []struct {
	Label string
	Value interface{}
	Test  func(interface{})
}{
	{
		Label: "string",
		Value: "bob",
		Test: func(value interface{}) {
			item, ok := value.(string)
			Expect(ok).To(BeTrue())
			Expect(item).To(Equal("bob"))
		},
	}, {
		Label: "int",
		Value: 5,
		Test: func(value interface{}) {
			item, ok := value.(int)
			Expect(ok).To(BeTrue())
			Expect(item).To(Equal(5))
		},
	}, {
		Label: "float64",
		Value: 1.1,
		Test: func(value interface{}) {
			item, ok := value.(float64)
			Expect(ok).To(BeTrue())
			Expect(item).To(Equal(1.1))
		},
	}, {
		Label: "struct",
		Value: name{
			First: "Sean",
			Last:  "Dolphin",
		},
		Test: func(value interface{}) {
			item, ok := value.(name)
			Expect(ok).To(BeTrue())
			Expect(item).To(BeEquivalentTo(name{
				First: "Sean",
				Last:  "Dolphin",
			}))
		},
	},
}

var _ = Describe("Set", func() {
	for _, t := range tests {
		test := t
		It(fmt.Sprintf("should set the %s value to the one provided.", test.Label), func() {
			ctx := Set(context.Background(), "key", test.Value)
			test.Test(Get(ctx, "key"))
		})
	}
})
