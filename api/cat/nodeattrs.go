// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Nodeattrs - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-nodeattrs.html for more info.
//
// options: optional parameters. Supports the following functional options: WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Nodeattrs(options ...*Option) (*NodeattrsResponse, error) {
	req := c.transport.NewRequest("GET")
	methodOptions := supportedOptions["Nodeattrs"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &NodeattrsResponse{resp}, err
}

// NodeattrsResponse is the response for Nodeattrs.
type NodeattrsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *NodeattrsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
