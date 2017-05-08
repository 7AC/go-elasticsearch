// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package ingest

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Simulate - the ingest plugins extend Elasticsearch by providing additional ingest node capabilities. See https://www.elastic.co/guide/en/elasticsearch/plugins/5.x/ingest.html for more info.
//
// body: the simulate definition.
//
// options: optional parameters. Supports the following functional options: WithID, WithVerbose, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Ingest) Simulate(body map[string]interface{}, options ...*Option) (*SimulateResponse, error) {
	req := i.transport.NewRequest("GET")
	methodOptions := supportedOptions["Simulate"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &SimulateResponse{resp}, err
}

// SimulateResponse is the response for Simulate.
type SimulateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *SimulateResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
