package datapipeline

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"net/http"
)

const ServiceName = "datapipeline"

// DataPipelineService describes the API interface.
// Instantiate with swf.NewService().
type DataPipelineService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *DataPipelineService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *DataPipelineService) RegionName() string {
	return s.region.Name()
}

func (s *DataPipelineService) Endpoint() string {
	return s.endpoint
}

// Low-level request to Data Pipeline service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *DataPipelineService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *DataPipelineService) wrapperSignAndDo(target string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	if err == nil {
		h := req.Header()
		h.Set("Content-Type", "application/x-amz-json-1.1")
		h.Set("X-Amz-Target", target)
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * Data Pipeline Service Methods.
 */

// Validates a pipeline and initiates processing.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ActivatePipeline.html]
func (s *DataPipelineService) ActivatePipeline(req *ActivatePipelineRequest) (result *ActivatePipelineResult, err error) {

	result = new(ActivatePipelineResult)
	err = s.wrapperSignAndDo("DataPipelineService.ActivatePipeline", req, result)
	return
}

// Creates a new empty pipeline. When this action succeeds, you can then use the
// PutPipelineDefinition action to populate the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_CreatePipeline.html]
func (s *DataPipelineService) CreatePipeline(req *CreatePipelineRequest) (result *CreatePipelineResult, err error) {

	result = new(CreatePipelineResult)
	err = s.wrapperSignAndDo("DataPipelineService.CreatePipeline", req, result)
	return
}

// Permanently deletes a pipeline, its pipeline definition and its run history.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DeletePipeline.html]
func (s *DataPipelineService) DeletePipeline(req *DeletePipelineRequest) (err error) {

	err = s.wrapperSignAndDo("DataPipelineService.DeletePipeline", req, nil)
	return
}

// Returns the object definitions for a set of objects associated with the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribeObjects.html]
func (s *DataPipelineService) DescribeObjects(req *DescribeObjectsRequest) (result *DescribeObjectsResult, err error) {

	result = new(DescribeObjectsResult)
	err = s.wrapperSignAndDo("DataPipelineService.DescribeObjects", req, result)
	return
}

// Retrieve metadata about one or more pipelines.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_DescribePipelines.html]
func (s *DataPipelineService) DescribePipelines(req *DescribePipelinesRequest) (result *DescribePipelinesResult, err error) {

	result = new(DescribePipelinesResult)
	err = s.wrapperSignAndDo("DataPipelineService.DescribePipelines", req, result)
	return
}

// Evaluates a string in the context of a specified object.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_EvaluateExpression.html]
func (s *DataPipelineService) EvaluateExpression(req *EvaluateExpressionRequest) (result *EvaluateExpressionResult, err error) {

	result = new(EvaluateExpressionResult)
	err = s.wrapperSignAndDo("DataPipelineService.EvaluateExpression", req, result)
	return
}

// Returns the definition of the specified pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinition.html]
func (s *DataPipelineService) GetPipelineDefinition(req *GetPipelineDefinitionRequest) (result *GetPipelineDefinitionResult, err error) {

	result = new(GetPipelineDefinitionResult)
	err = s.wrapperSignAndDo("DataPipelineService.GetPipelineDefinition", req, result)
	return
}

// Returns a list of pipeline identifiers for all active pipelines.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ListPipelines.html]
func (s *DataPipelineService) ListPipelines(req *ListPipelinesRequest) (result *ListPipelinesResult, err error) {

	result = new(ListPipelinesResult)
	err = s.wrapperSignAndDo("DataPipelineService.ListPipelines", req, result)
	return
}

// Task runners call this action to receive a task to perform from AWS Data Pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PollForTask.html]
func (s *DataPipelineService) PollForTask(req *PollForTaskRequest) (result *PollForTaskResult, err error) {

	result = new(PollForTaskResult)
	err = s.wrapperSignAndDo("DataPipelineService.PollForTask", req, result)
	return
}

// Adds tasks, schedules, and preconditions that control the behavior of the pipeline.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html]
func (s *DataPipelineService) PutPipelineDefinition(req *PutPipelineDefinitionRequest) (result *PutPipelineDefinitionResult, err error) {

	result = new(PutPipelineDefinitionResult)
	err = s.wrapperSignAndDo("DataPipelineService.PutPipelineDefinition", req, result)
	return
}

// Queries a pipeline for the names of objects that match a specified set of conditions.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_QueryObjects.html]
func (s *DataPipelineService) QueryObjects(req *QueryObjectsRequest) (result *QueryObjectsResult, err error) {

	result = new(QueryObjectsResult)
	err = s.wrapperSignAndDo("DataPipelineService.QueryObjects", req, result)
	return
}

// Updates the AWS Data Pipeline service on the progress of the calling task runner.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskProgress.html]
func (s *DataPipelineService) ReportTaskProgress(req *ReportTaskProgressRequest) (result *ReportTaskProgressResult, err error) {

	result = new(ReportTaskProgressResult)
	err = s.wrapperSignAndDo("DataPipelineService.ReportTaskProgress", req, result)
	return
}

// Task runners call ReportTaskRunnerHeartbeat every 15 minutes to indicate that they are operational.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ReportTaskRunnerHeartbeat.html]
func (s *DataPipelineService) ReportTaskRunnerHeartbeat(req *ReportTaskRunnerHeartbeatRequest) (result *ReportTaskRunnerHeartbeatResult, err error) {

	result = new(ReportTaskRunnerHeartbeatResult)
	err = s.wrapperSignAndDo("DataPipelineService.ReportTaskRunnerHeartbeat", req, result)
	return
}

// Notifies AWS Data Pipeline that a task is completed and provides information about the final status.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_SetTaskStatus.html]
func (s *DataPipelineService) SetStatus(req *SetStatusRequest) (err error) {

	err = s.wrapperSignAndDo("DataPipelineService.SetStatus", req, nil)
	return
}

// Tests the pipeline definition with a set of validation checks to ensure
// that it is well formed and can run without error.
// [http://docs.aws.amazon.com/datapipeline/latest/APIReference/API_ValidatePipelineDefinition.html]
func (s *DataPipelineService) SetTaskStatus(req *SetTaskStatusRequest) (result *SetTaskStatusResult, err error) {

	result = new(SetTaskStatusResult)
	err = s.wrapperSignAndDo("DataPipelineService.SetTaskStatus", req, result)
	return
}

//
// []
func (s *DataPipelineService) ValidatePipelineDefinition(req *ValidatePipelineDefinitionRequest) (result *ValidatePipelineDefinitionResult, err error) {

	result = new(ValidatePipelineDefinitionResult)
	err = s.wrapperSignAndDo("DataPipelineService.ValidatePipelineDefinition", req, result)
	return
}

// Creates a new Simple Workflow Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *DataPipelineService {
	return &DataPipelineService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
