package sqs

import ()

/*****************************************************************************/

// Adds a permission to a queue for a specific principal. This allows for
// sharing access to the queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_AddPermission.html]
type AddPermissionRequest struct {
	AWSAccountIds []string `name:"AWSAccountId.#" base:"1"`
	ActionNames   []Action `name:"ActionName.#" base:"1"`
	Label         string   `name:"Label"`
	QueueUrl      string   `name:"-"`
}

// Creates a new AddPermissionRequest.
func NewAddPermissionRequest(queueUrl, label, awsAccountId string, action Action) *AddPermissionRequest {
	return &AddPermissionRequest{[]string{awsAccountId}, []Action{action}, label, queueUrl}
}

// Add an Action.
func (r *AddPermissionRequest) AddAction(awsAccountId string, action Action) {
	r.ActionNames = append(r.ActionNames, action)
	r.AWSAccountIds = append(r.AWSAccountIds, awsAccountId)
}

/*****************************************************************************/

// Changes the visibility timeout of a specified message in a queue to a new value.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibility.html]
type ChangeMessageVisibilityRequest struct {
	QueueUrl          string `name:"-"`
	ReceiptHandle     string `name:"ReceiptHandle"`
	VisibilityTimeout int    `name:"VisibilityTimeout"`
}

// Creates a new ChangeMessageVisibilityRequest.
func NewChangeMessageVisibilityRequest(queueUrl, receiptHandle string, visibilityTimeout int) *ChangeMessageVisibilityRequest {
	return &ChangeMessageVisibilityRequest{queueUrl, receiptHandle, visibilityTimeout}
}

/*****************************************************************************/

// Changes the visibility timeout of multiple messages.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatch.html]
type ChangeMessageVisibilityBatchRequest struct {
	ChangeMessageVisibilityBatchRequestEntry []ChangeMessageVisibilityBatchRequestEntry `name:"ChangeMessageVisibilityBatchRequestEntry.#." base:"1"`
	QueueUrl                                 string                                     `name:"-"`
}

// Creates a new ChangeMessageVisibilityBatchRequest.
func NewChangeMessageVisibilityBatchRequest(queueUrl string, entry ...ChangeMessageVisibilityBatchRequestEntry) *ChangeMessageVisibilityBatchRequest {
	return &ChangeMessageVisibilityBatchRequest{entry, queueUrl}
}

/*****************************************************************************/

// Creates a new queue, or returns the URL of an existing one.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html]
type CreateQueueRequest struct {
	AttributeNames  []Attribute `name:"Attribute.#.Name,omitempty" base:"1"`
	AttributeValues []string    `name:"Attribute.#.Value,omitempty" base:"1"`
	QueueUrl        string      `name:"-"`
}

// Creates a new CreateQueueRequest.
func NewCreateQueueRequest(queueUrl string) *CreateQueueRequest {
	return &CreateQueueRequest{QueueUrl: queueUrl}
}

// Add an Attribute.
func (r *CreateQueueRequest) AddAttribute(name Attribute, value string) {
	r.AttributeNames = append(r.AttributeNames, name)
	r.AttributeValues = append(r.AttributeValues, value)
}

/*****************************************************************************/

// Deletes the specified message from the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessage.html]
type DeleteMessageRequest struct {
	QueueUrl      string `xml:"-"`
	ReceiptHandle string `name:"ReceiptHandle"`
}

// Creates a new DeleteMessageRequest.
func NewDeleteMessageRequest() *DeleteMessageRequest {
	return &DeleteMessageRequest{}
}

/*****************************************************************************/

// Deletes multiple messages.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatch.html]
type DeleteMessageBatchRequest struct {
	DeleteMessageBatchRequestEntry []DeleteMessageBatchRequestEntry `name:"DeleteMessageBatchRequestEntry.#." base:"1"`
	QueueUrl                       string                           `name:"-"`
}

// Creates a new DeleteMessageBatchRequest.
func NewDeleteMessageBatchRequest(queueUrl string, entry ...DeleteMessageBatchRequestEntry) *DeleteMessageBatchRequest {
	return &DeleteMessageBatchRequest{entry, queueUrl}
}

/*****************************************************************************/

// Deletes the queue specified by the queue URL, regardless of whether the queue is empty.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteQueue.html]
type DeleteQueueRequest struct {
	QueueUrl string `name:"-"`
}

// Creates a new DeleteQueueRequest.
func NewDeleteQueueRequest(queueUrl string) *DeleteQueueRequest {
	return &DeleteQueueRequest{queueUrl}
}

/*****************************************************************************/

// Gets attributes for the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html]
type GetQueueAttributesRequest struct {
	AttributeName []Attribute `name:"AttributeName.#,omitempty" base:"1"`
	QueueUrl      string      `name:"-"`
}

// Creates a new GetQueueAttributesRequest.
func NewGetQueueAttributesRequest(queueUrl string) *GetQueueAttributesRequest {
	return &GetQueueAttributesRequest{QueueUrl: queueUrl}
}

// Add one or more Attributes.
func (r *GetQueueAttributesRequest) Add(attributes ...Attribute) {
	r.AttributeName = append(r.AttributeName, attributes...)
}

/*****************************************************************************/

// Returns the URL of an existing queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueUrl.html]
type GetQueueUrlRequest struct {
	QueueName              string `name:"QueueName"`
	QueueOwnerAWSAccountId string `name:"QueueOwnerAWSAccountId,omitempty"`
}

// Creates a new GetQueueUrlRequest.
func NewGetQueueUrlRequest(queueName string) *GetQueueUrlRequest {
	return &GetQueueUrlRequest{QueueName: queueName}
}

/*****************************************************************************/

// Returns a list of your queues that have the RedrivePolicy queue attribute
// configured with a dead letter queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListDeadLetterSourceQueues.html]
type ListDeadLetterSourceQueuesRequest struct {
	QueueUrl string `name:"-"`
}

// Creates a new ListDeadLetterSourceQueuesRequest.
func NewListDeadLetterSourceQueuesRequest(queueUrl string) *ListDeadLetterSourceQueuesRequest {
	return &ListDeadLetterSourceQueuesRequest{queueUrl}
}

/*****************************************************************************/

// Returns a list of your queues.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueues.html]
type ListQueuesRequest struct {
	QueueNamePrefix string `name:"QueueNamePrefix"`
}

// Creates a new ListQueuesRequest.
func NewListQueuesRequest(queueNamePrefix string) *ListQueuesRequest {
	return &ListQueuesRequest{queueNamePrefix}
}

/*****************************************************************************/

// Retrieves one or more messages, with a maximum limit of 10 messages,
// from the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html]
type ReceiveMessageRequest struct {
	AttributeName        []Attribute `name:"AttributeName.#,omitempty" base:"1"`
	MaxNumberOfMessages  int         `name:"MaxNumberOfMessages,omitempty"`
	MessageAttributeName []string    `name:"MessageAttributeName.#,omitempty" base:"1"`
	QueueUrl             string      `name:"-"`
	VisibilityTimeout    int         `name:"VisibilityTimeout,omitempty"`
	WaitTimeSeconds      int         `name:"WaitTimeSeconds,omitempty"`
}

// Creates a new ReceiveMessageRequest.
func NewReceiveMessageRequest(queueUrl string) *ReceiveMessageRequest {
	return &ReceiveMessageRequest{QueueUrl: queueUrl}
}

/*****************************************************************************/

// Revokes any permissions in the queue policy that matches the specified Label parameter.
// Only the owner of the queue can remove permissions.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_RemovePermission.html]
type RemovePermissionRequest struct {
	Label    string `name:"Label"`
	QueueUrl string `name:"-"`
}

// Creates a new RemovePermissionRequest.
func NewRemovePermissionRequest(queueUrl string, label string) *RemovePermissionRequest {
	return &RemovePermissionRequest{label, queueUrl}
}

/*****************************************************************************/

// Delivers a message to the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html]
type SendMessageRequest struct {
	DelaySeconds          int                     `name:"DelaySeconds,omitempty"`
	MessageAttributeName  []string                `name:"MessageAttribute.#.Name.,omitempty" base:"1"`
	MessageAttributeValue []MessageAttributeValue `name:"MessageAttribute.#.Value.,omitempty" base:"1"`
	MessageBody           string                  `name:"MessageBody"`
	QueueUrl              string                  `name:"-"`
}

// Creates a new .
func NewSendMessageRequest(queueUrl string, messageBody string) *SendMessageRequest {
	return &SendMessageRequest{MessageBody: messageBody, QueueUrl: queueUrl}
}

// Add MessageAttributeValue.
func (r *SendMessageRequest) AddMessageAttributeValue(name string, value MessageAttributeValue) {
	r.MessageAttributeName = append(r.MessageAttributeName, name)
	r.MessageAttributeValue = append(r.MessageAttributeValue, value)
}

/*****************************************************************************/

// Delivers up to ten messages to the specified queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatch.html]
type SendMessageBatchRequest struct {
	SendMessageBatchRequestEntry []SendMessageBatchRequestEntry `name:"SendMessageBatchRequestEntry.#." base:"1"`
	QueueUrl                     string                         `name:"-"`
}

// Creates a new SendMessageBatchRequest.
func NewSendMessageBatchRequest(queueUrl string, entry ...SendMessageBatchRequestEntry) *SendMessageBatchRequest {
	return &SendMessageBatchRequest{entry, queueUrl}
}

/*****************************************************************************/

// Sets the value of one or more queue attributes.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SetQueueAttributes.html]
type SetQueueAttributesRequest struct {
	AttributeNames  []Attribute `name:"Attribute.Name.#" base:"1"`
	AttributeValues []string    `name:"Attribute.Value.#" base:"1"`
	QueueUrl        string      `name:"-"`
}

// Creates a new SetQueueAttributesRequest.
func NewSetQueueAttributesRequest(queueUrl string) *SetQueueAttributesRequest {
	return &SetQueueAttributesRequest{}
}

// Add an Attribute.
func (r *SetQueueAttributesRequest) AddAttribute(name Attribute, value string) {
	r.AttributeNames = append(r.AttributeNames, name)
	r.AttributeValues = append(r.AttributeValues, value)
}
