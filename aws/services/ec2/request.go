package ec2

import (
	"time"
)

/*****************************************************************************/

// Accepts a VPC peering connection request. To accept a request, the VPC peering
// connection must be in the pending-acceptance state, and you must be the owner
// of the peer VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AcceptVpcPeeringConnection.html]
type AcceptVpcPeeringConnectionRequest struct {
	VpcPeeringConnectionId string `name:"VpcPeeringConnectionId"`
}

// Creates a new AcceptVpcPeeringConnectionRequest.
func NewAcceptVpcPeeringConnectionRequest(vpcPeeringConnectionId string) *AcceptVpcPeeringConnectionRequest {
	return &AcceptVpcPeeringConnectionRequest{vpcPeeringConnectionId}
}

/*****************************************************************************/

// Acquires an Elastic IP address.
// An Elastic IP address is for use either in the EC2-Classic platform or in a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AllocateAddress.html]
type AllocateAddressRequest struct {
	Domain Domain `name:"Domain,omitempty"`
}

// Creates a new AllocateAddressRequest.
func NewAllocateAddressRequest() *AllocateAddressRequest {
	return &AllocateAddressRequest{}
}

/*****************************************************************************/

// Assigns one or more secondary private IP addresses to the specified network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssignPrivateIpAddresses.html]
type AssignPrivateIpAddressesRequest struct {
	NetworkInterfaceId             string `name:"NetworkInterfaceId"`
	PrivateIpAddress               string `name:"PrivateIpAddress.#,omitempty"`
	SecondaryPrivateIpAddressCount int    `name:"SecondaryPrivateIpAddressCount,omitempty"`
	AllowReassignment              bool   `name:"AllowReassignment,omitempty"`
}

// Creates a new AssignPrivateIpAddressesRequest.
func NewAssignPrivateIpAddressesRequest(networkInterfaceId string) *AssignPrivateIpAddressesRequest {
	return &AssignPrivateIpAddressesRequest{NetworkInterfaceId: networkInterfaceId}
}

/*****************************************************************************/

// Associates an Elastic IP address with an instance or a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateAddress.html]
type AssociateAddressRequest struct {
	PublicIp           string `name:"PublicIp,omitempty"`
	InstanceId         string `name:"InstanceId,omitempty"`
	AllocationId       string `name:"AllocationId,omitempty"`
	NetworkInterfaceId string `name:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   string `name:"PrivateIpAddress,omitempty"`
	AllowReassociation bool   `name:"AllowReassociation,omitempty"`
}

// Creates a new AssociateAddressRequest.
func NewAssociateAddressRequest() *AssociateAddressRequest {
	return &AssociateAddressRequest{}
}

/*****************************************************************************/

// Associates a set of DHCP options (that you've previously created) with the
// specified VPC, or associates no DHCP options with the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateDhcpOptions.html]
type AssociateDhcpOptionsRequest struct {
	DhcpOptionsId string `name:"DhcpOptionsId"`
	VpcId         string `name:"VpcId"`
}

// Creates a new AssociateDhcpOptionsRequest.
func NewAssociateDhcpOptionsRequest(vpcId, dhcpOptionsId string) *AssociateDhcpOptionsRequest {
	return &AssociateDhcpOptionsRequest{dhcpOptionsId, vpcId}
}

/*****************************************************************************/

// Associates a subnet with a route table. The subnet and route table must be in the same VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateRouteTable.html]
type AssociateRouteTableRequest struct {
	RouteTableId string `name:"RouteTableId"`
	SubnetId     string `name:"SubnetId"`
}

// Creates a new AssociateRouteTableRequest.
func NewAssociateRouteTableRequest(routeTableId, subnetId string) *AssociateRouteTableRequest {
	return &AssociateRouteTableRequest{routeTableId, subnetId}
}

/*****************************************************************************/

// Attaches an Internet gateway to a VPC, enabling connectivity between the Internet and the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachInternetGateway.html]
type AttachInternetGatewayRequest struct {
	InternetGatewayId string `name:"InternetGatewayId"`
	VpcId             string `name:"VpcId"`
}

// Creates a new AttachInternetGatewayRequest.
func NewAttachInternetGatewayRequest(internetGatewayId, vpcId string) *AttachInternetGatewayRequest {
	return &AttachInternetGatewayRequest{internetGatewayId, vpcId}
}

/*****************************************************************************/

// Attaches a network interface to an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachNetworkInterface.html]
type AttachNetworkInterfaceRequest struct {
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
	InstanceId         string `name:"InstanceId"`
	DeviceIndex        int    `name:"DeviceIndex"`
}

// Creates a new AttachNetworkInterfaceRequest.
func NewAttachNetworkInterfaceRequest(networkInterfaceId, instanceId string, deviceIndex int) *AttachNetworkInterfaceRequest {
	return &AttachNetworkInterfaceRequest{networkInterfaceId, instanceId, deviceIndex}
}

/*****************************************************************************/

// Attaches an Amazon EBS volume to a running or stopped instance and exposes
// it to the instance with the specified device name.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVolume.html]
type AttachVolumeRequest struct {
	VolumeId   string `name:"VolumeId"`
	InstanceId string `name:"InstanceId"`
	Device     string `name:"Device"`
}

// Creates a new AttachVolumeRequest.
func NewAttachVolumeRequest(volumeId, instanceId, device string) *AttachVolumeRequest {
	return &AttachVolumeRequest{volumeId, instanceId, device}
}

/*****************************************************************************/

// Attaches a virtual private gateway to a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVpnGateway.html]
type AttachVpnGatewayRequest struct {
	VpnGatewayId string `name:"VpnGatewayId"`
	VpcId        string `name:"VpcId"`
}

// Creates a new AttachVpnGatewayRequest.
func NewAttachVpnGatewayRequest(vpcId, vpnGatewayId string) *AttachVpnGatewayRequest {
	return &AttachVpnGatewayRequest{vpnGatewayId, vpcId}
}

/*****************************************************************************/

// Adds one or more egress rules to a security group for use with a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupEgress.html]
type AuthorizeSecurityGroupEgressRequest struct {
	GroupId       string             `name:"GroupId"`
	IpPermissions []IpPermissionType `name:"IpPermissions.#." base:"1"`
}

// Creates a new AuthorizeSecurityGroupEgressRequest.
func NewAuthorizeSecurityGroupEgressRequest(groupId string, ipPermissions ...IpPermissionType) *AuthorizeSecurityGroupEgressRequest {
	return &AuthorizeSecurityGroupEgressRequest{groupId, ipPermissions}
}

/*****************************************************************************/

// Adds one or more ingress rules to a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html]
type AuthorizeSecurityGroupIngressRequest struct {
	GroupId       string            `name:"GroupId,omitempty"`
	GroupName     string            `name:"GroupName,omitempty"`
	IpPermissions *IpPermissionType `name:"IpPermissions.#.,omitempty" base:"1"`
}

// Creates a new AuthorizeSecurityGroupIngressRequest.
func NewAuthorizeSecurityGroupIngressRequest() *AuthorizeSecurityGroupIngressRequest {
	return &AuthorizeSecurityGroupIngressRequest{}
}

/*****************************************************************************/

// Bundles an Amazon instance store-backed Windows instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-BundleInstance.html]
type BundleInstanceRequest struct {
	InstanceId string                      `name:"InstanceId"`
	S3Storage  BundleInstanceS3StorageType `name:"Storage.S3."`
}

// Creates a new BundleInstanceRequest.
func NewBundleInstanceRequest(instanceId string, s3Storage BundleInstanceS3StorageType) *BundleInstanceRequest {
	return &BundleInstanceRequest{instanceId, s3Storage}
}

/*****************************************************************************/

// Cancels a bundling operation for an instance store-backed Windows instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelBundleTask.html]
type CancelBundleTaskRequest struct {
	BundleId string `name:"BundleId"`
}

// Creates a new CancelBundleTaskRequest.
func NewCancelBundleTaskRequest(bundleId string) *CancelBundleTaskRequest {
	return &CancelBundleTaskRequest{bundleId}
}

/*****************************************************************************/

// Cancels an active conversion task. The task can be the import of an instance or volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelConversionTask.html]
type CancelConversionTaskRequest struct {
	ConversionTaskId string `name:"ConversionTaskId"`
}

// Creates a new CancelConversionTaskRequest.
func NewCancelConversionTaskRequest(conversionTaskId string) *CancelConversionTaskRequest {
	return &CancelConversionTaskRequest{conversionTaskId}
}

/*****************************************************************************/

// Cancels an active export task. The request removes all artifacts of the export,
// including any partially-created Amazon S3 objects.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelExportTask.html]
type CancelExportTaskRequest struct {
	ExportTaskId string `name:"ExportTaskId"`
}

// Creates a new CancelExportTaskRequest.
func NewCancelExportTaskRequest(exportTaskId string) *CancelExportTaskRequest {
	return &CancelExportTaskRequest{exportTaskId}
}

/*****************************************************************************/

// Cancels the specified Reserved Instance listing in the Reserved Instance Marketplace.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelReservedInstancesListing.html]
type CancelReservedInstancesListingRequest struct {
	ReservedInstancesListingId string `name:"reservedInstancesListingId"`
}

// Creates a new CancelReservedInstancesListingRequest.
func NewCancelReservedInstancesListingRequest(reservedInstancesListingId string) *CancelReservedInstancesListingRequest {
	return &CancelReservedInstancesListingRequest{reservedInstancesListingId}
}

/*****************************************************************************/

// Cancels one or more Spot Instance requests. Spot Instances are instances that
// Amazon EC2 starts on your behalf when the maximum price that you specify exceeds
// the current Spot Price. Amazon EC2 periodically sets the Spot Price based on
// available Spot Instance capacity and current Spot Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelSpotInstanceRequests.html]
type CancelSpotInstanceRequestsRequest struct {
	SpotInstanceRequestIds []struct {
		Name  string `name:Name`
		Value string `name:Value`
	} `name:"SpotInstanceRequestId.#" base:"1"`
}

// Creates a new CancelSpotInstanceRequestsRequest.
func NewCancelSpotInstanceRequestsRequest(spotInstanceName, spotInstanceValue string) *CancelSpotInstanceRequestsRequest {
	r := &CancelSpotInstanceRequestsRequest{}
	r.AddSpotInstanceRequestId(spotInstanceName, spotInstanceValue)
	return r
}

// Add a SpotInstanceRequest.
func (r *CancelSpotInstanceRequestsRequest) AddSpotInstanceRequestId(spotInstanceName, spotInstanceValue string) {
	r.SpotInstanceRequestIds = append(r.SpotInstanceRequestIds, struct {
		Name  string `name:Name`
		Value string `name:Value`
	}{spotInstanceName, spotInstanceValue})
}

/*****************************************************************************/

// Determines whether a product code is associated with an instance. This action
// can only be used by the owner of the product code. It is useful when a product
// code owner needs to verify whether another user's instance is eligible for support.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ConfirmProductInstance.html]
type ConfirmProductInstanceRequest struct {
	ProductCode string `name:"ProductCode"`
	InstanceId  string `name:"InstanceId"`
}

// Creates a new ConfirmProductInstanceRequest.
func NewConfirmProductInstanceRequest(instanceId, productCode string) *ConfirmProductInstanceRequest {
	return &ConfirmProductInstanceRequest{productCode, instanceId}
}

/*****************************************************************************/

// Initiates the copy of an AMI from the specified source region to the current region.
// You specify the destination region by using its endpoint when making the request.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopyImage.html]
type CopyImageRequest struct {
	SourceRegion  string `name:"SourceRegion"`
	SourceImageId string `name:"SourceImageId"`
	Name          string `name:"Name,omitempty"`
	Description   string `name:"Description,omitempty"`
	ClientToken   string `name:"ClientToken,omitempty"`
}

// Creates a new CopyImageRequest.
func NewCopyImageRequest(sourceRegion, sourceImageId string) *CopyImageRequest {
	return &CopyImageRequest{SourceRegion: sourceRegion, SourceImageId: sourceImageId}
}

/*****************************************************************************/

// Copies a point-in-time snapshot of an Amazon Elastic Block Store (Amazon EBS) volume and
// stores it in Amazon Simple Storage Service (Amazon S3). You can copy the snapshot within
// the same region or from one region to another. The snapshot is copied to the regional
// endpoint that you send the HTTP request to.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopySnapshot.html]
type CopySnapshotRequest struct {
	SourceRegion      string `name:"SourceRegion"`
	SourceSnapshotId  string `name:"SourceSnapshotId"`
	Description       string `name:"Description,omitempty"`
	DestinationRegion string `name:"DestinationRegion,omitempty"`
	PresignedUrl      string `name:"DestinationRegion,omitempty"`
}

// Creates a new CopySnapshotRequest.
func NewCopySnapshotRequest(sourceRegion, sourceSnapshotId string) *CopySnapshotRequest {
	return &CopySnapshotRequest{SourceRegion: sourceRegion, SourceSnapshotId: sourceSnapshotId}
}

/*****************************************************************************/

// Provides information to AWS about your VPN customer gateway device. The customer
// gateway is the appliance at your end of the VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateCustomerGateway.html]
type CreateCustomerGatewayRequest struct {
	Type      string `name:"Type"`
	IpAddress string `name:"IpAddress"`
	BgpAsn    int    `name:"BgpAsn,omitempty"`
}

// Creates a new CreateCustomerGatewayRequest.
func NewCreateCustomerGatewayRequest(t, ipAddress string) *CreateCustomerGatewayRequest {
	return &CreateCustomerGatewayRequest{Type: t, IpAddress: ipAddress}
}

/*****************************************************************************/

// Creates a set of DHCP options for your VPC. After creating the set, you must associate
// it with the VPC, causing all existing and new instances that you launch in the VPC
// to use this set of DHCP options. The following are the individual DHCP options you can specify.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateDhcpOptions.html]
type CreateDhcpOptionsRequest struct {
	DhcpConfigurations []DhcpConfigurationItemType `name:"DhcpConfiguration.#." base:"1"`
}

// Creates a new CreateDhcpOptionsRequest.
func NewCreateDhcpOptionsRequest(dhcpConfigurations ...DhcpConfigurationItemType) *CreateDhcpOptionsRequest {
	return &CreateDhcpOptionsRequest{dhcpConfigurations}
}

/*****************************************************************************/

// Creates an Amazon EBS-backed AMI from an Amazon EBS-backed instance that is either running or stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateImage.html]
type CreateImageRequest struct {
	InstanceId         string                       `name:"InstanceId"`
	Name               string                       `name:"Name"`
	Description        string                       `name:"Description,omitempty"`
	NoReboot           bool                         `name:"NoReboot,omitempty"`
	BlockDeviceMapping []BlockDeviceMappingItemType `name:"BlockDeviceMapping.#.,omitempty" base:"1"`
}

// Creates a new CreateImageRequest.
func NewCreateImageRequest(instanceId, name string) *CreateImageRequest {
	return &CreateImageRequest{InstanceId: instanceId, Name: name}
}

/*****************************************************************************/

// Exports a running or stopped instance to an Amazon S3 bucket.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInstanceExportTask.html]
type CreateInstanceExportTaskRequest struct {
	Description       string                     `name:"Description,omitempty"`
	InstanceId        string                     `name:"InstanceId"`
	TargetEnvironment VirtualizationType         `name:"TargetEnvironment"`
	ExportToS3        ExportToS3TaskResponseType `name:"ExportToS3."`
}

// Creates a new CreateInstanceExportTaskRequest.
func NewCreateInstanceExportTaskRequest(instanceId string, targetEnvironment VirtualizationType, s3Bucket string) *CreateInstanceExportTaskRequest {
	return &CreateInstanceExportTaskRequest{
		InstanceId:        instanceId,
		TargetEnvironment: targetEnvironment,
		ExportToS3:        ExportToS3TaskResponseType{S3Bucket: s3Bucket},
	}
}

/*****************************************************************************/

// Creates an Internet gateway for use with a VPC. After creating the Internet
// gateway, you attach it to a VPC using AttachInternetGateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInternetGateway.html]
type CreateInternetGatewayRequest struct{}

/*****************************************************************************/

// Creates a 2048-bit RSA key pair with the specified name. Amazon EC2 stores the
// public key and displays the private key for you to save to a file. The private
// key is returned as an unencrypted PEM encoded PKCS#8 private key. If a key with
// the specified name already exists, Amazon EC2 returns an error.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateKeyPair.html]
type CreateKeyPairRequest struct {
	KeyName string `name:"KeyName"`
}

// Creates a new CreateKeyPairRequest.
func NewCreateKeyPairRequest(keyName string) *CreateKeyPairRequest {
	return &CreateKeyPairRequest{keyName}
}

/*****************************************************************************/

// Creates a network ACL in a VPC. Network ACLs provide an optional layer of security
// (in addition to security groups) for the instances in your VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAcl.html]
type CreateNetworkAclRequest struct {
	VpcId string `name:"VpcId"`
}

// Creates a new CreateNetworkAclRequest.
func NewCreateNetworkAclRequest(vpcId string) *CreateNetworkAclRequest {
	return &CreateNetworkAclRequest{vpcId}
}

/*****************************************************************************/

// Creates an entry (a rule) in a network ACL with the specified rule number. Each network ACL
// has a set of numbered ingress rules and a separate set of numbered egress rules. When
// determining whether a packet should be allowed in or out of a subnet associated with the
// ACL, we process the entries in the ACL according to the rule numbers, in ascending order.
// Each network ACL has a set of ingress rules and a separate set of egress rules.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAclEntry.html]
type CreateNetworkAclEntryRequest struct {
	NetworkAclId string            `name:"NetworkAclId"`
	RuleNumber   int               `name:"RuleNumber"`
	Protocol     int               `name:"Protocol"`
	RuleAction   RuleAction        `name:"RuleAction"`
	Egress       bool              `name:"Egress,omitempty"`
	CidrBlock    string            `name:"CidrBlock"`
	Icmp         *IcmpTypeCodeType `name:"Icmp,omitempty"`
	PortRange    *PortRangeType    `name:"PortRange,omitempty"`
}

// Creates a new CreateNetworkAclEntryRequest.
func NewCreateNetworkAclEntryRequest(networkAclId, cidrBlock string, ruleAction RuleAction, ruleNumber, protocol int) *CreateNetworkAclEntryRequest {
	return &CreateNetworkAclEntryRequest{
		NetworkAclId: networkAclId,
		RuleNumber:   ruleNumber,
		Protocol:     protocol,
		RuleAction:   ruleAction,
		CidrBlock:    cidrBlock,
	}
}

/*****************************************************************************/

// Creates a network interface in the specified subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkInterface.html]
type CreateNetworkInterfaceRequest struct {
	SubnetId                       string                                 `name:"SubnetId"`
	PrivateIpAddress               string                                 `name:"PrivateIpAddress,omitempty"`
	PrivateIpAddresses             []PrivateIpAddressesSetItemRequestType `name:"PrivateIpAddresses.#.,omitempty"`
	SecondaryPrivateIpAddressCount int                                    `name:"SecondaryPrivateIpAddressCount,omitempty"`
	Description                    string                                 `name:"Description,omitempty"`
	SecurityGroupId                []string                               `name:"SecurityGroupId.#,omitempty"`
}

// Creates a new CreateNetworkInterfaceRequest.
func NewCreateNetworkInterfaceRequest(subnetId string) *CreateNetworkInterfaceRequest {
	return &CreateNetworkInterfaceRequest{SubnetId: subnetId}
}

/*****************************************************************************/

// Creates a placement group that you launch cluster instances into. You must
// give the group a name that's unique within the scope of your account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreatePlacementGroup.html]
type CreatePlacementGroupRequest struct {
	GroupName string `name:"GroupName"`
	Strategy  string `name:"Strategy"`
}

// Creates a new CreatePlacementGroupRequest.
func NewCreatePlacementGroupRequest(groupName, strategy string) *CreatePlacementGroupRequest {
	return &CreatePlacementGroupRequest{groupName, strategy}
}

/*****************************************************************************/

// Creates a listing for Amazon EC2 Reserved Instances to be sold in the Reserved
// Instance Marketplace. You can submit one Reserved Instance listing at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateReservedInstancesListing.html]
type CreateReservedInstancesListingRequest struct {
	ReservedInstancesId string                            `name:"ReservedInstancesId"`
	InstanceCount       int                               `name:"InstanceCount"`
	PriceSchedules      []PriceScheduleRequestSetItemType `name:"PriceSchedules.#."`
	ClientToken         string                            `name:"ClientToken"`
}

// Creates a new CreateReservedInstancesListingRequest.
func NewCreateReservedInstancesListingRequest(reservedInstancesId, clientToken string, instanceCount int, priceSchedules ...PriceScheduleRequestSetItemType) *CreateReservedInstancesListingRequest {
	return &CreateReservedInstancesListingRequest{reservedInstancesId, instanceCount, priceSchedules, clientToken}
}

/*****************************************************************************/

// Creates a route in a route table within a VPC. The route's target can be an
// Internet gateway or virtual private gateway attached to the VPC, a VPC peering
// connection, or a NAT instance in the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRoute.html]
type CreateRouteRequest struct {
	RouteTableId           string `name:"RouteTableId"`
	DestinationCidrBlock   string `name:"DestinationCidrBlock"`
	GatewayId              string `name:"GatewayId,omitempty"`
	InstanceId             string `name:"InstanceId,omitempty"`
	NetworkInterfaceId     string `name:"NetworkInterfaceId,omitempty"`
	VpcPeeringConnectionId string `name:"VpcPeeringConnectionId,omitempty"`
}

// Creates a new CreateRouteRequest.
func NewCreateRouteRequest(routeTableId, destinationCidrBlock string) *CreateRouteRequest {
	return &CreateRouteRequest{RouteTableId: routeTableId, DestinationCidrBlock: destinationCidrBlock}
}

/*****************************************************************************/

// Creates a route table for the specified VPC. After you create a route table,
// you can add routes and associate the table with a subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRouteTable.html]
type CreateRouteTableRequest struct {
	VpcId string `name:"VpcId"`
}

// Creates a new CreateRouteTableRequest.
func NewCreateRouteTableRequest(vpcId string) *CreateRouteTableRequest {
	return &CreateRouteTableRequest{vpcId}
}

/*****************************************************************************/

// Creates a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSecurityGroup.html]
type CreateSecurityGroupRequest struct {
	GroupName        string `name:"GroupName"`
	GroupDescription string `name:"GroupDescription"`
	VpcId            string `name:"VpcId,omitempty"`
}

// Creates a new CreateSecurityGroupRequest.
func NewCreateSecurityGroupRequest(groupName, groupDescription string) *CreateSecurityGroupRequest {
	return &CreateSecurityGroupRequest{GroupName: groupName, GroupDescription: groupDescription}
}

/*****************************************************************************/

// Creates a snapshot of an Amazon EBS volume and stores it in Amazon S3. You can use snapshots
// for backups, to make copies of instance store volumes, and to save data before shutting
// down an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSnapshot.html]
type CreateSnapshotRequest struct {
	VolumeId    string `name:"VolumeId"`
	Description string `name:"Description,omitempty"`
}

// Creates a new CreateSnapshotRequest.
func NewCreateSnapshotRequest(volumeId string) *CreateSnapshotRequest {
	return &CreateSnapshotRequest{VolumeId: volumeId}
}

/*****************************************************************************/

// Creates the datafeed for Spot Instances, enabling you to view Spot Instance usage logs.
// You can create one data feed per account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSpotDatafeedSubscription.html]
type CreateSpotDatafeedSubscriptionRequest struct {
	Bucket string `name:"Bucket"`
	Prefix string `name:"Prefix,omitempty"`
}

// Creates a new CreateSpotDatafeedSubscriptionRequest.
func NewCreateSpotDatafeedSubscriptionRequest(bucket string) *CreateSpotDatafeedSubscriptionRequest {
	return &CreateSpotDatafeedSubscriptionRequest{Bucket: bucket}
}

/*****************************************************************************/

// Creates a subnet in an existing VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSubnet.html]
type CreateSubnetRequest struct {
	VpcId            string `name:"VpcId"`
	CidrBlock        string `name:"CidrBlock"`
	AvailabilityZone string `name:"AvailabilityZone,omitempty"`
}

// Creates a new CreateSubnetRequest.
func NewCreateSubnetRequest(vpcId, cidrBlock string) *CreateSubnetRequest {
	return &CreateSubnetRequest{VpcId: vpcId, CidrBlock: cidrBlock}
}

/*****************************************************************************/

// Adds or overwrites one or more tags for the specified Amazon EC2 resource or resources.
// Each resource can have a maximum of 10 tags. Each tag consists of a key and optional value.
// Tag keys must be unique per resource.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateTags.html]
type CreateTagsRequest struct {
	ResourceIds []string         `name:"ResourceId.#" base:"1"`
	Tags        []TagSetItemType `name:"Tag.#." base:"1"`
}

// Creates a new CreateTagsRequest.
func NewCreateTagsRequest(resourceId, tagKey, tagValue string) *CreateTagsRequest {
	return &CreateTagsRequest{
		[]string{resourceId},
		[]TagSetItemType{TagSetItemType{Key: tagKey, Value: tagValue}},
	}
}

// Add a tag.
func (r *CreateTagsRequest) AddTag(resourceId, tagKey, tagValue string) {
	r.ResourceIds = append(r.ResourceIds, resourceId)
	r.Tags = append(r.Tags, TagSetItemType{Key: tagKey, Value: tagValue})
}

/*****************************************************************************/

// Creates an Amazon EBS volume that can be attached to an instance in the same Availability Zone.
// The volume is created in the regional endpoint that you send the HTTP request to.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html]
type CreateVolumeRequest struct {
	Size             string     `name:"Size,omitempty"`
	SnapshotId       string     `name:"SnapshotId,omitempty"`
	AvailabilityZone string     `name:"AvailabilityZone"`
	VolumeType       VolumeType `name:"VolumeType,omitempty"`
	Iops             int        `name:"Iops,omitempty"`
	Encrypted        bool       `name:"Encrypted,omitempty"`
}

// Creates a new CreateVolumeRequest.
func NewCreateVolumeRequest(availabilityZone string) *CreateVolumeRequest {
	return &CreateVolumeRequest{AvailabilityZone: availabilityZone}
}

/*****************************************************************************/

// Creates a VPC with the specified CIDR block.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpc.html]
type CreateVpcRequest struct {
	CidrBlock       string          `name:"CidrBlock"`
	InstanceTenancy InstanceTenancy `name:"InstanceTenancy,omitempty"`
}

// Creates a new CreateVpcRequest.
func NewCreateVpcRequest(cidrBlock string) *CreateVpcRequest {
	return &CreateVpcRequest{CidrBlock: cidrBlock}
}

/*****************************************************************************/

// Requests a VPC peering connection between two VPCs: a requester VPC that you own and a
// peer VPC with which to create the connection. The peer VPC can belong to another
// AWS account. The requester VPC and peer VPC cannot have overlapping CIDR blocks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpcPeeringConnection.html]
type CreateVpcPeeringConnectionRequest struct {
	VpcId       string `name:"VpcId"`
	PeerVpcId   string `name:"PeerVpcId"`
	PeerOwnerId string `name:"PeerOwnerId,omitempty"`
}

// Creates a new CreateVpcPeeringConnectionRequest.
func NewCreateVpcPeeringConnectionRequest(vpcId, peerVpcId string) *CreateVpcPeeringConnectionRequest {
	return &CreateVpcPeeringConnectionRequest{VpcId: vpcId, PeerVpcId: peerVpcId}
}

/*****************************************************************************/

// Creates a VPN connection between an existing virtual private gateway and a VPN customer gateway.
// The only supported connection type is ipsec.1.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnection.html]
type CreateVpnConnectionRequest struct {
	Type              string                            `name:"Type"`
	CustomerGatewayId string                            `name:"CustomerGatewayId"`
	VpnGatewayId      string                            `name:"VpnGatewayId"`
	Options           *VpnConnectionOptionsResponseType `name:"Options.StaticRoutesOnly,omitempty"`
}

// Creates a new CreateVpnConnectionRequest.
func NewCreateVpnConnectionRequest(connType, vpnGatewayId, customerGatewayId string) *CreateVpnConnectionRequest {
	return &CreateVpnConnectionRequest{Type: connType, CustomerGatewayId: customerGatewayId, VpnGatewayId: vpnGatewayId}
}

/*****************************************************************************/

// Creates a static route associated with a VPN connection between an existing virtual
// private gateway and a VPN customer gateway. The static route allows traffic to be
// routed from the virtual private gateway to the VPN customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnectionRoute.html]
type CreateVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `name:"DestinationCidrBlock"`
	VpnConnectionId      string `name:"VpnConnectionId"`
}

// Creates a new CreateVpnConnectionRouteRequest.
func NewCreateVpnConnectionRouteRequest(vpnConnectionId, destinationCidrBlock string) *CreateVpnConnectionRouteRequest {
	return &CreateVpnConnectionRouteRequest{destinationCidrBlock, vpnConnectionId}
}

/*****************************************************************************/

// Creates a virtual private gateway. A virtual private gateway is the endpoint on the VPC side
// of your VPN connection. You can create a virtual private gateway before creating the VPC itself.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnGateway.html]
type CreateVpnGatewayRequest struct {
	Type string `name:"Type"`
}

// Creates a new CreateVpnGatewayRequest.
func NewCreateVpnGatewayRequest(vpnGatewayType string) *CreateVpnGatewayRequest {
	return &CreateVpnGatewayRequest{vpnGatewayType}
}

/*****************************************************************************/

// Deletes the specified customer gateway. You must delete the VPN connection
// before you can delete the customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteCustomerGateway.html]
type DeleteCustomerGatewayRequest struct {
	CustomerGatewayId string `name:"CustomerGatewayId"`
}

// Creates a new DeleteCustomerGatewayRequest.
func NewDeleteCustomerGatewayRequest(customerGatewayId string) *DeleteCustomerGatewayRequest {
	return &DeleteCustomerGatewayRequest{customerGatewayId}
}

/*****************************************************************************/

// Deletes the specified set of DHCP options. You must disassociate the set of DHCP options
// before you can delete it. You can disassociate the set of DHCP options by associating
// either a new set of options or the default set of options with the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteDhcpOptions.html]
type DeleteDhcpOptionsRequest struct {
	DhcpOptionsId string `name:"DhcpOptionsId"`
}

// Creates a new DeleteDhcpOptionsRequest.
func NewDeleteDhcpOptionsRequest(dhcpOptionsId string) *DeleteDhcpOptionsRequest {
	return &DeleteDhcpOptionsRequest{dhcpOptionsId}
}

/*****************************************************************************/

// Deletes the specified Internet gateway. You must detach the Internet gateway
// from the VPC before you can delete it.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteInternetGateway.html]
type DeleteInternetGatewayRequest struct {
	InternetGatewayId string `name:"InternetGatewayId"`
}

// Creates a new DeleteInternetGatewayRequest.
func NewDeleteInternetGatewayRequest(internetGatewayId string) *DeleteInternetGatewayRequest {
	return &DeleteInternetGatewayRequest{internetGatewayId}
}

/*****************************************************************************/

// Deletes the specified key pair, by removing the public key from Amazon EC2.
// You must own the key pair.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteKeyPair.html]
type DeleteKeyPairRequest struct {
	KeyName string `name:"KeyName"`
}

// Creates a new DeleteKeyPairRequest.
func NewDeleteKeyPairRequest(keyName string) *DeleteKeyPairRequest {
	return &DeleteKeyPairRequest{keyName}
}

/*****************************************************************************/

// Deletes the specified network ACL. You can't delete the ACL if it's associated
// with any subnets. You can't delete the default network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAcl.html]
type DeleteNetworkAclRequest struct {
	NetworkAclId string `name:"NetworkAclId"`
}

// Creates a new DeleteNetworkAclRequest.
func NewDeleteNetworkAclRequest(networkAclId string) *DeleteNetworkAclRequest {
	return &DeleteNetworkAclRequest{networkAclId}
}

/*****************************************************************************/

// Deletes the specified ingress or egress entry (rule) from the specified network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAclEntry.html]
type DeleteNetworkAclEntryRequest struct {
	NetworkAclId string `name:"NetworkAclId"`
	RuleNumber   int    `name:"RuleNumber"`
	Egress       bool   `name:"Egress,omitempty"`
}

// Creates a new DeleteNetworkAclEntryRequest.
func NewDeleteNetworkAclEntryRequest(networkAclId string, ruleNumber int) *DeleteNetworkAclEntryRequest {
	return &DeleteNetworkAclEntryRequest{NetworkAclId: networkAclId, RuleNumber: ruleNumber}
}

/*****************************************************************************/

// Deletes the specified network interface. You must detach the network interface
// before you can delete it.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkInterface.html]
type DeleteNetworkInterfaceRequest struct {
	NetworkInterfaceId string `name:"NetworkInterfaceId"`
}

// Creates a new DeleteNetworkInterfaceRequest.
func NewDeleteNetworkInterfaceRequest(networkInterfaceId string) *DeleteNetworkInterfaceRequest {
	return &DeleteNetworkInterfaceRequest{networkInterfaceId}
}

/*****************************************************************************/

// Deletes the specified placement group. You must terminate all instances in the
// placement group before you can delete the placement group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeletePlacementGroup.html]
type DeletePlacementGroupRequest struct {
	GroupName string `name:"GroupName"`
}

// Creates a new DeletePlacementGroupRequest.
func NewDeletePlacementGroupRequest(groupName string) *DeletePlacementGroupRequest {
	return &DeletePlacementGroupRequest{groupName}
}

/*****************************************************************************/

// Deletes the specified route from the specified route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRoute.html]
type DeleteRouteRequest struct {
	RouteTableId         string `name:"RouteTableId"`
	DestinationCidrBlock string `name:"DestinationCidrBlock"`
}

// Creates a new DeleteRouteRequest.
func NewDeleteRouteRequest(routeTableId, destinationCidrBlock string) *DeleteRouteRequest {
	return &DeleteRouteRequest{routeTableId, destinationCidrBlock}
}

/*****************************************************************************/

// Deletes the specified route table. You must disassociate the route table from any
// subnets before you can delete it. You can't delete the main route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRouteTable.html]
type DeleteRouteTableRequest struct {
	RouteTableId string `name:"RouteTableId"`
}

// Creates a new DeleteRouteTableRequest.
func NewDeleteRouteTableRequest(routeTableId string) *DeleteRouteTableRequest {
	return &DeleteRouteTableRequest{routeTableId}
}

/*****************************************************************************/

// Deletes a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSecurityGroup.html]
type DeleteSecurityGroupRequest struct {
	GroupName string `name:"GroupName,omitempty"`
	GroupId   string `name:"GroupId,omitempty"`
}

// Creates a new DeleteSecurityGroupRequest.
func NewDeleteSecurityGroupRequest() *DeleteSecurityGroupRequest {
	return &DeleteSecurityGroupRequest{}
}

/*****************************************************************************/

// Deletes the specified snapshot. When you make periodic snapshots of a volume,
// the snapshots are incremental, and only the blocks on the device that have
// changed since your last snapshot are saved in the new snapshot. When you delete
// a snapshot, only the data not needed for any other snapshot is removed. So
// regardless of which prior snapshots have been deleted, all active snapshots will
// have access to all the information needed to restore the volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSnapshot.html]
type DeleteSnapshotRequest struct {
	SnapshotId string `name:"SnapshotId"`
}

// Creates a new DeleteSnapshotRequest.
func NewDeleteSnapshotRequest(snapshotId string) *DeleteSnapshotRequest {
	return &DeleteSnapshotRequest{snapshotId}
}

/*****************************************************************************/

// Deletes the datafeed for Spot Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSpotDatafeedSubscription.html]

/*****************************************************************************/

// Deletes the specified subnet. You must terminate all running instances in the
// subnet before you can delete the subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSubnet.html]
type DeleteSubnetRequest struct {
	SubnetId string `name:"SubnetId"`
}

// Creates a new DeleteSubnetRequest.
func NewDeleteSubnetRequest(subnetId string) *DeleteSubnetRequest {
	return &DeleteSubnetRequest{subnetId}
}

/*****************************************************************************/

// Deletes the specified set of tags from the specified set of resources. This call is designed to follow a DescribeTags call.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteTags.html]
type DeleteTagsRequest struct {
	ResourceIds []string         `name:"ResourceId.#" base:"1"`
	Tags        []TagSetItemType `name:"Tag.#." base:"1"`
}

// Creates a new DeleteTagsRequest.
func NewDeleteTagsRequest(resourceId, tagKey, tagValue string) *DeleteTagsRequest {
	return &DeleteTagsRequest{
		[]string{resourceId},
		[]TagSetItemType{TagSetItemType{Key: tagKey, Value: tagValue}},
	}
}

// Add a tag.
func (r *DeleteTagsRequest) AddTag(resourceId, tagKey, tagValue string) {
	r.ResourceIds = append(r.ResourceIds, resourceId)
	r.Tags = append(r.Tags, TagSetItemType{Key: tagKey, Value: tagValue})
}

/*****************************************************************************/

// Deletes the specified Amazon EBS volume. The volume must be in the available
// state (not attached to an instance).
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVolume.html]
type DeleteVolumeRequest struct {
	VolumeId string `name:"VolumeId"`
}

// Creates a new DeleteVolumeRequest.
func NewDeleteVolumeRequest(volumeId string) *DeleteVolumeRequest {
	return &DeleteVolumeRequest{volumeId}
}

/*****************************************************************************/

// Deletes the specified VPC. You must detach or delete all gateways and resources that
// are associated with the VPC before you can delete it. For example, you must terminate
// all instances running in the VPC, delete all security groups associated with the VPC
// (except the default one), delete all route tables associated with the VPC
// (except the default one), and so on.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpc.html]
type DeleteVpcRequest struct {
	VpcId string `name:"VpcId"`
}

// Creates a new DeleteVpcRequest.
func NewDeleteVpcRequest(vpcId string) *DeleteVpcRequest {
	return &DeleteVpcRequest{vpcId}
}

/*****************************************************************************/

// Deletes a VPC peering connection. Either the owner of the requester VPC or the
// owner of the peer VPC can delete the VPC peering connection if it's in the
// active state. The owner of the requester VPC can delete a VPC peering connection
// in the pending-acceptance state.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpcPeeringConnection.html]
type DeleteVpcPeeringConnectionRequest struct {
	VpcPeeringConnectionId string `name:"VpcPeeringConnectionId"`
}

// Creates a new DeleteVpcPeeringConnectionRequest.
func NewDeleteVpcPeeringConnectionRequest(vpcPeeringConnectionId string) *DeleteVpcPeeringConnectionRequest {
	return &DeleteVpcPeeringConnectionRequest{vpcPeeringConnectionId}
}

/*****************************************************************************/

// Deletes the specified VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnection.html]
type DeleteVpnConnectionRequest struct {
	VpnConnectionId string `name:"VpnConnectionId"`
}

// Creates a new DeleteVpnConnectionRequest.
func NewDeleteVpnConnectionRequest(vpnConnectionId string) *DeleteVpnConnectionRequest {
	return &DeleteVpnConnectionRequest{vpnConnectionId}
}

/*****************************************************************************/

// Deletes the specified static route associated with a VPN connection between an
// existing virtual private gateway and a VPN customer gateway. The static route
// allows traffic to be routed from the virtual private gateway to the VPN customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnectionRoute.html]
type DeleteVpnConnectionRouteRequest struct {
	DestinationCidrBlock string `name:"DestinationCidrBlock"`
	VpnConnectionId      string `name:"VpnConnectionId"`
}

// Creates a new DeleteVpnConnectionRouteRequest.
func NewDeleteVpnConnectionRouteRequest(vpnConnectionId, destinationCidrBlock string) *DeleteVpnConnectionRouteRequest {
	return &DeleteVpnConnectionRouteRequest{destinationCidrBlock, vpnConnectionId}
}

/*****************************************************************************/

// Deletes the specified virtual private gateway. We recommend that before you delete
// a virtual private gateway, you detach it from the VPC and delete the VPN connection.
// Note that you don't need to delete the virtual private gateway if you plan to delete
// and recreate the VPN connection between your VPC and your network.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnGateway.html]
type DeleteVpnGatewayRequest struct {
	VpnGatewayId string `name:"VpnGatewayId"`
}

// Creates a new DeleteVpnGatewayRequest.
func NewDeleteVpnGatewayRequest(vpnGatewayId string) *DeleteVpnGatewayRequest {
	return &DeleteVpnGatewayRequest{vpnGatewayId}
}

/*****************************************************************************/

// Deregisters the specified AMI. After you deregister an AMI, it can't be used
// to launch new instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeregisterImage.html]
type DeregisterImageRequest struct {
	ImageId string `name:"ImageId"`
}

// Creates a new DeregisterImageRequest.
func NewDeregisterImageRequest(imageId string) *DeregisterImageRequest {
	return &DeregisterImageRequest{imageId}
}

/*****************************************************************************/

// Describes the specified attribute of your AWS account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAccountAttributes.html]
type DescribeAccountAttributesRequest struct {
	AttributeNames []string `name:"AttributeName.#" base:"1"`
}

// Creates a new DescribeAccountAttributesRequest.
func NewDescribeAccountAttributesRequest(attributeNames ...string) *DescribeAccountAttributesRequest {
	return &DescribeAccountAttributesRequest{attributeNames}
}

/*****************************************************************************/

// Describes one or more of your Elastic IP addresses.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAddresses.html]
type DescribeAddressesRequest struct {
	PublicIps     []string         `name:"PublicIp.#,omitempty" base:"1"`
	AllocationIds []string         `name:"AllocationId.#,omitempty" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeAddressesRequest.
func NewDescribeAddressesRequest() *DescribeAddressesRequest {
	return &DescribeAddressesRequest{}
}

/*****************************************************************************/

// Describes one or more of the Availability Zones that are available to you. The results include
// zones only for the region you're currently using. If there is an event impacting an
// Availability Zone, you can use this request to view the state and any provided message for
// that Availability Zone.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAvailabilityZones.html]
type DescribeAvailabilityZonesRequest struct {
	ZoneNames []string         `name:"ZoneName.#,omitempty" base:"1"`
	Filters   []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeAvailabilityZonesRequest.
func NewDescribeAvailabilityZonesRequest() *DescribeAvailabilityZonesRequest {
	return &DescribeAvailabilityZonesRequest{}
}

/*****************************************************************************/

// Describes one or more of your bundling tasks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeBundleTasks.html]
type DescribeBundleTasksRequest struct {
	BundleIds []string         `name:"BundleId.#,omitempty" base:"1"`
	Filters   []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeBundleTasksRequest.
func NewDescribeBundleTasksRequest() *DescribeBundleTasksRequest {
	return &DescribeBundleTasksRequest{}
}

/*****************************************************************************/

// Describes one or more of your conversion tasks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeConversionTasks.html]
type DescribeConversionTasksRequest struct {
	ConversionTaskIds []string `name:"ConversionTaskId.#,omitempty" base:"1"`
}

// Creates a new DescribeConversionTasksRequest.
func NewDescribeConversionTasksRequest(conversionTaskIds ...string) *DescribeConversionTasksRequest {
	return &DescribeConversionTasksRequest{conversionTaskIds}
}

/*****************************************************************************/

// Describes one or more of your VPN customer gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeCustomerGateways.html]
type DescribeCustomerGatewaysRequest struct {
	CustomerGatewayIds []string         `name:"CustomerGatewayId.#,omitempty" base:"1"`
	Filters            []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeCustomerGatewaysRequest.
func NewDescribeCustomerGatewaysRequest() *DescribeCustomerGatewaysRequest {
	return &DescribeCustomerGatewaysRequest{}
}

// Add a CustomerGatewayId.
func (r *DescribeCustomerGatewaysRequest) AddCustomerGatewayId(customerGatewayId string) {
	r.CustomerGatewayIds = append(r.CustomerGatewayIds, customerGatewayId)
}

// Add a filter.
func (r *DescribeCustomerGatewaysRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your DHCP options sets.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeDhcpOptions.html]
type DescribeDhcpOptionsRequest struct {
	DhcpOptionsIds []string         `name:"DhcpOptionsId.#" base:"1"`
	Filters        []FilterItemType `name:"Filter.#." base:"1"`
}

// Creates a new DescribeDhcpOptionsRequest.
func NewDescribeDhcpOptionsRequest() *DescribeDhcpOptionsRequest {
	return &DescribeDhcpOptionsRequest{}
}

// Add a DhcpOptionsId.
func (r *DescribeDhcpOptionsRequest) AddDhcpOptionsId(dhcpOptionsId string) {
	r.DhcpOptionsIds = append(r.DhcpOptionsIds, dhcpOptionsId)
}

// Add a filter.
func (r *DescribeDhcpOptionsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your export tasks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeExportTasks.html]
type DescribeExportTasksRequest struct {
	ExportTaskIds []string `name:"ExportTaskId.#" base:"1"`
}

// Creates a new DescribeExportTasksRequest.
func NewDescribeExportTasksRequest(exportTaskIds ...string) *DescribeExportTasksRequest {
	return &DescribeExportTasksRequest{exportTaskIds}
}

/*****************************************************************************/

// Describes the specified attribute of the specified AMI. You can specify only
// one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImageAttribute.html]
type DescribeImageAttributeRequest struct {
	ImageId   string         `name:"ImageId"`
	Attribute ImageAttribute `name:"Attribute"`
}

// Creates a new DescribeImageAttributeRequest.
func NewDescribeImageAttributeRequest(imageId string, attribute ImageAttribute) *DescribeImageAttributeRequest {
	return &DescribeImageAttributeRequest{imageId, attribute}
}

/*****************************************************************************/

// Describes one or more of the images (AMIs, AKIs, and ARIs) available to you.
// Images available to you include public images, private images that you own,
// and private images owned by other AWS accounts but for which you have explicit
// launch permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImages.html]
type DescribeImagesRequest struct {
	ExecutableBys []string         `name:"ExecutableBy.#,omitempty" base:"1"`
	ImageIds      []string         `name:"ImageId.#,omitempty" base:"1"`
	Owners        []string         `name:"Owner.#,omitempty" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeImagesRequest.
func NewDescribeImagesRequest() *DescribeImagesRequest {
	return &DescribeImagesRequest{}
}

// Add a ExecutableBy.
func (r *DescribeImagesRequest) AddExecutableBy(executableBy string) {
	r.ExecutableBys = append(r.ExecutableBys, executableBy)
}

// Add a ImageId.
func (r *DescribeImagesRequest) AddImageId(imageId string) {
	r.ImageIds = append(r.ImageIds, imageId)
}

// Add a Owner.
func (r *DescribeImagesRequest) AddOwner(owner string) {
	r.Owners = append(r.Owners, owner)
}

// Add a filter.
func (r *DescribeImagesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the specified attribute of the specified instance.
// You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceAttribute.html]
type DescribeInstanceAttributeRequest struct {
	InstanceId string            `name:"InstanceId"`
	Attribute  InstanceAttribute `name:"Attribute"`
}

// Creates a new DescribeInstanceAttributeRequest.
func NewDescribeInstanceAttributeRequest(instanceId string, attribute InstanceAttribute) *DescribeInstanceAttributeRequest {
	return &DescribeInstanceAttributeRequest{instanceId, attribute}
}

/*****************************************************************************/

// Describes one or more of your instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstances.html]
type DescribeInstancesRequest struct {
	InstanceIds []string         `name:"InstanceId.#,omitempty" base:"1"`
	MaxResults  int              `name:"MaxResults,omitempty"`
	NextToken   string           `name:"NextToken,omitempty"`
	Filters     []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeInstancesRequest.
func NewDescribeInstancesRequest() *DescribeInstancesRequest {
	return &DescribeInstancesRequest{}
}

// Add a InstanceId.
func (r *DescribeInstancesRequest) AddInstanceId(instanceId string) {
	r.InstanceIds = append(r.InstanceIds, instanceId)
}

// Add a filter.
func (r *DescribeInstancesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the status of one or more instances, including any scheduled events.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceStatus.html]
type DescribeInstanceStatusRequest struct {
	InstanceIds         []string         `name:"InstanceId.#,omitempty" base:"1"`
	IncludeAllInstances bool             `name:"IncludeAllInstances,omitempty"`
	MaxResults          int              `name:"MaxResults,omitempty"`
	NextToken           string           `name:"NextToken,omitempty"`
	Filters             []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeInstanceStatusRequest.
func NewDescribeInstanceStatusRequest() *DescribeInstanceStatusRequest {
	return &DescribeInstanceStatusRequest{}
}

// Add a InstanceId.
func (r *DescribeInstanceStatusRequest) AddInstanceId(instanceId string) {
	r.InstanceIds = append(r.InstanceIds, instanceId)
}

// Add a filter.
func (r *DescribeInstanceStatusRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your Internet gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInternetGateways.html]
type DescribeInternetGatewaysRequest struct {
	InternetGatewayIds []string         `name:"InternetGatewayId.#,omitempt" base:"1"`
	Filters            []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeInternetGatewaysRequest.
func NewDescribeInternetGatewaysRequest() *DescribeInternetGatewaysRequest {
	return &DescribeInternetGatewaysRequest{}
}

// Add a InternetGatewayId.
func (r *DescribeInternetGatewaysRequest) AddInternetGatewayId(internetGatewayId string) {
	r.InternetGatewayIds = append(r.InternetGatewayIds, internetGatewayId)
}

// Add a filter.
func (r *DescribeInternetGatewaysRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your key pairs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeKeyPairs.html]
type DescribeKeyPairsRequest struct {
	KeyNames []string         `name:"KeyName.#,omitempt" base:"1"`
	Filters  []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeKeyPairsRequest.
func NewDescribeKeyPairsRequest() *DescribeKeyPairsRequest {
	return &DescribeKeyPairsRequest{}
}

// Add a KeyName.
func (r *DescribeKeyPairsRequest) AddKeyName(keyName string) {
	r.KeyNames = append(r.KeyNames, keyName)
}

// Add a filter.
func (r *DescribeKeyPairsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your network ACLs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkAcls.html]
type DescribeNetworkAclsRequest struct {
	NetworkAclIds []string         `name:"NetworkAclId.#,omitempt" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeNetworkAclsRequest.
func NewDescribeNetworkAclsRequest() *DescribeNetworkAclsRequest {
	return &DescribeNetworkAclsRequest{}
}

// Add a NetworkAclId.
func (r *DescribeNetworkAclsRequest) AddNetworkAclId(networkAclId string) {
	r.NetworkAclIds = append(r.NetworkAclIds, networkAclId)
}

// Add a filter.
func (r *DescribeNetworkAclsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the specified attribute of the specified network interface. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaceAttribute.html]
type DescribeNetworkInterfaceAttributeRequest struct {
	NetworkInterfaceId string            `name:"NetworkInterfaceId"`
	Attribute          InstanceAttribute `name:"Attribute"`
}

// Creates a new DescribeNetworkInterfaceAttributeRequest.
func NewDescribeNetworkInterfaceAttributeRequest(networkInterfaceId string, attribute InstanceAttribute) *DescribeNetworkInterfaceAttributeRequest {
	return &DescribeNetworkInterfaceAttributeRequest{networkInterfaceId, attribute}
}

/*****************************************************************************/

// Describes one or more of your network interfaces.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaces.html]
type DescribeNetworkInterfacesRequest struct {
	NetworkInterfaceIds []string         `name:"NetworkInterfaceId.#,omitempt" base:"1"`
	Filters             []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeNetworkInterfacesRequest.
func NewDescribeNetworkInterfacesRequest() *DescribeNetworkInterfacesRequest {
	return &DescribeNetworkInterfacesRequest{}
}

// Add a NetworkInterfaceId.
func (r *DescribeNetworkInterfacesRequest) AddNetworkInterfaceId(networkInterfaceId string) {
	r.NetworkInterfaceIds = append(r.NetworkInterfaceIds, networkInterfaceId)
}

// Add a filter.
func (r *DescribeNetworkInterfacesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your placement groups.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribePlacementGroups.html]
type DescribePlacementGroupsRequest struct {
	GroupNames []string         `name:"GroupName.#,omitempty" base:"1"`
	Filters    []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribePlacementGroupsRequest.
func NewDescribePlacementGroupsRequest() *DescribePlacementGroupsRequest {
	return &DescribePlacementGroupsRequest{}
}

// Add a GroupName.
func (r *DescribePlacementGroupsRequest) AddGroupName(groupName string) {
	r.GroupNames = append(r.GroupNames, groupName)
}

// Add a filter.
func (r *DescribePlacementGroupsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more regions that are currently available to you.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRegions.html]
type DescribeRegionsRequest struct {
	RegionNames []string         `name:"RegionName.#,omitempty" base:"1"`
	Filters     []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeRegionsRequest.
func NewDescribeRegionsRequest() *DescribeRegionsRequest {
	return &DescribeRegionsRequest{}
}

// Add a RegionName.
func (r *DescribeRegionsRequest) AddRegionName(regionName string) {
	r.RegionNames = append(r.RegionNames, regionName)
}

// Add a filter.
func (r *DescribeRegionsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of the Reserved Instances that you purchased.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstances.html]
type DescribeReservedInstancesRequest struct {
	ReservedInstancesIds []string         `name:"ReservedInstancesId.#,omitempty" base:"1"`
	OfferingType         OfferingType     `name:"OfferingType,omitempty"`
	Filters              []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeReservedInstancesRequest.
func NewDescribeReservedInstancesRequest() *DescribeReservedInstancesRequest {
	return &DescribeReservedInstancesRequest{}
}

// Add a ReservedInstancesId.
func (r *DescribeReservedInstancesRequest) AddReservedInstancesId(reservedInstancesId string) {
	r.ReservedInstancesIds = append(r.ReservedInstancesIds, reservedInstancesId)
}

// Add a filter.
func (r *DescribeReservedInstancesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes your account's Reserved Instance listings in the Reserved Instance Marketplace.
// This call returns information, such as the ID of the Reserved Instance with which a
// listing is associated.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesListings.html]
type DescribeReservedInstancesListingsRequest struct {
	ReservedInstancesListingIds []string         `name:"ReservedInstancesListingId.#,omitempty" base:"1"`
	ReservedInstancesIds        []string         `name:"ReservedInstancesId.#,omitempty" base:"1"`
	Filters                     []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeReservedInstancesListingsRequest.
func NewDescribeReservedInstancesListingsRequest() *DescribeReservedInstancesListingsRequest {
	return &DescribeReservedInstancesListingsRequest{}
}

// Add a ReservedInstancesId.
func (r *DescribeReservedInstancesListingsRequest) AddReservedInstancesId(reservedInstancesId string) {
	r.ReservedInstancesIds = append(r.ReservedInstancesIds, reservedInstancesId)
}

// Add a ReservedInstancesListingId.
func (r *DescribeReservedInstancesListingsRequest) AddReservedInstancesListingId(reservedInstancesListingId string) {
	r.ReservedInstancesListingIds = append(r.ReservedInstancesListingIds, reservedInstancesListingId)
}

// Add a filter.
func (r *DescribeReservedInstancesListingsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the modifications made to your Reserved Instances. If no parameter is specified,
// information about all your Reserved Instances modification requests is returned. If a
// modification ID is specified, only information about the specific modification is returned.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesModifications.html]
type DescribeReservedInstancesModificationsRequest struct {
	ReservedInstancesModificationIds []string         `name:"ReservedInstancesModificationId.#,omitempty" base:"1"`
	NextToken                        string           `name:"NextToken,omitempty"`
	Filters                          []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeReservedInstancesModificationsRequest.
func NewDescribeReservedInstancesModificationsRequest() *DescribeReservedInstancesModificationsRequest {
	return &DescribeReservedInstancesModificationsRequest{}
}

// Add a ReservedInstancesModificationId.
func (r *DescribeReservedInstancesModificationsRequest) AddReservedInstancesModificationId(reservedInstancesModificationId string) {
	r.ReservedInstancesModificationIds = append(r.ReservedInstancesModificationIds, reservedInstancesModificationId)
}

// Add a filter.
func (r *DescribeReservedInstancesModificationsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes Reserved Instance offerings that are available for purchase. With Amazon EC2 Reserved
// Instances, you purchase the right to launch Amazon EC2 instances for a period of time. During
// that time period, you do not receive insufficient capacity errors, and you pay a lower usage
// rate than the rate charged for On-Demand instances for the actual time used.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesOfferings.html]
type DescribeReservedInstancesOfferingsRequest struct {
	ReservedInstancesOfferingIds []string           `name:"ReservedInstancesOfferingId.#,omitempty" base:"1"`
	InstanceType                 InstanceType       `name:"InstanceType,omitempty"`
	AvailabilityZone             string             `name:"AvailabilityZone,omitempty"`
	ProductDescription           ProductDescription `name:"ProductDescription,omitempty"`
	Filters                      []FilterItemType   `name:"Filter.#.,omitempty"`
	InstanceTenancy              InstanceTenancy    `name:"InstanceTenancy,omitempty"`
	OfferingType                 OfferingType       `name:"OfferingType,omitempty"`
	IncludeMarketplace           bool               `name:"IncludeMarketplace,omitempty"`
	MinDuration                  int64              `name:"MinDuration,omitempty"`
	MaxDuration                  int64              `name:"MaxDuration,omitempty"`
	MaxInstanceCount             int                `name:"MaxInstanceCount,omitempty"`
	NextToken                    string             `name:"NextToken,omitempty"`
	MaxResults                   int                `name:"MaxResults,omitempty"`
}

// Creates a new DescribeReservedInstancesOfferingsRequest.
func NewDescribeReservedInstancesOfferingsRequest() *DescribeReservedInstancesOfferingsRequest {
	return &DescribeReservedInstancesOfferingsRequest{}
}

// Add a ReservedInstancesOfferingId.
func (r *DescribeReservedInstancesOfferingsRequest) AddReservedInstancesOfferingId(reservedInstancesOfferingId string) {
	r.ReservedInstancesOfferingIds = append(r.ReservedInstancesOfferingIds, reservedInstancesOfferingId)
}

// Add a filter.
func (r *DescribeReservedInstancesOfferingsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your route tables.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRouteTables.html]
type DescribeRouteTablesRequest struct {
	RouteTableIds []string         `name:"RouteTableId.#,omitempty" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty"`
}

// Creates a new DescribeRouteTablesRequest.
func NewDescribeRouteTablesRequest() *DescribeRouteTablesRequest {
	return &DescribeRouteTablesRequest{}
}

// Add a RouteTableId.
func (r *DescribeRouteTablesRequest) AddRouteTableId(routeTableId string) {
	r.RouteTableIds = append(r.RouteTableIds, routeTableId)
}

// Add a filter.
func (r *DescribeRouteTablesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your security groups.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSecurityGroups.html]
type DescribeSecurityGroupsRequest struct {
	GroupNames []string         `name:"GroupName.#,omitempty" base:"1"`
	GroupIds   []string         `name:"GroupId.#,omitempty" base:"1"`
	Filters    []FilterItemType `name:"Filter.#.,omitempty"`
}

// Creates a new DescribeSecurityGroupsRequest.
func NewDescribeSecurityGroupsRequest() *DescribeSecurityGroupsRequest {
	return &DescribeSecurityGroupsRequest{}
}

// Add a GroupId.
func (r *DescribeSecurityGroupsRequest) AddGroupId(groupId string) {
	r.GroupIds = append(r.GroupIds, groupId)
}

// Add a GroupName.
func (r *DescribeSecurityGroupsRequest) AddGroupName(groupName string) {
	r.GroupNames = append(r.GroupNames, groupName)
}

// Add a filter.
func (r *DescribeSecurityGroupsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the specified attribute of the specified snapshot. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshotAttribute.html]
type DescribeSnapshotAttributeRequest struct {
	SnapshotId string            `name:"SnapshotId"`
	Attribute  InstanceAttribute `name:"Attribute"`
}

// Creates a new Request.
func NewDescribeSnapshotAttributeRequest(snapshotId string, attribute InstanceAttribute) *DescribeSnapshotAttributeRequest {
	return &DescribeSnapshotAttributeRequest{snapshotId, attribute}
}

/*****************************************************************************/

// Describes one or more of the Amazon EBS snapshots available to you. Snapshots available
// to you include public snapshots available for any AWS account to launch, private snapshots
// you own, and private snapshots owned by another AWS account but for which you've been
// given explicit create volume permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html]
type DescribeSnapshotsRequest struct {
	SnapshotIds   []string         `name:"SnapshotId.#" base:"1"`
	Owners        []string         `name:"Owner.#" base:"1"`
	RestorableBys []string         `name:"RestorableBy.#" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeSnapshotsRequest.
func NewDescribeSnapshotsRequest() *DescribeSnapshotsRequest {
	return &DescribeSnapshotsRequest{}
}

// Add a SnapshotId.
func (r *DescribeSnapshotsRequest) AddSnapshotId(snapshotId string) {
	r.SnapshotIds = append(r.SnapshotIds, snapshotId)
}

// Add a Owner.
func (r *DescribeSnapshotsRequest) AddOwner(owner string) {
	r.Owners = append(r.Owners, owner)
}

// Add a RestorableBy.
func (r *DescribeSnapshotsRequest) AddCustomerGatewayId(restorableBy string) {
	r.RestorableBys = append(r.RestorableBys, restorableBy)
}

// Add a filter.
func (r *DescribeSnapshotsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the datafeed for Spot Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotDatafeedSubscription.html]
type DescribeSpotDatafeedSubscriptionRequest struct{}

/*****************************************************************************/

// Describes the Spot Instance requests that belong to your account. Spot Instances are instances
// that Amazon EC2 starts on your behalf when the maximum price that you specify exceeds the
// current Spot Price. Amazon EC2 periodically sets the Spot Price based on available Spot
// Instance capacity and current Spot Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotInstanceRequests.html]
type DescribeSpotInstanceRequestsRequest struct {
	SpotInstanceRequestIds []string         `name:"SpotInstanceRequestId.#,omitempty"`
	Filters                []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeSpotInstanceRequestsRequest.
func NewDescribeSpotInstanceRequestsRequest() *DescribeSpotInstanceRequestsRequest {
	return &DescribeSpotInstanceRequestsRequest{}
}

// Add a SpotInstanceRequestId.
func (r *DescribeSpotInstanceRequestsRequest) AddSpotInstanceRequestId(spotInstanceRequestId string) {
	r.SpotInstanceRequestIds = append(r.SpotInstanceRequestIds, spotInstanceRequestId)
}

// Add a filter.
func (r *DescribeSpotInstanceRequestsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the Spot Price history. Spot Instances are instances that Amazon EC2 starts on your
// behalf when the maximum price that you specify exceeds the current Spot Price. Amazon EC2
// periodically sets the Spot Price based on available Spot Instance capacity and current Spot
// Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotPriceHistory.html]
type DescribeSpotPriceHistoryRequest struct {
	StartTime           time.Time            `name:"StartTime,omitempty" format:"2006-01-02T15:04:05.999Z"`
	EndTime             time.Time            `name:"EndTime,omitempty" format:"2006-01-02T15:04:05.999Z"`
	InstanceTypes       []string             `name:"InstanceType.#,omitempty" base:"1"`
	ProductDescriptions []ProductDescription `name:"ProductDescription.#,omitempty" base:"1"`
	Filters             []FilterItemType     `name:"Filter.#.,omitempty" base:"1"`
	AvailabilityZone    string               `name:"AvailabilityZone,omitempty"`
	MaxResults          int                  `name:"MaxResults,omitempty"`
	NextToken           string               `name:"NextToken,omitempty"`
}

// Creates a new DescribeSpotPriceHistoryRequest.
func NewDescribeSpotPriceHistoryRequest() *DescribeSpotPriceHistoryRequest {
	return &DescribeSpotPriceHistoryRequest{}
}

// Add a InstanceType.
func (r *DescribeSpotPriceHistoryRequest) AddInstanceType(instanceType string) {
	r.InstanceTypes = append(r.InstanceTypes, instanceType)
}

// Add a ProductDescription.
func (r *DescribeSpotPriceHistoryRequest) AddProductDescription(productDescription ProductDescription) {
	r.ProductDescriptions = append(r.ProductDescriptions, productDescription)
}

// Add a filter.
func (r *DescribeSpotPriceHistoryRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your subnets.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSubnets.html]
type DescribeSubnetsRequest struct {
	SubnetIds []string         `name:"SubnetId.#,omitempty"`
	Filters   []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeSubnetsRequest.
func NewDescribeSubnetsRequest() *DescribeSubnetsRequest {
	return &DescribeSubnetsRequest{}
}

// Add a SubnetId.
func (r *DescribeSubnetsRequest) AddSubnetId(subnetId string) {
	r.SubnetIds = append(r.SubnetIds, subnetId)
}

// Add a filter.
func (r *DescribeSubnetsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of the tags for your EC2 resources.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeTags.html]
type DescribeTagsRequest struct {
	MaxResults int              `name:"MaxResults,omitempty"`
	NextToken  string           `name:"NextToken,omitempty"`
	Filters    []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeTagsRequest.
func NewDescribeTagsRequest() *DescribeTagsRequest {
	return &DescribeTagsRequest{}
}

// Add a filter.
func (r *DescribeTagsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the specified attribute of the specified volume. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeAttribute.html]
type DescribeVolumeAttributeRequest struct {
	VolumeId  string            `name:"VolumeId"`
	Attribute InstanceAttribute `name:"Attribute"`
}

// Creates a new DescribeVolumeAttributeRequest.
func NewDescribeVolumeAttributeRequest(volumeId string, attribute InstanceAttribute) *DescribeVolumeAttributeRequest {
	return &DescribeVolumeAttributeRequest{volumeId, attribute}
}

/*****************************************************************************/

// Describes the specified Amazon EBS volumes.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumes.html]
type DescribeVolumesRequest struct {
	VolumeIds []string         `name:"VolumeId,VolumeId.#" base:"1"`
	Filters   []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVolumesRequest.
func NewDescribeVolumesRequest() *DescribeVolumesRequest {
	return &DescribeVolumesRequest{}
}

// Add a VolumeId.
func (r *DescribeVolumesRequest) AddVolumeId(volumeId string) {
	r.VolumeIds = append(r.VolumeIds, volumeId)
}

// Add a filter.
func (r *DescribeVolumesRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the status of the specified volumes. Volume status provides the result of the
// checks performed on your volumes to determine events that can impair the performance of
// your volumes. The performance of a volume can be affected if an issue occurs on the volume's
// underlying host. If the volume's underlying host experiences a power outage or system issue,
// once the system is restored there could be data inconsistencies on the volume. Volume events
// notify you if this occurs. Volume actions notify you if any action needs to be taken in
// response to the event.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeStatus.html]
type DescribeVolumeStatusRequest struct {
	VolumeIds []string         `name:"VolumeId,VolumeId.#" base:"1"`
	Filters   []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVolumeStatusRequest.
func NewDescribeVolumeStatusRequest() *DescribeVolumeStatusRequest {
	return &DescribeVolumeStatusRequest{}
}

// Add a VolumeId.
func (r *DescribeVolumeStatusRequest) AddVolumeId(volumeId string) {
	r.VolumeIds = append(r.VolumeIds, volumeId)
}

// Add a filter.
func (r *DescribeVolumeStatusRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes the specified attribute of the specified VPC. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcAttribute.html]
type DescribeVpcAttributeRequest struct {
	RequestId string       `name:"RequestId"`
	Attribute VpcAttribute `name:"Attribute"`
}

// Creates a new DescribeVpcAttributeRequest.
func NewDescribeVpcAttributeRequest() *DescribeVpcAttributeRequest {
	return &DescribeVpcAttributeRequest{}
}

/*****************************************************************************/

// Describes one or more of your VPC peering connections.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcPeeringConnections.html]
type DescribeVpcPeeringConnectionsRequest struct {
	VpcPeeringConnectionId string           `name:"VpcPeeringConnectionId,omitempty"`
	Filters                []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVpcPeeringConnectionsRequest.
func NewDescribeVpcPeeringConnectionsRequest() *DescribeVpcPeeringConnectionsRequest {
	return &DescribeVpcPeeringConnectionsRequest{}
}

// Add a filter.
func (r *DescribeVpcPeeringConnectionsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your VPCs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcs.html]
type DescribeVpcsRequest struct {
	VpcIds  []string         `name:"VpcId.#" base:"1"`
	Filters []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVpcsRequest.
func NewDescribeVpcsRequest() *DescribeVpcsRequest {
	return &DescribeVpcsRequest{}
}

// Add a VpcId.
func (r *DescribeVpcsRequest) AddVpcId(vpcId string) {
	r.VpcIds = append(r.VpcIds, vpcId)
}

// Add a filter.
func (r *DescribeVpcsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your VPN connections.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnConnections.html]
type DescribeVpnConnectionsRequest struct {
	VpnConnectionIds []string         `name:"VpnConnectionId.#" base:"1"`
	Filters          []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVpnConnectionsRequest.
func NewDescribeVpnConnectionsRequest() *DescribeVpnConnectionsRequest {
	return &DescribeVpnConnectionsRequest{}
}

// Add a VpnConnectionId.
func (r *DescribeVpnConnectionsRequest) AddVpnConnectionId(vpnConnectionId string) {
	r.VpnConnectionIds = append(r.VpnConnectionIds, vpnConnectionId)
}

// Add a filter.
func (r *DescribeVpnConnectionsRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Describes one or more of your virtual private gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnGateways.html]
type DescribeVpnGatewaysRequest struct {
	VpnGatewayIds []string         `name:"VpnGatewayId.#,omitempty" base:"1"`
	Filters       []FilterItemType `name:"Filter.#.,omitempty" base:"1"`
}

// Creates a new DescribeVpnGatewaysRequest.
func NewDescribeVpnGatewaysRequest() *DescribeVpnGatewaysRequest {
	return &DescribeVpnGatewaysRequest{}
}

// Add a VpnGatewayId.
func (r *DescribeVpnGatewaysRequest) AddVpnGatewayId(vpnGatewayId string) {
	r.VpnGatewayIds = append(r.VpnGatewayIds, vpnGatewayId)
}

// Add a filter.
func (r *DescribeVpnGatewaysRequest) AddFilter(name string, values ...string) {
	r.Filters = append(r.Filters, FilterItemType{Name: name, ValueSet: values})
}

/*****************************************************************************/

// Detaches an Internet gateway from a VPC, disabling connectivity between the Internet and the VPC.
// The VPC must not contain any running instances with Elastic IP addresses.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachInternetGateway.html]
type DetachInternetGatewayRequest struct {
	InternetGatewayId string `name:"InternetGatewayId"`
	VpcId             string `name:"VpcId"`
}

// Creates a new DetachInternetGatewayRequest.
func NewDetachInternetGatewayRequest(vpcId, internetGatewayId string) *DetachInternetGatewayRequest {
	return &DetachInternetGatewayRequest{internetGatewayId, vpcId}
}

/*****************************************************************************/

// Detaches a network interface from an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachNetworkInterface.html]
type DetachNetworkInterfaceRequest struct {
	AttachmentId string `name:"AttachmentId"`
	Force        bool   `name:"Force,omitempty"`
}

// Creates a new DetachNetworkInterfaceRequest.
func NewDetachNetworkInterfaceRequest(attachmentId string) *DetachNetworkInterfaceRequest {
	return &DetachNetworkInterfaceRequest{attachmentId, false}
}

/*****************************************************************************/

// Detaches an Amazon EBS volume from an instance. Make sure to unmount any file systems
// on the device within your operating system before detaching the volume. Failure to do
// so results in the volume being stuck in a busy state while detaching.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVolume.html]
type DetachVolumeRequest struct {
	VolumeId   string `name:"VolumeId"`
	InstanceId string `name:"InstanceId,omitempty"`
	Device     string `name:"Device,omitempty"`
	Force      bool   `name:"Force,omitempty"`
}

// Creates a new DetachVolumeRequest.
func NewDetachVolumeRequest(volumeId string) *DetachVolumeRequest {
	return &DetachVolumeRequest{VolumeId: volumeId}
}

/*****************************************************************************/

// Detaches a virtual private gateway from a VPC. You do this if you're planning to turn off
// the VPC and not use it anymore. You can confirm a virtual private gateway has been
// completely detached from a VPC by describing the virtual private gateway (any attachments
// to the virtual private gateway are also described).
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVpnGateway.html]
type DetachVpnGatewayRequest struct {
	VpnGatewayId string `name:"VpnGatewayId"`
	VpcId        string `name:"VpcId"`
}

// Creates a new DetachVpnGatewayRequest.
func NewDetachVpnGatewayRequest(vpcId, vpnGatewayId string) *DetachVpnGatewayRequest {
	return &DetachVpnGatewayRequest{vpnGatewayId, vpcId}
}

/*****************************************************************************/

// Disables a virtual private gateway (VGW) from propagating routes to the routing tables of a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisableVgwRoutePropagation.html]
type DisableVgwRoutePropagationRequest struct {
	RouteTableId string `name:"RouteTableId"`
	GatewayId    string `name:"GatewayId"`
}

// Creates a new DisableVgwRoutePropagationRequest.
func NewDisableVgwRoutePropagationRequest(gatewayId, routeTableId string) *DisableVgwRoutePropagationRequest {
	return &DisableVgwRoutePropagationRequest{routeTableId, gatewayId}
}

/*****************************************************************************/

// Disassociates an Elastic IP address from the instance or network interface it's associated with.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateAddress.html]
type DisassociateAddressRequest struct {
	PublicIp      string `name:"PublicIp,omitempty"`
	AssociationId string `name:"AssociationId,omitempty"`
}

// Creates a new DisassociateAddressRequest.
func NewDisassociateAddressRequest() *DisassociateAddressRequest {
	return &DisassociateAddressRequest{}
}

/*****************************************************************************/

// Disassociates a subnet from a route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateRouteTable.html]
type DisassociateRouteTableRequest struct {
	AssociationId string `name:"AssociationId"`
}

// Creates a new DisassociateRouteTableRequest.
func NewDisassociateRouteTableRequest(associationId string) *DisassociateRouteTableRequest {
	return &DisassociateRouteTableRequest{associationId}
}

/*****************************************************************************/

// Enables a virtual private gateway (VGW) to propagate routes to the routing tables of a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVgwRoutePropagation.html]
type EnableVgwRoutePropagationRequest struct {
	RouteTableId string `name:"RouteTableId"`
	GatewayId    string `name:"GatewayId"`
}

// Creates a new EnableVgwRoutePropagationRequest.
func NewEnableVgwRoutePropagationRequest(gatewayId, routeTableId string) *EnableVgwRoutePropagationRequest {
	return &EnableVgwRoutePropagationRequest{routeTableId, gatewayId}
}

/*****************************************************************************/

// Enables I/O operations for a volume that had I/O operations disabled because
// the data on the volume was potentially inconsistent.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVolumeIO.html]
type EnableVolumeIORequest struct {
	VolumeId string `name:"VolumeId"`
}

// Creates a new EnableVolumeIORequest.
func NewEnableVolumeIORequest(volumeId string) *EnableVolumeIORequest {
	return &EnableVolumeIORequest{volumeId}
}

/*****************************************************************************/

// Gets the console output for the specified instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetConsoleOutput.html]
type GetConsoleOutputRequest struct {
	InstanceId string `name:"InstanceId"`
}

// Creates a new GetConsoleOutputRequest.
func NewGetConsoleOutputRequest(instanceId string) *GetConsoleOutputRequest {
	return &GetConsoleOutputRequest{instanceId}
}

/*****************************************************************************/

// Retrieves the encrypted administrator password for an instance running Windows.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetPasswordData.html]
type GetPasswordDataRequest struct {
	InstanceId string `name:"InstanceId"`
}

// Creates a new GetPasswordDataRequest.
func NewGetPasswordDataRequest(instanceId string) *GetPasswordDataRequest {
	return &GetPasswordDataRequest{instanceId}
}

/*****************************************************************************/

// Creates an import instance task using metadata from the specified disk image.
// After importing the image, you then upload it using the ec2-import-volume command.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportInstance.html]
type ImportInstanceRequest struct {
	Description         string              `name:"Description,omitempty"`
	LaunchSpecification LaunchSpecification `name:"LaunchSpecification."`
	DiskImage           []DiskImage         `name:"DiskImage.#."`
	Platform            string              `name:"Platform,omitempty"`
}

type LaunchSpecification struct {
	Architecture                      Architecture `name:"Architecture"`
	GroupNames                        []string     `name:"GroupName.#,omitempty" base:"1"`
	UserData                          string       `name:"UserData,omitempty"`
	InstanceType                      InstanceType `name:"InstanceType"`
	PlacementAvailabilityZone         string       `name:"Placement.AvailabilityZone,omitempty"`
	MonitoringEnabled                 bool         `name:"Monitoring.Enabled,omitempty"`
	SubnetId                          string       `name:"SubnetId,omitempty"`
	InstanceInitiatedShutdownBehavior string       `name:"InstanceInitiatedShutdownBehavior,omitempty"`
	PrivateIpAddress                  string       `name:"PrivateIpAddress,omitempty"`
}

type DiskImage struct {
	Format            string `name:"Image.Format"`
	Bytes             int64  `name:"Image.Bytes"`
	ImportManifestUrl string `name:"Image.ImportManifestUrl"`
	Description       string `name:"Image.Description,omitempty"`
	VolumSize         int    `name:"Volume.Size"`
}

// Creates a new ImportInstanceRequest.
func NewImportInstanceRequest(architecture Architecture, format, importManifestUrl string, bytes int64, volumeSize int) *ImportInstanceRequest {
	return &ImportInstanceRequest{
		LaunchSpecification: LaunchSpecification{
			Architecture: architecture,
		},
		DiskImage: []DiskImage{DiskImage{
			Format:            format,
			Bytes:             bytes,
			ImportManifestUrl: importManifestUrl,
			VolumSize:         volumeSize,
		}},
	}
}

/*****************************************************************************/

// Imports the public key from an RSA key pair that you created with a third-party tool.
// Compare this with CreateKeyPair, in which AWS creates the key pair and gives the keys
// to you (AWS keeps a copy of the public key). With ImportKeyPair, you create the key
// pair and give AWS just the public key. The private key is never transferred between
// you and AWS.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportKeyPair.html]
type ImportKeyPairRequest struct {
	KeyName           string `name:"KeyName"`
	PublicKeyMaterial string `name:"PublicKeyMaterial"`
}

// Creates a new ImportKeyPairRequest.
func NewImportKeyPairRequest(keyName, publicKeyMaterial string) *ImportKeyPairRequest {
	return &ImportKeyPairRequest{keyName, publicKeyMaterial}
}

/*****************************************************************************/

// Creates an import volume task using metadata from the specified disk image. After importing
// the image, you then upload it using the ec2-import-volume command in the Amazon EC2
// command-line interface (CLI) tools.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportVolume.html]
type ImportVolumeRequest struct {
	AvailabilityZone string    `name:"AvailabilityZone"`
	Image            DiskImage `name:"Image."`
	Description      string    `name:"Description,omitempty"`
	VolumeSize       int       `name:"Volume.Size"`
}

// Creates a new ImportVolumeRequest.
func NewImportVolumeRequest(availabilityZone, format, importManifestUrl string, bytes int64, volumeSize int) *ImportVolumeRequest {
	return &ImportVolumeRequest{
		AvailabilityZone: availabilityZone,
		VolumeSize:       volumeSize,
		Image: DiskImage{
			Format:            format,
			Bytes:             bytes,
			ImportManifestUrl: importManifestUrl,
		},
	}
}

/*****************************************************************************/

// Modifies the specified attribute of the specified AMI. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyImageAttribute.html]
type ModifyImageAttributeRequest struct {
	ImageId          string            `name:"ImageId"`
	LaunchPermission *LaunchPermission `name:"LaunchPermission.,omitempty"`
	ProductCode      []string          `name:"ProductCode.#,omitempty" base:"1"`
	Description      string            `name:"Description.Value,omitempty"`
}

type UserIdGroupPair struct {
	UserId string `name:"UserId,omitempty"`
	Group  string `name:"Group,omitempty"`
}

type LaunchPermission struct {
	Add    []UserIdGroupPair `name:"Add.#.,omitempty" base:"1"`
	Remove []UserIdGroupPair `name:"Remove.#.,omitempty" base:"1"`
}

// Creates a new ModifyImageAttributeRequest.
func NewModifyImageAttributeRequest(imageId string) *ModifyImageAttributeRequest {
	return &ModifyImageAttributeRequest{ImageId: imageId}
}

/*****************************************************************************/

// Modifies the specified attribute of the specified instance. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyInstanceAttribute.html]
type ModifyInstanceAttributeRequest struct {
	InstanceId                        string                              `name:"InstanceId"`
	BlockDeviceMapping                *InstanceBlockDeviceMappingItemType `name:"BlockDeviceMapping.Value.,omitempty"`
	DisableApiTermination             string                              `name:"DisableApiTermination.Value,omitempty"`
	EbsOptimized                      bool                                `name:"EbsOptimized,omitempty"`
	GroupId                           []string                            `name:"GroupId.#,omitempty" base:"1"`
	InstanceInitiatedShutdownBehavior string                              `name:"InstanceInitiatedShutdownBehavior.Value,omitempty"`
	InstanceType                      InstanceType                        `name:"InstanceType.Value,omitempty"`
	Kernel                            string                              `name:"Kernel.Value,omitempty"`
	Ramdisk                           string                              `name:"Ramdisk.Value,omitempty"`
	SourceDestCheck                   bool                                `name:"SourceDestCheck.Value,omitempty"`
	SriovNetSupport                   string                              `name:"SriovNetSupport.Value,omitempty"`
	UserData                          string                              `name:"UserData.Value,omitempty"`
}

// Creates a new ModifyInstanceAttributeRequest.
func NewModifyInstanceAttributeRequest(instanceId string) *ModifyInstanceAttributeRequest {
	return &ModifyInstanceAttributeRequest{InstanceId: instanceId}
}

/*****************************************************************************/

// Modifies the specified network interface attribute. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyNetworkInterfaceAttribute.html]
type ModifyNetworkInterfaceAttributeRequest struct {
	NetworkInterfaceId string                                     `name:"NetworkInterfaceId"`
	Description        string                                     `name:"Description.Value,omitempty"`
	SecurityGroupIds   []string                                   `name:"SecurityGroupId.#,omitempty" base:"1"`
	SourceDestCheck    bool                                       `name:"SourceDestCheck.Value,omitempty"`
	Attachment         *ModifyNetworkInterfaceAttributeAttachment `name:"Attachment.,omitempty"`
}

type ModifyNetworkInterfaceAttributeAttachment struct {
	AttachmentId        string `name:"AttachmentId,omitempty"`
	DeleteOnTermination bool   `name:"DeleteOnTermination,omitempty"`
}

// Creates a new ModifyNetworkInterfaceAttributeRequest.
func NewModifyNetworkInterfaceAttributeRequest(networkInterfaceId string) *ModifyNetworkInterfaceAttributeRequest {
	return &ModifyNetworkInterfaceAttributeRequest{NetworkInterfaceId: networkInterfaceId}
}

/*****************************************************************************/

// Modifies the Availability Zone, instance count, instance type, or network platform
// (EC2-Classic or EC2-VPC) of your Reserved Instances. The Reserved Instances to be
// modified must be identical, except for Availability Zone, network platform,
// and instance type.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyReservedInstances.html]
type ModifyReservedInstancesRequest struct {
	ReservedInstancesID string                                    `name:"ReservedInstancesID"`
	ClientToken         string                                    `name:"ClientToken,omitempty"`
	TargetConfiguration ReservedInstancesConfigurationSetItemType `name:"TargetConfiguration"`
}

// Creates a new ModifyReservedInstancesRequest.
func NewModifyReservedInstancesRequest(reservedInstancesID string, targetConfiguration ReservedInstancesConfigurationSetItemType) *ModifyReservedInstancesRequest {
	return &ModifyReservedInstancesRequest{
		ReservedInstancesID: reservedInstancesID,
		TargetConfiguration: targetConfiguration,
	}
}

/*****************************************************************************/

// Adds or removes permission settings for the specified snapshot. You may add or remove specified
// AWS account IDs from a snapshot's list of create volume permissions, but you cannot do both in
// a single API call. If you need to both add and remove account IDs for a snapshot, you must use
// multiple API calls.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySnapshotAttribute.html]
type ModifySnapshotAttributeRequest struct {
	SnapshotId             string
	CreateVolumePermission *CreateVolumePermission `name:"CreateVolumePermission.,omitempty"`
}

type CreateVolumePermission struct {
	Add    []UserIdGroupPair `name:"Add.#.,omitempty" base:"1"`
	Remove []UserIdGroupPair `name:"Remove.#.,omitempty" base:"1"`
}

// Creates a new ModifySnapshotAttributeRequest.
func NewModifySnapshotAttributeRequest(snapshotId string) *ModifySnapshotAttributeRequest {
	return &ModifySnapshotAttributeRequest{SnapshotId: snapshotId}
}

/*****************************************************************************/

// Modifies a subnet attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySubnetAttribute.html]
type ModifySubnetAttributeRequest struct {
	SubnetId            string `name:"SubnetId"`
	MapPublicIpOnLaunch bool   `name:"MapPublicIpOnLaunch.Value"`
}

// Creates a new ModifySubnetAttributeRequest.
func NewModifySubnetAttributeRequest(subnetId string, mapPublicIpOnLaunch bool) *ModifySubnetAttributeRequest {
	return &ModifySubnetAttributeRequest{subnetId, mapPublicIpOnLaunch}
}

/*****************************************************************************/

// Modifies a volume attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVolumeAttribute.html]
type ModifyVolumeAttributeRequest struct {
	VolumeId     string `name:"VolumeId"`
	AutoEnableIO bool   `name:"AutoEnableIO.Value"`
}

// Creates a new ModifyVolumeAttributeRequest.
func NewModifyVolumeAttributeRequest(volumeId string, autoEnableIO bool) *ModifyVolumeAttributeRequest {
	return &ModifyVolumeAttributeRequest{volumeId, autoEnableIO}
}

/*****************************************************************************/

// Modifies the specified attribute of the specified VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVpcAttribute.html]
type ModifyVpcAttributeRequest struct {
	VpcId              string `name:"VpcId"`
	EnableDnsSupport   bool   `name:"EnableDnsSupport,omitempty"`
	EnableDnsHostnames bool   `name:"EnableDnsHostnames,omitempty"`
}

// Creates a new ModifyVpcAttributeRequest.
func NewModifyVpcAttributeRequest(vpcId string) *ModifyVpcAttributeRequest {
	return &ModifyVpcAttributeRequest{VpcId: vpcId}
}

/*****************************************************************************/

// Enables monitoring for a running instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-MonitorInstances.html]
type MonitorInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
}

// Creates a new MonitorInstancesRequest.
func NewMonitorInstancesRequest(instanceIds ...string) *MonitorInstancesRequest {
	return &MonitorInstancesRequest{instanceIds}
}

/*****************************************************************************/

// Purchases a Reserved Instance for use with your account. With Amazon EC2 Reserved Instances,
// you obtain a capacity reservation for a certain instance configuration over a specified
// period of time. You pay a lower usage rate than with On-Demand instances for the time that
// you actually use the capacity reservation.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-PurchaseReservedInstancesOffering.html]
type PurchaseReservedInstancesOfferingRequest struct {
	ReservedInstancesOfferingId string                          `name:"ReservedInstancesOfferingId"`
	InstanceCount               int                             `name:"InstanceCount"`
	LimitPrice                  *ReservedInstanceLimitPriceType `name:"LimitPrice.,omitempty"`
}

// Creates a new PurchaseReservedInstancesOfferingRequest.
func NewPurchaseReservedInstancesOfferingRequest(reservedInstancesOfferingId string, instanceCount int) *PurchaseReservedInstancesOfferingRequest {
	return &PurchaseReservedInstancesOfferingRequest{
		ReservedInstancesOfferingId: reservedInstancesOfferingId,
		InstanceCount:               instanceCount,
	}
}

/*****************************************************************************/

// Requests a reboot of one or more instances. This operation is asynchronous; it only queues a
// request to reboot the specified instances. The operation succeeds if the instances are valid
// and belong to you. Requests to reboot terminated instances are ignored.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RebootInstances.html]
type RebootInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
}

// Creates a new RebootInstancesRequest.
func NewRebootInstancesRequest(instanceIds ...string) *RebootInstancesRequest {
	return &RebootInstancesRequest{instanceIds}
}

/*****************************************************************************/

// Registers an AMI. When you're creating an AMI, this is the final step you must
// complete before you can launch an instance from the AMI.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RegisterImage.html]
type RegisterImageRequest struct {
	ImageLocation      string                       `name:"ImageLocation,omitempty"`
	Name               string                       `name:"Name"`
	Description        string                       `name:"Description,omitempty"`
	Architecture       Architecture                 `name:"Architecture,omitempty"`
	RootDeviceName     string                       `name:"RootDeviceName,omitempty"`
	BlockDeviceMapping []BlockDeviceMappingItemType `name:"BlockDeviceMapping.#.,omitempty" base:"1"`
	VirtualizationType VirtualizationType           `name:"VirtualizationType,omitempty"`
	KernelId           string                       `name:"KernelId,omitempty"`
	RamdiskId          string                       `name:"RamdiskId,omitempty"`
	SriovNetSupport    string                       `name:"SriovNetSupport,omitempty"`
}

// Creates a new RegisterImageRequest.
func NewRegisterImageRequest(name string) *RegisterImageRequest {
	return &RegisterImageRequest{Name: name}
}

/*****************************************************************************/

// Rejects a VPC peering connection request. The VPC peering connection must be
// in the pending-acceptance state.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RejectVpcPeeringConnection.html]
type RejectVpcPeeringConnectionRequest struct {
	VpcPeeringConnectionId string `name:"VpcPeeringConnectionId"`
}

// Creates a new RejectVpcPeeringConnectionRequest.
func NewRejectVpcPeeringConnectionRequest(vpcPeeringConnectionId string) *RejectVpcPeeringConnectionRequest {
	return &RejectVpcPeeringConnectionRequest{vpcPeeringConnectionId}
}

/*****************************************************************************/

// Releases the specified Elastic IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReleaseAddress.html]
type ReleaseAddressRequest struct {
	PublicIp     string `name:"PublicIp"`
	AllocationId string `name:"AllocationId"`
}

// Creates a new ReleaseAddressRequest.
func NewReleaseAddressRequest(publicIp, allocationId string) *ReleaseAddressRequest {
	return &ReleaseAddressRequest{publicIp, allocationId}
}

/*****************************************************************************/

// Changes which network ACL a subnet is associated with. By default when you create
// a subnet, it's automatically associated with the default network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html]
type ReplaceNetworkAclAssociationRequest struct {
	AssociationId string `name:"AssociationId"`
	NetworkAclId  string `name:"NetworkAclId"`
}

// Creates a new ReplaceNetworkAclAssociationRequest.
func NewReplaceNetworkAclAssociationRequest(networkAclId, associationId string) *ReplaceNetworkAclAssociationRequest {
	return &ReplaceNetworkAclAssociationRequest{associationId, networkAclId}
}

/*****************************************************************************/

// Replaces an entry (rule) in a network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclEntry.html]
type ReplaceNetworkAclEntryRequest struct {
	NetworkAclId string            `name:"NetworkAclId"`
	RuleNumber   int               `name:"RuleNumber"`
	Protocol     int               `name:"Protocol"`
	RuleAction   string            `name:"RuleAction"`
	Egress       bool              `name:"Egress,omitempty"`
	CidrBlock    string            `name:"CidrBlock,omitempty"`
	Icmp         *IcmpTypeCodeType `name:"Icmp.,omitempty"`
	PortRange    *PortRangeType    `name:"PortRange.,omitempty"`
}

// Creates a new ReplaceNetworkAclEntryRequest.
func NewReplaceNetworkAclEntryRequest(networkAclId, ruleAction string, ruleNumber, protocol int) *ReplaceNetworkAclEntryRequest {
	return &ReplaceNetworkAclEntryRequest{
		NetworkAclId: networkAclId,
		RuleNumber:   ruleNumber,
		Protocol:     protocol,
		RuleAction:   ruleAction,
	}
}

/*****************************************************************************/

// Replaces an existing route within a route table in a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRoute.html]
type ReplaceRouteRequest struct {
	RouteTableId           string `name:"RouteTableId"`
	DestinationCidrBlock   string `name:"DestinationCidrBlock"`
	GatewayId              string `name:"GatewayId,omitempty"`
	InstanceId             string `name:"InstanceId,omitempty"`
	NetworkInterfaceId     string `name:"NetworkInterfaceId,omitempty"`
	VpcPeeringConnectionId string `name:"VpcPeeringConnectionId,omitempty"`
}

// Creates a new ReplaceRouteRequest.
func NewReplaceRouteRequest(routeTableId, destinationCidrBlock string) *ReplaceRouteRequest {
	return &ReplaceRouteRequest{RouteTableId: routeTableId, DestinationCidrBlock: destinationCidrBlock}
}

/*****************************************************************************/

// Changes the route table associated with a given subnet in a VPC. After the operation completes,
// the subnet uses the routes in the new route table it's associated with.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRouteTableAssociation.html]
type ReplaceRouteTableAssociationRequest struct {
	AssociationId string `name:"AssociationId"`
	RouteTableId  string `name:"RouteTableId"`
}

// Creates a new ReplaceRouteTableAssociationRequest.
func NewReplaceRouteTableAssociationRequest(routeTableId, associationId string) *ReplaceRouteTableAssociationRequest {
	return &ReplaceRouteTableAssociationRequest{associationId, routeTableId}
}

/*****************************************************************************/

// Submits feedback about an instance's status. The instance must be in the running state.
// If your experience with the instance differs from the instance status returned by
// DescribeInstanceStatus, use ReportInstanceStatus to report your experience with the
// instance. Amazon EC2 collects this information to improve the accuracy of status checks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReportInstanceStatus.html]
type ReportInstanceStatusRequest struct {
	InstanceIds []string      `name:"InstanceId.#"`
	Status      State         `name:"Status"`
	StartTime   time.Time     `name:"StartTime,omitempty" format:"2006-01-02T15:04:05.999Z"`
	EndTime     time.Time     `name:"EndTime,omitempty" format:"2006-01-02T15:04:05.999Z"`
	ReasonCodes []HealthState `name:"ReasonCode.#"`
	Description string        `name:"Description,omitempty"`
}

// Creates a new ReportInstanceStatusRequest.
func NewReportInstanceStatusRequest() *ReportInstanceStatusRequest {
	return &ReportInstanceStatusRequest{}
}

// Add a InstanceId.
func (r *ReportInstanceStatusRequest) AddInstanceId(instanceId string) {
	r.InstanceIds = append(r.InstanceIds, instanceId)
}

// Add a ReasonCode.
func (r *ReportInstanceStatusRequest) AddReasonCode(reasonCodes HealthState) {
	r.ReasonCodes = append(r.ReasonCodes, reasonCodes)
}

/*****************************************************************************/

// Creates a Spot Instance request. Spot Instances are instances that Amazon EC2 starts on your
// behalf when the maximum price that you specify exceeds the current Spot Price. Amazon EC2
// periodically sets the Spot Price based on available Spot Instance capacity and current Spot
// Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RequestSpotInstances.html]
type RequestSpotInstancesRequest struct {
	SpotPrice             string                          `name:"SpotPrice"`
	InstanceCount         int                             `name:"InstanceCount,omitempty"`
	Type                  InstanceType                    `name:"Type,omitempty"`
	ValidFrom             time.Time                       `name:"ValidFrom,omitempty" format:"2006-01-02T15:04:05.999Z"`
	ValidUntil            time.Time                       `name:"ValidUntil,omitempty" format:"2006-01-02T15:04:05.999Z"`
	LaunchGroup           string                          `name:"LaunchGroup,omitempty"`
	AvailabilityZoneGroup string                          `name:"AvailabilityZoneGroup,omitempty"`
	LaunchSpecification   *LaunchSpecificationRequestType `name:"LaunchSpecification."`
}

// Creates a new RequestSpotInstancesRequest.
func NewRequestSpotInstancesRequest(spotPrice, imageId, instanceType string) *RequestSpotInstancesRequest {
	return &RequestSpotInstancesRequest{
		SpotPrice: spotPrice,
		LaunchSpecification: &LaunchSpecificationRequestType{
			ImageId:      imageId,
			InstanceType: instanceType,
		},
	}
}

/*****************************************************************************/

// Resets an attribute of an AMI to its default value.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetImageAttribute.html]
type ResetImageAttributeRequest struct {
	ImageId   string         `name:"ImageId"`
	Attribute ImageAttribute `name:"Attribute"`
}

// Creates a new ResetImageAttributeRequest.
func NewResetImageAttributeRequest(imageId string, attribute ImageAttribute) *ResetImageAttributeRequest {
	return &ResetImageAttributeRequest{imageId, attribute}
}

/*****************************************************************************/

// Resets an attribute of an instance to its default value. To reset the kernel or RAM disk,
// the instance must be in a stopped state. To reset the SourceDestCheck, the instance can
// be either running or stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetInstanceAttribute.html]
type ResetInstanceAttributeRequest struct {
	ImageId   string            `name:"ImageId"`
	Attribute InstanceAttribute `name:"Attribute"`
}

// Creates a new ResetInstanceAttributeRequest.
func NewResetInstanceAttributeRequest(imageId string, attribute InstanceAttribute) *ResetInstanceAttributeRequest {
	return &ResetInstanceAttributeRequest{imageId, attribute}
}

/*****************************************************************************/

// Resets a network interface attribute. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetNetworkInterfaceAttribute.html]
type ResetNetworkInterfaceAttributeRequest struct {
	NetworkInterfaceId string            `name:"NetworkInterfaceId"`
	Attribute          InstanceAttribute `name:"Attribute"`
}

// Creates a new ResetNetworkInterfaceAttributeRequest.
func NewResetNetworkInterfaceAttributeRequest(networkInterfaceId string, attribute InstanceAttribute) *ResetNetworkInterfaceAttributeRequest {
	return &ResetNetworkInterfaceAttributeRequest{networkInterfaceId, attribute}
}

/*****************************************************************************/

// Resets permission settings for the specified snapshot.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetSnapshotAttribute.html]
type ResetSnapshotAttributeRequest struct {
	SnapshotId string            `name:"SnapshotId"`
	Attribute  InstanceAttribute `name:"Attribute"`
}

// Creates a new ResetSnapshotAttributeRequest.
func NewResetSnapshotAttributeRequest(snapshotId string, attribute InstanceAttribute) *ResetSnapshotAttributeRequest {
	return &ResetSnapshotAttributeRequest{snapshotId, attribute}
}

/*****************************************************************************/

// Removes one or more egress rules from a security group for EC2-VPC. The values that you specify
// in the revoke request (for example, ports) must match the existing rule's values for the rule to be revoked.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupEgress.html]
type RevokeSecurityGroupEgressRequest struct {
	GroupId       string             `name:"GroupId"`
	IpPermissions []IpPermissionType `name:"IpPermissions.#." base:"1"`
}

// Creates a new RevokeSecurityGroupEgressRequest.
func NewRevokeSecurityGroupEgressRequest(groupId, ipProtocol, userGroupId string) *RevokeSecurityGroupEgressRequest {
	return &RevokeSecurityGroupEgressRequest{
		GroupId: groupId,
		IpPermissions: []IpPermissionType{
			IpPermissionType{
				IpProtocol: ipProtocol,
				Groups: []UserIdGroupPairType{
					UserIdGroupPairType{
						GroupId: userGroupId,
					},
				},
			},
		},
	}
}

/*****************************************************************************/

// Removes one or more ingress rules from a security group. The values that you specify in the revoke request
// (for example, ports) must match the existing rule's values for the rule to be removed.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupIngress.html]
type RevokeSecurityGroupIngressRequest struct {
	GroupId       string             `name:"GroupId,omitempty"`
	GroupName     string             `name:"GroupName,omitempty"`
	IpPermissions []IpPermissionType `name:"IpPermissions.#.,omitempty" base:"1"`
}

// Creates a new RevokeSecurityGroupIngressRequest.
func NewRevokeSecurityGroupIngressRequest() *RevokeSecurityGroupIngressRequest {
	return &RevokeSecurityGroupIngressRequest{}
}

/*****************************************************************************/

// Launches the specified number of instances using an AMI for which you have permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RunInstances.html]
type RunInstancesRequest struct {
	ImageId                           string                                       `name:"ImageId"`
	MinCount                          int                                          `name:"MinCount"`
	MaxCount                          int                                          `name:"MaxCount"`
	KeyName                           string                                       `name:"KeyName,omitempty"`
	SecurityGroupIds                  []string                                     `name:"SecurityGroupId.#,omitempty" base:"1"`
	SecurityGroups                    []string                                     `name:"SecurityGroup.#,omitempty" base:"1"`
	UserData                          string                                       `name:"UserData,omitempty"`
	InstanceType                      string                                       `name:"InstanceType,omitempty"`
	Placement                         *PlacementResponseType                       `name:"Placement.,omitempty"`
	KernelId                          string                                       `name:"KernelId,omitempty"`
	RamdiskId                         string                                       `name:"RamdiskId,omitempty"`
	BlockDeviceMappings               []InstanceBlockDeviceMappingItemType         `name:"BlockDeviceMappings.#.,omitempty" base:"1"`
	MonitoringEnabled                 bool                                         `name:"Monitoring.Enabled,omitempty"`
	SubnetId                          string                                       `name:"SubnetId,omitempty"`
	DisableApiTermination             bool                                         `name:"DisableApiTermination,omitempty"`
	InstanceInitiatedShutdownBehavior string                                       `name:"InstanceInitiatedShutdownBehavior,omitempty"`
	PrivateIpAddress                  string                                       `name:"PrivateIpAddress,omitempty"`
	ClientToken                       string                                       `name:"ClientToken"`
	NetworkInterfaces                 []InstanceNetworkInterfaceSetItemRequestType `name:"NetworkInterface.#.,omitempty" base:"0"`
	IamInstanceProfile                *IamInstanceProfileRequestType               `name:"IamInstanceProfile.,omitempty"`
	EbsOptimized                      bool                                         `name:"EbsOptimized,omitempty"`
}

// Creates a new RunInstancesRequest.
func NewRunInstancesRequest(imageId string, minCount, maxCount int) *RunInstancesRequest {
	return &RunInstancesRequest{
		ImageId:  imageId,
		MinCount: minCount,
		MaxCount: maxCount,
	}
}

/*****************************************************************************/

// Starts an Amazon EBS-backed AMI that you've previously stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StartInstances.html]
type StartInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
}

// Creates a new StartInstancesRequest.
func NewStartInstancesRequest(instanceIds ...string) *StartInstancesRequest {
	return &StartInstancesRequest{instanceIds}
}

/*****************************************************************************/

// Stops an Amazon EBS-backed instance. Each time you transition an instance from stopped to started,
// Amazon EC2 charges a full instance hour, even if transitions happen multiple times within a single hour.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StopInstances.html]
type StopInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
	Force       bool     `name:"Force,omitempty"`
}

// Creates a new StopInstancesRequest.
func NewStopInstancesRequest(instanceIds ...string) *StopInstancesRequest {
	return &StopInstancesRequest{InstanceIds: instanceIds}
}

/*****************************************************************************/

// Shuts down one or more instances. This operation is idempotent; if you terminate an
// instance more than once, each call succeeds.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-TerminateInstances.html]
type TerminateInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
}

// Creates a new TerminateInstancesRequest.
func NewTerminateInstancesRequest(instanceIds ...string) *TerminateInstancesRequest {
	return &TerminateInstancesRequest{instanceIds}
}

/*****************************************************************************/

// Unassigns one or more secondary private IP addresses from a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnassignPrivateIpAddresses.html]
type UnassignPrivateIpAddressesRequest struct {
	NetworkInterfaceId string   `name:"NetworkInterfaceId"`
	PrivateIpAddresses []string `name:"PrivateIpAddress.#"`
}

// Creates a new UnassignPrivateIpAddressesRequest.
func NewUnassignPrivateIpAddressesRequest(networkInterfaceId string, privateIpAddresses ...string) *UnassignPrivateIpAddressesRequest {
	return &UnassignPrivateIpAddressesRequest{networkInterfaceId, privateIpAddresses}
}

/*****************************************************************************/

// Disables monitoring for a running instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnmonitorInstances.html]
type UnmonitorInstancesRequest struct {
	InstanceIds []string `name:"InstanceId.#" base:"1"`
}

// Creates a new UnmonitorInstancesRequest.
func NewUnmonitorInstancesRequest(instanceIds ...string) *UnmonitorInstancesRequest {
	return &UnmonitorInstancesRequest{instanceIds}
}
