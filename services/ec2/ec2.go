//
// Amazon Elastic Compute Cloud (Amazon EC2) is a web service that provides resizeable
// computing capacity—literally, servers in Amazon's data centers—that you use to build
// and host your software systems.
//
// [http://aws.amazon.com/documentation/ec2/]
//
package ec2

import (
	"github.com/twhello/aws-to-go/auth"
	"github.com/twhello/aws-to-go/interfaces"
	"github.com/twhello/aws-to-go/regions"
	"github.com/twhello/aws-to-go/services"
	"github.com/twhello/aws-to-go/util/netutil"
	"net/http"
)

const ServiceName = "ec2"

// The EC2 Service object. Use ec2.NewService().
type EC2Service struct {
	cred     interfaces.IAWSCredentials
	region   *regions.Region
	endpoint string
}

// Returns the name of the service.
func (s *EC2Service) ServiceName() string {
	return ServiceName
}

// Returns the region name the service will call.
func (s *EC2Service) RegionName() string {
	return s.region.Name()
}

// Returns the endpoint to the service.
func (s *EC2Service) Endpoint() string {
	return s.endpoint
}

// Low-level request to EC2 service.
func (s *EC2Service) SignAndDo(req interfaces.IAWSRequest, dto interface{}) (resp *http.Response, err error) {

	signer := auth.V4Signer{s.cred, s}
	signer.Sign(req)

	resp, err = services.DoRequest(req, dto, services.NewEvalXmlServiceResponse())

	return
}

func (s *EC2Service) wrapperSignAndDo(action string, request, result interface{}) (err error) {

	qs := netutil.MarshalValues(request)
	qs.Add("Version", "2014-06-15")
	qs.Add("Action", action)

	req, err := services.NewClientRequest("GET", s.Endpoint(), qs)
	if err == nil {
		_, err = s.SignAndDo(req, result)
	}

	return
}

/******************************************************************************
 * EC2 Service Methods.
 */

// Accepts a VPC peering connection request. To accept a request, the VPC peering
// connection must be in the pending-acceptance state, and you must be the owner
// of the peer VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AcceptVpcPeeringConnection.html]
func (s *EC2Service) AcceptVpcPeeringConnection(req *AcceptVpcPeeringConnectionRequest) (result *AcceptVpcPeeringConnectionResponse, err error) {

	result = new(AcceptVpcPeeringConnectionResponse)
	err = s.wrapperSignAndDo("AcceptVpcPeeringConnection", req, result)
	return
}

// Acquires an Elastic IP address.
// An Elastic IP address is for use either in the EC2-Classic platform or in a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AllocateAddress.html]
func (s *EC2Service) AllocateAddress(req *AllocateAddressRequest) (result *AllocateAddressResponse, err error) {

	result = new(AllocateAddressResponse)
	err = s.wrapperSignAndDo("AllocateAddress", req, result)
	return
}

// Assigns one or more secondary private IP addresses to the specified network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssignPrivateIpAddresses.html]
func (s *EC2Service) AssignPrivateIpAddresses(req *AssignPrivateIpAddressesRequest) (result *AssignPrivateIpAddressesResponse, err error) {

	result = new(AssignPrivateIpAddressesResponse)
	err = s.wrapperSignAndDo("AssignPrivateIpAddresses", req, result)
	return
}

// Associates an Elastic IP address with an instance or a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateAddress.html]
func (s *EC2Service) AssociateAddress(req *AssociateAddressRequest) (result *AssociateAddressResponse, err error) {

	result = new(AssociateAddressResponse)
	err = s.wrapperSignAndDo("", req, result)
	return
}

// Associates a set of DHCP options (that you've previously created) with the
// specified VPC, or associates no DHCP options with the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateDhcpOptions.html]
func (s *EC2Service) AssociateDhcpOptions(req *AssociateDhcpOptionsRequest) (result *AssociateDhcpOptionsResponse, err error) {

	result = new(AssociateDhcpOptionsResponse)
	err = s.wrapperSignAndDo("AssociateDhcpOptions", req, result)
	return
}

// Associates a subnet with a route table. The subnet and route table must be in the same VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AssociateRouteTable.html]
func (s *EC2Service) AssociateRouteTable(req *AssociateRouteTableRequest) (result *AssociateRouteTableResponse, err error) {

	result = new(AssociateRouteTableResponse)
	err = s.wrapperSignAndDo("AssociateRouteTable", req, result)
	return
}

// Attaches an Internet gateway to a VPC, enabling connectivity between the Internet and the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachInternetGateway.html]
func (s *EC2Service) AttachInternetGateway(req *AttachInternetGatewayRequest) (result *AttachInternetGatewayResponse, err error) {

	result = new(AttachInternetGatewayResponse)
	err = s.wrapperSignAndDo("AttachInternetGateway", req, result)
	return
}

// Attaches a network interface to an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachNetworkInterface.html]
func (s *EC2Service) AttachNetworkInterface(req *AttachNetworkInterfaceRequest) (result *AttachNetworkInterfaceResponse, err error) {

	result = new(AttachNetworkInterfaceResponse)
	err = s.wrapperSignAndDo("AttachNetworkInterface", req, result)
	return
}

// Attaches an Amazon EBS volume to a running or stopped instance and exposes
// it to the instance with the specified device name.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVolume.html]
func (s *EC2Service) AttachVolume(req *AttachVolumeRequest) (result *AttachVolumeResponse, err error) {

	result = new(AttachVolumeResponse)
	err = s.wrapperSignAndDo("AttachVolume", req, result)
	return
}

// Attaches a virtual private gateway to a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVpnGateway.html]
func (s *EC2Service) AttachVpnGateway(req *AttachVpnGatewayRequest) (result *AttachVpnGatewayResponse, err error) {

	result = new(AttachVpnGatewayResponse)
	err = s.wrapperSignAndDo("AttachVpnGateway", req, result)
	return
}

// Adds one or more egress rules to a security group for use with a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupEgress.html]
func (s *EC2Service) AuthorizeSecurityGroupEgress(req *AuthorizeSecurityGroupEgressRequest) (result *AuthorizeSecurityGroupEgressResponse, err error) {

	result = new(AuthorizeSecurityGroupEgressResponse)
	err = s.wrapperSignAndDo("AuthorizeSecurityGroupEgress", req, result)
	return
}

// Adds one or more ingress rules to a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AuthorizeSecurityGroupIngress.html]
func (s *EC2Service) AuthorizeSecurityGroupIngress(req *AuthorizeSecurityGroupIngressRequest) (result *AuthorizeSecurityGroupIngressResponse, err error) {

	result = new(AuthorizeSecurityGroupIngressResponse)
	err = s.wrapperSignAndDo("AuthorizeSecurityGroupIngress", req, result)
	return
}

// Bundles an Amazon instance store-backed Windows instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-BundleInstance.html]
func (s *EC2Service) BundleInstance(req *BundleInstanceRequest) (result *BundleInstanceResponse, err error) {

	result = new(BundleInstanceResponse)
	err = s.wrapperSignAndDo("BundleInstance", req, result)
	return
}

// Cancels a bundling operation for an instance store-backed Windows instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelBundleTask.html]
func (s *EC2Service) CancelBundleTask(req *CancelBundleTaskRequest) (result *CancelBundleTaskResponse, err error) {

	result = new(CancelBundleTaskResponse)
	err = s.wrapperSignAndDo("CancelBundleTask", req, result)
	return
}

// Cancels an active conversion task. The task can be the import of an instance or volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelConversionTask.html]
func (s *EC2Service) CancelConversionTask(req *CancelConversionTaskRequest) (result *CancelConversionTaskResponse, err error) {

	result = new(CancelConversionTaskResponse)
	err = s.wrapperSignAndDo("CancelConversionTask", req, result)
	return
}

// Cancels an active export task. The request removes all artifacts of the export,
// including any partially-created Amazon S3 objects.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelExportTask.html]
func (s *EC2Service) CancelExportTask(req *CancelExportTaskRequest) (result *CancelExportTaskResponse, err error) {

	result = new(CancelExportTaskResponse)
	err = s.wrapperSignAndDo("CancelExportTask", req, result)
	return
}

// Cancels the specified Reserved Instance listing in the Reserved Instance Marketplace.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelReservedInstancesListing.html]
func (s *EC2Service) CancelReservedInstancesListing(req *CancelReservedInstancesListingRequest) (result *CancelReservedInstancesListingResponse, err error) {

	result = new(CancelReservedInstancesListingResponse)
	err = s.wrapperSignAndDo("CancelReservedInstancesListing", req, result)
	return
}

// Cancels one or more Spot Instance requests. Spot Instances are instances that
// Amazon EC2 starts on your behalf when the maximum price that you specify exceeds
// the current Spot Price. Amazon EC2 periodically sets the Spot Price based on
// available Spot Instance capacity and current Spot Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CancelSpotInstanceRequests.html]
func (s *EC2Service) CancelSpotInstanceRequests(req *CancelSpotInstanceRequestsRequest) (result *CancelSpotInstanceRequestsResponse, err error) {

	result = new(CancelSpotInstanceRequestsResponse)
	err = s.wrapperSignAndDo("CancelSpotInstanceRequests", req, result)
	return
}

// Determines whether a product code is associated with an instance. This action
// can only be used by the owner of the product code. It is useful when a product
// code owner needs to verify whether another user's instance is eligible for support.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ConfirmProductInstance.html]
func (s *EC2Service) ConfirmProductInstance(req *ConfirmProductInstanceRequest) (result *ConfirmProductInstanceResponse, err error) {

	result = new(ConfirmProductInstanceResponse)
	err = s.wrapperSignAndDo("ConfirmProductInstance", req, result)
	return
}

// Initiates the copy of an AMI from the specified source region to the current region.
// You specify the destination region by using its endpoint when making the request.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopyImage.html]
func (s *EC2Service) CopyImage(req *CopyImageRequest) (result *CopyImageResponse, err error) {

	result = new(CopyImageResponse)
	err = s.wrapperSignAndDo("CopyImage", req, result)
	return
}

// Copies a point-in-time snapshot of an Amazon Elastic Block Store (Amazon EBS) volume and
// stores it in Amazon Simple Storage Service (Amazon S3). You can copy the snapshot within
// the same region or from one region to another. The snapshot is copied to the regional
// endpoint that you send the HTTP request to.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CopySnapshot.html]
func (s *EC2Service) CopySnapshot(req *CopySnapshotRequest) (result *CopySnapshotResponse, err error) {

	result = new(CopySnapshotResponse)
	err = s.wrapperSignAndDo("CopySnapshot", req, result)
	return
}

// Provides information to AWS about your VPN customer gateway device. The customer
// gateway is the appliance at your end of the VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateCustomerGateway.html]
func (s *EC2Service) CreateCustomerGateway(req *CreateCustomerGatewayRequest) (result *CreateCustomerGatewayResponse, err error) {

	result = new(CreateCustomerGatewayResponse)
	err = s.wrapperSignAndDo("CreateCustomerGateway", req, result)
	return
}

// Creates a set of DHCP options for your VPC. After creating the set, you must associate
// it with the VPC, causing all existing and new instances that you launch in the VPC
// to use this set of DHCP options. The following are the individual DHCP options you can specify.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateDhcpOptions.html]
func (s *EC2Service) CreateDhcpOptions(req *CreateDhcpOptionsRequest) (result *CreateDhcpOptionsResponse, err error) {

	result = new(CreateDhcpOptionsResponse)
	err = s.wrapperSignAndDo("CreateDhcpOptions", req, result)
	return
}

// Creates an Amazon EBS-backed AMI from an Amazon EBS-backed instance that is either running or stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateImage.html]
func (s *EC2Service) CreateImage(req *CreateImageRequest) (result *CreateImageResponse, err error) {

	result = new(CreateImageResponse)
	err = s.wrapperSignAndDo("CreateImage", req, result)
	return
}

// Exports a running or stopped instance to an Amazon S3 bucket.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInstanceExportTask.html]
func (s *EC2Service) CreateInstanceExportTask(req *CreateInstanceExportTaskRequest) (result *CreateInstanceExportTaskResponse, err error) {

	result = new(CreateInstanceExportTaskResponse)
	err = s.wrapperSignAndDo("CreateInstanceExportTask", req, result)
	return
}

// Creates an Internet gateway for use with a VPC. After creating the Internet
// gateway, you attach it to a VPC using AttachInternetGateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInternetGateway.html]
func (s *EC2Service) CreateInternetGateway(req *CreateInternetGatewayRequest) (result *CreateInternetGatewayResponse, err error) {

	result = new(CreateInternetGatewayResponse)
	err = s.wrapperSignAndDo("CreateInternetGateway", req, result)
	return
}

// Creates a 2048-bit RSA key pair with the specified name. Amazon EC2 stores the
// public key and displays the private key for you to save to a file. The private
// key is returned as an unencrypted PEM encoded PKCS#8 private key. If a key with
// the specified name already exists, Amazon EC2 returns an error.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateKeyPair.html]
func (s *EC2Service) CreateKeyPair(req *CreateKeyPairRequest) (result *CreateKeyPairResponse, err error) {

	result = new(CreateKeyPairResponse)
	err = s.wrapperSignAndDo("CreateKeyPair", req, result)
	return
}

// Creates a network ACL in a VPC. Network ACLs provide an optional layer of security
// (in addition to security groups) for the instances in your VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAcl.html]
func (s *EC2Service) CreateNetworkAcl(req *CreateNetworkAclRequest) (result *CreateNetworkAclResponse, err error) {

	result = new(CreateNetworkAclResponse)
	err = s.wrapperSignAndDo("CreateNetworkAcl", req, result)
	return
}

// Creates an entry (a rule) in a network ACL with the specified rule number. Each network ACL
// has a set of numbered ingress rules and a separate set of numbered egress rules. When
// determining whether a packet should be allowed in or out of a subnet associated with the
// ACL, we process the entries in the ACL according to the rule numbers, in ascending order.
// Each network ACL has a set of ingress rules and a separate set of egress rules.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkAclEntry.html]
func (s *EC2Service) CreateNetworkAclEntry(req *CreateNetworkAclEntryRequest) (result *CreateNetworkAclEntryResponse, err error) {

	result = new(CreateNetworkAclEntryResponse)
	err = s.wrapperSignAndDo("CreateNetworkAclEntry", req, result)
	return
}

// Creates a network interface in the specified subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateNetworkInterface.html]
func (s *EC2Service) CreateNetworkInterface(req *CreateNetworkInterfaceRequest) (result *CreateNetworkInterfaceResponse, err error) {

	result = new(CreateNetworkInterfaceResponse)
	err = s.wrapperSignAndDo("CreateNetworkInterface", req, result)
	return
}

// Creates a placement group that you launch cluster instances into. You must
// give the group a name that's unique within the scope of your account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreatePlacementGroup.html]
func (s *EC2Service) CreatePlacementGroup(req *CreatePlacementGroupRequest) (result *CreatePlacementGroupResponse, err error) {

	result = new(CreatePlacementGroupResponse)
	err = s.wrapperSignAndDo("CreatePlacementGroup", req, result)
	return
}

// Creates a listing for Amazon EC2 Reserved Instances to be sold in the Reserved
// Instance Marketplace. You can submit one Reserved Instance listing at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateReservedInstancesListing.html]
func (s *EC2Service) CreateReservedInstancesListing(req *CreateReservedInstancesListingRequest) (result *CreateReservedInstancesListingResponse, err error) {

	result = new(CreateReservedInstancesListingResponse)
	err = s.wrapperSignAndDo("CreateReservedInstancesListing", req, result)
	return
}

// Creates a route in a route table within a VPC. The route's target can be an
// Internet gateway or virtual private gateway attached to the VPC, a VPC peering
// connection, or a NAT instance in the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRoute.html]
func (s *EC2Service) CreateRoute(req *CreateRouteRequest) (result *CreateRouteResponse, err error) {

	result = new(CreateRouteResponse)
	err = s.wrapperSignAndDo("CreateRoute", req, result)
	return
}

// Creates a route table for the specified VPC. After you create a route table,
// you can add routes and associate the table with a subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateRouteTable.html]
func (s *EC2Service) CreateRouteTable(req *CreateRouteTableRequest) (result *CreateRouteTableResponse, err error) {

	result = new(CreateRouteTableResponse)
	err = s.wrapperSignAndDo("CreateRouteTable", req, result)
	return
}

// Creates a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSecurityGroup.html]
func (s *EC2Service) CreateSecurityGroup(req *CreateSecurityGroupRequest) (result *CreateSecurityGroupResponse, err error) {

	result = new(CreateSecurityGroupResponse)
	err = s.wrapperSignAndDo("CreateSecurityGroup", req, result)
	return
}

// Creates a snapshot of an Amazon EBS volume and stores it in Amazon S3. You can use snapshots
// for backups, to make copies of instance store volumes, and to save data before shutting
// down an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSnapshot.html]
func (s *EC2Service) CreateSnapshot(req *CreateSnapshotRequest) (result *CreateSnapshotResponse, err error) {

	result = new(CreateSnapshotResponse)
	err = s.wrapperSignAndDo("CreateSnapshot", req, result)
	return
}

// Creates the datafeed for Spot Instances, enabling you to view Spot Instance usage logs.
// You can create one data feed per account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSpotDatafeedSubscription.html]
func (s *EC2Service) CreateSpotDatafeedSubscription(req *CreateSpotDatafeedSubscriptionRequest) (result *CreateSpotDatafeedSubscriptionResponse, err error) {

	result = new(CreateSpotDatafeedSubscriptionResponse)
	err = s.wrapperSignAndDo("CreateSpotDatafeedSubscription", req, result)
	return
}

// Creates a subnet in an existing VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateSubnet.html]
func (s *EC2Service) CreateSubnet(req *CreateSubnetRequest) (result *CreateSubnetResponse, err error) {

	result = new(CreateSubnetResponse)
	err = s.wrapperSignAndDo("CreateSubnet", req, result)
	return
}

// Adds or overwrites one or more tags for the specified Amazon EC2 resource or resources.
// Each resource can have a maximum of 10 tags. Each tag consists of a key and optional value.
// Tag keys must be unique per resource.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateTags.html]
func (s *EC2Service) CreateTags(req *CreateTagsRequest) (result *CreateTagsResponse, err error) {

	result = new(CreateTagsResponse)
	err = s.wrapperSignAndDo("CreateTags", req, result)
	return
}

// Creates an Amazon EBS volume that can be attached to an instance in the same Availability Zone.
// The volume is created in the regional endpoint that you send the HTTP request to.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVolume.html]
func (s *EC2Service) CreateVolume(req *CreateVolumeRequest) (result *CreateVolumeResponse, err error) {

	result = new(CreateVolumeResponse)
	err = s.wrapperSignAndDo("CreateVolume", req, result)
	return
}

// Creates a VPC with the specified CIDR block.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpc.html]
func (s *EC2Service) CreateVpc(req *CreateVpcRequest) (result *CreateVpcResponse, err error) {

	result = new(CreateVpcResponse)
	err = s.wrapperSignAndDo("CreateVpc", req, result)
	return
}

// Requests a VPC peering connection between two VPCs: a requester VPC that you own and a
// peer VPC with which to create the connection. The peer VPC can belong to another
// AWS account. The requester VPC and peer VPC cannot have overlapping CIDR blocks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpcPeeringConnection.html]
func (s *EC2Service) CreateVpcPeeringConnection(req *CreateVpcPeeringConnectionRequest) (result *CreateVpcPeeringConnectionResponse, err error) {

	result = new(CreateVpcPeeringConnectionResponse)
	err = s.wrapperSignAndDo("CreateVpcPeeringConnection", req, result)
	return
}

// Creates a VPN connection between an existing virtual private gateway and a VPN customer gateway.
// The only supported connection type is ipsec.1.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnection.html]
func (s *EC2Service) CreateVpnConnection(req *CreateVpnConnectionRequest) (result *CreateVpnConnectionResponse, err error) {

	result = new(CreateVpnConnectionResponse)
	err = s.wrapperSignAndDo("CreateVpnConnection", req, result)
	return
}

// Creates a static route associated with a VPN connection between an existing virtual
// private gateway and a VPN customer gateway. The static route allows traffic to be
// routed from the virtual private gateway to the VPN customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnConnectionRoute.html]
func (s *EC2Service) CreateVpnConnectionRoute(req *CreateVpnConnectionRouteRequest) (result *CreateVpnConnectionRouteResponse, err error) {

	result = new(CreateVpnConnectionRouteResponse)
	err = s.wrapperSignAndDo("CreateVpnConnectionRoute", req, result)
	return
}

// Creates a virtual private gateway. A virtual private gateway is the endpoint on the VPC side
// of your VPN connection. You can create a virtual private gateway before creating the VPC itself.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateVpnGateway.html]
func (s *EC2Service) CreateVpnGateway(req *CreateVpnGatewayRequest) (result *CreateVpnGatewayResponse, err error) {

	result = new(CreateVpnGatewayResponse)
	err = s.wrapperSignAndDo("CreateVpnGateway", req, result)
	return
}

// Deletes the specified customer gateway. You must delete the VPN connection
// before you can delete the customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteCustomerGateway.html]
func (s *EC2Service) DeleteCustomerGateway(req *DeleteCustomerGatewayRequest) (result *DeleteCustomerGatewayResponse, err error) {

	result = new(DeleteCustomerGatewayResponse)
	err = s.wrapperSignAndDo("DeleteCustomerGateway", req, result)
	return
}

// Deletes the specified set of DHCP options. You must disassociate the set of DHCP options
// before you can delete it. You can disassociate the set of DHCP options by associating
// either a new set of options or the default set of options with the VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteDhcpOptions.html]
func (s *EC2Service) DeleteDhcpOptions(req *DeleteDhcpOptionsRequest) (result *DeleteDhcpOptionsResponse, err error) {

	result = new(DeleteDhcpOptionsResponse)
	err = s.wrapperSignAndDo("DeleteDhcpOptions", req, result)
	return
}

// Deletes the specified Internet gateway. You must detach the Internet gateway
// from the VPC before you can delete it.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteInternetGateway.html]
func (s *EC2Service) DeleteInternetGateway(req *DeleteInternetGatewayRequest) (result *DeleteInternetGatewayResponse, err error) {

	result = new(DeleteInternetGatewayResponse)
	err = s.wrapperSignAndDo("DeleteInternetGateway", req, result)
	return
}

// Deletes the specified key pair, by removing the public key from Amazon EC2.
// You must own the key pair.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteKeyPair.html]
func (s *EC2Service) DeleteKeyPair(req *DeleteKeyPairRequest) (result *DeleteKeyPairResponse, err error) {

	result = new(DeleteKeyPairResponse)
	err = s.wrapperSignAndDo("DeleteKeyPair", req, result)
	return
}

// Deletes the specified network ACL. You can't delete the ACL if it's associated
// with any subnets. You can't delete the default network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAcl.html]
func (s *EC2Service) DeleteNetworkAcl(req *DeleteNetworkAclRequest) (result *DeleteNetworkAclResponse, err error) {

	result = new(DeleteNetworkAclResponse)
	err = s.wrapperSignAndDo("DeleteNetworkAcl", req, result)
	return
}

// Deletes the specified ingress or egress entry (rule) from the specified network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkAclEntry.html]
func (s *EC2Service) DeleteNetworkAclEntry(req *DeleteNetworkAclEntryRequest) (result *DeleteNetworkAclEntryResponse, err error) {

	result = new(DeleteNetworkAclEntryResponse)
	err = s.wrapperSignAndDo("DeleteNetworkAclEntry", req, result)
	return
}

// Deletes the specified network interface. You must detach the network interface
// before you can delete it.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteNetworkInterface.html]
func (s *EC2Service) DeleteNetworkInterface(req *DeleteNetworkInterfaceRequest) (result *DeleteNetworkInterfaceResponse, err error) {

	result = new(DeleteNetworkInterfaceResponse)
	err = s.wrapperSignAndDo("DeleteNetworkInterface", req, result)
	return
}

// Deletes the specified placement group. You must terminate all instances in the
// placement group before you can delete the placement group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeletePlacementGroup.html]
func (s *EC2Service) DeletePlacementGroup(req *DeletePlacementGroupRequest) (result *DeletePlacementGroupResponse, err error) {

	result = new(DeletePlacementGroupResponse)
	err = s.wrapperSignAndDo("DeletePlacementGroup", req, result)
	return
}

// Deletes the specified route from the specified route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRoute.html]
func (s *EC2Service) DeleteRoute(req *DeleteRouteRequest) (result *DeleteRouteResponse, err error) {

	result = new(DeleteRouteResponse)
	err = s.wrapperSignAndDo("DeleteRoute", req, result)
	return
}

// Deletes the specified route table. You must disassociate the route table from any
// subnets before you can delete it. You can't delete the main route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteRouteTable.html]
func (s *EC2Service) DeleteRouteTable(req *DeleteRouteTableRequest) (result *DeleteRouteTableResponse, err error) {

	result = new(DeleteRouteTableResponse)
	err = s.wrapperSignAndDo("DeleteRouteTable", req, result)
	return
}

// Deletes a security group.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSecurityGroup.html]
func (s *EC2Service) DeleteSecurityGroup(req *DeleteSecurityGroupRequest) (result *DeleteSecurityGroupResponse, err error) {

	result = new(DeleteSecurityGroupResponse)
	err = s.wrapperSignAndDo("DeleteSecurityGroup", req, result)
	return
}

// Deletes the specified snapshot. When you make periodic snapshots of a volume,
// the snapshots are incremental, and only the blocks on the device that have
// changed since your last snapshot are saved in the new snapshot. When you delete
// a snapshot, only the data not needed for any other snapshot is removed. So
// regardless of which prior snapshots have been deleted, all active snapshots will
// have access to all the information needed to restore the volume.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSnapshot.html]
func (s *EC2Service) DeleteSnapshot(req *DeleteSnapshotRequest) (result *DeleteSnapshotResponse, err error) {

	result = new(DeleteSnapshotResponse)
	err = s.wrapperSignAndDo("DeleteSnapshot", req, result)
	return
}

// Deletes the datafeed for Spot Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSpotDatafeedSubscription.html]
func (s *EC2Service) DeleteSpotDatafeedSubscription() (result *DeleteSpotDatafeedSubscriptionResponse, err error) {

	result = new(DeleteSpotDatafeedSubscriptionResponse)
	err = s.wrapperSignAndDo("DeleteSpotDatafeedSubscription", nil, result)
	return
}

// Deletes the specified subnet. You must terminate all running instances in the
// subnet before you can delete the subnet.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteSubnet.html]
func (s *EC2Service) DeleteSubnet(req *DeleteSubnetRequest) (result *DeleteSubnetResponse, err error) {

	result = new(DeleteSubnetResponse)
	err = s.wrapperSignAndDo("DeleteSubnet", req, result)
	return
}

// Deletes the specified set of tags from the specified set of resources. This call is designed to follow a DescribeTags call.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteTags.html]
func (s *EC2Service) DeleteTags(req *DeleteTagsRequest) (result *DeleteTagsResponse, err error) {

	result = new(DeleteTagsResponse)
	err = s.wrapperSignAndDo("DeleteTags", req, result)
	return
}

// Deletes the specified Amazon EBS volume. The volume must be in the available
// state (not attached to an instance).
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVolume.html]
func (s *EC2Service) DeleteVolume(req *DeleteVolumeRequest) (result *DeleteVolumeResponse, err error) {

	result = new(DeleteVolumeResponse)
	err = s.wrapperSignAndDo("DeleteVolume", req, result)
	return
}

// Deletes the specified VPC. You must detach or delete all gateways and resources that
// are associated with the VPC before you can delete it. For example, you must terminate
// all instances running in the VPC, delete all security groups associated with the VPC
// (except the default one), delete all route tables associated with the VPC
// (except the default one), and so on.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpc.html]
func (s *EC2Service) DeleteVpc(req *DeleteVpcRequest) (result *DeleteVpcResponse, err error) {

	result = new(DeleteVpcResponse)
	err = s.wrapperSignAndDo("DeleteVpc", req, result)
	return
}

// Deletes a VPC peering connection. Either the owner of the requester VPC or the
// owner of the peer VPC can delete the VPC peering connection if it's in the
// active state. The owner of the requester VPC can delete a VPC peering connection
// in the pending-acceptance state.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpcPeeringConnection.html]
func (s *EC2Service) DeleteVpcPeeringConnection(req *DeleteVpcPeeringConnectionRequest) (result *DeleteVpcPeeringConnectionResponse, err error) {

	result = new(DeleteVpcPeeringConnectionResponse)
	err = s.wrapperSignAndDo("DeleteVpcPeeringConnection", req, result)
	return
}

// Deletes the specified VPN connection.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnection.html]
func (s *EC2Service) DeleteVpnConnection(req *DeleteVpnConnectionRequest) (result *DeleteVpnConnectionResponse, err error) {

	result = new(DeleteVpnConnectionResponse)
	err = s.wrapperSignAndDo("DeleteVpnConnection", req, result)
	return
}

// Deletes the specified static route associated with a VPN connection between an
// existing virtual private gateway and a VPN customer gateway. The static route
// allows traffic to be routed from the virtual private gateway to the VPN customer gateway.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnConnectionRoute.html]
func (s *EC2Service) DeleteVpnConnectionRoute(req *DeleteVpnConnectionRouteRequest) (result *DeleteVpnConnectionRouteResponse, err error) {

	result = new(DeleteVpnConnectionRouteResponse)
	err = s.wrapperSignAndDo("DeleteVpnConnectionRoute", req, result)
	return
}

// Deletes the specified virtual private gateway. We recommend that before you delete
// a virtual private gateway, you detach it from the VPC and delete the VPN connection.
// Note that you don't need to delete the virtual private gateway if you plan to delete
// and recreate the VPN connection between your VPC and your network.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeleteVpnGateway.html]
func (s *EC2Service) DeleteVpnGateway(req *DeleteVpnGatewayRequest) (result *DeleteVpnGatewayResponse, err error) {

	result = new(DeleteVpnGatewayResponse)
	err = s.wrapperSignAndDo("DeleteVpnGateway", req, result)
	return
}

// Deregisters the specified AMI. After you deregister an AMI, it can't be used
// to launch new instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DeregisterImage.html]
func (s *EC2Service) DeregisterImage(req *DeregisterImageRequest) (result *DeregisterImageResponse, err error) {

	result = new(DeregisterImageResponse)
	err = s.wrapperSignAndDo("DeregisterImage", req, result)
	return
}

// Describes the specified attribute of your AWS account.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAccountAttributes.html]
func (s *EC2Service) DescribeAccountAttributes(req *DescribeAccountAttributesRequest) (result *DescribeAccountAttributesResponse, err error) {

	result = new(DescribeAccountAttributesResponse)
	err = s.wrapperSignAndDo("DescribeAccountAttributes", req, result)
	return
}

// Describes one or more of your Elastic IP addresses.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAddresses.html]
func (s *EC2Service) DescribeAddresses(req *DescribeAddressesRequest) (result *DescribeAddressesResponse, err error) {

	result = new(DescribeAddressesResponse)
	err = s.wrapperSignAndDo("DescribeAddresses", req, result)
	return
}

// Describes one or more of the Availability Zones that are available to you. The results include
// zones only for the region you're currently using. If there is an event impacting an
// Availability Zone, you can use this request to view the state and any provided message for
// that Availability Zone.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeAvailabilityZones.html]
func (s *EC2Service) DescribeAvailabilityZones(req *DescribeAvailabilityZonesRequest) (result *DescribeAvailabilityZonesResponse, err error) {

	result = new(DescribeAvailabilityZonesResponse)
	err = s.wrapperSignAndDo("DescribeAvailabilityZones", req, result)
	return
}

// Describes one or more of your bundling tasks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeBundleTasks.html]
func (s *EC2Service) DescribeBundleTasks(req *DescribeBundleTasksRequest) (result *DescribeBundleTasksResponse, err error) {

	result = new(DescribeBundleTasksResponse)
	err = s.wrapperSignAndDo("DescribeBundleTasks", req, result)
	return
}

// Describes one or more of your conversion tasks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeConversionTasks.html]
func (s *EC2Service) DescribeConversionTasks(req *DescribeConversionTasksRequest) (result *DescribeConversionTasksResponse, err error) {

	result = new(DescribeConversionTasksResponse)
	err = s.wrapperSignAndDo("DescribeConversionTasks", req, result)
	return
}

// Describes one or more of your VPN customer gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeCustomerGateways.html]
func (s *EC2Service) DescribeCustomerGateways(req *DescribeCustomerGatewaysRequest) (result *DescribeCustomerGatewaysResponse, err error) {

	result = new(DescribeCustomerGatewaysResponse)
	err = s.wrapperSignAndDo("DescribeCustomerGateways", req, result)
	return
}

// Describes one or more of your DHCP options sets.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeDhcpOptions.html]
func (s *EC2Service) DescribeDhcpOptions(req *DescribeDhcpOptionsRequest) (result *DescribeDhcpOptionsResponse, err error) {

	result = new(DescribeDhcpOptionsResponse)
	err = s.wrapperSignAndDo("DescribeDhcpOptions", req, result)
	return
}

// Describes the specified attribute of the specified AMI. You can specify only
// one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImageAttribute.html]
func (s *EC2Service) DescribeImageAttribute(req *DescribeImageAttributeRequest) (result *DescribeImageAttributeResponse, err error) {

	result = new(DescribeImageAttributeResponse)
	err = s.wrapperSignAndDo("DescribeImageAttribute", req, result)
	return
}

// Describes one or more of the images (AMIs, AKIs, and ARIs) available to you.
// Images available to you include public images, private images that you own,
// and private images owned by other AWS accounts but for which you have explicit
// launch permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeImages.html]
func (s *EC2Service) DescribeImages(req *DescribeImagesRequest) (result *DescribeImagesResponse, err error) {

	result = new(DescribeImagesResponse)
	err = s.wrapperSignAndDo("DescribeImages", req, result)
	return
}

// Describes the specified attribute of the specified instance.
// You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceAttribute.html]
func (s *EC2Service) DescribeInstanceAttribute(req *DescribeInstanceAttributeRequest) (result *DescribeInstanceAttributeResponse, err error) {

	result = new(DescribeInstanceAttributeResponse)
	err = s.wrapperSignAndDo("DescribeInstanceAttribute", req, result)
	return
}

// Describes one or more of your instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstances.html]
func (s *EC2Service) DescribeInstances(req *DescribeInstancesRequest) (result *DescribeInstancesResponse, err error) {

	result = new(DescribeInstancesResponse)
	err = s.wrapperSignAndDo("DescribeInstances", req, result)
	return
}

// Describes the status of one or more instances, including any scheduled events.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInstanceStatus.html]
func (s *EC2Service) DescribeInstanceStatus(req *DescribeInstanceStatusRequest) (result *DescribeInstanceStatusResponse, err error) {

	result = new(DescribeInstanceStatusResponse)
	err = s.wrapperSignAndDo("DescribeInstanceStatus", req, result)
	return
}

// Describes one or more of your Internet gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeInternetGateways.html]
func (s *EC2Service) DescribeInternetGateways(req *DescribeInternetGatewaysRequest) (result *DescribeInternetGatewaysResponse, err error) {

	result = new(DescribeInternetGatewaysResponse)
	err = s.wrapperSignAndDo("DescribeInternetGateways", req, result)
	return
}

// Describes one or more of your key pairs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeKeyPairs.html]
func (s *EC2Service) DescribeKeyPairs(req *DescribeKeyPairsRequest) (result *DescribeKeyPairsResponse, err error) {

	result = new(DescribeKeyPairsResponse)
	err = s.wrapperSignAndDo("DescribeKeyPairs", req, result)
	return
}

// Describes one or more of your network ACLs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkAcls.html]
func (s *EC2Service) DescribeNetworkAcls(req *DescribeNetworkAclsRequest) (result *DescribeNetworkAclsResponse, err error) {

	result = new(DescribeNetworkAclsResponse)
	err = s.wrapperSignAndDo("DescribeNetworkAcls", req, result)
	return
}

// Describes the specified attribute of the specified network interface. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaceAttribute.html]
func (s *EC2Service) DescribeNetworkInterfaceAttribute(req *DescribeNetworkInterfaceAttributeRequest) (result *DescribeNetworkInterfaceAttributeResponse, err error) {

	result = new(DescribeNetworkInterfaceAttributeResponse)
	err = s.wrapperSignAndDo("DescribeNetworkInterfaceAttribute", req, result)
	return
}

// Describes one or more of your network interfaces.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeNetworkInterfaces.html]
func (s *EC2Service) DescribeNetworkInterfaces(req *DescribeNetworkInterfacesRequest) (result *DescribeNetworkInterfacesResponse, err error) {

	result = new(DescribeNetworkInterfacesResponse)
	err = s.wrapperSignAndDo("DescribeNetworkInterfaces", req, result)
	return
}

// Describes one or more of your placement groups.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribePlacementGroups.html]
func (s *EC2Service) DescribePlacementGroups(req *DescribePlacementGroupsRequest) (result *DescribePlacementGroupsResponse, err error) {

	result = new(DescribePlacementGroupsResponse)
	err = s.wrapperSignAndDo("DescribePlacementGroups", req, result)
	return
}

// Describes one or more regions that are currently available to you.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRegions.html]
func (s *EC2Service) DescribeRegions(req *DescribeRegionsRequest) (result *DescribeRegionsResponse, err error) {

	result = new(DescribeRegionsResponse)
	err = s.wrapperSignAndDo("DescribeRegions", req, result)
	return
}

// Describes one or more of the Reserved Instances that you purchased.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstances.html]
func (s *EC2Service) DescribeReservedInstances(req *DescribeReservedInstancesRequest) (result *DescribeReservedInstancesResponse, err error) {

	result = new(DescribeReservedInstancesResponse)
	err = s.wrapperSignAndDo("DescribeReservedInstances", req, result)
	return
}

// Describes your account's Reserved Instance listings in the Reserved Instance Marketplace.
// This call returns information, such as the ID of the Reserved Instance with which a
// listing is associated.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesListings.html]
func (s *EC2Service) DescribeReservedInstancesListings(req *DescribeReservedInstancesListingsRequest) (result *DescribeReservedInstancesListingsResponse, err error) {

	result = new(DescribeReservedInstancesListingsResponse)
	err = s.wrapperSignAndDo("DescribeReservedInstancesListings", req, result)
	return
}

// Describes the modifications made to your Reserved Instances. If no parameter is specified,
// information about all your Reserved Instances modification requests is returned. If a
// modification ID is specified, only information about the specific modification is returned.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesModifications.html]
func (s *EC2Service) DescribeReservedInstancesModifications(req *DescribeReservedInstancesModificationsRequest) (result *DescribeReservedInstancesModificationsResponse, err error) {

	result = new(DescribeReservedInstancesModificationsResponse)
	err = s.wrapperSignAndDo("DescribeReservedInstancesModifications", req, result)
	return
}

// Describes Reserved Instance offerings that are available for purchase. With Amazon EC2 Reserved
// Instances, you purchase the right to launch Amazon EC2 instances for a period of time. During
// that time period, you do not receive insufficient capacity errors, and you pay a lower usage
// rate than the rate charged for On-Demand instances for the actual time used.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeReservedInstancesOfferings.html]
func (s *EC2Service) DescribeReservedInstancesOfferings(req *DescribeReservedInstancesOfferingsRequest) (result *DescribeReservedInstancesOfferingsResponse, err error) {

	result = new(DescribeReservedInstancesOfferingsResponse)
	err = s.wrapperSignAndDo("DescribeReservedInstancesOfferings", req, result)
	return
}

// Describes one or more of your route tables.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeRouteTables.html]
func (s *EC2Service) DescribeRouteTables(req *DescribeRouteTablesRequest) (result *DescribeRouteTablesResponse, err error) {

	result = new(DescribeRouteTablesResponse)
	err = s.wrapperSignAndDo("DescribeRouteTables", req, result)
	return
}

// Describes one or more of your security groups.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSecurityGroups.html]
func (s *EC2Service) DescribeSecurityGroups(req *DescribeSecurityGroupsRequest) (result *DescribeSecurityGroupsResponse, err error) {

	result = new(DescribeSecurityGroupsResponse)
	err = s.wrapperSignAndDo("DescribeSecurityGroups", req, result)
	return
}

// Describes the specified attribute of the specified snapshot. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshotAttribute.html]
func (s *EC2Service) DescribeSnapshotAttribute(req *DescribeSnapshotAttributeRequest) (result *DescribeSnapshotAttributeResponse, err error) {

	result = new(DescribeSnapshotAttributeResponse)
	err = s.wrapperSignAndDo("DescribeSnapshotAttribute", req, result)
	return
}

// Describes one or more of the Amazon EBS snapshots available to you. Snapshots available
// to you include public snapshots available for any AWS account to launch, private snapshots
// you own, and private snapshots owned by another AWS account but for which you've been
// given explicit create volume permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html]
func (s *EC2Service) DescribeSnapshots(req *DescribeSnapshotsRequest) (result *DescribeSnapshotsResponse, err error) {

	result = new(DescribeSnapshotsResponse)
	err = s.wrapperSignAndDo("DescribeSnapshots", req, result)
	return
}

// Describes the datafeed for Spot Instances.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotDatafeedSubscription.html]
func (s *EC2Service) DescribeSpotDatafeedSubscription(req *DescribeSpotDatafeedSubscriptionRequest) (result *DescribeSpotDatafeedSubscriptionResponse, err error) {

	result = new(DescribeSpotDatafeedSubscriptionResponse)
	err = s.wrapperSignAndDo("DescribeSpotDatafeedSubscription", req, result)
	return
}

// Describes the Spot Instance requests that belong to your account. Spot Instances are instances
// that Amazon EC2 starts on your behalf when the maximum price that you specify exceeds the
// current Spot Price. Amazon EC2 periodically sets the Spot Price based on available Spot
// Instance capacity and current Spot Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotInstanceRequests.html]
func (s *EC2Service) DescribeSpotInstanceRequests(req *DescribeSpotInstanceRequestsRequest) (result *DescribeSpotInstanceRequestsResponse, err error) {

	result = new(DescribeSpotInstanceRequestsResponse)
	err = s.wrapperSignAndDo("DescribeSpotInstanceRequests", req, result)
	return
}

// Describes the Spot Price history. Spot Instances are instances that Amazon EC2 starts on your
// behalf when the maximum price that you specify exceeds the current Spot Price. Amazon EC2
// periodically sets the Spot Price based on available Spot Instance capacity and current Spot
// Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSpotPriceHistory.html]
func (s *EC2Service) DescribeSpotPriceHistory(req *DescribeSpotPriceHistoryRequest) (result *DescribeSpotPriceHistoryResponse, err error) {

	result = new(DescribeSpotPriceHistoryResponse)
	err = s.wrapperSignAndDo("DescribeSpotPriceHistory", req, result)
	return
}

// Describes one or more of your subnets.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSubnets.html]
func (s *EC2Service) DescribeSubnets(req *DescribeSubnetsRequest) (result *DescribeSubnetsResponse, err error) {

	result = new(DescribeSubnetsResponse)
	err = s.wrapperSignAndDo("DescribeSubnets", req, result)
	return
}

// Describes one or more of the tags for your EC2 resources.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeTags.html]
func (s *EC2Service) DescribeTags(req *DescribeTagsRequest) (result *DescribeTagsResponse, err error) {

	result = new(DescribeTagsResponse)
	err = s.wrapperSignAndDo("DescribeTags", req, result)
	return
}

// Describes the specified attribute of the specified volume. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeAttribute.html]
func (s *EC2Service) DescribeVolumeAttribute(req *DescribeVolumeAttributeRequest) (result *DescribeVolumeAttributeResponse, err error) {

	result = new(DescribeVolumeAttributeResponse)
	err = s.wrapperSignAndDo("DescribeVolumeAttribute", req, result)
	return
}

// Describes the specified Amazon EBS volumes.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumes.html]
func (s *EC2Service) DescribeVolumes(req *DescribeVolumesRequest) (result *DescribeVolumesResponse, err error) {

	result = new(DescribeVolumesResponse)
	err = s.wrapperSignAndDo("DescribeVolumes", req, result)
	return
}

// Describes the status of the specified volumes. Volume status provides the result of the
// checks performed on your volumes to determine events that can impair the performance of
// your volumes. The performance of a volume can be affected if an issue occurs on the volume's
// underlying host. If the volume's underlying host experiences a power outage or system issue,
// once the system is restored there could be data inconsistencies on the volume. Volume events
// notify you if this occurs. Volume actions notify you if any action needs to be taken in
// response to the event.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVolumeStatus.html]
func (s *EC2Service) DescribeVolumeStatus(req *DescribeVolumeStatusRequest) (result *DescribeVolumeStatusResponse, err error) {

	result = new(DescribeVolumeStatusResponse)
	err = s.wrapperSignAndDo("DescribeVolumeStatus", req, result)
	return
}

// Describes the specified attribute of the specified VPC. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcAttribute.html]
func (s *EC2Service) DescribeVpcAttribute(req *DescribeVpcAttributeRequest) (result *DescribeVpcAttributeResponse, err error) {

	result = new(DescribeVpcAttributeResponse)
	err = s.wrapperSignAndDo("DescribeVpcAttribute", req, result)
	return
}

// Describes one or more of your VPC peering connections.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcPeeringConnections.html]
func (s *EC2Service) DescribeVpcPeeringConnections(req *DescribeVpcPeeringConnectionsRequest) (result *DescribeVpcPeeringConnectionsResponse, err error) {

	result = new(DescribeVpcPeeringConnectionsResponse)
	err = s.wrapperSignAndDo("DescribeVpcPeeringConnections", req, result)
	return
}

// Describes one or more of your VPCs.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpcs.html]
func (s *EC2Service) DescribeVpcs(req *DescribeVpcsRequest) (result *DescribeVpcsResponse, err error) {

	result = new(DescribeVpcsResponse)
	err = s.wrapperSignAndDo("DescribeVpcs", req, result)
	return
}

// Describes one or more of your VPN connections.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnConnections.html]
func (s *EC2Service) DescribeVpnConnections(req *DescribeVpnConnectionsRequest) (result *DescribeVpnConnectionsResponse, err error) {

	result = new(DescribeVpnConnectionsResponse)
	err = s.wrapperSignAndDo("DescribeVpnConnections", req, result)
	return
}

// Describes one or more of your virtual private gateways.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeVpnGateways.html]
func (s *EC2Service) DescribeVpnGateways(req *DescribeVpnGatewaysRequest) (result *DescribeVpnGatewaysResponse, err error) {

	result = new(DescribeVpnGatewaysResponse)
	err = s.wrapperSignAndDo("DescribeVpnGateways", req, result)
	return
}

// Detaches an Internet gateway from a VPC, disabling connectivity between the Internet and the VPC.
// The VPC must not contain any running instances with Elastic IP addresses.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachInternetGateway.html]
func (s *EC2Service) DetachInternetGateway(req *DetachInternetGatewayRequest) (result *DetachInternetGatewayResponse, err error) {

	result = new(DetachInternetGatewayResponse)
	err = s.wrapperSignAndDo("DetachInternetGateway", req, result)
	return
}

// Detaches a network interface from an instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachNetworkInterface.html]
func (s *EC2Service) DetachNetworkInterface(req *DetachNetworkInterfaceRequest) (result *DetachNetworkInterfaceResponse, err error) {

	result = new(DetachNetworkInterfaceResponse)
	err = s.wrapperSignAndDo("DetachNetworkInterface", req, result)
	return
}

// Detaches an Amazon EBS volume from an instance. Make sure to unmount any file systems
// on the device within your operating system before detaching the volume. Failure to do
// so results in the volume being stuck in a busy state while detaching.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVolume.html]
func (s *EC2Service) DetachVolume(req *DetachVolumeRequest) (result *DetachVolumeResponse, err error) {

	result = new(DetachVolumeResponse)
	err = s.wrapperSignAndDo("DetachVolume", req, result)
	return
}

// Detaches a virtual private gateway from a VPC. You do this if you're planning to turn off
// the VPC and not use it anymore. You can confirm a virtual private gateway has been
// completely detached from a VPC by describing the virtual private gateway (any attachments
// to the virtual private gateway are also described).
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVpnGateway.html]
func (s *EC2Service) DetachVpnGateway(req *DetachVpnGatewayRequest) (result *DetachVpnGatewayResponse, err error) {

	result = new(DetachVpnGatewayResponse)
	err = s.wrapperSignAndDo("DetachVpnGateway", req, result)
	return
}

// Disables a virtual private gateway (VGW) from propagating routes to the routing tables of a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisableVgwRoutePropagation.html]
func (s *EC2Service) DisableVgwRoutePropagation(req *DisableVgwRoutePropagationRequest) (result *DisableVgwRoutePropagationResponse, err error) {

	result = new(DisableVgwRoutePropagationResponse)
	err = s.wrapperSignAndDo("DisableVgwRoutePropagation", req, result)
	return
}

// Disassociates an Elastic IP address from the instance or network interface it's associated with.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateAddress.html]
func (s *EC2Service) DisassociateAddress(req *DisassociateAddressRequest) (result *DisassociateAddressResponse, err error) {

	result = new(DisassociateAddressResponse)
	err = s.wrapperSignAndDo("DisassociateAddress", req, result)
	return
}

// Disassociates a subnet from a route table.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DisassociateRouteTable.html]
func (s *EC2Service) DisassociateRouteTable(req *DisassociateRouteTableRequest) (result *DisassociateRouteTableResponse, err error) {

	result = new(DisassociateRouteTableResponse)
	err = s.wrapperSignAndDo("DisassociateRouteTable", req, result)
	return
}

// Enables a virtual private gateway (VGW) to propagate routes to the routing tables of a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVgwRoutePropagation.html]
func (s *EC2Service) EnableVgwRoutePropagation(req *EnableVgwRoutePropagationRequest) (result *EnableVgwRoutePropagationResponse, err error) {

	result = new(EnableVgwRoutePropagationResponse)
	err = s.wrapperSignAndDo("EnableVgwRoutePropagation", req, result)
	return
}

// Enables I/O operations for a volume that had I/O operations disabled because
// the data on the volume was potentially inconsistent.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-EnableVolumeIO.html]
func (s *EC2Service) EnableVolumeIO(req *EnableVolumeIORequest) (result *EnableVolumeIOResponse, err error) {

	result = new(EnableVolumeIOResponse)
	err = s.wrapperSignAndDo("EnableVolumeIO", req, result)
	return
}

// Gets the console output for the specified instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetConsoleOutput.html]
func (s *EC2Service) GetConsoleOutput(req *GetConsoleOutputRequest) (result *GetConsoleOutputResponse, err error) {

	result = new(GetConsoleOutputResponse)
	err = s.wrapperSignAndDo("GetConsoleOutput", req, result)
	return
}

// Retrieves the encrypted administrator password for an instance running Windows.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-GetPasswordData.html]
func (s *EC2Service) GetPasswordData(req *GetPasswordDataRequest) (result *GetPasswordDataResponse, err error) {

	result = new(GetPasswordDataResponse)
	err = s.wrapperSignAndDo("GetPasswordData", req, result)
	return
}

// Creates an import instance task using metadata from the specified disk image.
// After importing the image, you then upload it using the ec2-import-volume command.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportInstance.html]
func (s *EC2Service) ImportInstance(req *ImportInstanceRequest) (result *ImportInstanceResponse, err error) {

	result = new(ImportInstanceResponse)
	err = s.wrapperSignAndDo("ImportInstance", req, result)
	return
}

// Imports the public key from an RSA key pair that you created with a third-party tool.
// Compare this with CreateKeyPair, in which AWS creates the key pair and gives the keys
// to you (AWS keeps a copy of the public key). With ImportKeyPair, you create the key
// pair and give AWS just the public key. The private key is never transferred between
// you and AWS.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportKeyPair.html]
func (s *EC2Service) ImportKeyPair(req *ImportKeyPairRequest) (result *ImportKeyPairResponse, err error) {

	result = new(ImportKeyPairResponse)
	err = s.wrapperSignAndDo("ImportKeyPair", req, result)
	return
}

// Creates an import volume task using metadata from the specified disk image. After importing
// the image, you then upload it using the ec2-import-volume command in the Amazon EC2
// command-line interface (CLI) tools.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ImportVolume.html]
func (s *EC2Service) ImportVolume(req *ImportVolumeRequest) (result *ImportVolumeResponse, err error) {

	result = new(ImportVolumeResponse)
	err = s.wrapperSignAndDo("ImportVolume", req, result)
	return
}

// Modifies the specified attribute of the specified AMI. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyImageAttribute.html]
func (s *EC2Service) ModifyImageAttribute(req *ModifyImageAttributeRequest) (result *ModifyImageAttributeResponse, err error) {

	result = new(ModifyImageAttributeResponse)
	err = s.wrapperSignAndDo("ModifyImageAttribute", req, result)
	return
}

// Modifies the specified attribute of the specified instance. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyInstanceAttribute.html]
func (s *EC2Service) ModifyInstanceAttribute(req *ModifyInstanceAttributeRequest) (result *ModifyInstanceAttributeResponse, err error) {

	result = new(ModifyInstanceAttributeResponse)
	err = s.wrapperSignAndDo("ModifyInstanceAttribute", req, result)
	return
}

// Modifies the specified network interface attribute. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyNetworkInterfaceAttribute.html]
func (s *EC2Service) ModifyNetworkInterfaceAttribute(req *ModifyNetworkInterfaceAttributeRequest) (result *ModifyNetworkInterfaceAttributeResponse, err error) {

	result = new(ModifyNetworkInterfaceAttributeResponse)
	err = s.wrapperSignAndDo("ModifyNetworkInterfaceAttribute", req, result)
	return
}

// Modifies the Availability Zone, instance count, instance type, or network platform
// (EC2-Classic or EC2-VPC) of your Reserved Instances. The Reserved Instances to be
// modified must be identical, except for Availability Zone, network platform,
// and instance type.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyReservedInstances.html]
func (s *EC2Service) ModifyReservedInstances(req *ModifyReservedInstancesRequest) (result *ModifyReservedInstancesResponse, err error) {

	result = new(ModifyReservedInstancesResponse)
	err = s.wrapperSignAndDo("ModifyReservedInstances", req, result)
	return
}

// Adds or removes permission settings for the specified snapshot. You may add or remove specified
// AWS account IDs from a snapshot's list of create volume permissions, but you cannot do both in
// a single API call. If you need to both add and remove account IDs for a snapshot, you must use
// multiple API calls.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySnapshotAttribute.html]
func (s *EC2Service) ModifySnapshotAttribute(req *ModifySnapshotAttributeRequest) (result *ModifySnapshotAttributeResponse, err error) {

	result = new(ModifySnapshotAttributeResponse)
	err = s.wrapperSignAndDo("ModifySnapshotAttribute", req, result)
	return
}

// Modifies a subnet attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifySubnetAttribute.html]
func (s *EC2Service) ModifySubnetAttribute(req *ModifySubnetAttributeRequest) (result *ModifySubnetAttributeResponse, err error) {

	result = new(ModifySubnetAttributeResponse)
	err = s.wrapperSignAndDo("ModifySubnetAttribute", req, result)
	return
}

// Modifies a volume attribute.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVolumeAttribute.html]
func (s *EC2Service) ModifyVolumeAttribute(req *ModifyVolumeAttributeRequest) (result *ModifyVolumeAttributeResponse, err error) {

	result = new(ModifyVolumeAttributeResponse)
	err = s.wrapperSignAndDo("ModifyVolumeAttribute", req, result)
	return
}

// Modifies the specified attribute of the specified VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ModifyVpcAttribute.html]
func (s *EC2Service) ModifyVpcAttribute(req *ModifyVpcAttributeRequest) (result *ModifyVpcAttributeResponse, err error) {

	result = new(ModifyVpcAttributeResponse)
	err = s.wrapperSignAndDo("ModifyVpcAttribute", req, result)
	return
}

// Enables monitoring for a running instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-MonitorInstances.html]
func (s *EC2Service) MonitorInstances(req *MonitorInstancesRequest) (result *MonitorInstancesResponse, err error) {

	result = new(MonitorInstancesResponse)
	err = s.wrapperSignAndDo("MonitorInstances", req, result)
	return
}

// Purchases a Reserved Instance for use with your account. With Amazon EC2 Reserved Instances,
// you obtain a capacity reservation for a certain instance configuration over a specified
// period of time. You pay a lower usage rate than with On-Demand instances for the time that
// you actually use the capacity reservation.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-PurchaseReservedInstancesOffering.html]
func (s *EC2Service) PurchaseReservedInstancesOffering(req *PurchaseReservedInstancesOfferingRequest) (result *PurchaseReservedInstancesOfferingResponse, err error) {

	result = new(PurchaseReservedInstancesOfferingResponse)
	err = s.wrapperSignAndDo("PurchaseReservedInstancesOffering", req, result)
	return
}

// Requests a reboot of one or more instances. This operation is asynchronous; it only queues a
// request to reboot the specified instances. The operation succeeds if the instances are valid
// and belong to you. Requests to reboot terminated instances are ignored.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RebootInstances.html]
func (s *EC2Service) RebootInstances(req *RebootInstancesRequest) (result *RebootInstancesResponse, err error) {

	result = new(RebootInstancesResponse)
	err = s.wrapperSignAndDo("RebootInstances", req, result)
	return
}

// Registers an AMI. When you're creating an AMI, this is the final step you must
// complete before you can launch an instance from the AMI.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RegisterImage.html]
func (s *EC2Service) RegisterImage(req *RegisterImageRequest) (result *RegisterImageResponse, err error) {

	result = new(RegisterImageResponse)
	err = s.wrapperSignAndDo("RegisterImage", req, result)
	return
}

// Rejects a VPC peering connection request. The VPC peering connection must be
// in the pending-acceptance state.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RejectVpcPeeringConnection.html]
func (s *EC2Service) RejectVpcPeeringConnection(req *RejectVpcPeeringConnectionRequest) (result *RejectVpcPeeringConnectionResponse, err error) {

	result = new(RejectVpcPeeringConnectionResponse)
	err = s.wrapperSignAndDo("RejectVpcPeeringConnection", req, result)
	return
}

// Releases the specified Elastic IP address.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReleaseAddress.html]
func (s *EC2Service) ReleaseAddress(req *ReleaseAddressRequest) (result *ReleaseAddressResponse, err error) {

	result = new(ReleaseAddressResponse)
	err = s.wrapperSignAndDo("ReleaseAddress", req, result)
	return
}

// Changes which network ACL a subnet is associated with. By default when you create
// a subnet, it's automatically associated with the default network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html]
func (s *EC2Service) ReplaceNetworkAclAssociation(req *ReplaceNetworkAclAssociationRequest) (result *ReplaceNetworkAclAssociationResponse, err error) {

	result = new(ReplaceNetworkAclAssociationResponse)
	err = s.wrapperSignAndDo("ReplaceNetworkAclAssociation", req, result)
	return
}

// Replaces an entry (rule) in a network ACL.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclEntry.html]
func (s *EC2Service) ReplaceNetworkAclEntry(req *ReplaceNetworkAclEntryRequest) (result *ReplaceNetworkAclEntryResponse, err error) {

	result = new(ReplaceNetworkAclEntryResponse)
	err = s.wrapperSignAndDo("ReplaceNetworkAclEntry", req, result)
	return
}

// Replaces an existing route within a route table in a VPC.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRoute.html]
func (s *EC2Service) ReplaceRoute(req *ReplaceRouteRequest) (result *ReplaceRouteResponse, err error) {

	result = new(ReplaceRouteResponse)
	err = s.wrapperSignAndDo("ReplaceRoute", req, result)
	return
}

// Changes the route table associated with a given subnet in a VPC. After the operation completes,
// the subnet uses the routes in the new route table it's associated with.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceRouteTableAssociation.html]
func (s *EC2Service) ReplaceRouteTableAssociation(req *ReplaceRouteTableAssociationRequest) (result *ReplaceRouteTableAssociationResponse, err error) {

	result = new(ReplaceRouteTableAssociationResponse)
	err = s.wrapperSignAndDo("ReplaceRouteTableAssociation", req, result)
	return
}

// Submits feedback about an instance's status. The instance must be in the running state.
// If your experience with the instance differs from the instance status returned by
// DescribeInstanceStatus, use ReportInstanceStatus to report your experience with the
// instance. Amazon EC2 collects this information to improve the accuracy of status checks.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReportInstanceStatus.html]
func (s *EC2Service) ReportInstanceStatus(req *ReportInstanceStatusRequest) (result *ReportInstanceStatusResponse, err error) {

	result = new(ReportInstanceStatusResponse)
	err = s.wrapperSignAndDo("ReportInstanceStatus", req, result)
	return
}

// Creates a Spot Instance request. Spot Instances are instances that Amazon EC2 starts on your
// behalf when the maximum price that you specify exceeds the current Spot Price. Amazon EC2
// periodically sets the Spot Price based on available Spot Instance capacity and current Spot
// Instance requests.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RequestSpotInstances.html]
func (s *EC2Service) RequestSpotInstances(req *RequestSpotInstancesRequest) (result *RequestSpotInstancesResponse, err error) {

	result = new(RequestSpotInstancesResponse)
	err = s.wrapperSignAndDo("RequestSpotInstances", req, result)
	return
}

// Resets an attribute of an AMI to its default value.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetImageAttribute.html]
func (s *EC2Service) ResetImageAttribute(req *ResetImageAttributeRequest) (result *ResetImageAttributeResponse, err error) {

	result = new(ResetImageAttributeResponse)
	err = s.wrapperSignAndDo("ResetImageAttribute", req, result)
	return
}

// Resets an attribute of an instance to its default value. To reset the kernel or RAM disk,
// the instance must be in a stopped state. To reset the SourceDestCheck, the instance can
// be either running or stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetInstanceAttribute.html]
func (s *EC2Service) ResetInstanceAttribute(req *ResetInstanceAttributeRequest) (result *ResetInstanceAttributeResponse, err error) {

	result = new(ResetInstanceAttributeResponse)
	err = s.wrapperSignAndDo("ResetInstanceAttribute", req, result)
	return
}

// Resets a network interface attribute. You can specify only one attribute at a time.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetNetworkInterfaceAttribute.html]
func (s *EC2Service) ResetNetworkInterfaceAttribute(req *ResetNetworkInterfaceAttributeRequest) (result *ResetNetworkInterfaceAttributeResponse, err error) {

	result = new(ResetNetworkInterfaceAttributeResponse)
	err = s.wrapperSignAndDo("ResetNetworkInterfaceAttribute", req, result)
	return
}

// Resets permission settings for the specified snapshot.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ResetSnapshotAttribute.html]
func (s *EC2Service) ResetSnapshotAttribute(req *ResetSnapshotAttributeRequest) (result *ResetSnapshotAttributeResponse, err error) {

	result = new(ResetSnapshotAttributeResponse)
	err = s.wrapperSignAndDo("ResetSnapshotAttribute", req, result)
	return
}

// Removes one or more egress rules from a security group for EC2-VPC. The values that you specify
// in the revoke request (for example, ports) must match the existing rule's values for the rule to be revoked.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupEgress.html]
func (s *EC2Service) RevokeSecurityGroupEgress(req *RevokeSecurityGroupEgressRequest) (result *RevokeSecurityGroupEgressResponse, err error) {

	result = new(RevokeSecurityGroupEgressResponse)
	err = s.wrapperSignAndDo("RevokeSecurityGroupEgress", req, result)
	return
}

// Removes one or more ingress rules from a security group. The values that you specify in the revoke request
// (for example, ports) must match the existing rule's values for the rule to be removed.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RevokeSecurityGroupIngress.html]
func (s *EC2Service) RevokeSecurityGroupIngress(req *RevokeSecurityGroupIngressRequest) (result *RevokeSecurityGroupIngressResponse, err error) {

	result = new(RevokeSecurityGroupIngressResponse)
	err = s.wrapperSignAndDo("RevokeSecurityGroupIngress", req, result)
	return
}

// Launches the specified number of instances using an AMI for which you have permissions.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-RunInstances.html]
func (s *EC2Service) RunInstances(req *RunInstancesRequest) (result *RunInstancesResponse, err error) {

	result = new(RunInstancesResponse)
	err = s.wrapperSignAndDo("RunInstances", req, result)
	return
}

// Starts an Amazon EBS-backed AMI that you've previously stopped.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StartInstances.html]
func (s *EC2Service) StartInstances(req *StartInstancesRequest) (result *StartInstancesResponse, err error) {

	result = new(StartInstancesResponse)
	err = s.wrapperSignAndDo("StartInstances", req, result)
	return
}

// Stops an Amazon EBS-backed instance. Each time you transition an instance from stopped to started,
// Amazon EC2 charges a full instance hour, even if transitions happen multiple times within a single hour.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-StopInstances.html]
func (s *EC2Service) StopInstances(req *StopInstancesRequest) (result *StopInstancesResponse, err error) {

	result = new(StopInstancesResponse)
	err = s.wrapperSignAndDo("StopInstances", req, result)
	return
}

// Shuts down one or more instances. This operation is idempotent; if you terminate an
// instance more than once, each call succeeds.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-TerminateInstances.html]
func (s *EC2Service) TerminateInstances(req *TerminateInstancesRequest) (result *TerminateInstancesResponse, err error) {

	result = new(TerminateInstancesResponse)
	err = s.wrapperSignAndDo("TerminateInstances", req, result)
	return
}

// Unassigns one or more secondary private IP addresses from a network interface.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnassignPrivateIpAddresses.html]
func (s *EC2Service) UnassignPrivateIpAddresses(req *UnassignPrivateIpAddressesRequest) (result *UnassignPrivateIpAddressesResponse, err error) {

	result = new(UnassignPrivateIpAddressesResponse)
	err = s.wrapperSignAndDo("UnassignPrivateIpAddresses", req, result)
	return
}

// Disables monitoring for a running instance.
// [http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-UnmonitorInstances.html]
func (s *EC2Service) UnmonitorInstances(req *UnmonitorInstancesRequest) (result *UnmonitorInstancesResponse, err error) {

	result = new(UnmonitorInstancesResponse)
	err = s.wrapperSignAndDo("UnmonitorInstances", req, result)
	return
}

// Creates a new EC2 Service.
func NewService(cred interfaces.IAWSCredentials) *EC2Service {

	return &EC2Service{cred, nil, "https://" + ServiceName + ".amazonaws.com/"}
}
