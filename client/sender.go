package client // import "danyelmorales.com/wc-api-gogo/client"

import (
	"net/http"
	"danyelmorales.com/wc-api-gogo/request"
)

// Sender interface
type Sender interface {
	Send(req request.Request) (resp *http.Response, err error)
}
