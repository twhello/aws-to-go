package s3

import ()

/*****************************************************************************/

// This implementation of the PUT operation creates a new bucket.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUT.html]
type CreateBucketRequest struct {
	BucketName                string
	CannedACL                 CannedACL
	CreateBucketConfiguration *CreateBucketConfiguration
}

// Creates a new CreateBucketRequest.
func NewCreateBucketRequest(bucketName string) *CreateBucketRequest {
	return &CreateBucketRequest{BucketName: bucketName}
}

/*****************************************************************************/

// This implementation of the DELETE operation deletes the bucket named in the URI.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketDELETE.html]
type DeleteBucketRequest struct {
	BucketName string
}

// Creates a new DeleteBucketRequest.
func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
	return &DeleteBucketRequest{bucketName}
}

/*****************************************************************************/

// This implementation of the GET operation returns some or all (up to 1000) of the objects in a bucket.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html]
type ListObjectsRequest struct {
	BucketName  string
	ListObjects *ListObjects
}

// Creates a new ListObjectsRequest.
func NewListObjectsRequest(bucketName string) *ListObjectsRequest {
	return &ListObjectsRequest{BucketName: bucketName}
}
