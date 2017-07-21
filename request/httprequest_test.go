package request_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/EscherAuth/escher/request"
	"github.com/stretchr/testify/assert"
)

func TestNewFromHTTPRequest(t *testing.T) {

	httpRequest, err := http.NewRequest("GET", "/?k=p", bytes.NewBuffer([]byte("Hello, World!")))

	if err != nil {
		t.Fatal(err)
	}

	httpRequest.Header.Set("X-Testing", "OK")

	escherReqest, err := request.NewFromHTTPRequest(httpRequest)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, escherReqest.Path(), "/")
	assert.Equal(t, escherReqest.Body(), "Hello, World!")
	assert.Equal(t, escherReqest.Method(), "GET")
	assert.Equal(t, escherReqest.RawURL(), "/?k=p")
	assert.Equal(t, escherReqest.Expires(), 0)
	assert.Equal(t, request.Query{[2]string{"k", "p"}}, escherReqest.Query())
	assert.Equal(t, request.Headers{[2]string{"X-Testing", "OK"}}, escherReqest.Headers())

}

func TestNewFromHTTPRequest_EscherRequestMade_HTTPBodyStillContainsValueLikeItIsUnTouched(t *testing.T) {

	expectedBodyString := "Hello, World!"
	httpRequest, err := http.NewRequest("GET", "/?k=p", bytes.NewBuffer([]byte(expectedBodyString)))

	if err != nil {
		t.Fatal(err)
	}

	httpRequest.Header.Set("X-Testing", "OK")

	escherReqest, err := request.NewFromHTTPRequest(httpRequest)

	if err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadAll(httpRequest.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(content), expectedBodyString)
	assert.Equal(t, escherReqest.Body(), expectedBodyString)
	assert.Equal(t, string(content), escherReqest.Body())

}

func TestHTTPRequest(t *testing.T) {

	bodyIO := bytes.NewBuffer([]byte("Hello you awesome!"))
	expectedHTTPRequest, err := http.NewRequest("GET", "/?k=p", bodyIO)

	if err != nil {
		t.Fatal(err)
	}

	expectedHTTPRequest.Header.Set("X-Testing", "OK")

	escherReqest, err := request.NewFromHTTPRequest(expectedHTTPRequest)

	if err != nil {
		t.Fatal(err)
	}

	actuallyHTTPRequest, err := escherReqest.HTTPRequest()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedHTTPRequest.Method, actuallyHTTPRequest.Method)
	assert.Equal(t, expectedHTTPRequest.URL, actuallyHTTPRequest.URL)
	assert.Equal(t, expectedHTTPRequest.Proto, actuallyHTTPRequest.Proto)
	assert.Equal(t, expectedHTTPRequest.ProtoMajor, actuallyHTTPRequest.ProtoMajor)
	assert.Equal(t, expectedHTTPRequest.ProtoMinor, actuallyHTTPRequest.ProtoMinor)
	assert.Equal(t, expectedHTTPRequest.Header, actuallyHTTPRequest.Header)
	assert.Equal(t, expectedHTTPRequest.ContentLength, actuallyHTTPRequest.ContentLength)
	assert.Equal(t, expectedHTTPRequest.TransferEncoding, actuallyHTTPRequest.TransferEncoding)
	assert.Equal(t, expectedHTTPRequest.Close, actuallyHTTPRequest.Close)
	assert.Equal(t, expectedHTTPRequest.Form, actuallyHTTPRequest.Form)
	assert.Equal(t, expectedHTTPRequest.PostForm, actuallyHTTPRequest.PostForm)
	assert.Equal(t, expectedHTTPRequest.MultipartForm, actuallyHTTPRequest.MultipartForm)
	assert.Equal(t, expectedHTTPRequest.Trailer, actuallyHTTPRequest.Trailer)
	assert.Equal(t, expectedHTTPRequest.RemoteAddr, actuallyHTTPRequest.RemoteAddr)
	assert.Equal(t, expectedHTTPRequest.RequestURI, actuallyHTTPRequest.RequestURI)
	assert.Equal(t, expectedHTTPRequest.TLS, actuallyHTTPRequest.TLS)
	assert.Equal(t, expectedHTTPRequest.Cancel, actuallyHTTPRequest.Cancel)
	assert.Equal(t, expectedHTTPRequest.Response, actuallyHTTPRequest.Response)

	eBodyBuffer, _ := expectedHTTPRequest.GetBody()
	expectedBody, _ := ioutil.ReadAll(eBodyBuffer)

	aBodyBuffer, _ := actuallyHTTPRequest.GetBody()
	actuallyBody, _ := ioutil.ReadAll(aBodyBuffer)

	assert.Equal(t, expectedBody, actuallyBody)

}
