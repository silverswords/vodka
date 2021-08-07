package vodka

import (
	"github.com/yusank/vodka/client"
	"github.com/yusank/vodka/cmd"
	"github.com/yusank/vodka/config"
	"github.com/yusank/vodka/registry"
	"github.com/yusank/vodka/selector"
	"github.com/yusank/vodka/server"
)

// Service is vodka service
type Service interface {
	Init()
	Option()
	Start()
	Stop()
	Client() client.Client
	Server() []server.Server
	Config() config.Config
	Selector() selector.Selector
	Registry() registry.Registry
	Cmd() cmd.Cmd
}
