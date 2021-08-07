package config

import "github.com/yusank/vodka/api"

type Config interface {
	Init()
	Read()
	Watch()
	api.IName
}
