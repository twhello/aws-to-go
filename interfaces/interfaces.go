package interfaces

import (
	"net/http"
	"net/url"
)

// AWSService Interface
type IAWSService interface {
	Endpoint() string
	RegionName() string
	ServiceName() string
	SignAndDo(req IAWSRequest, dto interface{}) (*http.Response, error)
}

// AWSRequest Interface
type IAWSRequest interface {
	BuildRequest() *http.Request
	FormPostValues() *url.Values
	Header() *http.Header
	Method() string
	QueryStringValues() *url.Values
	Url() string
}

// AWSCredentials Interface
type IAWSCredentials interface {
	AccessKeyId() string
	SecretKey() string
}

// AWS Service Error Interface
type IServiceError interface {
	Error() string
	Code() int
	Status() string
	ErrorType() string
	ErrorMessage() string
	IsRetry() bool
}

// Interface for XML and JSON decoding.
type IDecoder interface {
	Decode(v interface{}) error
}
