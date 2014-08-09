package simpledb

import ()

/*****************************************************************************/

// Performs multiple DeleteAttributes operations in a single call,
// which reduces round trips and latencies.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchDeleteAttributes.html]
type BatchDeleteAttributesRequest struct {
	DomainName string `name:"DomainName"`
	Items      []Item `name:"Item.#.,omitempty" base:"1"`
}

// Creates a new BatchDeleteAttributesRequest.
func NewBatchDeleteAttributesRequest(domainName string) *BatchDeleteAttributesRequest {
	return &BatchDeleteAttributesRequest{DomainName: domainName}
}

// Add an Item. Do not set Replace in the Attributes.
func (r *BatchDeleteAttributesRequest) AddItem(name string, attributes ...Attribute) {
	item := Item{Name: name, Attributes: attributes}
	r.Items = append(r.Items, item)
}

/*****************************************************************************/

// With the BatchPutAttributes operation, you can perform multiple
// PutAttribute operations in a single call.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchPutAttributes.html]
type BatchPutAttributesRequest struct {
	DomainName string `name:"DomainName"`
	Items      []Item `name:"Item.#." base:"1"`
}

// Creates a new BatchPutAttributesRequest.
func NewBatchPutAttributesRequest(domainName string) *BatchPutAttributesRequest {
	return &BatchPutAttributesRequest{DomainName: domainName}
}

// Add an Item. Optionally use Replace in the Attributes.
func (r *BatchPutAttributesRequest) AddItem(name string, attributes ...Attribute) {
	item := Item{Name: name, Attributes: attributes}
	r.Items = append(r.Items, item)
}

/*****************************************************************************/

// The CreateDomain operation creates a new domain.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_CreateDomain.html]
type CreateDomainRequest struct {
	DomainName string `name:"DomainName"`
}

// Creates a new CreateDomainRequest.
func NewCreateDomainRequest(domainName string) *CreateDomainRequest {
	return &CreateDomainRequest{domainName}
}

/*****************************************************************************/

// Deletes one or more attributes associated with the item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteAttributes.html]
type DeleteAttributesRequest struct {
	Attributes []Attribute `name:"Attribute.#.,omitempty" base:"1"`
	DomainName string      `name:"DomainName"`
	Expected   *Expected    `name:"Expected.,omitempty"`
	ItemName   string      `name:"ItemName"`
}

// Creates a new DeleteAttributesRequest.
func NewDeleteAttributesRequest(domainName, itemName string) *DeleteAttributesRequest {
	return &DeleteAttributesRequest{DomainName: domainName, ItemName: itemName}
}

// Add an Attribute.
func (r *DeleteAttributesRequest) AddAttribute(name, value string) {
	r.Attributes = append(r.Attributes, Attribute{name, value, false})
}

// Set the Expected check.
func (r *DeleteAttributesRequest) SetExpected(name, value string, exists bool) {
	r.Expected = &Expected{name, value, exists}
}

/*****************************************************************************/

// The DeleteDomain operation deletes a domain.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteDomain.html]
type DeleteDomainRequest struct {
	DomainName string `name:"DomainName"`
}

// Creates a new DeleteDomainRequest.
func NewDeleteDomainRequest(domainName string) *DeleteDomainRequest {
	return &DeleteDomainRequest{domainName}
}

/*****************************************************************************/

// Returns information about the domain, including when the domain was created,
// the number of items and attributes, and the size of attribute names and values.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DomainMetadata.html]
type DomainMetadataRequest struct {
	DomainName string `name:"DomainName"`
}

// Creates a new DomainMetadataRequest.
func NewDomainMetadataRequest(domainName string) *DomainMetadataRequest {
	return &DomainMetadataRequest{domainName}
}

/*****************************************************************************/

// Returns all of the attributes associated with the item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_GetAttributes.html]
type GetAttributesRequest struct {
	AttributeName  string `name:"AttributeName,omitempty"`
	ConsistentRead bool   `name:"ConsistentRead,omitempty"`
	DomainName     string `name:"DomainName"`
	ItemName       string `name:"ItemName"`
}

// Creates a new GetAttributesRequest.
func NewGetAttributesRequest(domainName, itemName string) *GetAttributesRequest {
	return &GetAttributesRequest{DomainName: domainName, ItemName: itemName}
}

/*****************************************************************************/

// The ListDomains operation lists all domains associated with the Access Key ID.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_ListDomains.html]
type ListDomainsRequest struct {
	MaxNumberOfDomains int    `name:"MaxNumberOfDomains,omitempty"`
	NextToken          string `name:"NextToken,omitempty"`
}

// Creates a new ListDomainsRequest.
func NewListDomainsRequest() *ListDomainsRequest {
	return &ListDomainsRequest{}
}

/*****************************************************************************/

// The PutAttributes operation creates or replaces attributes in an item.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_PutAttributes.html]
type PutAttributesRequest struct {
	Attributes []Attribute `name:"Attribute.#." base:"1"`
	DomainName string      `name:"DomainName"`
	Expected   *Expected    `name:"Expected.,omitempty"`
	ItemName   string      `name:"ItemName"`
}

// Creates a new PutAttributesRequest.
func NewPutAttributesRequest(domainName, itemName string) *PutAttributesRequest {
	return &PutAttributesRequest{DomainName: domainName, ItemName: itemName}
}

// Add an Attribute.
func (r *PutAttributesRequest) AddAttribute(name, value string, replace bool) {
	r.Attributes = append(r.Attributes, Attribute{name, value, replace})
}

// Set the Expected check.
func (r *PutAttributesRequest) SetExpected(name, value string, exists bool) {
	r.Expected = &Expected{name, value, exists}
}

/*****************************************************************************/

// The Select operation returns a set of Attributes for ItemNames that match the select expression.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_Select.html]
type SelectRequest struct {
	ConsistentRead   bool   `name:"ConsistentRead,omitempty"`
	NextToken        string `name:"NextToken,omitempty"`
	SelectExpression string `name:"SelectExpression"`
}

// Creates a new SelectRequest.
func NewSelectRequest(selectExpression string) *SelectRequest {
	return &SelectRequest{SelectExpression: selectExpression}
}
