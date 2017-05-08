// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Health - the cluster health API allows to get a very simple status on the health of the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-health.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithLevel, WithLocal, WithMasterTimeout, WithTimeout, WithWaitForActiveShards, WithWaitForEvents, WithWaitForNoRelocatingShards, WithWaitForNodes, WithWaitForStatus, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cluster) Health(options ...*Option) (*HealthResponse, error) {
	req := c.transport.NewRequest("GET")
	methodOptions := supportedOptions["Health"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &HealthResponse{resp}, err
}

// HealthResponse is the response for Health.
type HealthResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *HealthResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
