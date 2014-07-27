package s3

import (
	"encoding/xml"
	"time"
)

/*	Canned ACL Constants.
	PRIVATE = "private"
	PUBLIC_READ = "public-read"
	PUBLIC_READ_WRITE = "public-read-write"
	AUTHENTICATE_READ = "authenticated-read"
	BUCKET_OWNER_READ = "bucket-owner-read"
	BUCKET_OWNER_FULL_CONTROL = "bucket-owner-full-control"
	LOG_DELIVERY_WRITE = "log-delivery-write"
	[http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl] */
type CannedACL string

const (
	PRIVATE                   CannedACL = "private"
	PUBLIC_READ               CannedACL = "public-read"
	PUBLIC_READ_WRITE         CannedACL = "public-read-write"
	AUTHENTICATE_READ         CannedACL = "authenticated-read"
	BUCKET_OWNER_READ         CannedACL = "bucket-owner-read"
	BUCKET_OWNER_FULL_CONTROL CannedACL = "bucket-owner-full-control"
	LOG_DELIVERY_WRITE        CannedACL = "log-delivery-write"
)

type Owner struct {
	ID          string `xml:"ID"`
	DisplayName string `xml:"DisplayName"`
}

type Bucket struct {
	Name         string    `xml:"Name"`
	CreationDate time.Time `xml:"CreationDate"`
}

type Contents struct {
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	ETag         string    `xml:"ETag"`
	Size         int64     `xml:"Size"`
	StorageClass string    `xml:"StorageClass"`
	Owner        Owner     `xml:"Owner"`
}

type Delete struct {
	Object []Object `xml:"Object"`
	Quiet  bool     `xml:"Quiet"`
}

type Object struct {
	Key       string `xml:"Name"`
	VersionId string `xml:"VersionId,omitempty"`
}

// Type: XML
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUT.html]
type CreateBucketConfiguration struct {
	XMLName            xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CreateBucketConfiguration"`
	LocationConstraint string   `xml:"LocationConstraint"`
}

// Type: Header Values
// [http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html]
type DeleteMultipleObjectsHeaders struct {
	MFA string `name:"X-Amz-Mfa"`
}

// Type: Header Values
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html]
type ObjectMetadata struct {
	CacheControl                          string            `name:"Cache-Control"`
	ContentDisposition                    string            `name:"Content-Disposition"`
	ContentEncoding                       string            `name:"Content-Encoding"`
	ContentType                           string            `name:"Content-Type" default:"binary/octet-stream"`
	Expect                                string            `name:"Expect"`
	Expires                               *time.Time        `name:"Expires" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	Metadata                              map[string]string `name:"X-Amz-Meta-*"`
	StorageClass                          string            `name:"X-Amz-Storage-Class" default:"STANDARD"`
	WebsiteRedirectionLocation            string            `name:"X-Amz-Website-Redirect-Location"`
	CannedACL                             string            `name:"X-Amz-Acl"`
	ServerSideEncryption                  string            `name:"X-Amz-Server-Side-Encryption"`
	ServerSideEncryptionCustomerAlgorithm string            `name:"X-Amz-Server-Side-Encryption-Customer-Algorithm"`
	ServerSideEncryptionCustomerKey       string            `name:"X-Amz-server-Side-Encryption-Customer-Key"`
	ServerSideEncryptionCustomerKeyMD5    string            `name:"X-Amz-server-Side-Encryption-Customer-Key-Md5"`
}

//
// []
type ListObjects struct {
	Delimiter    string `name:"delimiter"`
	EncodingType string `name:"encoding-type"`
	Marker       string `name:"marker"`
	MaxKeys      string `name:"max-keys" default:"1000"`
	Prefix       string `name:"prefix"`
}

// There are times when you want to override certain response header values in a GET response.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html]
type ResponseHeaderOverrides struct {
	ResponseContentType        string     `name:"Response-Content-Type"`
	ResponseContentLanguage    string     `name:"Response-Content-Language"`
	ResponseExpires            *time.Time `name:"Response-Expires" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	ResponseCacheControl       string     `name:Response-Cache-Control"`
	ResponseContentDisposition string     `name:"Response-Content-Disposition"`
	ResponseContentEncoding    string     `name:"Response-Content-Encoding"`
}

// Type: Header Values
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html]
type Constraints struct {
	Range                                 string     `name:"Range"`
	IfModifiedSince                       *time.Time `name:"If-Modified-Since" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	IfUnmodifiedSince                     *time.Time `name:"If-Unmodified-Since" format:"Mon, 02 Jan 2006 15:04:05 MST"`
	IfMatch                               string     `name:"If-Match"`
	InNoneMatch                           string     `name:"If-None-Match"`
	ServerSideEncryptionCustomerAlgorithm string     `name:"X-Amz-Server-Side-Encryption-Customer-Algorithm"`
	ServerSideEncryptionCustomerKey       string     `name:"X-Amz-server-Side-Encryption-Customer-Key"`
	ServerSideEncryptionCustomerKeyMD5    string     `name:"X-Amz-server-Side-Encryption-Customer-Key-Md5"`
}
