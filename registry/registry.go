package registry

import "github.com/yusank/vodka/api"

type Registry interface {
	Init()
	Registry()
	UnRegistry()
	GetService()
	api.IName
}
