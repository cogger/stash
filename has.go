package stash

import "golang.org/x/net/context"

func Has(ctx context.Context, key string) bool {
	_, ok := c(ctx)[key]
	return ok
}
