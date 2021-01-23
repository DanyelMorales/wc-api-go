package request // import "github.com/DanyelMorales/wc-api-go/request"

import (
	"encoding/json"
	"net/url"
)

// Request ...
type Request struct {
	Method   string
	Endpoint string
	Values   url.Values
	Body     []byte
}

type JsonMap map[string][]interface{}

func (j *JsonMap) Add(key string, value interface{}) {
	(*j)[key] = append((*j)[key], value)
}

func (j *JsonMap) Set(key string, value interface{}) {
	(*j)[key] = []interface{}{value}
}

func (j *JsonMap) Get(key string) interface{} {
	v, ok := (*j)[key]
	if !ok {
		return nil
	}
	return v
}

func (j *JsonMap) ToByte() []byte {
	b, err := json.Marshal(*j)
	if err != nil {
		return nil
	}
	return b
}
