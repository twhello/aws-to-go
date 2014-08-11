//
// Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications that
// coordinate work across distributed components. In Amazon SWF, a task represents a
// logical unit of work that is performed by a component of your application. Coordinating
// tasks across the application involves managing intertask dependencies, scheduling,
// and concurrency in accordance with the logical flow of the application. Amazon SWF
// gives you full control over implementing tasks and coordinating them without worrying
// about underlying complexities such as tracking their progress and maintaining their state.
//
// [http://aws.amazon.com/documentation/swf/]
//
package swf

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"net/http"
)

const ServiceName = "swf"

// SWFService describes the API interface. Instantiate with swf.NewService().
type SWFService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *SWFService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *SWFService) RegionName() string {
	return s.region.Name()
}

func (s *SWFService) Endpoint() string {
	return s.endpoint
}

// Low-level request to S3 service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *SWFService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)
	
	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *SWFService) wrapperSignAndDo(target string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	if err == nil {
		h := req.Header()
		h.Set("Connection", "Keep-Alive")
		h.Set("Content-Type", "application/x-amz-json-1.0")
		h.Set("Content-Encoding", "amz-1.0")
		h.Set("Pragma", "no-cache")
		h.Set("Cache-Control", "no-cache")
		h.Set("X-Amz-Target", target)
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * Simple Workflow Service Methods.
 */

// Returns the number of closed workflow executions within the given domain that meet the specified filtering criteria.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountClosedWorkflowExecutions.html]
func (s *SWFService) CountClosedWorkflowExecutions(req *CountClosedWorkflowExecutionsRequest) (result *WorkflowExecutionCount, err error) {

	result = new(WorkflowExecutionCount)
	err = s.wrapperSignAndDo("SimpleWorkflowService.CountClosedWorkflowExecutions", req, result)
	return
}

// Returns the number of open workflow executions within the given domain that meet the specified filtering criteria.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountOpenWorkflowExecutions.html]
func (s *SWFService) CountOpenWorkflowExecutions(req *CountOpenWorkflowExecutionsRequest) (result *WorkflowExecutionCount, err error) {

	result = new(WorkflowExecutionCount)
	err = s.wrapperSignAndDo("SimpleWorkflowService.CountOpenWorkflowExecutions", req, result)
	return
}

// Returns the estimated number of activity tasks in the specified task list. The count returned is an approximation and
// is not guaranteed to be exact. If you specify a task list that no activity task was ever scheduled in then 0 will be returned.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingActivityTasks.html]
func (s *SWFService) CountPendingActivityTasks(req *CountPendingActivityTasksRequest) (result *WorkflowExecutionCount, err error) {

	result = new(WorkflowExecutionCount)
	err = s.wrapperSignAndDo("SimpleWorkflowService.CountPendingActivityTasks", req, result)
	return
}

// Returns the estimated number of decision tasks in the specified task list. The count returned is an approximation
// and is not guaranteed to be exact. If you specify a task list that no decision task was ever scheduled in then 0 will be returned.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingDecisionTasks.html]
func (s *SWFService) CountPendingDecisionTasks(req *CountPendingDecisionTasksRequest) (result *WorkflowExecutionCount, err error) {

	result = new(WorkflowExecutionCount)
	err = s.wrapperSignAndDo("SimpleWorkflowService.CountPendingDecisionTasks", req, result)
	return
}

// Deprecates the specified activity type. After an activity type has been deprecated, you cannot create new tasks of that activity type.
// Tasks of this type that were scheduled before the type was deprecated will continue to run.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateActivityType.html]
func (s *SWFService) DeprecateActivityType(req *DeprecateActivityTypeRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.DeprecateActivityType", req, nil)
	return
}

// Deprecates the specified domain. After a domain has been deprecated it cannot be used to create new workflow executions or register
// new types. However, you can still use visibility actions on this domain. Deprecating a domain also deprecates all activity and
// workflow types registered in the domain. Executions that were started before the domain was deprecated will continue to run.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateDomain.html]
func (s *SWFService) DeprecateDomain(req *DeprecateDomainRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.DeprecateDomain", req, nil)
	return
}

// Deprecates the specified workflow type. After a workflow type has been deprecated, you cannot create new executions of that type.
// Executions that were started before the type was deprecated will continue to run. A deprecated workflow type may still be used
// when calling visibility actions.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateWorkflowType.html]
func (s *SWFService) DeprecateWorkflowType(req *DeprecateWorkflowTypeRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.DeprecateWorkflowType", req, nil)
	return
}

// Returns information about the specified activity type. This includes configuration settings provided at registration time as
// well as other general information about the type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeActivityType.html]
func (s *SWFService) DescribeActivityType(req *DescribeActivityTypeRequest) (result *ActivityTypeDetail, err error) {

	result = new(ActivityTypeDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.DescribeActivityType", req, result)
	return
}

// Returns information about the specified domain including description and status.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeDomain.html]
func (s *SWFService) DescribeDomain(req *DescribeDomainRequest) (result *DomainDetail, err error) {

	result = new(DomainDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.DomainDetail", req, result)
	return
}

// Returns information about the specified workflow execution including its type and some statistics.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeWorkflowExecution.html]
func (s *SWFService) DescribeWorkflowExecution(req *DescribeWorkflowExecutionRequest) (result *WorkflowExecutionDetail, err error) {

	result = new(WorkflowExecutionDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.DescribeWorkflowExecution", req, result)
	return
}

// Returns information about the specified workflow type. This includes configuration settings specified when the type was
// registered and other information such as creation date, current status, etc.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeWorkflowType.html]
func (s *SWFService) DescribeWorkflowType(req *DescribeWorkflowTypeRequest) (result *WorkflowTypeDetail, err error) {

	result = new(WorkflowTypeDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.DescribeWorkflowType", req, result)
	return
}

// Returns the history of the specified workflow execution. The results may be split into multiple pages. To retrieve subsequent
// pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_GetWorkflowExecutionHistory.html]
func (s *SWFService) GetWorkflowExecutionHistory(req *GetWorkflowExecutionHistoryRequest) (result *WorkflowTypeDetail, err error) {

	result = new(WorkflowTypeDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.GetWorkflowExecutionHistory", req, result)
	return
}

// Returns information about all activities registered in the specified domain that match the specified name and registration status.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListActivityTypes.html]
func (s *SWFService) ListActivityTypes(req *ListActivityTypesRequest) (result *ActivityTypeDetail, err error) {

	result = new(ActivityTypeDetail)
	err = s.wrapperSignAndDo("SimpleWorkflowService.ListActivityTypes", req, result)
	return
}

// Returns a list of closed workflow executions in the specified domain that meet the filtering criteria. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListClosedWorkflowExecutions.html]
func (s *SWFService) ListClosedWorkflowExecutions(req *ListClosedWorkflowExecutionsRequest) (result *WorkflowExecutionInfos, err error) {

	result = new(WorkflowExecutionInfos)
	err = s.wrapperSignAndDo("SimpleWorkflowService.ListClosedWorkflowExecutions", req, result)
	return
}

// Returns the list of domains registered in the account. The results may be split into multiple pages.
// To retrieve subsequent pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListDomains.html]
func (s *SWFService) ListDomains(req *ListDomainsRequest) (result *DomainInfos, err error) {

	result = new(DomainInfos)
	err = s.wrapperSignAndDo("SimpleWorkflowService.ListDomains", req, result)
	return
}

// Returns a list of open workflow executions in the specified domain that meet the filtering criteria. The results may be split into
// multiple pages. To retrieve subsequent pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListOpenWorkflowExecutions.html]
func (s *SWFService) ListOpenWorkflowExecutions(req *ListOpenWorkflowExecutionsRequest) (result *WorkflowExecutionInfos, err error) {

	result = new(WorkflowExecutionInfos)
	err = s.wrapperSignAndDo("SimpleWorkflowService.ListOpenWorkflowExecutions", req, result)
	return
}

// Returns information about workflow types in the specified domain. The results may be split into
// multiple pages that can be retrieved by making the call repeatedly.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListWorkflowTypes.html]
func (s *SWFService) ListWorkflowTypes(req *ListWorkflowTypesRequest) (result *WorkflowTypeInfos, err error) {

	result = new(WorkflowTypeInfos)
	err = s.wrapperSignAndDo("SimpleWorkflowService.ListWorkflowTypes", req, result)
	return
}

// Used by workers to get an ActivityTask from the specified activity taskList. This initiates a long poll,
// where the service holds the HTTP connection open and responds as soon as a task becomes available.
// The maximum time the service holds on to the request before responding is 60 seconds. If no task is available
// within 60 seconds, the poll will return an empty result. An empty result, in this context, means that an
// ActivityTask is returned, but that the value of taskToken is an empty string. If a task is returned, the
// worker should use its type to identify and process it correctly.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_PollForActivityTask.html]
func (s *SWFService) PollForActivityTask(req *PollForActivityTaskRequest) (result *ActivityTask, err error) {

	result = new(ActivityTask)
	err = s.wrapperSignAndDo("SimpleWorkflowService.PollForActivityTask", req, result)
	return
}

// Used by deciders to get a DecisionTask from the specified decision taskList. A decision task may be returned
// for any open workflow execution that is using the specified task list. The task includes a paginated view of
// the history of the workflow execution. The decider should use the workflow type and the history to determine
// how to properly handle the task.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_PollForDecisionTask.html]
func (s *SWFService) PollForDecisionTask(req *PollForDecisionTaskRequest) (result *DecisionTask, err error) {

	result = new(DecisionTask)
	err = s.wrapperSignAndDo("SimpleWorkflowService.PollForDecisionTask", req, result)
	return
}

// Used by activity workers to report to the service that the ActivityTask represented by the specified taskToken is
// still making progress. The worker can also (optionally) specify details of the progress, for example percent complete,
// using the details parameter. This action can also be used by the worker as a mechanism to check if cancellation is
// being requested for the activity task. If a cancellation is being attempted for the specified task, then the boolean
// cancelRequested flag returned by the service is set to true.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RecordActivityTaskHeartbeat.html]
func (s *SWFService) RecordActivityTaskHeartbeat(req *RecordActivityTaskHeartbeatRequest) (result *ActivityTaskStatus, err error) {

	result = new(ActivityTaskStatus)
	err = s.wrapperSignAndDo("SimpleWorkflowService.RecordActivityTaskHeartbeat", req, result)
	return
}

// Registers a new domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RegisterDomain.html]
func (s *SWFService) RegisterDomain(req *RegisterDomainRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RegisterDomain", req, nil)
	return
}

// Registers a new workflow type and its configuration settings in the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RegisterWorkflowType.html]
func (s *SWFService) RegisterWorkflowType(req *RegisterWorkflowTypeRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RegisterWorkflowType", req, nil)
	return
}

// Records a WorkflowExecutionCancelRequested event in the currently running workflow execution identified by the
// given domain, workflowId, and runId. This logically requests the cancellation of the workflow execution as a whole.
// It is up to the decider to take appropriate actions when it receives an execution history with this event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelWorkflowExecution.html]
func (s *SWFService) RequestCancelWorkflowExecution(req *RequestCancelWorkflowExecutionRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RequestCancelWorkflowExecution", req, nil)
	return
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken was successfully canceled.
// Additional details can be optionally provided using the details argument.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskCanceled.html]
func (s *SWFService) RespondActivityTaskCanceled(req *RespondActivityTaskCanceledRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RespondActivityTaskCanceled", req, nil)
	return
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken completed successfully
// with a result (if provided). The result appears in the ActivityTaskCompleted event in the workflow history.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskCompleted.html]
func (s *SWFService) RespondActivityTaskCompleted(req *RespondActivityTaskCompletedRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RespondActivityTaskCompleted", req, nil)
	return
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken has failed with reason (if specified).
// The reason and details appear in the ActivityTaskFailed event added to the workflow history.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskFailed.html]
func (s *SWFService) RespondActivityTaskFailed(req *RespondActivityTaskFailedRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RespondActivityTaskFailed", req, nil)
	return
}

// Used by deciders to tell the service that the DecisionTask identified by the taskToken has successfully completed.
// The decisions argument specifies the list of decisions made while processing the task.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondDecisionTaskCompleted.html]
func (s *SWFService) RespondDecisionTaskCompleted(req *RespondDecisionTaskCompletedRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.RespondDecisionTaskCompleted", req, nil)
	return
}

// Records a WorkflowExecutionSignaled event in the workflow execution history and creates a decision task
// for the workflow execution identified by the given domain, workflowId and runId. The event is recorded with
// the specified user defined signalName and input (if provided).
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_SignalWorkflowExecution.html]
func (s *SWFService) SignalWorkflowExecution(req *SignalWorkflowExecutionRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.SignalWorkflowExecution", req, nil)
	return
}

// Starts an execution of the workflow type in the specified domain using the provided workflowId and input data.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartWorkflowExecution.html]
func (s *SWFService) StartWorkflowExecution(req *StartWorkflowExecutionRequest) (result *Run, err error) {

	result = new(Run)
	err = s.wrapperSignAndDo("SimpleWorkflowService.StartWorkflowExecution", req, result)
	return
}

// Records a WorkflowExecutionTerminated event and forces closure of the workflow execution identified by the
// given domain, runId, and workflowId. The child policy, registered with the workflow type or specified when
// starting this execution, is applied to any open child workflow executions of this workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TerminateWorkflowExecution.html]
func (s *SWFService) TerminateWorkflowExecution(req *TerminateWorkflowExecutionRequest) (err error) {

	err = s.wrapperSignAndDo("SimpleWorkflowService.TerminateWorkflowExecution", req, nil)
	return
}

// Creates a new Simple Workflow Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *SWFService {
	return &SWFService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
