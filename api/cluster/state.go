// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// State - the cluster state API allows to get a comprehensive state information of the whole cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-state.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithMetric, WithAllowNoIndices, WithExpandWildcards, WithFlatSettings, WithIgnoreUnavailable, WithLocal, WithMasterTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cluster) State(options ...*Option) (*StateResponse, error) {
	req := c.transport.NewRequest("GET")
	methodOptions := supportedOptions["State"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &StateResponse{resp}, err
}

// StateResponse is the response for State.
type StateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *StateResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
