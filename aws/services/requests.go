package services

import (
	"github.com/twhello/aws-to-go/aws/util/netutil"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Content-Types for AWSRequest.
const (
	CONTENT_TYPE_TEXT_PLAIN                  = "text/plain"
	CONTENT_TYPE_APPLICATION_JSON            = "application/json"
	CONTENT_TYPE_APPLICATION_AMZ_JSON        = "application/x-amz-json-1.0"
	CONTENT_TYPE_APPLICATION_XML             = "application/xml"
	CONTENT_TYPE_APPLICATION_FORM_URLENCODED = "application/x-www-form-urlencoded; charset=utf-8"
	CONTENT_TYPE_BINARY_OCTET_STREAM         = "binary/octet-stream"
)

/******************************************************************************
 * AWS Request object implements interfaces.IAWSRequest
 */

// AWSRequest struct.
type AWSRequest struct {
	method    string
	u         *url.URL
	queryVals *url.Values
	postVals  *url.Values
	header    *http.Header
	body      interface{}
}

// Creates a new AWSRequest.
// method expects GET, DELETE or HEAD
// The arg 'queryString' can be a struct or nil.
func NewClientRequest(method, rawurl string, queryString interface{}) (*AWSRequest, error) {

	u, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return nil, err
	}

	var tv url.Values
	fmt.Println(queryString)
	if queryString != nil {
		tv = netutil.MarshalValues(queryString)
	}

	hdr := &http.Header{}
	hdr.Add("Content-Type", CONTENT_TYPE_TEXT_PLAIN)

	return &AWSRequest{
		strings.ToUpper(method),
		u,
		&tv,
		nil,
		hdr,
		nil,
	}, nil
}

// method expects POST or PUT
// contentType expects a constant, e.g. CONTENT_TYPE_TEXT_PLAIN.
// body expects string, io.Reader, []byte or a struct. Will panic otherwise on AWSRequest.BuildRequest().
func NewServerRequest(method, rawurl string, body interface{}) (*AWSRequest, error) {

	u, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return nil, err
	}

	return &AWSRequest{
		strings.ToUpper(method),
		u,
		nil,
		nil,
		&http.Header{},
		body,
	}, nil
}

func (r *AWSRequest) Method() string {
	return r.method
}

func (r *AWSRequest) Url() string {
	return r.u.RequestURI()
}

func (r *AWSRequest) QueryStringValues() *url.Values {
	return r.queryVals
}

func (r *AWSRequest) FormPostValues() *url.Values {
	return r.postVals
}

func (r *AWSRequest) Header() *http.Header {
	return r.header
}

func (r *AWSRequest) BuildRequest() *http.Request {

	if r.queryVals != nil {
		r.u.RawQuery = r.queryVals.Encode()
	}

	body, length := r.buildBody()

	return &http.Request{
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Close:         true,
		Method:        r.method,
		URL:           r.u,
		Host:          r.u.Host,
		Header:        *r.header,
		ContentLength: length,
		Body:          body,
	}
}

func (r *AWSRequest) buildBody() (io.ReadCloser, int64) {

	if r.body == nil {
		return ioutil.NopCloser(strings.NewReader(``)), 0
	}

	switch rb := r.body.(type) {
	case []byte:
		return ioutil.NopCloser(bytes.NewReader(rb)), int64(len(rb))

	case string:
		return ioutil.NopCloser(strings.NewReader(rb)), int64(len([]byte(rb)))

	case io.Reader:
		b, _ := ioutil.ReadAll(rb)
		r.body = b
		return ioutil.NopCloser(bytes.NewReader(b)), int64(len(b))

	case io.ReadCloser:
		b, _ := ioutil.ReadAll(rb)
		rb.Close()
		r.body = b
		return ioutil.NopCloser(bytes.NewReader(b)), int64(len(b))
	}

	var b []byte
	var err error

	switch r.Header().Get("Content-Type") {
	case CONTENT_TYPE_APPLICATION_XML:
		b, err = xml.Marshal(r.body)
	default:
		b, err = json.Marshal(r.body)
	}

	fmt.Printf("services.requests.go: %s \n", b)

	if err != nil {
		panic(err.Error())
	}

	return ioutil.NopCloser(bytes.NewReader(b)), int64(len(b))
}
