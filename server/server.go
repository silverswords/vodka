package server

import "github.com/yusank/vodka/api"

type Server interface {
	Init()
	Start()
	Stop()
	Handle()
	api.IName
}
