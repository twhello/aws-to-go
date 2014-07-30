//
// Amazon CloudWatch Logs enables you to monitor, store, and access your system,
// application, and custom log files.
//
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/Welcome.html]
//
package cloudwatchlogs

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"net/http"
)

const ServiceName = "logs"

// CloudWatchLogsService describes the API interface. Instantiate with cloudwatchlogs.NewService().
type CloudWatchLogsService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *CloudWatchLogsService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *CloudWatchLogsService) RegionName() string {
	return s.region.Name()
}

func (s *CloudWatchLogsService) Endpoint() string {
	return s.endpoint
}

// Low-level request to CloudWatchLogs service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *CloudWatchLogsService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *CloudWatchLogsService) wrapperSignAndDo(target string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	if err == nil {
		h := req.Header()
		h.Set("Accept", "application/json")
		h.Set("Connection", "Keep-Alive")
		h.Set("Content-Type", "application/x-amz-json-1.1")
		h.Set("X-Amz-Target", target)
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * CloudWatch Logs Service Methods.
 */

// Creates a new log group with the specified name. The name of the log group
// must be unique within a region for an AWS account. You can create up to 500
// log groups per account.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogGroup.html]
func (s *CloudWatchLogsService) CreateLogGroup(req *CreateLogGroupRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.CreateLogGroup", req, nil)
	return
}

// Creates a new log stream in the specified log group. The name of the log
// stream must be unique within the log group. There is no limit on the number
// of log streams that can exist in a log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogStream.html]
func (s *CloudWatchLogsService) CreateLogStream(req *CreateLogStreamRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.CreateLogStream", req, nil)
	return
}

// Deletes the log group with the specified name and permanently deletes all the archived log events associated with it.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteLogGroup.html]
func (s *CloudWatchLogsService) DeleteLogGroup(req *DeleteLogGroupRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.DeleteLogGroup", req, nil)
	return
}

// Deletes a log stream and permanently deletes all the archived log events associated with it.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteLogStream.html]
func (s *CloudWatchLogsService) DeleteLogStream(req *DeleteLogStreamRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.DeleteLogStream", req, nil)
	return
}

// Deletes a metric filter associated with the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteMetricFilter.html]
func (s *CloudWatchLogsService) DeleteMetricFilter(req *DeleteMetricFilterRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.DeleteMetricFilter", req, nil)
	return
}

// Deletes the retention policy of the specified log group. Log events would not expire
// if they belong to log groups without a retention policy.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html]
func (s *CloudWatchLogsService) DeleteRetentionPolicy(req *DeleteRetentionPolicyRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.DeleteRetentionPolicy", req, nil)
	return
}

// Returns all the log groups that are associated with the AWS account making the request.
// The list returned in the response is ASCII-sorted by log group name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogGroups.html]
func (s *CloudWatchLogsService) DescribeLogGroups(req *DescribeLogGroupsRequest) (result *DescribeLogGroupsResult, err error) {

	result = new(DescribeLogGroupsResult)
	err = s.wrapperSignAndDo("Logs_20140328.DescribeLogGroups", req, result)
	return
}

// Returns all the log streams that are associated with the specified log group.
// The list returned in the response is ASCII-sorted by log stream name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogStreams.html]
func (s *CloudWatchLogsService) DescribeLogStreams(req *DescribeLogStreamsRequest) (result *DescribeLogStreamsResult, err error) {

	result = new(DescribeLogStreamsResult)
	err = s.wrapperSignAndDo("Logs_20140328.DescribeLogStreams", req, result)
	return
}

// Returns all the metrics filters associated with the specified log group.
// The list returned in the response is ASCII-sorted by filter name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeMetricFilters.html]
func (s *CloudWatchLogsService) DescribeMetricFilters(req *DescribeMetricFiltersRequest) (result *DescribeMetricFiltersResult, err error) {

	result = new(DescribeMetricFiltersResult)
	err = s.wrapperSignAndDo("Logs_20140328.DescribeMetricFilters", req, result)
	return
}

// Retrieves log events from the specified log stream. You can provide an optional
// time range to filter the results on the event timestamp.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html]
func (s *CloudWatchLogsService) GetLogEvents(req *GetLogEventsRequest) (result *GetLogEventsResult, err error) {

	result = new(GetLogEventsResult)
	err = s.wrapperSignAndDo("Logs_20140328.GetLogEvents", req, result)
	return
}

// Uploads a batch of log events to the specified log stream.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html]
func (s *CloudWatchLogsService) PutLogEvents(req *PutLogEventsRequest) (result *PutLogEventsResult, err error) {

	result = new(PutLogEventsResult)
	err = s.wrapperSignAndDo("Logs_20140328.PutLogEvents", req, result)
	return
}

// Creates or updates a metric filter and associates it with the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutMetricFilter.html]
func (s *CloudWatchLogsService) PutMetricFilter(req *PutMetricFilterRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.PutMetricFilter", req, nil)
	return
}

// Sets the retention of the specified log group. A retention policy allows you to configure
// the number of days you want to retain log events in the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html]
func (s *CloudWatchLogsService) PutRetentionPolicy(req *PutRetentionPolicyRequest) (err error) {

	err = s.wrapperSignAndDo("Logs_20140328.PutRetentionPolicy", req, nil)
	return
}

// Tests the filter pattern of a metric filter against a sample of log event messages.
// You can use this operation to validate the correctness of a metric filter pattern.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TestMetricFilter.html]
func (s *CloudWatchLogsService) TestMetricFilter(req *TestMetricFilterRequest) (result *TestMetricFilterResult, err error) {

	result = new(TestMetricFilterResult)
	err = s.wrapperSignAndDo("Logs_20140328.TestMetricFilter", req, result)
	return
}

// Creates a new CloudWatch Logs Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *CloudWatchLogsService {
	return &CloudWatchLogsService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
