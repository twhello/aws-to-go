//
// Amazon SES is an outbound-only email-sending service that provides an easy,
// cost-effective way for you to send email.
//
// [http://aws.amazon.com/documentation/ses/]
//
package ses

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"github.com/twhello/aws-to-go/util/netutil"
	"net/http"
	"time"
)

const ServiceName = "ses"

// The SES Service object. Use ses.NewService().
type SESService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *SESService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *SESService) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *SESService) Endpoint() string {
	return s.endpoint
}

// Low-level request to SES service.
func (s *SESService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *SESService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	vals := req.FormPostValues()
	vals.Add("AWSAccessKeyId", s.cred.AccessKeyId())
	vals.Add("Timestamp", time.Now().UTC().Format(time.RFC3339)) // ISO 8601 2006-01-02T15:04:05.999Z
	vals.Add("Version", "2010-12-01")
	vals.Add("Action", action)

	if request != nil {
		netutil.MergeValues(vals, netutil.MarshalValues(request))
	}

	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * SES Service Methods.
 */

// Deletes the specified identity (email address or domain) from the list of verified identities.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteIdentity.html]
func (s *SESService) DeleteIdentity(req *DeleteIdentityRequest) (result *DeleteIdentityResponse, err error) {

	result = new(DeleteIdentityResponse)
	err = s.wrapperSignAndDo("DeleteIdentity", req, result)
	return
}

// Deletes the specified email address from the list of verified addresses.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteVerifiedEmailAddress.html]
func (s *SESService) DeleteVerifiedEmailAddress(req *DeleteVerifiedEmailAddressRequest) (result *DeleteVerifiedEmailAddressResponse, err error) {

	result = new(DeleteVerifiedEmailAddressResponse)
	err = s.wrapperSignAndDo("DeleteVerifiedEmailAddress", req, result)
	return
}

// Returns the current status of Easy DKIM signing for an entity. For domain name identities,
// this action also returns the DKIM tokens that are required for Easy DKIM signing, and
// whether Amazon SES has successfully verified that these tokens have been published.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityDkimAttributes.html]
func (s *SESService) GetIdentityDkimAttributes(req *GetIdentityDkimAttributesRequest) (result *GetIdentityDkimAttributesResponse, err error) {

	result = new(GetIdentityDkimAttributesResponse)
	err = s.wrapperSignAndDo("GetIdentityDkimAttributes", req, result)
	return
}

// Given a list of verified identities (email addresses and/or domains), returns a
// structure describing identity notification attributes.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityNotificationAttributes.html]
func (s *SESService) GetIdentityNotificationAttributes(req *GetIdentityNotificationAttributesRequest) (result *GetIdentityNotificationAttributesResponse, err error) {

	result = new(GetIdentityNotificationAttributesResponse)
	err = s.wrapperSignAndDo("GetIdentityNotificationAttributes", req, result)
	return
}

// Given a list of identities (email addresses and/or domains), returns the verification
// status and (for domain identities) the verification token for each identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityVerificationAttributes.html]
func (s *SESService) GetIdentityVerificationAttributes(req *GetIdentityVerificationAttributesRequest) (result *GetIdentityVerificationAttributesResponse, err error) {

	result = new(GetIdentityVerificationAttributesResponse)
	err = s.wrapperSignAndDo("GetIdentityVerificationAttributes", req, result)
	return
}

// Returns the user's current sending limits.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendQuota.html]
func (s *SESService) GetSendQuota() (result *GetSendQuotaResponse, err error) {

	result = new(GetSendQuotaResponse)
	err = s.wrapperSignAndDo("GetSendQuota", nil, result)
	return
}

// Returns the user's sending statistics. The result is a list of data points,
// representing the last two weeks of sending activity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendStatistics.html]
func (s *SESService) GetSendStatistics() (result *GetSendStatisticsResponse, err error) {

	result = new(GetSendStatisticsResponse)
	err = s.wrapperSignAndDo("GetSendStatistics", nil, result)
	return
}

// Returns a list containing all of the identities (email addresses and domains) for a specific AWS Account, regardless of verification status.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListIdentities.html]
func (s *SESService) ListIdentities(req *ListIdentitiesRequest) (result *ListIdentitiesResponse, err error) {

	result = new(ListIdentitiesResponse)
	err = s.wrapperSignAndDo("ListIdentities", req, result)
	return
}

// Returns a list containing all of the email addresses that have been verified.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListVerifiedEmailAddresses.html]
func (s *SESService) ListVerifiedEmailAddresses() (result *ListVerifiedEmailAddressesResponse, err error) {

	result = new(ListVerifiedEmailAddressesResponse)
	err = s.wrapperSignAndDo("ListVerifiedEmailAddresses", nil, result)
	return
}

// Composes an email message based on input data, and then immediately queues the message for sending.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendEmail.html]
func (s *SESService) SendEmail(req *SendEmailRequest) (result *SendEmailResponse, err error) {

	result = new(SendEmailResponse)
	err = s.wrapperSignAndDo("SendEmail", req, result)
	return
}

// Sends an email message, with header and content specified by the client. The SendRawEmail
// action is useful for sending multipart MIME emails. The raw text of the message must
// comply with Internet email standards; otherwise, the message cannot be sent.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendRawEmail.html]
func (s *SESService) SendRawEmail(req *SendRawEmailRequest) (result *SendRawEmailResponse, err error) {

	result = new(SendRawEmailResponse)
	err = s.wrapperSignAndDo("SendRawEmail", req, result)
	return
}

// Enables or disables Easy DKIM signing of email sent from an identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityDkimEnabled.html]
func (s *SESService) SetIdentityDkimEnabled(req *SetIdentityDkimEnabledRequest) (result *SetIdentityDkimEnabledResponse, err error) {

	result = new(SetIdentityDkimEnabledResponse)
	err = s.wrapperSignAndDo("SetIdentityDkimEnabled", req, result)
	return
}

// Given an identity (email address or domain), enables or disables whether Amazon SES
// forwards bounce and complaint notifications as email. Feedback forwarding can only
// be disabled when Amazon Simple Notification Service (Amazon SNS) topics are specified
// for both bounces and complaints.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityFeedbackForwardingEnabled.html]
func (s *SESService) SetIdentityFeedbackForwardingEnabled(req *SetIdentityFeedbackForwardingEnabledRequest) (result *SetIdentityFeedbackForwardingEnabledResponse, err error) {

	result = new(SetIdentityFeedbackForwardingEnabledResponse)
	err = s.wrapperSignAndDo("SetIdentityFeedbackForwardingEnabled", req, result)
	return
}

// Given an identity (email address or domain), sets the Amazon Simple Notification Service (Amazon SNS)
// topic to which Amazon SES will publish bounce, complaint, and/or delivery notifications
// for emails sent with that identity as the Source.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityNotificationTopic.html]
func (s *SESService) SetIdentityNotificationTopic(req *SetIdentityNotificationTopicRequest) (result *SetIdentityNotificationTopicResponse, err error) {

	result = new(SetIdentityNotificationTopicResponse)
	err = s.wrapperSignAndDo("SetIdentityNotificationTopic", req, result)
	return
}

// Returns a set of DKIM tokens for a domain.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainDkim.html]
func (s *SESService) VerifyDomainDkim(req *VerifyDomainDkimRequest) (result *VerifyDomainDkimResponse, err error) {

	result = new(VerifyDomainDkimResponse)
	err = s.wrapperSignAndDo("VerifyDomainDkim", req, result)
	return
}

// Verifies a domain.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainIdentity.html]
func (s *SESService) VerifyDomainIdentity(req *VerifyDomainIdentityRequest) (result *VerifyDomainIdentityResponse, err error) {

	result = new(VerifyDomainIdentityResponse)
	err = s.wrapperSignAndDo("VerifyDomainIdentity", req, result)
	return
}

// Verifies an email address. This action causes a confirmation email message to be sent
// to the specified address.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailAddress.html]
func (s *SESService) VerifyEmailAddress(req *VerifyEmailAddressRequest) (result *VerifyEmailAddressResponse, err error) {

	result = new(VerifyEmailAddressResponse)
	err = s.wrapperSignAndDo("VerifyEmailAddress", req, result)
	return
}

// Verifies an email address. This action causes a confirmation email message to be
// sent to the specified address.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailIdentity.html]
func (s *SESService) VerifyEmailIdentity(req *VerifyEmailIdentityRequest) (result *VerifyEmailIdentityResponse, err error) {

	result = new(VerifyEmailIdentityResponse)
	err = s.wrapperSignAndDo("VerifyEmailIdentity", req, result)
	return
}

// Creates a new SES Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *SESService {

	subdomain := ServiceName
	if region.Name() != regions.DEFAULT_REGION {
		subdomain += "-" + region.Name()
	}
	return &SESService{cred, region, "https://" + subdomain + ".amazonaws.com"}
}
