package vodka

import (
	"golang.org/x/sync/errgroup"
)

// Service is vodka service
type Service struct {
	opts     *Options
	errGroup *errgroup.Group
}

// Option Service option func
type Option func(*Options)

func NewService(opts ...Option) *Service {
	return &Service{}
}

func (s *Service) Start() error {
	for _, srv := range s.opts.Server {
		s.errGroup.Go(srv.Start)
	}

	return s.errGroup.Wait()
}
