//
// Amazon Cognito is a web service that facilitates the delivery of scoped, temporary credentials
// to mobile devices or other untrusted environments. Amazon Cognito uniquely identifies a device
// or user and supplies the user with a consistent identity throughout the lifetime of an application.
//
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/Welcome.html]
//
package cognito

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"net/http"
)

const ServiceName = "cognito"

// CognitoService describes the API interface. Instantiate with cognito.NewService().
type CognitoService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *CognitoService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *CognitoService) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *CognitoService) Endpoint() string {
	return s.endpoint
}

// Low-level request to Cognito Service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *CognitoService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *CognitoService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	if err == nil {
		h := req.Header()
		h.Set("Content-Type", "application/x-amz-json-1.1")
		h.Set("Action", action)
		h.Set("Version", "2014-06-30")
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * Cognito Service Methods.
 */

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The limit on identity
// pools is 60 per account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CreateIdentityPool.html]
func (s *CognitoService) CreateIdentityPool(req *CreateIdentityPoolRequest) (result *CreateIdentityPoolResult, err error) {

	result = new(CreateIdentityPoolResult)
	err = s.wrapperSignAndDo("CreateIdentityPool", req, result)
	return
}

// Deletes a user pool. Once a pool is deleted, users will not be able to
// authenticate with the pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DeleteIdentityPool.html]
func (s *CognitoService) DeleteIdentityPool(req *DeleteIdentityPoolRequest) (err error) {

	err = s.wrapperSignAndDo("DeleteIdentityPool", req, nil)
	return
}

// Gets details about a particular identity pool, including the pool name,
// ID description, creation date, and current number of users.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_DescribeIdentityPool.html]
func (s *CognitoService) DescribeIdentityPool(req *DescribeIdentityPoolRequest) (result *DescribeIdentityPoolResult, err error) {

	result = new(DescribeIdentityPoolResult)
	err = s.wrapperSignAndDo("DescribeIdentityPool", req, result)
	return
}

// Generates (or retrieves) a Cognito ID. Supplying multiple logins will create
// an implicit linked account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html]
func (s *CognitoService) GetId(req *GetIdRequest) (result *GetIdResult, err error) {

	result = new(GetIdResult)
	err = s.wrapperSignAndDo("GetId", req, result)
	return
}

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned from GetId. You can optionally add additional logins for the identity.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetOpenIdToken.html]
func (s *CognitoService) GetOpenIdToken(req *GetOpenIdTokenRequest) (result *GetOpenIdTokenResult, err error) {

	result = new(GetOpenIdTokenResult)
	err = s.wrapperSignAndDo("GetOpenIdToken", req, result)
	return
}

// Lists the identities in a pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentities.html]
func (s *CognitoService) ListIdentities(req *ListIdentitiesRequest) (result *ListIdentitiesResult, err error) {

	result = new(ListIdentitiesResult)
	err = s.wrapperSignAndDo("ListIdentities", req, result)
	return
}

// Lists all of the Cognito identity pools registered for your account.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_ListIdentityPools.html]
func (s *CognitoService) ListIdentityPools(req *ListIdentityPoolsRequest) (result *ListIdentityPoolsResult, err error) {

	result = new(ListIdentityPoolsResult)
	err = s.wrapperSignAndDo("ListIdentityPools", req, result)
	return
}

// Unlinks a federated identity from an existing account. Unlinked logins will
// be considered new identities next time they are seen.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_UnlinkIdentity.html]
func (s *CognitoService) UnlinkIdentity(req *UnlinkIdentityRequest) (err error) {

	err = s.wrapperSignAndDo("UnlinkIdentity", req, nil)
	return
}

// Updates a user pool.
// [http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_UpdateIdentityPool.html]
func (s *CognitoService) UpdateIdentityPool(req *UpdateIdentityPoolRequest) (result *UpdateIdentityPoolResult, err error) {

	result = new(UpdateIdentityPoolResult)
	err = s.wrapperSignAndDo("UpdateIdentityPool", req, result)
	return
}

// Creates a new Cognito Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *CognitoService {
	return &CognitoService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
