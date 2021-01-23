package net

import (
	"danyelmorales.com/wc-api-gogo/request"
)

// URLBuilder interface
type URLBuilder interface {
	GetURL(req request.Request) string
}
