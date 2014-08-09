package ec2

import (
	"time"
)

/******************************************************************************
 * Constants
 */

// The following are the supported account attributes.
// 	DEFAULT_VPC - The ID of the default VPC for your account, or none.
//	SUPPORTED_PLATFORMS - Indicates whether your account can launch instances into EC2-Classic and EC2-VPC, or only into EC2-VPC.
type AccountAttribute string

const (
	DEFAULT_VPC         AccountAttribute = "default-vpc"
	SUPPORTED_PLATFORMS AccountAttribute = "supported-platforms"
)

// The architecture of the image.
//	I386 Architecture = "i386"
//	X86_64 Architecture = "x86_64"
type Architecture string

const (
	I386   Architecture = "i386"
	X86_64 Architecture = "x86_64"
)

// The container format used to combine disk images with metadata (such as OVF).
//	OVA = "ova"
type ContainerFormat string

const (
	OVA ContainerFormat = "ova"
)

// The currency of the Reserved Instance offering you are purchasing.
//	USD CurrencyCode = "USD"
type CurrencyCode string

const (
	USD CurrencyCode = "USD"
)

// The format for the exported image.
//	VHD = "vhd"
//	VMDK = "vdk"
type DiskImageFormat string

const (
	VHD  DiskImageFormat = "vhd"
	VMDK DiskImageFormat = "vdk"
)

// The domain type.
//	STANDARD = "standard"
//	VPC = "vpc"
type Domain string

const (
	STANDARD Domain = "standard"
	VPC      Domain = "vpc"
)

// The frequency of the recurring charge.
// HOURLY = "hourly"
type Frequency string

const (
	HOURLY Frequency = "hourly"
)

// A reason code that describes a specific instance's health state.
//	INSTANCE_STUCK_IN_STATE HealthState = "instance-stuck-in-state"
//	UNRESPONSIVE HealthState = "unresponsive"
//	NOT_ACCEPTING_CREDENTIALS HealthState = "not-accepting-credentials"
//	PASSWORD_NOT_AVAILABLE HealthState = "password-not-available"
//	PERFORMANCE_NETWORK HealthState = "performance-network"
//	PERFORMANCE_INSTANCE_STORE HealthState = "performance-instance-store"
//	PERFORMANCE_EBS_VOLUME HealthState = "performance-ebs-volume"
//	PERFORMANCE_OTHER HealthState = "performance-other"
//	OTHER HealthState = "other"
type HealthState string

const (
	INSTANCE_STUCK_IN_STATE    HealthState = "instance-stuck-in-state"
	UNRESPONSIVE               HealthState = "unresponsive"
	NOT_ACCEPTING_CREDENTIALS  HealthState = "not-accepting-credentials"
	PASSWORD_NOT_AVAILABLE     HealthState = "password-not-available"
	PERFORMANCE_NETWORK        HealthState = "performance-network"
	PERFORMANCE_INSTANCE_STORE HealthState = "performance-instance-store"
	PERFORMANCE_EBS_VOLUME     HealthState = "performance-ebs-volume"
	PERFORMANCE_OTHER          HealthState = "performance-other"
	OTHER                      HealthState = "other"
)

// The hypervisor type of the image.
//	OVM = "ovm"
//	XEN = "xen"
type Hypervisor string

const (
	OVM Hypervisor = "ovm"
	XEN Hypervisor = "xen"
)

// The AMI attribute.
//	AMI_DESCRIPTION = "description"
//	AMI_BLOCK_DEVICE_MAPPING = "blockDeviceMapping"
//	AMI_LAUNCH_PERMISSION = "launchPermission"
//	AMI_PRODUCT_CODES = "productCodes"
//	AMI_KERNEL = "kernel"
//	AMI_RAMDISK = "ramdisk"
//	AMI_SRIOV_NET_SUPPORT = "sriovNetSupport"
type ImageAttribute string

const (
	AMI_DESCRIPTION          ImageAttribute = "description"
	AMI_BLOCK_DEVICE_MAPPING ImageAttribute = "blockDeviceMapping"
	AMI_LAUNCH_PERMISSION    ImageAttribute = "launchPermission"
	AMI_PRODUCT_CODES        ImageAttribute = "productCodes"
	AMI_KERNEL               ImageAttribute = "kernel"
	AMI_RAMDISK              ImageAttribute = "ramdisk"
	AMI_SRIOV_NET_SUPPORT    ImageAttribute = "sriovNetSupport"
)

// The type of image.
//	KERNEL = "kernel"
//	MACHINE = "machine"
//	RAMDISK = "ramdisk"
type ImageType string

const (
	KERNEL  ImageType = "kernel"
	MACHINE ImageType = "machine"
	RAMDISK ImageType = "ramdisk"
)

// The instance attribute.
//	ATTR_BLOCK_DEVICE_MAPPING = "blockDeviceMapping"
//	ATTR_DISABLE_API_TERMINATION = "disableApiTermination"
//	ATTR_EBS_OPTIMIZED = "ebsOptimized"
//	ATTR_GROUP_SET = "groupSet"
//	ATTR_INSTANCE_INITIATED_SHUTDOWN_BEHAVIOR = "instanceInitiatedShutdownBehavior"
//	ATTR_INSTANCE_TYPE = "instanceType"
//	ATTR_KERNEL = "kernel"
//	ATTR_PRODUCT_CODES = "productCodes"
//	ATTR_RAMDISK = "ramdisk"
//	ATTR_ROOT_DEVICE_NAME = "rootDeviceName"
//	ATTR_SOURCE_DEST_CHECK = "sourceDestCheck"
//	ATTR_SRIOV_NET_SUPPORT = "sriovNetSupport"
//	ATTR_USER_DATA = "userData"
type InstanceAttribute string

const (
	INSTANCE_ATTACHMENT                           InstanceAttribute = "attachment"
	INSTANCE_AUTO_ENABLE_IO                       InstanceAttribute = "autoEnableIO"
	INSTANCE_BLOCK_DEVICE_MAPPING                 InstanceAttribute = "blockDeviceMapping"
	INSTANCE_CREATE_VOLUME_PERMISSION             InstanceAttribute = "createVolumePermission"
	INSTANCE_DESCRIPTION                          InstanceAttribute = "description"
	INSTANCE_DISABLE_API_TERMINATION              InstanceAttribute = "disableApiTermination"
	INSTANCE_EBS_OPTIMIZED                        InstanceAttribute = "ebsOptimized"
	INSTANCE_GROUP_SET                            InstanceAttribute = "groupSet"
	INSTANCE_INSTANCE_INITIATED_SHUTDOWN_BEHAVIOR InstanceAttribute = "instanceInitiatedShutdownBehavior"
	INSTANCE_INSTANCE_TYPE                        InstanceAttribute = "instanceType"
	INSTANCE_KERNEL                               InstanceAttribute = "kernel"
	INSTANCE_PRODUCT_CODES                        InstanceAttribute = "productCodes"
	INSTANCE_RAMDISK                              InstanceAttribute = "ramdisk"
	INSTANCE_ROOT_DEVICE_NAME                     InstanceAttribute = "rootDeviceName"
	INSTANCE_SOURCE_DEST_CHECK                    InstanceAttribute = "sourceDestCheck"
	INSTANCE_SRIOV_NET_SUPPORT                    InstanceAttribute = "sriovNetSupport"
	INSTANCE_USER_DATA                            InstanceAttribute = "userData"
)

// Indicates whether this is a Spot Instance.
//	NONE = ""
//	SPOT = "spot"
type InstanceLifecycle string

const (
	NONE InstanceLifecycle = ""
	SPOT InstanceLifecycle = "spot"
)

// The associated code of the event.
//	INSTANCE_REBOOT = "instance-reboot"
//	SYSTEM_REBOOT = "system-reboot"
//	SYSTEM_MAINTENANCE = "system-maintenance"
//	INSTANCE_RETIREMENT = "instance-retirement"
//	INSTANCE_STOP = "instance-stop"
type InstanceStatus string

const (
	INSTANCE_REBOOT     InstanceStatus = "instance-reboot"
	SYSTEM_REBOOT       InstanceStatus = "system-reboot"
	SYSTEM_MAINTENANCE  InstanceStatus = "system-maintenance"
	INSTANCE_RETIREMENT InstanceStatus = "instance-retirement"
	INSTANCE_STOP       InstanceStatus = "instance-stop"
)

// The tenancy of the reserved instance.
//	DEDICATED InstanceTenancy = "dedicated"
//	DEFAULT InstanceTenancy = "default"
type InstanceTenancy string

const (
	DEDICATED InstanceTenancy = "dedicated"
	DEFAULT   InstanceTenancy = "default"
)

// The Instance request type.
//	ONE_TIME = "one-time"
//	PERSISTENT = "persistent"
type InstanceType string

const (
	ONE_TIME   InstanceType = "one-time"
	PERSISTENT InstanceType = "persistent"
)

// The Reserved Instance offering type.
//	 HEAVY_UTILIZATION = "Heavy Utilization"
//	 LIGHT_UTILIZATION = "Light Utilization"
//	 MEDIUM_UTILIZATION = "Medium Utilization"
type OfferingType string

const (
	HEAVY_UTILIZATION  OfferingType = "Heavy Utilization"
	LIGHT_UTILIZATION  OfferingType = "Light Utilization"
	MEDIUM_UTILIZATION OfferingType = "Medium Utilization"
)

// Describes how the route was created.
//	CREATE_ROUTE = "CreateRoute"
//	CREATE_ROUTE_TABLE = "CreateRouteTable"
//	ENABLE_VGW_ROUTE_PROPOGATION = "EnableVgwRoutePropagation"
type Origin string

const (
	CREATE_ROUTE                 Origin = "CreateRoute"
	CREATE_ROUTE_TABLE           Origin = "CreateRouteTable"
	ENABLE_VGW_ROUTE_PROPOGATION Origin = "EnableVgwRoutePropagation"
)

// The Reserved Instance description.
//	LINUX_UNIX = "Linux/UNIX"
//	LINUX_UNIX_AMAZON_VPC = "Linux/UNIX (Amazon VPC)"
//	WINDOWS = "Windows"
//	WINDOWS_AMAZON_VPC = "Windows (Amazon VPC)"
type ProductDescription string

const (
	LINUX_UNIX            ProductDescription = "Linux/UNIX"
	LINUX_UNIX_AMAZON_VPC ProductDescription = "Linux/UNIX (Amazon VPC)"
	SUSE_LINUX            ProductDescription = "SUSE Linux"
	SUSE_LINUX_AMAZON_VPC ProductDescription = "SUSE Linux (Amazon VPC)"
	WINDOWS               ProductDescription = "Windows"
	WINDOWS_AMAZON_VPC    ProductDescription = "Windows (Amazon VPC)"
)

// The type of product code.
//	DEVPAY = "devpay"
//	MARKETPLACE = "marketplace"
type ProductCodeType string

const (
	DEVPAY      ProductCodeType = "devpay"
	MARKETPLACE ProductCodeType = "marketplace"
)

// The reason code for the state change.
//	CLIENT_INTERNAL_ERROR = "Client.InternalError"
//	CLIENT_INSTANCE_INITIATED_SHUTDOWN = "Client.InstanceInitiatedShutdown"
//	CLIENT_INVALID_SNAPSHOT_NOT_FOUND = "Client.InvalidSnapshot.NotFound"
//	CLIENT_USER_INITIATED_SHUTDOWN = "Client.UserInitiatedShutdown"
//	CLIENT_VOLUME_LIMIT_EXCEEDED = "Client.VolumeLimitExceeded"
//	SERVER_INSUFFICIENT_INSTANCE_CAPACITY = "Server.InsufficientInstanceCapacity"
//	SERVER_INTERNAL_ERROR = "Server.InternalError"
//	SERVER_SPOT_INSTANCE_TERMINATION = "Server.SpotInstanceTermination"
type ReasonCode string

const (
	CLIENT_INTERNAL_ERROR                 ReasonCode = "Client.InternalError"
	CLIENT_INSTANCE_INITIATED_SHUTDOWN    ReasonCode = "Client.InstanceInitiatedShutdown"
	CLIENT_INVALID_SNAPSHOT_NOT_FOUND     ReasonCode = "Client.InvalidSnapshot.NotFound"
	CLIENT_USER_INITIATED_SHUTDOWN        ReasonCode = "Client.UserInitiatedShutdown"
	CLIENT_VOLUME_LIMIT_EXCEEDED          ReasonCode = "Client.VolumeLimitExceeded"
	SERVER_INSUFFICIENT_INSTANCE_CAPACITY ReasonCode = "Server.InsufficientInstanceCapacity"
	SERVER_INTERNAL_ERROR                 ReasonCode = "Server.InternalError"
	SERVER_SPOT_INSTANCE_TERMINATION      ReasonCode = "Server.SpotInstanceTermination"
)

// The type of resource.
//	CUSTOMER_GATEWAY = "customer-gateway"
//	DHCP_OPTIONS = "dhcp-options"
//	IMAGE = "image"
//	INSTANCE = "instance"
//	INTERNET_GATEWAY = "internet-gateway"
//	NETWORK_ACL = "network-acl"
//	NETWORL_INTERFACE = "network-interface"
//	RESERVED_INSTANCES = "reserved-instances"
//	ROUTE_TABLE = "route-table"
//	SECURITY_GROUP = "security-group"
//	SNAPSHOT = "snapshot"
//	SPOT_INSTANCES_REQUEST = "spot-instances-request"
//	SUBNET = "subnet"
//	VOLUME = "volume"
//	VIRTUAL_PRIVATE_CLOUD = "vpc"
//	VPN_CONNECTION = "vpn-connection"
//	VPN_GATEWAY = "vpn-gateway"
type ResourceType string

const (
	CUSTOMER_GATEWAY       ResourceType = "customer-gateway"
	DHCP_OPTIONS           ResourceType = "dhcp-options"
	IMAGE                  ResourceType = "image"
	INSTANCE               ResourceType = "instance"
	INTERNET_GATEWAY       ResourceType = "internet-gateway"
	NETWORK_ACL            ResourceType = "network-acl"
	NETWORL_INTERFACE      ResourceType = "network-interface"
	RESERVED_INSTANCES     ResourceType = "reserved-instances"
	ROUTE_TABLE            ResourceType = "route-table"
	SECURITY_GROUP         ResourceType = "security-group"
	SNAPSHOT               ResourceType = "snapshot"
	SPOT_INSTANCES_REQUEST ResourceType = "spot-instances-request"
	SUBNET                 ResourceType = "subnet"
	VOLUME                 ResourceType = "volume"
	VIRTUAL_PRIVATE_CLOUD  ResourceType = "vpc"
	VPN_CONNECTION         ResourceType = "vpn-connection"
	VPN_GATEWAY            ResourceType = "vpn-gateway"
)

// The type of root device used by the AMI.
//	EBS = "ebs"
//	INSTANCE_STORE = "instance-store"
type RootDeviceType string

const (
	EBS            RootDeviceType = "ebs"
	INSTANCE_STORE RootDeviceType = "instance-store"
)

// Type of rule action.
//	ALLOW = "allow"
//	DENY = "deny"
type RuleAction string

const (
	ALLOW RuleAction = "allow"
	DENY  RuleAction = "deny"
)

// The state of the item.
//	ACTIVE = "active"
//	AVAILABLE = "available"
//	BLACKHOLE = "blackhole"
//	BUNDLING = "bundling"
//	CANCELLED = "cancelled"
//	CANCELLING = "cancelling"
//	CLOSED = "closed"
//	COMPLETE = "complete"
//	COMPLETED = "completed"
//	DEGRADED = "degraded"
//	DELETED = "deleted"
//	DELETING = "deleting"
//	DISABLED = "disabled"
//	ENABLED = "enabled"
//	ERROR = "error"
//	EXPIRED = "expired"
//	FAILED = "failed"
//	IMPAIRED = "impaired"
//	INITIATING_REQUEST = "initiating-request"
//	INSUFFICIENT_DATA = "insufficient-data"
//	NORMAL = "normal"
//	NOT_APPLICABLE = "not-applicable"
//	OK = "ok"
//	OPEN = "open"
//	PASSED = "passed"
//	PAYMENT_FAILED = "payment-pending"
//	PAYMENT_PENDING = "payment-failed"
//	PENDING = "pending"
//	PENDING_ACCEPTANCE = "pending-acceptance"
//	PROVISIONING = "provisioning"
//	REJECTED = "rejected"
//	RETIRED = "retired"
//	RUNNING = "running"
//	SEVERELY_DEGRADED = "severely-degraded"
//	SHUTTING_DOWN = "shutting-down"
//	SOLD = "sold"
//	STALLED = "stalled"
//	STOPPED = "stopped"
//	STOPPING = "stopping"
//	STORING = "storing"
//	TERMINATED = "terminated"
//	UNAVAILABLE = "unavailable"
//	WAITING_FOR_SHUTDOWN = "waiting-for-shutdown"
type State string

const (
	ACTIVE               State = "active"
	AVAILABLE            State = "available"
	BLACKHOLE            State = "blackhole"
	BUNDLING             State = "bundling"
	CANCELLED            State = "cancelled"
	CANCELLING           State = "cancelling"
	CLOSED               State = "closed"
	COMPLETE             State = "complete"
	COMPLETED            State = "completed"
	DEGRADED             State = "degraded"
	DELETED              State = "deleted"
	DELETING             State = "deleting"
	DISABLED             State = "disabled"
	ENABLED              State = "enabled"
	ERROR                State = "error"
	EXPIRED              State = "expired"
	FAILED               State = "failed"
	IMPAIRED             State = "impaired"
	INITIATING_REQUEST   State = "initiating-request"
	INSUFFICIENT_DATA    State = "insufficient-data"
	NORMAL               State = "normal"
	NOT_APPLICABLE       State = "not-applicable"
	OK                   State = "ok"
	OPEN                 State = "open"
	PASSED               State = "passed"
	PAYMENT_FAILED       State = "payment-pending"
	PAYMENT_PENDING      State = "payment-failed"
	PENDING              State = "pending"
	PENDING_ACCEPTANCE   State = "pending-acceptance"
	PROVISIONING         State = "provisioning"
	REJECTED             State = "rejected"
	RETIRED              State = "retired"
	RUNNING              State = "running"
	SEVERELY_DEGRADED    State = "severely-degraded"
	SHUTTING_DOWN        State = "shutting-down"
	SOLD                 State = "sold"
	STALLED              State = "stalled"
	STOPPED              State = "stopped"
	STOPPING             State = "stopping"
	STORING              State = "storing"
	TERMINATED           State = "terminated"
	UNAVAILABLE          State = "unavailable"
	WAITING_FOR_SHUTDOWN State = "waiting-for-shutdown"
)

// The status of the item.
//	ATTACHED = "attached"
//	ATTACHING = "attaching"
//	AVAILABLE_NOW = "available"
//	DETACHED = "detached"
//	DETACHING = "detaching"
//	IN_USE = "in-use"
type Status string

const (
	ATTACHED      Status = "attached"
	ATTACHING     Status = "attaching"
	AVAILABLE_NOW Status = "available"
	DETACHED      Status = "detached"
	DETACHING     Status = "detaching"
	IN_USE        Status = "in-use"
)

// The placement strategy.
//	CLUSTER = "cluster"
type StrategyType string

const (
	CLUSTER StrategyType = "cluster"
)

// The type of support.
// 	SIMPLE = "simple"
type SupportType string

const (
	SIMPLE SupportType = "simple"
)

// The type of virtualization of the AMI.
//	CITRIX = "citrix"
//	HVM = "hvm"
//	PARAVIRTUAL = "paravirtual"
//	VMWARE = "vmware"
type VirtualizationType string

const (
	CITRIX      VirtualizationType = "citrix"
	HVM         VirtualizationType = "hvm"
	PARAVIRTUAL VirtualizationType = "paravirtual"
	MICROSOFT   VirtualizationType = "microsoft"
	VMWARE      VirtualizationType = "vmware"
)

// The name of the volume status.
//	IO_ENABLED = "io-enabled"
//	IO_PERFORMANCE = "io-performance"
type VolumeStatus string

const (
	IO_ENABLED     VolumeStatus = "io-enabled"
	IO_PERFORMANCE VolumeStatus = "io-performance"
)

// The volume type.
// 	GENERAL_PURPOSE = "gp2"
// 	MAGNETIC = "standard"
// 	PROVISIONED_IOPS = "io1"
type VolumeType string

const (
	GENERAL_PURPOSE  VolumeType = "gp2"
	MAGNETIC         VolumeType = "standard"
	PROVISIONED_IOPS VolumeType = "io1"
)

// The VPC attribute.
//	VPC_ENABLE_DNS_HOSTNAMES = "enableDnsHostnames"
//	VPC_ENABLE_DNS_SUPPORT = "enableDnsSupport"
type VpcAttribute string

const (
	VPC_ENABLE_DNS_HOSTNAMES VpcAttribute = "enableDnsHostnames"
	VPC_ENABLE_DNS_SUPPORT   VpcAttribute = "enableDnsSupport"
)

/******************************************************************************
 * Data Types
 */

// Describes an account attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AccountAttributeSetItemType.html]
type AccountAttributeSetItemType struct {
	AttributeName      string                             `xml:"attributeName"`
	AttributeValueList []AccountAttributeValueSetItemType `xml:"attributeValueList>item"`
}

// Describes a value of an account attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AccountAttributeValueSetItemType.html]
type AccountAttributeValueSetItemType struct {
	AttributeValue string `xml:"attributeValue"`
}

// Describes a private IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AssignPrivateIpAddressesSetItemRequestType.html]
type AssignPrivateIpAddressesSetItemRequestType struct {
	PrivateIpAddress string `xml:"privateIpAddress"`
}

// Describes an attachment between a volume and an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AttachmentSetItemResponseType.html]
type AttachmentSetItemResponseType struct {
	VolumeId            string    `xml:"volumeId"`
	InstanceId          string    `xml:"instanceId"`
	Device              string    `xml:"device"`
	Status              Status    `xml:"status"`
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
}

// Describes an attachment between a virtual private gateway and a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AttachmentType.html]
type AttachmentType struct {
	VpcId string `xml:"vpcId"`
	State Status `xml:"state"`
}

// Describes an Availability Zone.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AvailabilityZoneItemType.html]
type AvailabilityZoneItemType struct {
	ZoneName   string                        `xml:"zoneName"`
	ZoneState  State                         `xml:"zoneState"`
	RegionName string                        `xml:"regionName"`
	MessageSet []AvailabilityZoneMessageType `xml:"messageSet>item"`
}

// Describes a message about an Availability Zone.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-AvailabilityZoneMessageType.html]
type AvailabilityZoneMessageType struct {
	Message string `xml:"message"`
}

// Describes a block device mapping.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-BlockDeviceMappingItemType.html]
type BlockDeviceMappingItemType struct {
	DeviceName  string             `xml:"deviceName" name:"DeviceName"`
	VirtualName string             `xml:"virtualName" name:"VirtualName"`
	Ebs         EbsBlockDeviceType `xml:"ebs" name:"Ebs."`
	NoDevice    string             `xml:"noDevice" name:"NoDevice"`
}

// Describes the Amazon S3 bucket for an instance store-backed AMI.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-BundleInstanceS3StorageType.html]
type BundleInstanceS3StorageType struct {
	AwsAccessKeyId        string `xml:"awsAccessKeyId" name:"AwsAccessKeyId"`
	Bucket                string `xml:"bucket" name:"Bucket"`
	Prefix                string `xml:"prefix" name:"Prefix"`
	UploadPolicy          string `xml:"uploadPolicy" name:"UploadPolicy"`
	UploadPolicySignature string `xml:"uploadPolicySignature" name:"UploadPolicySignature"`
}

// Describes an error for BundleInstance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-BundleInstanceTaskErrorType.html]
type BundleInstanceTaskErrorType struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// Describes the storage location for an instance store-backed AMI.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-BundleInstanceTaskStorageType.html]
type BundleInstanceTaskStorageType struct {
	S3 BundleInstanceS3StorageType `xml:"S3"`
}

// Describes a bundle task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-BundleInstanceTaskType.html]
type BundleInstanceTaskType struct {
	InstanceId string                        `xml:"instanceId"`
	BundleId   string                        `xml:"bundleId"`
	State      State                         `xml:"state"`
	StartTime  time.Time                     `xml:"startTime"`
	UpdateTime time.Time                     `xml:"updateTime"`
	Storage    BundleInstanceTaskStorageType `xml:"storage"`
	Progress   string                        `xml:"progress"`
	Error      BundleInstanceTaskErrorType   `xml:"error"`
}

// Describes a request to cancel a Spot Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-CancelSpotInstanceRequestsResponseSetItemType.html]
type CancelSpotInstanceRequestsResponseSetItemType struct {
	SpotInstanceRequestId string `xml:"spotInstanceRequestId"`
	State                 State  `xml:"state"`
}

// Describes a conversion task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ConversionTaskType.html]
type ConversionTaskType struct {
	ConversionTaskId string                        `xml:"conversionTaskId"`
	ExpirationTime   time.Time                     `xml:"expirationTime"`
	ImportVolume     ImportVolumeTaskDetailsType   `xml:"importVolume"`
	ImportInstance   ImportInstanceTaskDetailsType `xml:"importInstance"`
	State            State                         `xml:"state"`
	StatusMessage    string                        `xml:"statusMessage"`
}

// Describes volume creation permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-CreateVolumePermissionItemType.html]
type CreateVolumePermissionItemType struct {
	UserId string `xml:"userId"`
	Group  string `xml:"group"`
}

// Describes a customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-CustomerGatewayType.html]
type CustomerGatewayType struct {
	CustomerGatewayId string                   `xml:"customerGatewayId"`
	State             State                    `xml:"state"`
	Type              string                   `xml:"type"`
	IpAddress         string                   `xml:"ipAddress"`
	BgpAsn            int                      `xml:"bgpAsn"`
	TagSet            []ResourceTagSetItemType `xml:"tagSet>item"`
}

// Describes an IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeAddressesResponseItemType.html]
type DescribeAddressesResponseItemType struct {
	PublicIp                string `xml:"PublicIp"`
	AllocationId            string `xml:"AllocationId"`
	Domain                  Domain `xml:"Domain"`
	InstanceId              string `xml:"InstanceId"`
	AssociationId           string `xml:"AssociationId"`
	NetworkInterfaceId      string `xml:"NetworkInterfaceId"`
	NetworkInterfaceOwnerId string `xml:"NetworkInterfaceOwnerId"`
}

// Describes an image.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeImagesResponseItemType.html]
type DescribeImagesResponseItemType struct {
	ImageId            string                     `xml:"imageId"`
	ImageLocation      string                     `xml:"imageLocation"`
	ImageState         State                      `xml:"imageState"`
	ImageOwnerId       string                     `xml:"imageOwnerId"`
	IsPublic           bool                       `xml:"isPublic"`
	ProductCodes       []ProductCodesSetItemType  `xml:"productCodes>item"`
	Architecture       Architecture               `xml:"architecture"`
	ImageType          ImageType                  `xml:"imageType"`
	KernalId           string                     `xml:"kernalId"`
	RamdiskId          string                     `xml:"ramdiskId"`
	Platform           string                     `xml:"platform"`
	SriovNetSupport    string                     `xml:"sriovNetSupport"`
	StateReason        StateReasonType            `xml:"stateReason"`
	ImageOwnerAlias    string                     `xml:"imageOwnerAlias"`
	Name               string                     `xml:"name"`
	Derscription       string                     `xml:"derscription"`
	RootDeviceType     RootDeviceType             `xml:"rootDeviceType"`
	RootDeviceName     string                     `xml:"rootDeviceName"`
	BlockDeviceMapping BlockDeviceMappingItemType `xml:"blockDeviceMapping"`
	VirtualizationType VirtualizationType         `xml:"virtualizationType"`
	TagSet             []ResourceTagSetItemType   `xml:"tagSet>item"`
	Hypervisor         Hypervisor                 `xml:"hypervisor"`
}

// Describes a key pair.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeKeyPairsResponseItemType.html]
type DescribeKeyPairsResponseItemType struct {
	KeyName        string `xml:"keyName"`
	KeyFingerprint string `xml:"keyFingerprint"`
}

// Describes a Reserved Instance listing.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeReservedInstancesListingsResponseSetItemType.html]
type DescribeReservedInstancesListingsResponseSetItemType struct {
	ReservedInstancesListingId string                   `xml:"reservedInstancesListingId"`
	ReservedInstancesId        string                   `xml:"reservedInstancesId"`
	CreateDate                 time.Time                `xml:"createDate"`
	UpdateDate                 time.Time                `xml:"updateDate"`
	Status                     State                    `xml:"status"`
	StatusMessage              string                   `xml:"statusMessage"`
	InstanceCounts             []InstanceCountsSetType  `xml:"instanceCounts>item"`
	PriceSchedules             []PriceScheduleSetType   `xml:"priceSchedules>item"`
	TagSet                     []ResourceTagSetItemType `xml:"tagSet>item"`
	ClientToken                string                   `xml:"clientToken"`
}

// Describes a Reserved Instance listing.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeReservedInstancesListingSetItemType.html]
type DescribeReservedInstancesListingSetItemType struct {
	ReservedInstancesListingId string `xml:"reservedInstancesListingId"`
}

// Describes a Reserved Instance modification.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ReservedInstancesModificationType.html]
type DescribeReservedInstancesModificationsResponseSetItemType struct {
	ReservedInstancesModificationId string                                      `xml:"reservedInstancesModificationId"`
	ClientToken                     string                                      `xml:"clientToken"`
	ReservedInstancesId             string                                      `xml:"reservedInstancesId"`
	ModificationResults             []ReservedInstancesConfigurationSetItemType `xml:"modificationResults>item"`
	CreateDate                      time.Time                                   `xml:"createDate"`
	UpdateDate                      time.Time                                   `xml:"updateDate"`
	EffectiveDate                   time.Time                                   `xml:"effectiveDate"`
	Status                          string                                      `xml:"status"`
	StatusMessage                   string                                      `xml:"statusMessage"`
}

// Describes a Reserved Instance offering.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeReservedInstancesOfferingsResponseSetItemType.html]
type DescribeReservedInstancesOfferingsResponseSetItemType struct {
	ReservedInstancesOfferingId string                        `xml:"reservedInstancesOfferingId"`
	InstanceType                string                        `xml:"instanceType"`
	AvailabilityZone            string                        `xml:"availabilityZone"`
	Duration                    int64                         `xml:"duration"`
	FixedPrice                  float64                       `xml:"fixedPrice"`
	UsagePrice                  float64                       `xml:"usagePrice"`
	ProductDescription          ProductDescription            `xml:"productDescription"`
	InstanceTenancy             InstanceTenancy               `xml:"instanceTenancy"`
	CurrencyCode                CurrencyCode                  `xml:"currencyCode"`
	OfferingType                OfferingType                  `xml:"offeringType"`
	RecurringCharges            []RecurringChargesSetItemType `xml:"recurringCharges>item"`
	Marketplace                 bool                          `xml:"marketplace"`
	PricingDetailsSet           []PricingDetailsSetItemType   `xml:"pricingDetailsSet>item"`
}

// Describes a Reserved Instance offering.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeReservedInstancesOfferingsResponseType1.html]
type DescribeReservedInstancesOfferingsResponseType struct {
	RequestId                     string                                                  `xml:"requestId"`
	ReservedInstancesOfferingsSet []DescribeReservedInstancesOfferingsResponseSetItemType `xml:"reservedInstancesOfferingsSet>item"`
	NextToken                     string                                                  `xml:"nextToken"`
}

// Describes a Reserved Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeReservedInstancesResponseSetItemType.html]
type DescribeReservedInstancesResponseSetItemType struct {
	ReservedInstancesId string                        `xml:"reservedInstancesId"`
	InstanceType        string                        `xml:"instanceType"`
	AvailabilityZone    string                        `xml:"availabilityZone"`
	Start               time.Time                     `xml:"start"`
	Duration            int64                         `xml:"duration"`
	End                 time.Time                     `xml:"end"`
	FixedPrice          float64                       `xml:"fixedPrice"`
	UsagePrice          float64                       `xml:"usagePrice"`
	InstanceCount       int                           `xml:"instanceCount"`
	ProductDescription  ProductDescription            `xml:"productDescription"`
	State               State                         `xml:"state"`
	TagSet              []ResourceTagSetItemType      `xml:"tagSet>item"`
	InstanceTenancy     InstanceTenancy               `xml:"instanceTenancy"`
	CurrencyCode        CurrencyCode                  `xml:"currencyCode"`
	OfferingType        OfferingType                  `xml:"offeringType"`
	RecurringCharges    []RecurringChargesSetItemType `xml:"recurringCharges>item"`
}

// Describes a snapshot.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeSnapshotsSetItemResponseType.html]
type DescribeSnapshotsSetItemResponseType struct {
	SnapshotId  string                   `xml:"snapshotId"`
	VolumeId    string                   `xml:"volumeId"`
	Status      State                    `xml:"status"`
	StartTime   time.Time                `xml:"startTime"`
	Progress    string                   `xml:"progress"`
	OwnerId     string                   `xml:"ownerId"`
	VolumeSize  string                   `xml:"volumeSize"`
	Description string                   `xml:"description"`
	OwnerAlias  string                   `xml:"ownerAlias"`
	TagSet      []ResourceTagSetItemType `xml:"tagSet>item"`
}

// Describes an Amazon EBS volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DescribeVolumesSetItemResponseType.html]
type DescribeVolumesSetItemResponseType struct {
	VolumeId         string                          `xml:"volumeId"`
	Size             string                          `xml:"size"`
	SnapshotId       string                          `xml:"snapshotId"`
	AvailabilityZone string                          `xml:"availabilityZone"`
	Status           State                           `xml:"status"`
	CreateTime       time.Time                       `xml:"createTime"`
	AttachmentSet    []AttachmentSetItemResponseType `xml:"attachmentSet>item"`
	TagSet           []ResourceTagSetItemType        `xml:"tagSet>item"`
	VolumeType       string                          `xml:"volumeType"`
	Iops             int                             `xml:"iops"`
	Encrypted        bool                            `xml:"encrypted"`
}

// Describes a DHCP configuration option.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DhcpConfigurationItemType.html]
type DhcpConfigurationItemType struct {
	Key      string   `xml:"key" name:"Key"`
	ValueSet []string `xml:"valueSet>item" name:"Value.#" base:"1"`
}

// Describes a set of DHCP options.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DhcpOptionsType.html]
type DhcpOptionsType struct {
	DhcpOptionsId        string                      `xml:"dhcpOptionsId"`
	DhcpConfigurationSet []DhcpConfigurationItemType `xml:"dhcpConfigurationSet>item"`
	TagSet               []ResourceTagSetItemType    `xml:"tagSet>item"`
}

// Describes the value of a DHCP option.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DhcpValueType.html]
type DhcpValueType struct {
	Value string `xml:"value"`
}

// Describes a disk image.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DiskImageDescriptionType.html]
type DiskImageDescriptionType struct {
	Format            string `xml:"format"`
	Size              int64  `xml:"size"`
	ImportManifestUrl string `xml:"importManifestUrl"`
	Checksum          string `xml:"checksum"`
}

// Describes the disk image for a volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-DiskImageVolumeDescriptionType.html]
type DiskImageVolumeDescriptionType struct {
	Size int    `xml:"size"`
	Id   string `xml:"id"`
}

// Describe an Amazon EBS block device.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-EbsBlockDeviceType.html]
type EbsBlockDeviceType struct {
	SnapshotId          string `xml:"snapshotId" name:"SnapshotId"`
	VolumeSize          int    `xml:"volumeSize" name:"VolumeSize"`
	DeleteOnTermination bool   `xml:"deleteOnTermination" name:"DeleteOnTermination"`
	VolumeType          string `xml:"volumeType" name:"VolumeType"`
	Iops                int    `xml:"iops,omitempty" name:"Iops"`
	Encrypted           bool   `xml:"encrypted" name:"Encrypted"`
}

// Describes a parameter used to set up an Amazon EBS volume in a block device mapping.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-EbsInstanceBlockDeviceMappingResponseType.html]
type EbsInstanceBlockDeviceMappingResponseType struct {
	VolumeId            string    `xml:"volumeId"`
	Status              Status    `xml:"status"`
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
}

// Describes an export task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ExportTaskResponseType.html]
type ExportTaskResponseType struct {
	ExportTaskId   string                         `xml:"exportTaskId"`
	Description    string                         `xml:"description"`
	State          State                          `xml:"state"`
	StatusMessage  string                         `xml:"statusMessage"`
	InstanceExport InstanceExportTaskResponseType `xml:"instanceExport"`
	ExportToS3     ExportToS3TaskResponseType     `xml:"exportToS3"`
}

// Describes an export task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ExportToS3TaskResponseType.html]
type ExportToS3TaskResponseType struct {
	DiskImageFormat DiskImageFormat `xml:"diskImageFormat" name:"DiskImageFormat,omitempty"`
	ContainerFormat ContainerFormat `xml:"containerFormat" name:"ContainerFormat,omitempty"`
	S3Bucket        string          `xml:"s3Bucket" name:"S3Bucket"`
	S3Key           string          `xml:"s3Key" name:"S3Key,omitempty"`
}

// Describes a filter configuration option.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAddresses.html]
type FilterItemType struct {
	Name     string   `xml:"name" name:"Name"`
	ValueSet []string `xml:"valueSet" name:"Value.#" base:"1"`
}

// Describes a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-GroupItemType.html]
type GroupItemType struct {
	GroupId   string `xml:"groupId" name:"GroupId,omitempty"`
	GroupName string `xml:"groupName" name:"GroupName,omitempty"`
}

// Describes an IAM instance profile.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-IamInstanceProfileRequestType.html]
type IamInstanceProfileRequestType struct {
	Arn  string `xml:"arn" name:"Arn,omitempty"`
	Name string `xml:"name" name:"Name,omitempty"`
}

// Describes an IAM instance profile.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-IamInstanceProfileResponseType.html]
type IamInstanceProfileResponseType struct {
	Arn string `xml:"arn"`
	Id  string `xml:"id"`
}

// Describes the ICMP type and code.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-IcmpTypeCodeType.html]
type IcmpTypeCodeType struct {
	Code int `xml:"code" name:"Code"`
	Type int `xml:"type" name:"Type"`
}

// Describes an import instance task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ImportInstanceTaskDetailsType.html]
type ImportInstanceTaskDetailsType struct {
	Volumes     []ImportInstanceVolumeDetailItemType `xml:"volumes>item"`
	InstanceId  string                               `xml:"instanceId"`
	Platform    string                               `xml:"platform"`
	Description string                               `xml:"description"`
}

// Describes an import instance volume task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ImportInstanceVolumeDetailItemType.html]
type ImportInstanceVolumeDetailItemType struct {
	BytesConverted   int64                          `xml:"bytesConverted"`
	AvailabilityZone string                         `xml:"availabilityZone"`
	Image            DiskImageDescriptionType       `xml:"image"`
	Description      string                         `xml:"description"`
	Volume           DiskImageVolumeDescriptionType `xml:"volume"`
	Status           string                         `xml:"status"`
	StatusMessage    string                         `xml:"statusMessage"`
}

// Describes an import volume task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ImportVolumeTaskDetailsType.html]
type ImportVolumeTaskDetailsType struct {
	BytesConverted   int64                          `xml:"bytesConverted"`
	AvailabilityZone string                         `xml:"availabilityZone"`
	Description      string                         `xml:"description"`
	Image            DiskImageDescriptionType       `xml:"image"`
	Volume           DiskImageVolumeDescriptionType `xml:"volume"`
}

// Describes a block device mapping.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceBlockDeviceMappingItemType.html]
type InstanceBlockDeviceMappingItemType struct {
	DeviceName  string                     `xml:"deviceName" name:"DeviceName"`
	VirtualName string                     `xml:"virtualName" name:"VirtualName"`
	Ebs         InstanceEbsBlockDeviceType `xml:"ebs" name:"Ebs."`
	NoDevice    string                     `xml:"noDevice" name:"NoDevice"`
}

// Describes a block device mapping.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceBlockDeviceMappingResponseItemType.html]
type InstanceBlockDeviceMappingResponseItemType struct {
	DeviceName string                                    `xml:"deviceName" name:"DeviceName"`
	Ebs        EbsInstanceBlockDeviceMappingResponseType `xml:"ebs" name:"Ebs."`
}

//
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceCountsSetItemType.html]
type InstanceCountsSetItemType struct {
	State         State `xml:"state"`
	InstanceCount int   `xml:"instanceCount"`
}

// Contains a set of Reserved Instance listing states.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceCountsSetType.html]
type InstanceCountsSetType struct {
	Item InstanceCountsSetItemType `xml:"item"`
}

// Describes parameters used to set up an Amazon EBS volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceEbsBlockDeviceType.html]
type InstanceEbsBlockDeviceType struct {
	DeleteOnTermination bool   `xml:"deleteOnTermination" name:"DeleteOnTermination"`
	VolumeId            string `xml:"volumeId" name:"VolumeId"`
}

// Describes an instance export task.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceExportTaskResponseType.html]
type InstanceExportTaskResponseType struct {
	InstanceId        string             `xml:"instanceId"`
	TargetEnvironment VirtualizationType `xml:"targetEnvironment"`
}

// Describes the monitoring information for an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceMonitoringStateType.html]
type InstanceMonitoringStateType struct {
	State State `xml:"state"`
}

// Describes association information for an Elastic IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceNetworkInterfaceAssociationType.html]
type InstanceNetworkInterfaceAssociationType struct {
	PublicIp      string `xml:"publicIp"`
	PublicDnsName string `xml:"publicDnsName"`
	IpOwnerId     string `xml:"ipOwnerId"`
}

// Describes a network interface attachment.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceNetworkInterfaceAttachmentType.html]
type InstanceNetworkInterfaceAttachmentType struct {
	AttachmentID        string    `xml:"attachmentID"`
	DeviceIndex         int       `xml:"deviceIndex"`
	Status              Status    `xml:"status"`
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
}

// Describes a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceNetworkInterfaceSetItemRequestType.html]
type InstanceNetworkInterfaceSetItemRequestType struct {
	NetworkInterfaceId             string                                 `xml:"networkInterfaceId" name:"NetworkInterfaceId,omitempty"`
	DeviceIndex                    int                                    `xml:"deviceIndex" name:"DeviceIndex,omitempty"`
	SubnetId                       string                                 `xml:"subnetId" name:"SubnetId,omitempty"`
	Description                    string                                 `xml:"description" name:"Description,omitempty"`
	PrivateIpAddress               string                                 `xml:"privateIpAddress" name:"PrivateIpAddress,omitempty"`
	GroupSet                       []SecurityGroupIdSetItemType           `xml:"groupSet>item" name:"GroupSet.,omitempty"`
	DeleteOnTermination            bool                                   `xml:"deleteOnTermination" name:"DeleteOnTermination,omitempty"`
	PrivateIpAddressesSet          []PrivateIpAddressesSetItemRequestType `xml:"privateIpAddressesSet>item" name:"PrivateIpAddressesSet.,omitempty"`
	SecondaryPrivateIpAddressCount int                                    `xml:"secondaryPrivateIpAddressCount" name:"SecondaryPrivateIpAddressCount,omitempty"`
}

// Describes a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceNetworkInterfaceSetItemType.html]
type InstanceNetworkInterfaceSetItemType struct {
	NetworkInterfaceId    string                                  `xml:"networkInterfaceId"`
	SubnetId              string                                  `xml:"subnetId"`
	VpcId                 string                                  `xml:"vpcId"`
	Description           string                                  `xml:"description"`
	OwnerId               string                                  `xml:"ownerId"`
	Status                State                                   `xml:"status"`
	MacAddress            string                                  `xml:"macAddress"`
	PrivateIpAddress      string                                  `xml:"privateIpAddress"`
	PrivateDnsName        string                                  `xml:"privateDnsName"`
	SourceDestCheck       bool                                    `xml:"sourceDestCheck"`
	GroupSet              []GroupItemType                         `xml:"groupSet>item"`
	Attachment            InstanceNetworkInterfaceAttachmentType  `xml:"attachment"`
	Association           InstanceNetworkInterfaceAssociationType `xml:"association"`
	PrivateIpAddressesSet []InstancePrivateIpAddressesSetItemType `xml:"privateIpAddressesSet>item"`
}

// Describes a private IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstancePrivateIpAddressesSetItemType.html]
type InstancePrivateIpAddressesSetItemType struct {
	PrivateIpAddress string                                  `xml:"privateIpAddress"`
	PrivateDnsName   string                                  `xml:"privateDnsName"`
	Primary          bool                                    `xml:"primary"`
	Association      InstanceNetworkInterfaceAssociationType `xml:"association"`
}

// Describes an instance state change.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStateChangeType.html]
type InstanceStateChangeType struct {
	InstanceId    string            `xml:"instanceId"`
	CurrentState  InstanceStateType `xml:"currentState"`
	PreviousState InstanceStateType `xml:"previousState"`
}

// Describes the current state of the instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStateType.html]
type InstanceStateType struct {
	Code uint16 `xml:"code"`
	Name State  `xml:"name"`
}

// Describes the instance status with the cause and more detail.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusDetailsSetType.html]
type InstanceStatusDetailsSetType struct {
	Name          string    `xml:"name"`
	Status        string    `xml:"status"`
	ImpairedSince time.Time `xml:"impairedSince"`
}

// Describes a set of instance events.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusEventSetType.html]
type InstanceStatusEventsSetType struct {
	Item InstanceStatusEventType `xml:"item"`
}

// Describes an instance event.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusEventType.html]
type InstanceStatusEventType struct {
	Code        InstanceStatus `xml:"code"`
	Description string         `xml:"description"`
	NotBefore   time.Time      `xml:"notBefore"`
	NotAfter    time.Time      `xml:"notAfter"`
}

// Describes the instance status, cause, details, and potential actions to take in response.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusItemType.html]
type InstanceStatusItemType struct {
	InstanceId       string                        `xml:"instanceId"`
	AvailabilityZone string                        `xml:"availabilityZone"`
	EventsSet        []InstanceStatusEventsSetType `xml:"eventsSet>item"`
	InstanceState    InstanceStateType             `xml:"instanceState"`
	SystemStatus     InstanceStatusType            `xml:"systemStatus"`
	InstanceStatus   InstanceStatusType            `xml:"instanceStatus"`
}

// Describes the status of an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusSetType.html]
type InstanceStatusSetType struct {
	Item InstanceStatusItemType `xml:"item"`
}

// Describes the status of an instance with details.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InstanceStatusType.html]
type InstanceStatusType struct {
	Status  State                          `xml:"status"`
	Details []InstanceStatusDetailsSetType `xml:"details>item"`
}

// Describes the attachment of a VPC to an Internet gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InternetGatewayAttachmentType.html]
type InternetGatewayAttachmentType struct {
	VpcId string `xml:"vpcId"`
	State Status `xml:"state"`
}

// Describes an Internet gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-InternetGatewayType.html]
type InternetGatewayType struct {
	InternetGatewayId string                          `xml:"internetGatewayId"`
	AttachmentSet     []InternetGatewayAttachmentType `xml:"attachmentSet>item"`
	TagSet            []ResourceTagSetItemType        `xml:"tagSet>item"`
}

// Describes a security group rule.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-IpPermissionType.html]
type IpPermissionType struct {
	IpProtocol string                `xml:"ipProtocol" name:"IpProtocol,omitempty"`
	FromPort   int                   `xml:"fromPort" name:"FromPort,omitempty"`
	ToPort     int                   `xml:"toPort" name:"ToPort,omitempty"`
	Groups     []UserIdGroupPairType `xml:"groups" name:"Groups.#.,omitempty"`
	IpRanges   []IpRangeItemType     `xml:"ipRanges" name:"IpRanges.#.,omitempty"`
}

// Describes an IP range.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-IpRangeItemType.html]
type IpRangeItemType struct {
	CidrIp string `xml:"cidrIp" name:"CidrIp"`
}

// Describes a launch permission.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-LaunchPermissionItemType.html]
type LaunchPermissionItemType struct {
	Group  string `xml:"group"`
	UserId string `xml:"userId"`
}

// Describes the launch specification of a Spot Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-LaunchSpecificationRequestType.html]
type LaunchSpecificationRequestType struct {
	ImageId             string                                       `xml:"imageId" name:"ImageId,omitempty"`
	KeyName             string                                       `xml:"keyName" name:"KeyName,omitempty"`
	GroupSet            []GroupItemType                              `xml:"groupSet>item" name:"GroupSet.,omitempty"`
	UserData            *UserDataType                                `xml:"userData" name:"UserData.,omitempty"`
	InstanceType        string                                       `xml:"instanceType" name:"InstanceType"`
	Placement           *PlacementRequestType                        `xml:"placement" name:"Placement.,omitempty"`
	KernelId            string                                       `xml:"kernelId" name:"KernelId,omitempty"`
	RamdiskId           string                                       `xml:"ramdiskId" name:"RamdiskId,omitempty"`
	BlockDeviceMapping  *BlockDeviceMappingItemType                  `xml:"blockDeviceMapping" name:"BlockDeviceMapping.,omitempty"`
	Monitoring          *MonitoringInstanceType                      `xml:"monitoring" name:"Monitoring.,omitempty"`
	SubnetId            string                                       `xml:"subnetId" name:"SubnetId,omitempty"`
	NetworkInterfaceSet []InstanceNetworkInterfaceSetItemRequestType `xml:"networkInterfaceSet>item" name:"NetworkInterfaceSet.,omitempty"`
	IamInstanceProfile  *IamInstanceProfileRequestType               `xml:"iamInstanceProfile" name:"IamInstanceProfile.,omitempty"`
	EbsOptimized        bool                                         `xml:"ebsOptimized" name:"EbsOptimized,omitempty"`
}

// Describes the launch specification of a Spot Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-LaunchSpecificationResponseType.html]
type LaunchSpecificationResponseType struct {
	ImageId             string                                       `xml:"imageId"`
	KeyName             string                                       `xml:"keyName"`
	GroupSet            []GroupItemType                              `xml:"groupSet>item"`
	UserData            UserDataType                                 `xml:"userData"`
	InstanceType        string                                       `xml:"instanceType"`
	Placement           PlacementRequestType                         `xml:"placement"`
	KernelId            string                                       `xml:"kernelId"`
	RamdiskId           string                                       `xml:"ramdiskId"`
	BlockDeviceMapping  BlockDeviceMappingItemType                   `xml:"blockDeviceMapping"`
	Monitoring          MonitoringInstanceType                       `xml:"monitoring"`
	SubnetId            string                                       `xml:"subnetId"`
	NetworkInterfaceSet []InstanceNetworkInterfaceSetItemRequestType `xml:"networkInterfaceSet>item"`
	IamInstanceProfile  IamInstanceProfileRequestType                `xml:"iamInstanceProfile"`
	EbsOptimized        bool                                         `xml:"ebsOptimized"`
}

// Describes the monitoring for the instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-MonitoringInstanceType.html]
type MonitoringInstanceType struct {
	Enabled bool `xml:"enabled" name:"Enabled"`
}

// Describes the monitoring for the instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-MonitorInstancesResponseSetItemType.html]
type MonitorInstancesResponseSetItemType struct {
	InstanceId string                      `xml:"instanceId"`
	Monitoring InstanceMonitoringStateType `xml:"monitoring"`
}

// Describes an association between a network ACL and a subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkAclAssociationType.html]
type NetworkAclAssociationType struct {
	NetworkAclAssociationId string `xml:"networkAclAssociationId"`
	NetworkAclId            string `xml:"networkAclId"`
	SubnetId                string `xml:"subnetId"`
}

// Describes an entry in a network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkAclEntryType.html]
type NetworkAclEntryType struct {
	RuleNumber   int              `xml:"ruleNumber"`
	Protocol     int              `xml:"protocol"`
	RuleAction   string           `xml:"ruleAction"`
	Egress       bool             `xml:"egress"`
	CidrBlock    string           `xml:"cidrBlock"`
	IcmpTypeCode IcmpTypeCodeType `xml:"icmpTypeCode"`
	PortRange    PortRangeType    `xml:"portRange"`
}

// Describes a network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkAclType.html]
type NetworkAclType struct {
	NetworkAclId   string                      `xml:"networkAclId"`
	VpcId          string                      `xml:"vpcId"`
	Default        bool                        `xml:"default"`
	EntrySet       []NetworkAclEntryType       `xml:"entrySet>item"`
	AssociationSet []NetworkAclAssociationType `xml:"associationSet>item"`
	TagSet         []ResourceTagSetItemType    `xml:"tagSet>item"`
}

// Describes association information for an Elastic IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkInterfaceAssociationType.html]
type NetworkInterfaceAssociationType struct {
	PublicIp      string `xml:"publicIp"`
	PublicDnsName string `xml:"publicDnsName"`
	IpOwnerId     string `xml:"ipOwnerId"`
	AllocationId  string `xml:"allocationId"`
	AssociationId string `xml:"associationId"`
}

// Describes a network interface attachment.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkInterfaceAttachmentType.html]
type NetworkInterfaceAttachmentType struct {
	AttachmentId        string    `xml:"attachmentId"`
	InstanceId          string    `xml:"instanceId"`
	InstanceOwnerId     string    `xml:"instanceOwnerId"`
	DeviceIndex         int       `xml:"deviceIndex"`
	Status              State     `xml:"status"`
	AttachTime          time.Time `xml:"attachTime"`
	DeleteOnTermination bool      `xml:"deleteOnTermination"`
}

// Describes the private IP address of a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkInterfacePrivateIpAddressesSetItemType.html]
type NetworkInterfacePrivateIpAddressesSetItemType struct {
	PrivateIpAddress string                          `xml:"privateIpAddress"`
	PrivateDnsName   string                          `xml:"privateDnsName"`
	Primary          bool                            `xml:"primary"`
	Association      NetworkInterfaceAssociationType `xml:"association"`
}

// Describes a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-NetworkInterfaceType.html]
type NetworkInterfaceType struct {
	NetworkInterfaceId    string                                          `xml:"networkInterfaceId"`
	SubnetId              string                                          `xml:"subnetId"`
	VpcId                 string                                          `xml:"vpcId"`
	AvailabilityZone      string                                          `xml:"availabilityZone"`
	Description           string                                          `xml:"description"`
	OwnerId               string                                          `xml:"ownerId"`
	RequesterId           string                                          `xml:"requesterId"`
	RequesterManaged      string                                          `xml:"requesterManaged"`
	Status                Status                                          `xml:"status"`
	MacAddress            string                                          `xml:"macAddress"`
	PrivateIpAddress      string                                          `xml:"privateIpAddress"`
	PrivateDnsName        string                                          `xml:"privateDnsName"`
	SourceDestCheck       bool                                            `xml:"sourceDestCheck"`
	GroupSet              []GroupItemType                                 `xml:"groupSet>item"`
	Attachment            NetworkInterfaceAttachmentType                  `xml:"attachment"`
	Association           NetworkInterfaceAssociationType                 `xml:"association"`
	TagSet                []ResourceTagSetItemType                        `xml:"tagSet>item"`
	PrivateIpAddressesSet []NetworkInterfacePrivateIpAddressesSetItemType `xml:"privateIpAddressesSet>item"`
}

// Describes a placement group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PlacementGroupInfoType.html]
type PlacementGroupInfoType struct {
	GroupName string       `xml:"groupName"`
	Strategy  StrategyType `xml:"strategy"`
	State     State        `xml:"state"`
}

// Describes a placement group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PlacementRequestType.html]
type PlacementRequestType struct {
	AvailabilityZone string `xml:"availabilityZone" name:"AvailabilityZone,omitempty"`
	GroupName        string `xml:"groupName" name:"GroupName,omitempty"`
}

// Describes a placement group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PlacementResponseType.html]
type PlacementResponseType struct {
	AvailabilityZone string          `xml:"availabilityZone"`
	GroupName        string          `xml:"groupName"`
	Tenancy          InstanceTenancy `xml:"tenancy"`
}

// Describes a range of ports.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PortRangeType.html]
type PortRangeType struct {
	From int `xml:"from" name:"From,omitempty"`
	To   int `xml:"to" name:"To,omitempty"`
}

// Describes the price for a Reserved Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PriceScheduleRequestSetItemType.html]
type PriceScheduleRequestSetItemType struct {
	Term         int64        `xml:term""`
	Price        float64      `xml:price""`
	CurrencyCode CurrencyCode `xml:currencyCode""`
}

// escribes the price for a Reserved Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PriceScheduleSetItemType.html]
type PriceScheduleSetItemType struct {
	Term         int64        `xml:"term" name:"Term"`
	Price        float64      `xml:"price" name:"Price"`
	CurrencyCode CurrencyCode `xml:"currencyCode" name:"CurrencyCode"`
	Active       bool         `xml:"active" name:"Active"`
}

// Describes the price for a Reserved Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PriceScheduleSetType.html]
type PriceScheduleSetType struct {
	Item PriceScheduleSetItemType `xml:"item"`
}

// Describes a Reserved Instance offering.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PricingDetailsSetItemType.html]
type PricingDetailsSetItemType struct {
	Price float64 `xml:"price"`
	Count int     `xml:"count"`
}

// Describes a secondary private IP address for a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PrivateIpAddressesSetItemRequestType.html]
type PrivateIpAddressesSetItemRequestType struct {
	PrivateIpAddressesSet []string `xml:"privateIpAddressesSet>item" name:"PrivateIpAddress"`
	Primary               bool     `xml:"primary" name:"Primary"`
}

// Describes a product code.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ProductCodeItemType.html]
type ProductCodeItemType struct {
	ProductCode string `xml:"productCode"`
}

// Describes a product code.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ProductCodesSetItemType.html]
type ProductCodesSetItemType struct {
	ProductCode string `xml:"productCode"`
	Type        string `xml:"type"`
}

// Specifies a basic product description.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ProductDescriptionSetItemType.html]
type ProductDescriptionSetItemType struct {
	ProductDescription ProductDescription `xml:"productDescription"`
}

// Describes a virtual private gateway propagating route.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-PropagatingVGWType.html]
type PropagatingVgwType struct {
	GatewayID string `xml:"gatewayID"`
}

// Describes a recurring charge.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RecurringChargesSetItemType.html]
type RecurringChargesSetItemType struct {
	Frequency Frequency `xml:"frequency"`
	Amount    float64   `xml:"amount"`
}

// Describes a region.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RegionItemType.html]
type RegionItemType struct {
	RegionName     string `xml:"regionName"`
	RegionEndpoint string `xml:"regionEndpoint"`
}

// Describes a reservation.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ReservationInfoType.html]
type ReservationInfoType struct {
	ReservationId string                     `xml:"reservationId"`
	OwnerId       string                     `xml:"ownerId"`
	GroupSet      []GroupItemType            `xml:"groupSet>item"`
	InstancesSet  []RunningInstancesItemType `xml:"instancesSet>item"`
	RequesterId   string                     `xml:"requesterId"`
}

// Describes the limit price of a Reserved Instance offering.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ReservedInstanceLimitPriceType.html]
type ReservedInstanceLimitPriceType struct {
	Amount       float64      `xml:"amount" name:"Amount"`
	CurrencyCode CurrencyCode `xml:"currencyCode" name:"CurrencyCode"`
}

// The configuration settings for the modified Reserved Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ReservedInstancesConfiguration.html]
type ReservedInstancesConfigurationSetItemType struct {
	AvailabilityZone string `xml:"availabilityZone" name:"AvailabilityZone"`
	Platform         string `xml:"platform" name:"Platform"`
	InstanceCount    int    `xml:"instanceCount" name:"InstanceCount"`
	InstanceType     string `xml:"instanceType" name:"InstanceType"`
}

// Describes a Reserved Instance modification.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ReservedInstancesModificationResultType.html]
type ReservedInstancesModificationResultSetItemType struct {
	ReservedInstancesId string                                    `xml:"reservedInstancesId"`
	TargetConfiguration ReservedInstancesConfigurationSetItemType `xml:"targetConfiguration"`
}

// Describes the tags assigned to an Amazon EC2 resource.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-ResourceTagSetItemType.html]
type ResourceTagSetItemType struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// Describes an association between a route table and a subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RouteTableAssociationType.html]
type RouteTableAssociationType struct {
	RouteTableAssociationId string `xml:"routeTableAssociationId"`
	RouteTableId            string `xml:"routeTableId"`
	SubnetId                string `xml:"subnetId"`
	Main                    bool   `xml:"main"`
}

// Describes a route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RouteTableType.html]
type RouteTableType struct {
	RouteTableId      string                      `xml:"routeTableId"`
	VpcId             string                      `xml:"vpcId"`
	RouteSet          []RouteType                 `xml:"routeSet>item"`
	AssociationSet    []RouteTableAssociationType `xml:"associationSet>item"`
	PropagatingVgwSet []PropagatingVgwType        `xml:"propagatingVgwSet>item"`
	TagSet            []ResourceTagSetItemType    `xml:"tagSet>item"`
}

// Describes a route in a route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RouteType.html]
type RouteType struct {
	DestinationCidrBlock   string `xml:"destinationCidrBlock"`
	GatewayId              string `xml:"gatewayId"`
	InstanceId             string `xml:"instanceId"`
	InstanceOwnerId        string `xml:"instanceOwnerId"`
	NetworkInterfaceId     string `xml:"networkInterfaceId"`
	State                  State  `xml:"state"`
	Origin                 string `xml:"origin"`
	VpcPeeringConnectionId string `xml:"vpcPeeringConnectionId"`
}

// Describes a running instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-RunningInstancesItemType.html]
type RunningInstancesItemType struct {
	InstanceId            string                                     `xml:"instanceId"`
	ImageId               string                                     `xml:"imageId"`
	InstanceState         InstanceStateType                          `xml:"instanceState"`
	PrivateDnsName        string                                     `xml:"privateDnsName"`
	DnsName               string                                     `xml:"dnsName"`
	Reason                string                                     `xml:"reason"`
	KeyName               string                                     `xml:"keyName"`
	AmiLaunchIndex        string                                     `xml:"amiLaunchIndex"`
	ProductCodes          []ProductCodesSetItemType                  `xml:"productCodes>item"`
	InstanceType          string                                     `xml:"instanceType"`
	LaunchTime            time.Time                                  `xml:"launchTime"`
	Placement             PlacementResponseType                      `xml:"placement"`
	KernelId              string                                     `xml:"kernelId"`
	RamdiskId             string                                     `xml:"ramdiskId"`
	Platform              string                                     `xml:"platform"`
	Onitoring             InstanceMonitoringStateType                `xml:"monitoring"`
	SubnetId              string                                     `xml:"subnetId"`
	VpcId                 string                                     `xml:"vpcId"`
	PrivateIpAddress      string                                     `xml:"privateIpAddress"`
	IpAddress             string                                     `xml:"ipAddress"`
	SourceDestCheck       bool                                       `xml:"sourceDestCheck"`
	GroupSet              []GroupItemType                            `xml:"groupSet>item"`
	StateReason           StateReasonType                            `xml:"stateReason"`
	Architecture          Architecture                               `xml:"architecture"`
	RootDeviceType        RootDeviceType                             `xml:"rootDeviceType"`
	RootDeviceName        string                                     `xml:"rootDeviceName"`
	BlockDeviceMapping    InstanceBlockDeviceMappingResponseItemType `xml:"blockDeviceMapping"`
	InstanceLifecycle     string                                     `xml:"instanceLifecycle"`
	SpotInstanceRequestId string                                     `xml:"spotInstanceRequestId"`
	VirtualizationType    VirtualizationType                         `xml:"virtualizationType"`
	ClientToken           string                                     `xml:"clientToken"`
	TagSet                ResourceTagSetItemType                     `xml:"tagSet"`
	Hypervisor            Hypervisor                                 `xml:"hypervisor"`
	NetworkInterfaceSet   []InstanceNetworkInterfaceSetItemType      `xml:"networkInterfaceSet>item"`
	IamInstanceProfile    IamInstanceProfileResponseType             `xml:"iamInstanceProfile"`
	EbsOptimized          bool                                       `xml:"ebsOptimized"`
	SriovNetSupport       string                                     `xml:"sriovNetSupport"`
}

// Describes a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SecurityGroupIdSetItemType.html]
type SecurityGroupIdSetItemType struct {
	GroupId string `xml:"groupId" name:"GroupId"`
}

// Describes a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SecurityGroupItemType.html]
type SecurityGroupItemType struct {
	OwnerId             string                   `xml:"ownerId"`
	GroupId             string                   `xml:"groupId"`
	GroupName           string                   `xml:"groupName"`
	GroupDescription    string                   `xml:"groupDescription"`
	VpcId               string                   `xml:"vpcId"`
	IpPermissions       []IpPermissionType       `xml:"ipPermissions>item"`
	IpPermissionsEgress IpPermissionType         `xml:"ipPermissionsEgress"`
	TagSet              []ResourceTagSetItemType `xml:"tagSet>item"`
}

// Describes the datafeed for a Spot Instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SpotDatafeedSubscriptionType.html]
type SpotDatafeedSubscriptionType struct {
	OwnerId string                     `xml:"OwnerId"`
	Bucket  string                     `xml:"Bucket"`
	Prefix  string                     `xml:"Prefix"`
	State   State                      `xml:"State"`
	Fault   SpotInstanceStateFaultType `xml:"Fault"`
}

// Describe a Spot Instance request.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SpotInstanceRequestSetItemType.html]
type SpotInstanceRequestSetItemType struct {
	SpotInstanceRequestId    string                        `xml:"spotInstanceRequestId"`
	SpotPrice                string                        `xml:"spotInstanceRequestId"`
	Type                     InstanceType                  `xml:"spotPrice"`
	State                    State                         `xml:"state"`
	Fault                    SpotInstanceStateFaultType    `xml:"fault"`
	Status                   SpotInstanceStatusMessageType `xml:"status"`
	ValidFrom                time.Time                     `xml:"validFrom"`
	ValidUntil               time.Time                     `xml:"validUntil"`
	LaunchGroup              string                        `xml:"launchGroup"`
	AvailabilityZoneGroup    string                        `xml:"availabilityZoneGroup"`
	LaunchedAvailabilityZone string                        `xml:"launchedAvailabilityZone"`
	InstanceId               string                        `xml:"instanceId"`
	CreateTime               time.Time                     `xml:"createTime"`
	ProductDescription       string                        `xml:"productDescription"`
	TagSet                   []ResourceTagSetItemType      `xml:"tagSet>item"`
}

// Describes a Spot Instance state change.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SpotInstanceStateFaultType.html]
type SpotInstanceStateFaultType struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

// Describes a Spot Instance request.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SpotInstanceStatusMessageType.html]
type SpotInstanceStatusMessageType struct {
	Code       string    `xml:"code"`
	Message    string    `xml:"message"`
	UpdateTime time.Time `xml:"updateTime"`
}

// Describes the Spot Price history.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SpotPriceHistorySetItemType.html]
type SpotPriceHistorySetItemType struct {
	InstanceType       string             `xml:"instanceType"`
	ProductDescription ProductDescription `xml:"productDescription"`
	SpotPrice          string             `xml:"spotPrice"`
	Timestamp          time.Time          `xml:"timestamp"`
	AvailabilityZone   string             `xml:"availabilityZone"`
}

// Describes a state change.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-StateReasonType.html]
type StateReasonType struct {
	Code    ReasonCode `xml:"code"`
	Message string     `xml:"message"`
}

// Describes a subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-SubnetType.html]
type SubnetType struct {
	SubnetId                string                   `xml:"subnetId"`
	State                   State                    `xml:"state"`
	VpcId                   string                   `xml:"vpcId"`
	CidrBlock               string                   `xml:"cidrBlock"`
	AvailableIpAddressCount int                      `xml:"availableIpAddressCount"`
	AvailabilityZone        string                   `xml:"availabilityZone"`
	DefaultForAz            bool                     `xml:"defaultForAz"`
	MapPublicIpOnLaunch     bool                     `xml:"mapPublicIpOnLaunch"`
	TagSet                  []ResourceTagSetItemType `xml:"tagSet>item"`
}

// Describes a tag.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-TagSetItemType.html]
type TagSetItemType struct {
	ResourceId   string       `xml:"resourceId" name:"ResourceId,omitempty"`
	ResourceType ResourceType `xml:"resourceType" name:"ResourceType,omitempty"`
	Key          string       `xml:"key" name:"Key"`
	Value        string       `xml:"value" name:"Value"`
}

// Specifies user data.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-UserDataType.html]
type UserDataType struct {
	Date string `xml:"data" name:"Data"`
}

// Describes a security group and AWS account ID pair.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-UserIdGroupPairType.html]
type UserIdGroupPairType struct {
	UserId    string `xml:"userId"`
	GroupId   string `xml:"groupId"`
	GroupName string `xml:"groupName"`
}

// Describes the volume status, cause, details, and potential actions to take in response.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VolumeStatusItemType.html]
type VolumeStatusItemType struct {
	VolumeId         string                       `xml:"volumeId"`
	AvailabilityZone string                       `xml:"availabilityZone"`
	VolumeStatus     VolumeStatusInfoType         `xml:"volumeStatus"`
	EventSet         []VolumeStatusEventItemType  `xml:"eventSet>item"`
	ActionSet        []VolumeStatusActionItemType `xml:"actionSet>item"`
}

// Describes the volume status with details.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VolumeStatusInfoType.html]
type VolumeStatusInfoType struct {
	Status  State                       `xml:"status"`
	Details VolumeStatusDetailsItemType `xml:"details"`
}

// Describes the cause and more detail for a volume status.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VolumeStatusDetailsItemType.html]
type VolumeStatusDetailsItemType struct {
	Name   VolumeStatus `xml:"name"`
	Status State        `xml:"status"`
}

// Describes a volume status event.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VolumeStatusEventItemType.html]
type VolumeStatusEventItemType struct {
	EventType   string    `xml:"eventType"`
	EventId     string    `xml:"eventId"`
	Description string    `xml:"description"`
	NotBefore   time.Time `xml:"notBefore"`
	NotAfter    time.Time `xml:"notAfter"`
}

// Describes a volume status action code.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VolumeStatusActionItemType.html]
type VolumeStatusActionItemType struct {
	Code        string `xml:"code"`
	EventType   string `xml:"eventType"`
	EventId     string `xml:"eventId"`
	Description string `xml:"description"`
}

// Describes a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpcType.html]
type VpcType struct {
	vpcId           string                   `xml:"vpcId"`
	State           State                    `xml:"state"`
	CidrBlock       string                   `xml:"cidrBlock"`
	DhcpOptionsId   string                   `xml:"dhcpOptionsId"`
	TagSet          []ResourceTagSetItemType `xml:"tagSet>item"`
	InstanceTenancy InstanceTenancy          `xml:"instanceTenancy"`
	IsDefault       bool                     `xml:"IsDefault"`
}

// Describes a VPC peering connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpcPeeringConnectionType.html]
type VpcPeeringConnectionType struct {
	VpcPeeringConnectionId string                              `xml:"vpcPeeringConnectionId"`
	RequesterVpcInfo       VpcPeeringConnectionVpcInfoType     `xml:"requesterVpcInfo"`
	AccepterVpcInfo        VpcPeeringConnectionVpcInfoType     `xml:"accepterVpcInfo"`
	Status                 VpcPeeringConnectionStateReasonType `xml:"status"`
	ExpirationTime         time.Time                           `xml:"expirationTime"`
	TagSet                 []ResourceTagSetItemType            `xml:"tagSet>item"`
}

// Describes the status of a VPC peering connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpcPeeringConnectionStateReasonType.html]
type VpcPeeringConnectionStateReasonType struct {
	Code    State  `xml:"code"`
	Message string `xml:"message"`
}

// Describes a VPC in a VPC peering connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpcPeeringConnectionVpcInfoType.html]
type VpcPeeringConnectionVpcInfoType struct {
	VpcId     string `xml:"vpcId"`
	OwnerId   string `xml:"ownerId"`
	CidrBlock string `xml:"cidrBlock"`
}

// Describes VPN connection options.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpnConnectionOptionsResponseType.html]
type VpnConnectionOptionsResponseType struct {
	StaticRoutesOnly bool `xml:"staticRoutesOnly" name:"StaticRoutesOnly"`
}

// Describes a VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpnConnectionType.html]
type VpnConnectionType struct {
	VpnConnectionId              string                             `xml:"vpnConnectionId"`
	State                        State                              `xml:"state"`
	CustomerGatewayConfiguration string                             `xml:"customerGatewayConfiguration"`
	Type                         string                             `xml:"type"`
	CustomerGatewayId            string                             `xml:"customerGatewayId"`
	VpnGatewayId                 string                             `xml:"vpnGatewayId"`
	TagSet                       []ResourceTagSetItemType           `xml:"tagSet>item"`
	VgwTelemetry                 VpnTunnelTelemetryType             `xml:"vgwTelemetry"`
	Options                      []VpnConnectionOptionsResponseType `xml:"options>item"`
	Routes                       []VpnStaticRouteType               `xml:"routes>item"`
}

// Describes a virtual private gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpnGatewayType.html]
type VpnGatewayType struct {
	VpnGatewayId     string                   `xml:"vpnGatewayId"`
	State            State                    `xml:"state"`
	Type             string                   `xml:"type"`
	AvailabilityZone string                   `xml:"availabilityZone"`
	Attachments      []AttachmentType         `xml:"attachments>item"`
	TagSet           []ResourceTagSetItemType `xml:"tagSet>item"`
}

// Describes a static route for a VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpnStaticRouteType.html]
type VpnStaticRouteType struct {
	DestinationCidrBlock string `xml:"destinationCidrBlock"`
	Source               string `xml:"source"`
	State                State  `xml:"state"`
}

// Describes telemetry for a VPN tunnel.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-ItemType-VpnTunnelTelemetryType.html]
type VpnTunnelTelemetryType struct {
	OutsideIpAddress   string    `xml:"outsideIpAddress"`
	Status             string    `xml:"status"`
	LastStatusChange   time.Time `xml:"lastStatusChange"`
	StatusMessage      string    `xml:"statusMessage"`
	AcceptedRouteCount int       `xml:"acceptedRouteCount"`
}

/******************************************************************************
 * Responses
 */

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AcceptVpcPeeringConnection.html]
type AcceptVpcPeeringConnectionResponse struct {
	RequestId            string                   `xml:"requestId"`
	VpcPeeringConnection VpcPeeringConnectionType `xml:"vpcPeeringConnection"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AllocateAddress.html]
type AllocateAddressResponse struct {
	RequestId    string `xml:"requestId"`
	PublicIp     string `xml:"publicIp"`
	Domain       string `xml:"domain"`
	AllocationId string `xml:"allocationId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssignPrivateIpAddresses.html]
type AssignPrivateIpAddressesResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateAddress.html]
type AssociateAddressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateDhcpOptions.html]
type AssociateDhcpOptionsResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateRouteTable.html]
type AssociateRouteTableResponse struct {
	RequestId     string `xml:"requestId"`
	AssociationId string `xml:"associationId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachInternetGateway.html]
type AttachInternetGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachNetworkInterface.html]
type AttachNetworkInterfaceResponse struct {
	RequestId    string `xml:"requestId"`
	AttachmentId string `xml:"attachmentId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVolume.html]
type AttachVolumeResponse struct {
	RequestId  string    `xml:"requestId"`
	VolumeId   string    `xml:"volumeId"`
	InstanceId string    `xml:"instanceId"`
	Device     string    `xml:"device"`
	Status     Status    `xml:"status"`
	AttachTime time.Time `xml:"attachTime"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVpnGateway.html]
type AttachVpnGatewayResponse struct {
	RequestId  string         `xml:"requestId"`
	Attachment AttachmentType `xml:"attachment"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupEgress.html]
type AuthorizeSecurityGroupEgressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html]
type AuthorizeSecurityGroupIngressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-BundleInstance.html]
type BundleInstanceResponse struct {
	RequestId          string                 `xml:"requestId"`
	BundleInstanceTask BundleInstanceTaskType `xml:"bundleInstanceTask"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelBundleTask.html]
type CancelBundleTaskResponse struct {
	RequestId          string                 `xml:"requestId"`
	BundleInstanceTask BundleInstanceTaskType `xml:"bundleInstanceTask"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelConversionTask.html]
type CancelConversionTaskResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelExportTask.html]
type CancelExportTaskResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelReservedInstancesListing.html]
type CancelReservedInstancesListingResponse struct {
	RequestId                    string                                                 `xml:"requestId"`
	ReservedInstancesListingsSet []DescribeReservedInstancesListingsResponseSetItemType `xml:"reservedInstancesListingsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelSpotInstanceRequests.html]
type CancelSpotInstanceRequestsResponse struct {
	RequestId              string                           `xml:"requestId"`
	SpotInstanceRequestSet []SpotInstanceRequestSetItemType `xml:"spotInstanceRequestSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ConfirmProductInstance.html]
type ConfirmProductInstanceResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
	OwnerId   string `xml:"ownerId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopyImage.html]
type CopyImageResponse struct {
	RequestId string `xml:"requestId"`
	ImageId   string `xml:"imageId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopySnapshot.html]
type CopySnapshotResponse struct {
	RequestId  string `xml:"requestId"`
	SnapshotId string `xml:"snapshotId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateCustomerGateway.html]
type CreateCustomerGatewayResponse struct {
	RequestId       string              `xml:"requestId"`
	CustomerGateway CustomerGatewayType `xml:"customerGateway"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateDhcpOptions.html]
type CreateDhcpOptionsResponse struct {
	RequestId   string          `xml:"requestId"`
	DhcpOptions DhcpOptionsType `xml:"dhcpOptions"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateImage.html]
type CreateImageResponse struct {
	RequestId string `xml:"requestId"`
	ImageId   string `xml:"imageId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInstanceExportTask.html]
type CreateInstanceExportTaskResponse struct {
	RequestId  string                 `xml:"requestId"`
	ExportTask ExportTaskResponseType `xml:"exportTask"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInternetGateway.html]
type CreateInternetGatewayResponse struct {
	RequestId       string              `xml:"requestId"`
	InternetGateway InternetGatewayType `xml:"internetGateway"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateKeyPair.html]
type CreateKeyPairResponse struct {
	RequestId      string `xml:"requestId"`
	KeyName        string `xml:"keyName"`
	KeyFingerprint string `xml:"keyFingerprint"`
	KeyMaterial    string `xml:"keyMaterial"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAcl.html]
type CreateNetworkAclResponse struct {
	RequestId  string         `xml:"requestId"`
	NetworkAcl NetworkAclType `xml:"networkAcl"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAclEntry.html]
type CreateNetworkAclEntryResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkInterface.html]
type CreateNetworkInterfaceResponse struct {
	RequestId        string               `xml:"requestId"`
	NetworkInterface NetworkInterfaceType `xml:"networkInterface"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreatePlacementGroup.html]
type CreatePlacementGroupResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateReservedInstancesListing.html]
type CreateReservedInstancesListingResponse struct {
	RequestId            string                                         `xml:"requestId"`
	ReservedInstancesSet []DescribeReservedInstancesResponseSetItemType `xml:"reservedInstancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRoute.html]
type CreateRouteResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRouteTable.html]
type CreateRouteTableResponse struct {
	RequestId  string         `xml:"requestId"`
	RouteTable RouteTableType `xml:"routeTable"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSecurityGroup.html]
type CreateSecurityGroupResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
	GroupId   string `xml:"groupId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSnapshot.html]
type CreateSnapshotResponse struct {
	RequestId   string                 `xml:"requestId"`
	SnapshotId  string                 `xml:"snapshotId"`
	VolumeId    string                 `xml:"volumeId"`
	Status      State                  `xml:"status"`
	StartTime   time.Time              `xml:"startTime"`
	Progress    string                 `xml:"progress"`
	OwnerId     string                 `xml:"ownerId"`
	VolumeSize  string                 `xml:"volumeSize"`
	Description string                 `xml:"description"`
	OwnerAlias  string                 `xml:"ownerAlias"`
	TagSet      ResourceTagSetItemType `xml:"tagSet"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSpotDatafeedSubscription.html]
type CreateSpotDatafeedSubscriptionResponse struct {
	RequestId                string                       `xml:"requestId"`
	SpotDatafeedSubscription SpotDatafeedSubscriptionType `xml:"spotDatafeedSubscription"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSubnet.html]
type CreateSubnetResponse struct {
	RequestId string     `xml:"requestId"`
	Subnet    SubnetType `xml:"subnet"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateTags.html]
type CreateTagsResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html]
type CreateVolumeResponse struct {
	RequestId        string                        `xml:"requestId"`
	VolumeId         string                        `xml:"volumeId"`
	Size             string                        `xml:"size"`
	SnapshotId       string                        `xml:"snapshotId"`
	AvailabilityZone string                        `xml:"availabilityZone"`
	Status           State                         `xml:"status"`
	CreateTime       time.Time                     `xml:"createTime"`
	AttachmentSet    AttachmentSetItemResponseType `xml:"attachmentSet"`
	TagSet           ResourceTagSetItemType        `xml:"tagSet"`
	VolumeType       string                        `xml:"volumeType"`
	Iops             int                           `xml:"iops"`
	Encrypted        bool                          `xml:"encrypted"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpc.html]
type CreateVpcResponse struct {
	RequestId string  `xml:"requestId"`
	Vpc       VpcType `xml:"vpc"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpcPeeringConnection.html]
type CreateVpcPeeringConnectionResponse struct {
	RequestId            string                   `xml:"requestId"`
	VpcPeeringConnection VpcPeeringConnectionType `xml:"vpcPeeringConnection"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnection.html]
type CreateVpnConnectionResponse struct {
	RequestId     string            `xml:"requestId"`
	VpnConnection VpnConnectionType `xml:"vpnConnection"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnectionRoute.html]
type CreateVpnConnectionRouteResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnGateway.html]
type CreateVpnGatewayResponse struct {
	RequestId  string         `xml:"requestId"`
	VpnGateway VpnGatewayType `xml:"vpnGateway"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteCustomerGateway.html]
type DeleteCustomerGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteDhcpOptions.html]
type DeleteDhcpOptionsResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteInternetGateway.html]
type DeleteInternetGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteKeyPair.html]
type DeleteKeyPairResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAcl.html]
type DeleteNetworkAclResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAclEntry.html]
type DeleteNetworkAclEntryResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkInterface.html]
type DeleteNetworkInterfaceResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeletePlacementGroup.html]
type DeletePlacementGroupResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRoute.html]
type DeleteRouteResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRouteTable.html]
type DeleteRouteTableResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSecurityGroup.html]
type DeleteSecurityGroupResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSnapshot.html]
type DeleteSnapshotResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// Deletes the datafeed for Spot Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSpotDatafeedSubscription.html]
type DeleteSpotDatafeedSubscriptionResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSubnet.html]
type DeleteSubnetResponse struct {
	RequestId string `xml:"requestId"`
}

// Deletes the specified set of tags from the specified set of resources. This call is designed to follow a DescribeTags call.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteTags.html]
type DeleteTagsResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVolume.html]
type DeleteVolumeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpc.html]
type DeleteVpcResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpcPeeringConnection.html]
type DeleteVpcPeeringConnectionResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnection.html]
type DeleteVpnConnectionResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnectionRoute.html]
type DeleteVpnConnectionRouteResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnGateway.html]
type DeleteVpnGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeregisterImage.html]
type DeregisterImageResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAccountAttributes.html]
type DescribeAccountAttributesResponse struct {
	RequestId           string                        `xml:"requestId"`
	AccountAttributeSet []AccountAttributeSetItemType `xml:"accountAttributeSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAddresses.html]
type DescribeAddressesResponse struct {
	RequestId    string                              `xml:"requestId"`
	AddressesSet []DescribeAddressesResponseItemType `xml:"addressesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAvailabilityZones.html]
type DescribeAvailabilityZonesResponse struct {
	RequestId            string                     `xml:"requestId"`
	AvailabilityZoneInfo []AvailabilityZoneItemType `xml:"availabilityZoneInfo>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeBundleTasks.html]
type DescribeBundleTasksResponse struct {
	RequestId              string                   `xml:"requestId"`
	BundleInstanceTasksSet []BundleInstanceTaskType `xml:"bundleInstanceTasksSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeConversionTasks.html]
type DescribeConversionTasksResponse struct {
	RequestId       string               `xml:"requestId"`
	conversionTasks []ConversionTaskType `xml:"conversionTasks>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeCustomerGateways.html]
type DescribeCustomerGatewaysResponse struct {
	RequestId          string                `xml:"requestId"`
	CustomerGatewaySet []CustomerGatewayType `xml:"customerGatewaySet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeDhcpOptions.html]
type DescribeDhcpOptionsResponse struct {
	RequestId      string            `xml:"requestId"`
	DhcpOptionsSet []DhcpOptionsType `xml:"dhcpOptionsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeExportTasks.html]
type DescribeExportTasksResponse struct {
	RequestId     string                   `xml:"requestId"`
	ExportTaskSet []ExportTaskResponseType `xml:"exportTaskSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImageAttribute.html]
type DescribeImageAttributeResponse struct {
	RequestId        string                     `xml:"requestId"`
	LaunchPermission []LaunchPermissionItemType `xml:"launchPermission>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImages.html]
type DescribeImagesResponse struct {
	RequestId string                           `xml:"requestId"`
	ImagesSet []DescribeImagesResponseItemType `xml:"imagesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceAttribute.html]
type DescribeInstanceAttributeResponse struct {
	RequestId                         string `xml:"requestId"`
	InstanceId                        string `xml:"instanceId"`
	InstanceInitiatedShutdownBehavior string `xml:"instanceInitiatedShutdownBehavior>value"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstances.html]
type DescribeInstancesResponse struct {
	RequestId      string                `xml:"requestId"`
	ReservationSet []ReservationInfoType `xml:"reservationSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceStatus.html]
type DescribeInstanceStatusResponse struct {
	RequestId         string                   `xml:"requestId"`
	InstanceStatusSet []InstanceStatusItemType `xml:"instanceStatusSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInternetGateways.html]
type DescribeInternetGatewaysResponse struct {
	RequestId          string                `xml:"requestId"`
	InternetGatewaySet []InternetGatewayType `xml:"internetGatewaySet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeKeyPairs.html]
type DescribeKeyPairsResponse struct {
	RequestId string                             `xml:"requestId"`
	KeySet    []DescribeKeyPairsResponseItemType `xml:"keySet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkAcls.html]
type DescribeNetworkAclsResponse struct {
	RequestId     string           `xml:"requestId"`
	NetworkAclSet []NetworkAclType `xml:"networkAclSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaceAttribute.html]
type DescribeNetworkInterfaceAttributeResponse struct {
	RequestId          string `xml:"requestId"`
	NetworkInterfaceId string `xml:"networkInterfaceId"`
	SourceDestCheck    string `xml:"sourceDestCheck>value"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaces.html]
type DescribeNetworkInterfacesResponse struct {
	RequestId           string                 `xml:"requestId"`
	NetworkInterfaceSet []NetworkInterfaceType `xml:"networkInterfaceSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribePlacementGroups.html]
type DescribePlacementGroupsResponse struct {
	RequestId         string                   `xml:"requestId"`
	PlacementGroupSet []PlacementGroupInfoType `xml:"placementGroupSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRegions.html]
type DescribeRegionsResponse struct {
	RequestId  string           `xml:"requestId"`
	RegionInfo []RegionItemType `xml:"regionInfo>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstances.html]
type DescribeReservedInstancesResponse struct {
	RequestId            string                                         `xml:"requestId"`
	ReservedInstancesSet []DescribeReservedInstancesResponseSetItemType `xml:"reservedInstancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesListings.html]
type DescribeReservedInstancesListingsResponse struct {
	RequestId                    string                                        `xml:"requestId"`
	ReservedInstancesListingsSet []DescribeReservedInstancesListingSetItemType `xml:"reservedInstancesListingsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesModifications.html]
type DescribeReservedInstancesModificationsResponse struct {
	RequestId                         string                                                      `xml:"requestId"`
	ReservedInstancesModificationsSet []DescribeReservedInstancesModificationsResponseSetItemType `xml:"reservedInstancesModificationsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesOfferings.html]
type DescribeReservedInstancesOfferingsResponse struct {
	RequestId                     string                                                  `xml:"requestId"`
	ReservedInstancesOfferingsSet []DescribeReservedInstancesOfferingsResponseSetItemType `xml:"reservedInstancesOfferingsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRouteTables.html]
type DescribeRouteTablesResponse struct {
	RequestId     string           `xml:"requestId"`
	RouteTableSet []RouteTableType `xml:"routeTableSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSecurityGroups.html]
type DescribeSecurityGroupsResponse struct {
	RequestId         string                `xml:"requestId"`
	securityGroupInfo SecurityGroupItemType `xml:"securityGroupInfo>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshotAttribute.html]
type DescribeSnapshotAttributeResponse struct {
	RequestId              string                         `xml:"requestId"`
	SnapshotId             string                         `xml:"snapshotId"`
	CreateVolumePermission CreateVolumePermissionItemType `xml:"createVolumePermission>item"`
	ProductCodes           []ProductCodesSetItemType      `xml:"productCodes>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html]
type DescribeSnapshotsResponse struct {
	RequestId   string                                 `xml:"requestId"`
	SnapshotSet []DescribeSnapshotsSetItemResponseType `xml:"snapshotSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotDatafeedSubscription.html]
type DescribeSpotDatafeedSubscriptionResponse struct {
	RequestId                string                         `xml:"requestId"`
	SpotDatafeedSubscription []SpotDatafeedSubscriptionType `xml:"spotDatafeedSubscription"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotInstanceRequests.html]
type DescribeSpotInstanceRequestsResponse struct {
	RequestId      string                `xml:"requestId"`
	ReservationSet []ReservationInfoType `xml:"reservationSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotPriceHistory.html]
type DescribeSpotPriceHistoryResponse struct {
	RequestId           string                        `xml:"requestId"`
	SpotPriceHistorySet []SpotPriceHistorySetItemType `xml:"spotPriceHistorySet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSubnets.html]
type DescribeSubnetsResponse struct {
	RequestId string       `xml:"requestId"`
	SubnetSet []SubnetType `xml:"subnetSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeTags.html]
type DescribeTagsResponse struct {
	RequestId string           `xml:"requestId"`
	TagSet    []TagSetItemType `xml:"tagSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeAttribute.html]
type DescribeVolumeAttributeResponse struct {
	RequestId    string                `xml:"requestId"`
	VolumeId     string                `xml:"volumeId"`
	AutoEnableIO string                `xml:"autoEnableIO>value"`
	ProductCodes []ProductCodeItemType `xml:"productCodes>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumes.html]
type DescribeVolumesResponse struct {
	RequestId string                               `xml:"requestId"`
	VolumeSet []DescribeVolumesSetItemResponseType `xml:"volumeSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeStatus.html]
type DescribeVolumeStatusResponse struct {
	RequestId       string                 `xml:"requestId"`
	volumeStatusSet []VolumeStatusItemType `xml:"volumeStatusSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcAttribute.html]
type DescribeVpcAttributeResponse struct {
	RequestId          string   `xml:"requestId"`
	VpcId              string   `xml:"vpcId"`
	EnableDnsSupport   bool     `xml:"enableDnsSupport>value"`
	EnableDnsHostnames []string `xml:"enableDnsHostnames>value"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcPeeringConnections.html]
type DescribeVpcPeeringConnectionsResponse struct {
	RequestId               string                     `xml:"requestId"`
	VpcPeeringConnectionSet []VpcPeeringConnectionType `xml:"vpcPeeringConnectionSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcs.html]
type DescribeVpcsResponse struct {
	RequestId string    `xml:"requestId"`
	VpcSet    []VpcType `xml:"vpcSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnConnections.html]
type DescribeVpnConnectionsResponse struct {
	RequestId        string              `xml:"requestId"`
	VpnConnectionSet []VpnConnectionType `xml:"vpnConnectionSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnGateways.html]
type DescribeVpnGatewaysResponse struct {
	RequestId     string           `xml:"requestId"`
	VpnGatewaySet []VpnGatewayType `xml:"vpnGatewaySet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachInternetGateway.html]
type DetachInternetGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachNetworkInterface.html]
type DetachNetworkInterfaceResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVolume.html]
type DetachVolumeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVpnGateway.html]
type DetachVpnGatewayResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisableVgwRoutePropagation.html]
type DisableVgwRoutePropagationResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateAddress.html]
type DisassociateAddressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateRouteTable.html]
type DisassociateRouteTableResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVgwRoutePropagation.html]
type EnableVgwRoutePropagationResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVolumeIO.html]
type EnableVolumeIOResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetConsoleOutput.html]
type GetConsoleOutputResponse struct {
	RequestId  string    `xml:"requestId"`
	InstanceId string    `xml:"instanceId"`
	Timestamp  time.Time `xml:"timestamp"`
	Output     string    `xml:"output"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetPasswordData.html]
type GetPasswordDataResponse struct {
	RequestId    string    `xml:"requestId"`
	InstanceId   string    `xml:"instanceId"`
	Timestamp    time.Time `xml:"timestamp"`
	PasswordDate string    `xml:"passwordData"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportInstance.html]
type ImportInstanceResponse struct {
	RequestId      string             `xml:"requestId"`
	ConversionTask ConversionTaskType `xml:"conversionTask"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportKeyPair.html]
type ImportKeyPairResponse struct {
	RequestId      string `xml:"requestId"`
	KeyName        string `xml:"keyName"`
	KeyFingerprint string `xml:"keyFingerprint"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportVolume.html]
type ImportVolumeResponse struct {
	RequestId      string             `xml:"requestId"`
	ConversionTask ConversionTaskType `xml:"conversionTask"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyImageAttribute.html]
type ModifyImageAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyInstanceAttribute.html]
type ModifyInstanceAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyNetworkInterfaceAttribute.html]
type ModifyNetworkInterfaceAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyReservedInstances.html]
type ModifyReservedInstancesResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySnapshotAttribute.html]
type ModifySnapshotAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySubnetAttribute.html]
type ModifySubnetAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVolumeAttribute.html]
type ModifyVolumeAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVpcAttribute.html]
type ModifyVpcAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-MonitorInstances.html]
type MonitorInstancesResponse struct {
	RequestId    string                        `xml:"requestId"`
	InstancesSet []InstanceMonitoringStateType `xml:"instancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-PurchaseReservedInstancesOffering.html]
type PurchaseReservedInstancesOfferingResponse struct {
	RequestId                     string                                         `xml:"requestId"`
	ReservedInstancesId           string                                         `xml:"reservedInstancesId"`
	ReservedInstancesSet          []DescribeReservedInstancesResponseSetItemType `xml:"reservedInstancesSet>item"`
	ReservedInstancesOfferingsSet []DescribeReservedInstancesResponseSetItemType `xml:"reservedInstancesOfferingsSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RebootInstances.html]
type RebootInstancesResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RegisterImage.html]
type RegisterImageResponse struct {
	RequestId string `xml:"requestId"`
	ImageId   string `xml:"imageId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RejectVpcPeeringConnection.html]
type RejectVpcPeeringConnectionResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReleaseAddress.html]
type ReleaseAddressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html]
type ReplaceNetworkAclAssociationResponse struct {
	RequestId        string `xml:"requestId"`
	NewAssociationId string `xml:"newAssociationId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclEntry.html]
type ReplaceNetworkAclEntryResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRoute.html]
type ReplaceRouteResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRouteTableAssociation.html]
type ReplaceRouteTableAssociationResponse struct {
	RequestId        string `xml:"requestId"`
	NewAssociationId string `xml:"newAssociationId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReportInstanceStatus.html]
type ReportInstanceStatusResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RequestSpotInstances.html]
type RequestSpotInstancesResponse struct {
	RequestId              string                           `xml:"requestId"`
	SpotInstanceRequestSet []SpotInstanceRequestSetItemType `xml:"spotInstanceRequestSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetImageAttribute.html]
type ResetImageAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetInstanceAttribute.html]
type ResetInstanceAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetNetworkInterfaceAttribute.html]
type ResetNetworkInterfaceAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetSnapshotAttribute.html]
type ResetSnapshotAttributeResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupEgress.html]
type RevokeSecurityGroupEgressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupIngress.html]
type RevokeSecurityGroupIngressResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RunInstances.html]
type RunInstancesResponse struct {
	RequestId     string                     `xml:"requestId"`
	reservationId string                     `xml:"reservationId"`
	ownerId       string                     `xml:"ownerId"`
	groupSet      []GroupItemType            `xml:"groupSet>item"`
	InstancesSet  []RunningInstancesItemType `xml:"instancesSet>item"`
	RequesterId   string                     `xml:"requesterId"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StartInstances.html]
type StartInstancesResponse struct {
	RequestId    string                    `xml:"requestId"`
	InstancesSet []InstanceStateChangeType `xml:"instancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StopInstances.html]
type StopInstancesResponse struct {
	RequestId    string                    `xml:"requestId"`
	InstancesSet []InstanceStateChangeType `xml:"instancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-TerminateInstances.html]
type TerminateInstancesResponse struct {
	RequestId    string                    `xml:"requestId"`
	InstancesSet []InstanceStateChangeType `xml:"instancesSet>item"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnassignPrivateIpAddresses.html]
type UnassignPrivateIpAddressesResponse struct {
	RequestId string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnmonitorInstances.html]
type UnmonitorInstancesResponse struct {
	RequestId    string                    `xml:"requestId"`
	InstancesSet []InstanceStateChangeType `xml:"instancesSet>item"`
}
