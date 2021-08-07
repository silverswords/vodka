package server

import "github.com/yusank/vodka/api"

// Server defines handle rpc request
type Server interface {
	Start() error
	Stop()
	Handle(interface{}) error
	api.IName
}

// Option apply option func fot Server
type Option func(*Options)

func Addr(addr string) Option {
	return func(o *Options) {
		o.Address = addr
		return
	}
}
