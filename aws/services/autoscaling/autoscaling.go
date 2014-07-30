//
// Auto Scaling is a web service designed to automatically launch or terminate
// Amazon Elastic Compute Cloud (Amazon EC2) instances based on user-defined policies,
// schedules, and health checks. This service is used in conjunction with
// Amazon CloudWatch and Elastic Load Balancing services.
//
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/Welcome.html]
//
package autoscaling

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"github.com/twhello/aws-to-go/aws/util/netutil"
	"net/http"
)

const ServiceName = "autoscaling"

// The Auto Scaling Service object. Use autoscaling.NewService().
type AutoScalingService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *AutoScalingService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *AutoScalingService) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *AutoScalingService) Endpoint() string {
	return s.endpoint
}

// Low-level request to Auto Scaling service.
func (s *AutoScalingService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *AutoScalingService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	qs := netutil.MarshalValues(request)
	qs.Add("Version", "2011-01-01")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", s.Endpoint(), qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * Auto Scaling Service Methods.
 */

// Attaches one or more Amazon EC2 instances to an existing Auto Scaling group.
// After the instance(s) is attached, it becomes a part of the Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_AttachInstances.html]
func (s *AutoScalingService) AttachInstances(req *AttachInstancesRequest) (result *AttachInstancesResponse, err error) {

	result = new(AttachInstancesResponse)
	err = s.wrapperSignAndDo("AttachInstances", req, result)
	return
}

// Creates a new Auto Scaling group with the specified name and other attributes.
// When the creation request is completed, the Auto Scaling group is ready to be
// used in other calls.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateAutoScalingGroup.html]
func (s *AutoScalingService) CreateAutoScalingGroup(req *CreateAutoScalingGroupRequest) (result *CreateAutoScalingGroupResponse, err error) {

	result = new(CreateAutoScalingGroupResponse)
	err = s.wrapperSignAndDo("CreateAutoScalingGroup", req, result)
	return
}

// Creates a new launch configuration. The launch configuration name must be unique within
// the scope of the client's AWS account. The maximum limit of launch configurations,
// which by default is 100, must not yet have been met; otherwise, the call will fail.
// When created, the new launch configuration is available for immediate use.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html]
func (s *AutoScalingService) CreateLaunchConfiguration(req *CreateLaunchConfigurationRequest) (result *CreateLaunchConfigurationResponse, err error) {

	result = new(CreateLaunchConfigurationResponse)
	err = s.wrapperSignAndDo("CreateLaunchConfiguration", req, result)
	return
}

// Creates new tags or updates existing tags for an Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateOrUpdateTags.html]
func (s *AutoScalingService) CreateOrUpdateTags(req *CreateOrUpdateTagsRequest) (result *CreateOrUpdateTagsResponse, err error) {

	result = new(CreateOrUpdateTagsResponse)
	err = s.wrapperSignAndDo("CreateOrUpdateTags", req, result)
	return
}

// Deletes the specified Auto Scaling group if the group has no instances and
// no scaling activities in progress.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteAutoScalingGroup.html]
func (s *AutoScalingService) DeleteAutoScalingGroup(req *DeleteAutoScalingGroupRequest) (result *DeleteAutoScalingGroupResponse, err error) {

	result = new(DeleteAutoScalingGroupResponse)
	err = s.wrapperSignAndDo("DeleteAutoScalingGroup", req, result)
	return
}

// Deletes the specified LaunchConfiguration.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteLaunchConfiguration.html]
func (s *AutoScalingService) DeleteLaunchConfiguration(req *DeleteLaunchConfigurationRequest) (result *DeleteLaunchConfigurationResponse, err error) {

	result = new(DeleteLaunchConfigurationResponse)
	err = s.wrapperSignAndDo("DeleteLaunchConfiguration", req, result)
	return
}

// Deletes notifications created by PutNotificationConfiguration.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteNotificationConfiguration.html]
func (s *AutoScalingService) DeleteNotificationConfiguration(req *DeleteNotificationConfigurationRequest) (result *DeleteNotificationConfigurationResponse, err error) {

	result = new(DeleteNotificationConfigurationResponse)
	err = s.wrapperSignAndDo("DeleteNotificationConfiguration", req, result)
	return
}

// Deletes a policy created by PutScalingPolicy.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeletePolicy.html]
func (s *AutoScalingService) DeletePolicy(req *DeletePolicyRequest) (result *DeletePolicyResponse, err error) {

	result = new(DeletePolicyResponse)
	err = s.wrapperSignAndDo("DeletePolicy", req, result)
	return
}

// Deletes a scheduled action previously created using the PutScheduledUpdateGroupAction.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteScheduledAction.html]
func (s *AutoScalingService) DeleteScheduledAction(req *DeleteScheduledActionRequest) (result *DeleteScheduledActionResponse, err error) {

	result = new(DeleteScheduledActionResponse)
	err = s.wrapperSignAndDo("DeleteScheduledAction", req, result)
	return
}

// Removes the specified tags or a set of tags from a set of resources.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DeleteTags.html]
func (s *AutoScalingService) DeleteTags(req *DeleteTagsRequest) (result *DeleteTagsResponse, err error) {

	result = new(DeleteTagsResponse)
	err = s.wrapperSignAndDo("DeleteTags", req, result)
	return
}

// Returns the limits for the Auto Scaling resources currently allowed for your AWS account.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAccountLimits.html]
func (s *AutoScalingService) DescribeAccountLimits() (result *DescribeAccountLimitsResponse, err error) {

	result = new(DescribeAccountLimitsResponse)
	err = s.wrapperSignAndDo("DescribeAccountLimits", nil, result)
	return
}

// Returns policy adjustment types for use in the PutScalingPolicy action.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAdjustmentTypes.html]
func (s *AutoScalingService) DescribeAdjustmentTypes() (result *DescribeAdjustmentTypesResponse, err error) {

	result = new(DescribeAdjustmentTypesResponse)
	err = s.wrapperSignAndDo("DescribeAdjustmentTypes", nil, result)
	return
}

// Returns a full description of each Auto Scaling group in the given list. This includes all
// Amazon EC2 instances that are members of the group. If a list of names is not provided,
// the service returns the full details of all Auto Scaling groups.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingGroups.html]
func (s *AutoScalingService) DescribeAutoScalingGroups(req *DescribeAutoScalingGroupsRequest) (result *DescribeAutoScalingGroupsResponse, err error) {

	result = new(DescribeAutoScalingGroupsResponse)
	err = s.wrapperSignAndDo("DescribeAutoScalingGroups", req, result)
	return
}

// Returns a description of each Auto Scaling instance in the InstanceIds list.
// If a list is not provided, the service returns the full details of all instances
// up to a maximum of 50. By default, the service returns a list of 20 items.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingInstances.html]
func (s *AutoScalingService) DescribeAutoScalingInstances(req *DescribeAutoScalingInstancesRequest) (result *DescribeAutoScalingInstancesResponse, err error) {

	result = new(DescribeAutoScalingInstancesResponse)
	err = s.wrapperSignAndDo("DescribeAutoScalingInstances", req, result)
	return
}

// Returns a list of all notification types that are supported by Auto Scaling.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeAutoScalingNotificationTypes.html]
func (s *AutoScalingService) DescribeAutoScalingNotificationTypes() (result *DescribeAutoScalingNotificationTypesResponse, err error) {

	result = new(DescribeAutoScalingNotificationTypesResponse)
	err = s.wrapperSignAndDo("DescribeAutoScalingNotificationTypes", nil, result)
	return
}

// Returns a full description of the launch configurations, or the specified launch configurations, if they exist.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeLaunchConfigurations.html]
func (s *AutoScalingService) DescribeLaunchConfigurations(req *DescribeLaunchConfigurationsRequest) (result *DescribeLaunchConfigurationsResponse, err error) {

	result = new(DescribeLaunchConfigurationsResponse)
	err = s.wrapperSignAndDo("DescribeLaunchConfigurations", req, result)
	return
}

// Returns a list of metrics and a corresponding list of granularities for each metric.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeMetricCollectionTypes.html]
func (s *AutoScalingService) DescribeMetricCollectionTypes() (result *DescribeMetricCollectionTypesResponse, err error) {

	result = new(DescribeMetricCollectionTypesResponse)
	err = s.wrapperSignAndDo("DescribeMetricCollectionTypes", nil, result)
	return
}

// Returns a list of notification actions associated with Auto Scaling groups for specified events.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html]
func (s *AutoScalingService) DescribeNotificationConfigurations(req *DescribeNotificationConfigurationsRequest) (result *DescribeNotificationConfigurationsResponse, err error) {

	result = new(DescribeNotificationConfigurationsResponse)
	err = s.wrapperSignAndDo("DescribeNotificationConfigurations", req, result)
	return
}

// Returns descriptions of what each policy does. This action supports pagination.
// If the response includes a token, there are more records available. To get the
// additional records, repeat the request with the response token as the NextToken parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribePolicies.html]
func (s *AutoScalingService) DescribePolicies(req *DescribePoliciesRequest) (result *DescribePoliciesResponse, err error) {

	result = new(DescribePoliciesResponse)
	err = s.wrapperSignAndDo("DescribePolicies", req, result)
	return
}

// Returns the scaling activities for the specified Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingActivities.html]
func (s *AutoScalingService) DescribeScalingActivities(req *DescribeScalingActivitiesRequest) (result *DescribeScalingActivitiesResponse, err error) {

	result = new(DescribeScalingActivitiesResponse)
	err = s.wrapperSignAndDo("DescribeScalingActivities", req, result)
	return
}

// Returns scaling process types for use in the ResumeProcesses and SuspendProcesses actions.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScalingProcessTypes.html]
func (s *AutoScalingService) DescribeScalingProcessTypes() (result *DescribeScalingProcessTypesResponse, err error) {

	result = new(DescribeScalingProcessTypesResponse)
	err = s.wrapperSignAndDo("DescribeScalingProcessTypes", nil, result)
	return
}

// Lists all the actions scheduled for your Auto Scaling group that haven't been executed.
// To see a list of actions already executed, see the activity record returned in DescribeScalingActivities.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeScheduledActions.html]
func (s *AutoScalingService) DescribeScheduledActions(req *DescribeScheduledActionsRequest) (result *DescribeScheduledActionsResponse, err error) {

	result = new(DescribeScheduledActionsResponse)
	err = s.wrapperSignAndDo("DescribeScheduledActions", req, result)
	return
}

// Lists the Auto Scaling group tags.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTags.html]
func (s *AutoScalingService) DescribeTags(req *DescribeTagsRequest) (result *DescribeTagsResponse, err error) {

	result = new(DescribeTagsResponse)
	err = s.wrapperSignAndDo("DescribeTags", req, result)
	return
}

// Returns a list of all termination policies supported by Auto Scaling.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeTerminationPolicyTypes.html]
func (s *AutoScalingService) DescribeTerminationPolicyTypes() (result *DescribeTerminationPolicyTypesResponse, err error) {

	result = new(DescribeTerminationPolicyTypesResponse)
	err = s.wrapperSignAndDo("DescribeTerminationPolicyTypes", nil, result)
	return
}

// Disables monitoring of group metrics for the Auto Scaling group specified in AutoScalingGroupName.
// You can specify the list of affected metrics with the Metrics parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DisableMetricsCollection.html]
func (s *AutoScalingService) DisableMetricsCollection(req *DisableMetricsCollectionRequest) (result *DisableMetricsCollectionResponse, err error) {

	result = new(DisableMetricsCollectionResponse)
	err = s.wrapperSignAndDo("DisableMetricsCollection", req, result)
	return
}

// Enables monitoring of group metrics for the Auto Scaling group specified in AutoScalingGroupName.
// You can specify the list of enabled metrics with the Metrics parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_EnableMetricsCollection.html]
func (s *AutoScalingService) EnableMetricsCollection(req *EnableMetricsCollectionRequest) (result *EnableMetricsCollectionResponse, err error) {

	result = new(EnableMetricsCollectionResponse)
	err = s.wrapperSignAndDo("EnableMetricsCollection", req, result)
	return
}

// Executes the specified policy.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ExecutePolicy.html]
func (s *AutoScalingService) ExecutePolicy(req *ExecutePolicyRequest) (result *ExecutePolicyResponse, err error) {

	result = new(ExecutePolicyResponse)
	err = s.wrapperSignAndDo("ExecutePolicy", req, result)
	return
}

// Configures an Auto Scaling group to send notifications when specified events take place.
// Subscribers to this topic can have messages for events delivered to an endpoint such as
// a web server or email address.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutNotificationConfiguration.html]
func (s *AutoScalingService) PutNotificationConfiguration(req *PutNotificationConfigurationRequest) (result *PutNotificationConfigurationResponse, err error) {

	result = new(PutNotificationConfigurationResponse)
	err = s.wrapperSignAndDo("PutNotificationConfiguration", req, result)
	return
}

// Creates or updates a policy for an Auto Scaling group. To update an existing policy,
// use the existing policy name and set the parameter(s) you want to change. Any existing
// parameter not changed in an update to an existing policy is not changed in this update request.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScalingPolicy.html]
func (s *AutoScalingService) PutScalingPolicy(req *PutScalingPolicyRequest) (result *PutScalingPolicyResponse, err error) {

	result = new(PutScalingPolicyResponse)
	err = s.wrapperSignAndDo("PutScalingPolicy", req, result)
	return
}

// Creates or updates a scheduled scaling action for an Auto Scaling group. When updating
// a scheduled scaling action, if you leave a parameter unspecified, the corresponding
// value remains unchanged in the affected Auto Scaling group.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_PutScheduledUpdateGroupAction.html]
func (s *AutoScalingService) PutScheduledUpdateGroupAction(req *PutScheduledUpdateGroupActionRequest) (result *PutScheduledUpdateGroupActionResponse, err error) {

	result = new(PutScheduledUpdateGroupActionResponse)
	err = s.wrapperSignAndDo("PutScheduledUpdateGroupAction", req, result)
	return
}

// Resumes all suspended Auto Scaling processes for an Auto Scaling group. For information
// on suspending and resuming Auto Scaling process, see Suspend and Resume Auto Scaling Process.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_ResumeProcesses.html]
func (s *AutoScalingService) ResumeProcesses(req *ResumeProcessesRequest) (result *ResumeProcessesResponse, err error) {

	result = new(ResumeProcessesResponse)
	err = s.wrapperSignAndDo("ResumeProcesses", req, result)
	return
}

// Sets the desired size of the specified AutoScalingGroup.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetDesiredCapacity.html]
func (s *AutoScalingService) SetDesiredCapacity(req *SetDesiredCapacityRequest) (result *SetDesiredCapacityResponse, err error) {

	result = new(SetDesiredCapacityResponse)
	err = s.wrapperSignAndDo("SetDesiredCapacity", req, result)
	return
}

// Sets the health status of a specified instance that belongs to any of your Auto Scaling groups.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SetInstanceHealth.html]
func (s *AutoScalingService) SetInstanceHealth(req *SetInstanceHealthRequest) (result *SetInstanceHealthResponse, err error) {

	result = new(SetInstanceHealthResponse)
	err = s.wrapperSignAndDo("SetInstanceHealth", req, result)
	return
}

// Suspends Auto Scaling processes for an Auto Scaling group. To suspend specific process types,
// specify them by name with the ScalingProcesses.member.N parameter. To suspend all process
// types, omit the ScalingProcesses.member.N parameter.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_SuspendProcesses.html]
func (s *AutoScalingService) SuspendProcesses(req *SuspendProcessesRequest) (result *SuspendProcessesResponse, err error) {

	result = new(SuspendProcessesResponse)
	err = s.wrapperSignAndDo("SuspendProcesses", req, result)
	return
}

// Terminates the specified instance. Optionally, the desired group size can be adjusted.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_TerminateInstanceInAutoScalingGroup.html]
func (s *AutoScalingService) TerminateInstanceInAutoScalingGroup(req *TerminateInstanceInAutoScalingGroupRequest) (result *TerminateInstanceInAutoScalingGroupResponse, err error) {

	result = new(TerminateInstanceInAutoScalingGroupResponse)
	err = s.wrapperSignAndDo("TerminateInstanceInAutoScalingGroup", req, result)
	return
}

// Updates the configuration for the specified AutoScalingGroup.
// [http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_UpdateAutoScalingGroup.html]
func (s *AutoScalingService) UpdateAutoScalingGroup(req *UpdateAutoScalingGroupRequest) (result *UpdateAutoScalingGroupResponse, err error) {

	result = new(UpdateAutoScalingGroupResponse)
	err = s.wrapperSignAndDo("UpdateAutoScalingGroup", req, result)
	return
}

// Creates a new Auto Scaling Service.
func NewService(cred interfaces.IAWSCredentials) *AutoScalingService {

	return &AutoScalingService{cred, nil, "https://" + ServiceName + ".amazonaws.com/"}
}
