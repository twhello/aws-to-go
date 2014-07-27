package dynamodb

import ()

/*****************************************************************************/

// The BatchGetItem operation returns the attributes of one or more items from one or more tables.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItem.html]
type BatchGetItemRequest struct {
	RequestItems           map[string]KeysAndAttributes `json:"RequestItems"`
	ReturnConsumedCapacity ReturnConsumedCapacity       `json:"ReturnConsumedCapacity,omitempty"`
}

// Creates a new BatchGetItemRequest.
// (requestItems map[string]KeysAndAttributes) A map of one or more table names and corresponding primary.
func NewBatchGetItemRequest(requestItems map[string]KeysAndAttributes) *BatchGetItemRequest {
	return &BatchGetItemRequest{RequestItems: requestItems}
}

/*****************************************************************************/

// The BatchWriteItem operation puts or deletes multiple items in one or more tables.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html]
type BatchWriteItemRequest struct {
	RequestItems                map[string][]WriteRequest   `json:"RequestItems"`
	ReturnConsumedCapacity      ReturnConsumedCapacity      `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics `json:"ReturnItemCollectionMetrics,omitempty"`
}

// Creates a new BatchWriteItemRequest.
// (map[string][]WriteRequest) A map of one or more table names and lists of operations (DeleteRequest or PutRequest).
func NewBatchWriteItemRequest(requestItems map[string][]WriteRequest) *BatchWriteItemRequest {
	return &BatchWriteItemRequest{RequestItems: requestItems}
}

/*****************************************************************************/

// The CreateTable operation adds a new table to your account.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html]
type CreateTableRequest struct {
	AttributeDefinitions   []AttributeDefinition  `json:"AttributeDefinitions"`
	GlobalSecondaryIndexes []GlobalSecondaryIndex `json:"GlobalSecondaryIndexes,omitempty"`
	LocalSecondaryIndexes  []LocalSecondaryIndex  `json:"LocalSecondaryIndexes,omitempty"`
	ProvisionedThroughput  *ProvisionedThroughput `json:"ProvisionedThroughput"`
	TableName              string                 `json:"TableName"`
}

// Creates a new CreateTableRequest.
// (tableName string) The name of the table to create.
// (readCapacityUnits int64) The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException.
// (writeCapacityUnits int64) The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.
func NewCreateTableRequest(tableName string, readCapacityUnits int64, writeCapacityUnits int64) *CreateTableRequest {
	return &CreateTableRequest{
		TableName: tableName,
		ProvisionedThroughput: &ProvisionedThroughput{
			ReadCapacityUnits:  readCapacityUnits,
			WriteCapacityUnits: writeCapacityUnits,
		},
	}
}

// Adds an AttributeDefinition to the AttributeDefinitions array.
// (attributeName string) A name for the attribute.
// (attributeType AttributeType) The data type for the attribute. Valid Values: S | N | B
func (r *CreateTableRequest) AddAttributeDefinition(attributeName string, attributeType AttributeType) {
	r.AttributeDefinitions = append(r.AttributeDefinitions, AttributeDefinition{attributeName, attributeType})
}

/*****************************************************************************/

// Deletes a single item in a table by primary key.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html]
type DeleteItemRequest struct {
	ConditionalOperator         ConditionalOperator               `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	Key                         map[string]AttributeValue         `json:"Key"`
	ReturnConsumedCapacity      ReturnConsumedCapacity            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics       `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                ReturnValues                      `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
}

// Creates a new DeleteItemRequest with the primary key hash attribute.
// (tableName string) The name of the table from which to delete the item.
func NewDeleteItemRequest(tableName string) *DeleteItemRequest {
	return &DeleteItemRequest{TableName: tableName, Key: make(map[string]AttributeValue)}
}

// Adds an attribute for the primary key. For the primary key, you must provide all of the attributes.
// (name string) The attribute name.
// (attribute AttributeValue) Represents the data for an attribute.
func (r *DeleteItemRequest) AddKey(name string, attribute AttributeValue) {
	r.Key[name] = attribute
}

/*****************************************************************************/

// The DeleteTable operation deletes a table and all of its items.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteTable.html]
type DeleteTableRequest struct {
	TableName string `json:"TableName"`
}

// Creates a new DeleteTableRequest.
// (tableName string) The name of the table to delete.
func NewDeleteTableRequest(tableName string) *DeleteTableRequest {
	return &DeleteTableRequest{tableName}
}

/*****************************************************************************/

// Returns information about the table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTable.html]
type DescribeTableRequest struct {
	TableName string `json:"TableName"`
}

// Creates a new DescribeTableRequest.
// (tableName string) The name of the table to describe.
func NewDescribeTableRequest(tableName string) *DescribeTableRequest {
	return &DescribeTableRequest{tableName}
}

/*****************************************************************************/

// The GetItem operation returns a set of attributes for the item with the given primary key.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html]
type GetItemRequest struct {
	AttributesToGet        []string                  `json:"AttributesToGet,omitempty"`
	ConsistentRead         bool                      `json:"ConsistentRead,omitempty"`
	Key                    map[string]AttributeValue `json:"Key"`
	ReturnConsumedCapacity string                    `json:"ReturnConsumedCapacity,omitempty"`
	TableName              string                    `json:"TableName"`
}

// Creates a new GetItemRequest with the primary key hash attribute.
// (tableName string) The name of the table containing the requested item.
func NewGetItemRequest(tableName string) *GetItemRequest {
	return &GetItemRequest{TableName: tableName, Key: make(map[string]AttributeValue)}
}

// Adds an attribute for the primary key. For the primary key, you must provide all of the attributes.
// (name string) The attribute name.
// (attribute AttributeValue) Represents the data for an attribute.
func (r *GetItemRequest) AddKey(name string, attribute AttributeValue) {
	r.Key[name] = attribute
}

/*****************************************************************************/

// Returns an array of table names associated with the current account and endpoint.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListTables.html]
type ListTablesRequest struct {
	ExclusiveStartTableName string `json:"ExclusiveStartTableName,omitempty"`
	Limit                   int    `json:"Limit,omitempty"`
}

// Creates a new ListTablesRequest.
// (limit int) A maximum number of table names to return.
func NewListTablesRequest(limit int) *ListTablesRequest {
	return &ListTablesRequest{Limit: limit}
}

/*****************************************************************************/

// Creates a new item, or replaces an old item with a new item.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html]
type PutItemRequest struct {
	ConditionalOperator         ConditionalOperator               `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	Item                        map[string]AttributeValue         `json:"Item"`
	ReturnConsumedCapacity      ReturnConsumedCapacity            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics       `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                ReturnValues                      `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
}

// Creates a new PutItemRequest.
// (tableName string) The name of the table to contain the item.
func NewPutItemRequest(tableName string) *PutItemRequest {
	return &PutItemRequest{TableName: tableName, Item: make(map[string]AttributeValue)}
}

// Add an Item.
// (name string) The attribute name.
// (attribute AttributeValue) Represents the data for an attribute.
func (r *PutItemRequest) AddItem(name string, attribute AttributeValue) {
	r.Item[name] = attribute
}

/*****************************************************************************/

// QueryRequest struct.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Query.html]
type QueryRequest struct {
	AttributesToGet        []string                  `json:"AttributesToGet,omitempty"`
	ConditionalOperator    ConditionalOperator       `json:"ConditionalOperator,omitempty"`
	ConsistentRead         bool                      `json:"ConsistentRead,omitempty"`
	ExclusiveStartKey      map[string]AttributeValue `json:"ExclusiveStartKey,omitempty"`
	IndexName              string                    `json:"IndexName,omitempty"`
	KeyConditions          map[string]Condition      `json:"KeyConditions"`
	Limit                  int                       `json:"Limit,omitempty"`
	QueryFilter            map[string]Condition      `json:"QueryFilter,omitempty"`
	ReturnConsumedCapacity ReturnConsumedCapacity    `json:"ReturnConsumedCapacity,omitempty"`
	ScanIndexForward       bool                      `json:"ScanIndexForward,omitempty"`
	Select                 SelectAttributes          `json:"Select,omitempty"`
	TableName              string                    `json:"TableName"`
}

// Create a new QueryRequest object.
// (tableName string) The name of the table containing the requested items.
func NewQueryRequest(tableName string) *QueryRequest {
	return &QueryRequest{TableName: tableName, KeyConditions: map[string]Condition{}}
}

// Add a KeyCondition.
// (name string) The name of the KeyCondition.
// (attr AttributeValue) Values to evaluate against the supplied attribute.
// (op ComparisonOperator) The comparator for evaluating attributes.
func (r *QueryRequest) AddKeyCondition(name string, attribute AttributeValue, op ComparisonOperator) {
	if len(r.KeyConditions) == 0 {
		r.KeyConditions = make(map[string]Condition)
	}
	r.KeyConditions[name] = Condition{[]AttributeValue{attribute}, op}
}

// Add a QueryFilter.
// (name string) The name of the QueryFilter.
// (attr AttributeValue) Values to evaluate against the supplied attribute.
// (op ComparisonOperator) The comparator for evaluating attributes.
func (r *QueryRequest) AddQueryFilter(name string, attribute AttributeValue, op ComparisonOperator) {
	if len(r.QueryFilter) == 0 {
		r.QueryFilter = make(map[string]Condition)
	}
	r.QueryFilter[name] = Condition{[]AttributeValue{attribute}, op}
}

/*****************************************************************************/

// The Scan operation returns one or more items and item attributes by accessing every item in the table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Scan.html]
type ScanRequest struct {
	AttributesToGet        []string                  `json:"AttributesToGet,omitempty"`
	ConditionalOperator    ConditionalOperator       `json:"ConditionalOperator,omitempty"`
	ExclusiveStartKey      map[string]AttributeValue `json:"ExclusiveStartKey,omitempty"`
	Limit                  int                       `json:"Limit,omitempty"`
	ReturnConsumedCapacity ReturnConsumedCapacity    `json:"ReturnConsumedCapacity,omitempty"`
	ScanFilter             map[string]Condition      `json:"ScanFilter,omitempty"`
	Segment                int64                     `json:"Segment,omitempty"`
	Select                 SelectAttributes          `json:"Select,omitempty"`
	TableName              string                    `json:"TableName"`
	TotalSegments          int                       `json:"TotalSegments,omitempty"`
}

// Create a new QueryRequest object.
// (tableName string) The name of the table containing the requested items.
func NewScanRequest(tableName string) *ScanRequest {
	return &ScanRequest{TableName: tableName}
}

// Add a ScanFilter.
// (name string) The name of the ScanFilter.
// (attr AttributeValue) Values to evaluate against the supplied attribute.
// (op ComparisonOperator) The comparator for evaluating attributes.
func (r *ScanRequest) AddScanFilter(name string, attribute AttributeValue, op ComparisonOperator) {
	if len(r.ScanFilter) == 0 {
		r.ScanFilter = make(map[string]Condition)
	}
	r.ScanFilter[name] = Condition{[]AttributeValue{attribute}, op}
}

/*****************************************************************************/

// Edits an existing item's attributes, or inserts a new item if it does not already exist.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html]
type UpdateItemRequest struct {
	ConditionalOperator         ConditionalOperator               `json:"ConditionalOperator,omitempty"`
	Expected                    map[string]ExpectedAttributeValue `json:"Expected,omitempty"`
	Key                         map[string]AttributeValue         `json:"Key"`
	ReturnConsumedCapacity      ReturnConsumedCapacity            `json:"ReturnConsumedCapacity,omitempty"`
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics       `json:"ReturnItemCollectionMetrics,omitempty"`
	ReturnValues                ReturnValues                      `json:"ReturnValues,omitempty"`
	TableName                   string                            `json:"TableName"`
}

// Creates a new UpdateItemRequest.
// (tableName string)
func NewUpdateItemRequest(tableName string) *UpdateItemRequest {
	return &UpdateItemRequest{TableName: tableName, Key: make(map[string]AttributeValue)}
}

// Adds an attribute for the primary key. For the primary key, you must provide all of the attributes.
// (name string) The attribute name.
// (attribute AttributeValue) Represents the data for an attribute.
func (r *UpdateItemRequest) AddKey(name string, attribute AttributeValue) {
	r.Key[name] = attribute
}

/*****************************************************************************/

// Updates the provisioned throughput for the given table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html]
type UpdateTableRequest struct {
	GlobalSecondaryIndexUpdates *GlobalSecondaryIndexUpdate `json:"GlobalSecondaryIndexUpdates,omitempty"`
	ProvisionedThroughput       *ProvisionedThroughput      `json:"ProvisionedThroughput,omitempty"`
	TableName                   string                      `json:"TableName"`
}

// Creates a new UpdateTableRequest.
// (tableName string) The name of the table to update.
// (readCapacityUnits int64) The maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException.
// (writeCapacityUnits int64) The maximum number of writes consumed per second before DynamoDB returns a ThrottlingException.
func NewUpdateTableRequest(tableName string, readCapacityUnits int64, writeCapacityUnits int64) *UpdateTableRequest {
	return &UpdateTableRequest{
		TableName: tableName,
		ProvisionedThroughput: &ProvisionedThroughput{
			ReadCapacityUnits:  readCapacityUnits,
			WriteCapacityUnits: writeCapacityUnits,
		},
	}
}
