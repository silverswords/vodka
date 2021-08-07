package selector

import "github.com/yusank/vodka/api"

type Selector interface {
	Init()
	Get()
	SetUnHealth()
	api.IName
}
