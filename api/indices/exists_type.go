// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// ExistsType - used to check if a type/types exists in an index/indices. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-types-exists.html for more info.
//
// index: a comma-separated list of index names; use "_all" to check the types across all indices.
//
// documentType: a comma-separated list of document types to check.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ExistsType(index []string, documentType []string, options ...*Option) (*ExistsTypeResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "HEAD",
	}
	methodOptions := supportedOptions["ExistsType"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ExistsTypeResponse{resp}, err
}

// ExistsTypeResponse is the response for ExistsType.
type ExistsTypeResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ExistsTypeResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
