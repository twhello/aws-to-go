//
// aws-to-go - Go packages to interact with the Amazon Web Services.
//
// https://github.com/twhello/aws-to-go/
//
// Copyright (c) 2014 TwHello, Inc.
//
// Written by Patrick Martin <info@twhello.com>
//
package aws

import (
	"github.com/twhello/aws-to-go/aws/auth"
	"github.com/twhello/aws-to-go/aws/regions"
	"github.com/twhello/aws-to-go/aws/services/datapipeline"
	"github.com/twhello/aws-to-go/aws/services/dynamodb"
	"github.com/twhello/aws-to-go/aws/services/kinesis"
	"github.com/twhello/aws-to-go/aws/services/s3"
	"github.com/twhello/aws-to-go/aws/services/ses"
	"github.com/twhello/aws-to-go/aws/services/simpledb"
	"github.com/twhello/aws-to-go/aws/services/sns"
	"github.com/twhello/aws-to-go/aws/services/sqs"
	"github.com/twhello/aws-to-go/aws/services/swf"
	"os"
)

// The AWS Client struct.
type Client struct {
	AccessKeyId string
	SecretKey   string
	RegionName  string
}

// Creates a new Client from specified credentials and region.
// (accessKeyId string) AWS access key.
// (secretKey string) AWS secret key.
// (regionName *string) Name of region. nil defaults to US_EAST_1.
func NewClient(accessKeyId, secretKey string, regionName *string) Client {
	return Client{accessKeyId, secretKey, *regionName}
}

// Creates a new Client from environmental variables AWS_ACCESS_KEY_ID and AWS_SECRET_KEY.
// (regionName *string) Name of region. nil defaults to env var AWS_REGION or US_EAST_1.
func NewClientByEnvars(regionName *string) Client {
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	if regionName == nil {
		rn := os.Getenv("AWS_REGION")
		if rn == "" {
			rn = "us-east-1"
		}
		regionName = &rn
	}
	return Client{accessKeyId, secretKey, *regionName}
}

// Creates a new Client from a local PROPERTIES file.
// (path *string) Path to file. nil defaults to "AWSCredentials.properties".
// (regionName *string) Name of region. nil defaults to US_EAST_1 or file setting.
func NewClientByFile(path, regionName *string) Client {
	var accessKeyId string
	var secretKey string
	return Client{accessKeyId, secretKey, *regionName}
}

/******************************************************************************
 * Service Access Methods
 */

// AWS Data Pipeline is a web service that you can use to automate the movement and transformation of data.
// With AWS Data Pipeline, you can define data-driven workflows, so that tasks can be dependent on the
// successful completion of previous tasks.
// [http://aws.amazon.com/documentation/data-pipeline/]
func (c Client) DataPipeling() *datapipeline.DataPipelineService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return datapipeline.NewService(cred, regions.Config(c.RegionName))
}


// Amazon DynamoDB is a fully managed NoSQL database service that provides fast and
// predictable performance with seamless scalability. You can use Amazon DynamoDB to
// create a database table that can store and retrieve any amount of data, and serve
// any level of request traffic. Amazon DynamoDB automatically spreads the data and
// traffic for the table over a sufficient number of servers to handle the request
// capacity specified by the customer and the amount of data stored, while maintaining
// consistent and fast performance.
// [http://aws.amazon.com/documentation/dynamodb/]
func (c Client) DynamoDB() *dynamodb.DynamoDBService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return dynamodb.NewService(cred, regions.Config(c.RegionName))
}

// Amazon Kinesis is a managed service that scales elastically for real-time processing
// of streaming big data. The service takes in large streams of data records that can
// then be consumed in real time by multiple data-processing applications that can be
// run on Amazon EC2 instances.
// [http://aws.amazon.com/documentation/kinesis/]
func (c Client) Kinesis() *kinesis.KinesisService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return kinesis.NewService(cred, regions.Config(c.RegionName))
}

// Amazon Simple Storage Service (Amazon S3) is storage for the Internet. You can use
// Amazon S3 to store and retrieve any amount of data at any time, from anywhere on
// the web. You can accomplish these tasks using the simple and intuitive web
// interface of the AWS Management Console.
// [http://aws.amazon.com/documentation/s3/]
func (c Client) S3() *s3.S3Service {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return s3.NewService(cred, regions.Config(c.RegionName))
}

// Amazon SES is an outbound-only email-sending service that provides an easy,
// cost-effective way for you to send email.
// [http://aws.amazon.com/documentation/ses/]
func (c Client) SES() *ses.SESService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return ses.NewService(cred, regions.Config(c.RegionName))
}

// Amazon Simple Notification Service (Amazon SNS) is a web service that enables applications,
// end-users, and devices to instantly send and receive notifications from the cloud.
// [http://aws.amazon.com/documentation/sns/]
func (c Client) SNS() *sns.SNSService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return sns.NewService(cred, regions.Config(c.RegionName))
}

// Amazon SimpleDB is a web service for running queries on structured data in real time.
// This service works in close conjunction with Amazon Simple Storage Service (Amazon S3)
// and Amazon Elastic Compute Cloud (Amazon EC2), collectively providing the ability to
// store, process and query data sets in the cloud. These services are designed to make
// web-scale computing easier and more cost-effective for developers.
// [http://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/Welcome.html]
func (c Client) SimpleDB() *simpledb.SDBService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return simpledb.NewService(cred, regions.Config(c.RegionName))
}

// Amazon Simple Queue Service (Amazon SQS)  is a messaging queue service that handles message
// or workflows between other components in a system.
// [http://aws.amazon.com/documentation/sqs/]
func (c Client) SQS() *sqs.SQSService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return sqs.NewService(cred, regions.Config(c.RegionName))
}

// Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications that
// coordinate work across distributed components. In Amazon SWF, a task represents a
// logical unit of work that is performed by a component of your application. Coordinating
// tasks across the application involves managing intertask dependencies, scheduling,
// and concurrency in accordance with the logical flow of the application. Amazon SWF
// gives you full control over implementing tasks and coordinating them without worrying
// about underlying complexities such as tracking their progress and maintaining their state.
// [http://aws.amazon.com/documentation/swf/]
func (c Client) SWF() *swf.SWFService {
	cred := auth.NewCredentials(c.AccessKeyId, c.SecretKey)
	return swf.NewService(cred, regions.Config(c.RegionName))
}
