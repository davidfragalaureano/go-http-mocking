package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/http-mocking/common"
)

// Holds the value of our http client implementation,
// whether is a real implementation or a mock
var (
	Client common.HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Get sends a post request to the URL with the body
func Get(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(jsonBytes))

	if err != nil {
		return nil, err
	}

	request.Header = headers
	return Client.Do(request)
}
