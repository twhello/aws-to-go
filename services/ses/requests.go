package ses

import ()

/*****************************************************************************/

// Deletes the specified identity (email address or domain) from the list of verified identities.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteIdentity.html]
type DeleteIdentityRequest struct {
	Identity string `name:"Identity"`
}

// Creates a new DeleteIdentityRequest.
func NewDeleteIdentityRequest(identity string) *DeleteIdentityRequest {
	return &DeleteIdentityRequest{identity}
}

/*****************************************************************************/

// Deletes the specified email address from the list of verified addresses.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_DeleteVerifiedEmailAddress.html]
type DeleteVerifiedEmailAddressRequest struct {
	EmailAddress string `name:"EmailAddress"`
}

// Creates a new DeleteVerifiedEmailAddressRequest.
func NewDeleteVerifiedEmailAddressRequest(emailAddress string) *DeleteVerifiedEmailAddressRequest {
	return &DeleteVerifiedEmailAddressRequest{emailAddress}
}

/*****************************************************************************/

// Returns the current status of Easy DKIM signing for an entity. For domain name identities,
// this action also returns the DKIM tokens that are required for Easy DKIM signing, and
// whether Amazon SES has successfully verified that these tokens have been published.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityDkimAttributes.html]
type GetIdentityDkimAttributesRequest struct {
	Identities []string `name:"Identities.member.#"`
}

// Creates a new GetIdentityDkimAttributesRequest.
func NewGetIdentityDkimAttributesRequest(identities ...string) *GetIdentityDkimAttributesRequest {
	return &GetIdentityDkimAttributesRequest{identities}
}

/*****************************************************************************/

// Given a list of verified identities (email addresses and/or domains), returns a
// structure describing identity notification attributes.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityNotificationAttributes.html]
type GetIdentityNotificationAttributesRequest struct {
	Identities []string `name:"Identities.member.#"`
}

// Creates a new GetIdentityNotificationAttributesRequest.
func NewGetIdentityNotificationAttributesRequest(identities ...string) *GetIdentityNotificationAttributesRequest {
	return &GetIdentityNotificationAttributesRequest{identities}
}

/*****************************************************************************/

// Given a list of identities (email addresses and/or domains), returns the verification
// status and (for domain identities) the verification token for each identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_GetIdentityVerificationAttributes.html]
type GetIdentityVerificationAttributesRequest struct {
	Identities []string `name:"Identities.member.#"`
}

// Creates a new GetIdentityVerificationAttributesRequest.
func NewGetIdentityVerificationAttributesRequest(identities ...string) *GetIdentityVerificationAttributesRequest {
	return &GetIdentityVerificationAttributesRequest{identities}
}

/*****************************************************************************/

// Returns a list containing all of the identities (email addresses and domains)
// for a specific AWS Account, regardless of verification status.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_ListIdentities.html]
type ListIdentitiesRequest struct {
	IdentityType IdentityType `name:"IdentityType,omitempty"`
	MaxItems     int          `name:"MaxItems,omitempty"`
	NextToken    string       `name:"NextToken,omitempty"`
}

// Creates a new ListIdentitiesRequest.
func NewListIdentitiesRequest(identityType IdentityType) *ListIdentitiesRequest {
	return &ListIdentitiesRequest{IdentityType: identityType, MaxItems: 100}
}

/*****************************************************************************/

// Composes an email message based on input data, and then immediately queues the message for sending.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendEmail.html]
type SendEmailRequest struct {
	Destination      Destination `name:"Destination."`
	Message          Message     `name:"Message."`
	ReplyToAddresses []string    `name:"ReplyToAddresses.member.#,omitempty"`
	ReturnPath       string      `name:"ReturnPath,omitempty"`
	Source           string      `name:"Source"`
}

// Creates a new SendEmailRequest.
func NewSendEmailRequest(source, destination string, message Message) *SendEmailRequest {
	return &SendEmailRequest{
		Destination: Destination{ToAddresses: []string{source}},
		Message:     message,
		Source:      source,
	}
}

/*****************************************************************************/

// Sends an email message, with header and content specified by the client. The SendRawEmail
// action is useful for sending multipart MIME emails. The raw text of the message must
// comply with Internet email standards; otherwise, the message cannot be sent.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SendRawEmail.html]
type SendRawEmailRequest struct {
	Destinations []string   `name:"Destinations.member.#,omitempty"`
	RawMessage   RawMessage `name:"RawMessage."`
	Source       string     `name:"Source,omitempty"`
}

// Creates a new SendRawEmailRequest.
func NewSendRawEmailRequest(rawMessage RawMessage) *SendRawEmailRequest {
	return &SendRawEmailRequest{RawMessage: rawMessage}
}

/*****************************************************************************/

// Enables or disables Easy DKIM signing of email sent from an identity.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityDkimEnabled.html]
type SetIdentityDkimEnabledRequest struct {
	DkimEnabled bool   `name:"DkimEnabled"`
	Identity    string `name:"Identity"`
}

// Creates a new SetIdentityDkimEnabledRequest.
func NewSetIdentityDkimEnabledRequest(dkimEnabled bool, identity string) *SetIdentityDkimEnabledRequest {
	return &SetIdentityDkimEnabledRequest{dkimEnabled, identity}
}

/*****************************************************************************/

// Given an identity (email address or domain), enables or disables whether Amazon SES
// forwards bounce and complaint notifications as email. Feedback forwarding can only
// be disabled when Amazon Simple Notification Service (Amazon SNS) topics are specified
// for both bounces and complaints.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityFeedbackForwardingEnabled.html]
type SetIdentityFeedbackForwardingEnabledRequest struct {
	ForwardingEnabled bool   `name:"ForwardingEnabled"`
	Identity          string `name:"Identity"`
}

// Creates a new SetIdentityFeedbackForwardingEnabledRequest.
func NewSetIdentityFeedbackForwardingEnabledRequest(forwardingEnabled bool, identity string) *SetIdentityFeedbackForwardingEnabledRequest {
	return &SetIdentityFeedbackForwardingEnabledRequest{forwardingEnabled, identity}
}

/*****************************************************************************/

// Given an identity (email address or domain), sets the Amazon Simple Notification Service (Amazon SNS)
// topic to which Amazon SES will publish bounce, complaint, and/or delivery notifications
// for emails sent with that identity as the Source.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityNotificationTopic.html]
type SetIdentityNotificationTopicRequest struct {
	Identity         string           `name:"Identity"`
	NotificationType NotificationType `name:"NotificationType"`
	SnsTopic         string           `name:"SnsTopic,omitempty"`
}

// Creates a new SetIdentityNotificationTopicRequest.
func NewSetIdentityNotificationTopicRequest(identity string, notificationType NotificationType) *SetIdentityNotificationTopicRequest {
	return &SetIdentityNotificationTopicRequest{Identity: identity, NotificationType: notificationType}
}

/*****************************************************************************/

// Returns a set of DKIM tokens for a domain.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainDkim.html]
type VerifyDomainDkimRequest struct {
	Domain string `name:"Domain"`
}

// Creates a new VerifyDomainDkimRequest.
func NewVerifyDomainDkimRequest(domain string) *VerifyDomainDkimRequest {
	return &VerifyDomainDkimRequest{domain}
}

/*****************************************************************************/

// Verifies a domain.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyDomainIdentity.html]
type VerifyDomainIdentityRequest struct {
	Domain string `name:"Domain"`
}

// Creates a new VerifyDomainIdentityRequest.
func NewVerifyDomainIdentityRequest(domain string) *VerifyDomainIdentityRequest {
	return &VerifyDomainIdentityRequest{domain}
}

/*****************************************************************************/

// Verifies an email address. This action causes a confirmation email message to be sent
// to the specified address.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailAddress.html]
type VerifyEmailAddressRequest struct {
	EmailAddress string `name:"EmailAddress"`
}

// Creates a new VerifyEmailAddressRequest.
func NewVerifyEmailAddressRequest(emailAddress string) *VerifyEmailAddressRequest {
	return &VerifyEmailAddressRequest{emailAddress}
}

/*****************************************************************************/

// Verifies an email address. This action causes a confirmation email message to be
// sent to the specified address.
// [http://docs.aws.amazon.com/ses/latest/APIReference/API_VerifyEmailIdentity.html]
type VerifyEmailIdentityRequest struct {
	EmailAddress string `name:"EmailAddress"`
}

// Creates a new VerifyEmailIdentityRequest.
func NewVerifyEmailIdentityRequest(emailAddress string) *VerifyEmailIdentityRequest {
	return &VerifyEmailIdentityRequest{emailAddress}
}
