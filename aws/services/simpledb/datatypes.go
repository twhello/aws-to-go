package simpledb

import ()

/******************************************************************************
 * Data Types
 */

type Attribute struct {
	Name    string `name:"Name" xml:"Name"`
	Value   string `name:"Value" xml:"Value"`
	Replace bool   `name:"Replace,omitempty" xml:"Replace,omitempty"`
}

type DomainMetadataResult struct {
	ItemCount                int64 `xml:"ItemCount"`
	ItemNamesSizeBytes       int64 `xml:"ItemNamesSizeBytes"`
	AttributeNameCount       int64 `xml:"AttributeNameCount"`
	AttributeNamesSizeBytes  int64 `xml:"AttributeNamesSizeBytes"`
	AttributeValueCount      int64 `xml:"AttributeValueCount"`
	AttributeValuesSizeBytes int64 `xml:"AttributeValuesSizeBytes"`
	Timestamp                int64 `xml:"Timestamp"`
}

type GetAttributesResult struct {
	Attributes []Attribute `xml:"Attribute"`
}

type Expected struct {
	Name   string `name:"Name"`
	Value  string `name:"Value"`
	Exists bool   `name:"Exists"`
}

type ListDomainsResult struct {
	DomainNames []string `xml:"DomainName"`
	NextToken   string   `xml:"NextToken"`
}

type Item struct {
	Name       string      `name:"ItemName" xml:"Name"`
	Attributes []Attribute `name:"Attribute.#.,omitempty" xml:"Attribute"`
}

type ResponseMetadata struct {
	RequestId string  `xml:"RequestId"`
	BoxUsage  float64 `xml:"BoxUsage"`
}

type SelectResult struct {
	Items []Item `xml:"Item"`
}

/******************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchDeleteAttributes.html]
type BatchDeleteAttributesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_BatchPutAttributes.html]
type BatchPutAttributesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_CreateDomain.html]
type CreateDomainResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteAttributes.html]
type DeleteAttributesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DeleteDomain.html]
type DeleteDomainResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_DomainMetadata.html]
type DomainMetadataResponse struct {
	DomainMetadataResult DomainMetadataResult `xml:"DomainMetadataResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_GetAttributes.html]
type GetAttributesResponse struct {
	GetAttributesResult GetAttributesResult `xml:"GetAttributesResult"`
	ResponseMetadata    ResponseMetadata    `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_ListDomains.html]
type ListDomainsResponse struct {
	ListDomainsResult ListDomainsResult `xml:"ListDomainsResult"`
	ResponseMetadata  ResponseMetadata  `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_PutAttributes.html]
type PutAttributesResponse struct {
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/SDB_API_Select.html]
type SelectResponse struct {
	SelectResult     SelectResult     `xml:"SelectResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}
