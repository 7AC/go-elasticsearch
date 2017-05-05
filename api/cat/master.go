// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Master - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-master.html for more info.
//
// options: optional parameters. Supports the following functional options: WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Master(options ...*Option) (*MasterResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Master"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &MasterResponse{resp}, err
}

// MasterResponse is the response for Master.
type MasterResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MasterResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
