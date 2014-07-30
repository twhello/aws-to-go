package autoscaling

import (
	"time"
)

/*****************************************************************************/

// Attaches one or more Amazon EC2 instances to an existing Auto Scaling group.
// After the instance(s) is attached, it becomes a part of the Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AttachInstances.html]
type AttachInstancesRequest struct {
	AutoScalingGroupName string   `name:"AutoScalingGroupName"`
	InstanceIds          []string `name:"InstanceIds.member.#,omitempty" base:"1"`
}

// Creates a new AttachInstancesRequest.
func NewAttachInstancesRequest(autoScalingGroupName string) *AttachInstancesRequest {
	return &AttachInstancesRequest{AutoScalingGroupName: autoScalingGroupName}
}

/*****************************************************************************/

// Creates a new Auto Scaling group with the specified name and other attributes.
// When the creation request is completed, the Auto Scaling group is ready to be
// used in other calls.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateAutoScalingGroup.html]
type CreateAutoScalingGroupRequest struct {
	AutoScalingGroupName    string          `name:"AutoScalingGroupName"`
	AvailabilityZones       []string        `name:"AvailabilityZones.member.#,omitempty" base:"1"`
	DefaultCooldown         int             `name:"DefaultCooldown,omitempty"`
	DesiredCapacity         int             `name:"DesiredCapacity,omitempty"`
	HealthCheckGracePeriod  int             `name:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType         HealthCheckType `name:"HealthCheckType,omitempty"`
	InstanceId              string          `name:"InstanceId,omitempty"`
	LaunchConfigurationName string          `name:"LaunchConfigurationName,omitempty"`
	LoadBalancerNames       []string        `name:"LoadBalancerNames.member.#,omitempty" base:"1"`
	MaxSize                 int             `name:"MaxSize"`
	MinSize                 int             `name:"MinSize"`
	PlacementGroup          string          `name:"PlacementGroup,omitempty"`
	Tags                    []Tag           `name:"Tags.member.#.,omitempty" base:"1"`
	TerminationPolicies     []string        `name:"TerminationPolicies.member.#,omitempty" base:"1"`
	VPCZoneIdentifier       string          `name:"VPCZoneIdentifier,omitempty"`
}

// Creates a new CreateAutoScalingGroupRequest.
func NewCreateAutoScalingGroupRequest(autoScalingGroupName string, maxSize, minSize int) *CreateAutoScalingGroupRequest {
	return &CreateAutoScalingGroupRequest{
		AutoScalingGroupName: autoScalingGroupName,
		MaxSize:              maxSize, MinSize: minSize,
	}
}

/*****************************************************************************/

// Creates a new launch configuration. The launch configuration name must be unique within
// the scope of the client's AWS account. The maximum limit of launch configurations,
// which by default is 100, must not yet have been met; otherwise, the call will fail.
// When created, the new launch configuration is available for immediate use.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html]
type CreateLaunchConfigurationRequest struct {
	AssociatePublicIpAddress bool                 `name:"AssociatePublicIpAddress,omitempty"`
	BlockDeviceMappings      []BlockDeviceMapping `name:"BlockDeviceMapping.member.#.,omitempty" base:"1"`
	EbsOptimized             bool                 `name:"EbsOptimized,omitempty"`
	IamInstanceProfile       string               `name:"IamInstanceProfile,omitempty"`
	ImageId                  string               `name:"ImageId,omitempty"`
	InstanceId               string               `name:"InstanceId,omitempty"`
	InstanceMonitoring       *InstanceMonitoring  `name:"InstanceMonitoring,omitempty"`
	InstanceType             string               `name:"InstanceType,omitempty"`
	KernelId                 string               `name:"KernelId,omitempty"`
	KeyName                  string               `name:"KeyName,omitempty"`
	LaunchConfigurationName  string               `name:"LaunchConfigurationName"`
	PlacementTenancy         PlacementTenancy     `name:"PlacementTenancy,omitempty"`
	RamdiskId                string               `name:"RamdiskId,omitempty"`
	SecurityGroups           []string             `name:"SecurityGroups.member.#,omitempty" base:"1"`
	SpotPrice                string               `name:"SpotPrice,omitempty"`
	UserData                 string               `name:"UserData,omitempty"`
}

// Creates a new CreateLaunchConfigurationRequest.
func NewCreateLaunchConfigurationRequest(launchConfigurationName string) *CreateLaunchConfigurationRequest {
	return &CreateLaunchConfigurationRequest{LaunchConfigurationName: launchConfigurationName}
}

/*****************************************************************************/

// Creates new tags or updates existing tags for an Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateOrUpdateTags.html]
type CreateOrUpdateTagsRequest struct {
	Tags []Tag `name:"Tags.member.#." base:"1"`
}

// Creates a new CreateOrUpdateTagsRequest.
func NewCreateOrUpdateTagsRequest(tags ...Tag) *CreateOrUpdateTagsRequest {
	return &CreateOrUpdateTagsRequest{tags}
}

/*****************************************************************************/

// Deletes the specified Auto Scaling group if the group has no instances and
// no scaling activities in progress.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteAutoScalingGroup.html]
type DeleteAutoScalingGroupRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName"`
	ForceDelete          bool   `name:"ForceDelete,omitempty"`
}

// Creates a new DeleteAutoScalingGroupRequest.
func NewDeleteAutoScalingGroupRequest(autoScalingGroupName string) *DeleteAutoScalingGroupRequest {
	return &DeleteAutoScalingGroupRequest{AutoScalingGroupName: autoScalingGroupName}
}

/*****************************************************************************/

// Deletes the specified LaunchConfiguration.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteLaunchConfiguration.html]
type DeleteLaunchConfigurationRequest struct {
	LaunchConfigurationName string `name:"LaunchConfigurationName"`
}

// Creates a new DeleteLaunchConfigurationRequest.
func NewDeleteLaunchConfigurationRequest(launchConfigurationName string) *DeleteLaunchConfigurationRequest {
	return &DeleteLaunchConfigurationRequest{LaunchConfigurationName: launchConfigurationName}
}

/*****************************************************************************/

// Deletes notifications created by PutNotificationConfiguration.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteNotificationConfiguration.html]
type DeleteNotificationConfigurationRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName"`
	TopicARN             string `name:"TopicARN"`
}

// Creates a new DeleteNotificationConfigurationRequest.
func NewDeleteNotificationConfigurationRequest(autoScalingGroupName, topicArn string) *DeleteNotificationConfigurationRequest {
	return &DeleteNotificationConfigurationRequest{autoScalingGroupName, topicArn}
}

/*****************************************************************************/

// Deletes a policy created by PutScalingPolicy.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeletePolicy.html]
type DeletePolicyRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName"`
	PolicyName           string `name:"PolicyName"`
}

// Creates a new DeletePolicyRequest.
func NewDeletePolicyRequest(autoScalingGroupName, policyName string) *DeletePolicyRequest {
	return &DeletePolicyRequest{autoScalingGroupName, policyName}
}

/*****************************************************************************/

// Deletes a scheduled action previously created using the PutScheduledUpdateGroupAction.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteScheduledAction.html]
type DeleteScheduledActionRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName,omitempty"`
	ScheduledActionName  string `name:"ScheduledActionName"`
}

// Creates a new DeleteScheduledActionRequest.
func NewDeleteScheduledActionRequest(scheduledActionName string) *DeleteScheduledActionRequest {
	return &DeleteScheduledActionRequest{ScheduledActionName: scheduledActionName}
}

/*****************************************************************************/

// Removes the specified tags or a set of tags from a set of resources.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteTags.html]
type DeleteTagsRequest struct {
	Tags []Tag `name:"Tags.member.#." base:"1"`
}

// Creates a new DeleteTagsRequest.
func NewDeleteTagsRequest(tags ...Tag) *DeleteTagsRequest {
	return &DeleteTagsRequest{tags}
}

/*****************************************************************************/

// Returns the limits for the Auto Scaling resources currently allowed for your AWS account.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAccountLimits.html]
type DescribeAccountLimitsRequest struct{}

/*****************************************************************************/

// Returns policy adjustment types for use in the PutScalingPolicy action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAdjustmentTypes.html]
type DescribeAdjustmentTypesRequest struct{}

/*****************************************************************************/

// Returns a full description of each Auto Scaling group in the given list. This includes all
// Amazon EC2 instances that are members of the group. If a list of names is not provided,
// the service returns the full details of all Auto Scaling groups.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingGroups.html]
type DescribeAutoScalingGroupsRequest struct {
	AutoScalingGroupNames []string `name:"AutoScalingGroupNames.member.#,omitempty" base:"1"`
	MaxRecords            int      `name:"MaxRecords,omitempty"`
	NextToken             string   `name:"NextToken,omitempty"`
}

// Creates a new DescribeAutoScalingGroupsRequest.
func NewDescribeAutoScalingGroupsRequest() *DescribeAutoScalingGroupsRequest {
	return &DescribeAutoScalingGroupsRequest{}
}

/*****************************************************************************/

// Returns a description of each Auto Scaling instance in the InstanceIds list.
// If a list is not provided, the service returns the full details of all instances
// up to a maximum of 50. By default, the service returns a list of 20 items.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingInstances.html]
type DescribeAutoScalingInstancesRequest struct {
	InstanceIds string `name:"InstanceIds.member.#,omitempty" base:"1"`
	MaxRecords  int    `name:"MaxRecords,omitempty"`
	NextToken   string `name:"NextToken,omitempty"`
}

// Creates a new DescribeAutoScalingInstancesRequest.
func NewDescribeAutoScalingInstancesRequest() *DescribeAutoScalingInstancesRequest {
	return &DescribeAutoScalingInstancesRequest{}
}

/*****************************************************************************/

// Returns a list of all notification types that are supported by Auto Scaling.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingNotificationTypes.html]
type DescribeAutoScalingNotificationTypesRequest struct{}

/*****************************************************************************/

// Returns a full description of the launch configurations, or the specified launch configurations, if they exist.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeLaunchConfigurations.html]
type DescribeLaunchConfigurationsRequest struct {
	LaunchConfigurationNames []string `name:"LaunchConfigurationNames.member.#,omitempty" base:"1"`
	MaxRecords               int      `name:"MaxRecords,omitempty"`
	NextToken                string   `name:"NextToken,omitempty"`
}

// Creates a new DescribeLaunchConfigurationsRequest.
func NewDescribeLaunchConfigurationsRequest() *DescribeLaunchConfigurationsRequest {
	return &DescribeLaunchConfigurationsRequest{}
}

/*****************************************************************************/

// Returns a list of metrics and a corresponding list of granularities for each metric.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeMetricCollectionTypes.html]
type DescribeMetricCollectionTypesRequest struct{}

/*****************************************************************************/

// Returns a list of notification actions associated with Auto Scaling groups for specified events.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html]
type DescribeNotificationConfigurationsRequest struct {
	AutoScalingGroupNames []string `name:"AutoScalingGroupNames.member.#,omitempty" base:"1"`
	MaxRecords            int      `name:"MaxRecords,omitempty"`
	NextToken             string   `name:"NextToken,omitempty"`
}

// Creates a new DescribeNotificationConfigurationsRequest.
func NewDescribeNotificationConfigurationsRequest() *DescribeNotificationConfigurationsRequest {
	return &DescribeNotificationConfigurationsRequest{}
}

/*****************************************************************************/

// Returns descriptions of what each policy does. This action supports pagination.
// If the response includes a token, there are more records available. To get the
// additional records, repeat the request with the response token as the NextToken parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribePolicies.html]
type DescribePoliciesRequest struct {
	AutoScalingGroupName string   `name:"AutoScalingGroupName,omitempty"`
	MaxRecords           int      `name:"MaxRecords,omitempty"`
	NextToken            string   `name:"NextToken,omitempty"`
	PolicyNames          []string `name:"PolicyNames.member.#,omitempty" base:"1"`
}

// Creates a new DescribePoliciesRequest.
func NewDescribePoliciesRequest() *DescribePoliciesRequest {
	return &DescribePoliciesRequest{}
}

/*****************************************************************************/

// Returns the scaling activities for the specified Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingActivities.html]
type DescribeScalingActivitiesRequest struct {
	ActivityIds          []string `name:"ActivityIds.member.#,omitempty" base:"1"`
	AutoScalingGroupName string   `name:"AutoScalingGroupName,omitempty"`
	MaxRecords           int      `name:"MaxRecords,omitempty"`
	NextToken            string   `name:"NextToken,omitempty"`
}

// Creates a new DescribeScalingActivitiesRequest.
func NewDescribeScalingActivitiesRequest() *DescribeScalingActivitiesRequest {
	return &DescribeScalingActivitiesRequest{}
}

/*****************************************************************************/

// Returns scaling process types for use in the ResumeProcesses and SuspendProcesses actions.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingProcessTypes.html]
type DescribeScalingProcessTypesRequest struct{}

/*****************************************************************************/

// Lists all the actions scheduled for your Auto Scaling group that haven't been executed.
// To see a list of actions already executed, see the activity record returned in DescribeScalingActivities.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScheduledActions.html]
type DescribeScheduledActionsRequest struct {
	AutoScalingGroupName string    `name:"AutoScalingGroupName,omitempty"`
	EndTime              time.Time `name:"EndTime,omitempty" format:"2006-01-02T15:04:05Z"`
	MaxRecords           int       `name:"MaxRecords,omitempty"`
	NextToken            string    `name:"NextToken,omitempty"`
	ScheduledActionNames []string  `name:"ScheduledActionNames.member.#,omitempty" base:"1"`
	StartTime            time.Time `name:"StartTime,omitempty" format:"2006-01-02T15:04:05Z"`
}

// Creates a new DescribeScheduledActionsRequest.
func NewDescribeScheduledActionsRequest() *DescribeScheduledActionsRequest {
	return &DescribeScheduledActionsRequest{}
}

/*****************************************************************************/

// Lists the Auto Scaling group tags.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTags.html]
type DescribeTagsRequest struct {
	Filters    []Filter `name:"Filters.member.#.,omitempty" base:"1"`
	MaxRecords int      `name:"MaxRecords,omitempty"`
	NextToken  string   `name:"NextToken,omitempty"`
}

// Creates a new DescribeTagsRequest.
func NewDescribeTagsRequest() *DescribeTagsRequest {
	return &DescribeTagsRequest{}
}

/*****************************************************************************/

// Returns a list of all termination policies supported by Auto Scaling.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTerminationPolicyTypes.html]
type DescribeTerminationPolicyTypesRequest struct{}

/*****************************************************************************/

// Disables monitoring of group metrics for the Auto Scaling group specified in AutoScalingGroupName.
// You can specify the list of affected metrics with the Metrics parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DisableMetricsCollection.html]
type DisableMetricsCollectionRequest struct {
	AutoScalingGroupName string                 `name:"AutoScalingGroupName"`
	Metrics              []MetricCollectionType `name:"Metrics.member.#,omitempty" base:"1"`
}

// Creates a new DisableMetricsCollectionRequest.
func NewDisableMetricsCollectionRequest(autoScalingGroupName string) *DisableMetricsCollectionRequest {
	return &DisableMetricsCollectionRequest{AutoScalingGroupName: autoScalingGroupName}
}

/*****************************************************************************/

// Enables monitoring of group metrics for the Auto Scaling group specified in AutoScalingGroupName.
// You can specify the list of enabled metrics with the Metrics parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_EnableMetricsCollection.html]
type EnableMetricsCollectionRequest struct {
	AutoScalingGroupName string                 `name:"AutoScalingGroupName"`
	Granularity          string                 `name:"Granularity"`
	Metrics              []MetricCollectionType `name:"Metrics.member.#,omitempty" base:"1"`
}

// Creates a new EnableMetricsCollectionRequest.
func NewEnableMetricsCollectionRequest(autoScalingGroupName, granularity string) *EnableMetricsCollectionRequest {
	return &EnableMetricsCollectionRequest{AutoScalingGroupName: autoScalingGroupName, Granularity: granularity}
}

/*****************************************************************************/

// Executes the specified policy.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ExecutePolicy.html]
type ExecutePolicyRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName,omitempty"`
	HonorCooldown        bool   `name:"HonorCooldown,omitempty"`
	PolicyName           string `name:"PolicyName"`
}

// Creates a new ExecutePolicyRequest.
func NewExecutePolicyRequest(policyName string) *ExecutePolicyRequest {
	return &ExecutePolicyRequest{PolicyName: policyName}
}

/*****************************************************************************/

// Configures an Auto Scaling group to send notifications when specified events take place.
// Subscribers to this topic can have messages for events delivered to an endpoint such as
// a web server or email address.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutNotificationConfiguration.html]
type PutNotificationConfigurationRequest struct {
	AutoScalingGroupName string   `name:"AutoScalingGroupName"`
	NotificationTypes    []string `name:"NotificationTypes.member.#" base:"1"`
	TopicArn             string   `name:"TopicArn"`
}

// Creates a new PutNotificationConfigurationRequest.
func NewPutNotificationConfigurationRequest(autoScalingGroupName, topicArn string, notificationTypes ...string) *PutNotificationConfigurationRequest {
	return &PutNotificationConfigurationRequest{autoScalingGroupName, notificationTypes, topicArn}
}

/*****************************************************************************/

// Creates or updates a policy for an Auto Scaling group. To update an existing policy,
// use the existing policy name and set the parameter(s) you want to change. Any existing
// parameter not changed in an update to an existing policy is not changed in this update request.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScalingPolicy.html]
type PutScalingPolicyRequest struct {
	AdjustmentType       AdjustmentType `name:"AdjustmentType"`
	AutoScalingGroupName string         `name:"AutoScalingGroupName"`
	Cooldown             bool           `name:"Cooldown,omitempty"`
	MinAdjustmentStep    int            `name:"MinAdjustmentStep,omitempty"`
	PolicyName           string         `name:"PolicyName"`
	ScalingAdjustment    int            `name:"ScalingAdjustment"`
}

// Creates a new PutScalingPolicyRequest.
func NewPutScalingPolicyRequest(autoScalingGroupName, policyName string, adjustmentType AdjustmentType, scalingAdjustment int) *PutScalingPolicyRequest {
	return &PutScalingPolicyRequest{
		AutoScalingGroupName: autoScalingGroupName,
		PolicyName:           policyName,
		AdjustmentType:       adjustmentType,
		ScalingAdjustment:    scalingAdjustment,
	}
}

/*****************************************************************************/

// Creates or updates a scheduled scaling action for an Auto Scaling group. When updating
// a scheduled scaling action, if you leave a parameter unspecified, the corresponding
// value remains unchanged in the affected Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScheduledUpdateGroupAction.html]
type PutScheduledUpdateGroupActionRequest struct {
	AutoScalingGroupName string    `name:"AutoScalingGroupName"`
	DesiredCapacity      int       `name:"DesiredCapacity,omitempty"`
	EndTime              time.Time `name:"EndTime,omitempty" format:"2006-01-02T15:04:05Z"`
	MaxSize              int       `name:"MaxSize,omitempty"`
	MinSize              int       `name:"MinSize,omitempty"`
	Recurrence           string    `name:"Recurrence,omitempty"`
	ScheduledActionName  string    `name:"ScheduledActionName"`
	StartTime            time.Time `name:"StartTime,omitempty" format:"2006-01-02T15:04:05Z"`
}

// Creates a new PutScheduledUpdateGroupActionRequest.
func NewPutScheduledUpdateGroupActionRequest(autoScalingGroupName, scheduledActionName string) *PutScheduledUpdateGroupActionRequest {
	return &PutScheduledUpdateGroupActionRequest{AutoScalingGroupName: autoScalingGroupName, ScheduledActionName: scheduledActionName}
}

/*****************************************************************************/

// Resumes all suspended Auto Scaling processes for an Auto Scaling group. For information
// on suspending and resuming Auto Scaling process, see Suspend and Resume Auto Scaling Process.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ResumeProcesses.html]
type ResumeProcessesRequest struct {
	AutoScalingGroupName string        `name:"AutoScalingGroupName"`
	ScalingProcesses     []ProcessType `name:"ScalingProcesses.member.#,omitempty" base:"1"`
}

// Creates a new ResumeProcessesRequest.
func NewResumeProcessesRequest(autoScalingGroupName string) *ResumeProcessesRequest {
	return &ResumeProcessesRequest{AutoScalingGroupName: autoScalingGroupName}
}

/*****************************************************************************/

// Sets the desired size of the specified AutoScalingGroup.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetDesiredCapacity.html]
type SetDesiredCapacityRequest struct {
	AutoScalingGroupName string `name:"AutoScalingGroupName"`
	DesiredCapacity      int    `name:"DesiredCapacity"`
	HonorCooldown        bool   `name:"HonorCooldown,omitempty"`
}

// Creates a new SetDesiredCapacityRequest.
func NewSetDesiredCapacityRequest(autoScalingGroupName string, desiredCapacity int) *SetDesiredCapacityRequest {
	return &SetDesiredCapacityRequest{AutoScalingGroupName: autoScalingGroupName, DesiredCapacity: desiredCapacity}
}

/*****************************************************************************/

// Sets the health status of a specified instance that belongs to any of your Auto Scaling groups.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetInstanceHealth.html]
type SetInstanceHealthRequest struct {
	HealthStatus             HealthStatus `name:"HealthStatus"`
	InstanceId               string       `name:"InstanceId"`
	ShouldRespectGracePeriod bool         `name:"ShouldRespectGracePeriod,omitempty"`
}

// Creates a new SetInstanceHealthRequest.
func NewSetInstanceHealthRequest(instanceId string, healthStatus HealthStatus) *SetInstanceHealthRequest {
	return &SetInstanceHealthRequest{HealthStatus: healthStatus, InstanceId: instanceId}
}

/*****************************************************************************/

// Suspends Auto Scaling processes for an Auto Scaling group. To suspend specific process types,
// specify them by name with the ScalingProcesses.member.N parameter. To suspend all process
// types, omit the ScalingProcesses.member.N parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SuspendProcesses.html]
type SuspendProcessesRequest struct {
	AutoScalingGroupName string        `name:"AutoScalingGroupName"`
	ScalingProcesses     []ProcessType `name:"ScalingProcesses.member.#,omitempty" base:"1"`
}

// Creates a new SuspendProcessesRequest.
func NewSuspendProcessesRequest(autoScalingGroupName string) *SuspendProcessesRequest {
	return &SuspendProcessesRequest{AutoScalingGroupName: autoScalingGroupName}
}

/*****************************************************************************/

// Terminates the specified instance. Optionally, the desired group size can be adjusted.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_TerminateInstanceInAutoScalingGroup.html]
type TerminateInstanceInAutoScalingGroupRequest struct {
	InstanceId                     string `name:"InstanceId"`
	ShouldDecrementDesiredCapacity bool   `name:"ShouldDecrementDesiredCapacity"`
}

// Creates a new .
func NewTerminateInstanceInAutoScalingGroupRequest(instanceId string, shouldDecrementDesiredCapacity bool) *TerminateInstanceInAutoScalingGroupRequest {
	return &TerminateInstanceInAutoScalingGroupRequest{instanceId, shouldDecrementDesiredCapacity}
}

/*****************************************************************************/

// Updates the configuration for the specified AutoScalingGroup.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_UpdateAutoScalingGroup.html]
type UpdateAutoScalingGroupRequest struct {
	AutoScalingGroupName    string          `name:"AutoScalingGroupName"`
	AvailabilityZones       []string        `name:"AvailabilityZones.member.#,omitempty" base:"1"`
	DefaultCooldown         int             `name:"DefaultCooldown,omitempty"`
	DesiredCapacity         int             `name:"DesiredCapacity,omitempty"`
	HealthCheckGracePeriod  int             `name:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType         HealthCheckType `name:"HealthCheckType,omitempty"`
	LaunchConfigurationName string          `name:"LaunchConfigurationName,omitempty"`
	MaxSize                 int             `name:"MaxSize"`
	MinSize                 int             `name:"MinSize"`
	PlacementGroup          string          `name:"PlacementGroup,omitempty"`
	TerminationPolicies     []string        `name:"TerminationPolicies.member.#,omitempty" base:"1"`
	VPCZoneIdentifier       string          `name:"VPCZoneIdentifier,omitempty"`
}

// Creates a new UpdateAutoScalingGroupRequest.
func NewUpdateAutoScalingGroupRequest(autoScalingGroupName string, maxSize, minSize int) *UpdateAutoScalingGroupRequest {
	return &UpdateAutoScalingGroupRequest{
		AutoScalingGroupName: autoScalingGroupName,
		MaxSize:              maxSize, MinSize: minSize,
	}
}
