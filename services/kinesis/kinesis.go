//
// Amazon Kinesis is a managed service that scales elastically for real-time processing
// of streaming big data. The service takes in large streams of data records that can
// then be consumed in real time by multiple data-processing applications that can be
// run on Amazon EC2 instances.
// 
// [http://aws.amazon.com/documentation/kinesis/]
package kinesis

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"net/http"
)

const ServiceName = "kinesis"

// SWFService describes the API interface. Instantiate with swf.NewService().
type KinesisService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *KinesisService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *KinesisService) RegionName() string {
	return s.region.Name()
}

func (s *KinesisService) Endpoint() string {
	return s.endpoint
}

// Low-level request to S3 service.
// (req interfaces.IAWSRequest)
// (dto interface{})
func (s *KinesisService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

// Creates the IAWSRequest and sets required headers.
// (target string) Sets the X-Amz-Target header.
// (request interface{}) The interface to marshal into the request body.
// (result interface{}) The interface for the unmarshalled API result, or nil.
func (s *KinesisService) wrapperSignAndDo(target string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", s.Endpoint(), request)

	if err == nil {
		h := req.Header()
		h.Set("Connection", "Keep-Alive")
		h.Set("Content-Type", "application/x-amz-json-1.1")
		h.Set("X-Amz-Target", target)
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * Kinesis Service Methods.
 */

// This operation adds a new Amazon Kinesis stream to your AWS account.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_CreateStream.html]
func (s *KinesisService) CreateStream(req *CreateStreamRequest) (err error) {

	err = s.wrapperSignAndDo("Kinesis_20131202.CreateStream ", req, nil)
	return
}

// This operation deletes a stream and all of its shards and data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_DeleteStream.html]
func (s *KinesisService) DeleteStream(req *DeleteStreamRequest) (err error) {

	err = s.wrapperSignAndDo("Kinesis_20131202.DeleteStream", req, nil)
	return
}

// This operation returns the following information about the stream:
// the current status of the stream, the stream Amazon Resource Name (ARN),
// and an array of shard objects that comprise the stream.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_DescribeStream.html]
func (s *KinesisService) DescribeStream(req *DescribeStreamRequest) (result *DescribeStreamResult, err error) {

	result = new(DescribeStreamResult)
	err = s.wrapperSignAndDo("Kinesis_20131202.DescribeStream", req, result)
	return
}

// This operation returns one or more data records from a shard.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecords.html]
func (s *KinesisService) GetRecords(req *GetRecordsRequest) (result *GetRecordsResult, err error) {

	result = new(GetRecordsResult)
	err = s.wrapperSignAndDo("Kinesis_20131202.GetRecords", req, result)
	return
}

// This operation returns a shard iterator in ShardIterator.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html]
func (s *KinesisService) GetShardIterator(req *GetShardIteratorRequest) (result *GetShardIteratorResult, err error) {

	result = new(GetShardIteratorResult)
	err = s.wrapperSignAndDo("Kinesis_20131202.GetShardIterator", req, result)
	return
}

// This operation returns an array of the names of all the streams that are associated
// with the AWS account making the ListStreams request.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_ListStreams.html]
func (s *KinesisService) ListStreams(req *ListStreamsRequest) (result *ListStreamsResult, err error) {

	result = new(ListStreamsResult)
	err = s.wrapperSignAndDo("Kinesis_20131202.ListStreams", req, result)
	return
}

// This operation merges two adjacent shards in a stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_MergeShards.html]
func (s *KinesisService) MergeShards(req *MergeShardsRequest) (err error) {

	err = s.wrapperSignAndDo("Kinesis_20131202.MergeShards", req, nil)
	return
}

// This operation puts a data record into an Amazon Kinesis stream from a producer.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html]
func (s *KinesisService) PutRecord(req *PutRecordRequest) (result *PutRecordResult, err error) {

	result = new(PutRecordResult)
	err = s.wrapperSignAndDo("Kinesis_20131202.PutRecord", req, result)
	return
}

// This operation splits a shard into two new shards in the stream, to increase the stream's
// capacity to ingest and transport data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_SplitShard.html]
func (s *KinesisService) SplitShard(req *SplitShardRequest) (err error) {

	err = s.wrapperSignAndDo("Kinesis_20131202.SplitShard", req, nil)
	return
}

// Creates a new Kinesis Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *KinesisService {
	return &KinesisService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
