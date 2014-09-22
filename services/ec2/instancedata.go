package ec2

import (
	"github.com/twhello/aws-to-go/services"
	"io/ioutil"
)

const (
	urlDynamic = "http://169.254.169.254/latest/dynamic/"
	urlMeta    = "http://169.254.169.254/latest/metadata/"
	urlUser    = "http://169.254.169.254/latest/user-data"
)

// The following table lists the categories of instance metadata.
const (
	META_DATA_AMI_ID            = "ami-id"            // The AMI ID used to launch the instance.
	META_DATA_AMI_LAUNCH_INDEX  = "ami-launch-index"  // Indicates the order in which the instance was launched.
	META_DATA_AMI_MANIFEST_PATH = "ami-manifest-path" // The path to the AMI's manifest file in Amazon S3.
	META_DATA_HOSTNAME          = "hostname"          // The private hostname of the instance.
	META_DATA_INSTANCE_ID       = "instance-id"       // The ID of this instance.
	META_DATA_INSTANCE_TYPE     = "instance-type"     // The type of instance.
	META_DATA_KERNAL_ID         = "kernal-id"         // The ID of the kernel launched with this instance, if applicable.
	META_DATA_LOCAL_HOSTNAME    = "local-hostname"    // The private DNS hostname of the instance.
	META_DATA_LOCAL_IPV4        = "local-ipv4"        // The private IP address of the instance.
	META_DATA_MAC               = "mac"               // The instance's media access control (MAC) address.
	META_DATA_PUBLIC_HOSTNAME   = "public-hostname"   // The instance's public DNS.
	META_DATA_PUBLIC_IPV4       = "public-ipv4"       // The public IP address.
)

// EC2 instances can also include dynamic data, such as an instance identity document
// that is generated when the instance is launched. For more information, see Dynamic Data Categories.
// [http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html#dynamic-data-categories]
// (item string) Item name. http://169.254.169.254/latest/dynamic/{item}
func InstanceDynamicData(item string) (string, error) {
	return getInstanceData(urlDynamic + item)
}

// Instance metadata is data about your instance that you can use to configure or
// manage the running instance. Instance metadata is divided into categories.
// For more information, see Instance Metadata Categories.
// [http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html#instancedata-data-categories]
// (item string) Item name. http://169.254.169.254/latest/meta-data/{item}
func InstanceMetaData(item string) (string, error) {
	return getInstanceData(urlMeta + item)
}

// You can also access the user data that you supplied when launching your instance.
// For example, you can specify parameters for configuring your instance, or attach
// a simple script.
func InstanceUserData() (string, error) {
	return getInstanceData(urlUser)
}

// Make request and process response.
func getInstanceData(path string) (v string, err error) {
	resp, err := services.HttpClient().Get(path)
	defer resp.Body.Close()
	if err == nil {
		var b []byte
		b, err = ioutil.ReadAll(resp.Body)
		v = string(b)
	}
	return
}
