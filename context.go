package stash

import (
	"sync"

	"golang.org/x/net/context"
)

var stashkey = "stash"

func c(ctx context.Context) (context.Context, *values) {
	value, ok := ctx.Value(&stashkey).(*values)
	if !ok {
		value = newValues()
		ctx = context.WithValue(ctx, &stashkey, value)
	}
	return ctx, value
}

type values struct {
	sync.RWMutex
	values map[string]interface{}
}

func newValues() *values {

	return &values{
		values: map[string]interface{}{},
	}

}
