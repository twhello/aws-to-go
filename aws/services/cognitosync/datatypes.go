package cognitosync

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/******************************************************************************
 * Constants
 */

//	REMOVE = "remove"
//	REPLACE = "replace"
type Operation string

const (
	REMOVE  Operation = "remove"
	REPLACE Operation = "replace"
)

/******************************************************************************
 * Data Types
 */

// A collection of data for an identity pool. An identity pool can have multiple datasets. A dataset is per identity and can be general or associated with a particular entity in an application (like a saved game). Datasets are automatically created if they don't exist. Data is synced by dataset, and a dataset can hold up to 1MB of key-value pairs.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_Dataset.html]
type Dataset struct {
	CreationDate     *datetime.JsonDate `json:"CreationDate,omitempty"`
	DataStorage      int64              `json:"DataStorage,omitempty"`
	DatasetName      string             `json:"DatasetName,omitempty"`
	IdentityId       string             `json:"IdentityId,omitempty"`
	LastModifiedBy   string             `json:"LastModifiedBy,omitempty"`
	LastModifiedDate *datetime.JsonDate `json:"LastModifiedDate,omitempty"`
	NumRecords       int64              `json:"NumRecords,omitempty"`
}

// Response to a successful DeleteDataset request.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DeleteDatasetResult.html]
type DeleteDatasetResult struct {
	Dataset *Dataset `json:"Dataset,omitempty"`
}

// Response to a successful DescribeDataset request.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeDatasetResult.html]
type DescribeDatasetResult struct {
	Dataset *Dataset `json:"Dataset,omitempty"`
}

// Response to a successful DescribeIdentityPoolUsage request.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityPoolUsageResult.html]
type DescribeIdentityPoolUsageResult struct {
	IdentityPoolUsage *IdentityPoolUsage `json:" IdentityPoolUsage,omitempty"`
}

// The response to a successful DescribeIdentityUsage request.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_DescribeIdentityUsageResult.html]
type DescribeIdentityUsageResult struct {
	IdentityUsage *IdentityUsage `json:"IdentityUsage,omitempty"`
}

// Usage information for the identity pool.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_IdentityPoolUsage.html]
type IdentityPoolUsage struct {
	DataStorage       int64              `json:"DataStorage,omitempty"`
	IdentityPoolId    string             `json:"IdentityPoolId,omitempty"`
	LastModifiedDate  *datetime.JsonDate `json:"LastModifiedDate,omitempty"`
	SyncSessionsCount int64              `json:SyncSessionsCount,omitempty"`
}

// Usage information for the identity.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_IdentityUsage.html]
type IdentityUsage struct {
	DataStorage      int64              `json:"DataStorage,omitempty"`
	DatasetCount     int                `json:"DatasetCount,omitempty"`
	IdentityId       string             `json:"IdentityId,omitempty"`
	IdentityPoolId   string             `json:"IdentityPoolId,omitempty"`
	LastModifiedDate *datetime.JsonDate `json:"LastModifiedDate,omitempty"`
}

// Returned for a successful ListDatasets request.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListDatasetsResult.html]
type ListDatasetsResult struct {
	Count     int       `json:"Count,omitempty"`
	Datasets  []Dataset `json:"Datasets,omitempty"`
	NextToken string    `json:"NextToken,omitempty"`
}

//
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListIdentityPoolUsageResult.html]
type ListIdentityPoolUsageResult struct {
	Count              int                 `json:"Count,omitempty"`
	IdentityPoolUsages []IdentityPoolUsage `json:",omitempty"`
	MaxResults         int                 `json:",omitempty"`
	NextToken          string              `json:"NextToken,omitempty"`
}

// Returned for a successful ListRecordsRequest.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_ListRecordsResult.html]
type ListRecordsResult struct {
	Count                                 int      `json:"Count,omitempty"`
	DatasetDeletedAfterRequestedSyncCount bool     `json:"DatasetDeletedAfterRequestedSyncCount,omitempty"`
	DatasetExists                         bool     `json:"DatasetExists,omitempty"`
	DatasetSyncCount                      int64    `json:"DatasetSyncCount,omitempty"`
	LastModifiedBy                        string   `json:"LastModifiedBy,omitempty"`
	MergedDatasetNames                    []string `json:"MergedDatasetNames,omitempty"`
	NextToken                             string   `json:"NextToken,omitempty"`
	Records                               []Record `json:"Records,omitempty"`
	SyncSessionToken                      string   `json:"SyncSessionToken,omitempty"`
}

// The basic data structure of a dataset.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_Record.html]
type Record struct {
	DeviceLastModifiedDate *datetime.JsonDate `json:"DeviceLastModifiedDate,omitempty"`
	Key                    string             `json:"Key,omitempty"`
	LastModifiedBy         string             `json:"LastModifiedBy,omitempty"`
	LastModifiedDate       *datetime.JsonDate `json:"LastModifiedDate,omitempty"`
	SyncCount              int64              `json:"SyncCount,omitempty"`
	Value                  string             `json:"Value,omitempty"`
}

// An update operation for a record.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_RecordPatch.html]
type RecordPatch struct {
	DeviceLastModifiedDate *datetime.JsonDate `json:"DeviceLastModifiedDate,omitempty"`
	Key                    string             `json:"Key"`
	Op                     Operation          `json:""`
	SyncCount              int64              `json:"SyncCount"`
	Value                  string             `json:"Value,omitempty"`
}

// Returned for a successful UpdateRecordsRequest.
// [http://docs.aws.amazon.com/cognitosync/latest/APIReference/API_UpdateRecordsResult.html]
type UpdateRecordsResult struct {
	Records []Record `json:"Records,omitempty"`
}
