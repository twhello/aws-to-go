package cloudwatchlogs

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/*****************************************************************************/

// Creates a new log group with the specified name. The name of the log group
// must be unique within a region for an AWS account. You can create up to 500
// log groups per account.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogGroup.html]
type CreateLogGroupRequest struct {
	LogGroupName string `json:"logGroupName"`
}

// Creates a new CreateLogGroupRequest.
func NewCreateLogGroupRequest(logGroupName string) *CreateLogGroupRequest {
	return &CreateLogGroupRequest{logGroupName}
}

/*****************************************************************************/

// Creates a new log stream in the specified log group. The name of the log
// stream must be unique within the log group. There is no limit on the number
// of log streams that can exist in a log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogStream.html]
type CreateLogStreamRequest struct {
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

// Creates a new CreateLogStreamRequest.
func NewCreateLogStreamRequest(logGroupName, logStreamName string) *CreateLogStreamRequest {
	return &CreateLogStreamRequest{logGroupName, logStreamName}
}

/*****************************************************************************/

// Deletes the log group with the specified name and permanently deletes all the archived log events associated with it.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteLogGroup.html]
type DeleteLogGroupRequest struct {
	LogGroupName string `json:"logGroupName"`
}

// Creates a new DeleteLogGroupRequest.
func NewDeleteLogGroupRequest(logGroupName string) *DeleteLogGroupRequest {
	return &DeleteLogGroupRequest{logGroupName}
}

/*****************************************************************************/

// Deletes a log stream and permanently deletes all the archived log events associated with it.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteLogStream.html]
type DeleteLogStreamRequest struct {
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

// Creates a new DeleteLogStreamRequest.
func NewDeleteLogStreamRequest(logGroupName, logStreamName string) *DeleteLogStreamRequest {
	return &DeleteLogStreamRequest{logGroupName, logStreamName}
}

/*****************************************************************************/

// Deletes a metric filter associated with the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteMetricFilter.html]
type DeleteMetricFilterRequest struct {
	FilterName   string `json:"fFilterName"`
	LogGroupName string `json:"logGroupName"`
}

// Creates a new DeleteMetricFilterRequest.
func NewDeleteMetricFilterRequest(logGroupName, filterName string) *DeleteMetricFilterRequest {
	return &DeleteMetricFilterRequest{filterName, logGroupName}
}

/*****************************************************************************/

// Deletes the retention policy of the specified log group. Log events would not expire
// if they belong to log groups without a retention policy.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html]
type DeleteRetentionPolicyRequest struct {
	LogGroupName string `json:"logGroupName"`
}

// Creates a new DeleteRetentionPolicyRequest.
func NewDeleteRetentionPolicyRequest(logGroupName string) *DeleteRetentionPolicyRequest {
	return &DeleteRetentionPolicyRequest{logGroupName}
}

/*****************************************************************************/

// Returns all the log groups that are associated with the AWS account making the request.
// The list returned in the response is ASCII-sorted by log group name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogGroups.html]
type DescribeLogGroupsRequest struct {
	Limit              int    `json:"limit,omitempty"`
	LogGroupNamePrefix string `json:"logGroupNamePrefix,omitempty"`
	NextToken          string `json:"nextToken,omitempty"`
}

// Creates a new DescribeLogGroupsRequest.
func NewDescribeLogGroupsRequest() *DescribeLogGroupsRequest {
	return &DescribeLogGroupsRequest{}
}

/*****************************************************************************/

// Returns all the log streams that are associated with the specified log group.
// The list returned in the response is ASCII-sorted by log stream name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogStreams.html]
type DescribeLogStreamsRequest struct {
	Limit              int    `json:"limit,omitempty"`
	LogGroupName       string `json:"logGroupName"`
	LogGroupNamePrefix string `json:"logGroupNamePrefix,omitempty"`
	NextToken          string `json:"nextToken,omitempty"`
}

// Creates a new DescribeLogStreamsRequest.
func NewDescribeLogStreamsRequest(logGroupName string) *DescribeLogStreamsRequest {
	return &DescribeLogStreamsRequest{LogGroupName: logGroupName}
}

/*****************************************************************************/

// Returns all the metrics filters associated with the specified log group.
// The list returned in the response is ASCII-sorted by filter name.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeMetricFilters.html]
type DescribeMetricFiltersRequest struct {
	FilterNamePrefix string `json:"filterNamePrefix,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	LogGroupName     string `json:"logGroupName"`
	NextToken        string `json:"nextToken,omitempty"`
}

// Creates a new DescribeMetricFiltersRequest.
func NewDescribeMetricFiltersRequest() *DescribeMetricFiltersRequest {
	return &DescribeMetricFiltersRequest{}
}

/*****************************************************************************/

// Retrieves log events from the specified log stream. You can provide an optional
// time range to filter the results on the event timestamp.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html]
type GetLogEventsRequest struct {
	EndTime       *datetime.JsonDate `json:"endTime,omitempty"`
	Limit         int                `json:"lLimit,omitempty"`
	LogGroupName  string             `json:"logGroupName"`
	LogStreamName string             `json:"logStreamName"`
	NextToken     string             `json:"nextToken,omitempty"`
	StartFromHead bool               `json:"startFromHead,omitempty"`
	StartTime     *datetime.JsonDate `json:"startTime,omitempty"`
}

// Creates a new GetLogEventsRequest.
func NewGetLogEventsRequest(logGroupName, logStreamName string) *GetLogEventsRequest {
	return &GetLogEventsRequest{LogGroupName: logGroupName, LogStreamName: logStreamName}
}

/*****************************************************************************/

// Uploads a batch of log events to the specified log stream.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html]
type PutLogEventsRequest struct {
	LogEvents     []InputLogEvent `json:"logEvents"`
	LogGroupName  string          `json:"logGroupName"`
	LogStreamName string          `json:"logStreamName"`
	SequenceToken string          `json:"sequenceToken,omitempty"`
}

// Creates a new PutLogEventsRequest.
func NewPutLogEventsRequest() *PutLogEventsRequest {
	return &PutLogEventsRequest{}
}

/*****************************************************************************/

// Creates or updates a metric filter and associates it with the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutMetricFilter.html]
type PutMetricFilterRequest struct {
	FilterName            string                 `json:"filterName"`
	FilterPattern         string                 `json:"filterPattern"`
	LogGroupName          string                 `json:"logGroupName"`
	MetricTransformations []MetricTransformation `json:"metricTransformations"`
}

// Creates a new PutMetricFilterRequest.
func NewPutMetricFilterRequest(logGroupName, filterName, filterPattern string, metricTransformations ...MetricTransformation) *PutMetricFilterRequest {
	return &PutMetricFilterRequest{logGroupName, filterName, filterPattern, metricTransformations}
}

/*****************************************************************************/

// Sets the retention of the specified log group. A retention policy allows you to configure
// the number of days you want to retain log events in the specified log group.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html]
type PutRetentionPolicyRequest struct {
	LogGroupName    string `json:"logGroupName"`
	RetentionInDays int    `json:"retentionInDays"`
}

// Creates a new PutRetentionPolicyRequest.
func NewPutRetentionPolicyRequest(logGroupName string, retentionInDays int) *PutRetentionPolicyRequest {
	return &PutRetentionPolicyRequest{logGroupName, retentionInDays}
}

/*****************************************************************************/

// Tests the filter pattern of a metric filter against a sample of log event messages.
// You can use this operation to validate the correctness of a metric filter pattern.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TestMetricFilter.html]
type TestMetricFilterRequest struct {
	FilterPattern    string   `json:"filterPattern"`
	LogEventMessages []string `json:"logEventMessages"`
}

// Creates a new TestMetricFilterRequest.
func NewTestMetricFilterRequest(filterPattern string, logEventMessages ...string) *TestMetricFilterRequest {
	return &TestMetricFilterRequest{filterPattern, logEventMessages}
}
