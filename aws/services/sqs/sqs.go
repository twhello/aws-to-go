//
// Amazon Simple Queue Service (Amazon SQS)  is a messaging queue service that handles message
// or workflows between other components in a system.
//
// [http://aws.amazon.com/documentation/sqs/]
//
package sqs

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"github.com/twhello/aws-to-go/aws/util/netutil"
	"net/http"
	"time"
)

const ServiceName = "sqs"

// The SQS Service object. Use ses.NewService().
type SQSService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *SQSService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *SQSService) RegionName() string {
	return s.region.Name()
}

func (s *SQSService) Endpoint() string {
	return s.endpoint
}

// Low-level request to SQS service.
func (s *SQSService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *SQSService) wrapperSignAndDo(url, action string, request, result interface{}) (err error) {
	
	qs := netutil.MarshalValues(request)
	qs.Add("AWSAccessKeyId", s.cred.AccessKeyId())
	qs.Add("Expires", time.Now().Add(time.Second*300).Format("2006-01-02T15:04:05MST")) // 2011-10-24T22:52:43PST
	qs.Add("Version", "2012-11-05")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", url, qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * SQS Service Methods.
 */

// Adds a permission to a queue for a specific principal. This allows for sharing access to the queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_AddPermission.html]
func (s *SQSService) AddPermission(req *AddPermissionRequest) (result *AddPermissionResponse, err error) {

	result = new(AddPermissionResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "AddPermission", req, result)
	return
}

// Changes the visibility timeout of a specified message in a queue to a new value.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibility.html]
func (s *SQSService) ChangeMessageVisibility(req *ChangeMessageVisibilityRequest) (result *ChangeMessageVisibilityResponse, err error) {

	result = new(ChangeMessageVisibilityResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "ChangeMessageVisibility", req, result)
	return
}

// Changes the visibility timeout of multiple messages.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatch.html]
func (s *SQSService) ChangeMessageVisibilityBatch(req *ChangeMessageVisibilityBatchRequest) (result *ChangeMessageVisibilityBatchResponse, err error) {

	result = new(ChangeMessageVisibilityBatchResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "ChangeMessageVisibilityBatch", req, result)
	return
}

// Creates a new queue, or returns the URL of an existing one.
// When you request CreateQueue, you provide a name for the queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html]
func (s *SQSService) CreateQueue(req *CreateQueueRequest) (result *CreateQueueResponse, err error) {

	result = new(CreateQueueResponse)
	err = s.wrapperSignAndDo(s.Endpoint(), "CreateQueue",  req, result)
	return
}

// Deletes the specified message from the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessage.html]
func (s *SQSService) DeleteMessage(req *DeleteMessageRequest) (result *DeleteMessageResponse, err error) {

	result = new(DeleteMessageResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "DeleteMessage", req, result)
	return
}

// Deletes multiple messages. This is a batch version of DeleteMessage.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatch.html]
func (s *SQSService) DeleteMessageBatch(req *DeleteMessageBatchRequest) (result *DeleteMessageBatchResponse, err error) {

	result = new(DeleteMessageBatchResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "DeleteMessageBatch", req, result)
	return
}

// Deletes the queue specified by the queue URL, regardless of whether the queue is empty.
// If the specified queue does not exist, Amazon SQS returns a successful response. 
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteQueue.html]
func (s *SQSService) DeleteQueue(req *DeleteQueueRequest) (result *DeleteQueueResponse, err error) {

	result = new(DeleteQueueResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "DeleteQueue", req, result)
	return
}

// Gets attributes for the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html]
func (s *SQSService) GetQueueAttributes(req *GetQueueAttributesRequest) (result *GetQueueAttributesResponse, err error) {

	result = new(GetQueueAttributesResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "GetQueueAttributes", req, result)
	return
}


// Returns the URL of an existing queue. This action provides a simple way to retrieve
// the URL of an Amazon SQS queue. 
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueUrl.html]
func (s *SQSService) GetQueueUrl(req *GetQueueUrlRequest) (result *GetQueueUrlResponse, err error) {

	result = new(GetQueueUrlResponse)
	err = s.wrapperSignAndDo(s.Endpoint(), "GetQueueUrl", req, result)
	return
}

// Returns a list of your queues that have the RedrivePolicy queue attribute configured with a dead letter queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListDeadLetterSourceQueues.html]
func (s *SQSService) ListDeadLetterSourceQueues() (result *ListDeadLetterSourceQueuesResponse, err error) {

	result = new(ListDeadLetterSourceQueuesResponse)
	err = s.wrapperSignAndDo(s.Endpoint(), "ListDeadLetterSourceQueues", nil, result)
	return
}

// Returns a list of your queues. The maximum number of queues that can be returned is 1000. If you specify a value for the optional QueueNamePrefix parameter, only queues with a name beginning with the specified value are returned.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueues.html]
func (s *SQSService) ListQueues(req *ListQueuesRequest) (result *ListQueuesResponse, err error) {

	result = new(ListQueuesResponse)
	err = s.wrapperSignAndDo(s.Endpoint(), "ListQueues", req, result)
	return
}

// Retrieves one or more messages, with a maximum limit of 10 messages, from the specified queue.
// Long poll support is enabled by using the WaitTimeSeconds parameter.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html]
func (s *SQSService) ReceiveMessage(req *ReceiveMessageRequest) (result *ReceiveMessageResponse, err error) {

	result = new(ReceiveMessageResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "ReceiveMessage", req, result)
	return
}

// Revokes any permissions in the queue policy that matches the specified Label parameter.
// Only the owner of the queue can remove permissions.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_RemovePermission.html]
func (s *SQSService) RemovePermission(req *RemovePermissionRequest) (result *RemovePermissionResponse, err error) {

	result = new(RemovePermissionResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "RemovePermission", req, result)
	return
}

// Delivers a message to the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html]
func (s *SQSService) SendMessage(req *SendMessageRequest) (result *SendMessageResponse, err error) {

	result = new(SendMessageResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "SendMessage", req, result)
	return
}

// Delivers up to ten messages to the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatch.html]
func (s *SQSService) SendMessageBatch(req *SendMessageBatchRequest) (result *SendMessageBatchResponse, err error) {

	result = new(SendMessageBatchResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "SendMessageBatch", req, result)
	return
}

// Sets the value of one or more queue attributes.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SetQueueAttributes.html]
func (s *SQSService) SetQueueAttributes(req *SetQueueAttributesRequest) (result *SetQueueAttributesResponse, err error) {

	result = new(SetQueueAttributesResponse)
	err = s.wrapperSignAndDo(req.QueueUrl, "SetQueueAttributes", req, result)
	return
}

// Creates a new SES Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *SQSService {

	subdomain := ServiceName
	if region.Name() != regions.DEFAULT_REGION {
		subdomain += "." + region.Name()
	}
	return &SQSService{cred, region, "https://" + subdomain + ".amazonaws.com"}
}
