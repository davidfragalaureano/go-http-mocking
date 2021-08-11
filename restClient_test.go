package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/http-mocking/mocks"
	"github.com/http-mocking/restclient"

	"github.com/stretchr/testify/assert"
)

func TestMakeGetRequestSucess(t *testing.T) {

	body := createHTTPBody(`{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`)
	mockHTTPClient(&http.Response{
		StatusCode: 200,
		Body:       body,
	}, false)
	resp, error := MakeGetRequest("some-url")
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.EqualValues(t, `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`, resp)
}

func TestMakeGetRequestFailedHttpRquest(t *testing.T) {

	mockHTTPClient(&http.Response{
		StatusCode: 500,
		Status:     "Something bad happened",
	}, true)
	_, error := MakeGetRequest("some-url")

	if error == nil {
		t.Fatal(fmt.Printf("Error is nil: Expected: nil, Got: %s", error.Error()))
	}

	assert.EqualError(t, error, "Unable to request via GET: Something bad happended")
}

func TestMakeGetRequestInvalidBody(t *testing.T) {

	mockHTTPClient(&http.Response{
		StatusCode: 500,
		Status:     "Something bad happened",
	}, true)
	_, error := MakeGetRequest("some-url")

	if error == nil {
		t.Fatal(fmt.Printf("Error is nil: Expected: nil, Got: %s", error.Error()))
	}

	assert.EqualError(t, error, "Unable to request via GET: Something bad happended")
}

// Set http client mock
func mockHTTPClient(httpResponse *http.Response, errorThrown bool) {
	restclient.Client = &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if errorThrown {
				return httpResponse, errors.New("Something bad happended")
			}
			return httpResponse, nil
		},
	}
}

func createHTTPBody(json string) io.ReadCloser {
	// create a new reader with JSON value
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
