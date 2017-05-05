// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// GetSettings - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-get-settings.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithName, WithAllowNoIndices, WithExpandWildcards, WithFlatSettings, WithIgnoreUnavailable, WithIncludeDefaults, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetSettings(options ...*Option) (*GetSettingsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["GetSettings"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &GetSettingsResponse{resp}, err
}

// GetSettingsResponse is the response for GetSettings.
type GetSettingsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetSettingsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
