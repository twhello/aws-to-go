package swf

import ()

/*****************************************************************************/

// Returns the number of closed workflow executions within the given domain that meet the specified filtering criteria.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountClosedWorkflowExecutions.html]
type CountClosedWorkflowExecutionsRequest struct {
	CloseStatusFilter *CloseStatusFilter       `json:"closeStatusFilter,omitempty"`
	CloseTimeFilter   ExecutionTimeFilter      `json:"closeTimeFilter,omitempty"`
	Domain            string                   `json:"domain"`
	ExecutionFilter   *WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	StartTimeFilter   *ExecutionTimeFilter     `json:"startTimeFilter,omitempty"`
	TagFilter         *TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter        *WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// Creates a new CountClosedWorkflowExecutionsRequest.
// (domain string) The name of the domain containing the workflow executions to count.
func NewCountClosedWorkflowExecutionsRequest(domain string) *CountClosedWorkflowExecutionsRequest {
	return &CountClosedWorkflowExecutionsRequest{Domain: domain}
}

/*****************************************************************************/

// Returns the number of open workflow executions within the given domain that meet the specified filtering criteria.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountOpenWorkflowExecutions.html]
type CountOpenWorkflowExecutionsRequest struct {
	Domain           string                   `json:"domain"`
	ExecutionFilter  *WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	StartTimerFilter ExecutionTimeFilter      `json:"startTimerFilter"`
	TagFilter        *TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter       *WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// Creates a new CountOpenWorkflowExecutionsRequest.
// (domain string) The name of the domain containing the workflow executions to count.
// (startTimerFilter ExecutionTimeFilter) Specifies the start time criteria that workflow executions must meet in order to be counted.
func NewCountOpenWorkflowExecutionsRequest(domain string, startTimerFilter ExecutionTimeFilter) *CountOpenWorkflowExecutionsRequest {
	return &CountOpenWorkflowExecutionsRequest{Domain: domain, StartTimerFilter: startTimerFilter}
}

/*****************************************************************************/

// Returns the estimated number of activity tasks in the specified task list.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingActivityTasks.html]
type CountPendingActivityTasksRequest struct {
	Domain   string    `json:"domain"`
	TaskList *TaskList `json:"taskList"`
}

// Creates a new CountPendingActivityTasksRequest.
// (domain string) The name of the domain that contains the task list.
// (taskList TaskList) The name of the task list.
func NewCountPendingActivityTasksRequest(domain string, taskList *TaskList) *CountPendingActivityTasksRequest {
	return &CountPendingActivityTasksRequest{domain, taskList}
}

/*****************************************************************************/

// Returns the estimated number of decision tasks in the specified task list.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingDecisionTasks.html]
type CountPendingDecisionTasksRequest struct {
	Domain   string    `json:"domain"`
	TaskList *TaskList `json:"taskList"`
}

// Creates a new CountPendingDecisionTasksRequest.
// (domain string) The name of the domain that contains the task list.
// (taskList TaskList) The name of the task list.
func NewCountPendingDecisionTasksRequest(domain string, taskList *TaskList) *CountPendingDecisionTasksRequest {
	return &CountPendingDecisionTasksRequest{domain, taskList}
}

/*****************************************************************************/

// Deprecates the specified activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateActivityType.html]
type DeprecateActivityTypeRequest struct {
	ActivityType *ActivityType `json:"activityType"`
	Domain       string        `json:"domain"`
}

// Creates a new DeprecateActivityTypeRequest.
// (domain string) The name of the domain in which the activity type is registered.
// (activityType ActivityType) The activity type to deprecate.
func NewDeprecateActivityTypeRequest(domain string, activityType *ActivityType) *DeprecateActivityTypeRequest {
	return &DeprecateActivityTypeRequest{activityType, domain}
}

/*****************************************************************************/

// Deprecates the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateDomain.html]
type DeprecateDomainRequest struct {
	Name string `json:"name"`
}

// Creates a new DeprecateDomainRequest.
// (name string) The name of the domain to deprecate.
func NewDeprecateDomainRequest(name string) *DeprecateDomainRequest {
	return &DeprecateDomainRequest{name}
}

/*****************************************************************************/

// Deprecates the specified workflow type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DeprecateWorkflowType.html]
type DeprecateWorkflowTypeRequest struct {
	Domain       string        `json:"domain"`
	WorkflowType *WorkflowType `json:"workflowType"`
}

// Creates a new DeprecateWorkflowTypeRequest.
// (domain string) The name of the domain in which the workflow type is registered.
// (workflowType WorkflowType) The workflow type to deprecate.
func NewDeprecateWorkflowTypeRequest(domain string, workflowType *WorkflowType) *DeprecateWorkflowTypeRequest {
	return &DeprecateWorkflowTypeRequest{domain, workflowType}
}

/*****************************************************************************/

// Returns information about the specified activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeActivityType.html]
type DescribeActivityTypeRequest struct {
	ActivityType *ActivityType `json:"activityType"`
	Domain       string        `json:"domain"`
}

// Creates a new DescribeActivityTypeRequest.
// (domain string) The name of the domain in which the activity type is registered.
// (activityType ActivityType) The activity type to describe.
func NewDescribeActivityTypeRequest(domain string, activityType *ActivityType) *DescribeActivityTypeRequest {
	return &DescribeActivityTypeRequest{activityType, domain}
}

/*****************************************************************************/

// Returns information about the specified domain including description and status.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeDomain.html]
type DescribeDomainRequest struct {
	Name string `json:"name"`
}

// Creates a new DescribeDomainRequest.
// (name string) The name of the domain to describe.
func NewDescribeDomainRequest(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{name}
}

/*****************************************************************************/

// Returns information about the specified workflow execution including its type and some statistics.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeWorkflowExecution.html]
type DescribeWorkflowExecutionRequest struct {
	Domain    string             `json:"domain"`
	Execution *WorkflowExecution `json:"execution"`
}

// Creates a new DescribeWorkflowExecutionRequest.
// (domain string) The name of the domain containing the workflow execution.
// (execution WorkflowExecution) The workflow execution to describe.
func NewDescribeWorkflowExecutionRequest(domain string, execution *WorkflowExecution) *DescribeWorkflowExecutionRequest {
	return &DescribeWorkflowExecutionRequest{domain, execution}
}

/*****************************************************************************/

// Returns information about the specified workflow type. This includes configuration settings specified when
// the type was registered and other information such as creation date, current status, etc.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DescribeWorkflowType.html]
type DescribeWorkflowTypeRequest struct {
	Domain       string        `json:"domain"`
	WorkflowType *WorkflowType `json:"workflowType"`
}

// Creates a new DescribeWorkflowTypeRequest.
// (domain string) The name of the domain in which this workflow type is registered.
// (workflowType WorkflowType) The workflow type to describe.
func NewDescribeWorkflowTypeRequest(domain string, workflowType *WorkflowType) *DescribeWorkflowTypeRequest {
	return &DescribeWorkflowTypeRequest{domain, workflowType}
}

/*****************************************************************************/

// Returns the history of the specified workflow execution. The results may be split into multiple pages.
// To retrieve subsequent pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_GetWorkflowExecutionHistory.html]
type GetWorkflowExecutionHistoryRequest struct {
	Domain          string             `json:"domain"`
	Execution       *WorkflowExecution `json:"execution"`
	MaximumPageSize int64              `json:"maximumPageSize,omitempty"`
	NextPageToken   string             `json:"nextPageToken,omitempty"`
	ReverseOrder    bool               `json:"reverseOrder,omitempty"`
}

// Creates a new GetWorkflowExecutionHistoryRequest.
// (domain string) The name of the domain containing the workflow execution.
// (execution WorkflowExecution) Specifies the workflow execution for which to return the history.
func NewGetWorkflowExecutionHistoryRequest(domain string, execution *WorkflowExecution) *GetWorkflowExecutionHistoryRequest {
	return &GetWorkflowExecutionHistoryRequest{Domain: domain, Execution: execution}
}

/*****************************************************************************/

// Returns information about all activities registered in the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListActivityTypes.html]
type ListActivityTypesRequest struct {
	Domain             string `json:"domain"`
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	Name               string `json:"name,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus Status `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// Creates a new ListActivityTypesRequest.
// (domain string) The name of the domain in which the activity types have been registered.
// (registrationStatus Status) Specifies the registration status of the activity types to list. (REGISTERED, DEPRECATED)
func NewListActivityTypesRequest(domain string, registrationStatus Status) *ListActivityTypesRequest {
	return &ListActivityTypesRequest{Domain: domain, RegistrationStatus: registrationStatus}
}

/*****************************************************************************/

// Returns a list of closed workflow executions in the specified domain that meet the filtering criteria.
// The results may be split into multiple pages. To retrieve subsequent pages, make the call again using
// the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListClosedWorkflowExecutions.html]
type ListClosedWorkflowExecutionsRequest struct {
	CloseStatusFilter *CloseStatusFilter       `json:"closeStatusFilter,omitempty"`
	CloseTimeFilter   ExecutionTimeFilter      `json:"closeTimeFilter,omitempty"`
	Domain            string                   `json:"domain"`
	ExecutionFilter   *WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	MaximumPageSize   int                      `json:"maximumPageSize,omitempty"`
	NextPageToken     string                   `json:"reverseOrder,omitempty"`
	ReverseOrder      bool                     `json:"reverseOrder,omitempty"`
	StartTimeFilter   *ExecutionTimeFilter     `json:"startTimeFilter,omitempty"`
	TagFilter         *TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter        *WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// Creates a new ListClosedWorkflowExecutionsRequest.
// (domain string) The name of the domain that contains the workflow executions to list.
func NewListClosedWorkflowExecutionsRequest(domain string) *ListClosedWorkflowExecutionsRequest {
	return &ListClosedWorkflowExecutionsRequest{Domain: domain}
}

/*****************************************************************************/

// Returns the list of domains registered in the account. The results may be split into multiple pages.
// To retrieve subsequent pages, make the call again using the nextPageToken returned by the initial call.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListDomains.html]
type ListDomainsRequest struct {
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus Status `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// Creates a new ListDomainsRequest.
// (registrationStatus Status) Specifies the registration status of the domains to list. (REGISTERED, DEPRECATED)
func NewListDomainsRequest(registrationStatus Status) *ListDomainsRequest {
	return &ListDomainsRequest{RegistrationStatus: registrationStatus}
}

/*****************************************************************************/

// Returns a list of open workflow executions in the specified domain that meet the filtering criteria.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListOpenWorkflowExecutions.html]
type ListOpenWorkflowExecutionsRequest struct {
	Domain          string                   `json:"domain"`
	ExecutionFilter *WorkflowExecutionFilter `json:"executionFilter,omitempty"`
	MaximumPageSize int                      `json:"maximumPageSize,omitempty"`
	NextPageToken   string                   `json:"nextPageToken,omitempty"`
	ReverseOrder    bool                     `json:"reverseOrder,omitempty"`
	StartTimeFilter *ExecutionTimeFilter     `json:"startTimeFilter,omitempty"`
	TagFilter       *TagFilter               `json:"tagFilter,omitempty"`
	TypeFilter      *WorkflowTypeFilter      `json:"typeFilter,omitempty"`
}

// Creates a new ListOpenWorkflowExecutionsRequest.
// (domain string) The name of the domain that contains the workflow executions to list.
func NewListOpenWorkflowExecutionsRequest(domain string) *ListOpenWorkflowExecutionsRequest {
	return &ListOpenWorkflowExecutionsRequest{Domain: domain}
}

/*****************************************************************************/

// Returns information about workflow types in the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ListWorkflowTypes.html]
type ListWorkflowTypesRequest struct {
	Domain             string `json:"domain"`
	MaximumPageSize    int    `json:"maximumPageSize,omitempty"`
	Name               string `json:"name,omitempty"`
	NextPageToken      string `json:"nextPageToken,omitempty"`
	RegistrationStatus Status `json:"registrationStatus"`
	ReverseOrder       bool   `json:"reverseOrder,omitempty"`
}

// Creates a new ListWorkflowTypesRequest.
// (domain string) The name of the domain in which the workflow types have been registered.
func NewListWorkflowTypesRequest(domain string) *ListWorkflowTypesRequest {
	return &ListWorkflowTypesRequest{Domain: domain}
}

/*****************************************************************************/

// Used by workers to get an ActivityTask from the specified activity taskList.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_PollForActivityTask.html]
type PollForActivityTaskRequest struct {
	Domain   string    `json:"domain"`
	Identity string    `json:"identity,omitempty"`
	TaskList *TaskList `json:"taskList"`
}

// Creates a new PollForActivityTaskRequest.
// (domain string) The name of the domain that contains the task lists being polled.
// (taskList TaskList) Specifies the task list to poll for activity tasks.
func NewPollForActivityTaskRequest(domain string, taskList *TaskList) *PollForActivityTaskRequest {
	return &PollForActivityTaskRequest{Domain: domain, TaskList: taskList}
}

/*****************************************************************************/

// Used by deciders to get a DecisionTask from the specified decision taskList.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_PollForDecisionTask.html]
type PollForDecisionTaskRequest struct {
	Domain          string    `json:"domain"`
	Identity        string    `json:"identity,omitempty"`
	MaximumPageSize int       `json:"maximumPageSize,omitempty"`
	NextPageToken   string    `json:"nextPageToken,omitempty"`
	ReverseOrder    bool      `json:"reverseOrder,omitempty"`
	TaskList        *TaskList `json:"taskList"`
}

// Creates a new PollForDecisionTaskRequest.
// (domain string) The name of the domain containing the task lists to poll.
// (taskList TaskList) Specifies the task list to poll for decision tasks.
func NewPollForDecisionTaskRequest(domain string, taskList *TaskList) *PollForDecisionTaskRequest {
	return &PollForDecisionTaskRequest{Domain: domain, TaskList: taskList}
}

/*****************************************************************************/

// Used by activity workers to report to the service that the ActivityTask represented by
// the specified taskToken is still making progress.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RecordActivityTaskHeartbeat.html]
type RecordActivityTaskHeartbeatRequest struct {
	Details   string `json:"details,omitempty"`
	TaskToken string `json:"taskToken"`
}

// Creates a new RecordActivityTaskHeartbeatRequest.
// (taskToken string) The taskToken of the ActivityTask.
func NewRecordActivityTaskHeartbeatRequest(taskToken string) *RecordActivityTaskHeartbeatRequest {
	return &RecordActivityTaskHeartbeatRequest{TaskToken: taskToken}
}

/*****************************************************************************/

// Registers a new activity type along with its configuration settings in the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RegisterActivityType.html]
type RegisterActivityTypeRequest struct {
	DefaultTaskHeartbeatTimeout       string    `json:"defaultTaskHeartbeatTimeout,omitempty"`
	DefaultTaskList                   *TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskScheduleToCloseTimeout string    `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	DefaultTaskScheduleToStartTimeout string    `json:"defaultTaskScheduleToStartTimeout,omitempty"`
	DefaultTaskStartToCloseTimeout    string    `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Description                       string    `json:"description,omitempty"`
	Domain                            string    `json:"domain"`
	Name                              string    `json:"name"`
	Version                           string    `json:"version"`
}

// Creates a new RegisterActivityTypeRequest.
// (domain string) The name of the domain in which this activity is to be registered.
// (name string) The name of the activity type within the domain.
// (version string) The version of the activity type.
func NewRegisterActivityTypeRequest(domain, name, version string) *RegisterActivityTypeRequest {
	return &RegisterActivityTypeRequest{Domain: domain, Name: name, Version: version}
}

/*****************************************************************************/

// Registers a new domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RegisterDomain.html]
type RegisterDomainRequest struct {
	Description                            string `json:"description,omitempty"`
	Name                                   string `json:"name"`
	WorkflowExecutionRetentionPeriodInDays string `json:"workflowExecutionRetentionPeriodInDays"`
}

// Creates a new RegisterDomainRequest.
// (name string) Name of the domain to register. The name must be unique in the region that the domain is registered in.
// (workflowExecutionRetentionPeriodInDays string) The duration (in days) that records should be kept.
// (description string) A text description of the domain.
func NewRegisterDomainRequest(name, workflowExecutionRetentionPeriodInDays, description string) *RegisterDomainRequest {
	return &RegisterDomainRequest{description, name, workflowExecutionRetentionPeriodInDays}
}

/*****************************************************************************/

// Registers a new workflow type and its configuration settings in the specified domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RegisterWorkflowType.html]
type RegisterWorkflowTypeRequest struct {
	DefaultChildPolicy                  ChildPolicy `json:"defaultChildPolicy,omitempty"`
	DefaultExecutionStartToCloseTimeout string      `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	DefaultTaskList                     *TaskList   `json:"defaultTaskList,omitempty"`
	DefaultTaskStartToCloseTimeout      string      `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Description                         string      `json:"description,omitempty"`
	Domain                              string      `json:"domain"`
	Name                                string      `json:"name"`
	Version                             string      `json:"version"`
}

// Creates a new RegisterWorkflowTypeRequest.
// (domain string) The name of the domain in which to register the workflow type.
// (name string) The name of the workflow type.
// (version string) The version of the workflow type.
func NewRegisterWorkflowTypeRequest(domain, name, version string) *RegisterWorkflowTypeRequest {
	return &RegisterWorkflowTypeRequest{Domain: domain, Name: name, Version: version}
}

/*****************************************************************************/

// Records a WorkflowExecutionCancelRequested event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelWorkflowExecution.html]
type RequestCancelWorkflowExecutionRequest struct {
	Domain     string `json:"domain"`
	RunId      string `json:"runId,omitempty"`
	WorkflowId string `json:"workflowId"`
}

// Creates a new RequestCancelWorkflowExecutionRequest.
// (domain string) The name of the domain containing the workflow execution to cancel.
// (workflowId string) The workflowId of the workflow execution to cancel.
func NewRequestCancelWorkflowExecutionRequest(domain, workflowId string) *RequestCancelWorkflowExecutionRequest {
	return &RequestCancelWorkflowExecutionRequest{Domain: domain, WorkflowId: workflowId}
}

/*****************************************************************************/

// Used by workers to tell the service that the ActivityTask identified by the taskToken was successfully canceled.
// Additional details can be optionally provided using the details argument.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskCanceled.html]
type RespondActivityTaskCanceledRequest struct {
	Details   string `json:"details,omitempty"`
	TaskToken string `json:"taskToken"`
}

// Creates a new RespondActivityTaskCanceledRequest.
// (taskToken string) The taskToken of the ActivityTask.
func NewRespondActivityTaskCanceledRequest(taskToken string) *RespondActivityTaskCanceledRequest {
	return &RespondActivityTaskCanceledRequest{TaskToken: taskToken}
}

/*****************************************************************************/

// Used by workers to tell the service that the ActivityTask identified by the taskToken completed
// successfully with a result (if provided).
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskCompleted.html]
type RespondActivityTaskCompletedRequest struct {
	Result    string `json:"result,omitempty"`
	TaskToken string `json:"taskToken"`
}

// Creates a new RespondActivityTaskCompletedRequest.
// (taskToken string) The taskToken of the ActivityTask.
func NewRespondActivityTaskCompletedRequest(taskToken string) *RespondActivityTaskCompletedRequest {
	return &RespondActivityTaskCompletedRequest{TaskToken: taskToken}
}

/*****************************************************************************/

// Used by workers to tell the service that the ActivityTask identified by the taskToken has failed with reason (if specified).
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondActivityTaskFailed.html]
type RespondActivityTaskFailedRequest struct {
	Details   string `json:"details,omitempty"`
	Reason    string `json:"reason,omitempty"`
	TaskToken string `json:"taskToken"`
}

// Creates a new RespondActivityTaskFailedRequest.
// (taskToken string) The taskToken of the ActivityTask.
func NewRespondActivityTaskFailedRequest(taskToken string) *RespondActivityTaskFailedRequest {
	return &RespondActivityTaskFailedRequest{TaskToken: taskToken}
}

/*****************************************************************************/

// Used by deciders to tell the service that the DecisionTask identified by the taskToken has successfully completed.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RespondDecisionTaskCompleted.html]
type RespondDecisionTaskCompletedRequest struct {
	Decisions        []Decision `json:"decisions,omitempty"`
	ExecutionContext string     `json:"executionContext,omitempty"`
	TaskToken        string     `json:"taskToken"`
}

// Creates a new RespondDecisionTaskCompletedRequest.
// (taskToken string) The taskToken from the DecisionTask.
func NewRespondDecisionTaskCompletedRequest(taskToken string) *RespondDecisionTaskCompletedRequest {
	return &RespondDecisionTaskCompletedRequest{TaskToken: taskToken}
}

/*****************************************************************************/

// Records a WorkflowExecutionSignaled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_SignalWorkflowExecution.html]
type SignalWorkflowExecutionRequest struct {
	Domain     string `json:"domain"`
	Input      string `json:"input,omitempty"`
	RunId      string `json:"runId,omitempty"`
	SignalName string `json:"signalName"`
	WorkflowId string `json:"workflowId"`
}

// Creates a new SignalWorkflowExecutionRequest.
// (domain string) The name of the domain containing the workflow execution to signal.
// (signalName string) The name of the signal. This name must be meaningful to the target workflow.
// (workflowId string) The workflowId of the workflow execution to signal.
func NewSignalWorkflowExecutionRequest(domain, signalName, workflowId string) *SignalWorkflowExecutionRequest {
	return &SignalWorkflowExecutionRequest{Domain: domain, SignalName: signalName, WorkflowId: workflowId}
}

/*****************************************************************************/

// Starts an execution of the workflow type in the specified domain using the provided workflowId and input data.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartWorkflowExecution.html]
type StartWorkflowExecutionRequest struct {
	ChildPolicy                  ChildPolicy   `json:"childPolicy,omitempty"`
	Domain                       string        `json:"domain"`
	ExecutionStartToCloseTimeout string        `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string        `json:"input,omitempty"`
	TagList                      []string      `json:"tagList,omitempty"`
	TaskList                     *TaskList     `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string        `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowId                   string        `json:"workflowId"`
	WorkflowType                 *WorkflowType `json:"workflowType"`
}

// Creates a new StartWorkflowExecutionRequest.
// (domain string) The name of the domain in which the workflow execution is created.
// (workflowId string) The user defined identifier associated with the workflow execution
// (workflowType WorkflowType) The type of the workflow to start.
func NewStartWorkflowExecutionRequest(domain, workflowId string, workflowType *WorkflowType) *StartWorkflowExecutionRequest {
	return &StartWorkflowExecutionRequest{Domain: domain, WorkflowId: workflowId, WorkflowType: workflowType}
}

/*****************************************************************************/

// Records a WorkflowExecutionTerminated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TerminateWorkflowExecution.html]
type TerminateWorkflowExecutionRequest struct {
	ChildPolicy ChildPolicy `json:"childPolicy,omitempty"`
	Details     string      `json:"details,omitempty"`
	Domain      string      `json:"domain"`
	Reason      string      `json:"reason,omitempty"`
	RunId       string      `json:"runId,omitempty"`
	WorkflowId  string      `json:"workflowId"`
}

// Creates a new TerminateWorkflowExecutionRequest.
// (domain string) The domain of the workflow execution to terminate.
// (workflowId string) The workflowId of the workflow execution to terminate.
func NewTerminateWorkflowExecutionRequest(domain, workflowId string) *TerminateWorkflowExecutionRequest {
	return &TerminateWorkflowExecutionRequest{Domain: domain, WorkflowId: workflowId}
}
