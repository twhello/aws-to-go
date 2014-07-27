//
// Amazon SimpleDB is a web service for running queries on structured data in real time.
// This service works in close conjunction with Amazon Simple Storage Service (Amazon S3)
// and Amazon Elastic Compute Cloud (Amazon EC2), collectively providing the ability to
// store, process and query data sets in the cloud. These services are designed to make
// web-scale computing easier and more cost-effective for developers.
//
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/Welcome.html]
//
package simpledb

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"github.com/twhello/aws-to-go/aws/util/netutil"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

const ServiceName = "sdb"

// The SimpleDB Service object. Use sdb.NewService().
type SDBService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *SDBService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *SDBService) RegionName() string {
	return s.region.Name()
}

func (s *SDBService) Endpoint() string {
	return s.endpoint
}

// Low-level request to SimpleDB service.
func (s *SDBService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto,
		services.NewEvalServiceResponse(
			func(r io.Reader, v interface{})error { return xml.NewDecoder(r).Decode(v) },
			[]int{408, 500, 503},
			nil,
		),
	)

	return
}

func (s *SDBService) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	qs := netutil.MarshalValues(request)
	qs.Add("AWSAccessKeyId", s.cred.AccessKeyId())
	qs.Add("Timestamp", time.Now().UTC().Format(time.RFC3339)) // ISO 8601 2006-01-02T15:04:05.999Z
	qs.Add("Version", "2009-04-15")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", s.Endpoint(), qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * SimpleDB Service Methods.
 */

// Performs multiple DeleteAttributes operations in a single call,
// which reduces round trips and latencies.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchDeleteAttributes.html]
func (s *SDBService) BatchDeleteAttributes(req *BatchDeleteAttributesRequest) (result *BatchDeleteAttributesResponse, err error) {

	result = new(BatchDeleteAttributesResponse)
	err = s.wrapperSignAndDo("BatchDeleteAttributes", req, result)
	return
}

// With the BatchPutAttributes operation, you can perform multiple
// PutAttribute operations in a single call.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchPutAttributes.html]
func (s *SDBService) BatchPutAttributes(req *BatchPutAttributesRequest) (result *BatchPutAttributesResponse, err error) {

	result = new(BatchPutAttributesResponse)
	err = s.wrapperSignAndDo("BatchPutAttributes", req, result)
	return
}

// The CreateDomain operation creates a new domain.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_CreateDomain.html]
func (s *SDBService) CreateDomain(req *CreateDomainRequest) (result *CreateDomainResponse, err error) {

	result = new(CreateDomainResponse)
	err = s.wrapperSignAndDo("CreateDomain", req, result)
	return
}

// Deletes one or more attributes associated with the item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteAttributes.html]
func (s *SDBService) DeleteAttributes(req *DeleteAttributesRequest) (result *DeleteAttributesResponse, err error) {

	result = new(DeleteAttributesResponse)
	err = s.wrapperSignAndDo("DeleteAttributes", req, result)
	return
}

// The DeleteDomain operation deletes a domain.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteDomain.html]
func (s *SDBService) DeleteDomain(req *DeleteDomainRequest) (result *DeleteDomainResponse, err error) {

	result = new(DeleteDomainResponse)
	err = s.wrapperSignAndDo("DeleteDomain", req, result)
	return
}

// Returns information about the domain, including when the domain was created,
// the number of items and attributes, and the size of attribute names and values.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DomainMetadata.html]
func (s *SDBService) DomainMetadata(req *DomainMetadataRequest) (result *DomainMetadataResponse, err error) {

	result = new(DomainMetadataResponse)
	err = s.wrapperSignAndDo("DomainMetadata", req, result)
	return
}

// Returns all of the attributes associated with the item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_GetAttributes.html]
func (s *SDBService) GetAttributes(req *GetAttributesRequest) (result *GetAttributesResponse, err error) {

	result = new(GetAttributesResponse)
	err = s.wrapperSignAndDo("GetAttributes", req, result)
	return
}

// The ListDomains operation lists all domains associated with the Access Key ID.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_ListDomains.html]
func (s *SDBService) ListDomains(req *ListDomainsRequest) (result *ListDomainsResponse, err error) {

	result = new(ListDomainsResponse)
	err = s.wrapperSignAndDo("ListDomains", req, result)
	return
}

// The PutAttributes operation creates or replaces attributes in an item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_PutAttributes.html]
func (s *SDBService) PutAttributes(req *PutAttributesRequest) (result *PutAttributesResponse, err error) {

	result = new(PutAttributesResponse)
	err = s.wrapperSignAndDo("PutAttributes", req, result)
	return
}

// The Select operation returns a set of Attributes for ItemNames that match the select expression.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_Select.html]
func (s *SDBService) Select(req *SelectRequest) (result *SelectResponse, err error) {

	result = new(SelectResponse)
	err = s.wrapperSignAndDo("Select", req, result)
	return
}

// Creates a new SSimpleDBS Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *SDBService {
	return &SDBService{cred, region, "https://" + ServiceName + ".amazonaws.com"}
}
