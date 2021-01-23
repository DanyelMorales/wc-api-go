package url // import "danyelmorales.com/wc-api-gogo/url"

import (
	"danyelmorales.com/wc-api-gogo/request"
	"net/url"
)

// QueryEnricher uses package auth to enrich existing query parameters with Authentication Based ones
type QueryEnricher interface {
	GetEnrichedQuery(url string, query url.Values, req request.Request) url.Values
}
