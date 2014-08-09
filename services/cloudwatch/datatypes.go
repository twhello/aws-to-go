package cloudwatch

import (
	"time"
)

/*****************************************************************************
 * Constants
 */

// The arithmetic operation to use when comparing the specified Statistic and Threshold.
// The specified Statistic value is used as the first operand.
//	GTE_THRESHOLD = "GreaterThanOrEqualToThreshold"
//	GT_THRESHOLD = "GreaterThanThreshold"
//	LTE_THRESHOLD = "LessThanOrEqualToThreshold"
//	LT_THRESHOLD = "LessThanThreshold"
type ComparisonOperator string

const (
	GTE_THRESHOLD ComparisonOperator = "GreaterThanOrEqualToThreshold"
	GT_THRESHOLD  ComparisonOperator = "GreaterThanThreshold"
	LTE_THRESHOLD ComparisonOperator = "LessThanOrEqualToThreshold"
	LT_THRESHOLD  ComparisonOperator = "LessThanThreshold"
)

// The type of alarm history item.
//	ACTION = "Action"
//	CONFIGURATION_UPDATE = "ConfigurationUpdate"
//	STATE_UPDATE = "StateUpdate"
type HistoryItemType string

const (
	ACTION               HistoryItemType = "Action"
	CONFIGURATION_UPDATE HistoryItemType = "ConfigurationUpdate"
	STATE_UPDATE         HistoryItemType = "StateUpdate"
)

// The state value for the alarm.
//	ALARM
//	INSUFFICIENT_DATA
//	OK
type StateValue string

const (
	ALARM             StateValue = "ALARM"
	INSUFFICIENT_DATA StateValue = "INSUFFICIENT_DATA"
	OK                StateValue = "OK"
)

// The statistic to apply to the alarm's associated metric.
//	AVERAGE = "Average"
//	MAXIMUM = "Maximum"
//	MINIMUM = "Minimum"
//	SAMPLE_COUNT = "SampleCount"
//	SUM = "Sum"
type Statistic string

const (
	AVERAGE      Statistic = "Average"
	MAXIMUM      Statistic = "Maximum"
	MINIMUM      Statistic = "Minimum"
	SAMPLE_COUNT Statistic = "SampleCount"
	SUM          Statistic = "Sum"
)

// The unit of the alarm's associated metric:
//	SECONDS = "Seconds"
//	MICROSECONDS = "Microseconds"
//	MILLISECONDS = "Milliseconds"
//	BYTES = "Bytes"
//	KILOBYTES = "Kilobytes"
//	MEGABYTES = "Megabytes"
//	GIGABYTES = "Gigabytes"
//	TERABYTES = "Terabytes"
//	BITS = "Bits"
//	KILOBITS = "Kilobits"
//	MEGABITS = "Megabits"
//	GIGABITS = "Gigabits"
//	TERABITS = "Terabits"
//	PERCENT = "Percent"
//	COUNT = "Count"
//	BYTES_SECOND = "Bytes/Second"
//	KILOBYTES_SECOND = "Kilobytes/Second"
//	MEGABYTES_SECOND = "Megabytes/Second"
//	GIGABYTES_SECOND = "Gigabytes/Second"
//	TERABYTES_SECOND = "Terabytes/Second"
//	BITS_SECOND = "Bits/Second"
//	KILOBITS_SECOND = "Kilobits/Second"
//	MEGABITS_SECOND = "Megabits/Second"
//	GIGABITS_SECOND = "Gigabits/Second"
//	TERABITS_SECOND = "Terabits/Second"
//	COUNT_SECOND = "Count/Second"
//	NONE = "None"
type Unit string

const (
	SECONDS          Unit = "Seconds"
	MICROSECONDS     Unit = "Microseconds"
	MILLISECONDS     Unit = "Milliseconds"
	BYTES            Unit = "Bytes"
	KILOBYTES        Unit = "Kilobytes"
	MEGABYTES        Unit = "Megabytes"
	GIGABYTES        Unit = "Gigabytes"
	TERABYTES        Unit = "Terabytes"
	BITS             Unit = "Bits"
	KILOBITS         Unit = "Kilobits"
	MEGABITS         Unit = "Megabits"
	GIGABITS         Unit = "Gigabits"
	TERABITS         Unit = "Terabits"
	PERCENT          Unit = "Percent"
	COUNT            Unit = "Count"
	BYTES_SECOND     Unit = "Bytes/Second"
	KILOBYTES_SECOND Unit = "Kilobytes/Second"
	MEGABYTES_SECOND Unit = "Megabytes/Second"
	GIGABYTES_SECOND Unit = "Gigabytes/Second"
	TERABYTES_SECOND Unit = "Terabytes/Second"
	BITS_SECOND      Unit = "Bits/Second"
	KILOBITS_SECOND  Unit = "Kilobits/Second"
	MEGABITS_SECOND  Unit = "Megabits/Second"
	GIGABITS_SECOND  Unit = "Gigabits/Second"
	TERABITS_SECOND  Unit = "Terabits/Second"
	COUNT_SECOND     Unit = "Count/Second"
	NONE             Unit = "None"
)

/*****************************************************************************
 * Data Types
 */

// The AlarmHistoryItem data type contains descriptive information about the history of a specific alarm.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_AlarmHistoryItem.html]
type AlarmHistoryItem struct {
	AlarmName       string          `xml:"AlarmName,omitempty"`
	HistoryData     string          `xml:"HistoryData,omitempty"`
	HistoryItemType HistoryItemType `xml:"HistoryItemType,omitempty"`
	HistorySummary  string          `xml:"HistorySummary,omitempty"`
	Timestamp       time.Time       `xml:"Timestamp,omitempty"`
}

// The Datapoint data type encapsulates the statistical data that Amazon CloudWatch computes from metric data.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Datapoint.html]
type Datapoint struct {
	Average     float64   `xml:"Average,omitempty"`
	Maximum     float64   `xml:"Maximum,omitempty"`
	Minimum     float64   `xml:"Minimum,omitempty"`
	SampleCount float64   `xml:"SampleCount,omitempty"`
	Sum         float64   `xml:"Sum,omitempty"`
	Timestamp   time.Time `xml:"Timestamp,omitempty"`
	Unit        Unit      `xml:"Unit,omitempty"`
}

// The output for the DescribeAlarmHistory action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistoryResult.html]
type DescribeAlarmHistoryResult struct {
	AlarmHistoryItems []AlarmHistoryItem `xml:"AlarmHistoryItems,omitempty"`
	NextToken         string             `xml:"NextToken,omitempty"`
}

// The output for the DescribeAlarmsForMetric action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmsForMetricResult.html]
type DescribeAlarmsForMetricResult struct {
	MetricAlarms []MetricAlarm `xml:"MetricAlarms,omitempty"`
}

// The output for the DescribeAlarms action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmsResult.html]
type DescribeAlarmsResult struct {
	MetricAlarms []MetricAlarm `xml:"MetricAlarms,omitempty"`
	NextToken    string        `xml:"NextToken,omitempty"`
}

// The Dimension data type further expands on the identity of a metric using a Name, Value pair.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Dimension.html]
type Dimension struct {
	Name  string `xml:"Name" name:"Name"`
	Value string `xml:"Value" name:"Value"`
}

// The DimensionFilter data type is used to filter ListMetrics results.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DimensionFilter.html]
type DimensionFilter struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value,omitempty"`
}

// The output for the GetMetricStatistics action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatisticsResult.html]
type GetMetricStatisticsResult struct {
	Datapoints []Datapoint `xml:"Datapoints,omitempty"`
	Label      string      `xml:"Label,omitempty"`
}

// The output for the ListMetrics action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetricsResult.html]
type ListMetricsResult struct {
	Metrics   []Metric `xml:"Metrics,omitempty"`
	NextToken string   `xml:"NextToken,omitempty"`
}

// The Metric data type contains information about a specific metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html]
type Metric struct {
	Dimensions []Dimension `xml:"Dimensions,omitempty"`
	MetricName string      `xml:"MetricName,omitempty"`
	Namespace  string      `xml:"Namespace,omitempty"`
}

// The MetricAlarm data type represents an alarm.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html]
type MetricAlarm struct {
	ActionsEnabled                     bool               `xml:"ActionsEnabled,omitempty"`
	AlarmActions                       string             `xml:"AlarmActions,omitempty"`
	AlarmArn                           string             `xml:"AlarmArn,omitempty"`
	AlarmConfigurationUpdatedTimestamp time.Time          `xml:"AlarmConfigurationUpdatedTimestamp,omitempty"`
	AlarmDescription                   string             `xml:"AlarmDescription,omitempty"`
	AlarmName                          string             `xml:"AlarmName,omitempty"`
	ComparisonOperator                 ComparisonOperator `xml:"ComparisonOperator,omitempty"`
	Dimensions                         []Dimension        `xml:"Dimensions,omitempty"`
	EvaluationPeriods                  int                `xml:"EvaluationPeriods,omitempty"`
	InsufficientDataActions            []string           `xml:"InsufficientDataActions,omitempty"`
	MetricName                         string             `xml:"MetricName,omitempty"`
	Namespace                          string             `xml:"Namespace,omitempty"`
	OKActions                          []string           `xml:"OKActions,omitempty"`
	Period                             int                `xml:"Period,omitempty"`
	StateReason                        string             `xml:"StateReason,omitempty"`
	StateReasonData                    string             `xml:"StateReasonData,omitempty"`
	StateUpdatedTimestamp              time.Time          `xml:"StateUpdatedTimestamp,omitempty"`
	StateValue                         StateValue         `xml:"StateValue,omitempty"`
	Statistic                          Statistic          `xml:"Statistic,omitempty"`
	Threshold                          float64            `xml:"Threshold,omitempty"`
	Unit                               Unit               `xml:"Unit,omitempty"`
}

// The MetricDatum data type encapsulates the information sent with PutMetricData to either
// create a new metric or add new values to be aggregated into an existing metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html]
type MetricDatum struct {
	Dimensions      []Dimension  `xml:"Dimensions,omitempty" name:"Dimensions,omitempty"`
	MetricName      string       `xml:"MetricName" name:"MetricName"`
	StatisticValues StatisticSet `xml:"StatisticValues,omitempty" name:"StatisticValues,omitempty"`
	Timestamp       time.Time    `xml:"Timestamp,omitempty" name:"Timestamp,omitempty" format:"2006-01-02T15:04:05Z"`
	Unit            Unit         `xml:"Unit,omitempty" name:"Unit,omitempty"`
}

type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

// The StatisticSet data type describes the StatisticValues component of MetricDatum,
// and represents a set of statistics that describes a specific metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_StatisticSet.html]
type StatisticSet struct {
	Maximum     float64 `xml:"Maximum"`
	Minimum     float64 `xml:"Minimum"`
	SampleCount float64 `xml:"SampleCount"`
	Sum         float64 `xml:"Sum"`
}

/*****************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DeleteAlarms.html]
type DeleteAlarmsResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html]
type DescribeAlarmHistoryResponse struct {
	DescribeAlarmHistoryResult DescribeAlarmHistoryResult `xml:"DescribeAlarmHistoryResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html]
type DescribeAlarmsResponse struct {
	DescribeAlarmsResult DescribeAlarmsResult `xml:"DescribeAlarmsResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmsForMetric.html]
type DescribeAlarmsForMetricResponse struct {
	DescribeAlarmsForMetricResult DescribeAlarmsForMetricResult `xml:"DescribeAlarmsForMetricResult"`
	ResponseMetadata              ResponseMetadata              `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DisableAlarmActions.html]
type DisableAlarmActionsResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_EnableAlarmActions.html]
type EnableAlarmActionsResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html]
type GetMetricStatisticsResponse struct {
	GetMetricStatisticsResult GetMetricStatisticsResult `xml:"GetMetricStatisticsResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html]
type ListMetricsResponse struct {
	ListMetricsResult ListMetricsResult `xml:"ListMetricsResult"`
	ResponseMetadata  ResponseMetadata  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html]
type PutMetricAlarmResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricData.html]
type PutMetricDataResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_SetAlarmState.html]
type SetAlarmStateResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}
