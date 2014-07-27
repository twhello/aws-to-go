package util

import (
	"github.com/twhello/aws-to-go/aws/services/dynamodb"
	"errors"
	"time"
)

// Checks if a specified table exists and is in ACTIVE state.
// (tableName string) The name of the table whose status is being checked.
func DoesTableExist(dynamo *dynamodb.DynamoDBService, tableName string) bool {

	dtr := dynamodb.NewDescribeTableRequest(tableName)
	result, err := dynamo.DescribeTable(dtr)
	return err == nil || result.Table.TableStatus == dynamodb.ACTIVE
}

// Blocks up to a specified amount of time for a specified AWS DynamoDB table to move into the ACTIVE state.
// If the table doesn't transition to the ACTIVE state, then an error is thrown.
// (dynamo *dynamodb.DynamoDBService) The AWS DynamoDB Service to use to make requests.
// (tableName string) The name of the table whose status is being checked.
// (timeout time.Duration) The maximum number of milliseconds to wait.
// (interval uint) The poll interval in milliseconds.
func WaitForTableToBecomeActive(dynamo *dynamodb.DynamoDBService, tableName string, timeout uint, interval uint) error {

	now := time.Now()

	for !DoesTableExist(dynamo, tableName) {

		if time.Since(now) > time.Millisecond*time.Duration(timeout) {
			return errors.New("")
		}
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
	return nil
}
