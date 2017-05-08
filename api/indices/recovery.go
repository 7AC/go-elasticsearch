// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
)

// Recovery - the indices recovery API provides insight into on-going index shard recoveries. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-recovery.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithActiveOnly, WithDetailed, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Recovery(options ...*Option) (*RecoveryResponse, error) {
	req := i.transport.NewRequest("GET")
	methodOptions := supportedOptions["Recovery"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &RecoveryResponse{resp}, err
}

// RecoveryResponse is the response for Recovery.
type RecoveryResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *RecoveryResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
