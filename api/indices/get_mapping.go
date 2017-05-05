// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// GetMapping - the get mapping API allows to retrieve mapping definitions for an index or index/type. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-get-mapping.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetMapping(options ...*Option) (*GetMappingResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["GetMapping"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &GetMappingResponse{resp}, err
}

// GetMappingResponse is the response for GetMapping.
type GetMappingResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetMappingResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
