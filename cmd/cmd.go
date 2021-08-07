package cmd

import "github.com/yusank/vodka/api"

type Cmd interface {
	Init()
	api.IName
}
