package kinesis

import ()

/*****************************************************************************/

// This operation adds a new Amazon Kinesis stream to your AWS account.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_CreateStream.html]
type CreateStreamRequest struct {
	ShardCount int    `json:"ShardCount"`
	StreamName string `json:"StreamName"`
}

// Creates a new CreateStreamRequest.
func NewCreateStreamRequest(streamName string, shardCount int) *CreateStreamRequest {
	return &CreateStreamRequest{shardCount, streamName}
}

/*****************************************************************************/

// This operation deletes a stream and all of its shards and data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_DeleteStream.html]
type DeleteStreamRequest struct {
	StreamName string `json:"StreamName"`
}

// Creates a new DeleteStreamRequest.
func NewDeleteStreamRequest(streamName string) *DeleteStreamRequest {
	return &DeleteStreamRequest{streamName}
}

/*****************************************************************************/

// This operation returns the following information about the stream:
// the current status of the stream, the stream Amazon Resource Name (ARN),
// and an array of shard objects that comprise the stream.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_DescribeStream.html]
type DescribeStreamRequest struct {
	ExclusiveStartShardId string `json:"ExclusiveStartShardId,omitempty"`
	Limit                 int    `json:"Limit,omitempty"`
	StreamName            string `json:"StreamName"`
}

// Creates a new DescribeStreamRequest.
func NewDescribeStreamRequest(streamName string) *DescribeStreamRequest {
	return &DescribeStreamRequest{StreamName: streamName}
}

/*****************************************************************************/

// This operation returns one or more data records from a shard.
// A GetRecords operation request can retrieve up to 10 MB of data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecords.html]
type GetRecordsRequest struct {
	Limit         int    `json:"Limit,omitempty"`
	ShardIterator string `json:"ShardIterator"`
}

// Creates a new GetRecordsRequest.
func NewGetRecordsRequest(shardIterator string) *GetRecordsRequest {
	return &GetRecordsRequest{ShardIterator: shardIterator}
}

/*****************************************************************************/

// This operation returns a shard iterator in ShardIterator.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html]
type GetShardIteratorRequest struct {
	ShardId                string            `json:"ShardId"`
	ShardIteratorType      ShardIteratorType `json:"ShardIteratorType"`
	StartingSequenceNumber string            `json:"StartingSequenceNumber,omitempty"`
	StreamName             string            `json:"StreamName"`
}

// Creates a new GetShardIteratorRequest.
func NewGetShardIteratorRequest(streamName, shardId string, shardIteratorType ShardIteratorType) *GetShardIteratorRequest {
	return &GetShardIteratorRequest{
		StreamName: streamName, ShardId: shardId, ShardIteratorType: shardIteratorType,
	}
}

/*****************************************************************************/

// This operation returns an array of the names of all the streams that are associated
// with the AWS account making the ListStreams request.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_ListStreams.html]
type ListStreamsRequest struct {
	ExclusiveStartStreamName string `json:"ExclusiveStartStreamName,omitempty"`
	Limit                    int    `json:"Limit,omitempty"`
}

// Creates a new ListStreamsRequest.
func NewListStreamsRequest() *ListStreamsRequest {
	return &ListStreamsRequest{}
}

/*****************************************************************************/

// This operation merges two adjacent shards in a stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_MergeShards.html]
type MergeShardsRequest struct {
	AdjacentShardToMerge string `json:"AdjacentShardToMerge"`
	ShardToMerge         string `json:"ShardToMerge"`
	StreamName           string `json:"StreamName"`
}

// Creates a new MergeShardsRequest.
func NewMergeShardsRequest(streamName, shardToMerge, adjacentShardToMerge string) *MergeShardsRequest {
	return &MergeShardsRequest{adjacentShardToMerge, shardToMerge, streamName}
}

/*****************************************************************************/

// This operation puts a data record into an Amazon Kinesis stream from a producer.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html]
type PutRecordRequest struct {
	Data                      []byte `json:"Data"`
	ExplicitHashKey           string `json:"ExplicitHashKey,omitempty"`
	PartitionKey              string `json:"PartitionKey"`
	SequenceNumberForOrdering string `json:"SequenceNumberForOrdering,omitempty"`
	StreamName                string `json:"StreamName"`
}

// Creates a new PutRecordRequest.
func NewPutRecordRequest(streamName, partitionKey string, data []byte) *PutRecordRequest {
	return &PutRecordRequest{Data: data, PartitionKey: partitionKey, StreamName: streamName}
}

/*****************************************************************************/

// This operation splits a shard into two new shards in the stream, to increase
// the stream's capacity to ingest and transport data.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_SplitShard.html]
type SplitShardRequest struct {
	NewStartingHashKey string `json:"NewStartingHashKey"`
	ShardToSplit       string `json:"ShardToSplit"`
	StreamName         string `json:"StreamName"`
}

// Creates a new SplitShardRequest.
func NewSplitShardRequest(streamName, shardToSplit, newStartingHashKey string) *SplitShardRequest {
	return &SplitShardRequest{newStartingHashKey, shardToSplit, streamName}
}
