// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type NetworkingSubnetV2HostRoutes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_subnet_v2#destination_cidr NetworkingSubnetV2#destination_cidr}.
	DestinationCidr *string `field:"required" json:"destinationCidr" yaml:"destinationCidr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_subnet_v2#next_hop NetworkingSubnetV2#next_hop}.
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
}

