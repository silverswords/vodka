package client

import "github.com/yusank/vodka/api"

type Client interface {
	Init(...Option)
	Call()
	api.IName
}
