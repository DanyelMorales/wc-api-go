package net

import (
	"bytes"
	"danyelmorales.com/wc-api-gogo/request"
	"io"
	"net/http"
)

// Sender provides HTTP Requests
type Sender struct {
	requestEnricher RequestEnricher
	urlBuilder      URLBuilder
	httpClient      Client
	requestCreator  RequestCreator
}

// Send method sends requests to WooCommerce API
func (s *Sender) Send(req request.Request) (resp *http.Response, err error) {
	request := s.prepareRequest(req)
	return s.httpClient.Do(request)
}

func (s *Sender) prepareRequest(req request.Request) *http.Request {
	URL := s.urlBuilder.GetURL(req)
	var bodyReader io.Reader
	if req.Body != nil && ("POST" == req.Method || "PUT" == req.Method) {
		bodyReader = bytes.NewBuffer(req.Body)
	}
	request, _ := s.requestCreator.NewRequest(req.Method, URL, bodyReader)
	s.requestEnricher.EnrichRequest(request, URL)
	request.Header.Set("Content-Type", "application/json")
	return request
}

// SetRequestEnricher ...
func (s *Sender) SetRequestEnricher(a RequestEnricher) {
	s.requestEnricher = a
}

// SetURLBuilder ...
func (s *Sender) SetURLBuilder(urlBuilder URLBuilder) {
	s.urlBuilder = urlBuilder
}

// SetHTTPClient ...
func (s *Sender) SetHTTPClient(c Client) {
	s.httpClient = c
}

// SetRequestCreator ...
func (s *Sender) SetRequestCreator(rc RequestCreator) {
	s.requestCreator = rc
}
