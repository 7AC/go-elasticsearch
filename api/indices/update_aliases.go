// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// UpdateAliases - APIs in Elasticsearch accept an index name when working against a specific index, and several indices when applicable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-aliases.html for more info.
//
// body: the definition of "actions" to perform.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) UpdateAliases(body map[string]interface{}, options ...*Option) (*UpdateAliasesResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["UpdateAliases"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &UpdateAliasesResponse{resp}, err
}

// UpdateAliasesResponse is the response for UpdateAliases.
type UpdateAliasesResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *UpdateAliasesResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
