//
// Amazon Simple Notification Service (Amazon SNS) is a web service that enables applications,
// end-users, and devices to instantly send and receive notifications from the cloud.
//
// [http://aws.amazon.com/documentation/sns/]
//
package sns

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"github.com/twhello/aws-to-go/util/netutil"
	"net/http"
	"time"
)

const ServiceName = "sns"

// The SNS Service object. Use sns.NewService().
type SNSService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *SNSService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *SNSService) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *SNSService) Endpoint() string {
	return s.endpoint
}

// Low-level request to SNS service.
func (s *SNSService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *SNSService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	qs := netutil.MarshalValues(request)
	qs.Add("AWSAccessKeyId", s.cred.AccessKeyId())
	qs.Add("Timestamp", time.Now().UTC().Format(time.RFC3339)) // ISO 8601 2006-01-02T15:04:05.999Z
	qs.Add("Version", "2010-03-31")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", s.Endpoint(), qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * SNS Service Methods.
 */

// Adds a statement to a topic's access control policy, granting access
// for the specified AWS accounts to the specified actions.
// [http://docs.aws.amazon.com/sns/latest/api/API_AddPermission.html]
func (s *SNSService) AddPermission(req *AddPermissionRequest) (result *AddPermissionResponse, err error) {

	result = new(AddPermissionResponse)
	err = s.wrapperSignAndDo("AddPermission", req, result)
	return
}

// Verifies an endpoint owner's intent to receive messages by validating
// the token sent to the endpoint by an earlier Subscribe action.
// [http://docs.aws.amazon.com/sns/latest/api/API_ConfirmSubscription.html]
func (s *SNSService) ConfirmSubscription(req *ConfirmSubscriptionRequest) (result *ConfirmSubscriptionResponse, err error) {

	result = new(ConfirmSubscriptionResponse)
	err = s.wrapperSignAndDo("ConfirmSubscription", req, result)
	return
}

// Creates a platform application object for one of the supported push notification
// services, such as APNS and GCM, to which devices and mobile apps may register.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html]
func (s *SNSService) CreatePlatformApplication(req *CreatePlatformApplicationRequest) (result *CreatePlatformApplicationResponse, err error) {

	result = new(CreatePlatformApplicationResponse)
	err = s.wrapperSignAndDo("CreatePlatformApplication", req, result)
	return
}

// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformEndpoint.html]
func (s *SNSService) CreatePlatformEndpoint(req *CreatePlatformEndpointRequest) (result *CreatePlatformEndpointResponse, err error) {

	result = new(CreatePlatformEndpointResponse)
	err = s.wrapperSignAndDo("CreatePlatformEndpoint", req, result)
	return
}

// Creates a topic to which notifications can be published.
// [http://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html]
func (s *SNSService) CreateTopic(req *CreateTopicRequest) (result *CreateTopicResponse, err error) {

	result = new(CreateTopicResponse)
	err = s.wrapperSignAndDo("CreateTopic", req, result)
	return
}

// Deletes the endpoint from Amazon SNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteEndpoint.html]
func (s *SNSService) DeleteEndpoint(req *DeleteEndpointRequest) (result *DeleteEndpointResponse, err error) {

	result = new(DeleteEndpointResponse)
	err = s.wrapperSignAndDo("DeleteEndpoint", req, result)
	return
}

// Deletes a platform application object for one of the supported push
// notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeletePlatformApplication.html]
func (s *SNSService) DeletePlatformApplication(req *DeletePlatformApplicationRequest) (result *DeletePlatformApplicationResponse, err error) {

	result = new(DeletePlatformApplicationResponse)
	err = s.wrapperSignAndDo("DeletePlatformApplication", req, result)
	return
}

// Deletes a topic and all its subscriptions. Deleting a topic might prevent
// some messages previously sent to the topic from being delivered to subscribers.
// [http://docs.aws.amazon.com/sns/latest/api/API_DeleteTopic.html]
func (s *SNSService) DeleteTopic(req *DeleteTopicRequest) (result *DeleteTopicResponse, err error) {

	result = new(DeleteTopicResponse)
	err = s.wrapperSignAndDo("DeleteTopic", req, result)
	return
}

// Retrieves the endpoint attributes for a device on one of the supported push notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetEndpointAttributes.html]
func (s *SNSService) GetEndpointAttributes(req *GetEndpointAttributesRequest) (result *GetEndpointAttributesResponse, err error) {

	result = new(GetEndpointAttributesResponse)
	err = s.wrapperSignAndDo("GetEndpointAttributes", req, result)
	return
}

// Retrieves the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetPlatformApplicationAttributes.html]
func (s *SNSService) GetPlatformApplicationAttributes(req *GetPlatformApplicationAttributesRequest) (result *GetPlatformApplicationAttributesResponse, err error) {

	result = new(GetPlatformApplicationAttributesResponse)
	err = s.wrapperSignAndDo("GetPlatformApplicationAttributes", req, result)
	return
}

// Returns all of the properties of a subscription.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html]
func (s *SNSService) GetSubscriptionAttributes(req *GetSubscriptionAttributesRequest) (result *GetSubscriptionAttributesResponse, err error) {

	result = new(GetSubscriptionAttributesResponse)
	err = s.wrapperSignAndDo("GetSubscriptionAttributes", req, result)
	return
}

// Returns all of the properties of a topic.
// [http://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html]
func (s *SNSService) GetTopicAttributes(req *GetTopicAttributesRequest) (result *GetTopicAttributesResponse, err error) {

	result = new(GetTopicAttributesResponse)
	err = s.wrapperSignAndDo("GetTopicAttributes", req, result)
	return
}

// Lists the endpoints and endpoint attributes for devices in a supported push notification
// service, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListEndpointsByPlatformApplication.html]
func (s *SNSService) ListEndpointsByPlatformApplication(req *ListEndpointsByPlatformApplicationRequest) (result *ListEndpointsByPlatformApplicationResponse, err error) {

	result = new(ListEndpointsByPlatformApplicationResponse)
	err = s.wrapperSignAndDo("ListEndpointsByPlatformApplication", req, result)
	return
}

// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListPlatformApplications.html]
func (s *SNSService) ListPlatformApplications(req *ListPlatformApplicationsRequest) (result *ListPlatformApplicationsResponse, err error) {

	result = new(ListPlatformApplicationsResponse)
	err = s.wrapperSignAndDo("ListPlatformApplications", req, result)
	return
}

// Returns a list of the requester's subscriptions.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptions.html]
func (s *SNSService) ListSubscriptions(req *ListSubscriptionsRequest) (result *ListSubscriptionsResponse, err error) {

	result = new(ListSubscriptionsResponse)
	err = s.wrapperSignAndDo("ListSubscriptions", req, result)
	return
}

// Returns a list of the subscriptions to a specific topic.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListSubscriptionsByTopic.html]
func (s *SNSService) ListSubscriptionsByTopic(req *ListSubscriptionsByTopicRequest) (result *ListSubscriptionsByTopicResponse, err error) {

	result = new(ListSubscriptionsByTopicResponse)
	err = s.wrapperSignAndDo("ListSubscriptionsByTopic", req, result)
	return
}

// Returns a list of the requester's topics.
// [http://docs.aws.amazon.com/sns/latest/api/API_ListTopics.html]
func (s *SNSService) ListTopics(req *ListTopicsRequest) (result *ListTopicsResponse, err error) {

	result = new(ListTopicsResponse)
	err = s.wrapperSignAndDo("ListTopics", req, result)
	return
}

// Sends a message to all of a topic's subscribed endpoints.
// [http://docs.aws.amazon.com/sns/latest/api/API_Publish.html]
func (s *SNSService) Publish(req *PublishRequest) (result *PublishResponse, err error) {

	post, err := services.NewServerRequest("POST", s.Endpoint(), nil)
	if err == nil {
		qs := post.FormPostValues()
		qs.Add("AWSAccessKeyId", s.cred.AccessKeyId())
		qs.Add("Timestamp", time.Now().UTC().Format(time.RFC3339)) // ISO 8601 2006-01-02T15:04:05.999Z
		qs.Add("Version", "2010-03-31")
		qs.Add("Action", "Publish")
		netutil.MergeValues(qs, netutil.MarshalValues(req))
	
		result = new(PublishResponse)
		_, err = s.SignAndDo(post, result)
	}	
	
	return
}

// Removes a statement from a topic's access control policy.
// [http://docs.aws.amazon.com/sns/latest/api/API_RemovePermission.html]
func (s *SNSService) RemovePermission(req *RemovePermissionRequest) (result *RemovePermissionResponse, err error) {

	result = new(RemovePermissionResponse)
	err = s.wrapperSignAndDo("RemovePermission", req, result)
	return
}

// Sets the attributes for an endpoint for a device on one of the supported
// push notification services, such as GCM and APNS.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html]
func (s *SNSService) SetEndpointAttributes(req *SetEndpointAttributesRequest) (result *SetEndpointAttributesResponse, err error) {

	result = new(SetEndpointAttributesResponse)
	err = s.wrapperSignAndDo("SetEndpointAttributes", req, result)
	return
}

// Sets the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetPlatformApplicationAttributes.html]
func (s *SNSService) SetPlatformApplicationAttributes(req *SetPlatformApplicationAttributesRequest) (result *SetPlatformApplicationAttributesResponse, err error) {

	result = new(SetPlatformApplicationAttributesResponse)
	err = s.wrapperSignAndDo("SetPlatformApplicationAttributes", req, result)
	return
}

// Allows a subscription owner to set an attribute of the topic to a new value.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetSubscriptionAttributes.html]
func (s *SNSService) SetSubscriptionAttributes(req *SetSubscriptionAttributesRequest) (result *SetSubscriptionAttributesResponse, err error) {

	result = new(SetSubscriptionAttributesResponse)
	err = s.wrapperSignAndDo("SetSubscriptionAttributes", req, result)
	return
}

// Allows a topic owner to set an attribute of the topic to a new value.
// [http://docs.aws.amazon.com/sns/latest/api/API_SetTopicAttributes.html]
func (s *SNSService) SetTopicAttributes(req *SetTopicAttributesRequest) (result *SetTopicAttributesResponse, err error) {

	result = new(SetTopicAttributesResponse)
	err = s.wrapperSignAndDo("SetTopicAttributes", req, result)
	return
}

// Prepares to subscribe an endpoint by sending the endpoint a confirmation message.
// [http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html]
func (s *SNSService) Subscribe(req *SubscribeRequest) (result *SubscribeResponse, err error) {

	result = new(SubscribeResponse)
	err = s.wrapperSignAndDo("Subscribe", req, result)
	return
}

// Deletes a subscription.
// [http://docs.aws.amazon.com/sns/latest/api/API_Unsubscribe.html]
func (s *SNSService) Unsubscribe(req *UnsubscribeRequest) (result *UnsubscribeResponse, err error) {

	result = new(UnsubscribeResponse)
	err = s.wrapperSignAndDo("Unsubscribe", req, result)
	return
}

// Creates a new SNS Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *SNSService {

	subdomain := ServiceName
	if region.Name() != regions.DEFAULT_REGION {
		subdomain += "." + region.Name()
	}
	return &SNSService{cred, region, "https://" + subdomain + ".amazonaws.com"}
}
