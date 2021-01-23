package url

import (
	"danyelmorales.com/wc-api-gogo/request"
	URL "net/url"
)

// QueryEnricherMock ...
type QueryEnricherMock struct {
	query URL.Values
}

// GetEnrichedQuery ...
func (q *QueryEnricherMock) GetEnrichedQuery(url string, query URL.Values, req request.Request) URL.Values {
	return q.query
}
