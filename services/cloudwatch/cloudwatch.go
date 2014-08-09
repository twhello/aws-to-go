// 
// Amazon CloudWatch is a web service that enables you to publish, monitor, and manage various metrics,
// as well as configure alarm actions based on data from metrics.
//
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/Welcome.html]
//
package cloudwatch

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"github.com/twhello/aws-to-go/util/netutil"
	"net/http"
	"time"
)

const ServiceName = "monitoring"

// The CloudWatch Service object. Use cloudwatch.NewService().
type CloudWatchService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *CloudWatchService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *CloudWatchService) RegionName() string {
	return s.region.Name()
}

func (s *CloudWatchService) Endpoint() string {
	return s.endpoint
}

// Low-level request to CloudWatch service.
func (s *CloudWatchService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *CloudWatchService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	qs := netutil.MarshalValues(request)
	qs.Add("AWSAccessKeyId", s.cred.AccessKeyId())
	qs.Add("Timestamp", time.Now().UTC().Format(time.RFC3339)) // ISO 8601 2006-01-02T15:04:05.999Z
	qs.Add("Version", "2014-03-28")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", s.Endpoint(), qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * CloudWatch Service Methods.
 */

// Deletes all specified alarms. In the event of an error, no alarms are deleted.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DeleteAlarms.html]
func (s *CloudWatchService) DeleteAlarms(req *DeleteAlarmsRequest) (result *DeleteAlarmsResponse, err error) {

	result = new(DeleteAlarmsResponse)
	err = s.wrapperSignAndDo("DeleteAlarms", req, result)
	return
}

// Retrieves history for the specified alarm. Filter alarms by date range or item type.
// If an alarm name is not specified, Amazon CloudWatch returns histories for all of
// the owner's alarms.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html]
func (s *CloudWatchService) DescribeAlarmHistory(req *DescribeAlarmHistoryRequest) (result *DescribeAlarmHistoryResponse, err error) {

	result = new(DescribeAlarmHistoryResponse)
	err = s.wrapperSignAndDo("DescribeAlarmHistory", req, result)
	return
}

// Retrieves alarms with the specified names. If no name is specified, all alarms
// for the user are returned. Alarms can be retrieved by using only a prefix for
// the alarm name, the alarm state, or a prefix for any action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html]
func (s *CloudWatchService) DescribeAlarms(req *DescribeAlarmsRequest) (result *DescribeAlarmsResponse, err error) {

	result = new(DescribeAlarmsResponse)
	err = s.wrapperSignAndDo("DescribeAlarms", req, result)
	return
}

// Retrieves all alarms for a single metric. Specify a statistic, period, or unit to filter the set of alarms further.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmsForMetric.html]
func (s *CloudWatchService) DescribeAlarmsForMetric(req *DescribeAlarmsForMetricRequest) (result *DescribeAlarmsForMetricResponse, err error) {

	result = new(DescribeAlarmsForMetricResponse)
	err = s.wrapperSignAndDo("DescribeAlarmsForMetric", req, result)
	return
}

// Disables actions for the specified alarms. When an alarm's actions are disabled
// the alarm's state may change, but none of the alarm's actions will execute.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DisableAlarmActions.html]
func (s *CloudWatchService) DisableAlarmActions(req *DisableAlarmActionsRequest) (result *DisableAlarmActionsResponse, err error) {

	result = new(DisableAlarmActionsResponse)
	err = s.wrapperSignAndDo("DisableAlarmActions", req, result)
	return
}

// Enables actions for the specified alarms.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_EnableAlarmActions.html]
func (s *CloudWatchService) EnableAlarmActions(req *EnableAlarmActionsRequest) (result *EnableAlarmActionsResponse, err error) {

	result = new(EnableAlarmActionsResponse)
	err = s.wrapperSignAndDo("EnableAlarmActions", req, result)
	return
}

// Gets statistics for the specified metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html]
func (s *CloudWatchService) GetMetricStatistics(req *GetMetricStatisticsRequest) (result *GetMetricStatisticsResponse, err error) {

	result = new(GetMetricStatisticsResponse)
	err = s.wrapperSignAndDo("GetMetricStatistics", req, result)
	return
}

// Returns a list of valid metrics stored for the AWS account owner. Returned metrics
// can be used with GetMetricStatistics to obtain statistical data for a given metric.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html]
func (s *CloudWatchService) ListMetrics(req *ListMetricsRequest) (result *ListMetricsResponse, err error) {

	result = new(ListMetricsResponse)
	err = s.wrapperSignAndDo("ListMetrics", req, result)
	return
}

// Creates or updates an alarm and associates it with the specified Amazon CloudWatch metric.
// Optionally, this operation can associate one or more Amazon Simple Notification Service
// resources with the alarm.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html]
func (s *CloudWatchService) PutMetricAlarm(req *PutMetricAlarmRequest) (result *PutMetricAlarmResponse, err error) {

	result = new(PutMetricAlarmResponse)
	err = s.wrapperSignAndDo("PutMetricAlarm", req, result)
	return
}

// Publishes metric data points to Amazon CloudWatch. Amazon CloudWatch associates the
// data points with the specified metric. If the specified metric does not exist,
// Amazon CloudWatch creates the metric. When Amazon CloudWatch creates a metric, it can
// take up to fifteen minutes for the metric to appear in calls to the ListMetrics action.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricData.html]
func (s *CloudWatchService) PutMetricData(req *PutMetricDataRequest) (result *PutMetricDataResponse, err error) {

	result = new(PutMetricDataResponse)
	err = s.wrapperSignAndDo("PutMetricData", req, result)
	return
}

// Temporarily sets the state of an alarm. When the updated StateValue differs from the
// previous value, the action configured for the appropriate state is invoked. This is not
// a permanent change. The next periodic alarm check (in about a minute) will set the
// alarm to its actual state.
// [http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_SetAlarmState.html]
func (s *CloudWatchService) SetAlarmState(req *SetAlarmStateRequest) (result *SetAlarmStateResponse, err error) {

	result = new(SetAlarmStateResponse)
	err = s.wrapperSignAndDo("SetAlarmState", req, result)
	return
}

// Creates a new SNS Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *CloudWatchService {

	subdomain := ServiceName
	if region.Name() != regions.DEFAULT_REGION {
		subdomain += "." + region.Name()
	}
	return &CloudWatchService{cred, region, "https://" + subdomain + ".amazonaws.com/doc/2010-08-01"}
}
