// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Bulk makes it possible to perform many index/delete operations in a single API call. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-bulk.html for more info.
//
// body: the operation definition and data (action-data pairs), separated by newlines.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithSource, WithSourceExclude, WithSourceInclude, WithFields, WithPipeline, WithRefresh, WithRouting, WithTimeout, WithTypeParam, WithWaitForActiveShards, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Bulk(body map[string]interface{}, options ...*Option) (*BulkResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Bulk"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &BulkResponse{resp}, err
}

// BulkResponse is the response for Bulk.
type BulkResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *BulkResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
