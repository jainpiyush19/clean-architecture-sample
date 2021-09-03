package user

import "context"

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key

// NewContext returns a new Context that carries value userID.
func NewContext(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userKey, &userID)
}

// FromContext returns the userID value stored in ctx, if any.
func FromContext(ctx context.Context) (*int64, bool) {
	d, ok := ctx.Value(userKey).(*int64)
	return d, ok
}
