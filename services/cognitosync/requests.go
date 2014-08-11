package cognitosync

import ()

/*****************************************************************************/

// Deletes the specific dataset. The dataset will be deleted permanently, and
// the action can't be undone. Datasets that this dataset was merged with will
// no longer report the merge. Any consequent operation on this dataset will
// result in a ResourceNotFoundException.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DeleteDataset.html]
type DeleteDatasetRequest struct {
	DatasetName    string `json:"DatasetName"`
	IdentityId     string `json:"IdentityId"`
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DeleteDatasetRequest.
func NewDeleteDatasetRequest(identityId, identityPoolId, datasetName string) *DeleteDatasetRequest {
	return &DeleteDatasetRequest{datasetName, identityId, identityPoolId}
}

/*****************************************************************************/

// Gets metadata about a dataset by identity and dataset name.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeDataset.html]
type DescribeDatasetRequest struct {
	DatasetName    string `json:"DatasetName"`
	IdentityId     string `json:"IdentityId"`
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DescribeDatasetRequest.
func NewDescribeDatasetRequest(identityId, identityPoolId, datasetName string) *DescribeDatasetRequest {
	return &DescribeDatasetRequest{datasetName, identityId, identityPoolId}
}

/*****************************************************************************/

// Gets usage details (for example, data storage) about a particular identity pool.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityPoolUsage.html]
type DescribeIdentityPoolUsageRequest struct {
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DescribeIdentityPoolUsageRequest.
func NewDescribeIdentityPoolUsageRequest(identityPoolId string) *DescribeIdentityPoolUsageRequest {
	return &DescribeIdentityPoolUsageRequest{identityPoolId}
}

/*****************************************************************************/

// Gets usage information for an identity, including number of datasets and data usage.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityUsage.html]
type DescribeIdentityUsageRequest struct {
	IdentityId     string `json:"IdentityId"`
	IdentityPoolId string `json:"IdentityPoolId"`
}

// Creates a new DescribeIdentityUsageRequest.
func NewDescribeIdentityUsageRequest(identityId, identityPoolId string) *DescribeIdentityUsageRequest {
	return &DescribeIdentityUsageRequest{identityId, identityPoolId}
}

/*****************************************************************************/

// Lists datasets for an identity.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListDatasets.html]
type ListDatasetsRequest struct {
	IdentityId     string `json:"IdentityId"`
	IdentityPoolId string `json:"IdentityPoolId,omitempty"`
	MaxResults     string `json:"MaxResults,omitempty"`
	NextToken      string `json:"NextToken,omitempty"`
}

// Creates a new ListDatasetsRequest.
func NewListDatasetsRequest(identityId string) *ListDatasetsRequest {
	return &ListDatasetsRequest{IdentityId: identityId}
}

/*****************************************************************************/

// Gets a list of identity pools registered with Cognito.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListIdentityPoolUsage.html]
type ListIdentityPoolUsageRequest struct {
	MaxResults string `json:"MaxResults,omitempty"`
	NextToken  string `json:"NextToken,omitempty"`
}

// Creates a new ListIdentityPoolUsageRequest.
func NewListIdentityPoolUsageRequest() *ListIdentityPoolUsageRequest {
	return &ListIdentityPoolUsageRequest{}
}

/*****************************************************************************/

// Gets paginated records, optionally changed after a particular sync count for a dataset and identity.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListRecords.html]
type ListRecordsRequest struct {
	DatasetName      string `json:"DatasetName"`
	IdentityId       string `json:"IdentityId"`
	IdentityPoolId   string `json:"IdentityPoolId"`
	LastSyncCount    string `json:"LastSyncCount,omitempty"`
	MaxResults       string `json:"MaxResults,omitempty"`
	NextToken        string `json:"NextToken,omitempty"`
	SyncSessionToken string `json:"SyncSessionToken,omitempty"`
}

// Creates a new ListRecordsRequest.
func NewListRecordsRequest(identityId, identityPoolId, datasetName string) *ListRecordsRequest {
	return &ListRecordsRequest{DatasetName: datasetName, IdentityId: identityId, IdentityPoolId: identityPoolId}
}

/*****************************************************************************/

// Posts updates to records and add and delete records for a dataset and user.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_UpdateRecords.html]
type UpdateRecordsRequest struct {
	ClientContext    string        `json:"ClientContext,omitempty"`
	DatasetName      string        `json:"DatasetName"`
	IdentityId       string        `json:"IdentityId"`
	IdentityPoolId   string        `json:"IdentityPoolId"`
	RecordPatches    []RecordPatch `json:"RecordPatches,omitempty"`
	SyncSessionToken string        `json:"SyncSessionToken"`
}

// Creates a new UpdateRecordsRequest.
func NewUpdateRecordsRequest(identityId, identityPoolId, datasetName, syncSessionToken string) *UpdateRecordsRequest {
	return &UpdateRecordsRequest{
		DatasetName: datasetName,
		IdentityId: identityId,
		IdentityPoolId: identityPoolId,
		SyncSessionToken: syncSessionToken,
	}
}

// Add RecordPatches.
func (r *UpdateRecordsRequest) AddRecordPatches(recordPatches ...RecordPatch) {
	r.RecordPatches = append(r.RecordPatches, recordPatches...)
}
