// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// GetFieldMapping - the get field mapping API allows you to retrieve mapping definitions for one or more fields. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-get-field-mapping.html for more info.
//
// fields: a comma-separated list of fields.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithIncludeDefaults, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetFieldMapping(fields []string, options ...*Option) (*GetFieldMappingResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["GetFieldMapping"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &GetFieldMappingResponse{resp}, err
}

// GetFieldMappingResponse is the response for GetFieldMapping
type GetFieldMappingResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *GetFieldMappingResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}