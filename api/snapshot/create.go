// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// CreateOption is a non-required Create option that gets applied to an HTTP request.
type CreateOption func(r *transport.Request)

// WithCreateMasterTimeout - explicit operation timeout for connection to master node.
func WithCreateMasterTimeout(masterTimeout time.Duration) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateWaitForCompletion - should this request wait until the operation has completed before returning.
func WithCreateWaitForCompletion(waitForCompletion bool) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateBody - the snapshot definition.
func WithCreateBody(body map[string]interface{}) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateErrorTrace - include the stack trace of returned errors.
func WithCreateErrorTrace(errorTrace bool) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithCreateFilterPath(filterPath []string) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateHuman - return human readable values for statistics.
func WithCreateHuman(human bool) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateIgnore - ignores the specified HTTP status codes.
func WithCreateIgnore(ignore []int) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreatePretty - pretty format the returned JSON response.
func WithCreatePretty(pretty bool) CreateOption {
	return func(r *transport.Request) {
	}
}

// WithCreateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithCreateSourceParam(sourceParam string) CreateOption {
	return func(r *transport.Request) {
	}
}

// Create - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a snapshot name.
//
// options: optional parameters.
func (s *Snapshot) Create(repository string, snapshot string, options ...CreateOption) (*CreateResponse, error) {
	req := s.transport.NewRequest("PUT")
	for _, option := range options {
		option(req)
	}
	resp, err := s.transport.Do(req)
	return &CreateResponse{resp}, err
}

// CreateResponse is the response for Create.
type CreateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *CreateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
