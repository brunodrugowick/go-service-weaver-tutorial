package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

type Blaber interface {
	Bla(ctx context.Context, blablabla string) error
}

type blaber struct {
	weaver.Implements[Blaber]
}

func (s *blaber) Bla(_ context.Context, blablabla string) error {
	s.Logger().Info("Blaber says: %s", blablabla)
	return nil
}
