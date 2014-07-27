package ses

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/******************************************************************************
 * Constants
 */

// The type of the identities to list.
// EMAIL_ADDRESS = "EmailAddress"
// DOMAIN = "Domain"
type IdentityType string

const (
	EMAIL_ADDRESS IdentityType = "EmailAddress"
	DOMAIN        IdentityType = "Domain"
)

// The type of notifications that will be published to the specified Amazon SNS topic.
// BOUNCE = "Bounce"
// COMPLAINT = "Complaint"
// DELIVERY = "Delivery"
type NotificationType string

const (
	BOUNCE    NotificationType = "Bounce"
	COMPLAINT NotificationType = "Complaint"
	DELIVERY  NotificationType = "Delivery"
)

// Describes whether Amazon SES has successfully verified the DKIM DNS records
// (tokens) published in the domain name's DNS.
// PENDING = "Pending"
// SUCCESS = "Success"
// FAILED = "Failed"
// TEMP_FAILURE = "TemporaryFailure"
// NOT_STARTED = "NotStarted"
type VerificationStatus string

const (
	PENDING      VerificationStatus = "Pending"
	SUCCESS      VerificationStatus = "Success"
	FAILED       VerificationStatus = "Failed"
	TEMP_FAILURE VerificationStatus = "TemporaryFailure"
	NOT_STARTED  VerificationStatus = "NotStarted"
)

/******************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteIdentity.html]
type DeleteIdentityResponse struct {
	DeleteIdentityResult DeleteIdentityResult `xml:"DeleteIdentityResult"`
	ResponseMetaData     ResponseMetaData     `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteVerifiedEmailAddress.html]
type DeleteVerifiedEmailAddressResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityDkimAttributes.html]
type GetIdentityDkimAttributesResponse struct {
	GetIdentityDkimAttributesResult GetIdentityDkimAttributesResult `xml:"GetIdentityDkimAttributesResult"`
	ResponseMetaData                ResponseMetaData                `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityNotificationAttributes.html]
type GetIdentityNotificationAttributesResponse struct {
	GetIdentityNotificationAttributesResult GetIdentityNotificationAttributesResult `xml:"GetIdentityNotificationAttributesResult"`
	ResponseMetaData                        ResponseMetaData                        `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityVerificationAttributes.html]
type GetIdentityVerificationAttributesResponse struct {
	GetIdentityVerificationAttributesResult GetIdentityVerificationAttributesResult `xml:"GetIdentityVerificationAttributesResult"`
	ResponseMetaData                        ResponseMetaData                        `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendQuota.html]
type GetSendQuotaResponse struct {
	GetSendQuotaResult GetSendQuotaResult `xml:"GetSendQuotaResult"`
	ResponseMetaData   ResponseMetaData   `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendStatistics.html]
type GetSendStatisticsResponse struct {
	GetSendStatisticsResult GetSendStatisticsResult `xml:"GetSendStatisticsResult"`
	ResponseMetaData        ResponseMetaData        `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListIdentities.html]
type ListIdentitiesResponse struct {
	ListIdentitiesResult ListIdentitiesResult `xml:"ListIdentitiesResult"`
	ResponseMetaData     ResponseMetaData     `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListVerifiedEmailAddresses.html]
type ListVerifiedEmailAddressesResponse struct {
	ListVerifiedEmailAddressesResult ListVerifiedEmailAddressesResult `xml:"ListVerifiedEmailAddressesResult"`
	ResponseMetaData                 ResponseMetaData                 `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendEmail.html]
type SendEmailResponse struct {
	SendEmailResult  SendEmailResult  `xml:"SendEmailResult"`
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendRawEmail.html]
type SendRawEmailResponse struct {
	SendRawEmailResult SendRawEmailResult `xml:"SendRawEmailResult"`
	ResponseMetaData   ResponseMetaData   `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityDkimEnabled.html]
type SetIdentityDkimEnabledResponse struct {
	SetIdentityDkimEnabledResult SetIdentityDkimEnabledResult `xml:"SetIdentityDkimEnabledResult"`
	ResponseMetaData             ResponseMetaData             `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityFeedbackForwardingEnabled.html]
type SetIdentityFeedbackForwardingEnabledResponse struct {
	SetIdentityFeedbackForwardingEnabledResult SetIdentityFeedbackForwardingEnabledResult `xml:"SetIdentityFeedbackForwardingEnabledResult"`
	ResponseMetaData                           ResponseMetaData                           `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityNotificationTopic.html]
type SetIdentityNotificationTopicResponse struct {
	SetIdentityNotificationTopicResult SetIdentityNotificationTopicResult `xml:"SetIdentityNotificationTopicResult"`
	ResponseMetaData                   ResponseMetaData                   `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainDkim.html]
type VerifyDomainDkimResponse struct {
	VerifyDomainDkimResult VerifyDomainDkimResult `xml:"VerifyDomainDkimResult"`
	ResponseMetaData       ResponseMetaData       `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainIdentity.html]
type VerifyDomainIdentityResponse struct {
	VerifyDomainIdentityResult VerifyDomainIdentityResult `xml:"VerifyDomainIdentityResult"`
	ResponseMetaData           ResponseMetaData           `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailAddress.html]
type VerifyEmailAddressResponse struct {
	ResponseMetaData ResponseMetaData `xml:"ResponseMetaData"`
}

// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailIdentity.html]
type VerifyEmailIdentityResponse struct {
	VerifyEmailIdentityResult VerifyEmailIdentityResult `xml:"VerifyEmailIdentityResult"`
	ResponseMetaData          ResponseMetaData          `xml:"ResponseMetaData"`
}

/******************************************************************************
 * Data Types
 */

// Represents the body of the message. You can specify text, HTML, or both. If you use both, then
// the message should display correctly in the widest variety of email clients.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_Body.html]
type Body struct {
	Html Content `xml:"Html,omitempty" name:"Html.,omitempty"`
	Text Content `xml:"Text,omitempty" name:"Text.,omitempty"`
}

// Represents textual data, plus an optional character set specification.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_Content.html]
type Content struct {
	CharSet string `xml:"CharSet,omitempty" name:"CharSet,omitempty"`
	Data    string `xml:"Data" name:"Data"`
}

// Represents the destination of the message, consisting of To:, CC:, and BCC: fields.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_Destination.html]
type Destination struct {
	BccAddresses []string `xml:"BccAddresses,omitempty" name:"BccAddress.member.#,omitempty"`
	CCAddresses  []string `xml:"CCAddresses,omitempty" name:"CCAddresses.member.#,omitempty"`
	ToAddresses  []string `xml:"ToAddresses,omitempty" name:"ToAddresses.member.#,omitempty"`
}

// Represents the DKIM attributes of a verified email address or a domain.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_IdentityDkimAttributes.html]
type IdentityDkimAttributes struct {
	DkimEnabled            bool               `xml:"DkimEnabled"`
	DkimTokens             []string           `xml:"DkimTokens,omitempty"`
	DkimVerificationStatus VerificationStatus `xml:"DkimVerificationStatus"`
}

// Represents the notification attributes of an identity, including whether an identity
// has Amazon Simple Notification Service (Amazon SNS) topics set for bounce, complaint,
// and/or delivery notifications, and whether feedback forwarding is enabled for bounce
// and complaint notifications.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_IdentityNotificationAttributes.html]
type IdentityNotificationAttributes struct {
	BounceTopic       string `xml:"BounceTopic"`
	ComplaintTopic    string `xml:"ComplaintTopic"`
	DeliveryTopic     string `xml:"DeliveryTopic"`
	ForwardingEnabled bool   `xml:"ForwardingEnabled"`
}

// Represents the verification attributes of a single identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_IdentityVerificationAttributes.html]
type IdentityVerificationAttributes struct {
	VerificationStatus VerificationStatus `xml:"BounceTopic"`
	VerificationToken  string             `xml:"BounceTopic,omitstatus"`
}

// Represents the message to be sent, composed of a subject and a body.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_Message.html]
type Message struct {
	Body    Body    `xml:"Body" name:"Body."`
	Subject Content `xml:"Subject" name:"Subject."`
}

// Represents the raw data of the message.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_RawMessage.html]
type RawMessage struct {
	Data string `xml:"Data" name:"Data"`
}

// Contains the Request ID of the response.
type ResponseMetaData struct {
	RequestId string `xml:"RequestId"`
}

// Represents sending statistics data. Each SendDataPoint contains statistics for a
// 15-minute period of sending activity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendDataPoint.html]
type SendDataPoint struct {
	Bounces          int64             `xml:"Bounces,omitempty"`
	Complaints       int64             `xml:"Complaints,omitempty"`
	DeliveryAttempts int64             `xml:"DeliveryAttempts,omitempty"`
	Rejects          int64             `xml:"Rejects,omitempty"`
	Timestamp        datetime.JsonDate `xml:"Timestamp,omitempty"`
}

/*****************************************************************************
 * Result Data Types
 */

// An empty element. Receiving this element indicates that the request completed successfully.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteIdentityResult.html]
type DeleteIdentityResult struct{}

// Represents a list of all the DKIM attributes for the specified identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityDkimAttributesResult.html]
type GetIdentityDkimAttributesResult struct {
	DkimAttributes map[string]IdentityDkimAttributes `xml:DkimAttributes`
}

// Describes whether an identity has Amazon Simple Notification Service (Amazon SNS) topics set for bounce,
// complaint, and/or delivery notifications, and specifies whether feedback forwarding is enabled for
// bounce and complaint notifications.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityNotificationAttributesResult.html]
type GetIdentityNotificationAttributesResult struct {
	NotificationAttributes map[string]IdentityNotificationAttributes `xml:"NotificationAttributes"`
}

// Represents the verification attributes for a list of identities.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityVerificationAttributesResult.html]
type GetIdentityVerificationAttributesResult struct {
	VerificationAttributes map[string]IdentityNotificationAttributes `xml:"VerificationAttributes"`
}

// Represents the user's current activity limits returned from a successful GetSendQuota request.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendQuotaResult.html]
type GetSendQuotaResult struct {
	Max24HourSend   float64 `xml:"Max24HourSend"`
	MaxSendRate     float64 `xml:"MaxSendRate"`
	SentLast24Hours float64 `xml:"SentLast24Hours"`
}

// Represents a list of SendDataPoint items returned from a successful GetSendStatistics request.
// This list contains aggregated data from the previous two weeks of sending activity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetSendStatisticsResult.html]
type GetSendStatisticsResult struct {
	SendDataPoints []SendDataPoint `xml:"SendDataPoints,omitempty"`
}

// Represents a list of all verified identities for the AWS Account.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListIdentitiesResult.html]
type ListIdentitiesResult struct {
	Identities []string `xml:"Identities"`
	NextToken  string   `xml:"NextToken,omitempty"`
}

// Represents a list of all the email addresses verified for the current user.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListVerifiedEmailAddressesResult.html]
type ListVerifiedEmailAddressesResult struct {
	VerifiedEmailAddresses []string `xml:"VerifiedEmailAddresses"`
}

// Represents a unique message ID returned from a successful SendEmail request.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendEmailResult.html]
type SendEmailResult struct {
	MessageId string `xml:"MessageId"`
}

// Represents a unique message ID returned from a successful SendRawEmail request.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendRawEmailResult.html]
type SendRawEmailResult struct {
	MessageId string `xml:"MessageId"`
}

// An empty element. Receiving this element indicates that the request completed successfully.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityDkimEnabledResult.html]
type SetIdentityDkimEnabledResult struct{}

// An empty element. Receiving this element indicates that the request completed successfully.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityFeedbackForwardingEnabledResult.html]
type SetIdentityFeedbackForwardingEnabledResult struct{}

// An empty element. Receiving this element indicates that the request completed successfully.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityNotificationTopicResult.html]
type SetIdentityNotificationTopicResult struct{}

// Represents the DNS records that must be published in the domain name's DNS to complete DKIM setup.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainDkimResult.html]
type VerifyDomainDkimResult struct {
	DkimTokens []string `xml:"DkimTokens"`
}

// Represents a token used for domain ownership verification.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainIdentityResult.html]
type VerifyDomainIdentityResult struct {
	VerificationToken string `xml:"VerificationToken"`
}

// An empty element. Receiving this element indicates that the request completed successfully.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailIdentityResult.html]
type VerifyEmailIdentityResult struct{}
