package autoscaling

import (
	"time"
)

/******************************************************************************
 * Constants
 */

// Specifies whether the PutScalingPolicy ScalingAdjustment parameter is an absolute
// number or a percentage of the current capacity.
//	CHANGE_IN_CAPACITY = "ChangeInCapacity"
//	EXACT_CAPACITY = "ExactCapacity"
//	PERCENT_CHANGE_IN_CAPACITY = "PercentChangeInCapacity"
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AdjustmentType.html]
type AdjustmentType string

const (
	CHANGE_IN_CAPACITY         AdjustmentType = "ChangeInCapacity"
	EXACT_CAPACITY             AdjustmentType = "ExactCapacity"
	PERCENT_CHANGE_IN_CAPACITY AdjustmentType = "PercentChangeInCapacity"
)

// Notification types supported by Auto Scaling.
//	EC2_INSTANCE_LAUNCH AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_LAUNCH"
//	EC2_INSTANCE_LAUNCH_ERROR AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_LAUNCH_ERROR"
//	EC2_INSTANCE_TERMINATE AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_TERMINATE"
//	EC2_INSTANCE_TERMINATE_ERROR AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_TERMINATE_ERROR"
//	TEST_NOTIFICATION  AutoScalingNotificationType = "autoscaling:TEST_NOTIFICATION"
type AutoScalingNotificationType string

const (
	EC2_INSTANCE_LAUNCH          AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_LAUNCH"
	EC2_INSTANCE_LAUNCH_ERROR    AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_LAUNCH_ERROR"
	EC2_INSTANCE_TERMINATE       AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_TERMINATE"
	EC2_INSTANCE_TERMINATE_ERROR AutoScalingNotificationType = "autoscaling:EC2_INSTANCE_TERMINATE_ERROR"
	TEST_NOTIFICATION            AutoScalingNotificationType = "autoscaling:TEST_NOTIFICATION"
)

// The name of the filter.
//	AUTO_SCALING_GROUP = "auto-scaling-group"
//	KEY = "key"
//	VALUE = "value"
//	PROPOGATE_AT_LAUNCH = "propagate-at-launch"
type FilterName string

const (
	AUTO_SCALING_GROUP  FilterName = "auto-scaling-group"
	KEY                 FilterName = "key"
	VALUE               FilterName = "value"
	PROPOGATE_AT_LAUNCH FilterName = "propagate-at-launch"
)

// The service of interest for the health status check.
//	EC2
//	ELB
type HealthCheckType string

const (
	EC2 HealthCheckType = "EC2"
	ELB HealthCheckType = "ELB"
)

// The health status of this instance.
//	HEALTHY = "Healthy"
//	UNHEALTHY = "Unhealthy"
type HealthStatus string

const (
	HEALTHY   HealthStatus = "Healthy"
	UNHEALTHY HealthStatus = "Unhealthy"
)

// The Amazon EC2 instances within your Auto Scaling group progresses
// through the following states over their lifespan.
//	IN_SERVICE - When the instance is live and running.
//	PENDING - When the instance is in the process of launching.
//	QUARANTINED - Not currently used.
//	TERMINATED - When the instance is no longer in service.
//	TERMINATING - When the instance is in the process of being terminated.
type LifecycleState string

const (
	IN_SERVICE  LifecycleState = "InService"
	PENDING     LifecycleState = "Pending"
	QUARANTINED LifecycleState = "Quarantined"
	TERMINATED  LifecycleState = "Terminating"
	TERMINATING LifecycleState = "Terminating"
)

// Supported Metric Collection Types.
//	GROUP_MIN_SIZE MetricCollectionType = "GroupMinSize"
//	GROUP_MAX_SIZE MetricCollectionType = "GroupMaxSize"
//	GROUP_DESCRIPTION_CAPACITY MetricCollectionType = "GroupDesiredCapacity"
//	GROUP_IN_SERVICE_INSTANCES MetricCollectionType = "GroupInServiceInstances"
//	GROUP_PENDING_INSTANCES MetricCollectionType = "GroupPendingInstances"
//	GROUP_TERMINATING_INSTANCES MetricCollectionType = "GroupTerminatingInstances"
//	GROUP_TOTAL_INSTANCES MetricCollectionType = "GroupTotalInstances"
type MetricCollectionType string

const (
	GROUP_MIN_SIZE              MetricCollectionType = "GroupMinSize"
	GROUP_MAX_SIZE              MetricCollectionType = "GroupMaxSize"
	GROUP_DESCRIPTION_CAPACITY  MetricCollectionType = "GroupDesiredCapacity"
	GROUP_IN_SERVICE_INSTANCES  MetricCollectionType = "GroupInServiceInstances"
	GROUP_PENDING_INSTANCES     MetricCollectionType = "GroupPendingInstances"
	GROUP_TERMINATING_INSTANCES MetricCollectionType = "GroupTerminatingInstances"
	GROUP_TOTAL_INSTANCES       MetricCollectionType = "GroupTotalInstances"
)

// The granularity of a Metric.
type MetricGranularityType string

// The tenancy of the instance.
//	DEFAULT = "default"
//	DEDICATED = "dedicated"
type PlacementTenancy string

const (
	DEFAULT   PlacementTenancy = "default"
	DEDICATED PlacementTenancy = "dedicated"
)

// Auto Scaling process types.
//	ADD_TO_LOAD_BALANCER ProcessType = "AddToLoadBalancer"
//	ALARM_NOTIFICATION ProcessType = "AlarmNotification"
//	AZ_REBALANCE ProcessType = "AZRebalance"
//	HEALTH_CHECK  ProcessType = "HealthCheck"
//	LAUNCH ProcessType = "Launch"
//	REPLACE_UNHEALTHY ProcessType = "ReplaceUnhealthy"
//	SCHEDULE_ACTIONS ProcessType = "ScheduledActions"
//	TERMINATE ProcessType = "Terminate"
type ProcessType string

const (
	ADD_TO_LOAD_BALANCER ProcessType = "AddToLoadBalancer"
	ALARM_NOTIFICATION   ProcessType = "AlarmNotification"
	AZ_REBALANCE         ProcessType = "AZRebalance"
	HEALTH_CHECK         ProcessType = "HealthCheck"
	LAUNCH               ProcessType = "Launch"
	REPLACE_UNHEALTHY    ProcessType = "ReplaceUnhealthy"
	SCHEDULE_ACTIONS     ProcessType = "ScheduledActions"
	TERMINATE            ProcessType = "Terminate"
)

// Contains the current status of the activity.
//	CANCELLED = "Cancelled"
//	FAILED = "Failed"
//	IN_PROGRESS = "InProgress"
//	PRE_IN_SERVICE = "PreInService"
//	SUCCESSFUL = "Successful"
//	WAITING_FOR_INSTANCE_ID = "WaitingForInstanceId"
//	WAITING_FOR_SPOT_INSTANCE_ID = "WaitingForSpotInstanceId"
//	WAITING_FOR_SPOT_INSTANCE_REQUEST_ID = "WaitingForSpotInstanceRequestId"
type StatusCode string

const (
	CANCELLED                            StatusCode = "Cancelled"
	FAILED                               StatusCode = "Failed"
	IN_PROGRESS                          StatusCode = "InProgress"
	PRE_IN_SERVICE                       StatusCode = "PreInService"
	SUCCESSFUL                           StatusCode = "Successful"
	WAITING_FOR_INSTANCE_ID              StatusCode = "WaitingForInstanceId"
	WAITING_FOR_SPOT_INSTANCE_ID         StatusCode = "WaitingForSpotInstanceId"
	WAITING_FOR_SPOT_INSTANCE_REQUEST_ID StatusCode = "WaitingForSpotInstanceRequestId"
)

// Termination policies supported by Auto Scaling.
//	CLOSEST_TO_NEXT_INSTANCE_HOUR = "ClosestToNextInstanceHour"
//	DEFAULT = "Default"
//	NEWEST_INSTANCE = "NewestInstance"
//	OLDEST_INSTANCE = "OldestInstance"
//	OLDEST_LAUNCH_CONFIGURATION = "OldestLaunchConfiguration"
type TerminationPolicyType string

const (
	CLOSEST_TO_NEXT_INSTANCE_HOUR TerminationPolicyType = "ClosestToNextInstanceHour"
	DEFAULT_POLICY                TerminationPolicyType = "Default"
	NEWEST_INSTANCE               TerminationPolicyType = "NewestInstance"
	OLDEST_INSTANCE               TerminationPolicyType = "OldestInstance"
	OLDEST_LAUNCH_CONFIGURATION   TerminationPolicyType = "OldestLaunchConfiguration"
)

// The volume type.
//	IO1 = "io1"
//	STANDARD = "standard"
type VolumeType string

const (
	IO1      VolumeType = "io1"
	STANDARD VolumeType = "standard"
)

/******************************************************************************
 * Data Types
 */

// A scaling Activity is a long-running process that represents a change to your AutoScalingGroup,
// such as changing the size of the group. It can also be a process to replace an instance, or a
// process to perform any other long-running operations supported by the API.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Activity.html]
type Activity struct {
	ActivityId           string     `xml:"ActivityId"`
	AutoScalingGroupName string     `xml:"AutoScalingGroupName"`
	Cause                string     `xml:"Cause"`
	Description          string     `xml:"Description,omitempty"`
	Details              string     `xml:"Details,omitempty"`
	EndTime              time.Time  `xml:"EndTime,omitempty"`
	Progress             int        `xml:"Progress,omitempty"`
	StartTime            time.Time  `xml:"StartTime"`
	StatusCode           StatusCode `xml:"StatusCode"`
	StatusMessage        string     `xml:"StatusMessage,omitempty"`
}

// The Alarm data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Alarm.html]
type Alarm struct {
	AlarmARN  string `xml:"AlarmARN"`
	AlarmName string `xml:"AlarmName"`
}

// The AutoScalingGroup data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AutoScalingGroup.html]
type AutoScalingGroup struct {
	AutoScalingGroupARN     string             `xml:"AutoScalingGroupARN,omitempty"`
	AutoScalingGroupName    string             `xml:"AutoScalingGroupName"`
	AvailabilityZones       []string           `xml:"AvailabilityZones>member"`
	CreatedTime             time.Time          `xml:"CreatedTime"`
	DefaultCooldown         int                `xml:"DefaultCooldown"`
	DesiredCapacity         int                `xml:"DesiredCapacity"`
	EnabledMetrics          []EnabledMetric    `xml:"EnabledMetrics>member,omitempty"`
	HealthCheckGracePeriod  int                `xml:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType         HealthCheckType    `xml:"HealthCheckType"`
	Instances               []Instance         `xml:"Instances>member,omitempty"`
	LaunchConfigurationName string             `xml:"LaunchConfigurationName"`
	LoadBalancerNames       []string           `xml:"LoadBalancerNames>member,omitempty"`
	MaxSize                 int                `xml:"MaxSize"`
	MinSize                 int                `xml:"MinSize"`
	PlacementGroup          string             `xml:"PlacementGroup,omitempty"`
	Status                  string             `xml:"Status,omitempty"`
	SuspendedProcesses      []SuspendedProcess `xml:"SuspendedProcesses>member,omitempty"`
	Tags                    []TagDescription   `xml:"Tags>member,omitempty"`
	TerminationPolicies     []string           `xml:"TerminationPolicies>member,omitempty"`
	VPCZoneIdentifier       string             `xml:"VPCZoneIdentifier,omitempty"`
}

// The AutoScalingInstanceDetails data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AutoScalingInstanceDetails.html]
type AutoScalingInstanceDetails struct {
	AutoScalingGroupName    string         `xml:"AutoScalingGroupName"`
	AvailabilityZone        string         `xml:"AvailabilityZone"`
	HealthStatus            HealthStatus   `xml:"HealthStatus"`
	InstanceId              string         `xml:"InstanceId"`
	LaunchConfigurationName string         `xml:"LaunchConfigurationName"`
	LifecycleState          LifecycleState `xml:"LifecycleState"`
}

// The BlockDeviceMapping data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_BlockDeviceMapping.html]
type BlockDeviceMapping struct {
	DeviceName  string `xml:"DeviceName" name:"DeviceName"`
	Ebs         Ebs    `xml:"Ebs,omitempty" name:"Ebs,omitempty"`
	NoDevice    bool   `xml:"NoDevice,omitempty" name:"NoDevice,omitempty"`
	VirtualName string `xml:"VirtualName,omitempty" name:"VirtualName,omitempty"`
}

// The output of the DescribeAccountLimitsResult action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAccountLimitsResult.html]
type DescribeAccountLimitsResult struct {
	MaxNumberOfAutoScalingGroups    int `xml:"MaxNumberOfAutoScalingGroups,omitempty"`
	MaxNumberOfLaunchConfigurations int `xml:"MaxNumberOfLaunchConfigurations,omitempty"`
}

// The output of the DescribeAdjustmentTypes action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAdjustmentTypesResult.html]
type DescribeAdjustmentTypesResult struct {
	AdjustmentTypes []AdjustmentType `xml:"AdjustmentTypes>member>AdjustmentType"`
}

// The AutoScalingGroupsType data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingGroupsResult.html]
type DescribeAutoScalingGroupsResult struct {
	AutoScalingGroups []AutoScalingGroup `xml:"AutoScalingGroups>member"`
	NextToken         string             `xml:"NextToken,omitempty"`
}

// The AutoScalingInstancesType data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingInstancesResult.html]
type DescribeAutoScalingInstancesResult struct {
	AutoScalingInstances []AutoScalingInstanceDetails `xml:"AutoScalingInstances>member,omitempty"`
	NextToken            string                       `xml:"NextToken,omitempty"`
}

// The AutoScalingNotificationTypes data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingNotificationTypesResult.html]
type DescribeAutoScalingNotificationTypesResult struct {
	AutoScalingNotificationTypes []AutoScalingNotificationType `xml:"AutoScalingNotificationTypes>member>Type,omitempty"`
}

// The LaunchConfigurationsType data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeLaunchConfigurationsResult.html]
type DescribeLaunchConfigurationsResult struct {
	LaunchConfigurations []LaunchConfiguration `xml:"LaunchConfigurations>member"`
	NextToken            string                `xml:"NextToken,omitempty"`
}

// The output of the DescribeMetricCollectionTypes action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeMetricCollectionTypesResult.html]
type DescribeMetricCollectionTypesResult struct {
	Granularities []MetricGranularityType `xml:"Granularities>member>Granularity,omitempty"`
	Metrics       []MetricCollectionType  `xml:"Metrics>member>Metric,omitempty"`
}

// The output of the DescribeNotificationConfigurations action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurationsResult.html]
type DescribeNotificationConfigurationsResult struct {
	NextToken                  string                      `xml:"NextToken,omitempty"`
	NotificationConfigurations []NotificationConfiguration `xml:"NotificationConfigurations>member"`
}

// The PoliciesType data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribePoliciesResult.html]
type DescribePoliciesResult struct {
	NextToken       string          `xml:"NextToken,omitempty"`
	ScalingPolicies []ScalingPolicy `xml:"ScalingPolicies>member,omitempty"`
}

// The output for the DescribeScalingActivities action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingActivitiesResult.html]
type DescribeScalingActivitiesResult struct {
	Activities []Activity `xml:"Activities>member"`
	NextToken  string     `xml:"NextToken,omitempty"`
}

// The output of the DescribeScalingProcessTypes action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingProcessTypesResult.html]
type DescribeScalingProcessTypesResult struct {
	Processes []ProcessType `xml:"Processes>member>ProcessName,omitempty"`
}

// A scaling action that is scheduled for a future time and date.
// An action can be scheduled up to thirty days in advance.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScheduledActionsResult.html]
type DescribeScheduledActionsResult struct {
	NextToken                   string                       `xml:"NextToken,omitempty"`
	ScheduledUpdateGroupActions []ScheduledUpdateGroupAction `xml:"ScheduledUpdateGroupActions>member,omitempty"`
}

// The output of the DescribeTags action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTagsResult.html]
type DescribeTagsResult struct {
	NextToken string           `xml:"NextToken,omitempty"`
	Tags      []TagDescription `xml:"Tags>member,omitempty"`
}

// The TerminationPolicyTypes data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTerminationPolicyTypesResult.html]
type DescribeTerminationPolicyTypesResult struct {
	TerminationPolicyTypes []TerminationPolicyType `xml:"TerminationPolicyTypes>member"`
}

// The Ebs data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Ebs.html]
type Ebs struct {
	DeleteOnTermination bool       `xml:"DeleteOnTermination,omitempty"`
	Iops                int        `xml:"Iops,omitempty"`
	SnapshotId          string     `xml:"SnapshotId,omitempty"`
	VolumeSize          int        `xml:"VolumeSize,omitempty"`
	VolumeType          VolumeType `xml:"VolumeType,omitempty"`
}

// The EnabledMetric data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_EnabledMetric.html]
type EnabledMetric struct {
	Granularity string `xml:"Granularity,omitempty"`
	Metric      string `xml:"Metric,omitempty"`
}

// The Filter data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html]
type Filter struct {
	Name   FilterName `xml:"Name,omitempty"`
	Values []string   `xml:"Values,omitempty"`
}

// The Instance data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Instance.html]
type Instance struct {
	AvailabilityZone        string         `xml:"AvailabilityZone"`
	HealthStatus            HealthStatus   `xml:"HealthStatus"`
	InstanceId              string         `xml:"InstanceId"`
	LaunchConfigurationName string         `xml:"LaunchConfigurationName"`
	LifecycleState          LifecycleState `xml:"LifecycleState"`
}

// The InstanceMonitoring data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_InstanceMonitoring.html]
type InstanceMonitoring struct {
	Enabled bool `xml:"Enabled,omitempty" name:"Enabled,omitempty"`
}

// The LaunchConfiguration data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_LaunchConfiguration.html]
type LaunchConfiguration struct {
	AssociatePublicIpAddress bool                 `xml:"AssociatePublicIpAddress,omitempty"`
	BlockDeviceMappings      []BlockDeviceMapping `xml:"BlockDeviceMappings>member,omitempty"`
	CreatedTime              time.Time            `xml:"CreatedTime"`
	EbsOptimized             bool                 `xml:"EbsOptimized,omitempty"`
	IamInstanceProfile       string               `xml:"IamInstanceProfile,omitempty"`
	ImageId                  string               `xml:"ImageId"`
	InstanceMonitoring       *InstanceMonitoring  `xml:"InstanceMonitoring,omitempty"`
	InstanceType             string               `xml:"InstanceType"`
	KernelId                 string               `xml:"KernelId,omitempty"`
	KeyName                  string               `xml:"KeyName,omitempty"`
	LaunchConfigurationARN   string               `xml:"LaunchConfigurationARN,omitempty"`
	LaunchConfigurationName  string               `xml:"LaunchConfigurationName"`
	PlacementTenancy         string               `xml:"PlacementTenancy,omitempty"`
	RamdiskId                string               `xml:"RamdiskId,omitempty"`
	SecurityGroups           []string             `xml:"SecurityGroups>member,omitempty"`
	SpotPrice                string               `xml:"SpotPrice,omitempty"`
	UserData                 string               `xml:"UserData,omitempty"`
}

// The NotificationConfiguration data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html]
type NotificationConfiguration struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName,omitempty"`
	NotificationType     string `xml:"NotificationType,omitempty"`
	TopicARN             string `xml:"TopicARN,omitempty"`
}

// The PolicyARNType data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScalingPolicyResult.html]
type PutScalingPolicyResult struct {
	PolicyARN string `xml:"PolicyARN,omitempty"`
}

type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

// The ScalingPolicy data type.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ScalingPolicy.html]
type ScalingPolicy struct {
	AdjustmentType       AdjustmentType `xml:"AdjustmentType,omitempty"`
	Alarms               []Alarm        `xml:"Alarms>member,omitempty"`
	AutoScalingGroupName string         `xml:"AutoScalingGroupName,omitempty"`
	Cooldown             int            `xml:"Cooldown,omitempty"`
	MinAdjustmentStep    int            `xml:"MinAdjustmentStep,omitempty"`
	PolicyARN            string         `xml:"PolicyARN,omitempty"`
	PolicyName           string         `xml:"PolicyName,omitempty"`
	ScalingAdjustment    int            `xml:"ScalingAdjustment,omitempty"`
}

// This data type stores information about a scheduled update to an Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ScheduledUpdateGroupAction.html]
type ScheduledUpdateGroupAction struct {
	AutoScalingGroupName string    `xml:"AutoScalingGroupName,omitempty"`
	DesiredCapacity      int       `xml:"DesiredCapacity,omitempty"`
	EndTime              time.Time `xml:"EndTime,omitempty"`
	MaxSize              int       `xml:"MaxSize,omitempty"`
	MinSize              int       `xml:"MinSize,omitempty"`
	Recurrence           string    `xml:"Recurrence,omitempty"`
	ScheduledActionARN   string    `xml:"ScheduledActionARN,omitempty"`
	ScheduledActionName  string    `xml:"ScheduledActionName,omitempty"`
	StartTime            time.Time `xml:"StartTime,omitempty"`
}

// An Auto Scaling process that has been suspended. For more information, see ProcessType.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SuspendedProcess.html]
type SuspendedProcess struct {
	ProcessName      string `xml:"ProcessName,omitempty"`
	SuspensionReason string `xml:"SuspensionReason,omitempty"`
}

// The tag applied to an Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Tag.html]
type Tag struct {
	Key               string `xml:"Key" name:"Key"`
	PropagateAtLaunch bool   `xml:"PropagateAtLaunch,omitempty" name:"PropagateAtLaunch,omitempty"`
	ResourceId        string `xml:"ResourceId,omitempty" name:"ResourceId,omitempty"`
	ResourceType      string `xml:"ResourceType,omitempty" name:"ResourceType,omitempty"`
	Value             string `xml:"Value,omitempty" name:"Value,omitempty"`
}

// The tag applied to an Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_TagDescription.html]
type TagDescription struct {
	Key               string `xml:"Key,ommitempty"`
	PropagateAtLaunch bool   `xml:"PropagateAtLaunch,omitempty"`
	ResourceId        string `xml:"ResourceId,omitempty"`
	ResourceType      string `xml:"ResourceType,omitempty"`
	Value             string `xml:"Value,omitempty"`
}

// The output for the TerminateInstanceInAutoScalingGroup action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_TerminateInstanceInAutoScalingGroupResult.html]
type TerminateInstanceInAutoScalingGroupResult struct {
	Activity Activity `xml:"Activity,omitempty"`
}

/******************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AttachInstances.html]
type AttachInstancesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateAutoScalingGroup.html]
type CreateAutoScalingGroupResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html]
type CreateLaunchConfigurationResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateOrUpdateTags.html]
type CreateOrUpdateTagsResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteAutoScalingGroup.html]
type DeleteAutoScalingGroupResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteLaunchConfiguration.html]
type DeleteLaunchConfigurationResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteNotificationConfiguration.html]
type DeleteNotificationConfigurationResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeletePolicy.html]
type DeletePolicyResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteScheduledAction.html]
type DeleteScheduledActionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteTags.html]
type DeleteTagsResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAccountLimits.html]
type DescribeAccountLimitsResponse struct {
	DescribeAccountLimitsResult DescribeAccountLimitsResult `xml:"DescribeAccountLimitsResult"`
	ResponseMetadata            ResponseMetadata            `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAdjustmentTypes.html]
type DescribeAdjustmentTypesResponse struct {
	DescribeAdjustmentTypesResult DescribeAdjustmentTypesResult `xml:"DescribeAdjustmentTypesResult"`
	ResponseMetadata              ResponseMetadata              `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingGroups.html]
type DescribeAutoScalingGroupsResponse struct {
	DescribeAutoScalingGroupsResult DescribeAutoScalingGroupsResult `xml:"DescribeAutoScalingGroupsResult"`
	ResponseMetadata                ResponseMetadata                `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingInstances.html]
type DescribeAutoScalingInstancesResponse struct {
	DescribeAutoScalingInstancesResult DescribeAutoScalingInstancesResult `xml:"DescribeAutoScalingInstancesResult"`
	ResponseMetadata                   ResponseMetadata                   `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingNotificationTypes.html]
type DescribeAutoScalingNotificationTypesResponse struct {
	DescribeAutoScalingNotificationTypesResult DescribeAutoScalingNotificationTypesResult `xml:"DescribeAutoScalingNotificationTypesResult"`
	ResponseMetadata                           ResponseMetadata                           `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeLaunchConfigurations.html]
type DescribeLaunchConfigurationsResponse struct {
	DescribeLaunchConfigurationsResult DescribeLaunchConfigurationsResult `xml:"DescribeLaunchConfigurationsResult"`
	ResponseMetadata                   ResponseMetadata                   `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeMetricCollectionTypes.html]
type DescribeMetricCollectionTypesResponse struct {
	DescribeMetricCollectionTypesResult DescribeMetricCollectionTypesResult `xml:"DescribeMetricCollectionTypesResult"`
	ResponseMetadata                    ResponseMetadata                    `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html]
type DescribeNotificationConfigurationsResponse struct {
	DescribeNotificationConfigurationsResult DescribeNotificationConfigurationsResult `xml:"DescribeNotificationConfigurationsResult"`
	ResponseMetadata                         ResponseMetadata                         `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribePolicies.html]
type DescribePoliciesResponse struct {
	DescribePoliciesResult DescribePoliciesResult `xml:"DescribePoliciesResult"`
	ResponseMetadata       ResponseMetadata       `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingActivities.html]
type DescribeScalingActivitiesResponse struct {
	DescribeScalingActivitiesResult DescribeScalingActivitiesResult `xml:"DescribeScalingActivitiesResult"`
	ResponseMetadata                ResponseMetadata                `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingProcessTypes.html]
type DescribeScalingProcessTypesResponse struct {
	DescribeScalingProcessTypesResult DescribeScalingProcessTypesResult `xml:"DescribeScalingProcessTypesResult"`
	ResponseMetadata                  ResponseMetadata                  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScheduledActions.html]
type DescribeScheduledActionsResponse struct {
	DescribeScheduledActionsResult DescribeScheduledActionsResult `xml:"DescribeScheduledActionsResult"`
	ResponseMetadata               ResponseMetadata               `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTags.html]
type DescribeTagsResponse struct {
	DescribeTagsResult DescribeTagsResult `xml:"DescribeTagsResult"`
	ResponseMetadata   ResponseMetadata   `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTerminationPolicyTypes.html]
type DescribeTerminationPolicyTypesResponse struct {
	DescribeTerminationPolicyTypesResult DescribeTerminationPolicyTypesResult `xml:"DescribeTerminationPolicyTypesResult"`
	ResponseMetadata                     ResponseMetadata                     `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DisableMetricsCollection.html]
type DisableMetricsCollectionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_EnableMetricsCollection.html]
type EnableMetricsCollectionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ExecutePolicy.html]
type ExecutePolicyResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutNotificationConfiguration.html]
type PutNotificationConfigurationResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScalingPolicy.html]
type PutScalingPolicyResponse struct {
	PutScalingPolicyResult PutScalingPolicyResult `xml:"PutScalingPolicyResult"`
	ResponseMetadata       ResponseMetadata       `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScheduledUpdateGroupAction.html]
type PutScheduledUpdateGroupActionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ResumeProcesses.html]
type ResumeProcessesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetDesiredCapacity.html]
type SetDesiredCapacityResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetInstanceHealth.html]
type SetInstanceHealthResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SuspendProcesses.html]
type SuspendProcessesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_TerminateInstanceInAutoScalingGroup.html]
type TerminateInstanceInAutoScalingGroupResponse struct {
	TerminateInstanceInAutoScalingGroupResult TerminateInstanceInAutoScalingGroupResult `xml:"TerminateInstanceInAutoScalingGroupResult"`
	ResponseMetadata                          ResponseMetadata                          `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_UpdateAutoScalingGroup.html]
type UpdateAutoScalingGroupResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}
