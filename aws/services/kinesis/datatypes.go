package kinesis

import ()

/******************************************************************************
 * Constants
 */

// The current status of the stream being described.
// ACTIVE - The stream exists and is ready for read and write operations or deletion.
// CREATING - The stream is being created.
// DELETING - The stream is being deleted.
// UPDATING - Shards in the stream are being merged or split.
type StreamStatus string

const (
	ACTIVE   StreamStatus = "ACTIVE"
	CREATING StreamStatus = "CREATING"
	DELETING StreamStatus = "DELETING"
	UPDATING StreamStatus = "UPDATING"
)

// Determines how the shard iterator is used to start reading data records from the shard.
// AT_SEQUENCE_NUMBER - Start reading exactly from the position denoted by a specific sequence number.
// AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted by a specific sequence number.
// TRIM_HORIZON - Start reading at the last untrimmed record in the shard in the system, which is the oldest data record in the shard.
// LATEST - Start reading just after the most recent record in the shard, so that you always read the most recent data in the shard.
type ShardIteratorType string

const (
	AFTER_SEQUENCE_NUMBER ShardIteratorType = "AFTER_SEQUENCE_NUMBER"
	AT_SEQUENCE_NUMBER    ShardIteratorType = "AT_SEQUENCE_NUMBER"
	LATEST                ShardIteratorType = "LATEST"
	TRIM_HORIZON          ShardIteratorType = "TRIM_HORIZON"
)

/******************************************************************************
 * Data Types
 */

// Represents the output of a DescribeStream operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_DescribeStreamResult.html]
type DescribeStreamResult struct {
	StreamDescription StreamDescription `json:"StreamDescription"`
}

// Represents the output of a GetRecords operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecordsResult.html]
type GetRecordsResult struct {
	NextShardIterator string   `json:"NextShardIterator,omitempty"`
	Records           []Record `json:"Records"`
}

// Represents the output of a GetShardIterator operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIteratorResult.html]
type GetShardIteratorResult struct {
	ShardIterator string `json:"ShardIterator,omitempty"`
}

// The range of possible hash key values for the shard, which is a set of ordered
// contiguous positive integers.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_HashKeyRange.html]
type HashKeyRange struct {
	EndingHashKey   string `json:"EndingHashKey"`
	StartingHashKey string `json:"StartingHashKey"`
}

// Represents the output of a ListStreams operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_ListStreamsResult.html]
type ListStreamsResult struct {
	HasMoreStreams bool     `json:"HasMoreStreams"`
	StreamNames    []string `json:"StreamNames"`
}

// Represents the output of a PutRecord operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecordResult.html]
type PutRecordResult struct {
	SequenceNumber string `json:"SequenceNumber"`
	ShardId        string `json:"SharedId"`
}

// The unit of data of the Amazon Kinesis stream, which is composed of a sequence
// number, a partition key, and a data blob.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_Record.html]
type Record struct {
	Data           []byte `json:"Data"`
	PartitionKey   string `json:"PartitionKey"`
	SequenceNumber string `json:"SequenceNumber"`
}

// The range of possible sequence numbers for the shard.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_SequenceNumberRange.html]
type SequenceNumberRange struct {
	EndingSequenceNumber   string `json:"EndingSequenceNumber,omitempty"`
	StartingSequenceNumber string `json:"StartingSequenceNumber"`
}

// A uniquely identified group of data records in an Amazon Kinesis stream.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_Shard.html]
type Shard struct {
	AdjacentParentShardId string              `json:"AdjacentParentShardId,omitempty"`
	HashKeyRange          HashKeyRange        `json:"HashKeyRange"`
	ParentShardId         string              `json:"ParentShardId,omitempty"`
	SequenceNumberRange   SequenceNumberRange `json:"SequenceNumberRange"`
	ShardId               string              `json:"ShardId"`
}

// Represents the output of a DescribeStream operation.
// [http://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescription.html]
type StreamDescription struct {
	HasMoreShards bool         `json:"HasMoreShards"`
	Shards        []Shard      `json:"Shards"`
	StreamARN     string       `json:"StreamARN"`
	StreamName    string       `json:"StreamName"`
	StreamStatus  StreamStatus `json:"StreamStatus"`
}
