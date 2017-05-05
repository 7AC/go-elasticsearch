// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Mget - multi GET API allows to get multiple documents based on an index, type (optional) and id (and possibly routing). See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-multi-get.html for more info.
//
// body: document identifiers; can be either "docs" (containing full document information) or "ids" (when index and type is provided in the URL.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithSource, WithSourceExclude, WithSourceInclude, WithPreference, WithRealtime, WithRefresh, WithRouting, WithStoredFields, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Mget(body map[string]interface{}, options ...*Option) (*MgetResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Mget"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &MgetResponse{resp}, err
}

// MgetResponse is the response for Mget.
type MgetResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MgetResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
