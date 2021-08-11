package common

import "net/http"

// Interface for http contract functionality
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
