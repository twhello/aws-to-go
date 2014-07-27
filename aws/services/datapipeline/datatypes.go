package datapipeline

import ()

/******************************************************************************
 * Constants
 */

// Specifies the status to be set on all the objects in objectIds:
//	CANCEL Status = "CANCEL"
//	MARK_FINISHED
//	PAUSE
//	RERUN
//	RESUME
type Status string

const (
	CANCEL        Status = "CANCEL"
	FALSE         Status = "FALSE"
	FAILED        Status = "FAILED"
	FINISHED      Status = "FINISHED"
	MARK_FINISHED Status = "MARK_FINISHED"
	PAUSE         Status = "PAUSE"
	RERUN         Status = "RERUN"
	RESUME        Status = "RESUME"
)

// The logical operation to be performed:
//	BETWEEN Type = "BETWEEN"
//	EQ - Equal
//	GE - Greater than or equal
//	LE - Less than or equal
//	REF_EQ - Equal reference
type Type string

const (
	BETWEEN Type = "BETWEEN"
	EQ      Type = "EQ"
	GE      Type = "GE"
	LE      Type = "LE"
	REF_EQ  Type = "REF_EQ"
)

/******************************************************************************
 * Data Types
 */

// Contains the output from the ActivatePipeline action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ActivatePipelineResult.html]
type ActivatePipelineResult struct{}

// Contains the output from the CreatePipeline action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_CreatePipelineResult.html]
type CreatePipelineResult struct {
	PipelineId string `json:"pipelineId"`
}

// If True, there are more results that can be returned in another call to DescribeObjects.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribeObjectsResult.html]
type DescribeObjectsResult struct {
	HasMoreResults  bool             `json:"hasMoreResults,omitempty"`
	Marker          string           `json:"marker,omitempty"`
	PipelineObjects []PipelineObject `json:"pipelineObjects"`
}

// Contains the output from the DescribePipelines action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribePipelinesResult.html]
type DescribePipelinesResult struct {
	PipelineDescriptionList []PipelineDescription `json:"pipelineDescriptionList"`
}

// Contains the output from the EvaluateExpression action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_EvaluateExpressionResult.html]
type EvaluateExpressionResult struct {
	EvaluatedExpression string `json:"evaluatedExpression"`
}

// A key-value pair that describes a property of a pipeline object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_Field.html]
type Field struct {
	Key         string `json:"key"`
	RefValue    string `json:"refValue,omitempty"`
	StringValue string `json:"stringValue,omitempty"`
}

// Contains the output from the GetPipelineDefinition action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinitionResult.html]
type GetPipelineDefinitionResult struct {
	PipelineObjects []PipelineObject `json:"pipelineObjects,omitempty"`
}

//
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_InstanceIdentity.html]
type InstanceIdentity struct {
	Document  string `json:"document,omitempty"`
	Signature string `json:"signature,omitempty"`
}

// Contains the output from the ListPipelines action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ListPipelinesResult.html]
type ListPipelinesResult struct {
	HasMoreResults bool             `json:"hasMoreResults,omitempty"`
	Marker         string           `json:"marker,omitempty"`
	PipelineIdList []PipelineIdName `json:"pipelineIdList"`
}

// Contains a logical operation for comparing the value of a field with a specified value.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_Operator.html]
type Operator struct {
	Type   Type     `json:"type,omitempty"`
	Values []string `json:"value,omitempty"`
}

// Contains pipeline metadata.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PipelineDescription.html]
type PipelineDescription struct {
	Description string  `json:"description,omitempty"`
	Fields      []Field `json:"fields"`
	Name        string  `json:"name"`
	PipelinId   string  `json:"pipelinId"`
}

// Contains the name and identifier of a pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PipelineIdName.html]
type PipelineIdName struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Contains information about a pipeline object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PipelineObject.html]
type PipelineObject struct {
	Fields []Field `json:"fields"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
}

// Contains the output from the PollForTask action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PollForTaskResult.html]
type PollForTaskResult struct {
	TaskObject TaskObject `json:"taskObject,omitempty"`
}

// Contains the output of the PutPipelineDefinition action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinitionResult.html]
type PutPipelineDefinitionResult struct {
	Errored            bool                `json:"errored"`
	ValidationErrors   []ValidationError   `json:"validationErrors,omitempty"`
	ValidationWarnings []ValidationWarning `json:"validationWarnings,omitempty"`
}

// Defines the query to run against an object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_Query.html]
type Query struct {
	Selectors []Selector `json:"selectors,omitempty"`
}

// Contains the output from the QueryObjects action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_QueryObjectsResult.html]
type QueryObjectsResult struct {
	HasMoreResults bool     `json:"hasMoreResults,omitempty"`
	Ids            []string `json:"ids,omitempty"`
	Marker         string   `json:"marker,omitempty"`
}

// Contains the output from the ReportTaskProgress action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskProgressResult.html]
type ReportTaskProgressResult struct {
	Canceled bool `json:"canceled"`
}

// http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskRunnerHeartbeatResult.html
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskRunnerHeartbeatResult.html]
type ReportTaskRunnerHeartbeatResult struct {
	Terminate bool `json:"terminate"`
}

// A comparision that is used to determine whether a query should return this object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_Selector.html]
type Selector struct {
	FieldName string    `json:"fieldName,omitempty"`
	Operator  *Operator `json:"operator,omitempty"`
}

// The output from the SetTaskStatus action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_SetTaskStatusResult.html]
type SetTaskStatusResult struct{}

// Contains information about a pipeline task that is assigned to a task runner.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_TaskObject.html]
type TaskObject struct {
	AttemptId  string                    `json:"attemptId,omitempty"`
	Objects    map[string]PipelineObject `json:"objects,omitempty"`
	PipelineId string                    `json:"pipelineId,omitempty"`
	TaskId     string                    `json:"taskId,omitempty"`
}

// Contains the output from the ValidatePipelineDefinition action.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ValidatePipelineDefinitionResult.html]
type ValidatePipelineDefinitionResult struct {
	Errored            bool                `json:"errored"`
	ValidationErrors   []ValidationError   `json:"validationErrors,omitempty"`
	ValidationWarnings []ValidationWarning `json:"validationWarnings,omitempty"`
}

// Defines a validation error returned by PutPipelineDefinition or ValidatePipelineDefinition.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ValidationError.html]
type ValidationError struct {
	Errors []string `json:"errors,omitempty"`
	Id     string   `json:"id,omitempty"`
}

// Defines a validation warning returned by PutPipelineDefinition or ValidatePipelineDefinition.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ValidationWarning.html]
type ValidationWarning struct {
	Id       string   `json:"id,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
