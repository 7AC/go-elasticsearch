// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// ExistsSource - the get API allows to get a typed JSON document from the index based on its id. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html for more info.
//
// index: the name of the index.
//
// documentType: the type of the document; use "_all" to fetch the first document matching the ID across all types.
//
// id: the document ID.
//
// options: optional parameters. Supports the following functional options: WithSource, WithSourceExclude, WithSourceInclude, WithParent, WithPreference, WithRealtime, WithRefresh, WithRouting, WithVersion, WithVersionType, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) ExistsSource(index string, documentType string, id string, options ...*Option) (*ExistsSourceResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "HEAD",
	}
	methodOptions := supportedOptions["ExistsSource"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &ExistsSourceResponse{resp}, err
}

// ExistsSourceResponse is the response for ExistsSource.
type ExistsSourceResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ExistsSourceResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
