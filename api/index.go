// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Index adds or updates a typed JSON document in a specific index, making it searchable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-index_.html for more info.
//
// index: the name of the index.
//
// documentType: the type of the document.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithID, WithOpType, WithParent, WithPipeline, WithRefresh, WithRouting, WithTimeout, WithTimestamp, WithTTL, WithVersion, WithVersionType, WithWaitForActiveShards, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Index(index string, documentType string, body map[string]interface{}, options ...*Option) (*IndexResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Index"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &IndexResponse{resp}, err
}

// IndexResponse is the response for Index.
type IndexResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *IndexResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
