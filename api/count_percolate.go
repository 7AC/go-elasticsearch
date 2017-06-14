// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// CountPercolateOption is a non-required CountPercolate option that gets applied to an HTTP request.
type CountPercolateOption func(r *transport.Request)

// WithCountPercolateID - substitute the document in the request body with a document that is known by the specified id. On top of the id, the index and type parameter will be used to retrieve the document from within the cluster.
func WithCountPercolateID(id string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithCountPercolateAllowNoIndices(allowNoIndices bool) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// CountPercolateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type CountPercolateExpandWildcards int

const (
	// CountPercolateExpandWildcardsOpen can be used to set CountPercolateExpandWildcards to "open"
	CountPercolateExpandWildcardsOpen = iota
	// CountPercolateExpandWildcardsClosed can be used to set CountPercolateExpandWildcards to "closed"
	CountPercolateExpandWildcardsClosed = iota
	// CountPercolateExpandWildcardsNone can be used to set CountPercolateExpandWildcards to "none"
	CountPercolateExpandWildcardsNone = iota
	// CountPercolateExpandWildcardsAll can be used to set CountPercolateExpandWildcards to "all"
	CountPercolateExpandWildcardsAll = iota
)

// WithCountPercolateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithCountPercolateExpandWildcards(expandWildcards CountPercolateExpandWildcards) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithCountPercolateIgnoreUnavailable(ignoreUnavailable bool) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolatePercolateIndex - the index to count percolate the document into. Defaults to index.
func WithCountPercolatePercolateIndex(percolateIndex string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolatePercolateType - the type to count percolate document into. Defaults to type.
func WithCountPercolatePercolateType(percolateType string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolatePreference - specify the node or shard the operation should be performed on (default: random).
func WithCountPercolatePreference(preference string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateRouting - a comma-separated list of specific routing values.
func WithCountPercolateRouting(routing []string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateVersion - explicit version number for concurrency control.
func WithCountPercolateVersion(version int) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// CountPercolateVersionType - specific version type.
type CountPercolateVersionType int

const (
	// CountPercolateVersionTypeInternal can be used to set CountPercolateVersionType to "internal"
	CountPercolateVersionTypeInternal = iota
	// CountPercolateVersionTypeExternal can be used to set CountPercolateVersionType to "external"
	CountPercolateVersionTypeExternal = iota
	// CountPercolateVersionTypeExternalGte can be used to set CountPercolateVersionType to "external_gte"
	CountPercolateVersionTypeExternalGte = iota
	// CountPercolateVersionTypeForce can be used to set CountPercolateVersionType to "force"
	CountPercolateVersionTypeForce = iota
)

// WithCountPercolateVersionType - specific version type.
func WithCountPercolateVersionType(versionType CountPercolateVersionType) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateBody - the count percolator request definition using the percolate DSL.
func WithCountPercolateBody(body map[string]interface{}) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateErrorTrace - include the stack trace of returned errors.
func WithCountPercolateErrorTrace(errorTrace bool) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithCountPercolateFilterPath(filterPath []string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateHuman - return human readable values for statistics.
func WithCountPercolateHuman(human bool) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateIgnore - ignores the specified HTTP status codes.
func WithCountPercolateIgnore(ignore []int) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolatePretty - pretty format the returned JSON response.
func WithCountPercolatePretty(pretty bool) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// WithCountPercolateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithCountPercolateSourceParam(sourceParam string) CountPercolateOption {
	return func(r *transport.Request) {
	}
}

// CountPercolate - for indices created on or after version 5.0.0-alpha1 the percolator automatically indexes the query terms with the percolator queries. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-percolate.html for more info.
//
// index: the index of the document being count percolated.
//
// documentType: the type of the document being count percolated.
//
// options: optional parameters.
func (a *API) CountPercolate(index string, documentType string, options ...CountPercolateOption) (*CountPercolateResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &CountPercolateResponse{resp}, err
}

// CountPercolateResponse is the response for CountPercolate.
type CountPercolateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *CountPercolateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
