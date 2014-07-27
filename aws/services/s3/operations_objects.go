package s3

import (
	"io"
)

/*****************************************************************************/

// The DELETE operation removes the null version (if there is one) of an object
// and inserts a delete marker, which becomes the current version of the object.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html]
type DeleteObjectRequest struct {
	BucketName string
	ObjectName string
}

// Creates a new DeleteObjectRequest.
func NewDeleteObjectRequest(bucketName, objectName string) *DeleteObjectRequest {
	return &DeleteObjectRequest{bucketName, objectName}
}

/*****************************************************************************/

// The Multi-Object Delete operation enables you to delete multiple objects from a
// bucket using a single HTTP request.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html]
type DeleteMultipleObjectsRequest struct {
	BucketName string
	Headers    *DeleteMultipleObjectsHeaders // Can be nil
	Delete     *Delete
}

// Creates a new DeleteMultipleObjectsRequest.
func NewDeleteMultipleObjectsRequest(bucketName string, keys ...string) *DeleteMultipleObjectsRequest {

	objects := make([]Object, len(keys))

	for i, key := range keys {
		objects[i] = Object{Key: key}
	}

	return &DeleteMultipleObjectsRequest{BucketName: bucketName, Delete: &Delete{Object: objects}}
}

/*****************************************************************************/

// This implementation of the GET operation retrieves objects from Amazon S3.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html]
type GetObjectRequest struct {
	BucketName              string
	ObjectName              string
	Constraints             *Constraints             // Can be nil
	ResponseHeaderOverrides *ResponseHeaderOverrides // Can be nil
}

// Creates a new GetObjectRequest.
func NewGetObjectRequest(bucketName, objectName string) *GetObjectRequest {
	return &GetObjectRequest{bucketName, objectName, nil, nil}
}

/*****************************************************************************/

// This implementation of the PUT operation adds an object to a bucket.
// You must have WRITE permissions on a bucket to add an object to it.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html]
type PutObjectRequest struct {
	BucketName     string
	ObjectName     string
	Content        io.Reader
	ObjectMetadata *ObjectMetadata
}

// Creates a new PutObjectRequest.
func NewPutObjectRequest(bucketName, objectName string, content io.Reader, metadata *ObjectMetadata) *PutObjectRequest {
	return &PutObjectRequest{bucketName, objectName, content, metadata}
}
