// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Status - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// options: optional parameters. Supports the following functional options: WithRepository, WithSnapshot, WithIgnoreUnavailable, WithMasterTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (s *Snapshot) Status(options ...*Option) (*StatusResponse, error) {
	req := s.transport.NewRequest("GET")
	methodOptions := supportedOptions["Status"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := s.transport.Do(req)
	return &StatusResponse{resp}, err
}

// StatusResponse is the response for Status.
type StatusResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *StatusResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
