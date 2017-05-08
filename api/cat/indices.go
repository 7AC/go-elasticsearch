// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Indices - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-indices.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithBytes, WithFormat, WithH, WithHealth, WithHelp, WithLocal, WithMasterTimeout, WithPri, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Indices(options ...*Option) (*IndicesResponse, error) {
	req := c.transport.NewRequest("GET")
	methodOptions := supportedOptions["Indices"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &IndicesResponse{resp}, err
}

// IndicesResponse is the response for Indices.
type IndicesResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *IndicesResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
