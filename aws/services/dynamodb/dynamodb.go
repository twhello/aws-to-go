//
// Amazon DynamoDB is a fully managed NoSQL database service that provides fast and
// predictable performance with seamless scalability. You can use Amazon DynamoDB to
// create a database table that can store and retrieve any amount of data, and serve
// any level of request traffic. Amazon DynamoDB automatically spreads the data and
// traffic for the table over a sufficient number of servers to handle the request
// capacity specified by the customer and the amount of data stored, while maintaining
// consistent and fast performance.
//
// [http://aws.amazon.com/documentation/dynamodb/]
//
package dynamodb

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/interfaces"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services"
	"net/http"
)

const ServiceName = "dynamodb"

// DynamoDB Service struct. Use [dynamodb.NewService(interfaces.IAWSCredentials, *regions.Region) *DynamoDBService].
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/Welcome.html]
type DynamoDBService struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (db *DynamoDBService) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (db *DynamoDBService) RegionName() string {
	return db.region.Name()
}

func (db *DynamoDBService) Endpoint() string {
	return db.endpoint
}

// Low-level request to S3 service.
func (db *DynamoDBService) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{db.cred, db}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalJsonServiceResponse())

	return
}

func (db *DynamoDBService) wrapperSignAndDo(target string, request, result interface{}) (err error) {

	req, err := services.NewServerRequest("POST", db.Endpoint(), request)

	if err == nil {
		req.Header().Set("Connection", "Keep-Alive")
		req.Header().Set("Content-Type", "application/x-amz-json-1.0")
		req.Header().Set("X-Amz-Target", target)
		_, err = db.SignAndDo(req, result)
	}

	return
}

// The BatchGetItem operation returns the attributes of one or more items from one or more tables.
// You identify requested items by primary key.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItem.html]
func (db *DynamoDBService) BatchGetItem(bgir *BatchGetItemRequest) (result *BatchGetItemResult, err error) {

	result = new(BatchGetItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.BatchGetItem", bgir, result)
	return
}

// The BatchWriteItem operation puts or deletes multiple items in one or more tables. A single call to BatchWriteItem
// can write up to 1 MB of data, which can comprise as many as 25 put or delete requests.
// Individual items to be written can be as large as 64 KB.
// Note: BatchWriteItem cannot update items. To update items, use the UpdateItem API.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html]
func (db *DynamoDBService) BatchWriteItem(bwir *BatchWriteItemRequest) (result *BatchWriteItemResult, err error) {

	result = new(BatchWriteItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.BatchWriteItem", bwir, result)
	return
}

// The CreateTable operation adds a new table to your account. In an AWS account, table names must be unique within each region.
// That is, you can have two tables with same name if you create the tables in different regions.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html]
func (db *DynamoDBService) CreateTable(ctr *CreateTableRequest) (result *CreateTableResult, err error) {

	result = new(CreateTableResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.CreateTable", ctr, result)
	return
}

// Deletes a single item in a table by primary key. You can perform a conditional delete operation that deletes the
// item if it exists, or if it has an expected attribute value.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html]
func (db *DynamoDBService) DeleteItem(dir *DeleteItemRequest) (result *DeleteItemResult, err error) {

	result = new(DeleteItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.DeleteItem", dir, result)
	return
}

// The DeleteTable operation deletes a table and all of its items. After a DeleteTable request, the specified table is in the
// DELETING state until DynamoDB completes the deletion.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteTable.html]
func (db *DynamoDBService) DeleteTable(dtr *DeleteTableRequest) (result *DeleteTableResult, err error) {

	result = new(DeleteTableResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.DeleteTable", dtr, result)
	return
}

// Returns information about the table, including the current status of the table, when it was created,
// the primary key schema, and any indexes on the table.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTable.html]
func (db *DynamoDBService) DescribeTable(dtr *DescribeTableRequest) (result *DescribeTableResult, err error) {

	result = new(DescribeTableResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.DescribeTable", dtr, result)
	return
}

// The GetItem operation returns a set of attributes for the item with the given primary key.
// If there is no matching item, GetItem does not return any data.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html]
func (db *DynamoDBService) GetItem(gir *GetItemRequest) (result *GetItemResult, err error) {

	result = new(GetItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.GetItem", gir, result)
	return
}

// Returns an array of table names associated with the current account and endpoint. The output from ListTables is paginated,
// with each page returning a maximum of 100 table names.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListTables.html]
func (db *DynamoDBService) ListTables(ltr *ListTablesRequest) (result *ListTablesResult, err error) {

	result = new(ListTablesResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.ListTables", ltr, result)
	return
}

// Creates a new item, or replaces an old item with a new item. If an item already exists in the specified table
// with the same primary key, the new item completely replaces the existing item. You can perform a conditional put
// (insert a new item if one with the specified primary key doesn't exist), or replace an existing item if it has
// certain attribute values.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html]
func (db *DynamoDBService) PutItem(pir *PutItemRequest) (result *PutItemResult, err error) {

	result = new(PutItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.PutItem", pir, result)
	return
}

// A Query operation directly accesses items from a table using the table primary key, or from an index
// using the index key. You must provide a specific hash key value. You can narrow the scope of the query
// by using comparison operators on the range key value, or on the index key. You can use the ScanIndexForward
// parameter to get results in forward or reverse order, by range key or by index key.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Query.html]
func (db *DynamoDBService) Query(qr *QueryRequest) (result *QueryResult, err error) {

	result = new(QueryResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.Query", qr, result)
	return
}

// The Scan operation returns one or more items and item attributes by accessing every item in the table.
// To have DynamoDB return fewer items, you can provide a ScanFilter.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Scan.html]
func (db *DynamoDBService) Scan(sr *ScanRequest) (result *ScanResult, err error) {

	result = new(ScanResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.Scan", sr, result)
	return
}

// Edits an existing item's attributes, or inserts a new item if it does not already exist. You can put, delete,
// or add attribute values. You can also perform a conditional update (insert a new attribute name-value pair if
// it doesn't exist, or replace an existing name-value pair if it has certain expected attribute values).
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html]
func (db *DynamoDBService) UpdateItem(uir *UpdateItemRequest) (result *UpdateItemResult, err error) {

	result = new(UpdateItemResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.UpdateItem", uir, result)
	return
}

// Updates the provisioned throughput for the given table. Setting the throughput for a table helps you manage performance and is
// part of the provisioned throughput feature of DynamoDB.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html]
func (db *DynamoDBService) UpdateTable(utr *UpdateTableRequest) (result *UpdateTableResult, err error) {

	result = new(UpdateTableResult)
	err = db.wrapperSignAndDo("DynamoDB_20120810.UpdateTable", utr, result)
	return
}

/******************************************************************************
 * DynamoDB Initializers.
 */

// Creates a new DynamoDB Service.
func NewService(cred interfaces.IAWSCredentials, region *regions.Region) *DynamoDBService {
	return &DynamoDBService{cred, region, "https://" + ServiceName + "." + region.Name() + ".amazonaws.com"}
}
