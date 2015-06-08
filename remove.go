package stash

import "golang.org/x/net/context"

func Remove(ctx context.Context, key string) context.Context {
	values := c(ctx)
	if _, ok := values[key]; ok {
		delete(values, key)
	}

	return context.WithValue(ctx, &stashkey, &values)
}
