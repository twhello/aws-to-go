package datamodeling

import (
	"github.com/twhello/aws-to-go/services/dynamodb"
)

// Struct mapper for domain-object interaction with DynamoDB.
//
// Default configuration uses UPDATE behavior for saves and EVENTUALly consistent reads,
// with no table name override and lazy-loading strategy.
type DynamoDBMapper struct {
	DynamoDBService      *dynamodb.DynamoDBService
	DynamoDBMapperConfig DynamoDBMapperConfig
}

/*****************************************************************************/

// Deletes the given object from its DynamoDB table.
func (m *DynamoDBMapper) Delete(v interface{}) (*dynamodb.DeleteItemResult, error) {
	return m.DeleteWithExpression(v, nil)
}

// Deletes the given object from its DynamoDB table using the provided deleteExpression.
func (m *DynamoDBMapper) DeleteWithExpression(v interface{}, expression *DeleteExpression) (*dynamodb.DeleteItemResult, error) {

	model := Marshal(v)

	dir := dynamodb.NewDeleteItemRequest(model.TableName)

	if attr, ok := model.Item[model.HashKey]; ok && !attr.IsEmpty() {
		dir.AddKey(model.HashKey, attr)
	}
	if attr, ok := model.Item[model.RangeKey]; ok && !attr.IsEmpty() {
		dir.AddKey(model.RangeKey, attr)
	}

	if expression != nil {
		dir.ConditionalOperator = expression.ConditionalOperator
		dir.Expected = expression.Expected
	}

	result, err := m.DynamoDBService.DeleteItem(dir)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*****************************************************************************/

// Loads an object with a hash and, if provided, range key.
func (m *DynamoDBMapper) Load(v interface{}) (*dynamodb.GetItemResult, error) {

	model := Marshal(v)

	gir := dynamodb.NewGetItemRequest(model.TableName)
	gir.ConsistentRead = m.DynamoDBMapperConfig.ConsistentReads == CONSISTANT

	if attr, ok := model.Item[model.HashKey]; ok && !attr.IsEmpty() {
		gir.AddKey(model.HashKey, attr)
	}
	if attr, ok := model.Item[model.RangeKey]; ok && !attr.IsEmpty() {
		gir.AddKey(model.RangeKey, attr)
	}
	/*
		if attr, ok := model.Item[model.IndexHashKey]; ok && !attr.IsEmpty() {
			gir.AddKey(model.IndexHashKey, attr)
		}
		if attr, ok := model.Item[model.IndexRangeKey]; ok && !attr.IsEmpty() {
			gir.AddKey(model.IndexRangeKey, attr)
		}
	*/
	result, err := m.DynamoDBService.GetItem(gir)
	if err != nil {
		return nil, err
	}

	Unmarshal(result.Item, v)

	return result, nil
}

/*****************************************************************************

//
func (m *DynamoDBMapper) Query(request dynamodb.QueryRequest, v interface{}) ([]interface{}, dynamodb.QueryResult, error) {

	return nil, dynamodb.QueryResult{}, nil
}

//
func (m *DynamoDBMapper) QueryWithExpression(request dynamodb.QueryRequest, v interface{}) ([]interface{}, dynamodb.QueryResult, error) {

	return nil, dynamodb.QueryResult{}, nil
}

//
func (m *DynamoDBMapper) QueryCount(request dynamodb.QueryRequest) (int, error) {

	return 0, nil
}
*/
/*****************************************************************************/

// Saves the object given into DynamoDB.
func (m *DynamoDBMapper) Save(v interface{}) error {
	return m.SaveWithExpression(v, nil)
}

// Saves the object given into DynamoDB, the specified saveExpression.
func (m *DynamoDBMapper) SaveWithExpression(v interface{}, expression *SaveExpression) error {

	switch m.DynamoDBMapperConfig.SaveBehavior {
	case CLOBBER:
		return m.saveClobber(v, expression)
	case UPDATE_SKIP_NULL_ATTRIBUTES:
		return m.saveUpdateSkipNullAttributes(v, expression)
	default:
		return m.saveUpdate(v, expression)
	}
}

// CLOBBER will clear and replace all attributes, included unmodeled ones, (delete and recreate) on save.
func (m *DynamoDBMapper) saveClobber(v interface{}, expression *SaveExpression) error {

	model := Marshal(v)

	pir := dynamodb.NewPutItemRequest(model.TableName)

	for k, v := range model.Item {
		if !v.IsEmpty() {
			pir.AddItem(k, v)
		}
	}

	if expression != nil {
		pir.ConditionalOperator = expression.ConditionalOperator
		pir.Expected = expression.Expected
	}

	_, err := m.DynamoDBService.PutItem(pir)
	if err != nil {
		return err
	}

	return nil
}

// UPDATE will not affect unmodeled attributes on a save operation and a null value for the
// modeled attribute will remove it from that item in DynamoDB.
func (m *DynamoDBMapper) saveUpdate(v interface{}, expression *SaveExpression) error {

	model := Marshal(v)

	uir := dynamodb.NewUpdateItemRequest(model.TableName)
	uir.Key = model.Item

	if expression != nil {
		uir.ConditionalOperator = expression.ConditionalOperator
		uir.Expected = expression.Expected
	}
	
	_, err := m.DynamoDBService.UpdateItem(uir)
	if err != nil {
		return err
	}

	return nil
}

// UPDATE_SKIP_NULL_ATTRIBUTES is similar to UPDATE, except that it ignores any null value attribute(s)
// and will NOT remove them from that item in DynamoDB.
func (m *DynamoDBMapper) saveUpdateSkipNullAttributes(v interface{}, expression *SaveExpression) error {

	model := Marshal(v)

	uir := dynamodb.NewUpdateItemRequest(model.TableName)

	for k, v := range model.Item {
		if !v.IsEmpty() {
			uir.AddKey(k, v)
		}
	}

	if expression != nil {
		uir.ConditionalOperator = expression.ConditionalOperator
		uir.Expected = expression.Expected
	}

	_, err := m.DynamoDBService.UpdateItem(uir)
	if err != nil {
		return err
	}

	return nil
}

/*****************************************************************************

//
func (m *DynamoDBMapper) Scan(request dynamodb.ScanRequest, v interface{}) ([]interface{}, dynamodb.ScanResult, error) {

	return nil, dynamodb.ScanResult{}, nil
}

//
func (m *DynamoDBMapper) ScanWithExpression(request dynamodb.ScanRequest, v interface{}) ([]interface{}, dynamodb.ScanResult, error) {

	return nil, dynamodb.ScanResult{}, nil
}

//
func (m *DynamoDBMapper) ScanCount(request dynamodb.ScanRequest) (int, error) {

	return 0, nil
}
*/