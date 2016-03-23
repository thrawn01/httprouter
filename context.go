package httprouter
import (
	"time"
	"golang.org/x/net/context"
)


type ParamContext interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	ByName(string) string
	Value(key interface{}) interface{}
}

type ParamContextImpl struct {
	context.Context
	Params
}

