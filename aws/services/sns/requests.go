package sns

import (
	"encoding/json"
)

/*****************************************************************************/

// Adds a statement to a topic's access control policy, granting access
// for the specified AWS accounts to the specified actions.
// [http://docs.aws.amazon.com/sns/latest/api/API_AddPermission.html]
type AddPermissionRequest struct {
	AWSAccountId []string `name:"AWSAccountId.member.#" base:"1"`
	ActionName   []string `name:"ActionName.member.#" base:"1"`
	Label        string   `name:"Label"`
	TopicArn     string   `name:"TopicArn"`
}

// Creates a new AddPermissionRequest.
func NewAddPermissionRequest(topicArn, label, actionName, awsAccountId string) *AddPermissionRequest {
	return &AddPermissionRequest{[]string{awsAccountId}, []string{actionName}, label, topicArn}
}

// Add an Action.
func (r *AddPermissionRequest) AddAction(name, awsAccountId string) {
	r.AWSAccountId = append(r.AWSAccountId, awsAccountId)
	r.ActionName = append(r.ActionName, name)
}

/*****************************************************************************/

// Verifies an endpoint owner's intent to receive messages by validating
// the token sent to the endpoint by an earlier Subscribe action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ConfirmSubscription.html]
type ConfirmSubscriptionRequest struct {
	AuthenticateOnUnsubscribe string `name:"AuthenticateOnUnsubscribe,omitempty"`
	Token                     string `name:"Token"`
	TopicArn                  string `name:"TopicArn"`
}

// Creates a new ConfirmSubscriptionRequest.
func NewConfirmSubscriptionRequest(topicArn, token string) *ConfirmSubscriptionRequest {
	return &ConfirmSubscriptionRequest{TopicArn: topicArn, Token: token}
}

/*****************************************************************************/

// Creates a platform application object for one of the supported push notification
// services, such as APNS and GCM, to which devices and mobile apps may register.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html]
type CreatePlatformApplicationRequest struct {
	AttributeKeys   []PlatformAttribute `name:"Attributes.entry.#.key" base:"1"`
	AttributeValues []string            `name:"Attributes.entry.#.value" base:"1"`
	Name            string              `name:"Name"`
	Platform        string              `name:"Platform"`
}

// Creates a new CreatePlatformApplicationRequest.
func NewCreatePlatformApplicationRequest(name, platform string, attributeKey PlatformAttribute, attributeValue string) *CreatePlatformApplicationRequest {
	return &CreatePlatformApplicationRequest{[]PlatformAttribute{attributeKey}, []string{attributeValue}, name, platform}
}

// Add a Platform Attribute.
func (r *CreatePlatformApplicationRequest) AddAttribute(name PlatformAttribute, value string) {
	r.AttributeKeys = append(r.AttributeKeys, name)
	r.AttributeValues = append(r.AttributeValues, value)
}

/*****************************************************************************/

// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformEndpoint.html]
type CreatePlatformEndpointRequest struct {
	AttributeKeys          []PlatformAttribute `name:"Attributes.entry.#.key,omitempty" base:"1"`
	AttributeValues        []string            `name:"Attributes.entry.#.value,omitempty" base:"1"`
	CustomUserData         string              `name:"CustomUserData,omitempty"`
	PlatformApplicationArn string              `name:"PlatformApplicationArn"`
}

// Creates a new CreatePlatformEndpointRequest.
func NewCreatePlatformEndpointRequest(platformApplicationArn string) *CreatePlatformEndpointRequest {
	return &CreatePlatformEndpointRequest{PlatformApplicationArn: platformApplicationArn}
}

// Add a Platform Attribute.
func (r *CreatePlatformEndpointRequest) AddAttribute(name PlatformAttribute, value string) {
	r.AttributeKeys = append(r.AttributeKeys, name)
	r.AttributeValues = append(r.AttributeValues, value)
}

/*****************************************************************************/

// Creates a topic to which notifications can be published.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html]
type CreateTopicRequest struct {
	Name string `name:"Name"`
}

// Creates a new CreateTopicRequest.
func NewCreateTopicRequest(name string) *CreateTopicRequest {
	return &CreateTopicRequest{name}
}

/*****************************************************************************/

// Deletes the endpoint from Amazon SNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteEndpoint.html]
type DeleteEndpointRequest struct {
	EndpointArn string `name:"EndpointArn"`
}

// Creates a new DeleteEndpointRequest.
func NewDeleteEndpointRequest(endpointArn string) *DeleteEndpointRequest {
	return &DeleteEndpointRequest{endpointArn}
}

/*****************************************************************************/

// Deletes a platform application object for one of the supported push
// notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeletePlatformApplication.html]
type DeletePlatformApplicationRequest struct {
	PlatformApplicationArn string `name:"PlatformApplicationArn"`
}

// Creates a new PlatformApplicationArn.
func NewDeletePlatformApplicationRequest(platformApplicationArn string) *DeletePlatformApplicationRequest {
	return &DeletePlatformApplicationRequest{platformApplicationArn}
}

/*****************************************************************************/

// Deletes a topic and all its subscriptions. Deleting a topic might prevent
// some messages previously sent to the topic from being delivered to subscribers.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteTopic.html]
type DeleteTopicRequest struct {
	TopicArn string `name:"TopicArn"`
}

// Creates a new DeleteTopicRequest.
func NewDeleteTopicRequest(topicArn string) *DeleteTopicRequest {
	return &DeleteTopicRequest{topicArn}
}

/*****************************************************************************/

// Retrieves the endpoint attributes for a device on one of the supported push notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetEndpointAttributes.html]
type GetEndpointAttributesRequest struct {
	EndpointArn string `name:"EndpointArn"`
}

// Creates a new GetEndpointAttributesRequest.
func NewGetEndpointAttributesRequest(endpointArn string) *GetEndpointAttributesRequest {
	return &GetEndpointAttributesRequest{endpointArn}
}

/*****************************************************************************/

// Retrieves the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetPlatformApplicationAttributes.html]
type GetPlatformApplicationAttributesRequest struct {
	PlatformApplicationArn string `name:"PlatformApplicationArn"`
}

// Creates a new GetPlatformApplicationAttributesRequest.
func NewGetPlatformApplicationAttributesRequest(platformApplicationArn string) *GetPlatformApplicationAttributesRequest {
	return &GetPlatformApplicationAttributesRequest{platformApplicationArn}
}

/*****************************************************************************/

// Returns all of the properties of a subscription.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html]
type GetSubscriptionAttributesRequest struct {
	SubscriptionArn string `name:"EndpoSubscriptionArnintArn"`
}

// Creates a new GetSubscriptionAttributesRequest.
func NewGetSubscriptionAttributesRequest(endpoSubscriptionArnintArn string) *GetSubscriptionAttributesRequest {
	return &GetSubscriptionAttributesRequest{endpoSubscriptionArnintArn}
}

/*****************************************************************************/

// Returns all of the properties of a topic.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html]
type GetTopicAttributesRequest struct {
	TopicArn string `name:"TopicArn"`
}

// Creates a new GetTopicAttributesRequest.
func NewGetTopicAttributesRequest(topicArn string) *GetTopicAttributesRequest {
	return &GetTopicAttributesRequest{topicArn}
}

/*****************************************************************************/

// Lists the endpoints and endpoint attributes for devices in a supported push notification
// service, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListEndpointsByPlatformApplication.html]
type ListEndpointsByPlatformApplicationRequest struct {
	NextToken              string `name:"NextToken,omitempty"`
	PlatformApplicationArn string `name:"PlatformApplicationArn"`
}

// Creates a new ListEndpointsByPlatformApplicationRequest.
func NewListEndpointsByPlatformApplicationRequest(platformApplicationArn string) *ListEndpointsByPlatformApplicationRequest {
	return &ListEndpointsByPlatformApplicationRequest{PlatformApplicationArn: platformApplicationArn}
}

/*****************************************************************************/

// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListPlatformApplications.html]
type ListPlatformApplicationsRequest struct {
	NextToken string `name:"NextToken,omitempty"`
}

// Creates a new ListPlatformApplicationsRequest.
func NewListPlatformApplicationsRequest() *ListPlatformApplicationsRequest {
	return &ListPlatformApplicationsRequest{}
}

/*****************************************************************************/

// Returns a list of the requester's subscriptions.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptions.html]
type ListSubscriptionsRequest struct {
	NextToken string `name:"NextToken,omitempty"`
}

// Creates a new ListSubscriptionsRequest.
func NewListSubscriptionsRequest() *ListSubscriptionsRequest {
	return &ListSubscriptionsRequest{}
}

/*****************************************************************************/

// Returns a list of the subscriptions to a specific topic.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptionsByTopic.html]
type ListSubscriptionsByTopicRequest struct {
	NextToken string `name:"NextToken,omitempty"`
	TopicArn  string `name:"TopicArn"`
}

// Creates a new ListSubscriptionsByTopicRequest.
func NewListSubscriptionsByTopicRequest(topicArn string) *ListSubscriptionsByTopicRequest {
	return &ListSubscriptionsByTopicRequest{TopicArn: topicArn}
}

/*****************************************************************************/

// Returns a list of the requester's topics.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html]
type ListTopicsRequest struct {
	NextToken string `name:"NextToken,omitempty"`
}

// Creates a new ListTopicsRequest.
func NewListTopicsRequest() *ListTopicsRequest {
	return &ListTopicsRequest{}
}

/*****************************************************************************/

// Sends a message to all of a topic's subscribed endpoints.
// [http://docs.aws.amazon.com/sns/latest/api/API_Publish.html]
type PublishRequest struct {
	Message           string                  `name:"Message"`
	MessageAttributes []MessageAttributeValue `name:"MessageAttributes.entry.#.,omitempty" base:"1"`
	MessageStructure  string                  `name:"MessageStructure,omitempty"`
	Subject           string                  `name:"Subject,omitempty"`
	TargetArn         string                  `name:"TargetArn,omitempty"`
	TopicArn          string                  `name:"TopicArn,omitempty"`
}

// Creates a new PublishRequest.
func NewPublishRequest(message string) *PublishRequest {
	return &PublishRequest{Message: message}
}

// Creates a new PublishRequest with message structure as JSON.
func NewPublishWithJsonMessageRequest(message Message) *PublishRequest {
	jsonMsg, _ := json.Marshal(message)
	return &PublishRequest{Message: string(jsonMsg), MessageStructure: "json"}
}

/*****************************************************************************/

// Removes a statement from a topic's access control policy.
// [http://docs.aws.amazon.com/sns/latest/api/API_RemovePermission.html]
type RemovePermissionRequest struct {
	Label    string `name:"Label"`
	TopicArn string `name:"TopicArn"`
}

// Creates a new RemovePermissionRequest.
func NewRemovePermissionRequest(topicArn, label string) *RemovePermissionRequest {
	return &RemovePermissionRequest{label, topicArn}
}

/*****************************************************************************/

// Sets the attributes for an endpoint for a device on one of the supported
// push notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html]
type SetEndpointAttributesRequest struct {
	Attributes  []EndpointAttribute `name:"Attributes.entry.#.value" base:"1"`
	EndpointArn string              `name:"EndpointArn"`
}

// Creates a new SetEndpointAttributesRequest.
func NewSetEndpointAttributesRequest(endpointArn string, attributes ...EndpointAttribute) *SetEndpointAttributesRequest {
	return &SetEndpointAttributesRequest{attributes, endpointArn}
}

/*****************************************************************************/

// Sets the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html]
type SetPlatformApplicationAttributesRequest struct {
	Attributes             []PlatformAttribute `name:"PlatformAttributes.entry.#.value" base:"1"`
	PlatformApplicationArn string              `name:"EndpointArn"`
}

// Creates a new SetPlatformApplicationAttributesRequest.
func NewSetPlatformApplicationAttributesRequest(platformApplicationArn string, attributes ...PlatformAttribute) *SetPlatformApplicationAttributesRequest {
	return &SetPlatformApplicationAttributesRequest{attributes, platformApplicationArn}
}

/*****************************************************************************/

// Allows a subscription owner to set an attribute of the topic to a new value.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetSubscriptionAttributes.html]
type SetSubscriptionAttributesRequest struct {
	AttributeName   string `name:"AttributeName"`
	AttributeValue  string `name:"AttributeValue,omitempty"`
	SubscriptionArn string `name:"SubscriptionArn"`
}

// Creates a new SetSubscriptionAttributesRequest.
func NewSetSubscriptionAttributesRequest(subscriptionArn, attributeName, attributeValue string) *SetSubscriptionAttributesRequest {
	return &SetSubscriptionAttributesRequest{attributeName, attributeValue, subscriptionArn}
}

/*****************************************************************************/

// Allows a topic owner to set an attribute of the topic to a new value.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetTopicAttributes.html]
type SetTopicAttributesRequest struct {
	AttributeName  string `name:"AttributeName"`
	AttributeValue string `name:"AttributeValue,omitempty"`
	TopicArn       string `name:"TopicArn"`
}

// Creates a new SetTopicAttributesRequest.
func NewSetTopicAttributesRequest(topicArn, attributeName string) *SetTopicAttributesRequest {
	return &SetTopicAttributesRequest{AttributeName: attributeName, TopicArn: topicArn}
}

/*****************************************************************************/

// Prepares to subscribe an endpoint by sending the endpoint a confirmation message.
// [http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html]
type SubscribeRequest struct {
	Endpoint string `name:"Endpoint,omitempty"`
	Protocol string `name:"Protocol"`
	TopicArn string `name:"TopicArn"`
}

// Creates a new SubscribeRequest.
func NewSubscribeRequest(topicArn, protocol string) *SubscribeRequest {
	return &SubscribeRequest{Protocol: protocol, TopicArn: topicArn}
}

/*****************************************************************************/

// Deletes a subscription.
// [http://docs.aws.amazon.com/sns/latest/api/API_Unsubscribe.html]
type UnsubscribeRequest struct {
	SubscriptionArn string `name:"SubscriptionArn"`
}

// Creates a new UnsubscribeRequest.
func NewUnsubscribeRequest(subscriptionArn string) *UnsubscribeRequest {
	return &UnsubscribeRequest{subscriptionArn}
}
