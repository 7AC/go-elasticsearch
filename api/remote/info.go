// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package remote

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Info - see http://www.elastic.co/guide/en/elasticsearch/reference/master/remote-info.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (r *Remote) Info(options ...*Option) (*InfoResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: r.transport.URL.Scheme,
			Host:   r.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Info"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := r.transport.Do(req)
	return &InfoResponse{resp}, err
}

// InfoResponse is the response for Info.
type InfoResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *InfoResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
