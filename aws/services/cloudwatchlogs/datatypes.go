package cloudwatchlogs

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/******************************************************************************
 * Data Types
 */

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogGroupsResult.html]
type DescribeLogGroupsResult struct {
	LogGroups []LogGroup `json:"logGroups,omitempty"`
	NextToken string     `json:"nextToken,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogStreamsResult.html]
type DescribeLogStreamsResult struct {
	LogStreams []LogStream `json:"logStreams,omitempty"`
	NextToken  string      `json:"nextToken,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeMetricFiltersResult.html]
type DescribeMetricFiltersResult struct {
	MetricFilters []MetricFilter `json:"metricFilters,omitempty"`
	NextToken     string         `json:"nextToken,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEventsResult.html]
type GetLogEventsResult struct {
	Events            []OutputLogEvent `json:"events,omitempty"`
	NextBackwardToken string           `json:"nextBackwardToken,omitempty"`
	NextForwardToken  string           `json:"nextForwardToken,omitempty"`
}

// A log event is a record of some activity that was recorded by the application or resource being monitored.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_InputLogEvent.html]
type InputLogEvent struct {
	Message   string             `json:"Message"`
	Timestamp *datetime.JsonDate `json:"timestamp"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html]
type LogGroup struct {
	Arn               string             `json:"arn,omitempty"`
	CreationTime      *datetime.JsonDate `json:"creationTime,omitempty"`
	LogGroupName      string             `json:"logGroupName,omitempty"`
	MetricFilterCount int                `json:"metricFilterCount,omitempty"`
	RetentionInDays   int                `json:"retentionInDays,omitempty"`
	StoredBytes       int64              `json:"storedBytes,omitempty"`
}

// A log stream is sequence of log events that share the same emitter.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogStream.html]
type LogStream struct {
	Arn                 string             `json:"arn,omitempty"`
	CreationTime        *datetime.JsonDate `json:"creationTime,omitempty"`
	FirstEventTimestamp *datetime.JsonDate `json:"firstEventTimestamp,omitempty"`
	LastEventTimestamp  *datetime.JsonDate `json:"lastEventTimestamp,omitempty"`
	LastIngestionTime   *datetime.JsonDate `json:"lastIngestionTime,omitempty"`
	LogStreamName       string             `json:"logStreamName,omitempty"`
	StoredBytes         int64              `json:"storedBytes,omitempty"`
	UploadSequenceToken string             `json:"uploadSequenceToken,omitempty"`
}

// Metric filters can be used to express how Amazon CloudWatch Logs would extract metric observations
// from ingested log events and transform them to metric data in a CloudWatch metric.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html]
type MetricFilter struct {
	CreationTime          *datetime.JsonDate     `json:"creationTime,omitempty"`
	FilterName            string                 `json:"filterName,omitempty"`
	FilterPattern         string                 `json:"filterPattern,omitempty"`
	MetricTransformations []MetricTransformation `json:"metricTransformations,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilterMatchRecord.html]
type MetricFilterMatchRecord struct {
	EventMessage    string            `json:"eventMessage,omitempty"`
	EventNumber     int64             `json:"eventNumber,omitempty"`
	ExtractedValues map[string]string `json:"extractedValues,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricTransformation.html]
type MetricTransformation struct {
	MetricName      string `json:"metricName"`
	MetricNamespace string `json:"metricNamespace"`
	MetricValue     string `json:"metricValue"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_OutputLogEvent.html]
type OutputLogEvent struct {
	IngestionTime *datetime.JsonDate `json:"ingestionTime,omitempty"`
	Message       string             `json:"message,omitempty"`
	Timestamp     *datetime.JsonDate `json:"timestamp,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEventsResult.html]
type PutLogEventsResult struct {
	NextSequenceToken string `json:"nextSequenceToken,omitempty"`
}

// No action documentation available.
// [http://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TestMetricFilterResult.html]
type TestMetricFilterResult struct {
	Matches []MetricFilterMatchRecord `json:"matches,omitempty"`
}
