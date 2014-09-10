/*
The following AWS services are available:

- Auto Scaling
- Amazon CloudWatch
- Amazon Cognito
- AWS Data Pipeline
- Amazon DynamoDB
- Amazon EC2
- Amazon Kinesis
- Amazon Simple Storage Service (S3)
- Amazon Simple Email Service (SES)
- Amazon SimpleDB
- Amazon Simple Notification Service (SNS)
- Amazon Simple Queue Service (SQS)
- Amazon Simple Workflow Service (SWF)
*/
package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/twhello/aws-to-go/interfaces"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost:   25,
		ResponseHeaderTimeout: 30 * time.Second,
	},
}

// Returns the shared http.Client.
// Default settings: 25 MaxIdleConnsPerHost and 30 second ResponseHeaderTimeout.
func HttpClient() *http.Client {
	configSetting.m.RLock()
	defer configSetting.m.RUnlock()
	return httpClient
}

// Submits the signed request to AWS and, if not nil, unmarshals the XML
// response to the 'dto' interface, or returns an error or ServiceError.
func DoRequest(awsreq interfaces.IAWSRequest, dto interface{}, eval *EvalServiceResponse) (resp *http.Response, err interfaces.IServiceError) {

	config := Config()
	isDebug := config.IsDebugging()
	resp = nil
	err = nil
	req := awsreq.BuildRequest()

	if isDebug {
		log.Printf("REQUEST  > %+v \n", req)
	}

	RETRY_ATTEMPTS := config.RetryAttempts()
	retries := uint(0)

RETRY:

	resp, e := HttpClient().Do(req)
	if e != nil {
		err = NewServiceError(100, "100 HTTP Error", "", e.Error())
		return nil, err
	}

	resp, err = consumeResponse(resp, eval)

	if isDebug {
		log.Printf("RESPONSE > %+v \nERROR    > %+v \n", resp, err)
	}

	if err == nil {

		if dto != nil {
			defer resp.Body.Close()
			e := eval.Decode(resp.Body, dto)
			if e != nil {
				err = NewServiceError(101, "101 IO Read Error", "Decode", e.Error())
			}
		}

	} else if err.IsRetry() && retries < RETRY_ATTEMPTS {

		if isDebug {
			log.Printf("RETRY   > %s of %s in %s milliseconds.\n", (retries + 1), RETRY_ATTEMPTS, (1 << retries * 100))
		}
		time.Sleep(time.Millisecond * (1 << retries * 100))
		retries++
		goto RETRY
	}

	return
}

func consumeResponse(response *http.Response, eval *EvalServiceResponse) (*http.Response, interfaces.IServiceError) {

	if response.StatusCode >= 400 {

		srvErr := NewServiceError(response.StatusCode, response.Status, "", "")
		defer response.Body.Close()
		eval.Decode(response.Body, srvErr)
		srvErr.SetRetry(eval.Matches(response.StatusCode, srvErr.ErrorType()))
		
		return response, srvErr
	}

	return response, nil
}

/*****************************************************************************/

// A collection of retriable codes and errors.
type EvalServiceResponse struct {
	Codes   []int
	Errors  []string
	Decoder func(io.Reader, interface{}) error
}

// Creates the default EvalServiceResponse for XML responses:
//	services.NewEvalServiceResponse(
//		func(r io.Reader, v interface{})error { return xml.NewDecoder(r).Decode(v) },
//		[]int{500, 503},
//		[]string{"Throttling"},
//	)
func NewEvalXmlServiceResponse() *EvalServiceResponse {
	return NewEvalServiceResponse(
		func(r io.Reader, v interface{}) error { return xml.NewDecoder(r).Decode(v) },
		[]int{500, 503},
		[]string{"Throttling"},
	)
}

// Creates the default EvalServiceResponse for JSON responses:
//	services.NewEvalServiceResponse(
//		func(r io.Reader, v interface{})error { return json.NewDecoder(r).Decode(v) },
//		[]int{500, 503},
//		[]string{"Throttling"},
//	)
func NewEvalJsonServiceResponse() *EvalServiceResponse {
	return NewEvalServiceResponse(
		func(r io.Reader, v interface{}) error { return json.NewDecoder(r).Decode(v) },
		[]int{500, 503},
		[]string{"Throttling"},
	)
}

// Creates a new EvalServiceResponse struct.
func NewEvalServiceResponse(decoder func(io.Reader, interface{}) error, codes []int, errors []string) *EvalServiceResponse {
	sort.Ints(codes)
	sort.Strings(errors)
	return &EvalServiceResponse{codes, errors, decoder}
}

// Decodes the service response.Body into the given response struct.
func (e EvalServiceResponse) Decode(r io.Reader, v interface{}) error {
	return e.Decoder(r, v)
}

// Returns true if the collection contains the code or error.
// Note: Errors match using strings.Contains().
func (r *EvalServiceResponse) Matches(code int, err string) bool {

	if r.Codes != nil && sort.SearchInts(r.Codes, code) < len(r.Codes) {
		return true
	}

	if r.Errors != nil {
		for _, e := range r.Errors {
			if strings.Contains(err, e) {
				return true
			}
		}
	}

	return false
}

/*****************************************************************************/

// General Service Error.
type ServiceError struct {
	ErrCode    int    `xml:"-" json:"-"`
	ErrStatus  string `xml:"-" json:"-"`
	ErrType    string `xml:"Error>Code" json:"__type"`
	ErrMessage string `xml:"Error>Message" json:"message"`
	isRetry    bool
}

// Creates a new ServiceError.
func NewServiceError(code int, status, errType, errMessage string) *ServiceError {
	return &ServiceError{code, status, errType, errMessage, false}
}

func (err ServiceError) SetRetry(val bool) {
	err.isRetry = val
}

func (err ServiceError) Code() int {
	return err.ErrCode
}

func (err ServiceError) Status() string {
	return err.ErrStatus
}

func (err ServiceError) ErrorType() string {
	return err.ErrType
}

func (err ServiceError) ErrorMessage() string {
	return err.ErrMessage
}

func (err ServiceError) IsRetry() bool {
	return err.isRetry
}

func (err ServiceError) Error() string {
	return fmt.Sprintf("Code: %d, Status: %s, Type: %s, Message: %s \n", err.ErrCode, err.ErrStatus, err.ErrType, err.ErrMessage)
}
