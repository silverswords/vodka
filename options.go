package vodka

import (
	"github.com/yusank/vodka/client"
	"github.com/yusank/vodka/cmd"
	"github.com/yusank/vodka/server"
)

// Options for Service
type Options struct {
	Cmd    cmd.Cmd
	Client client.Client
	Server []server.Server
}
