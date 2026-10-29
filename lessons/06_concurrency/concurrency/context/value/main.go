// Sample program showing how to store and retrieve
// values from a context.
package main

import (
	"context"
	"fmt"
)

// TraceID identifies one request across service boundaries.
type TraceID string

type traceIDKey struct{}

func main() {
	traceID := TraceID("f47ac10b-58cc-0372-8567-0e02b2c3d479")
	key := traceIDKey{}
	ctx := context.WithValue(context.Background(), key, traceID)

	if uuid, ok := ctx.Value(key).(TraceID); ok {
		fmt.Println("TraceID:", uuid)
	}

	if _, ok := ctx.Value(0).(TraceID); !ok {
		fmt.Println("TraceID Not Found")
	}
}
