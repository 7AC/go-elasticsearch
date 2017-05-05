// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// ShardStores - provides store information for shard copies of indices. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-shards-stores.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithOperationThreading, WithStatus, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ShardStores(options ...*Option) (*ShardStoresResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["ShardStores"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ShardStoresResponse{resp}, err
}

// ShardStoresResponse is the response for ShardStores.
type ShardStoresResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ShardStoresResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
