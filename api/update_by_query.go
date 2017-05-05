// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// UpdateByQuery - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-update-by-query.html for more info.
//
// index: a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
//
// options: optional parameters. Supports the following functional options: WithType, WithSource, WithSourceExclude, WithSourceInclude, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithConflicts, WithDefaultOperator, WithDf, WithExpandWildcards, WithFrom, WithIgnoreUnavailable, WithLenient, WithPipeline, WithPreference, WithQ, WithRefresh, WithRequestCache, WithRequestsPerSecond, WithRouting, WithScroll, WithScrollSize, WithSearchTimeout, WithSearchType, WithSize, WithSlices, WithSort, WithStats, WithTerminateAfter, WithTimeout, WithVersion, WithVersionType, WithWaitForActiveShards, WithWaitForCompletion, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) UpdateByQuery(index []string, options ...*Option) (*UpdateByQueryResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["UpdateByQuery"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &UpdateByQueryResponse{resp}, err
}

// UpdateByQueryResponse is the response for UpdateByQuery.
type UpdateByQueryResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *UpdateByQueryResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
