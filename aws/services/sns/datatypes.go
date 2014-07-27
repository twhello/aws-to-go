package sns

import ()

/******************************************************************************
 * Constants
 */

/*	Amazon SNS supports the following logical data types:
STRING - Strings are Unicode with UTF-8 binary encoding. For a list of code values, see http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters.
NUMBER - Numbers are positive or negative integers or floating-point numbers. Numbers have sufficient range and precision to encompass most of the possible values that integers, floats, and doubles typically support. A number can have up to 38 digits of precision, and it can be between 10^-128 to 10^+126. Leading and trailing zeroes are trimmed.
BINARY - Binary type attributes can store any binary data, for example, compressed data, encrypted data, or images. */
type DataType string

const (
	STRING DataType = "String"
	NUMBER DataType = "Number"
	BINARY DataType = "Binary"
)

/*	Endpoint Attributes:
CUSTOM_USER_DATA - arbitrary user data to associate with the endpoint. Amazon SNS does not use this data. The data must be in UTF-8 format and less than 2KB.
ENABLED - flag that enables/disables delivery to the endpoint. Amazon SNS will set this to false when a notification service indicates to Amazon SNS that the endpoint is invalid. Users can set it back to true, typically after updating Token.
TOKEN - device token, also referred to as a registration id, for an app and mobile device. This is returned from the notification service when an app and mobile device are registered with the notification service. */
type EndpointAttribute string

const (
	CUSTOM_USER_DATA EndpointAttribute = "CustomUserData"
	ENABLED          EndpointAttribute = "Enabled"
	TOKEN            EndpointAttribute = "Token"
)

/*	Platform Attribute:
EVENT_DELIVERY_FAILURE - Topic ARN to which DeliveryFailure event notifications should be sent upon Direct Publish delivery failure (permanent) to one of the application's endpoints.
EVENT_ENDPOINT_CREATED - Topic ARN to which EndpointCreated event notifications should be sent.
EVENT_ENDPOINT_DELETED - Topic ARN to which EndpointDeleted event notifications should be sent.
EVENT_ENDPOINT_UPDATED - Topic ARN to which EndpointUpdate event notifications should be sent.
PLATFORM_CREDENTIAL - The credential received from the notification service. For APNS/APNS_SANDBOX, PlatformCredential is "private key". For GCM, PlatformCredential is "API key". For ADM, PlatformCredential is "client secret".
PLATFORM_PRINCIPAL - The principal received from the notification service. For APNS/APNS_SANDBOX, PlatformPrincipal is "SSL certificate". For GCM, PlatformPrincipal is not applicable. For ADM, PlatformPrincipal is "client id". */
type PlatformAttribute string

const (
	EVENT_DELIVERY_FAILURE PlatformAttribute = "EventDeliveryFailure"
	EVENT_ENDPOINT_CREATED PlatformAttribute = "EventEndpointCreated"
	EVENT_ENDPOINT_DELETED PlatformAttribute = "EventEndpointDeleted"
	EVENT_ENDPOINT_UPDATED PlatformAttribute = "EventEndpointUpdated"
	PLATFORM_CREDENTIAL    PlatformAttribute = "PlatformCredential"
	PLATFORM_PRINCIPAL     PlatformAttribute = "PlatformPrincipal"
)

/*	Subscription Attributes:
    CONFIRMATION_WAS_AUTHENTICATED - true if the subscription confirmation request was authenticated.
    DELIVERY_POLICY - the JSON serialization of the subscription's delivery policy.
    DisplayName -- the human-readable name used in the "From" field for notifications to email and email-json endpoints.
    EFFECTIVE_DELIVERY_POLICY -- the JSON serialization of the effective delivery policy that takes into account the topic delivery policy and account system defaults.
    OWNER - the AWS account ID of the subscription's owner.
	POLICY - the JSON serialization of the topic's access control policy.
	RAW_MESSAGE_DELIVERY -
    SUBSCRIPTION_ARN - the subscription's ARN.
	SUBSCRIPTIONS_CONFIRMED - the number of confirmed subscriptions on this topic.
	SUBSCRIPTIONS_DELETED - the number of deleted subscriptions on this topic.
	SUBSCRIPTIONS_PENDING - the number of subscriptions pending confirmation on this topic
    TOPIC_ARN - the topic ARN that the subscription is associated with. */
type SubscriptionAttribute string

const (
	CONFIRMATION_WAS_AUTHENTICATED SubscriptionAttribute = "ConfirmationWasAuthenticated"
	DELIVERY_POLICY                SubscriptionAttribute = "DeliveryPolicy"
	DISPLAY_NAME                   SubscriptionAttribute = "DisplayName"
	EFFECTIVE_DELIVERY_POLICY      SubscriptionAttribute = "EffectiveDeliveryPolicy"
	OWNER                          SubscriptionAttribute = "Owner"
	POLICY                         SubscriptionAttribute = "Policy"
	RAW_MESSAGE_DELIVERY           SubscriptionAttribute = "RawMessageDelivery"
	SUBSCRIPTION_ARN               SubscriptionAttribute = "SubscriptionArn"
	SUBSCRIPTIONS_PENDING          SubscriptionAttribute = "SubscriptionsPending"
	SUBSCRIPTIONS_CONFIRMED        SubscriptionAttribute = "SubscriptionsConfirmed"
	SUBSCRIPTIONS_DELETED          SubscriptionAttribute = "SubscriptionsDeleted"
	TOPIC_ARN                      SubscriptionAttribute = "TopicArn"
)

/******************************************************************************
 * Data Types
 */

// Response for ConfirmSubscriptions action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ConfirmSubscriptionResult.html]
type ConfirmSubscriptionResult struct {
	SubscriptionArn string `xml:"SubscriptionArn,omitempty"`
}

// Response from CreatePlatformApplication action.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplicationResult.html]
type CreatePlatformApplicationResult struct {
	PlatformApplicationArn string `xml:"PlatformApplicationArn,omitempty"`
}

// Response from CreateEndpoint action.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformEndpointResult.html]
type CreatePlatformEndpointResult struct {
	EndpointArn string `xml:"EndpointArn"`
}

// Response from CreateTopic action.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreateTopicResult.html]
type CreateTopicResult struct {
	TopicArn string `xml:"TopicArn"`
}

// Endpoint for mobile app and device.
// [http://docs.aws.amazon.com/sns/latest/api/API_Endpoint.html]
type Endpoint struct {
	Attributes  map[EndpointAttribute]string `xml:"Attributes>entry,omitempty"`
	EndpointArn string                       `xml:"EndpointArn"`
}

// Response from GetEndpointAttributes of the EndpointArn.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetEndpointAttributesResult.html]
type GetEndpointAttributesResult struct {
	Attributes map[EndpointAttribute]string `xml:"Attributes>entry,omitempty"`
}

// Response for GetPlatformApplicationAttributes action.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetPlatformApplicationAttributesResult.html]
type GetPlatformApplicationAttributesResult struct {
	Attributes map[PlatformAttribute]string `xml:"Attributes>entry,omitempty"`
}

// Response for GetSubscriptionAttributes action.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributesResult.html]
type GetSubscriptionAttributesResult struct {
	Attributes map[SubscriptionAttribute]string `xml:"Attributes>entry,omitempty"`
}

// Response for GetTopicAttributes action.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributesResult.html]
type GetTopicAttributesResult struct {
	Attributes map[SubscriptionAttribute]string `xml:"Attributes>entry,omitempty"`
}

// Response for ListEndpointsByPlatformApplication action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListEndpointsByPlatformApplicationResult.html]
type ListEndpointsByPlatformApplicationResult struct {
	Endpoints []Endpoint `xml:"Endpoints,omitempty"`
	NextToken string     `xml:"NextToken,omitempty"`
}

// Response for ListPlatformApplications action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListPlatformApplicationsResult.html]
type ListPlatformApplicationsResult struct {
	NextToken            string                `xml:"NextToken,omitempty"`
	PlatformApplications []PlatformApplication `xml:"PlatformApplications,omitempty"`
}

// Response for ListSubscriptionsByTopic action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptionsByTopicResult.html]
type ListSubscriptionsByTopicResult struct {
	NextToken     string         `xml:"NextToken,omitempty"`
	Subscriptions []Subscription `xml:"Subscriptions>member,omitempty"`
}

// Response for ListSubscriptions action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptionsResult.html]
type ListSubscriptionsResult struct {
	NextToken     string         `xml:"NextToken,omitempty"`
	Subscriptions []Subscription `xml:"Subscriptions>member,omitempty"`
}

// Response for ListTopics action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListTopicsResult.html]
type ListTopicsResult struct {
	NextToken string  `xml:"NextToken,omitempty"`
	Topics    []Topic `xml:"Topics>member,omitempty"`
}

// Use the following JSON object format for the Message parameter to send
// different messages to each protocol.
// [http://docs.aws.amazon.com/sns/latest/api/API_Publish.html]
type Message struct {
	Default   string `json:"default,omitempty"`
	Email     string `json:"email,omitempty"`
	EmailJson string `json:"email-json,omitempty"`
	Http      string `json:"http,omitempty"`
	Https     string `json:"https,omitempty"`
	Sqs       string `json:"sqs,omitempty"`
}

// The user-specified message attribute value.
// [http://docs.aws.amazon.com/sns/latest/api/API_MessageAttributeValue.html]
type MessageAttributeValue struct {
	BinaryValue []byte   `name:"BinaryValue,omitempty"`
	DataType    DataType `name:"DataType"`
	StringValue string   `name:"StringValue,omitempty"`
}

// Platform application object.
// [http://docs.aws.amazon.com/sns/latest/api/API_PlatformApplication.html]
type PlatformApplication struct {
	Attributes             map[PlatformAttribute]string `xml:"Attributes>entry,omitempty"`
	PlatformApplicationArn string                       `xml:"PlatformApplicationArn,omitempty"`
}

// Response for Publish action.
// [http://docs.aws.amazon.com/sns/latest/api/API_PublishResult.html]
type PublishResult struct {
	MessageId string `xml:"MessageId"`
}

// Contains the Request ID of the response.
type ResponseMetaData struct {
	RequestId string `xml:"RequestId"`
}

// Response for Subscribe action.
// [http://docs.aws.amazon.com/sns/latest/api/API_SubscribeResult.html]
type SubscribeResult struct {
	SubscriptionArn string `xml:"SubscriptionArn"`
}

// A wrapper type for the attributes of an Amazon SNS subscription.
// [http://docs.aws.amazon.com/sns/latest/api/API_Subscription.html]
type Subscription struct {
	Endpoint        string `name:"Endpoint,omitempty"`
	Owner           string `name:"Owner,omitempty"`
	Protocol        string `name:"Protocol,omitempty"`
	SubscriptionArn string `name:"SubscriptionArn,omitempty"`
	TopicArn        string `name:"TopicArn,omitempty"`
}

// A wrapper type for the topic's Amazon Resource Name (ARN).
// To retrieve a topic's attributes, use GetTopicAttributes.
// [http://docs.aws.amazon.com/sns/latest/api/API_Topic.html]
type Topic struct {
	TopicArn string `name:"TopicArn,omitempty"`
}

/******************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/sns/latest/api/API_AddPermission.html]
type AddPermissionResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ConfirmSubscription.html]
type ConfirmSubscriptionResponse struct {
	ConfirmSubscriptionResult ConfirmSubscriptionResult `xml:"ConfirmSubscriptionResult"`
	ResponseMetaData          ResponseMetaData          `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html]
type CreatePlatformApplicationResponse struct {
	CreatePlatformApplicationResult CreatePlatformApplicationResult `xml:"CreatePlatformApplicationResult"`
	ResponseMetaData                ResponseMetaData                `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformEndpoint.html]
type CreatePlatformEndpointResponse struct {
	CreatePlatformEndpointResult CreatePlatformEndpointResult `xml:"CreatePlatformEndpointResult"`
	ResponseMetaData             ResponseMetaData             `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html]
type CreateTopicResponse struct {
	CreateTopicResult CreateTopicResult `xml:"CreateTopicResult"`
	ResponseMetaData  ResponseMetaData  `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteEndpoint.html]
type DeleteEndpointResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_DeletePlatformApplication.html]
type DeletePlatformApplicationResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteTopic.html]
type DeleteTopicResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_GetEndpointAttributes.html]
type GetEndpointAttributesResponse struct {
	GetEndpointAttributesResult GetEndpointAttributesResult `xml:"GetEndpointAttributesResult"`
	ResponseMetaData            ResponseMetaData            `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_GetPlatformApplicationAttributes.html]
type GetPlatformApplicationAttributesResponse struct {
	GetPlatformApplicationAttributesResult GetPlatformApplicationAttributesResult `xml:"GetPlatformApplicationAttributesResult"`
	ResponseMetaData                       ResponseMetaData                       `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html]
type GetSubscriptionAttributesResponse struct {
	GetSubscriptionAttributesResult GetSubscriptionAttributesResult `xml:"GetSubscriptionAttributesResult"`
	ResponseMetaData                ResponseMetaData                `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html]
type GetTopicAttributesResponse struct {
	GetTopicAttributesResult GetTopicAttributesResult `xml:"GetTopicAttributesResult"`
	ResponseMetaData         ResponseMetaData         `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ListEndpointsByPlatformApplication.html]
type ListEndpointsByPlatformApplicationResponse struct {
	ListEndpointsByPlatformApplicationResult ListEndpointsByPlatformApplicationResult `xml:"ListEndpointsByPlatformApplicationResult"`
	ResponseMetaData                         ResponseMetaData                         `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ListPlatformApplications.html]
type ListPlatformApplicationsResponse struct {
	ListPlatformApplicationsResult ListPlatformApplicationsResult `xml:"ListPlatformApplicationsResult"`
	ResponseMetaData               ResponseMetaData               `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptions.html]
type ListSubscriptionsResponse struct {
	ListSubscriptionsResult ListSubscriptionsResult `xml:"ListSubscriptionsResult"`
	ResponseMetaData        ResponseMetaData        `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptionsByTopic.html]
type ListSubscriptionsByTopicResponse struct {
	ListSubscriptionsByTopicResult ListSubscriptionsByTopicResult `xml:"ListSubscriptionsByTopicResult"`
	ResponseMetaData               ResponseMetaData               `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html]
type ListTopicsResponse struct {
	ListTopicsResult ListTopicsResult `xml:"ListTopicsResult"`
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_Publish.html]
type PublishResponse struct {
	PublishResult    PublishResult    `xml:"PublishResult"`
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_RemovePermission.html]
type RemovePermissionResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html]
type SetEndpointAttributesResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html]
type SetPlatformApplicationAttributesResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_SetSubscriptionAttributes.html]
type SetSubscriptionAttributesResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_SetTopicAttributes.html]
type SetTopicAttributesResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html]
type SubscribeResponse struct {
	SubscribeResult  SubscribeResult  `xml:"SubscribeResult"`
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/sns/latest/api/API_Unsubscribe.html]
type UnsubscribeResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}
