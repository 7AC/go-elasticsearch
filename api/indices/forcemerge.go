// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// ForcemergeOption is a non-required Forcemerge option that gets applied to an HTTP request.
type ForcemergeOption func(r *transport.Request)

// WithForcemergeIndex - a comma-separated list of index names; use "_all" or empty string to perform the operation on all indices.
func WithForcemergeIndex(index []string) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithForcemergeAllowNoIndices(allowNoIndices bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// ForcemergeExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type ForcemergeExpandWildcards int

const (
	// ForcemergeExpandWildcardsOpen can be used to set ForcemergeExpandWildcards to "open"
	ForcemergeExpandWildcardsOpen = iota
	// ForcemergeExpandWildcardsClosed can be used to set ForcemergeExpandWildcards to "closed"
	ForcemergeExpandWildcardsClosed = iota
	// ForcemergeExpandWildcardsNone can be used to set ForcemergeExpandWildcards to "none"
	ForcemergeExpandWildcardsNone = iota
	// ForcemergeExpandWildcardsAll can be used to set ForcemergeExpandWildcards to "all"
	ForcemergeExpandWildcardsAll = iota
)

// WithForcemergeExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithForcemergeExpandWildcards(expandWildcards ForcemergeExpandWildcards) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeFlush - specify whether the index should be flushed after performing the operation (default: true).
func WithForcemergeFlush(flush bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithForcemergeIgnoreUnavailable(ignoreUnavailable bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeMaxNumSegments - the number of segments the index should be merged into (default: dynamic).
func WithForcemergeMaxNumSegments(maxNumSegments int) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeOnlyExpungeDeletes - specify whether the operation should only expunge deleted documents.
func WithForcemergeOnlyExpungeDeletes(onlyExpungeDeletes bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeWaitForMerge - specify whether the request should block until the merge process is finished (default: true).
func WithForcemergeWaitForMerge(waitForMerge bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeErrorTrace - include the stack trace of returned errors.
func WithForcemergeErrorTrace(errorTrace bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeFilterPath - a comma-separated list of filters used to reduce the respone.
func WithForcemergeFilterPath(filterPath []string) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeHuman - return human readable values for statistics.
func WithForcemergeHuman(human bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeIgnore - ignores the specified HTTP status codes.
func WithForcemergeIgnore(ignore []int) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergePretty - pretty format the returned JSON response.
func WithForcemergePretty(pretty bool) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// WithForcemergeSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithForcemergeSourceParam(sourceParam string) ForcemergeOption {
	return func(r *transport.Request) {
	}
}

// Forcemerge - the force merge API allows to force merging of one or more indices through an API. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-forcemerge.html for more info.
//
// options: optional parameters.
func (i *Indices) Forcemerge(options ...ForcemergeOption) (*ForcemergeResponse, error) {
	req := i.transport.NewRequest("POST")
	for _, option := range options {
		option(req)
	}
	resp, err := i.transport.Do(req)
	return &ForcemergeResponse{resp}, err
}

// ForcemergeResponse is the response for Forcemerge.
type ForcemergeResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ForcemergeResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
