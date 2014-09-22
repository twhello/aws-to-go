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
	"math"
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

	resp = nil
	err = nil
	config := Config()
	isDebug := config.IsDebugging()

	RETRY_ATTEMPTS := config.RetryAttempts()
	retries := uint(0)

RETRY:

	req := awsreq.BuildRequest()

	if isDebug {
		log.Printf("\nREQUEST  > %+v \n", req)
	}

	resp, e := HttpClient().Do(req)
	if e != nil {
		return nil, NewServiceError(100, "100 HTTP Error", "", e.Error())
	}

	resp, err = evalResponse(resp, eval)

	if isDebug {
		log.Printf("\nRESPONSE > %+v \n", resp)
		log.Printf("\nERROR    > %+v \n", err)
	}

	if err == nil {

		if dto != nil {
			defer resp.Body.Close()
			e := eval.Decode(resp.Body, dto)
			if e != nil {
				err = NewServiceError(101, "101 IO Read Error", "Decode", e.Error())
			}
		}

	} else {

		if err.IsRetry() && retries < RETRY_ATTEMPTS {

			retries++
			sleep := math.Min(float64(2^retries*50), 2000)

			if isDebug {
				log.Printf("\nRETRY   > %d of %d in %d milliseconds.\n", retries, RETRY_ATTEMPTS, sleep)
			}
			time.Sleep(time.Millisecond * time.Duration(sleep))
			goto RETRY
		}
	}

	return
}

func evalResponse(response *http.Response, eval *EvalServiceResponse) (*http.Response, interfaces.IServiceError) {

	if response.StatusCode >= 400 {

		srvErr := NewServiceError(response.StatusCode, response.Status, "", "")
		eval.Decode(response.Body, srvErr)
		response.Body.Close()
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
func (r *EvalServiceResponse) Matches(code int, errorType string) bool {

	if r.Codes != nil {
		for _, v := range r.Codes {
			if v == code {
				return true
			}
		}
	}

	if r.Errors != nil {
		for _, e := range r.Errors {
			if strings.Contains(errorType, e) {
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
	isRetry    bool   `xml:"-" json:"-"`
}

// Creates a new ServiceError.
func NewServiceError(code int, status, errType, errMessage string) *ServiceError {
	return &ServiceError{code, status, errType, errMessage, false}
}

func (err *ServiceError) SetRetry(val bool) {
	err.isRetry = val
}

func (err *ServiceError) Code() int {
	return err.ErrCode
}

func (err *ServiceError) Status() string {
	return err.ErrStatus
}

func (err *ServiceError) ErrorType() string {
	return err.ErrType
}

func (err *ServiceError) ErrorMessage() string {
	return err.ErrMessage
}

func (err *ServiceError) IsRetry() bool {
	return err.isRetry
}

func (err *ServiceError) Error() string {
	return fmt.Sprintf("Code: %d, Status: %s, Type: %s, Message: %s, Retry: %t \n",
		err.ErrCode, err.ErrStatus, err.ErrType, err.ErrMessage, err.isRetry)
}
