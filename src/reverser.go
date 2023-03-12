package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
	"math/rand"
	"time"
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
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	// Get a client to the Reverser component.
	// Looks like a dependency injection...
	slpr, err := weaver.Get[Sleeper](r)
	if err != nil {
		log.Fatal(err)
	}

	blbr, err := weaver.Get[Blaber](r)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(500)
	if intn%2 == 0 {
		slpr.Sleep(ctx)
	} else {
		blbr.Bla(ctx, fmt.Sprintf("You're now going into Blaber, %s. Don't thank me!", s))
	}

	return string(runes), nil
}
