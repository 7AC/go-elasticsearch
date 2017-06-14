// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// HealthOption is a non-required Health option that gets applied to an HTTP request.
type HealthOption func(r *transport.Request)

// WithHealthIndex - limit the information returned to a specific index.
func WithHealthIndex(index []string) HealthOption {
	return func(r *transport.Request) {
	}
}

// HealthLevel - specify the level of detail for returned information.
type HealthLevel int

const (
	// HealthLevelCluster can be used to set HealthLevel to "cluster"
	HealthLevelCluster = iota
	// HealthLevelIndices can be used to set HealthLevel to "indices"
	HealthLevelIndices = iota
	// HealthLevelShards can be used to set HealthLevel to "shards"
	HealthLevelShards = iota
)

// WithHealthLevel - specify the level of detail for returned information.
func WithHealthLevel(level HealthLevel) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthLocal - return local information, do not retrieve the state from master node (default: false).
func WithHealthLocal(local bool) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthMasterTimeout - explicit operation timeout for connection to master node.
func WithHealthMasterTimeout(masterTimeout time.Duration) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthTimeout - explicit operation timeout.
func WithHealthTimeout(timeout time.Duration) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthWaitForActiveShards - wait until the specified number of shards is active.
func WithHealthWaitForActiveShards(waitForActiveShards string) HealthOption {
	return func(r *transport.Request) {
	}
}

// HealthWaitForEvents - wait until all currently queued events with the given priority are processed.
type HealthWaitForEvents int

const (
	// HealthWaitForEventsImmediate can be used to set HealthWaitForEvents to "immediate"
	HealthWaitForEventsImmediate = iota
	// HealthWaitForEventsUrgent can be used to set HealthWaitForEvents to "urgent"
	HealthWaitForEventsUrgent = iota
	// HealthWaitForEventsHigh can be used to set HealthWaitForEvents to "high"
	HealthWaitForEventsHigh = iota
	// HealthWaitForEventsNormal can be used to set HealthWaitForEvents to "normal"
	HealthWaitForEventsNormal = iota
	// HealthWaitForEventsLow can be used to set HealthWaitForEvents to "low"
	HealthWaitForEventsLow = iota
	// HealthWaitForEventsLanguid can be used to set HealthWaitForEvents to "languid"
	HealthWaitForEventsLanguid = iota
)

// WithHealthWaitForEvents - wait until all currently queued events with the given priority are processed.
func WithHealthWaitForEvents(waitForEvents HealthWaitForEvents) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthWaitForNoRelocatingShards - whether to wait until there are no relocating shards in the cluster.
func WithHealthWaitForNoRelocatingShards(waitForNoRelocatingShards bool) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthWaitForNodes - wait until the specified number of nodes is available.
func WithHealthWaitForNodes(waitForNodes string) HealthOption {
	return func(r *transport.Request) {
	}
}

// HealthWaitForStatus - wait until cluster is in a specific state.
type HealthWaitForStatus int

const (
	// HealthWaitForStatusGreen can be used to set HealthWaitForStatus to "green"
	HealthWaitForStatusGreen = iota
	// HealthWaitForStatusYellow can be used to set HealthWaitForStatus to "yellow"
	HealthWaitForStatusYellow = iota
	// HealthWaitForStatusRed can be used to set HealthWaitForStatus to "red"
	HealthWaitForStatusRed = iota
)

// WithHealthWaitForStatus - wait until cluster is in a specific state.
func WithHealthWaitForStatus(waitForStatus HealthWaitForStatus) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthErrorTrace - include the stack trace of returned errors.
func WithHealthErrorTrace(errorTrace bool) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthFilterPath - a comma-separated list of filters used to reduce the respone.
func WithHealthFilterPath(filterPath []string) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthHuman - return human readable values for statistics.
func WithHealthHuman(human bool) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthIgnore - ignores the specified HTTP status codes.
func WithHealthIgnore(ignore []int) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthPretty - pretty format the returned JSON response.
func WithHealthPretty(pretty bool) HealthOption {
	return func(r *transport.Request) {
	}
}

// WithHealthSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithHealthSourceParam(sourceParam string) HealthOption {
	return func(r *transport.Request) {
	}
}

// Health - the cluster health API allows to get a very simple status on the health of the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-health.html for more info.
//
// options: optional parameters.
func (c *Cluster) Health(options ...HealthOption) (*HealthResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &HealthResponse{resp}, err
}

// HealthResponse is the response for Health.
type HealthResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *HealthResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
