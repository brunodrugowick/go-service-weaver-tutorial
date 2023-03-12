package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/ServiceWeaver/weaver"
)

type Sleeper interface {
	Sleep(ctx context.Context) error
}

type sleeper struct {
	weaver.Implements[Sleeper]
}

func (s *sleeper) Sleep(_ context.Context) error {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(500)
	time.Sleep(time.Duration(r) * time.Millisecond)
	s.Logger().Info("I slept for %d", r)
	return nil
}
