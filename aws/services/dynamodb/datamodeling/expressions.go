package datamodeling

import (
	"github.com/twhello/aws-to-go/aws/services/dynamodb"
)

/******************************************************************************/

// Enables adding options to a delete operation. For example, you may want to delete only
// if an attribute has a particular value.
type DeleteExpression struct {
	ConditionalOperator dynamodb.ConditionalOperator
	Expected            map[string]dynamodb.ExpectedAttributeValue
}

// Creates a new DeleteExpression.
func NewDeleteExpression(conditionalOperator dynamodb.ConditionalOperator, attributeName string, expected dynamodb.ExpectedAttributeValue) *DeleteExpression {
	return &DeleteExpression{conditionalOperator, map[string]dynamodb.ExpectedAttributeValue{attributeName: expected}}
}

// Adds one entry to the expected conditions.
func (e *DeleteExpression) AddExpectedEntry(attributeName string, expected dynamodb.ExpectedAttributeValue) {
	e.Expected[attributeName] = expected
}

/******************************************************************************/

// A query expression.
type QueryExpression struct {
	ConditionalOperator dynamodb.ConditionalOperator
	ConsistentRead      bool
	ExclusiveStartKey   map[string]dynamodb.AttributeValue
	HashKeyValues       []dynamodb.AttributeValue
	IndexName           string
	Limit               int64
	QueryFilter         map[string]dynamodb.Condition
	RangeKeyConditions  map[string]dynamodb.Condition
	ScanIndexForward    bool
}

// Creates a new QueryExpressoin.
func NewQueryExpression(hashKeyValues ...dynamodb.AttributeValue) *QueryExpression {
	return &QueryExpression{HashKeyValues: hashKeyValues, Limit: 1000}
}

// Adds a  exclusive start key for this query.
func (e *QueryExpression) AddExclusiveStartKey(attributeName string, attribute dynamodb.AttributeValue) {
	if len(e.ExclusiveStartKey) == 0 {
		e.ExclusiveStartKey = make(map[string]dynamodb.AttributeValue)
	}
	e.ExclusiveStartKey[attributeName] = attribute
}

// Adds a new filter condition to the current query filter.
func (e *QueryExpression) AddQueryFilterEntry(attributeName string, condition dynamodb.Condition) {
	if len(e.QueryFilter) == 0 {
		e.QueryFilter = make(map[string]dynamodb.Condition)
	}
	e.QueryFilter[attributeName] = condition
}

// Adds a  exclusive start key for this query.
func (e *QueryExpression) AddRangeKeyCondition(attributeName string, condition dynamodb.Condition) {
	if len(e.RangeKeyConditions) == 0 {
		e.RangeKeyConditions = make(map[string]dynamodb.Condition)
	}
	e.RangeKeyConditions[attributeName] = condition
}

// Sets the logical operator on the filter conditions of this query.
func (e *QueryExpression) SetConditionalOperator(conditionalOperator dynamodb.ConditionalOperator) {
	e.ConditionalOperator = conditionalOperator
}

// Sets whether this query uses consistent reads.
func (e *QueryExpression) SetConsistentRead(consistentRead bool) {
	e.ConsistentRead = consistentRead
}

// Sets the name of the index to be used by this query.
func (e *QueryExpression) SetIndexName(indexName string) {
	e.IndexName = indexName
}

// Sets the limit of items to return from this query.
func (e *QueryExpression) SetLimit(limit int64) {
	e.Limit = limit
}

// Sets whether this query scans forward.
func (e *QueryExpression) SetScanIndexForward(scanIndexForward bool) {
	e.ScanIndexForward = scanIndexForward
}

/******************************************************************************/

// Enables adding options to a save operation. For example, you may want to save only if
// an attribute has a particular value.
type SaveExpression struct {
	ConditionalOperator dynamodb.ConditionalOperator
	Expected            map[string]dynamodb.ExpectedAttributeValue
}

// Creates a new SaveExpression.
func NewSaveExpression(conditionalOperator dynamodb.ConditionalOperator, attributeName string, expected dynamodb.ExpectedAttributeValue) *SaveExpression {
	return &SaveExpression{conditionalOperator, map[string]dynamodb.ExpectedAttributeValue{attributeName: expected}}
}

// Adds one entry to the expected conditions.
func (e *SaveExpression) AddExpectedEntry(attributeName string, expected dynamodb.ExpectedAttributeValue) {
	e.Expected[attributeName] = expected
}

/******************************************************************************/

// Options for filtering results from a scan operation. For example, callers can specify filter conditions
// so that only objects whose attributes match different conditions are returned (see ComparisonOperator
// for more information on the available comparison types).
type ScanExpression struct {
	ConditionalOperator dynamodb.ConditionalOperator
	ExclusiveStartKey   map[string]dynamodb.AttributeValue
	Limit               int64
	ScanFilter          map[string]dynamodb.Condition
	Segment             int64
	TotalSegments       int64
}

// Creates a new ScanExpression.
func NewScanExpression(attributeName string, condition dynamodb.Condition) *ScanExpression {
	return &ScanExpression{
		ScanFilter: map[string]dynamodb.Condition{attributeName: condition},
		Limit:      1000, Segment: 0, TotalSegments: 100,
	}
}

// Adds a  exclusive start key for this scan.
func (e *ScanExpression) AddExclusiveStartKeyEntry(attributeName string, attribute dynamodb.AttributeValue) {
	if len(e.ExclusiveStartKey) == 0 {
		e.ExclusiveStartKey = make(map[string]dynamodb.AttributeValue)
	}
	e.ExclusiveStartKey[attributeName] = attribute
}

// Adds a new filter condition to the current scan filter.
func (e *ScanExpression) AddFilterConditionEntry(attributeName string, condition dynamodb.Condition) {
	e.ScanFilter[attributeName] = condition
}

// Sets the logical operator on the filter conditions of this scan.
func (e *ScanExpression) SetConditionalOperator(conditionalOperator dynamodb.ConditionalOperator) {
	e.ConditionalOperator = conditionalOperator
}

// Sets the limit of items to return from this scan.
func (e *ScanExpression) SetLimit(limit int64) {
	e.Limit = limit
}

// Sets the ID of the segment to be scanned.
func (e *ScanExpression) SetSegment(segment int64) {
	e.Segment = segment
}

// Sets the total number of segments into which the scan will be divided.
func (e *ScanExpression) SetTotalSegments(totalSegments int64) {
	e.TotalSegments = totalSegments
}
