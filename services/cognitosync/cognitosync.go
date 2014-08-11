//
// Amazon Cognito Sync provides an AWS service and client library that enable cross-device
// syncing of application-related user data. High-level client libraries are available for
// both iOS and Android. You can use these libraries to persist data locally so that it's
// available even if the device is offline. Developer credentials don't need to be stored
// on the mobile device to access the service. You can use Amazon Cognito to obtain a
// normalized user ID and credentials. User data is persisted in a dataset that can store
// up to 1 MB of key-value pairs, and you can have up to 20 datasets per user identity.
// 
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/Welcome.html]
//
package cognitosync

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"net/http"
)

const ServiceName = "cognito"

// CognitoSyncService describes the API interface. Instantiate with cognitosync.NewService().
type CognitoSyncService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *CognitoSyncService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *CognitoSyncService) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *CognitoSyncService) Endpoint() string {
	return s.endpoint
}

// Low-level request to Cognito Sync Service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *CognitoSyncService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *CognitoSyncService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

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
 * Cognito Sync Service Methods.
 */

// Deletes the specific dataset. The dataset will be deleted permanently, and
// the action can't be undone. Datasets that this dataset was merged with will
// no longer report the merge. Any consequent operation on this dataset will
// result in a ResourceNotFoundException.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DeleteDataset.html]
func (s *CognitoSyncService) DeleteDataset(req *DeleteDatasetRequest) (result *DeleteDatasetResult, err error) {

	result = new(DeleteDatasetResult)
	err = s.wrapperSignAndDo("DeleteDataset", req, result)
	return
}

// Gets metadata about a dataset by identity and dataset name.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeDataset.html]
func (s *CognitoSyncService) DescribeDataset(req *DescribeDatasetRequest) (result *DescribeDatasetResult, err error) {

	result = new(DescribeDatasetResult)
	err = s.wrapperSignAndDo("DescribeDataset", req, result)
	return
}

// Gets usage details (for example, data storage) about a particular identity pool.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityPoolUsage.html]
func (s *CognitoSyncService) DescribeIdentityPoolUsage(req *DescribeIdentityPoolUsageRequest) (result *DescribeIdentityPoolUsageResult, err error) {

	result = new(DescribeIdentityPoolUsageResult)
	err = s.wrapperSignAndDo("DescribeIdentityPoolUsage", req, result)
	return
}

// Gets usage information for an identity, including number of datasets and data usage.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityUsage.html]
func (s *CognitoSyncService) DescribeIdentityUsage(req *DescribeIdentityUsageRequest) (result *DescribeIdentityUsageResult, err error) {

	result = new(DescribeIdentityUsageResult)
	err = s.wrapperSignAndDo("DescribeIdentityUsage", req, result)
	return
}

// Lists datasets for an identity.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListDatasets.html]
func (s *CognitoSyncService) ListDatasets(req *ListDatasetsRequest) (result *ListDatasetsResult, err error) {

	result = new(ListDatasetsResult)
	err = s.wrapperSignAndDo("ListDatasets", req, result)
	return
}

// Gets a list of identity pools registered with Cognito.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListIdentityPoolUsage.html]
func (s *CognitoSyncService) ListIdentityPoolUsage(req *ListIdentityPoolUsageRequest) (result *ListIdentityPoolUsageResult, err error) {

	result = new(ListIdentityPoolUsageResult)
	err = s.wrapperSignAndDo("ListIdentityPoolUsage", req, result)
	return
}

// Gets paginated records, optionally changed after a particular sync count for a dataset and identity.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListRecords.html]
func (s *CognitoSyncService) ListRecords(req *ListRecordsRequest) (result *ListRecordsResult, err error) {

	result = new(ListRecordsResult)
	err = s.wrapperSignAndDo("ListRecords", req, result)
	return
}

// Posts updates to records and add and delete records for a dataset and user.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_UpdateRecords.html]
func (s *CognitoSyncService) UpdateRecords(req *UpdateRecordsRequest) (result *UpdateRecordsResult, err error) {

	result = new(UpdateRecordsResult)
	err = s.wrapperSignAndDo("UpdateRecords", req, result)
	return
}

// Creates a new Cognito Sync Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *CognitoSyncService {
	return &CognitoSyncService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
