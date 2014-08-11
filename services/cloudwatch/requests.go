package cloudwatch

import (
	"time"
)

/*****************************************************************************/

// Deletes all specified alarms. In the event of an error, no alarms are deleted.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DeleteAlarms.html]
type DeleteAlarmsRequest struct {
	AlarmNames []string `name:"AlarmName.member.#" base:"1"`
}

// Creates a new DeleteAlarmsRequest.
func NewDeleteAlarmsRequest(alarmNames ...string) *DeleteAlarmsRequest {
	return &DeleteAlarmsRequest{alarmNames}
}

/*****************************************************************************/

// Retrieves history for the specified alarm. Filter alarms by date range or item type.
// If an alarm name is not specified, Amazon CloudWatch returns histories for all of
// the owner's alarms.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html]
type DescribeAlarmHistoryRequest struct {
	AlarmName       string          `name:"AlarmName,omitempty"`
	EndDate         time.Time       `name:"EndDate,omitempty" format:"2006-01-02T15:04:05Z"`
	HistoryItemType HistoryItemType `name:"HistoryItemType,omitempty"`
	MaxRecords      int             `name:"MaxRecords,omitempty"`
	NextToken       string          `name:"NextToken,omitempty"`
	StartDate       time.Time       `name:"StartDate,omitempty" format:"2006-01-02T15:04:05Z"`
}

// Creates a new DescribeAlarmHistoryRequest.
func NewDescribeAlarmHistoryRequest(alarmName string, historyItemType HistoryItemType, startDate time.Time) *DescribeAlarmHistoryRequest {
	return &DescribeAlarmHistoryRequest{
		AlarmName:       alarmName,
		StartDate:       startDate,
		HistoryItemType: historyItemType,
	}
}

/*****************************************************************************/

// Retrieves alarms with the specified names. If no name is specified, all alarms
// for the user are returned. Alarms can be retrieved by using only a prefix for
// the alarm name, the alarm state, or a prefix for any action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html]
type DescribeAlarmsRequest struct {
	ActionPrefix    string     `name:"ActionPrefix,omitempty"`
	AlarmNamePrefix string     `name:"AlarmNamePrefix,omitempty"`
	AlarmNames      []string   `name:"AlarmName.member.#,omitempty" base:"1"`
	MaxRecords      int        `name:"MaxRecords,omitempty"`
	NextToken       string     `name:"NextToken,omitempty"`
	StateValue      StateValue `name:"StateValue,omitempty"`
}

// Creates a new DescribeAlarmsRequest.
func NewDescribeAlarmsRequest() *DescribeAlarmsRequest {
	return &DescribeAlarmsRequest{}
}

// Add alarm names.
func (r *DescribeAlarmsRequest) AddAlarmNames(alarmNames ...string) {
	r.AlarmNames = append(r.AlarmNames, alarmNames...)
}

/*****************************************************************************/

// Retrieves all alarms for a single metric. Specify a statistic, period, or unit to filter the set of alarms further.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmsForMetric.html]
type DescribeAlarmsForMetricRequest struct {
	Dimensions []Dimension `name:"Dimension.member.#.,omitempty" base:"1"`
	MetricName string      `name:"MetricName"`
	Namespace  string      `name:"Namespace"`
	Period     int         `name:"Period,omitempty"`
	Statistic  Statistic   `name:"Statistic,omitempty"`
	Unit       Unit        `name:"Unit,omitempty"`
}

// Creates a new DescribeAlarmsForMetricRequest.
func NewDescribeAlarmsForMetricRequest(namespace, metricName string) *DescribeAlarmsForMetricRequest {
	return &DescribeAlarmsForMetricRequest{Namespace: namespace, MetricName: metricName}
}

/*****************************************************************************/

// Disables actions for the specified alarms. When an alarm's actions are disabled
// the alarm's state may change, but none of the alarm's actions will execute.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DisableAlarmActions.html]
type DisableAlarmActionsRequest struct {
	AlarmNames []string `name:"AlarmName.#,omitempty" base:"1"`
}

// Creates a new DisableAlarmActionsRequest.
func NewDisableAlarmActionsRequest(alarmNames ...string) *DisableAlarmActionsRequest {
	return &DisableAlarmActionsRequest{alarmNames}
}

/*****************************************************************************/

// Enables actions for the specified alarms.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_EnableAlarmActions.html]
type EnableAlarmActionsRequest struct {
	AlarmNames []string `name:"AlarmName.#,omitempty" base:"1"`
}

// Creates a new EnableAlarmActionsRequest.
func NewEnableAlarmActionsRequest(alarmNames ...string) *EnableAlarmActionsRequest {
	return &EnableAlarmActionsRequest{alarmNames}
}

/*****************************************************************************/

// Gets statistics for the specified metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html]
type GetMetricStatisticsRequest struct {
	Dimensions []Dimension `name:"Dimension.member.#.,omitempty" base:"1"`
	EndTime    time.Time   `name:"EndTime" format:"2006-01-02T15:04:05Z"`
	MetricName string      `name:"MetricName"`
	Namespace  string      `name:"Namespace"`
	Period     int         `name:"Period"`
	StartTime  time.Time   `name:"StartTime" format:"2006-01-02T15:04:05Z"`
	Statistics []Statistic `name:"Statistic.member.#"`
	Unit       Unit        `name:"Unit,omitempty"`
}

// Creates a new GetMetricStatisticsRequest.
func NewGetMetricStatisticsRequest(namespace, metricName string, period int, startTime, endTime time.Time, statistics ...Statistic) *GetMetricStatisticsRequest {
	return &GetMetricStatisticsRequest{
		EndTime:    endTime,
		MetricName: metricName,
		Namespace:  namespace,
		Period:     period,
		StartTime:  startTime,
		Statistics: statistics,
	}
}

/*****************************************************************************/

// Returns a list of valid metrics stored for the AWS account owner. Returned metrics
// can be used with GetMetricStatistics to obtain statistical data for a given metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html]
type ListMetricsRequest struct {
	Dimensions []Dimension `name:"Dimension.member.#.,omitempty" base:"1"`
	MetricName string      `name:"MetricName,omitempty"`
	Namespace  string      `name:"Namespace,omitempty"`
	NextToken  string      `name:"NextToken,omitempty"`
}

// Creates a new ListMetricsRequest.
func NewListMetricsRequest(namespace, metricName string) *ListMetricsRequest {
	return &ListMetricsRequest{Namespace: namespace, MetricName: metricName}
}

/*****************************************************************************/

// Creates or updates an alarm and associates it with the specified Amazon CloudWatch metric.
// Optionally, this operation can associate one or more Amazon Simple Notification Service
// resources with the alarm.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html]
type PutMetricAlarmRequest struct {
	ActionsEnabled          bool               `name:"ActionsEnabled,omitempty"`
	AlarmActions            []string           `name:"AlarmAction.member.#,omitempty" base:"1"`
	AlarmDescription        string             `name:"AlarmDescription,omitempty"`
	AlarmName               string             `name:"AlarmName"`
	ComparisonOperator      ComparisonOperator `name:"ComparisonOperator"`
	Dimensions              []Dimension        `name:"Dimension.member.#.,omitempty" base:"1"`
	EvaluationPeriods       int                `name:"EvaluationPeriods"`
	InsufficientDataActions []string           `name:"InsufficientDataAction.member.#,omitempty" base:"1"`
	MetricName              string             `name:"MetricName"`
	Namespace               string             `name:"Namespace"`
	OKActions               []string           `name:"OKAction.member.#,omitempty" base:"1"`
	Period                  int                `name:"Period"`
	Statistic               Statistic          `name:"Statistic"`
	Threshold               float64            `name:"Threshold"`
	Unit                    Unit               `name:"Unit,omitempty"`
}

// Creates a new PutMetricAlarmRequest.
func NewPutMetricAlarmRequest(alarmName, namespace, metricName string, period, evaluationPeriods int, statistic Statistic, threshold float64, op ComparisonOperator) *PutMetricAlarmRequest {
	return &PutMetricAlarmRequest{
		AlarmName:          alarmName,
		ComparisonOperator: op,
		EvaluationPeriods:  evaluationPeriods,
		MetricName:         metricName,
		Namespace:          namespace,
		Period:             period,
		Statistic:          statistic,
		Threshold:          threshold,
	}
}

/*****************************************************************************/

// Publishes metric data points to Amazon CloudWatch. Amazon CloudWatch associates the
// data points with the specified metric. If the specified metric does not exist,
// Amazon CloudWatch creates the metric. When Amazon CloudWatch creates a metric, it can
// take up to fifteen minutes for the metric to appear in calls to the ListMetrics action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricData.html]
type PutMetricDataRequest struct {
	MetricData []MetricDatum `name:"MetricData.member.#." base:"1"`
	Namespace  string        `name:"Namespace"`
}

// Creates a new PutMetricDataRequest.
func NewPutMetricDataRequest(namespace string, metricData ...MetricDatum) *PutMetricDataRequest {
	return &PutMetricDataRequest{metricData, namespace}
}

/*****************************************************************************/

// Temporarily sets the state of an alarm. When the updated StateValue differs from the
// previous value, the action configured for the appropriate state is invoked. This is not
// a permanent change. The next periodic alarm check (in about a minute) will set the
// alarm to its actual state.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_SetAlarmState.html]
type SetAlarmStateRequest struct {
	AlarmName       string     `name:"AlarmName"`
	StateReason     string     `name:"StateReason"`
	StateReasonData string     `name:"StateReasonData,omitempty"`
	StateValue      StateValue `name:"StateValue"`
}

// Creates a new SetAlarmStateRequest.
func NewSetAlarmStateRequest(alarmName, stateReason string, stateValue StateValue) *SetAlarmStateRequest {
	return &SetAlarmStateRequest{AlarmName: alarmName, StateReason: stateReason, StateValue: stateValue}
}
