package datamodeling

import ()

/*	Consistent read behavior.
	CONSISTANT
	EVENTUAL */
type ConsistentReads uint

const (
	EVENTUAL ConsistentReads = iota
	CONSISTANT
)

/*	Behaviors for the save operation.
	APPEND_SET Treats scalar attributes (String, Number, Binary) the same as UPDATE_SKIP_NULL_ATTRIBUTES does.
	CLOBBER Will clear and replace all attributes, included unmodeled ones, (delete and recreate) on save.
	UPDATE Will not affect unmodeled attributes on a save operation and a null value for the modeled attribute will remove it from that item in DynamoDB.
	UPDATE_SKIP_NULL_ATTRIBUTES Is similar to UPDATE, except that it ignores any null value attribute(s) and will NOT remove them from that item in DynamoDB. */
type SaveBehavior uint

const (
	UPDATE SaveBehavior = iota
	CLOBBER
	UPDATE_SKIP_NULL_ATTRIBUTES
)

/*	Pagination loading strategy.
	EAGER_LOADING Paginated list will eagerly load all the paginated results from DynamoDB as soon as the list is initialized.
	ITERATION_ONLY Only supports using iterator to read from the paginated list.
	LAZY_LOADING Paginated list is lazily loaded when possible, and all loaded results are kept in the memory. */
type PaginationLoadingStrategy uint

const (
	LAZY_LOADING PaginationLoadingStrategy = iota
	EAGER_LOADING
	ITERATION_ONLY
)

// Configuration struct for service call behavior. An instance of this configuration is supplied to every DynamoDBMapper at construction;
// if not provided explicitly, DEFAULT is used. New instances can be given to the mapper object on individual save, load, and delete
// operations to override the defaults.
type DynamoDBMapperConfig struct {
	ConsistentReads           ConsistentReads           // Defaults to EVENTUAL
	SaveBehavior              SaveBehavior              // Defaults to UPDATE
	PaginationLoadingStrategy PaginationLoadingStrategy // Defaults to LAZY_LOADING
}
