package swf

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/******************************************************************************
 * Constants
 */

/*	The cause of the failure to process the decision.
	ACTIVITY_ID_ALREADY_IN_USE
	ACTIVITY_ID_UNKNOWN
	ACTIVITY_TYPE_DEPRECATED
	ACTIVITY_TYPE_DOES_NOT_EXIST
	ACTIVITY_CREATION_RATE_EXCEEDED
	CHILD_POLICY_APPLIED
	DEFAULT_CHILD_POLICY_UNDEFINED
	DEFAULT_EXECUTION_START_TO_CLOSE_TIMEOUT_UNDEFINED
	DEFAULT_HEARTBEAT_TIMEOUT_UNDEFINED
	DEFAULT_SCHEDULE_TO_START_TIMEOUT_UNDEFINED
	DEFAULT_TASK_LIST_UNDEFINED
	DEFAULT_TASK_START_TO_CLOSE_TIMEOUT_UNDEFINED
	EVENT_LIMIT_EXCEEDED
	OPEN_ACTIVITIES_LIMIT_EXCEEDED
	OPEN_TIMERS_LIMIT_EXCEEDED
	OPERATION_NOT_PERMITTED
	OPERATOR_INITIATED
	TIMER_CREATION_RATE_EXCEEDED
	TIMER_ID_ALREADY_IN_USE
	TIMER_ID_UNKNOWN
	UNHANDLED_DECISION
	WORKFLOW_TYPE_DEPRECATED
	WORKFLOW_TYPE_DOES_NOT_EXIST */
type Cause string

const (
	ACTIVITY_ID_ALREADY_IN_USE                         Cause = "ACTIVITY_ID_ALREADY_IN_USE"
	ACTIVITY_ID_UNKNOWN                                Cause = "ACTIVITY_ID_UNKNOWN"
	ACTIVITY_TYPE_DEPRECATED                           Cause = "ACTIVITY_TYPE_DEPRECATED"
	ACTIVITY_TYPE_DOES_NOT_EXIST                       Cause = "ACTIVITY_TYPE_DOES_NOT_EXIST"
	ACTIVITY_CREATION_RATE_EXCEEDED                    Cause = "ACTIVITY_CREATION_RATE_EXCEEDED"
	CHILD_POLICY_APPLIED                               Cause = "CHILD_POLICY_APPLIED"
	DEFAULT_CHILD_POLICY_UNDEFINED                     Cause = "DEFAULT_CHILD_POLICY_UNDEFINED"
	DEFAULT_EXECUTION_START_TO_CLOSE_TIMEOUT_UNDEFINED Cause = "DEFAULT_EXECUTION_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	DEFAULT_HEARTBEAT_TIMEOUT_UNDEFINED                Cause = "DEFAULT_HEARTBEAT_TIMEOUT_UNDEFINED"
	DEFAULT_SCHEDULE_TO_START_TIMEOUT_UNDEFINED        Cause = "DEFAULT_SCHEDULE_TO_START_TIMEOUT_UNDEFINED"
	DEFAULT_TASK_LIST_UNDEFINED                        Cause = "DEFAULT_TASK_LIST_UNDEFINED"
	DEFAULT_TASK_START_TO_CLOSE_TIMEOUT_UNDEFINED      Cause = "DEFAULT_TASK_START_TO_CLOSE_TIMEOUT_UNDEFINED"
	EVENT_LIMIT_EXCEEDED                               Cause = "EVENT_LIMIT_EXCEEDED"
	OPEN_ACTIVITIES_LIMIT_EXCEEDED                     Cause = "OPEN_ACTIVITIES_LIMIT_EXCEEDED"
	OPEN_TIMERS_LIMIT_EXCEEDED                         Cause = "OPEN_TIMERS_LIMIT_EXCEEDED"
	OPERATION_NOT_PERMITTED                            Cause = "OPERATION_NOT_PERMITTED"
	OPERATOR_INITIATED                                 Cause = "OPERATOR_INITIATED"
	TIMER_CREATION_RATE_EXCEEDED                       Cause = "TIMER_CREATION_RATE_EXCEEDED"
	TIMER_ID_ALREADY_IN_USE                            Cause = "TIMER_ID_ALREADY_IN_USE"
	TIMER_ID_UNKNOWN                                   Cause = "TIMER_ID_UNKNOWN"
	UNHANDLED_DECISION                                 Cause = "UNHANDLED_DECISION"
	WORKFLOW_TYPE_DEPRECATED                           Cause = "WORKFLOW_TYPE_DEPRECATED"
	WORKFLOW_TYPE_DOES_NOT_EXIST                       Cause = "WORKFLOW_TYPE_DOES_NOT_EXIST"
)

/*	If set, specifies the policy to use for the child workflow executions of the new execution
	if it is terminated by calling the TerminateWorkflowExecution action explicitly or due to
	an expired timeout.
	ABANDON: no action will be taken. The child executions will continue to run.
	REQUEST_CANCEL: a request to cancel will be attempted for each child execution by recording a WorkflowExecutionCancelRequested event in its history. It is up to the decider to take appropriate actions when it receives an execution history with this event.
	TERMINATE: the child executions will be terminated.
	[http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ContinueAsNewWorkflowExecutionDecisionAttributes.html] */
type ChildPolicy string

const (
	ABANDON        ChildPolicy = "ABANDON"
	REQUEST_CANCEL ChildPolicy = "REQUEST_CANCEL"
	TERMINATE      ChildPolicy = "TERMINATE"
)

/*	Specifies the type of the decision.
	CANCEL_TIMER = "CancelTimer"
	CANCEL_WORKFLOW_EXECUTION = "CancelWorkflowExecution"
	COMPLETE_WORKFLOW_EXECUTION = "CompleteWorkflowExecution"
	CONTINUE_AS_NEW_WORKFLOW_EXECUTION = "ContinueAsNewWorkflowExecution"
	FAIL_WORKFLOW_EXECUTION = "FailWorkflowExecution"
	RECORD_MARKER = "RecordMarker"
	REQUEST_CANCEL_ACTIVITY_TASK = "RequestCancelActivityTask"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION = "RequestCancelExternalWorkflowExecution"
	SCHEDULE_ACTIVITY_TASK = "ScheduleActivityTask"
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION = "SignalExternalWorkflowExecution"
	START_CHILD_WORKFLOW_EXECUTION = "StartChildWorkflowExecution"
	START_TIMER = "StartTimer" */
type DecisionType string

const (
	CANCEL_TIMER                               DecisionType = "CancelTimer"
	CANCEL_WORKFLOW_EXECUTION                  DecisionType = "CancelWorkflowExecution"
	COMPLETE_WORKFLOW_EXECUTION                DecisionType = "CompleteWorkflowExecution"
	CONTINUE_AS_NEW_WORKFLOW_EXECUTION         DecisionType = "ContinueAsNewWorkflowExecution"
	FAIL_WORKFLOW_EXECUTION                    DecisionType = "FailWorkflowExecution"
	RECORD_MARKER                              DecisionType = "RecordMarker"
	REQUEST_CANCEL_ACTIVITY_TASK               DecisionType = "RequestCancelActivityTask"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION DecisionType = "RequestCancelExternalWorkflowExecution"
	SCHEDULE_ACTIVITY_TASK                     DecisionType = "ScheduleActivityTask"
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION         DecisionType = "SignalExternalWorkflowExecution"
	START_CHILD_WORKFLOW_EXECUTION             DecisionType = "StartChildWorkflowExecution"
	START_TIMER                                DecisionType = "StartTimer"
)

/*	The type of the history event.
	ACTIVITY_TASK_CANCELED = "ActivityTaskCanceled"
	ACTIVITY_TASK_COMPLETED = "ActivityTaskCompleted"
	ACTIVITY_TASK_FAILED = "ActivityTaskFailed"
	ACTIVITY_TASK_REQUESTED = "ActivityTaskCancelRequested"
	ACTIVITY_TASK_SCHEDULED = "ActivityTaskScheduled"
	ACTIVITY_TASK_STARTED = "ActivityTaskStarted"
	ACTIVITY_TASK_TIMED_OUT = "ActivityTaskTimedOut"

	CHILD_WORKFLOW_EXECUTION_CANCELED = "ChildWorkflowExecutionCanceled"
	CHILD_WORKFLOW_EXECUTION_COMPLETED = "ChildWorkflowExecutionCompleted"
	CHILD_WORKFLOW_EXECUTION_FAILED = "ChildWorkflowExecutionFailed"
	CHILD_WORKFLOW_EXECUTION_STARTED = "ChildWorkflowExecutionStarted"
	CHILD_WORKFLOW_EXECUTION_TERMINATED = "ChildWorkflowExecutionTerminated"
	CHILD_WORKFLOW_EXECUTION_TIMED_OUT = "ChildWorkflowExecutionTimedOut"

	COMPLETE_WORKFLOW_EXECUTION_FAILED = "CompleteWorkflowExecutionFailed"

	CONTINUED_AS_NEW_WORKFLOW_EXECUTION_FAILED = "ContinueAsNewWorkflowExecutionFailed"

	CANCEL_TIMER_FAILED = "CancelTimerFailed"
	CANCEL_WORKFLOW_EXECUTION_FAILED = "CancelWorkflowExecutionFailed"

	EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED = "ExternalWorkflowExecutionCancelRequested"
	EXTERNAL_WORKFLOW_EXECUTION_SIGNALED = "ExternalWorkflowExecutionSignaled"

	FAIL_WORKFLOW_EXECUTION_FAILED = "FailWorkflowExecutionFailed"

	DECISION_TASK_COMPLETED = "DecisionTaskCompleted"
	DECISION_TASK_SCHEDULED = "DecisionTaskScheduled"
	DECISION_TASK_STARTED = "DecisionTaskStarted"
	DECISION_TASK_TIMED_OUT = "DecisionTaskTimedOut"

	MARKER_RECORDED = "MarkerRecorded"

	RECORD_MARKER_FAILED = "RecordMarkerFailed"

	REQUEST_CANCEL_ACTIVITY_TASK_FAILED = "RequestCancelActivityTaskFailed"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED = "RequestCancelExternalWorkflowExecutionFailed"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED = "RequestCancelExternalWorkflowExecutionInitiated"

	SCHEDULE_ACTIVITY_TASK_FAILED = "ScheduleActivityTaskFailed"

	START_CHILD_WORKFLOW_EXECUTION_FAILED = "StartChildWorkflowExecutionFailed"
	START_CHILD_WORKFLOW_EXECUTION_INITIATED = "StartChildWorkflowExecutionInitiated"
	START_TIMER_FAILED = "StartTimerFailed"

	SINGAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED = "SignalExternalWorkflowExecutionInitiated"
	SINGAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED = "SignalExternalWorkflowExecutionFailed"

	TIMER_CANCELED = "TimerCanceled"
	TIMER_FIRED = "TimerFired"
	TIMER_STARTED = "TimerStarted"

	WORKFLOW_EXECUTION_CANCEL_REQUESTED = "WorkflowExecutionCancelRequested"
	WORKFLOW_EXECUTION_CANCELED = "WorkflowExecutionCanceled"
	WORKFLOW_EXECUTION_COMPLETED = "WorkflowExecutionCompleted"
	WORKFLOW_EXECUTION_CONTINUED_AS_NEW = "WorkflowExecutionContinuedAsNew"
	WORKFLOW_EXECUTION_FAILED = "WorkflowExecutionFailed"
	WORKFLOW_EXECUTION_SIGNALED = "WorkflowExecutionSignaled"
	WORKFLOW_EXECUTION_STARTED = "WorkflowExecutionStarted"
	WORKFLOW_EXECUTION_TERMINATED = "WorkflowExecutionTerminated"
	WORKFLOW_EXECUTION_TIMED_OUT = "WorkflowExecutionTimedOut" */
type EventType string

const (
	ACTIVITY_TASK_CANCELED  EventType = "ActivityTaskCanceled"
	ACTIVITY_TASK_COMPLETED EventType = "ActivityTaskCompleted"
	ACTIVITY_TASK_FAILED    EventType = "ActivityTaskFailed"
	ACTIVITY_TASK_REQUESTED EventType = "ActivityTaskCancelRequested"
	ACTIVITY_TASK_SCHEDULED EventType = "ActivityTaskScheduled"
	ACTIVITY_TASK_STARTED   EventType = "ActivityTaskStarted"
	ACTIVITY_TASK_TIMED_OUT EventType = "ActivityTaskTimedOut"

	CHILD_WORKFLOW_EXECUTION_CANCELED   EventType = "ChildWorkflowExecutionCanceled"
	CHILD_WORKFLOW_EXECUTION_COMPLETED  EventType = "ChildWorkflowExecutionCompleted"
	CHILD_WORKFLOW_EXECUTION_FAILED     EventType = "ChildWorkflowExecutionFailed"
	CHILD_WORKFLOW_EXECUTION_STARTED    EventType = "ChildWorkflowExecutionStarted"
	CHILD_WORKFLOW_EXECUTION_TERMINATED EventType = "ChildWorkflowExecutionTerminated"
	CHILD_WORKFLOW_EXECUTION_TIMED_OUT  EventType = "ChildWorkflowExecutionTimedOut"

	COMPLETE_WORKFLOW_EXECUTION_FAILED EventType = "CompleteWorkflowExecutionFailed"

	CONTINUED_AS_NEW_WORKFLOW_EXECUTION_FAILED EventType = "ContinueAsNewWorkflowExecutionFailed"

	CANCEL_TIMER_FAILED              EventType = "CancelTimerFailed"
	CANCEL_WORKFLOW_EXECUTION_FAILED EventType = "CancelWorkflowExecutionFailed"

	EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = "ExternalWorkflowExecutionCancelRequested"
	EXTERNAL_WORKFLOW_EXECUTION_SIGNALED         EventType = "ExternalWorkflowExecutionSignaled"

	FAIL_WORKFLOW_EXECUTION_FAILED EventType = "FailWorkflowExecutionFailed"

	DECISION_TASK_COMPLETED EventType = "DecisionTaskCompleted"
	DECISION_TASK_SCHEDULED EventType = "DecisionTaskScheduled"
	DECISION_TASK_STARTED   EventType = "DecisionTaskStarted"
	DECISION_TASK_TIMED_OUT EventType = "DecisionTaskTimedOut"

	MARKER_RECORDED EventType = "MarkerRecorded"

	RECORD_MARKER_FAILED EventType = "RecordMarkerFailed"

	REQUEST_CANCEL_ACTIVITY_TASK_FAILED                  EventType = "RequestCancelActivityTaskFailed"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED    EventType = "RequestCancelExternalWorkflowExecutionFailed"
	REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = "RequestCancelExternalWorkflowExecutionInitiated"

	SCHEDULE_ACTIVITY_TASK_FAILED EventType = "ScheduleActivityTaskFailed"

	START_CHILD_WORKFLOW_EXECUTION_FAILED    EventType = "StartChildWorkflowExecutionFailed"
	START_CHILD_WORKFLOW_EXECUTION_INITIATED EventType = "StartChildWorkflowExecutionInitiated"
	START_TIMER_FAILED                       EventType = "StartTimerFailed"

	SINGAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = "SignalExternalWorkflowExecutionInitiated"
	SINGAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED    EventType = "SignalExternalWorkflowExecutionFailed"

	TIMER_CANCELED EventType = "TimerCanceled"
	TIMER_FIRED    EventType = "TimerFired"
	TIMER_STARTED  EventType = "TimerStarted"

	WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = "WorkflowExecutionCancelRequested"
	WORKFLOW_EXECUTION_CANCELED         EventType = "WorkflowExecutionCanceled"
	WORKFLOW_EXECUTION_COMPLETED        EventType = "WorkflowExecutionCompleted"
	WORKFLOW_EXECUTION_CONTINUED_AS_NEW EventType = "WorkflowExecutionContinuedAsNew"
	WORKFLOW_EXECUTION_FAILED           EventType = "WorkflowExecutionFailed"
	WORKFLOW_EXECUTION_SIGNALED         EventType = "WorkflowExecutionSignaled"
	WORKFLOW_EXECUTION_STARTED          EventType = "WorkflowExecutionStarted"
	WORKFLOW_EXECUTION_TERMINATED       EventType = "WorkflowExecutionTerminated"
	WORKFLOW_EXECUTION_TIMED_OUT        EventType = "WorkflowExecutionTimedOut"
)

/*	The current status of the execution.
	CLOSED
	OPEN */
type ExecutionStatus string

const (
	CLOSED ExecutionStatus = "CLOSED"
	OPEN   ExecutionStatus = "OPEN"
)

/*	The current status of the activity type.
	CANCELED
	COMPLETED
	CONTINUED_AS_NEW
	DEPRECATED
	FAILED
	REGISTERED
	TERMINATED
	TIMED_OUT */
type Status string

const (
	CANCELED         Status = "CANCELED"
	COMPLETED        Status = "COMPLETED"
	CONTINUED_AS_NEW Status = "CONTINUED_AS_NEW"
	DEPRECATED       Status = "DEPRECATED"
	FAILED           Status = "FAILED"
	REGISTERED       Status = "REGISTERED"
	TERMINATED       Status = "TERMINATED"
	TIMED_OUT        Status = "TIMED_OUT"
)

/*	The type of the timeout that caused this event:
	HEARTBEAT
	SCHEDULE_TO_CLOSE
	SCHEDULE_TO_START
	START_TO_CLOSE */
type TimeoutType string

const (
	HEARTBEAT         TimeoutType = "HEARTBEAT"
	SCHEDULE_TO_CLOSE TimeoutType = "SCHEDULE_TO_CLOSE"
	SCHEDULE_TO_START TimeoutType = "SCHEDULE_TO_START"
	START_TO_CLOSE    TimeoutType = "START_TO_CLOSE"
)

/******************************************************************************
 * Date Types
 */

// Unit of work sent to an activity worker.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTask.html]
type ActivityTask struct {
	ActivityId        string             `json:"activityId"`
	ActivityType      *ActivityType      `json:activityType"`
	Input             string             `json:"input,omitempty"`
	StartedEventId    int64              `json:"startedEventId"`
	TaskToken         string             `json:"taskToken"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
}

// Provides details of the ActivityTaskCancelRequested event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskCancelRequestedEventAttributes.html]
type ActivityTaskCancelRequestedEventAttributes struct {
	ActivityId                   string `json:"activityId"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
}

// Provides details of the ActivityTaskCanceled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskCanceledEventAttributes.html]
type ActivityTaskCanceledEventAttributes struct {
	Details                      string `json:"details,omitempty"`
	LatestCancelRequestedEventId int64  `json:"latestCancelRequestedEventId,omitempty"`
	ScheduledEventId             int64  `json:"scheduledEventId"`
	StartedEventId               int64  `json:"startedEventId"`
}

// Provides details of the ActivityTaskCompleted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskCompletedEventAttributes.html]
type ActivityTaskCompletedEventAttributes struct {
	Result           string `json:"result,omitempty"`
	ScheduledEventId int64  `json:"scheduledEventId"`
	StartedEventId   int64  `json:"startedEventId"`
}

// Provides details of the ActivityTaskFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskFailedEventAttributes.html]
type ActivityTaskFailedEventAttributes struct {
	Details          string `json:"details,omitempty"`
	Reason           string `json:"reason,omitempty"`
	ScheduledEventId int64  `json:"scheduledEventId"`
	StartedEventId   int64  `json:"startedEventId"`
}

// Provides details of the ActivityTaskScheduled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskScheduledEventAttributes.html]
type ActivityTaskScheduledEventAttributes struct {
	ActivityId                   string        `json:"activityId"`
	ActivityType                 *ActivityType `json:"activityType"`
	Control                      string        `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64         `json:"decisionTaskCompletedEventId"`
	HeartbeatTimeout             string        `json:"heartbeatTimeout,omitempty"`
	Input                        string        `json:"input,omitempty"`
	ScheduleToCloseTimeout       string        `json:"scheduleToCloseTimeout,omitempty"`
	ScheduleToStartTimeout       string        `json:"scheduleToStartTimeout,omitempty"`
	StartToCloseTimeout          string        `json:"startToCloseTimeout,omitempty"`
	TaskList                     *TaskList     `json:"taskList"`
}

// Provides details of the ActivityTaskStarted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskStartedEventAttributes.html]
type ActivityTaskStartedEventAttributes struct {
	Identity         string `json:"identity,omitempty"`
	ScheduledEventId int64  `json:"scheduledEventId"`
}

// Status information about an activity task.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskStatus.html]
type ActivityTaskStatus struct {
	CancelRequested bool `json:"cancelRequested"`
}

// Provides details of the ActivityTaskTimedOut event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTaskTimedOutEventAttributes.html]
type ActivityTaskTimedOutEventAttributes struct {
	Details          string      `json:"details,omitempty"`
	ScheduledEventId int64       `json:"scheduledEventId"`
	StartedEventId   int64       `json:"startedEventId"`
	TimeoutType      TimeoutType `json:"timeoutType"`
}

// Represents an activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityType.html]
type ActivityType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Configuration settings registered with the activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTypeConfiguration.html]
type ActivityTypeConfiguration struct {
	DefaultTaskHeartbeatTimeout       string    `json:"defaultTaskHeartbeatTimeout,omitempty"`
	DefaultTaskList                   *TaskList `json:"defaultTaskList,omitempty"`
	DefaultTaskScheduleToCloseTimeout string    `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	DefaultTaskStartToCloseTimeout    string    `json:"defaultTaskStartToCloseTimeout,omitempty"`
}

// Detailed information about an activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTypeDetail.html]
type ActivityTypeDetail struct {
	Configuration *ActivityTypeConfiguration `json:"configuration"`
	TypeInfo      *ActivityTypeInfo          `json:"typeInfo"`
}

// Detailed information about an activity type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTypeInfo.html]
type ActivityTypeInfo struct {
	ActivityType    *ActivityType      `json:"configuration"`
	CreationDate    *datetime.JsonDate `json:"configuration"`
	DeprecationDate *datetime.JsonDate `json:"configuration,omitempty"`
	Description     string             `json:"configuration,omitempty"`
	Status          Status             `json:"configuration"`
}

// Contains a paginated list of activity type information structures.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ActivityTypeInfos.html]
type ActivityTypeInfos struct {
	NextPageToken string            `json:"nextPageToken,omitempty"`
	TypeInfos     *ActivityTypeInfo `json:"typeInfos"`
}

// Provides details of the CancelTimer decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CancelTimerDecisionAttributes.html]
type CancelTimerDecisionAttributes struct {
	TimerId string `json:"timerId"`
}

// Provides details of the CancelTimerFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CancelTimerFailedEventAttributes.html]
type CancelTimerFailedEventAttributes struct {
	Cause                        Cause  `json:"cause"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	TimerId                      string `json:"timerId"`
}

// Provides details of the CancelWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CancelWorkflowExecutionDecisionAttributes.html]
type CancelWorkflowExecutionDecisionAttributes struct {
	Details string `json:""details,omitempty`
}

// Provides details of the CancelWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CancelWorkflowExecutionFailedEventAttributes.html]
type CancelWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause `json:"cause"`
	DecisionTaskCompletedEventId int64 `json:"decisionTaskCompletedEventId"`
}

// Provide details of the ChildWorkflowExecutionCanceled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionCanceledEventAttributes.html]
type ChildWorkflowExecutionCanceledEventAttributes struct {
	Details           string             `json:"details,omitempty"`
	InitiatedEventId  int64              `json:"initiatedEventId"`
	StartedEventId    int64              `json:"startedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Provides details of the ChildWorkflowExecutionCompleted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionCompletedEventAttributes.html]
type ChildWorkflowExecutionCompletedEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	Result            string             `json:"result,omitempty"`
	StartedEventId    int64              `json:"startedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Provides details of the ChildWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionFailedEventAttributes.html]
type ChildWorkflowExecutionFailedEventAttributes struct {
	Details           string             `json:"details,omitempty"`
	InitiatedEventId  int64              `json:"initiatedEventId"`
	Reason            string             `json:"reason,omitempty"`
	StartedEventId    int64              `json:"startedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Provides details of the ChildWorkflowExecutionStarted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionStartedEventAttributes.html]
type ChildWorkflowExecutionStartedEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Provides details of the ChildWorkflowExecutionTerminated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionTerminatedEventAttributes.html]
type ChildWorkflowExecutionTerminatedEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	StartedEventId    int64              `json:"startedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Provides details of the ChildWorkflowExecutionTimedOut event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ChildWorkflowExecutionTimedOutEventAttributes.html]
type ChildWorkflowExecutionTimedOutEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	StartedEventId    int64              `json:"startedEventId"`
	TimeoutType       TimeoutType        `json:"timeoutType"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
	WorkflowType      *WorkflowType      `json:"workflowType"`
}

// Used to filter the closed workflow executions in visibility APIs by their close status.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CloseStatusFilter.html]
type CloseStatusFilter struct {
	Status string `json:"status"`
}

// Provides details of the CompleteWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CompleteWorkflowExecutionDecisionAttributes.html]
type CompleteWorkflowExecutionDecisionAttributes struct {
	Result string `json:"result,omitempty"`
}

// Provides details of the CompleteWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_CompleteWorkflowExecutionFailedEventAttributes.html]
type CompleteWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause `json:"cause"`
	DecisionTaskCompletedEventId int64 `json:"decisionTaskCompletedEventId"`
}

// Provides details of the ContinueAsNewWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ContinueAsNewWorkflowExecutionDecisionAttributes.html]
type ContinueAsNewWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  ChildPolicy `json:"childPolicy,omitempty"`
	ExecutionStartToCloseTimeout string      `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string      `json:"input,omitempty"`
	TagList                      []string    `json:"tagList,omitempty"`
	TaskList                     *TaskList   `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string      `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowTypeVersion          string      `json:"workflowTypeVersion,omitempty"`
}

// Provides details of the ContinueAsNewWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ContinueAsNewWorkflowExecutionFailedEventAttributes.html]
type ContinueAsNewWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause `json:"cause"`
	DecisionTaskCompletedEventId int64 `json:"decisionTaskCompletedEventId"`
}

// Specifies a decision made by the decider.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_Decision.html]
type Decision struct {
	CancelTimerDecisionAttributes                            *CancelTimerDecisionAttributes                            `json:"cancelTimerDecisionAttributes,omitempty"`
	CancelWorkflowExecutionDecisionAttributes                *CancelWorkflowExecutionDecisionAttributes                `json:"cancelWorkflowExecutionDecisionAttributes,omitempty"`
	CompleteWorkflowExecutionDecisionAttributes              *CompleteWorkflowExecutionDecisionAttributes              `json:"completeWorkflowExecutionDecisionAttributes,omitempty"`
	ContinueAsNewWorkflowExecutionDecisionAttributes         *ContinueAsNewWorkflowExecutionDecisionAttributes         `json:"continueAsNewWorkflowExecutionDecisionAttributes,omitempty"`
	DecisionType                                             DecisionType                                              `json:"decisionType"`
	FailWorkflowExecutionDecisionAttributes                  *FailWorkflowExecutionDecisionAttributes                  `json:"failWorkflowExecutionDecisionAttributes,omitempty"`
	RecordMarkerDecisionAttributes                           *RecordMarkerDecisionAttributes                           `json:"recordMarkerDecisionAttributes,omitempty"`
	RequestCancelActivityTaskDecisionAttributes              *RequestCancelActivityTaskDecisionAttributes              `json:"requestCancelActivityTaskDecisionAttributes,omitempty"`
	RequestCancelExternalWorkflowExecutionDecisionAttributes *RequestCancelExternalWorkflowExecutionDecisionAttributes `json:"requestCancelExternalWorkflowExecutionDecisionAttributes,omitempty"`
	ScheduleActivityTaskDecisionAttributes                   *ScheduleActivityTaskDecisionAttributes                   `json:"scheduleActivityTaskDecisionAttributes,omitempty"`
	SignalExternalWorkflowExecutionDecisionAttributes        *SignalExternalWorkflowExecutionDecisionAttributes        `json:"signalExternalWorkflowExecutionDecisionAttributes,omitempty"`
	StartChildWorkflowExecutionDecisionAttributes            *StartChildWorkflowExecutionDecisionAttributes            `json:"startChildWorkflowExecutionDecisionAttributes,omitempty"`
	StartTimerDecisionAttributes                             *StartTimerDecisionAttributes                             `json:"startTimerDecisionAttributes,omitempty"`
}

// A structure that represents a decision task. Decision tasks are sent to deciders in order for them to make decisions.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DecisionTask.html]
type DecisionTask struct {
	Events                 []HistoryEvent     `json:"events"`
	NextPageToken          string             `json:"nextPageToken,omitempty"`
	PreviousStartedEventId int64              `json:"previousStartedEventId,omitempty"`
	StartedEventId         int64              `json:"startedEventId"`
	TaskToken              string             `json:"taskToken"`
	WorkflowExecution      *WorkflowExecution `json:"workflowExecution"`
	WorkflowType           *WorkflowType      `json:"workflowType"`
}

// Provides details of the DecisionTaskCompleted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DecisionTaskCompletedEventAttributes.html]
type DecisionTaskCompletedEventAttributes struct {
	ExecutionContext string `json:"executionContext,omitempty"`
	ScheduledEventId int64  `json:"scheduledEventId"`
	StartedEventId   int64  `json:"startedEventId"`
}

// Provides details of the DecisionTaskScheduled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DecisionTaskScheduledEventAttributes.html]
type DecisionTaskScheduledEventAttributes struct {
	StartToCloseTimeout string    `json:"startToCloseTimeout"`
	TaskList            *TaskList `json:"TaskList"`
}

// Provides details of the DecisionTaskStarted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DecisionTaskStartedEventAttributes.html]
type DecisionTaskStartedEventAttributes struct {
	Identity         string `json:"identity,omitempty"`
	ScheduledEventId int64  `json:"scheduledEventId"`
}

// Provides details of the DecisionTaskTimedOut event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DecisionTaskTimedOutEventAttributes.html]
type DecisionTaskTimedOutEventAttributes struct {
	ScheduledEventId int64       `json:"scheduledEventId"`
	StartedEventId   int64       `json:"startedEventId"`
	TimeoutType      TimeoutType `json:"timeoutType"`
}

// Contains the configuration settings of a domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DomainConfiguration.html]
type DomainConfiguration struct {
	WorkflowExecutionRetentionPeriodInDays string `json:" workflowExecutionRetentionPeriodInDays"`
}

// Contains details of a domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DomainDetail.html]
type DomainDetail struct {
	Configuration *DomainConfiguration `json:"configuration"`
	DomainInfo    *DomainInfo          `json:"domainInfo"`
}

// Contains general information about a domain.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DomainInfo.html]
type DomainInfo struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Status      Status `json:"status"`
}

// Contains a paginated collection of DomainInfo structures.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_DomainInfos.html]
type DomainInfos struct {
	DomainInfos   []DomainInfo `json:"domainInfos"`
	NextPageToken string       `json:"nextPageToken,omitempty"`
}

// Used to filter the workflow executions in visibility APIs by various time-based rules.
// Each parameter, if specified, defines a rule that must be satisfied by each returned query result.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ExecutionTimeFilter.html]
type ExecutionTimeFilter struct {
	LatestDate *datetime.JsonDate `json:"latestDate,omitempty"`
	OldestDate *datetime.JsonDate `json:"oldestDate"`
}

// Provides details of the ExternalWorkflowExecutionCancelRequested event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ExternalWorkflowExecutionCancelRequestedEventAttributes.html]
type ExternalWorkflowExecutionCancelRequestedEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
}

// Provides details of the ExternalWorkflowExecutionSignaled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ExternalWorkflowExecutionSignaledEventAttributes.html]
type ExternalWorkflowExecutionSignaledEventAttributes struct {
	InitiatedEventId  int64              `json:"initiatedEventId"`
	WorkflowExecution *WorkflowExecution `json:"workflowExecution"`
}

// Provides details of the FailWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_FailWorkflowExecutionDecisionAttributes.html]
type FailWorkflowExecutionDecisionAttributes struct {
	Details string `json:"details,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// Provides details of the FailWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_FailWorkflowExecutionFailedEventAttributes.html]
type FailWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause `json:"cause"`
	DecisionTaskCompletedEventId int64 `json:"decisionTaskCompletedEventId"`
}

// Paginated representation of a workflow history for a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_History.html]
type History struct {
	Events        []HistoryEvent `json:"events"`
	NextPageToken string         `json:"nextPageToken,omitempty"`
}

// Event within a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_HistoryEvent.html]
type HistoryEvent struct {
	ActivityTaskCancelRequestedEventAttributes                     *ActivityTaskCancelRequestedEventAttributes                     `json:"activityTaskCancelRequestedEventAttributes,omitempty"`
	ActivityTaskCanceledEventAttributes                            *ActivityTaskCanceledEventAttributes                            `json:"activityTaskCanceledEventAttributes,omitempty"`
	ActivityTaskCompletedEventAttributes                           *ActivityTaskCompletedEventAttributes                           `json:"activityTaskCompletedEventAttributes,omitempty"`
	ActivityTaskFailedEventAttributes                              *ActivityTaskFailedEventAttributes                              `json:"activityTaskFailedEventAttributes,omitempty"`
	ActivityTaskScheduledEventAttributes                           *ActivityTaskScheduledEventAttributes                           `json:"activityTaskScheduledEventAttributes,omitempty"`
	ActivityTaskStartedEventAttributes                             *ActivityTaskStartedEventAttributes                             `json:"activityTaskStartedEventAttributes,omitempty"`
	ActivityTaskTimedOutEventAttributes                            *ActivityTaskTimedOutEventAttributes                            `json:"activityTaskTimedOutEventAttributes,omitempty"`
	CancelTimerFailedEventAttributes                               *CancelTimerFailedEventAttributes                               `json:"cancelTimerFailedEventAttributes,omitempty"`
	CancelWorkflowExecutionFailedEventAttributes                   *CancelWorkflowExecutionFailedEventAttributes                   `json:"cancelWorkflowExecutionFailedEventAttributes,omitempty"`
	ChildWorkflowExecutionCanceledEventAttributes                  *ChildWorkflowExecutionCanceledEventAttributes                  `json:"childWorkflowExecutionCanceledEventAttributes,omitempty"`
	ChildWorkflowExecutionCompletedEventAttributes                 *ChildWorkflowExecutionCompletedEventAttributes                 `json:"childWorkflowExecutionCompletedEventAttributes,omitempty"`
	ChildWorkflowExecutionFailedEventAttributes                    *ChildWorkflowExecutionFailedEventAttributes                    `json:"childWorkflowExecutionFailedEventAttributes,omitempty"`
	ChildWorkflowExecutionStartedEventAttributes                   *ChildWorkflowExecutionStartedEventAttributes                   `json:"childWorkflowExecutionStartedEventAttributes,omitempty"`
	ChildWorkflowExecutionTerminatedEventAttributes                *ChildWorkflowExecutionTerminatedEventAttributes                `json:"childWorkflowExecutionTerminatedEventAttributes,omitempty"`
	ChildWorkflowExecutionTimedOutEventAttributes                  *ChildWorkflowExecutionTimedOutEventAttributes                  `json:"childWorkflowExecutionTimedOutEventAttributes,omitempty"`
	CompleteWorkflowExecutionFailedEventAttributes                 *CompleteWorkflowExecutionFailedEventAttributes                 `json:"completeWorkflowExecutionFailedEventAttributes,omitempty"`
	ContinueAsNewWorkflowExecutionFailedEventAttributes            *ContinueAsNewWorkflowExecutionFailedEventAttributes            `json:"continueAsNewWorkflowExecutionFailedEventAttributes,omitempty"`
	DecisionTaskCompletedEventAttributes                           *DecisionTaskCompletedEventAttributes                           `json:"decisionTaskCompletedEventAttributes,omitempty"`
	DecisionTaskScheduledEventAttributes                           *DecisionTaskScheduledEventAttributes                           `json:"decisionTaskScheduledEventAttributes,omitempty"`
	DecisionTaskStartedEventAttributes                             *DecisionTaskStartedEventAttributes                             `json:"decisionTaskStartedEventAttributes,omitempty"`
	DecisionTaskTimedOutEventAttributes                            *DecisionTaskTimedOutEventAttributes                            `json:"decisionTaskTimedOutEventAttributes,omitempty"`
	EventId                                                        int64                                                           `json:"eventId"`
	EventTimestamp                                                 *datetime.JsonDate                                              `json:"eventTimestamp"`
	EventType                                                      EventType                                                       `json:"eventType"`
	ExternalWorkflowExecutionCancelRequestedEventAttributes        *ExternalWorkflowExecutionCancelRequestedEventAttributes        `json:"externalWorkflowExecutionCancelRequestedEventAttributes,omitempty"`
	ExternalWorkflowExecutionSignaledEventAttributes               *ExternalWorkflowExecutionSignaledEventAttributes               `json:"externalWorkflowExecutionSignaledEventAttributes,omitempty"`
	FailWorkflowExecutionFailedEventAttributes                     *FailWorkflowExecutionFailedEventAttributes                     `json:"failWorkflowExecutionFailedEventAttributes,omitempty"`
	MarkerRecordedEventAttributes                                  *MarkerRecordedEventAttributes                                  `json:"markerRecordedEventAttributes,omitempty"`
	RecordMarkerFailedEventAttributes                              *RecordMarkerFailedEventAttributes                              `json:"recordMarkerFailedEventAttributes,omitempty"`
	RequestCancelActivityTaskFailedEventAttributes                 *RequestCancelActivityTaskFailedEventAttributes                 `json:"requestCancelActivityTaskFailedEventAttributes,omitempty"`
	RequestCancelExternalWorkflowExecutionInitiatedEventAttributes *RequestCancelExternalWorkflowExecutionInitiatedEventAttributes `json:"requestCancelExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	ScheduleActivityTaskFailedEventAttributes                      *ScheduleActivityTaskFailedEventAttributes                      `json:"scheduleActivityTaskFailedEventAttributes,omitempty"`
	SignalExternalWorkflowExecutionFailedEventAttributes           *SignalExternalWorkflowExecutionFailedEventAttributes           `json:"signalExternalWorkflowExecutionFailedEventAttributes,omitempty"`
	SignalExternalWorkflowExecutionInitiatedEventAttributes        *SignalExternalWorkflowExecutionInitiatedEventAttributes        `json:"signalExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	StartChildWorkflowExecutionFailedEventAttributes               *StartChildWorkflowExecutionFailedEventAttributes               `json:"startChildWorkflowExecutionFailedEventAttributes,omitempty"`
	StartChildWorkflowExecutionInitiatedEventAttributes            *StartChildWorkflowExecutionInitiatedEventAttributes            `json:"startChildWorkflowExecutionInitiatedEventAttributes,omitempty"`
	StartTimerFailedEventAttributes                                *StartTimerFailedEventAttributes                                `json:"startTimerFailedEventAttributes,omitempty"`
	TimerCanceledEventAttributes                                   *TimerCanceledEventAttributes                                   `json:"timerCanceledEventAttributes,omitempty"`
	TimerFiredEventAttributes                                      *TimerFiredEventAttributes                                      `json:"timerFiredEventAttributes,omitempty"`
	TimerStartedEventAttributes                                    *TimerStartedEventAttributes                                    `json:"timerStartedEventAttributes,omitempty"`
	WorkflowExecutionCancelRequestedEventAttributes                *WorkflowExecutionCancelRequestedEventAttributes                `json:"workflowExecutionCancelRequestedEventAttributes,omitempty"`
	WorkflowExecutionCanceledEventAttributes                       *WorkflowExecutionCanceledEventAttributes                       `json:"workflowExecutionCanceledEventAttributes,omitempty"`
	WorkflowExecutionCompletedEventAttributes                      *WorkflowExecutionCompletedEventAttributes                      `json:"workflowExecutionCompletedEventAttributes,omitempty"`
	WorkflowExecutionContinuedAsNewEventAttributes                 *WorkflowExecutionContinuedAsNewEventAttributes                 `json:"workflowExecutionContinuedAsNewEventAttributes,omitempty"`
	WorkflowExecutionFailedEventAttributes                         *WorkflowExecutionFailedEventAttributes                         `json:"workflowExecutionFailedEventAttributes,omitempty"`
	WorkflowExecutionSignaledEventAttributes                       *WorkflowExecutionSignaledEventAttributes                       `json:"workflowExecutionSignaledEventAttributes,omitempty"`
	WorkflowExecutionStartedEventAttributes                        *WorkflowExecutionStartedEventAttributes                        `json:"workflowExecutionStartedEventAttributes,omitempty"`
	WorkflowExecutionTerminatedEventAttributes                     *WorkflowExecutionTerminatedEventAttributes                     `json:"workflowExecutionTerminatedEventAttributes,omitempty"`
	WorkflowExecutionTimedOutEventAttributes                       *WorkflowExecutionTimedOutEventAttributes                       `json:"workflowExecutionTimedOutEventAttributes,omitempty"`
}

// Provides details of the MarkerRecorded event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_MarkerRecordedEventAttributes.html]
type MarkerRecordedEventAttributes struct {
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
	MarkerName                   string `json:"markerName,omitempty"`
}

// Contains the count of tasks in a task list.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_PendingTaskCount.html]
type PendingTaskCount struct {
	Count     int  `json:"count"`
	Truncated bool `json:"truncated,omitempty"`
}

// Provides details of the RecordMarker decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RecordMarkerDecisionAttributes.html]
type RecordMarkerDecisionAttributes struct {
	Details    string `json:"details,omitempty"`
	MarkerName string `json:"markerName,omitempty"`
}

// Provides details of the RecordMarkerFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RecordMarkerFailedEventAttributes.html]
type RecordMarkerFailedEventAttributes struct {
	Cause                        Cause  `json:"cause"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	MarkerName                   string `json:"markerName,omitempty"`
}

// Provides details of the RequestCancelActivityTask decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelActivityTaskDecisionAttributes.html]
type RequestCancelActivityTaskDecisionAttributes struct {
	ActivityId string `json:"activityId"`
}

// Provides details of the RequestCancelActivityTaskFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelActivityTaskFailedEventAttributes.html]
type RequestCancelActivityTaskFailedEventAttributes struct {
	ActivityId                   string `json:"activityId"`
	Cause                        Cause  `json:"cause"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
}

// Provides details of the RequestCancelExternalWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelExternalWorkflowExecutionDecisionAttributes.html]
type RequestCancelExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control,omitempty"`
	RunId      string `json:"runId,omitempty"`
	WorkflowId string `json:"workflowId"`
}

// Provides details of the RequestCancelExternalWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelExternalWorkflowExecutionFailedEventAttributes.html]
type RequestCancelExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause  `json:"cause"`
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	InitiatedEventId             int64  `json:"initiatedEventId"`
	RunId                        string `json:"runId,omitempty"`
	WorkflowId                   string `json:"workflowId"`
}

// Provides details of the RequestCancelExternalWorkflowExecutionInitiated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_RequestCancelExternalWorkflowExecutionInitiatedEventAttributes.html]
type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	RunId                        string `json:"runId,omitempty"`
	WorkflowId                   string `json:"workflowId"`
}

// Specifies the runId of a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_Run.html]
type Run struct {
	RunId string `json:"runId,omitempty"`
}

// Provides details of the ScheduleActivityTask decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ScheduleActivityTaskDecisionAttributes.html]
type ScheduleActivityTaskDecisionAttributes struct {
	ActivityId             string        `json:"activityId,omitempty"`
	ActivityType           *ActivityType `json:"activityType,omitempty"`
	Control                string        `json:"control,omitempty"`
	HeartbeatTimeout       string        `json:"heartbeatTimeout,omitempty"`
	Input                  string        `json:"input,omitempty"`
	ScheduleToCloseTimeout string        `json:"scheduleToCloseTimeout,omitempty"`
	StartToCloseTimeout    string        `json:"startToCloseTimeout,omitempty"`
	TaskList               *TaskList     `json:"taskList,omitempty"`
}

// Provides details of the ScheduleActivityTaskFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_ScheduleActivityTaskFailedEventAttributes.html]
type ScheduleActivityTaskFailedEventAttributes struct {
	ActivityId                   string        `json:"activityId"`
	ActivityType                 *ActivityType `json:"activityType"`
	Cause                        Cause         `json:"cause"`
	DecisionTaskCompletedEventId int64         `json:"decisionTaskCompletedEventId"`
}

// Provides details of the SignalExternalWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_SignalExternalWorkflowExecutionDecisionAttributes.html]
type SignalExternalWorkflowExecutionDecisionAttributes struct {
	Control    string `json:"control,omitempty"`
	Input      string `json:"input,omitempty"`
	RunId      string `json:"runId,omitempty"`
	SignalName string `json:"signalName"`
	WorkflowId string `json:"workflowId"`
}

// Provides details of the SignalExternalWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_SignalExternalWorkflowExecutionFailedEventAttributes.html]
type SignalExternalWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause  `json:"cause"`
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	RunId                        string `json:"runId,omitempty"`
	WorkflowId                   string `json:"workflowId"`
}

// Provides details of the SignalExternalWorkflowExecutionInitiated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_SignalExternalWorkflowExecutionInitiatedEventAttributes.html]
type SignalExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	Input                        string `json:"input,omitempty"`
	RunId                        string `json:"runId,omitempty"`
	SignalName                   string `json:"signalName"`
	WorkflowId                   string `json:"workflowId"`
}

// Provides details of the StartChildWorkflowExecution decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartChildWorkflowExecutionDecisionAttributes.html]
type StartChildWorkflowExecutionDecisionAttributes struct {
	ChildPolicy                  ChildPolicy   `json:"childPolicy,omitempty"`
	Control                      string        `json:"control,omitempty"`
	ExecutionStartToCloseTimeout string        `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string        `json:"input,omitempty"`
	TagList                      string        `json:"tagList,omitempty"`
	TaskList                     *TaskList     `json:"taskList,omitempty"`
	TaskStartToCloseTimeout      string        `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowId                   string        `json:"workflowId"`
	WorkflowType                 *WorkflowType `json:"workflowType"`
}

// Provides details of the StartChildWorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartChildWorkflowExecutionFailedEventAttributes.html]
type StartChildWorkflowExecutionFailedEventAttributes struct {
	Cause                        Cause         `json:"cause"`
	Control                      string        `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64         `json:"decisionTaskCompletedEventId"`
	InitiatedEventId             int64         `json:"initiatedEventId"`
	WorkflowId                   string        `json:"workflowId"`
	WorkflowType                 *WorkflowType `json:"workflowType"`
}

// Provides details of the StartChildWorkflowExecutionInitiated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartChildWorkflowExecutionInitiatedEventAttributes.html]
type StartChildWorkflowExecutionInitiatedEventAttributes struct {
	ChildPolicy                  ChildPolicy   `json:"childPolicy"`
	Control                      string        `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64         `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string        `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string        `json:"input,omitempty"`
	TagList                      string        `json:"tagList,omitempty"`
	TaskList                     *TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string        `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowId                   string        `json:"workflowId"`
	WorkflowType                 *WorkflowType `json:"workflowType"`
}

// Provides details of the StartTimer decision.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartTimerDecisionAttributes.html]
type StartTimerDecisionAttributes struct {
	Control            string `json:"control,omitempty"`
	StartToFireTimeout string `json:"startToFireTimeout"`
	TimerId            string `json:"timerId"`
}

// Provides details of the StartTimerFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_StartTimerFailedEventAttributes.html]
type StartTimerFailedEventAttributes struct {
	Cause                        Cause  `json:"cause"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	TimerId                      string `json:"timerId"`
}

// Used to filter the workflow executions in visibility APIs based on a tag.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TagFilter.html]
type TagFilter struct {
	Tag string `json:"tag"`
}

// Represents a task list.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TaskList.html]
type TaskList struct {
	Name string `json:"name"`
}

// Provides details of the TimerCanceled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TimerCanceledEventAttributes.html]
type TimerCanceledEventAttributes struct {
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	StartedEventId               int64  `json:"startedEventId"`
	TimerId                      string `json:"timerId"`
}

// Provides details of the TimerFired event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TimerFiredEventAttributes.html]
type TimerFiredEventAttributes struct {
	StartedEventId int64  `json:"startedEventId"`
	TimerId        string `json:"timerId"`
}

// Provides details of the TimerStarted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_TimerStartedEventAttributes.html]
type TimerStartedEventAttributes struct {
	Control                      string `json:"control,omitempty"`
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	StartedEventId               int64  `json:"startedEventId"`
	TimerId                      string `json:"timerId"`
}

// Represents a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecution.html]
type WorkflowExecution struct {
	RunID      string `json:"runId"`
	WorkflowId string `json:"workflowId"`
}

// Provides details of the WorkflowExecutionCancelRequested event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionCancelRequestedEventAttributes.html]
type WorkflowExecutionCancelRequestedEventAttributes struct {
	Cause                     Cause              `json:"cause,omitempty"`
	ExternalInitiatedEventId  int64              `json:"externalInitiatedEventId,omitempty"`
	ExternalWorkflowExecution *WorkflowExecution `json:"externalWorkflowExecution,omitempty"`
}

// Provides details of the WorkflowExecutionCanceled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionCanceledEventAttributes.html]
type WorkflowExecutionCanceledEventAttributes struct {
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
}

// Provides details of the WorkflowExecutionCompleted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionCompletedEventAttributes.html]
type WorkflowExecutionCompletedEventAttributes struct {
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	Result                       string `json:"result,omitempty"`
}

// The configuration settings for a workflow execution including timeout values, tasklist etc.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionConfiguration.html]
type WorkflowExecutionConfiguration struct {
	ChildPolicy                  ChildPolicy `json:"childPolicy"`
	ExecutionStartToCloseTimeout string      `json:"executionStartToCloseTimeout"`
	TaskList                     *TaskList   `json:"taskList"`
	TaskStartToCloseTimeout      string      `json:"taskStartToCloseTimeout"`
}

// Provides details of the WorkflowExecutionContinuedAsNew event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionContinuedAsNewEventAttributes.html]
type WorkflowExecutionContinuedAsNewEventAttributes struct {
	ChildPolicy                  ChildPolicy   `json:"childPolicy"`
	DecisionTaskCompletedEventId int64         `json:"decisionTaskCompletedEventId"`
	ExecutionStartToCloseTimeout string        `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string        `json:"input,omitempty"`
	NewExecutionRunId            string        `json:"newExecutionRunId"`
	TagList                      string        `json:"tagList,omitempty"`
	TaskList                     *TaskList     `json:"taskList"`
	TaskStartToCloseTimeout      string        `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowType                 *WorkflowType `json:"workflowType"`
}

// Contains the count of workflow executions returned from CountOpenWorkflowExecutions or CountClosedWorkflowExecutions.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionCount.html]
type WorkflowExecutionCount struct {
	Count     int  `json:"count"`
	Truncated bool `json:"truncated,omitempty"`
}

// Contains details about a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionDetail.html]
type WorkflowExecutionDetail struct {
	ExecutionConfiguration      *WorkflowExecutionConfiguration `json:"executionConfiguration"`
	ExecutionInfo               *WorkflowExecutionInfo          `json:"executionInfo"`
	LatestActivityTaskTimestamp *datetime.JsonDate              `json:"latestActivityTaskTimestamp,omitempty"`
	LatestExecutionContext      string                          `json:"latestExecutionContext,omitempty"`
	OpenCounts                  *WorkflowExecutionOpenCounts    `json:"openCounts"`
}

// Provides details of the WorkflowExecutionFailed event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionFailedEventAttributes.html]
type WorkflowExecutionFailedEventAttributes struct {
	DecisionTaskCompletedEventId int64  `json:"decisionTaskCompletedEventId"`
	Details                      string `json:"details,omitempty"`
	Reason                       string `json:"reason,omitempty"`
}

// Used to filter the workflow executions in visibility APIs by their workflowId.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionFilter.html]
type WorkflowExecutionFilter struct {
	WorkflowId string `json:"workflowId"`
}

// Contains information about a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionInfo.html]
type WorkflowExecutionInfo struct {
	CancelRequest   bool               `json:"cancelRequest,omitempty"`
	CloseStatus     Status             `json:"closeStatus,omitempty"`
	CloseTimestamp  *datetime.JsonDate `json:"closeTimestamp,omitempty"`
	Execution       *WorkflowExecution `json:"execution"`
	ExecutionStatus ExecutionStatus    `json:"executionStatus"`
	Parent          *WorkflowExecution `json:"parent,omitempty"`
	StartTimestamp  *datetime.JsonDate `json:"startTimestamp"`
	TagList         []string           `json:"tagList,omitempty"`
	WorkflowType    *WorkflowType      `json:"workflowType"`
}

// Contains a paginated list of information about workflow executions.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionInfos.html]
type WorkflowExecutionInfos struct {
	ExecutionInfos *WorkflowExecutionInfo `json:"executionInfos"`
	NextPageToken  string                 `json:"nextPageToken,omitempty"`
}

// Contains the counts of open tasks, child workflow executions and timers for a workflow execution.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionOpenCounts.html]
type WorkflowExecutionOpenCounts struct {
	OpenActivityTasks           int `json:"openActivityTasks"`
	OpenChildWorkflowExecutions int `json:"openChildWorkflowExecutions"`
	OpenDecisionTasks           int `json:"openDecisionTasks"`
	OpenTimers                  int `json:"openTimers"`
}

// Provides details of the WorkflowExecutionSignaled event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionSignaledEventAttributes.html]
type WorkflowExecutionSignaledEventAttributes struct {
	ExternalInitiatedEventId  int64              `json:"externalInitiatedEventId,omitempty"`
	ExternalWorkflowExecution *WorkflowExecution `json:"externalWorkflowExecution,omitempty"`
	Input                     string             `json:"input,omitempty"`
	SignalName                string             `json:"signalName"`
}

// Provides details of WorkflowExecutionStarted event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionStartedEventAttributes.html]
type WorkflowExecutionStartedEventAttributes struct {
	ChildPolicy                  ChildPolicy        `json:"childPolicy"`
	ContinuedExecutionRunId      string             `json:"continuedExecutionRunId,omitempty"`
	ExecutionStartToCloseTimeout string             `json:"executionStartToCloseTimeout,omitempty"`
	Input                        string             `json:"input,omitempty"`
	ParentInitiatedEventId       int64              `json:"parentInitiatedEventId,omitempty"`
	ParentWorkflowExecution      *WorkflowExecution `json:"parentWorkflowExecution,omitempty"`
	TagList                      string             `json:"tagList,omitempty"`
	TaskList                     *TaskList          `json:"taskList"`
	TaskStartToCloseTimeout      string             `json:"taskStartToCloseTimeout,omitempty"`
	WorkflowType                 *WorkflowType      `json:"workflowType"`
}

// Provides details of the WorkflowExecutionTerminated event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionTerminatedEventAttributes.html]
type WorkflowExecutionTerminatedEventAttributes struct {
	Cause       Cause       `json:"cause,omitempty"`
	ChildPolicy ChildPolicy `json:"childPolicy"`
	Details     string      `json:"details,omitempty"`
	Reason      string      `json:"reason,omitempty"`
}

// Provides details of the WorkflowExecutionTimedOut event.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowExecutionTimedOutEventAttributes.html]
type WorkflowExecutionTimedOutEventAttributes struct {
	ChildPolicy ChildPolicy `json:"childPolicy"`
	TimeoutType TimeoutType `json:"timeoutType"`
}

// Represents a workflow type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowType.html]
type WorkflowType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// The configuration settings of a workflow type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowTypeConfiguration.html]
type WorkflowTypeConfiguration struct {
	defaultChildPolicy                  ChildPolicy `json:"defaultChildPolicy,omitempty"`
	defaultExecutionStartToCloseTimeout string      `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	defaultTaskList                     *TaskList   `json:"defaultTaskList,omitempty"`
	defaultTaskStartToCloseTimeout      string      `json:"defaultTaskStartToCloseTimeout,omitempty"`
}

// Contains details about a workflow type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowTypeDetail.html]
type WorkflowTypeDetail struct {
	Configuration *WorkflowTypeConfiguration `json:"configuration"`
	TypeInfo      *WorkflowTypeInfo          `json:"typeInfo"`
}

// Used to filter workflow execution query results by type. Each parameter, if specified,
// defines a rule that must be satisfied by each returned result.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowTypeFilter.html]
type WorkflowTypeFilter struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

// Contains information about a workflow type.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowTypeInfo.html]
type WorkflowTypeInfo struct {
	CreationDate    *datetime.JsonDate `json:"creationDate"`
	DeprecationDate *datetime.JsonDate `json:"deprecationDate,omitempty"`
	Description     string             `json:"description,omitempty"`
	Status          Status             `json:"status"`
	WorkflowType    *WorkflowType      `json:"workflowType"`
}

// Contains a paginated list of information structures about workflow types.
// [http://docs.aws.amazon.com/amazonswf/latest/apireference/API_WorkflowTypeInfos.html]
type WorkflowTypeInfos struct {
	NextPageToken string             `json:"nextPageToken,omitempty"`
	TypeInfos     []WorkflowTypeInfo `json:"typeInfos"`
}
