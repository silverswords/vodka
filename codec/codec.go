package codec

import "github.com/yusank/vodka/api"

type Codec interface {
	Marshall(interface{}) ([]byte, error)
	UnMarshall([]byte, interface{}) error
	api.IName
}
