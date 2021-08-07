package server

import (
	"context"
	"time"
)

type Options struct {
	Address      string
	CloseWait time.Duration
	Ctx       context.Context
}

func NewOptions() *Options {
	return &Options{}
}
