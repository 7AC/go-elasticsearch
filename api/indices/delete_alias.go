// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// DeleteAlias - APIs in Elasticsearch accept an index name when working against a specific index, and several indices when applicable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-aliases.html for more info.
//
// index: a comma-separated list of index names (supports wildcards); use "_all" for all indices.
//
// name: a comma-separated list of aliases to delete (supports wildcards); use "_all" to delete all aliases for the specified indices.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) DeleteAlias(index []string, name []string, options ...*Option) (*DeleteAliasResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "DELETE",
	}
	methodOptions := supportedOptions["DeleteAlias"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &DeleteAliasResponse{resp}, err
}

// DeleteAliasResponse is the response for DeleteAlias
type DeleteAliasResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *DeleteAliasResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}