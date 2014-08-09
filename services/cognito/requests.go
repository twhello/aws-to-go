package cognito

import ()

/*****************************************************************************/

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The limit on identity
// pools is 60 per account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CreateIdentityPool.html]
type CreateIdentityPoolRequest struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// Creates a new CreateIdentityPoolRequest.
func NewCreateIdentityPoolRequest(identityPoolName string, allowUnauthenticatedIdentities bool) *CreateIdentityPoolRequest {
	return &CreateIdentityPoolRequest{AllowUnauthenticatedIdentities: allowUnauthenticatedIdentities, IdentityPoolName: identityPoolName}
}

// Add a Supported Login Provider.
func (r *CreateIdentityPoolRequest) AddSupportedLoginProvider(name, provider string) {
	if r.SupportedLoginProviders == nil {
		r.SupportedLoginProviders = make(map[string]string)
	}
	r.SupportedLoginProviders[name] = provider
}

/*****************************************************************************/

// Deletes a user pool. Once a pool is deleted, users will not be able to
// authenticate with the pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DeleteIdentityPool.html]
type DeleteIdentityPoolRequest struct {
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DeleteIdentityPoolRequest.
func NewDeleteIdentityPoolRequest(identityPoolId string) *DeleteIdentityPoolRequest {
	return &DeleteIdentityPoolRequest{identityPoolId}
}

/*****************************************************************************/

// Gets details about a particular identity pool, including the pool name,
// ID description, creation date, and current number of users.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html]
type DescribeIdentityPoolRequest struct {
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DescribeIdentityPoolRequest.
func NewDescribeIdentityPoolRequest(identityPoolId string) *DescribeIdentityPoolRequest {
	return &DescribeIdentityPoolRequest{identityPoolId}
}

/*****************************************************************************/

// Generates (or retrieves) a Cognito ID. Supplying multiple logins will create
// an implicit linked account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html]
type GetIdRequest struct {
	AccountId      string            `json:"AccountId"`
	IdentityPoolId string            `json:"IdentityPoolId"`
	Logins         map[string]string `json:"Logins,omitempty"`
}

// Creates a new GetIdRequest.
func NewGetIdRequest(accountId, identityPoolId string) *GetIdRequest {
	return &GetIdRequest{AccountId: accountId, IdentityPoolId: identityPoolId}
}

// Add a Login.
func (r *GetIdRequest) AddLogin(name, login string) {
	if r.Logins == nil {
		r.Logins = make(map[string]string)
	}
	r.Logins[name] = login
}

/*****************************************************************************/

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned from GetId. You can optionally add additional logins for the identity.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetOpenIdToken.html]
type GetOpenIdTokenRequest struct {
	IdentityId string            `json:"IdentityId"`
	Logins     map[string]string `json:"Logins,omitempty"`
}

// Creates a new GetOpenIdTokenRequest.
func NewGetOpenIdTokenRequest(identityId string) *GetOpenIdTokenRequest {
	return &GetOpenIdTokenRequest{IdentityId: identityId}
}

// Add a Login.
func (r *GetOpenIdTokenRequest) AddLogin(name, login string) {
	if r.Logins == nil {
		r.Logins = make(map[string]string)
	}
	r.Logins[name] = login
}

/*****************************************************************************/

// Lists the identities in a pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentities.html]
type ListIdentitiesRequest struct {
	IdentityPoolId string `json:"IdentityPoolId"`
	MaxResults     int    `json:"MaxResults,omitempty"`
	NextToken      string `json:"NextToken,omitempty"`
}

// Creates a new ListIdentitiesRequest.
func NewListIdentitiesRequest(identityPoolId string) *ListIdentitiesRequest {
	return &ListIdentitiesRequest{IdentityPoolId: identityPoolId}
}

/*****************************************************************************/

// Lists all of the Cognito identity pools registered for your account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentityPools.html]
type ListIdentityPoolsRequest struct {
	MaxResults int    `json:"MaxResults,omitempty"`
	NextToken  string `json:"NextToken,omitempty"`
}

// Creates a new ListIdentityPoolsRequest.
func NewListIdentityPoolsRequest() *ListIdentityPoolsRequest {
	return &ListIdentityPoolsRequest{}
}

/*****************************************************************************/

// Unlinks a federated identity from an existing account. Unlinked logins will
// be considered new identities next time they are seen.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_UnlinkIdentity.html]
type UnlinkIdentityRequest struct {
	IdentityId     string            `json:"IdentityId"`
	Logins         map[string]string `json:"Logins"`
	LoginsToRemove []string          `json:"LoginsToRemove"`
}

// Creates a new UnlinkIdentityRequest.
func NewUnlinkIdentityRequest(identityId string) *UnlinkIdentityRequest {
	return &UnlinkIdentityRequest{IdentityId: identityId}
}

// Add a Login.
func (r *UnlinkIdentityRequest) AddLogin(name, login string) {
	if r.Logins == nil {
		r.Logins = make(map[string]string)
	}
	r.Logins[name] = login
}

// Add Logins to Remove.
func (r *UnlinkIdentityRequest) AddLoginsToRemove(logins ...string) {
	r.LoginsToRemove = append(r.LoginsToRemove, logins...)
}

/*****************************************************************************/

// Updates a user pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_UpdateIdentityPool.html]
type UpdateIdentityPoolRequest struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	IdentityPoolId                 string            `json:"IdentityPoolId"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// Creates a new UpdateIdentityPoolRequest.
func NewUpdateIdentityPoolRequest(identityPoolId, identityPoolName string, allowUnauthenticatedIdentities bool) *UpdateIdentityPoolRequest {
	return &UpdateIdentityPoolRequest{
		IdentityPoolId:                 identityPoolId,
		IdentityPoolName:               identityPoolName,
		AllowUnauthenticatedIdentities: allowUnauthenticatedIdentities,
	}
}

// Add a Supported Login Provider.
func (r *UnlinkIdentityRequest) AddSupportedLoginProvider(name, id string) {
	if r.Logins == nil {
		r.Logins = make(map[string]string)
	}
	r.Logins[name] = id
}
