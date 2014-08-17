aws-to-go
=====================

Go packages to interact with Amazon Web Services.

***Notice:** These packages are NOT production ready. Testing and code coverage is minimal. Most errors will likely be typos or misconfigured marshalling tags in the datatypes.*

Here at Twhello, we are unabashed fans of the Go language and Amazon Web Services.  Unfortunately, there is no official AWS SDK for Go. While there are some promising alternatives, they are inconsistent and incomplete. We decided to write our own AWS SDK for Go and share it with the community.

Our intention is to add services, write test files and make aws-to-go production-ready as quickly as possible. If you find (and fix) any issues, please report them. Contributors are welcome.

> **Goals of aws-to-go:**
>
> - Always ask: *What would Amazon do?*
> - Mirror the official [<i class="icon-share"></i> AWS Documentation][aws_docs] as closely as possible.
> - Seek completeness.
> - A unified interface for all API requests for easy extensibility.
> - First-class documentation.

### **Available Services**
The following packages are available:

- Auto Scaling [<i class="icon-share"></i>][aws_autoscaling_docs]
- Amazon CloudWatch [<i class="icon-share"></i>][aws_cloudwatch_docs]
- Amazon Cognito [<i class="icon-share"></i>][aws_cognito_docs]
- AWS Data Pipeline [<i class="icon-share"></i>][aws_datapipeline_docs]
- Amazon DynamoDB [<i class="icon-share"></i>][aws_dynamodb_docs]
- Amazon EC2 [<i class="icon-share"></i>][aws_ec2_docs]
- Amazon Kinesis [<i class="icon-share"></i>][aws_kinesis_docs]
- Amazon Simple Storage Service (S3) [<i class="icon-share"></i>][aws_s3_docs]
- Amazon Simple Email Service (SES) [<i class="icon-share"></i>][aws_ses_docs]
- Amazon SimpleDB [<i class="icon-share"></i>][aws_simpledb_docs]
- Amazon Simple Notification Service (SNS) [<i class="icon-share"></i>][aws_sns_docs]
- Amazon Simple Queue Service (SQS) [<i class="icon-share"></i>][aws_sqs_docs]
- Amazon Simple Workflow Service (SWF) [<i class="icon-share"></i>][aws_swf_docs]

### **How to build and install aws-to-go**
Use `go get` to pull aws-to-go packages from the repository.

To pull the convenience package for all services:
`$ go get github.com/twhello/aws-to-go/aws`

To pull an individual service:
`$ go get github.com/twhello/aws-to-go/services/s3`

### **Documentation**
Documentation is available at http://godoc.org/github.com/twhello/aws-to-go

### **Code Samples**
Setup your AWS credentials and client:

    client := aws.Client{"AWS_ACCESS_KEY", "AWS_SECRET_KEY", regions.US_EAST_1}

#### <i class="icon-file"></i>S3 Code Samples

Create a new bucket.
```go
req := s3.NewCreateBucketRequest("s3-bucket")
err := client.S3().CreateBucket(req)
```
List buckets.
```go
list, err := client.S3().ListBuckets()
```
List objects in the specified bucket.
```go
req := s3.NewListObjectsRequest("s3-bucket")
result2, err := client.S3().ListObjects(req)
```
Delete the specified bucket.
```go
req := s3.NewDeleteBucketRequest("s3-bucket")
err := client.S3().DeleteBucket(req)
```
Retrieve the contents and header info from the specified file.
```go
req := s3.NewGetObjectRequest("s3-bucket", "path/to/file.txt")
content, headers, err := client.S3().GetObject(req)
b, _ := ioutil.ReadAll(content)
content.Close()
```
Put a new file in the specified bucket and file path.
```go
req := s3.NewPutObjectRequest(
	"s3-bucket",
	"path/to/file.txt",
	bytes.NewReader([]byte("This is content of the file.")),
	&s3.ObjectMetadata{ContentType:"text/plain"},
)
headers, err := client.S3().PutObject(req)
```

#### <i class="icon-file"></i>DynamoDB Mapper Code Samples
The following code samples saves to and loads from a DynamoDB table.
```go
type StructVal struct {
	ValueS string
	ValueI int
	Value  bool
}

type MapperTable struct {
	Hash      string    `DynamoDBHashKey:"Hash"`
	Range     string    `DynamoDBRangeKey:"Range"`
	StrVal    string    `DynamoDBAttribute:"StrVal"`
	IntVal    int       `DynamoDBAttribute:"IntVal"`
	FloatVal  float64   `DynamoDBAttribute:"FloatVal"`
	BoolVal   bool      `DynamoDBAttribute:"BoolValS" DynamoDBType:"S"`
	BoolVal2  bool      `DynamoDBAttribute:"BoolValN" DynamoDBType:"N"`
	BoolVal3  bool      `DynamoDBAttribute:"BoolValB" DynamoDBType:"B"`
	TimeValS  time.Time `DynamoDBAttribute:"TimeValS" DynamoDBType:"S"`
	TimeValN  time.Time `DynamoDBAttribute:"TimeValN" DynamoDBType:"N"`
	TimeValB  time.Time `DynamoDBAttribute:"TimeValB" DynamoDBType:"B"`
	StructVal StructVal `DynamoDBAttribute:"StructVal"`
}

mapper := datamodeling.DynamoDBMapper{DynamoDBService: client.DynamoDB()}

dataIn := &MapperTable{
    "HashValue", "RangeValue123",
    "String Value", 123, 123.45,
    false, true, false,
    time.Now(), time.Now(), time.Now(),
    StructVal{"Value", 123, true}
}
err := mapper.Save(dataIn)

dataOut := &MapperTable{Hash: "HashValue", Range: "RangeValue123"}
result, err := mapper.Load(dataOut)
```

----------

The MIT License (MIT)
Copyright (c) 2014 Twhello, Inc.

[aws_docs]: http://aws.amazon.com/documentation/
[aws_autoscaling_docs]: http://aws.amazon.com/documentation/autoscaling/
[aws_cloudwatch_docs]: http://aws.amazon.com/documentation/cloudwatch/
[aws_cognito_docs]: http://aws.amazon.com/documentation/cognito/
[aws_datapipeline_docs]: http://aws.amazon.com/documentation/datapipeline/
[aws_dynamodb_docs]: http://aws.amazon.com/documentation/dynamodb/
[aws_ec2_docs]: http://aws.amazon.com/documentation/ec2/
[aws_kinesis_docs]: http://aws.amazon.com/documentation/kinesis/
[aws_s3_docs]: http://aws.amazon.com/documentation/s3/
[aws_ses_docs]: http://aws.amazon.com/documentation/ses/
[aws_simpledb_docs]: http://aws.amazon.com/documentation/simpledb/
[aws_sns_docs]: http://aws.amazon.com/documentation/sns/
[aws_sqs_docs]: http://aws.amazon.com/documentation/sqs/
[aws_swf_docs]: http://aws.amazon.com/documentation/swf/
