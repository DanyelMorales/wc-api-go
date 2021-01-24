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

func (j *JsonMap) Remove(key string) {
	delete(*j, key)
}
func (j *JsonMap) RemoveIndex(key string, index int) {
	if val := (*j).Get(key); val != nil {
		container := make([]interface{}, 0)
		for i, data := range val {
			if i != index {
				container = append(container, data)
			}
		}
		(*j)[key] = container
	}
}

func (j *JsonMap) Add(key string, value []interface{}) {
	(*j)[key] = value
}

func (j *JsonMap) Set(key string, value interface{}) {
	if (*j)[key] == nil {
		(*j)[key] = make([]interface{}, 0)
	}
	(*j)[key] = append((*j)[key], value)
}

func (j *JsonMap) Get(key string) []interface{} {
	v, ok := (*j)[key]
	if !ok {
		return nil
	}
	return v
}

func (j *JsonMap) ToByteArray() []byte {
	b, err := json.Marshal(*j)
	if err != nil {
		return nil
	}
	return b
}
