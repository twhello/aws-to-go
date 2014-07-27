package sqs

import ()

/******************************************************************************
 * Constants
 */

/*	The action the client wants to allow for the specified principal.
	ALL = "*"
	CHANGE_MESSAGE_VISIBILITY = "ChangeMessageVisibility"
	DELETE_MESSAGE = "DeleteMessage"
	GET_QUEUE_ATTRIBUTES = "GetQueueAttributes"
	GET_QUEUE_URL = "GetQueueUrl"
	RECEIVE_MESSAGE = "ReceiveMessage"
	SEND_MESSAGE = "SendMessage" */
type Action string

const (
	ALL_ACTIONS               Action = "*"
	CHANGE_MESSAGE_VISIBILITY Action = "ChangeMessageVisibility"
	DELETE_MESSAGE            Action = "DeleteMessage"
	GET_QUEUE_ATTRIBUTES      Action = "GetQueueAttributes"
	GET_QUEUE_URL             Action = "GetQueueUrl"
	RECEIVE_MESSAGE           Action = "ReceiveMessage"
	SEND_MESSAGE              Action = "SendMessage"
)

/*	A map of attributes to the respective values.
	APPROXIMATE_NUMBER_OF_MESSAGES = "ApproximateNumberOfMessages"
	APPROXIMATE_NUMBER_OF_MESSAGES_DELAYED = "ApproximateNumberOfMessagesDelayed"
	APPROXIMATE_NUMBER_OF_MESSAGES_NOT_VISIBILE = "ApproximateNumberOfMessagesNotVisible"
	CREATED_TIMESTAMP = "CreatedTimestamp"
	DELAY_SECONDS = "DelaySeconds"
	LAST_MODIFIED_TIMESTAMP = "LastModifiedTimestamp"
	MAXIMUM_MESSAGE_SIZE = "MaximumMessageSize"
	MESSAGE_RETENTION_PERIOD = "MessageRetentionPeriod"
	POLICY = "Policy"
	QUEUE_ARN = "QueueArn"
	RECEIVE_MESSAGE_WAIT_TIME_SECONDS = "ReceiveMessageWaitTimeSeconds"
	REDRIVE_POLICY Attrite = "RedrivePolicy"
	VISIBILITY_TIMEOUT = "VisibilityTimeout" */
type Attribute string

const (
	ALL_ATTRIBUTES                              Attribute = "All"
	APPROXIMATE_NUMBER_OF_MESSAGES              Attribute = "ApproximateNumberOfMessages"
	APPROXIMATE_NUMBER_OF_MESSAGES_DELAYED      Attribute = "ApproximateNumberOfMessagesDelayed"
	APPROXIMATE_NUMBER_OF_MESSAGES_NOT_VISIBILE Attribute = "ApproximateNumberOfMessagesNotVisible"
	CREATED_TIMESTAMP                           Attribute = "CreatedTimestamp"
	DELAY_SECONDS                               Attribute = "DelaySeconds"
	LAST_MODIFIED_TIMESTAMP                     Attribute = "LastModifiedTimestamp"
	MAXIMUM_MESSAGE_SIZE                        Attribute = "MaximumMessageSize"
	MESSAGE_RETENTION_PERIOD                    Attribute = "MessageRetentionPeriod"
	POLICY                                      Attribute = "Policy"
	QUEUE_ARN                                   Attribute = "QueueArn"
	RECEIVE_MESSAGE_WAIT_TIME_SECONDS           Attribute = "ReceiveMessageWaitTimeSeconds"
	REDRIVE_POLICY                              Attribute = "RedrivePolicy"
	VISIBILITY_TIMEOUT                          Attribute = "VisibilityTimeout"
)

/******************************************************************************
 * Data Types
 */

// This is used in the responses of batch API to give a detailed description
// of the result of an action on each entry in the request.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_BatchResultErrorEntry.html]
type BatchResultErrorEntry struct {
	Code        string `xml:"Code"`
	Id          string `xml:"Id"`
	Message     string `xml:"Message,omitempty"`
	SenderFault bool   `xml:"SenderFault"`
}

// Encloses a receipt handle and an entry id for each message in ChangeMessageVisibilityBatch.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatchRequestEntry.html]
type ChangeMessageVisibilityBatchRequestEntry struct {
	Id                string `name:"Id"`
	ReceiptHandle     string `name:"ReceiptHandle"`
	VisibilityTimeout int    `name:"VisibilityTimeout, omitempty"`
}

// For each message in the batch, the response contains a ChangeMessageVisibilityBatchResultEntry
// tag if the message succeeds or a BatchResultErrorEntry tag if the message fails.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatchResult.html]
type ChangeMessageVisibilityBatchResult struct {
	Failed     []BatchResultErrorEntry                   `name:"BatchResultErrorEntry.#."`
	Successful []ChangeMessageVisibilityBatchResultEntry `name:"ChangeMessageVisibilityBatchResultEntry.#."`
}

// Encloses the id of an entry in ChangeMessageVisibilityBatch.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatchResultEntry.html]
type ChangeMessageVisibilityBatchResultEntry struct {
	Id string `name:"ChangeMessageVisibilityBatchResultEntry"`
}

// Returns the QueueUrl element of the created queue.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueueResult.html]
type CreateQueueResult struct {
	QueueUrl string `name:QueueUrl,omitempty`
}

// Encloses a receipt handle and an identifier for it.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatchRequestEntry.html]
type DeleteMessageBatchRequestEntry struct {
	Id            string `name:"Id"`
	ReceiptHandle string `name:"ReceiptHandle"`
}

// For each message in the batch, the response contains a DeleteMessageBatchResultEntry tag
// if the message is deleted or a BatchResultErrorEntry tag if the message cannot be deleted.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatchResult.html]
type DeleteMessageBatchResult struct {
	Failed     []BatchResultErrorEntry                   `name:"BatchResultErrorEntry.#."`
	Successful []ChangeMessageVisibilityBatchResultEntry `name:"ChangeMessageVisibilityBatchResultEntry.#."`
}

// Encloses the id an entry in DeleteMessageBatch.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatchResultEntry.html]
type DeleteMessageBatchResultEntry struct {
	Id string `xml:"Id"`
}

// A list of returned queue attributes.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributesResult.html]
type GetQueueAttributesResult struct {
	Attributes map[Attribute]string `xml:"Attributes"`
}

// For more information, see Responses in the Amazon SQS Developer Guide.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/UnderstandingResponses.html]
type GetQueueUrlResult struct {
	QueueUrl string `xml:"QueueUrl,omitempty"`
}

// A list of your dead letter source queues.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListDeadLetterSourceQueuesResult.html]
type ListDeadLetterSourceQueuesResult struct {
	QueueUrls []string `xml:"QueueUrls"`
}

// A list of your queues.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueuesResult.html]
type ListQueuesResult struct {
	QueueUrls []string `xml:"QueueUrls"`
}

// An Amazon SQS message.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_Message.html]
type Message struct {
	Attributes             map[Attribute]string    `xml:"Attributes,omitempty"`
	Body                   string                  `xml:"Body,omitempty"`
	MD5OfBody              string                  `xml:"MD5OfBody,omitempty"`
	MD5OfMessageAttributes string                  `xml:"MD5OfMessageAttributes,omitempty"`
	MessageAttributes      []MessageAttributeValue `xml:"MessageAttributes,omitempty"`
	MessageId              string                  `xml:"MessageId,omitempty"`
	ReceiptHandle          string                  `xml:"ReceiptHandle,omitempty"`
}

// The user-specified message attribute value. For string data types, the value attribute has
// the same restrictions on the content as the message body.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_MessageAttributeValue.html]
// Message attribute data types:
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSMessageAttributes.html#SQSMessageAttributes.DataTypes]
type MessageAttributeValue struct {
	BinaryListValues [][]byte `xml:"BinaryListValues,omitempty"`
	BinaryValue      []byte   `xml:"BinaryValue,omitempty"`
	DataType         string   `xml:"DataType"`
	StringListValues []string `xml:"StringListValues,omitempty"`
	StringValue      string   `xml:"StringValue,omitempty"`
}

// A list of received messages.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessageResult.html]
type ReceiveMessageResult struct {
	Messages []Message `xml:"Messages"`
}

// Contains the details of a single Amazon SQS message along with a Id.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatchRequestEntry.html]
// Message Attribute Items and Validation
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSMessageAttributes.html#SQSMessageAttributesNTV]
type SendMessageBatchRequestEntry struct {
	DelaySeconds      int                              `name:"DelaySeconds,omitempty"`
	Id                string                           `name:"Id"`
	MessageAttributes map[string]MessageAttributeValue `name:"MessageAttributes,omitempty"`
	MessageBody       string                           `name:"MessageBody"`
}

// For each message in the batch, the response contains a SendMessageBatchResultEntry tag
// if the message succeeds or a BatchResultErrorEntry tag if the message fails.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatchResult.html]
type SendMessageBatchResult struct {
	Failed     []BatchResultErrorEntry       `xml:"BatchResultErrorEntry"`
	Successful []SendMessageBatchResultEntry `xml:"SendMessageBatchResultEntry"`
}

// Encloses a message ID for successfully enqueued message of a SendMessageBatch.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatchResultEntry.html]
type SendMessageBatchResultEntry struct {
	Id                     string `xml:"Id"`
	MD5OfMessageAttributes string `xml:"MD5OfMessageAttributes,omitempty"`
	MD5OfMessageBody       string `xml:"MD5OfMessageBody"`
	MessageId              string `xml:"MessageId,omitempty"`
}

// The MD5OfMessageBody and MessageId elements.
// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageResult.html]
type SendMessageResult struct {
	MD5OfMessageAttributes string `xml:"MD5OfMessageAttributes,omitempty"`
	MD5OfMessageBody       string `xml:"MD5OfMessageBody,omitempty"`
	MessageId              string `xml:"MessageId,omitempty,omitempty"`
}

/******************************************************************************
 * Responses
 */

type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_AddPermission.html]
type AddPermissionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibility.html]
type ChangeMessageVisibilityResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ChangeMessageVisibilityBatch.html]
type ChangeMessageVisibilityBatchResponse struct {
	ChangeMessageVisibilityBatchResult ChangeMessageVisibilityBatchResult `xml:"ChangeMessageVisibilityBatchResult"`
	ResponseMetadata                   ResponseMetadata                   `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html]
type CreateQueueResponse struct {
	CreateQueueResult CreateQueueResult `xml:"CreateQueueResult"`
	ResponseMetadata  ResponseMetadata  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessage.html]
type DeleteMessageResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteMessageBatch.html]
type DeleteMessageBatchResponse struct {
	DeleteMessageBatchResult DeleteMessageBatchResult `xml:"DeleteMessageBatchResult"`
	ResponseMetadata         ResponseMetadata         `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_DeleteQueue.html]
type DeleteQueueResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html]
type GetQueueAttributesResponse struct {
	GetQueueAttributesResult GetQueueAttributesResult `xml:"GetQueueAttributesResult"`
	ResponseMetadata         ResponseMetadata         `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueUrl.html]
type GetQueueUrlResponse struct {
	GetQueueUrlResult GetQueueUrlResult `xml:"GetQueueUrlResult"`
	ResponseMetadata  ResponseMetadata  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListDeadLetterSourceQueues.html]
type ListDeadLetterSourceQueuesResponse struct {
	ListDeadLetterSourceQueuesResult ListDeadLetterSourceQueuesResult `xml:"ListDeadLetterSourceQueuesResult"`
	ResponseMetadata                 ResponseMetadata                 `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueues.html]
type ListQueuesResponse struct {
	ListQueuesResult ListQueuesResult `xml:"ListQueuesResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ReceiveMessage.html]
type ReceiveMessageResponse struct {
	ReceiveMessageResult ReceiveMessageResult `xml:"ReceiveMessageResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_RemovePermission.html]
type RemovePermissionResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html]
type SendMessageResponse struct {
	SendMessageResult SendMessageResult `xml:"SendMessageResult"`
	ResponseMetadata  ResponseMetadata  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatch.html]
type SendMessageBatchResponse struct {
	SendMessageBatchResult SendMessageBatchResult `xml:"SendMessageBatchResult"`
	ResponseMetadata       ResponseMetadata       `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SetQueueAttributes.html]
type SetQueueAttributesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}
