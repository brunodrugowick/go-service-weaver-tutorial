package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

// Reverser component - this is just an interface, it does nothing for now
type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser] // Right... looks like this is where the "magic" is.
}

func (r *reverser) Reverse(ctx context.Context, s string) (string, error) {
	fmt.Println(ctx)
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
