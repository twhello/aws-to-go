package dynamodb

import (
	"github.com/twhello/aws-to-go/aws/util/datetime"
)

/*****************************************************************************
 * The Amazon DynamoDB API contains several data types that various actions use.
 * This section describes each data type in detail.
 */

/*****************************************************************************
 * Types and Constants
 */

/*
	ADD
	DELETE
	PUT */
type Action string

const (
	ADD    Action = "ADD"
	DELETE Action = "DELETE"
	PUT    Action = "PUT"
)

/*	ComparisonOperator Values:
	EQUAL = "EQ"
	LESS_THAN_EQUAL = "LE"
	LESS_THAN = "LT"
	GREATER_THAN_EQUAL = "GE"
	GREATER_THAN = "GT"
	NOT_NULL = "NOT_NULL"
	CONTAINS = "CONTAINS"
	NOT_CONTAINS = "NOT_CONTAINS"
	BEGINS_WITH = "BEGINS_WITH"
	IN = "IN"
	BETWEEN = "BETWEEN"
	[http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Condition.html] */
type ComparisonOperator string

const (
	EQUAL              ComparisonOperator = "EQ"
	LESS_THAN_EQUAL    ComparisonOperator = "LE"
	LESS_THAN          ComparisonOperator = "LT"
	GREATER_THAN_EQUAL ComparisonOperator = "GE"
	GREATER_THAN       ComparisonOperator = "GT"
	NOT_NULL           ComparisonOperator = "NOT_NULL"
	CONTAINS           ComparisonOperator = "CONTAINS"
	NOT_CONTAINS       ComparisonOperator = "NOT_CONTAINS"
	BEGINS_WITH        ComparisonOperator = "BEGINS_WITH"
	IN                 ComparisonOperator = "IN"
	BETWEEN            ComparisonOperator = "BETWEEN"
)

/*	ConditionalOperator Values:
	AND
	OR
	[http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Condition.html] */
type ConditionalOperator string

const (
	AND ConditionalOperator = "AND"
	OR  ConditionalOperator = "OR"
)

/*	The attribute data, consisting of the data type and the attribute value itself.
	HASH
	RANGE */
type KeyType string

const (
	HASH  KeyType = "HASH"
	RANGE KeyType = "RANGE"
)

/*	The set of attributes that are projected into the index:
	ALL - All of the table attributes are projected into the index.
	INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes are in NonKeyAttributes.
	KEYS_ONLY - Only the index and primary keys are projected into the index. */
type ProjectionType string

const (
	ALL       ProjectionType = "ALL"
	INCLUDE   ProjectionType = "INCLUDE"
	KEYS_ONLY ProjectionType = "KEYS_ONLY"
)

/*	The attributes to be returned in the result. You can retrieve all item attributes, specific item attributes,
	the count of matching items, or in the case of an index, some or all of the attributes projected into the index.
	ALL_ATTRIBUTES
	ALL_PROJECTED_ATTRIBUTES
	SPECIFIC_ATTRIBUTES
	COUNT */
type SelectAttributes string

const (
	ALL_ATTRIBUTES           SelectAttributes = "ALL_ATTRIBUTES"
	ALL_PROJECTED_ATTRIBUTES SelectAttributes = "ALL_PROJECTED_ATTRIBUTES"
	SPECIFIC_ATTRIBUTES      SelectAttributes = "SPECIFIC_ATTRIBUTES"
	COUNT                    SelectAttributes = "COUNT"
)

/*	If set to TOTAL, the response includes ConsumedCapacity data for tables and indexes. If set to INDEXES, the response includes
	ConsumedCapacity for indexes. If set to NONE (the default), ConsumedCapacity is not included in the response.
	ReturnConsumedCapacity Values:
	INDEXES
	TOTAL
	NONE_RCC */
type ReturnConsumedCapacity string

const (
	INDEXES  ReturnConsumedCapacity = "INDEXES"
	TOTAL    ReturnConsumedCapacity = "TOTAL"
	NONE_RCC ReturnConsumedCapacity = "NONE"
)

/*	If set to SIZE, statistics about item collections, if any, that were modified during the operation are returned in the response.
	If set to NONE (the default), no statistics are returned.
	ReturnItemCollectionMetrics Values:
	SIZE
	NONE_RICM */
type ReturnItemCollectionMetrics string

const (
	SIZE      ReturnItemCollectionMetrics = "SIZE"
	NONE_RICM ReturnItemCollectionMetrics = "NONE"
)

/*	Use ReturnValues if you want to get the item attributes as they appeared before they were updated with the PutItem request.
	For PutItem, the valid values are:
	NONE_RV
	ALL_OLD
	UPDATED_OLD
	ALL_NEW
	UPDATED_NEW */
type ReturnValues string

const (
	NONE_RV     ReturnValues = "NONE"
	ALL_OLD     ReturnValues = "ALL_OLD"
	UPDATED_OLD ReturnValues = "UPDATED_OLD"
	ALL_NEW     ReturnValues = "ALL_NEW"
	UPDATED_NEW ReturnValues = "UPDATED_NEW"
)

/*	The data type for the attribute.
	STRING = "S"
	NUMBER = "N"
	BINARY = "B" */
type AttributeType string

const (
	BINARY AttributeType = "B"
	NUMBER AttributeType = "N"
	STRING AttributeType = "S"
)

/*	The current state of the table or global secondary index:
	ACTIVE
	CREATING
	DELETING
	UPDATING */
type Status string

const (
	ACTIVE   Status = "ACTIVE"
	CREATING Status = "CREATING"
	DELETING Status = "DELETING"
	UPDATING Status = "UPDATING"
)

/*****************************************************************************
 * Data Types
 */

// Represents an attribute for describing the key schema for the table and indexes.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeDefinition.html]
type AttributeDefinition struct {
	AttributeName string        `json:"AttributeName"`
	AttributeType AttributeType `json:"AttributeType"`
}

// For the UpdateItem operation, represents the attributes to be modified, the action to perform on each, and the new value for each.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValueUpdate.html]
type AttributeValueUpdate struct {
	Action string          `json:"Action,omitempty"`
	Value  *AttributeValue `json:"Value,omitempty"`
}

// Represents the amount of provisioned throughput capacity consumed on a table or an index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Capacity.html]
type CapacityUnits float64

// Represents the selection criteria for a Query or Scan operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Condition.html]
type Condition struct {
	AttributeValueList []AttributeValue   `json:"AttributeValueList,omitempty"`
	ComparisonOperator ComparisonOperator `json:"ComparisonOperator"`
}

// Represents the capacity units consumed by an operation. The data returned includes the total provisioned throughput
// consumed, along with statistics for the table and any indexes involved in the operation. ConsumedCapacity is only
// returned if it was asked for in the request.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ConsumedCapacity.html]
type ConsumedCapacity struct {
	CapacityUnits          CapacityUnits            `json:"CapacityUnits,omitempty"`
	GlobalSecondaryIndexes map[string]CapacityUnits `json:"GlobalSecondaryIndexes,omitempty"`
	LocalSecondaryIndexes  map[string]CapacityUnits `json:"LocalSecondaryIndexes,omitempty"`
	Table                  CapacityUnits            `json:"Table,omitempty"`
	TableName              string                   `json:"TableName,omitempty"`
}

// Represents a request to perform a DeleteItem operation on an item.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteRequest.html]
type DeleteRequest struct {
	Key map[string]AttributeValue `json:"Key"`
}

// Represents a condition to be compared with an attribute value.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ExpectedAttributeValue.html]
type ExpectedAttributeValue struct {
	AttributeValueList []AttributeValue   `json:"AttributeValueList,omitempty"`
	ComparisonOperator ComparisonOperator `json:"ComparisonOperator,omitempty"`
	Exists             bool               `json:"Exists,omitempty"`
	Value              *AttributeValue    `json:"Value,omitempty"`
}

// Represents a global secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalSecondaryIndex.html]
type GlobalSecondaryIndex struct {
	IndexName             string                 `json:"IndexName"`
	KeySchema             []KeySchemaElement     `json:"KeySchema"`
	Projection            *Projection            `json:"Projection"`
	ProvisionedThroughput *ProvisionedThroughput `json:"ProvisionedThroughput"`
}

// Represents the properties of a global secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalSecondaryIndexDescription.html]
type GlobalSecondaryIndexDescription struct {
	IndexName             string                            `json:"IndexName,omitempty"`
	IndexSizeBytes        int64                             `json:"IndexSizeBytes,omitempty"`
	IndexStatus           Status                            `json:"IndexStatus,omitempty"`
	ItemCount             int64                             `json:"ItemCount,omitempty"`
	KeySchema             []KeySchemaElement                `json:"KeySchema,omitempty"`
	Projection            *Projection                       `json:"Projection,omitempty"`
	ProvisionedThroughput *ProvisionedThroughputDescription `json:"ProvisionedThroughput,omitempty"`
}

// Represents the new provisioned throughput settings to apply to a global secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalSecondaryIndexUpdate.html]
type GlobalSecondaryIndexUpdate struct {
	Update *UpdateGlobalSecondaryIndexAction `json:"Update,omitempty"`
}

// Information about item collections, if any, that were affected by the operation.
// ItemCollectionMetrics is only returned if it was asked for in the request.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ItemCollectionMetrics.html]
type ItemCollectionMetrics struct {
	ItemCollectionKey   map[string]AttributeValue `json:"ItemCollectionKey,omitempty"`
	SizeEstimateRangeGB []float64                 `json:"SizeEstimateRangeGB,omitempty"`
}

// Represents a single element of a key schema. A key schema specifies the attributes that
// make up the primary key of a table, or the key attributes of an index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_KeySchemaElement.html]
type KeySchemaElement struct {
	AttributeName string  `json:"AttributeName"`
	KeyType       KeyType `json:"KeyType"`
}

// Represents a set of primary keys and, for each key, the attributes to retrieve from the table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_KeysAndAttributes.html]
type KeysAndAttributes struct {
	AttributesToGet []string                    `json:"AttributesToGet,omitempty"`
	ConsistentRead  bool                        `json:"ConsistentRead,omitempty"`
	Keys            []map[string]AttributeValue `json:"Keys"`
}

// Represents a local secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_LocalSecondaryIndex.html]
type LocalSecondaryIndex struct {
	IndexName  string             `json:"IndexName"`
	KeySchema  []KeySchemaElement `json:"KeySchema"`
	Projection *Projection        `json:"Projection"`
}

// Represents the properties of a local secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_LocalSecondaryIndexDescription.html]
type LocalSecondaryIndexDescription struct {
	IndexName      string             `json:"IndexName,omitempty"`
	IndexSizeBytes int64              `json:"IndexSizeBytes,omitempty"`
	ItemCount      int64              `json:"ItemCount,omitempty"`
	KeySchema      []KeySchemaElement `json:"KeySchema,omitempty"`
	Projection     *Projection        `json:"Projection,omitempty"`
}

// Represents attributes that are copied (projected) from the table into an index. These are in
// addition to the primary key attributes and index key attributes, which are automatically projected.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Projection.html]
type Projection struct {
	NonKeyAttributes []string       `json:"NonKeyAttributes,omitempty"`
	ProjectionType   ProjectionType `json:"ProjectionType,omitempty"`
}

// Represents the provisioned throughput settings for a specified table or index.
// The settings can be modified using the UpdateTable operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html]
type ProvisionedThroughput struct {
	ReadCapacityUnits  int64 `json:"AttributeName"`
	WriteCapacityUnits int64 `json:"KeyType"`
}

// Represents the provisioned throughput settings for the table, consisting of read and write capacity
// units, along with data about increases and decreases.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughputDescription.html]
type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime   *datetime.JsonDate `json:"LastDecreaseDateTime,omitempty"`
	LastIncreaseDateTime   *datetime.JsonDate `json:"LastIncreaseDateTime,omitempty"`
	NumberOfDecreasesToday int64              `json:"NumberOfDecreasesToday,omitempty"`
	ReadCapacityUnits      int64              `json:"AttributeName,omitempty"`
	WriteCapacityUnits     int64              `json:"KeyType,omitempty"`
}

// Represents a request to perform a PutItem operation on an item.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutRequest.html]
type PutRequest struct {
	Item map[string]AttributeValue `json:"Item"`
}

// Represents the properties of a table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html]
type TableDescription struct {
	AttributeDefinitions   []AttributeDefinition             `json:"AttributeDefinitions,omitempty"`
	CreationDateTime       *datetime.JsonDate                `json:"CreationDateTime,omitempty"`
	GlobalSecondaryIndexes *GlobalSecondaryIndexDescription  `json:"GlobalSecondaryIndexes,omitempty"`
	ItemCount              int                               `json:"ItemCount,omitempty"`
	KeySchema              []KeySchemaElement                `json:"KeySchema,omitempty"`
	LocalSecondaryIndexes  *LocalSecondaryIndexDescription   `json:"LocalSecondaryIndexes,omitempty"`
	ProvisionedThroughput  *ProvisionedThroughputDescription `json:"ProvisionedThroughput,omitempty"`
	TableName              string                            `json:"TableName,omitempty"`
	TableSizeBytes         int64                             `json:"TableSizeBytes,omitempty"`
	TableStatus            Status                            `json:"TableStatus,omitempty"`
}

// Represents the new provisioned throughput settings to be applied to a global secondary index.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateGlobalSecondaryIndexAction.html]
type UpdateGlobalSecondaryIndexAction struct {
	IndexName             string                            `json:"IndexName"`
	ProvisionedThroughput *ProvisionedThroughputDescription `json:"ProvisionedThroughput"`
}

// Represents an operation to perform - either DeleteItem or PutItem.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_WriteRequest.html]
type WriteRequest struct {
	DeleteRequest *DeleteRequest `json:DeleteRequest,omitempty`
	PutRequest    *PutRequest    `json:PutRequest,omitempty`
}

/*****************************************************************************
 * Result Data Types
 */

// Represents the output of a BatchGetItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItemResult.html]
type BatchGetItemResult struct {
	ConsumedCapacity ConsumedCapacity             `json:"ConsumedCapacity,omitempty"`
	Responses        map[string]AttributeValue    `json:"Responses,omitempty"`
	UnprocessedKeys  map[string]KeysAndAttributes `json:"UnprocessedKeys,omitempty"`
}

// Represents the output of a BatchWriteItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItemResult.html]
type BatchWriteItemResult struct {
	ConsumedCapacity      ConsumedCapacity                 `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics map[string]ItemCollectionMetrics `json:"ItemCollectionMetrics,omitempty"`
	UnprocessedItems      map[string]WriteRequest          `json:"UnprocessedItems,omitempty"`
}

// Represents the output of a CreateTable operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTableResult.html]
type CreateTableResult struct {
	TableDescription *TableDescription `json:"TableDescription,omitempty"`
}

// Represents the output of a DeleteItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItemResult.html]
type DeleteItemResult struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics    `json:"ItemCollectionMetrics,omitempty"`
}

// Represents the output of a DeleteTable operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteTableResult.html]
type DeleteTableResult struct {
	TableDescription `json:"TableDescription,omitempty"`
}

// Represents the output of a DescribeTable operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTableResult.html]
type DescribeTableResult struct {
	Table TableDescription `json:"Table,omitempty"`
}

// Represents the output of a GetItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItemResult.html]
type GetItemResult struct {
	ConsumedCapacity ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	Item             map[string]AttributeValue `json:"Item,omitempty"`
}

// Represents the output of a ListTables operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListTablesResult.html]
type ListTablesResult struct {
	LastEvaluatedTableName string   `json:"LastEvaluatedTableName"`
	TableNames             []string `json:"TableNames"`
}

// Represents the output of a PutItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItemResult.html]
type PutItemResult struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics    `json:"ItemCollectionMetrics,omitempty"`
}

// Represents the output of a Query operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_QueryResult.html]
type QueryResult struct {
	ConsumedCapacity ConsumedCapacity            `json:"ConsumedCapacity,omitempty"`
	Count            int                         `json:"Items,omitempty"`
	Items            []map[string]AttributeValue `json:"Items,omitempty"`
	LastEvaluatedKey map[string]AttributeValue   `json:"LastEvaluatedKey,omitempty"`
	ScannedCount     int                         `json:"Items,omitempty"`
}

// Represents the output of a Scan operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ScanResult.html]
type ScanResult struct {
	ConsumedCapacity ConsumedCapacity            `json:"ConsumedCapacity,omitempty"`
	Count            int                         `json:"Items,omitempty"`
	Items            []map[string]AttributeValue `json:"Items,omitempty"`
	LastEvaluatedKey map[string]AttributeValue   `json:"LastEvaluatedKey,omitempty"`
	ScannedCount     int                         `json:"Items,omitempty"`
}

// Represents the output of an UpdateItem operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItemResult.html]
type UpdateItemResult struct {
	Attributes            map[string]AttributeValue `json:"Attributes,omitempty"`
	ConsumedCapacity      ConsumedCapacity          `json:"ConsumedCapacity,omitempty"`
	ItemCollectionMetrics *ItemCollectionMetrics    `json:"ItemCollectionMetrics,omitempty"`
}

// Represents the output of an UpdateTable operation.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTableResult.html]
type UpdateTableResult struct {
	TableDescription *TableDescription `json:"TableDescription,omitempty"`
}
