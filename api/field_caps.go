// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// FieldCaps - this functionality is experimental and may be changed or removed completely in a future release. See http://www.elastic.co/guide/en/elasticsearch/reference/master/search-field-caps.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithFields, WithIgnoreUnavailable, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) FieldCaps(options ...*Option) (*FieldCapsResponse, error) {
	req := a.transport.NewRequest("GET")
	methodOptions := supportedOptions["FieldCaps"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &FieldCapsResponse{resp}, err
}

// FieldCapsResponse is the response for FieldCaps.
type FieldCapsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *FieldCapsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
