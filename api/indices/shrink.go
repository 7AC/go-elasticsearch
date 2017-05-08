// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Shrink - the shrink index API allows you to shrink an existing index into a new index with fewer primary shards. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-shrink-index.html for more info.
//
// index: the name of the source index to shrink.
//
// target: the name of the target index to shrink into.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, WithWaitForActiveShards, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Shrink(index string, target string, options ...*Option) (*ShrinkResponse, error) {
	req := i.transport.NewRequest("PUT")
	methodOptions := supportedOptions["Shrink"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ShrinkResponse{resp}, err
}

// ShrinkResponse is the response for Shrink.
type ShrinkResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ShrinkResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
