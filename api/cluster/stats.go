// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// StatsOption is a non-required Stats option that gets applied to an HTTP request.
type StatsOption func(r *transport.Request)

// WithStatsNodeID - a comma-separated list of node IDs or names to limit the returned information; use "_local" to return information from the node you're connecting to, leave empty to get information from all nodes.
func WithStatsNodeID(nodeID []string) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsFlatSettings - return settings in flat format (default: false).
func WithStatsFlatSettings(flatSettings bool) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsTimeout - explicit operation timeout.
func WithStatsTimeout(timeout time.Duration) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsErrorTrace - include the stack trace of returned errors.
func WithStatsErrorTrace(errorTrace bool) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsFilterPath - a comma-separated list of filters used to reduce the respone.
func WithStatsFilterPath(filterPath []string) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsHuman - return human readable values for statistics.
func WithStatsHuman(human bool) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsIgnore - ignores the specified HTTP status codes.
func WithStatsIgnore(ignore []int) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsPretty - pretty format the returned JSON response.
func WithStatsPretty(pretty bool) StatsOption {
	return func(r *transport.Request) {
	}
}

// WithStatsSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithStatsSourceParam(sourceParam string) StatsOption {
	return func(r *transport.Request) {
	}
}

// Stats - the Cluster Stats API allows to retrieve statistics from a cluster wide perspective. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-stats.html for more info.
//
// options: optional parameters.
func (c *Cluster) Stats(options ...StatsOption) (*StatsResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &StatsResponse{resp}, err
}

// StatsResponse is the response for Stats.
type StatsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *StatsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
