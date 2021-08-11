package mocks

import "net/http"

// MockClient contract interface
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do contract method
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
