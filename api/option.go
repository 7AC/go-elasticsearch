// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"time"

	"github.com/elastic/go-elasticsearch/transport"
)

// Conflicts - what to do when the update by query hits version conflicts?
type Conflicts int

const (
	// ConflictsAbort can be used to set Conflicts to "abort"
	ConflictsAbort = iota
	// ConflictsProceed can be used to set Conflicts to "proceed"
	ConflictsProceed = iota
)

// DefaultOperator - the default operator for query string query (AND or OR).
type DefaultOperator int

const (
	// DefaultOperatorAND can be used to set DefaultOperator to "AND"
	DefaultOperatorAND = iota
	// DefaultOperatorOR can be used to set DefaultOperator to "OR"
	DefaultOperatorOR = iota
)

// ExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type ExpandWildcards int

const (
	// ExpandWildcardsOpen can be used to set ExpandWildcards to "open"
	ExpandWildcardsOpen = iota
	// ExpandWildcardsClosed can be used to set ExpandWildcards to "closed"
	ExpandWildcardsClosed = iota
	// ExpandWildcardsNone can be used to set ExpandWildcards to "none"
	ExpandWildcardsNone = iota
	// ExpandWildcardsAll can be used to set ExpandWildcards to "all"
	ExpandWildcardsAll = iota
)

// SearchType - search operation type.
type SearchType int

const (
	// SearchTypeQueryThenFetch can be used to set SearchType to "query_then_fetch"
	SearchTypeQueryThenFetch = iota
	// SearchTypeDfsQueryThenFetch can be used to set SearchType to "dfs_query_then_fetch"
	SearchTypeDfsQueryThenFetch = iota
)

// Option is a non-required API option that gets applied to an HTTP request.
type Option struct {
	name  string
	apply func(r *transport.Request)
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithAllowNoIndices(allowNoIndices bool) *Option {
	return &Option{
		name: "WithAllowNoIndices",
		apply: func(r *transport.Request) {
		},
	}
}

// WithAnalyzeWildcard - specify whether wildcard and prefix queries should be analyzed (default: false).
func WithAnalyzeWildcard(analyzeWildcard bool) *Option {
	return &Option{
		name: "WithAnalyzeWildcard",
		apply: func(r *transport.Request) {
		},
	}
}

// WithAnalyzer - the analyzer to use for the query string.
func WithAnalyzer(analyzer string) *Option {
	return &Option{
		name: "WithAnalyzer",
		apply: func(r *transport.Request) {
		},
	}
}

// WithBody - the search definition using the Query DSL.
func WithBody(body map[string]interface{}) *Option {
	return &Option{
		name: "WithBody",
		apply: func(r *transport.Request) {
		},
	}
}

// WithConflicts - what to do when the update by query hits version conflicts?
func WithConflicts(conflicts Conflicts) *Option {
	return &Option{
		name: "WithConflicts",
		apply: func(r *transport.Request) {
		},
	}
}

// WithDefaultOperator - the default operator for query string query (AND or OR).
func WithDefaultOperator(defaultOperator DefaultOperator) *Option {
	return &Option{
		name: "WithDefaultOperator",
		apply: func(r *transport.Request) {
		},
	}
}

// WithDf - the field to use as default where no field prefix is given in the query string.
func WithDf(df string) *Option {
	return &Option{
		name: "WithDf",
		apply: func(r *transport.Request) {
		},
	}
}

// WithType - a comma-separated list of document types to search; leave empty to perform the operation on all types.
func WithType(documentType []string) *Option {
	return &Option{
		name: "WithType",
		apply: func(r *transport.Request) {
		},
	}
}

// WithErrorTrace - include the stack trace of returned errors.
func WithErrorTrace(errorTrace bool) *Option {
	return &Option{
		name: "WithErrorTrace",
		apply: func(r *transport.Request) {
		},
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithExpandWildcards(expandWildcards ExpandWildcards) *Option {
	return &Option{
		name: "WithExpandWildcards",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFields - a comma-separated list of field names.
func WithFields(fields []string) *Option {
	return &Option{
		name: "WithFields",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFilterPath - a comma-separated list of filters used to reduce the respone.
func WithFilterPath(filterPath []string) *Option {
	return &Option{
		name: "WithFilterPath",
		apply: func(r *transport.Request) {
		},
	}
}

// WithFrom - starting offset (default: 0).
func WithFrom(from int) *Option {
	return &Option{
		name: "WithFrom",
		apply: func(r *transport.Request) {
		},
	}
}

// WithHuman - return human readable values for statistics.
func WithHuman(human bool) *Option {
	return &Option{
		name: "WithHuman",
		apply: func(r *transport.Request) {
		},
	}
}

// WithID - the id of the stored search template.
func WithID(id string) *Option {
	return &Option{
		name: "WithID",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIgnore - ignores the specified HTTP status codes.
func WithIgnore(ignore []string) *Option {
	return &Option{
		name: "WithIgnore",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithIgnoreUnavailable(ignoreUnavailable bool) *Option {
	return &Option{
		name: "WithIgnoreUnavailable",
		apply: func(r *transport.Request) {
		},
	}
}

// WithIndex - a comma-separated list of index names; use "_all" or empty string to perform the operation on all indices.
func WithIndex(index []string) *Option {
	return &Option{
		name: "WithIndex",
		apply: func(r *transport.Request) {
		},
	}
}

// WithLenient - specify whether format-based query failures (such as providing text to a numeric field) should be ignored.
func WithLenient(lenient bool) *Option {
	return &Option{
		name: "WithLenient",
		apply: func(r *transport.Request) {
		},
	}
}

// WithParent - the ID of the parent document.
func WithParent(parent string) *Option {
	return &Option{
		name: "WithParent",
		apply: func(r *transport.Request) {
		},
	}
}

// WithPipeline - ingest pipeline to set on index requests made by this action. (default: none).
func WithPipeline(pipeline string) *Option {
	return &Option{
		name: "WithPipeline",
		apply: func(r *transport.Request) {
		},
	}
}

// WithPreference - specify the node or shard the operation should be performed on (default: random).
func WithPreference(preference string) *Option {
	return &Option{
		name: "WithPreference",
		apply: func(r *transport.Request) {
		},
	}
}

// WithPretty - pretty format the returned JSON response.
func WithPretty(pretty bool) *Option {
	return &Option{
		name: "WithPretty",
		apply: func(r *transport.Request) {
		},
	}
}

// WithQ - query in the Lucene query string syntax.
func WithQ(q string) *Option {
	return &Option{
		name: "WithQ",
		apply: func(r *transport.Request) {
		},
	}
}

// WithRealtime - specify whether to perform the operation in realtime or search mode.
func WithRealtime(realtime bool) *Option {
	return &Option{
		name: "WithRealtime",
		apply: func(r *transport.Request) {
		},
	}
}

// WithRefresh - should the effected indexes be refreshed?
func WithRefresh(refresh bool) *Option {
	return &Option{
		name: "WithRefresh",
		apply: func(r *transport.Request) {
		},
	}
}

// WithRequestCache - specify if request cache should be used for this request or not, defaults to index level setting.
func WithRequestCache(requestCache bool) *Option {
	return &Option{
		name: "WithRequestCache",
		apply: func(r *transport.Request) {
		},
	}
}

// WithRequestsPerSecond - the throttle to set on this request in sub-requests per second. -1 means no throttle.
func WithRequestsPerSecond(requestsPerSecond int) *Option {
	return &Option{
		name: "WithRequestsPerSecond",
		apply: func(r *transport.Request) {
		},
	}
}

// WithRouting - a comma-separated list of specific routing values.
func WithRouting(routing []string) *Option {
	return &Option{
		name: "WithRouting",
		apply: func(r *transport.Request) {
		},
	}
}

// WithScroll - specify how long a consistent view of the index should be maintained for scrolled search.
func WithScroll(scroll time.Time) *Option {
	return &Option{
		name: "WithScroll",
		apply: func(r *transport.Request) {
		},
	}
}

// WithScrollID - a comma-separated list of scroll IDs to clear.
func WithScrollID(scrollID []string) *Option {
	return &Option{
		name: "WithScrollID",
		apply: func(r *transport.Request) {
		},
	}
}

// WithScrollSize - size on the scroll request powering the update_by_query.
func WithScrollSize(scrollSize int) *Option {
	return &Option{
		name: "WithScrollSize",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSearchTimeout - explicit timeout for each search request. Defaults to no timeout.
func WithSearchTimeout(searchTimeout time.Time) *Option {
	return &Option{
		name: "WithSearchTimeout",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSearchType - search operation type.
func WithSearchType(searchType SearchType) *Option {
	return &Option{
		name: "WithSearchType",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSize - number of hits to return (default: 10).
func WithSize(size int) *Option {
	return &Option{
		name: "WithSize",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSlices - the number of slices this task should be divided into. Defaults to 1 meaning the task isn't sliced into subtasks.
func WithSlices(slices int) *Option {
	return &Option{
		name: "WithSlices",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSort - a comma-separated list of <field>:<direction> pairs.
func WithSort(sort []string) *Option {
	return &Option{
		name: "WithSort",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSource - true or false to return the _source field or not, or a list of fields to return.
func WithSource(source []string) *Option {
	return &Option{
		name: "WithSource",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSourceExclude - a list of fields to exclude from the returned _source field.
func WithSourceExclude(sourceExclude []string) *Option {
	return &Option{
		name: "WithSourceExclude",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSourceInclude - a list of fields to extract and return from the _source field.
func WithSourceInclude(sourceInclude []string) *Option {
	return &Option{
		name: "WithSourceInclude",
		apply: func(r *transport.Request) {
		},
	}
}

// WithSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSourceParam(sourceParam string) *Option {
	return &Option{
		name: "WithSourceParam",
		apply: func(r *transport.Request) {
		},
	}
}

// WithStats - specific 'tag' of the request for logging and statistical purposes.
func WithStats(stats []string) *Option {
	return &Option{
		name: "WithStats",
		apply: func(r *transport.Request) {
		},
	}
}

// WithTaskID - the task id to rethrottle.
func WithTaskID(taskID string) *Option {
	return &Option{
		name: "WithTaskID",
		apply: func(r *transport.Request) {
		},
	}
}

// WithTerminateAfter - the maximum number of documents to collect for each shard, upon reaching which the query execution will terminate early.
func WithTerminateAfter(terminateAfter int) *Option {
	return &Option{
		name: "WithTerminateAfter",
		apply: func(r *transport.Request) {
		},
	}
}

// WithTimeout - time each individual bulk request should wait for shards that are unavailable.
func WithTimeout(timeout time.Time) *Option {
	return &Option{
		name: "WithTimeout",
		apply: func(r *transport.Request) {
		},
	}
}

// WithVersion - specify whether to return document version as part of a hit.
func WithVersion(version bool) *Option {
	return &Option{
		name: "WithVersion",
		apply: func(r *transport.Request) {
		},
	}
}

// WithVersionType - should the document increment the version number (internal) on hit or not (reindex).
func WithVersionType(versionType bool) *Option {
	return &Option{
		name: "WithVersionType",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForActiveShards - sets the number of shard copies that must be active before proceeding with the update by query operation. Defaults to 1, meaning the primary shard only. Set to "all" for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func WithWaitForActiveShards(waitForActiveShards string) *Option {
	return &Option{
		name: "WithWaitForActiveShards",
		apply: func(r *transport.Request) {
		},
	}
}

// WithWaitForCompletion - should the request should block until the update by query operation is complete.
func WithWaitForCompletion(waitForCompletion bool) *Option {
	return &Option{
		name: "WithWaitForCompletion",
		apply: func(r *transport.Request) {
		},
	}
}

var (
	supportedOptions = map[string]map[string]struct{}{
		"UpdateByQuery": map[string]struct{}{
			"WithType":                struct{}{},
			"WithSource":              struct{}{},
			"WithSourceExclude":       struct{}{},
			"WithSourceInclude":       struct{}{},
			"WithAllowNoIndices":      struct{}{},
			"WithAnalyzeWildcard":     struct{}{},
			"WithAnalyzer":            struct{}{},
			"WithConflicts":           struct{}{},
			"WithDefaultOperator":     struct{}{},
			"WithDf":                  struct{}{},
			"WithExpandWildcards":     struct{}{},
			"WithFrom":                struct{}{},
			"WithIgnoreUnavailable":   struct{}{},
			"WithLenient":             struct{}{},
			"WithPipeline":            struct{}{},
			"WithPreference":          struct{}{},
			"WithQ":                   struct{}{},
			"WithRefresh":             struct{}{},
			"WithRequestCache":        struct{}{},
			"WithRequestsPerSecond":   struct{}{},
			"WithRouting":             struct{}{},
			"WithScroll":              struct{}{},
			"WithScrollSize":          struct{}{},
			"WithSearchTimeout":       struct{}{},
			"WithSearchType":          struct{}{},
			"WithSize":                struct{}{},
			"WithSlices":              struct{}{},
			"WithSort":                struct{}{},
			"WithStats":               struct{}{},
			"WithTerminateAfter":      struct{}{},
			"WithTimeout":             struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithWaitForCompletion":   struct{}{},
			"WithBody":                struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"ClearScroll": map[string]struct{}{
			"WithScrollID":    struct{}{},
			"WithBody":        struct{}{},
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"DeleteScript": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"FieldCaps": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithFields":            struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"GetSource": map[string]struct{}{
			"WithSource":        struct{}{},
			"WithSourceExclude": struct{}{},
			"WithSourceInclude": struct{}{},
			"WithParent":        struct{}{},
			"WithPreference":    struct{}{},
			"WithRealtime":      struct{}{},
			"WithRefresh":       struct{}{},
			"WithRouting":       struct{}{},
			"WithVersion":       struct{}{},
			"WithVersionType":   struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"Mget": map[string]struct{}{
			"WithIndex":         struct{}{},
			"WithType":          struct{}{},
			"WithSource":        struct{}{},
			"WithSourceExclude": struct{}{},
			"WithSourceInclude": struct{}{},
			"WithPreference":    struct{}{},
			"WithRealtime":      struct{}{},
			"WithRefresh":       struct{}{},
			"WithRouting":       struct{}{},
			"WithStoredFields":  struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"RenderSearchTemplate": map[string]struct{}{
			"WithID":          struct{}{},
			"WithBody":        struct{}{},
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Index": map[string]struct{}{
			"WithID":                  struct{}{},
			"WithOpType":              struct{}{},
			"WithParent":              struct{}{},
			"WithPipeline":            struct{}{},
			"WithRefresh":             struct{}{},
			"WithRouting":             struct{}{},
			"WithTimeout":             struct{}{},
			"WithTimestamp":           struct{}{},
			"WithTTL":                 struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"CountPercolate": map[string]struct{}{
			"WithID":                struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithPercolateIndex":    struct{}{},
			"WithPercolateType":     struct{}{},
			"WithPreference":        struct{}{},
			"WithRouting":           struct{}{},
			"WithVersion":           struct{}{},
			"WithVersionType":       struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"FieldStats": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithFields":            struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithLevel":             struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"Explain": map[string]struct{}{
			"WithSource":          struct{}{},
			"WithSourceExclude":   struct{}{},
			"WithSourceInclude":   struct{}{},
			"WithAnalyzeWildcard": struct{}{},
			"WithAnalyzer":        struct{}{},
			"WithDefaultOperator": struct{}{},
			"WithDf":              struct{}{},
			"WithLenient":         struct{}{},
			"WithParent":          struct{}{},
			"WithPreference":      struct{}{},
			"WithQ":               struct{}{},
			"WithRouting":         struct{}{},
			"WithStoredFields":    struct{}{},
			"WithBody":            struct{}{},
			"WithErrorTrace":      struct{}{},
			"WithFilterPath":      struct{}{},
			"WithHuman":           struct{}{},
			"WithIgnore":          struct{}{},
			"WithPretty":          struct{}{},
			"WithSourceParam":     struct{}{},
		},
		"PutTemplate": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Count": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithType":              struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithAnalyzeWildcard":   struct{}{},
			"WithAnalyzer":          struct{}{},
			"WithDefaultOperator":   struct{}{},
			"WithDf":                struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithLenient":           struct{}{},
			"WithMinScore":          struct{}{},
			"WithPreference":        struct{}{},
			"WithQ":                 struct{}{},
			"WithRouting":           struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"Bulk": map[string]struct{}{
			"WithIndex":               struct{}{},
			"WithType":                struct{}{},
			"WithSource":              struct{}{},
			"WithSourceExclude":       struct{}{},
			"WithSourceInclude":       struct{}{},
			"WithFields":              struct{}{},
			"WithPipeline":            struct{}{},
			"WithRefresh":             struct{}{},
			"WithRouting":             struct{}{},
			"WithTimeout":             struct{}{},
			"WithTypeParam":           struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Suggest": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithPreference":        struct{}{},
			"WithRouting":           struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"Msearch": map[string]struct{}{
			"WithIndex":                 struct{}{},
			"WithType":                  struct{}{},
			"WithMaxConcurrentSearches": struct{}{},
			"WithSearchType":            struct{}{},
			"WithTypedKeys":             struct{}{},
			"WithErrorTrace":            struct{}{},
			"WithFilterPath":            struct{}{},
			"WithHuman":                 struct{}{},
			"WithIgnore":                struct{}{},
			"WithPretty":                struct{}{},
			"WithSourceParam":           struct{}{},
		},
		"PutScript": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"MsearchTemplate": map[string]struct{}{
			"WithIndex":       struct{}{},
			"WithType":        struct{}{},
			"WithSearchType":  struct{}{},
			"WithTypedKeys":   struct{}{},
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Get": map[string]struct{}{
			"WithSource":        struct{}{},
			"WithSourceExclude": struct{}{},
			"WithSourceInclude": struct{}{},
			"WithParent":        struct{}{},
			"WithPreference":    struct{}{},
			"WithRealtime":      struct{}{},
			"WithRefresh":       struct{}{},
			"WithRouting":       struct{}{},
			"WithStoredFields":  struct{}{},
			"WithVersion":       struct{}{},
			"WithVersionType":   struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"Reindex": map[string]struct{}{
			"WithRefresh":             struct{}{},
			"WithRequestsPerSecond":   struct{}{},
			"WithSlices":              struct{}{},
			"WithTimeout":             struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithWaitForCompletion":   struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Mpercolate": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithType":              struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"Create": map[string]struct{}{
			"WithParent":              struct{}{},
			"WithPipeline":            struct{}{},
			"WithRefresh":             struct{}{},
			"WithRouting":             struct{}{},
			"WithTimeout":             struct{}{},
			"WithTimestamp":           struct{}{},
			"WithTTL":                 struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Delete": map[string]struct{}{
			"WithParent":              struct{}{},
			"WithRefresh":             struct{}{},
			"WithRouting":             struct{}{},
			"WithTimeout":             struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"GetTemplate": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Mtermvectors": map[string]struct{}{
			"WithIndex":           struct{}{},
			"WithType":            struct{}{},
			"WithFieldStatistics": struct{}{},
			"WithFields":          struct{}{},
			"WithIds":             struct{}{},
			"WithOffsets":         struct{}{},
			"WithParent":          struct{}{},
			"WithPayloads":        struct{}{},
			"WithPositions":       struct{}{},
			"WithPreference":      struct{}{},
			"WithRealtime":        struct{}{},
			"WithRouting":         struct{}{},
			"WithTermStatistics":  struct{}{},
			"WithVersion":         struct{}{},
			"WithVersionType":     struct{}{},
			"WithBody":            struct{}{},
			"WithErrorTrace":      struct{}{},
			"WithFilterPath":      struct{}{},
			"WithHuman":           struct{}{},
			"WithIgnore":          struct{}{},
			"WithPretty":          struct{}{},
			"WithSourceParam":     struct{}{},
		},
		"Scroll": map[string]struct{}{
			"WithScrollID":      struct{}{},
			"WithScroll":        struct{}{},
			"WithScrollIDParam": struct{}{},
			"WithBody":          struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"Info": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"SearchShards": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithType":              struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithLocal":             struct{}{},
			"WithPreference":        struct{}{},
			"WithRouting":           struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"ExistsSource": map[string]struct{}{
			"WithSource":        struct{}{},
			"WithSourceExclude": struct{}{},
			"WithSourceInclude": struct{}{},
			"WithParent":        struct{}{},
			"WithPreference":    struct{}{},
			"WithRealtime":      struct{}{},
			"WithRefresh":       struct{}{},
			"WithRouting":       struct{}{},
			"WithVersion":       struct{}{},
			"WithVersionType":   struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"GetScript": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Ping": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"SearchTemplate": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithType":              struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithExplain":           struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithPreference":        struct{}{},
			"WithProfile":           struct{}{},
			"WithRouting":           struct{}{},
			"WithScroll":            struct{}{},
			"WithSearchType":        struct{}{},
			"WithTypedKeys":         struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"DeleteTemplate": map[string]struct{}{
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"ReindexRethrottle": map[string]struct{}{
			"WithTaskID":      struct{}{},
			"WithErrorTrace":  struct{}{},
			"WithFilterPath":  struct{}{},
			"WithHuman":       struct{}{},
			"WithIgnore":      struct{}{},
			"WithPretty":      struct{}{},
			"WithSourceParam": struct{}{},
		},
		"Update": map[string]struct{}{
			"WithSource":              struct{}{},
			"WithSourceExclude":       struct{}{},
			"WithSourceInclude":       struct{}{},
			"WithFields":              struct{}{},
			"WithLang":                struct{}{},
			"WithParent":              struct{}{},
			"WithRefresh":             struct{}{},
			"WithRetryOnConflict":     struct{}{},
			"WithRouting":             struct{}{},
			"WithTimeout":             struct{}{},
			"WithTimestamp":           struct{}{},
			"WithTTL":                 struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithBody":                struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Exists": map[string]struct{}{
			"WithSource":        struct{}{},
			"WithSourceExclude": struct{}{},
			"WithSourceInclude": struct{}{},
			"WithParent":        struct{}{},
			"WithPreference":    struct{}{},
			"WithRealtime":      struct{}{},
			"WithRefresh":       struct{}{},
			"WithRouting":       struct{}{},
			"WithStoredFields":  struct{}{},
			"WithVersion":       struct{}{},
			"WithVersionType":   struct{}{},
			"WithErrorTrace":    struct{}{},
			"WithFilterPath":    struct{}{},
			"WithHuman":         struct{}{},
			"WithIgnore":        struct{}{},
			"WithPretty":        struct{}{},
			"WithSourceParam":   struct{}{},
		},
		"Percolate": map[string]struct{}{
			"WithID":                  struct{}{},
			"WithAllowNoIndices":      struct{}{},
			"WithExpandWildcards":     struct{}{},
			"WithIgnoreUnavailable":   struct{}{},
			"WithPercolateFormat":     struct{}{},
			"WithPercolateIndex":      struct{}{},
			"WithPercolatePreference": struct{}{},
			"WithPercolateRouting":    struct{}{},
			"WithPercolateType":       struct{}{},
			"WithPreference":          struct{}{},
			"WithRouting":             struct{}{},
			"WithVersion":             struct{}{},
			"WithVersionType":         struct{}{},
			"WithBody":                struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
		"Termvectors": map[string]struct{}{
			"WithID":              struct{}{},
			"WithFieldStatistics": struct{}{},
			"WithFields":          struct{}{},
			"WithOffsets":         struct{}{},
			"WithParent":          struct{}{},
			"WithPayloads":        struct{}{},
			"WithPositions":       struct{}{},
			"WithPreference":      struct{}{},
			"WithRealtime":        struct{}{},
			"WithRouting":         struct{}{},
			"WithTermStatistics":  struct{}{},
			"WithVersion":         struct{}{},
			"WithVersionType":     struct{}{},
			"WithBody":            struct{}{},
			"WithErrorTrace":      struct{}{},
			"WithFilterPath":      struct{}{},
			"WithHuman":           struct{}{},
			"WithIgnore":          struct{}{},
			"WithPretty":          struct{}{},
			"WithSourceParam":     struct{}{},
		},
		"Search": map[string]struct{}{
			"WithIndex":             struct{}{},
			"WithType":              struct{}{},
			"WithSource":            struct{}{},
			"WithSourceExclude":     struct{}{},
			"WithSourceInclude":     struct{}{},
			"WithAllowNoIndices":    struct{}{},
			"WithAnalyzeWildcard":   struct{}{},
			"WithAnalyzer":          struct{}{},
			"WithBatchedReduceSize": struct{}{},
			"WithDefaultOperator":   struct{}{},
			"WithDf":                struct{}{},
			"WithDocvalueFields":    struct{}{},
			"WithExpandWildcards":   struct{}{},
			"WithExplain":           struct{}{},
			"WithFielddataFields":   struct{}{},
			"WithFrom":              struct{}{},
			"WithIgnoreUnavailable": struct{}{},
			"WithLenient":           struct{}{},
			"WithPreference":        struct{}{},
			"WithQ":                 struct{}{},
			"WithRequestCache":      struct{}{},
			"WithRouting":           struct{}{},
			"WithScroll":            struct{}{},
			"WithSearchType":        struct{}{},
			"WithSize":              struct{}{},
			"WithSort":              struct{}{},
			"WithStats":             struct{}{},
			"WithStoredFields":      struct{}{},
			"WithSuggestField":      struct{}{},
			"WithSuggestMode":       struct{}{},
			"WithSuggestSize":       struct{}{},
			"WithSuggestText":       struct{}{},
			"WithTerminateAfter":    struct{}{},
			"WithTimeout":           struct{}{},
			"WithTrackScores":       struct{}{},
			"WithTypedKeys":         struct{}{},
			"WithVersion":           struct{}{},
			"WithBody":              struct{}{},
			"WithErrorTrace":        struct{}{},
			"WithFilterPath":        struct{}{},
			"WithHuman":             struct{}{},
			"WithIgnore":            struct{}{},
			"WithPretty":            struct{}{},
			"WithSourceParam":       struct{}{},
		},
		"DeleteByQuery": map[string]struct{}{
			"WithType":                struct{}{},
			"WithSource":              struct{}{},
			"WithSourceExclude":       struct{}{},
			"WithSourceInclude":       struct{}{},
			"WithAllowNoIndices":      struct{}{},
			"WithAnalyzeWildcard":     struct{}{},
			"WithAnalyzer":            struct{}{},
			"WithConflicts":           struct{}{},
			"WithDefaultOperator":     struct{}{},
			"WithDf":                  struct{}{},
			"WithExpandWildcards":     struct{}{},
			"WithFrom":                struct{}{},
			"WithIgnoreUnavailable":   struct{}{},
			"WithLenient":             struct{}{},
			"WithPreference":          struct{}{},
			"WithQ":                   struct{}{},
			"WithRefresh":             struct{}{},
			"WithRequestCache":        struct{}{},
			"WithRequestsPerSecond":   struct{}{},
			"WithRouting":             struct{}{},
			"WithScroll":              struct{}{},
			"WithScrollSize":          struct{}{},
			"WithSearchTimeout":       struct{}{},
			"WithSearchType":          struct{}{},
			"WithSize":                struct{}{},
			"WithSlices":              struct{}{},
			"WithSort":                struct{}{},
			"WithStats":               struct{}{},
			"WithTerminateAfter":      struct{}{},
			"WithTimeout":             struct{}{},
			"WithVersion":             struct{}{},
			"WithWaitForActiveShards": struct{}{},
			"WithWaitForCompletion":   struct{}{},
			"WithErrorTrace":          struct{}{},
			"WithFilterPath":          struct{}{},
			"WithHuman":               struct{}{},
			"WithIgnore":              struct{}{},
			"WithPretty":              struct{}{},
			"WithSourceParam":         struct{}{},
		},
	}
)
