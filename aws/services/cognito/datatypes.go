package cognito

import ()

/******************************************************************************
 * Date Types
 */

// An object representing a Cognito identity pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CreateIdentityPoolResult.html]
type CreateIdentityPoolResult struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	IdentityPoolId                 string            `json:"IdentityPoolId"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// An object representing a Cognito identity pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPoolResult.html]
type DescribeIdentityPoolResult struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	IdentityPoolId                 string            `json:"IdentityPoolId"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}

// Returned in the response to a GetId request.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetIdResult.html]
type GetIdResult struct {
	IdentityId string `json:"IdentityId,omitempty"`
}

// Returned in response to a successful GetOpenIdToken request.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetOpenIdTokenResult.html]
type GetOpenIdTokenResult struct {
	IdentityId string `json:"IdentityId,omitempty"`
	Token      string `json:"Token,omitempty"`
}

// A description of the identity.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_IdentityDescription.html]
type IdentityDescription struct {
	IdentityId string   `json:"IdentityId,omitempty"`
	Logins     []string `json:"Logins,omitempty"`
}

// A description of the identity pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_IdentityPoolShortDescription.html]
type IdentityPoolShortDescription struct {
	IdentityPoolId   string `json:"IdentityPoolId,omitempty"`
	IdentityPoolName string `json:"IdentityPoolName,omitempty"`
}

// The response to a ListIdentities request.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentitiesResult.html]
type ListIdentitiesResult struct {
	Identities     []IdentityDescription `json:"Identities,omitempty"`
	IdentityPoolId string                `json:"IdentityPoolId,omitempty"`
	NextToken      string                `json:"NextToken,omitempty"`
}

// The result of a successful ListIdentityPools action.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentityPoolsResult.html]
type ListIdentityPoolsResult struct {
	IdentityPools []IdentityPoolShortDescription `json:"IdentityPools,omitempty"`
	NextToken     string                         `json:"NextToken,omitempty"`
}

// An object representing a Cognito identity pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_UpdateIdentityPoolResult.html]
type UpdateIdentityPoolResult struct {
	AllowUnauthenticatedIdentities bool              `json:"AllowUnauthenticatedIdentities"`
	IdentityPoolId                 string            `json:"IdentityPoolId"`
	IdentityPoolName               string            `json:"IdentityPoolName"`
	SupportedLoginProviders        map[string]string `json:"SupportedLoginProviders,omitempty"`
}
