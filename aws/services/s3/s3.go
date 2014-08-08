//
// Amazon Simple Storage Service (Amazon S3) is storage for the Internet. You can use
// Amazon S3 to store and retrieve any amount of data at any time, from anywhere on
// the web. You can accomplish these tasks using the simple and intuitive web
// interface of the AWS Management Console.
//
// [http://aws.amazon.com/documentation/s3/]
//
package s3

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"github.com/twhello/aws-to-go/aws/util/netutil"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

const ServiceName = "s3"

/******************************************************************************/

// The S3 Service object. Use s3.NewService().
type S3Service struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s3 *S3Service) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s3 *S3Service) RegionName() string {
	return s3.region.Name()
}

func (s3 *S3Service) Endpoint() string {
	return s3.endpoint
}

// Low-level request to S3 service.
func (s3 *S3Service) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s3.cred, s3}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto,
		services.NewEvalServiceResponse(
			func(r io.Reader, v interface{})error { return xml.NewDecoder(r).Decode(v) },
			[]int{500, 503},
			nil,
		),
	)

	return
}

/******************************************************************************
 * S3 Service Methods for Operations on Buckets
 */

// This method creates a new bucket where owner gets FULL_CONTROL. No one else has access rights (default).
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUT.html]
func (s3 *S3Service) CreateBucket(cbr *CreateBucketRequest) (err error) {

	if s3.RegionName() != regions.DEFAULT_REGION && cbr.CreateBucketConfiguration == nil {
		cbr.CreateBucketConfiguration = &CreateBucketConfiguration{
			LocationConstraint: s3.RegionName(),
		}
	}

	req, err := services.NewServerRequest("PUT", s3.Endpoint()+"/"+cbr.BucketName, cbr.CreateBucketConfiguration)
	if err == nil {
		req.Header().Set("Content-Type", services.CONTENT_TYPE_APPLICATION_XML)
		if cbr.CannedACL != "" {
			req.Header().Add("X-Amz-Acl", string(cbr.CannedACL))
		}
		_, err = s3.SignAndDo(req, nil)
	}
	return
}

// This implementation of the DELETE operation deletes the specified bucket.
// All objects (including all object versions and delete markers) in the bucket must be deleted before the bucket itself can be deleted.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketDELETE.html]
func (s3 *S3Service) DeleteBucket(dbr *DeleteBucketRequest) (err error) {

	endpoint := s3.Endpoint()[0:8] + dbr.BucketName + "." + s3.Endpoint()[8:]

	if strings.ContainsRune(dbr.BucketName, '.') {
		endpoint = s3.Endpoint() + "/" + dbr.BucketName
	}

	req, err := services.NewClientRequest("DELETE", endpoint, nil)
	if err == nil {
		_, err = s3.SignAndDo(req, nil)
	}

	return
}

// This method is useful to determine if a bucket exists and you have permission to access it.
// The method returns a nil error if the bucket exists and you have permission to access it.
// Otherwise, the operation might return responses such as 404 Not Found and 403 Forbidden.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketHEAD.html]
func (s3 *S3Service) DoesBucketExist(bucket *Bucket) (err error) {

	req, err := services.NewClientRequest("HEAD", s3.Endpoint()+"/"+bucket.Name, nil)
	if err == nil {
		_, err = s3.SignAndDo(req, nil)
	}
	return
}

// Returns a list of all Amazon S3 buckets that the authenticated sender of the request owns.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html]
func (s3 *S3Service) ListBuckets() (buckets *ListBucketsResult, err error) {

	req, err := services.NewClientRequest("GET", s3.Endpoint(), nil)
	if err == nil {
		buckets = new(ListBucketsResult)
		_, err = s3.SignAndDo(req, buckets)
	}
	return
}

// This implementation of the GET operation returns some or all (up to 1000) of the objects in a bucket.
// You can use the request parameters as selection criteria to return a subset of the objects in a bucket.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html]
func (s3 *S3Service) ListObjects(lor *ListObjectsRequest) (objs *ListObjectsResult, err error) {

	req, err := services.NewClientRequest("GET", s3.Endpoint()+"/"+lor.BucketName, lor)
	if err == nil {
		objs = new(ListObjectsResult)
		_, err = s3.SignAndDo(req, objs)
	}

	return
}

/******************************************************************************
 * S3 Service Methods for Operations on Objects
 */

// The DELETE operation removes the null version (if there is one) of an object and inserts a delete marker,
// which becomes the current version of the object.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html]
func (s3 *S3Service) DeleteObject(dor *DeleteObjectRequest) (hdrs *DeleteObjectHeaderResponse, err error) {

	req, err := services.NewClientRequest("DELETE", s3.Endpoint()+"/"+dor.BucketName+"/"+dor.ObjectName, nil)
	if err == nil {

		resp, err := s3.SignAndDo(req, nil)
		if err == nil {
			hdrs = new(DeleteObjectHeaderResponse)
			netutil.UnmarshalHeader(resp.Header, hdrs)
		}
	}
	return
}

// The Multi-Object Delete operation enables you to delete multiple objects from a bucket using a single HTTP request.
// If you know the object keys that you want to delete, then this operation provides a suitable alternative to sending
// individual delete requests (see DELETE Object), reducing per-request overhead.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html]
func (s3 *S3Service) DeleteMultipleObjects(dmor *DeleteMultipleObjectsRequest) (result *DeleteResult, err error) {

	req, err := services.NewServerRequest("POST", s3.Endpoint()+"/"+dmor.BucketName, dmor.Headers)
	if err == nil {

		req.Header().Set("Content-Type", services.CONTENT_TYPE_APPLICATION_XML)
		req.QueryStringValues().Add("delete", "1")
		result = new(DeleteResult)
		_, err = s3.SignAndDo(req, result)
	}
	return
}

// This implementation of the GET operation retrieves objects from Amazon S3. To use GET, you must have READ access to the object.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectGET.html]
func (s3 *S3Service) GetObject(gor *GetObjectRequest) (content io.ReadCloser, hdrs *GetObjectHeaderResponse, err error) {

	req, err := services.NewClientRequest("GET", s3.Endpoint()+"/"+gor.BucketName+"/"+gor.ObjectName, gor.Constraints)
	if err == nil {

		netutil.MergeHeaders(req.Header(), netutil.MarshalHeader(gor.Constraints))

		var resp *http.Response
		resp, err = s3.SignAndDo(req, nil)
		if err == nil {
			content = resp.Body
			hdrs = new(GetObjectHeaderResponse)
			netutil.UnmarshalHeader(resp.Header, hdrs)
		}
	}

	return
}

// Gets the metadata for the specified Amazon S3 object without actually fetching the object itself.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectHEAD.html]
func (s3 *S3Service) GetObjectMetadata(gor *GetObjectRequest) (hdrs *GetObjectHeaderResponse, err error) {

	req, err := services.NewClientRequest("GET", s3.Endpoint()+"/"+gor.BucketName+"/"+gor.ObjectName, gor.Constraints)
	if err == nil {

		netutil.MergeHeaders(req.Header(), netutil.MarshalHeader(gor.Constraints))

		var resp *http.Response
		resp, err = s3.SignAndDo(req, nil)
		if err == nil {
			hdrs = new(GetObjectHeaderResponse)
			netutil.UnmarshalHeader(resp.Header, hdrs)
		}
	}

	return
}

// Uploads the specified input stream and object metadata to Amazon S3 under the specified bucket and key name.
// This implementation of the PUT operation adds an object to a bucket. You must have WRITE permissions on a bucket to add an object to it.
// [http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html]
func (s3 *S3Service) PutObject(por *PutObjectRequest) (hdrs *PutObjectHeaderResponse, err error) {

	req, err := services.NewServerRequest("PUT", s3.Endpoint()+"/"+por.BucketName+"/"+por.ObjectName, por.Content)
	if err == nil {
		
		if por.ObjectMetadata != nil {
			netutil.MergeHeaders(req.Header(), netutil.MarshalHeader(por.ObjectMetadata))
		}
		var resp *http.Response
		resp, err = s3.SignAndDo(req, nil)
		if err == nil {
			hdrs = new(PutObjectHeaderResponse)
			netutil.UnmarshalHeader(resp.Header, hdrs)
		}
	}

	return
}

/******************************************************************************/

// Creates a new S3Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *S3Service {

	subdomain := ServiceName

	if region.Name() != regions.DEFAULT_REGION {
		subdomain += "-" + region.Name()
	}

	return &S3Service{cred, region, "https://" + subdomain + ".amazonaws.com"}
}
