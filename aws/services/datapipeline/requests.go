package datapipeline

import ()

/*****************************************************************************/

// Validates a pipeline and initiates processing.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ActivatePipeline.html]
type ActivatePipelineRequest struct {
	PipelinId string `json:"pipelineId"`
}

// Creates a new ActivatePipelineRequest.
func NewActivatePipelineRequest(pipelineId string) *ActivatePipelineRequest {
	return &ActivatePipelineRequest{pipelineId}
}

/*****************************************************************************/

// Creates a new empty pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_CreatePipeline.html]
type CreatePipelineRequest struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	UniqueId    string `json:"uniqueId"`
}

// Creates a new CreatePipelineRequest.
func NewCreatePipelineRequest(name, uniqueId string) *CreatePipelineRequest {
	return &CreatePipelineRequest{Name: name, UniqueId: uniqueId}
}

/*****************************************************************************/

// Permanently deletes a pipeline, its pipeline definition and its run history.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DeletePipeline.html]
type DeletePipelineRequest struct {
	PipelineId string `json:"pipelineId"`
}

// Creates a new DeletePipelineRequest.
func NewDeletePipelineRequest(pipelineId string) *DeletePipelineRequest {
	return &DeletePipelineRequest{pipelineId}
}

/*****************************************************************************/

// Returns the object definitions for a set of objects associated with the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribeObjects.html]
type DescribeObjectsRequest struct {
	EvaluateExpressions bool     `json:"evaluateExpressions,omitempty"`
	Marker              string   `json:"marker,omitempty"`
	ObjectIds           []string `json:"objectIds"`
	PipelineId          string   `json:"pipelineId"`
}

// Creates a new DescribeObjectsRequest.
func NewDescribeObjectsRequest(pipelineId string, objectIds ...string) *DescribeObjectsRequest {
	return &DescribeObjectsRequest{ObjectIds: objectIds, PipelineId: pipelineId}
}

/*****************************************************************************/

// Retrieve metadata about one or more pipelines.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribePipelines.html]
type DescribePipelinesRequest struct {
	PipelineIds []string `json:"pipelineIds"`
}

// Creates a new DescribePipelinesRequest.
func NewDescribePipelinesRequest(pipelineIds ...string) *DescribePipelinesRequest {
	return &DescribePipelinesRequest{pipelineIds}
}

/*****************************************************************************/

// Evaluates a string in the context of a specified object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_EvaluateExpression.html]
type EvaluateExpressionRequest struct {
	expression string   `json:"expression"`
	ObjectIds  []string `json:"objectIds"`
	PipelineId string   `json:"pipelineId"`
}

// Creates a new EvaluateExpressionRequest.
func NewEvaluateExpressionRequest(pipelineId, expression string, objectIds ...string) *EvaluateExpressionRequest {
	return &EvaluateExpressionRequest{expression, objectIds, pipelineId}
}

/*****************************************************************************/

// Returns the definition of the specified pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinition.html]
type GetPipelineDefinitionRequest struct {
	PipelineId string `json:"pipelineId"`
	Version    string `json:"version,omitempty"`
}

// Creates a new GetPipelineDefinitionRequest.
func NewGetPipelineDefinitionRequest(pipelineId string) *GetPipelineDefinitionRequest {
	return &GetPipelineDefinitionRequest{PipelineId: pipelineId}
}

/*****************************************************************************/

// Returns a list of pipeline identifiers for all active pipelines.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ListPipelines.html]
type ListPipelinesRequest struct {
	Marker string `json:"marker,omitempty"`
}

// Creates a new ListPipelinesRequest.
func NewListPipelinesRequest() *ListPipelinesRequest {
	return &ListPipelinesRequest{}
}

/*****************************************************************************/

// Task runners call this action to receive a task to perform from AWS Data Pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PollForTask.html]
type PollForTaskRequest struct {
	Hostname         string           `json:"hostname,omitempty"`
	InstanceIdentity InstanceIdentity `json:"instanceIdentity,omitempty"`
	WorkerGroup      string           `json:"workerGroup"`
}

// Creates a new PollForTaskRequest.
func NewPollForTaskRequest(workerGroup string) *PollForTaskRequest {
	return &PollForTaskRequest{WorkerGroup: workerGroup}
}

/*****************************************************************************/

// Adds tasks, schedules, and preconditions that control the behavior of the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html]
type PutPipelineDefinitionRequest struct {
	PipelineId      string           `json:"pipelineId"`
	PipelineObjects []PipelineObject `json:"pipelineObjects"`
}

// Creates a new PutPipelineDefinitionRequest.
func NewPutPipelineDefinitionRequest(pipelineId string, pipelineObjects ...PipelineObject) *PutPipelineDefinitionRequest {
	return &PutPipelineDefinitionRequest{pipelineId, pipelineObjects}
}

/*****************************************************************************/

// Queries a pipeline for the names of objects that match a specified set of conditions.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_QueryObjects.html]
type QueryObjectsRequest struct {
	Limit      int    `json:"limit,omitempty"`
	Marker     string `json:"marker,omitempty"`
	PipelineId string `json:"pipelineId"`
	Query      Query  `json:"query,omitempty"`
	Sphere     string `json:"sphere"`
}

// Creates a new QueryObjectsRequest.
func NewQueryObjectsRequest(pipelineId, sphere string) *QueryObjectsRequest {
	return &QueryObjectsRequest{PipelineId: pipelineId, Sphere: sphere}
}

/*****************************************************************************/

// Updates the AWS Data Pipeline service on the progress of the calling task runner.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskProgress.html]
type ReportTaskProgressRequest struct {
	TaskId string `json:"taskId"`
}

// Creates a new ReportTaskProgressRequest.
func NewReportTaskProgressRequest(taskId string) *ReportTaskProgressRequest {
	return &ReportTaskProgressRequest{taskId}
}

/*****************************************************************************/

// Task runners call ReportTaskRunnerHeartbeat every 15 minutes to indicate that they are operational.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskRunnerHeartbeat.html]
type ReportTaskRunnerHeartbeatRequest struct {
	HostName     string `json:"hostName,omitempty"`
	TaskrunnerId string `json:"taskrunnerId"`
	WorkerGroup  string `json:"workerGroup,omitempty"`
}

// Creates a new ReportTaskRunnerHeartbeatRequest.
func NewReportTaskRunnerHeartbeatRequest(taskrunnerId string) *ReportTaskRunnerHeartbeatRequest {
	return &ReportTaskRunnerHeartbeatRequest{TaskrunnerId: taskrunnerId}
}

/*****************************************************************************/

// Requests that the status of an array of physical or logical pipeline objects be updated in the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_SetStatus.html]
type SetStatusRequest struct {
	ObjectIds  []string `json:"objectIds"`
	PipelineId string   `json:"pipelineId"`
	Status     Status   `json:"status"`
}

// Creates a new SetStatusRequest.
func NewSetStatusRequest(pipelineId string, status Status, objectIds ...string) *SetStatusRequest {
	return &SetStatusRequest{objectIds, pipelineId, status}
}

/*****************************************************************************/

// Notifies AWS Data Pipeline that a task is completed and provides information about the final status.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_SetTaskStatus.html]
type SetTaskStatusRequest struct {
	ErrorId         string `json:"errorId,omitempty"`
	ErrorMessage    string `json:"errorMessage,omitempty"`
	ErrorStackTrace string `json:"errorStackTrace,omitempty"`
	TaskId          string `json:"taskId"`
	TaskStatus      Status `json:"taskStatus"`
}

// Creates a new SetTaskStatusRequest.
func NewSetTaskStatusRequest(taskId string, taskStatus Status) *SetTaskStatusRequest {
	return &SetTaskStatusRequest{TaskId: taskId, TaskStatus: taskStatus}
}

/*****************************************************************************/

// Tests the pipeline definition with a set of validation checks to ensure
// that it is well formed and can run without error.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ValidatePipelineDefinition.html]
type ValidatePipelineDefinitionRequest struct {
	PipelineId      string           `json:"pipelineId"`
	PipelineObjects []PipelineObject `json:"pipelineObjects"`
}

// Creates a new ValidatePipelineDefinitionRequest.
func NewValidatePipelineDefinitionRequest(pipelineId string, pipelineObjects ...PipelineObject) *ValidatePipelineDefinitionRequest {
	return &ValidatePipelineDefinitionRequest{pipelineId, pipelineObjects}
}
