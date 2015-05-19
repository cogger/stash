package stash

import "golang.org/x/net/context"

var stashkey = "stash"

func c(ctx context.Context) map[string]interface{} {
	value, ok := ctx.Value(&stashkey).(*map[string]interface{})
	if !ok {
		value = &map[string]interface{}{}
	}
	return *value
}
