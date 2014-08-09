package s3

import (
	"time"
)

// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html]
type ListBucketsResult struct {
	Owner   Owner    `xml:"Owner"`
	Buckets []Bucket `xml:"Buckets>Bucket"`
}

// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html]
type ListObjectsResult struct {
	Contents       []Contents `xml:"Contents"`
	CommonPrefixes []string   `xml:"CommonPrefixes>Prefix"`
	Delimiter      string     `xml:"Delimiter"`
	EncodingType   string     `xml:"Ecoding-Type"`
	IsTruncated    bool       `xml:"IsTruncated"`
	Marker         string     `xml:"Marker"`
	MaxKeys        string     `xml:"MaxKeys"`
	Name           string     `xml:"Name"`
	NextMarker     string     `xml:"NextMarker"`
	Prefix         string     `xml:"Prefix"`
}

// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html]
type GetObjectHeaderResponse struct {
	DeleteMarker                         string            `name:"X-Amz-Delete-Marker"`
	Expiration                           *time.Time        `name:"X-Amz-Expiration" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	Metadata                             map[string]string `name:"X-Amz-Meta-*"`
	MissingMetadata                      int               `name:"X-Amz-Missing-Meta"`
	ServerSideEncryption                 string            `name:"X-Amz-Server-Side-Encryption"`
	ServerSideEncryptionCustomerAlgorith string            `name:"X-Amz-Server-Side-Encryption-Customer-Algorithm"`
	ServerSideEncryptionCustomerKeyMD5   string            `name:"X-Amz-Server-Side-Encryption-Customer-key-Md5"`
	Restore                              string            `name:"X-Amz-Restore"`
	VersionId                            string            `name:"X-Amz-Version-Id"`
	WebsiteRedirectLocation              string            `name:"X-Amz-Website-Redirect-Location"`
}

// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html]
type DeleteObjectHeaderResponse struct {
	DeleteMarker string `name:"X-Amz-Delete-Marker"`
	VersionId    string `name:"X-Amz-Version-Id"`
}

// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html]
type PutObjectHeaderResponse struct {
	Expiration                           time.Time `name:"X-Amz-Expiration" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	ServerSideEncryption                 string    `name:"X-Amz-Server-Side-Encryption"`
	ServerSideEncryptionCustomerAlgorith string    `name:"X-Amz-Server-Side-Encryption-Customer-Algorithm"`
	ServerSideEncryptionCustomerKeyMD5   string    `name:"X-Amz-Server-Side-Encryption-Customer-key-Md5"`
	VersionId                            string    `name:"X-Amz-Version-Id"`
}

// [http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html]
type DeleteResult struct {
	Deleted []struct {
		Key                   string `xml:"Key"`
		VersionId             string `xml:"VersionId"`
		DeleteMarker          bool   `xml:"DeleteMarker"`
		DeleteMarkerVersionId string `xml:"DeleteMarkerVersionId"`
	} `xml:"Deleted"`
	Error []struct {
		Key       string `xml:"Key"`
		VersionId string `xml:"VersionId"`
		Code      string `xml:"Code"`
		Message   string `xml:"Message"`
	} `xml:"Error"`
}
