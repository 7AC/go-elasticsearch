// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Mtermvectors - multi termvectors API allows to get multiple termvectors at once. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-multi-termvectors.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithFieldStatistics, WithFields, WithIds, WithOffsets, WithParent, WithPayloads, WithPositions, WithPreference, WithRealtime, WithRouting, WithTermStatistics, WithVersion, WithVersionType, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Mtermvectors(options ...*Option) (*MtermvectorsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Mtermvectors"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &MtermvectorsResponse{resp}, err
}

// MtermvectorsResponse is the response for Mtermvectors.
type MtermvectorsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MtermvectorsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
