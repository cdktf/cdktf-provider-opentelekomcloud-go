package networkingsubnetv2


type NetworkingSubnetV2AllocationPools struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/networking_subnet_v2#end NetworkingSubnetV2#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/networking_subnet_v2#start NetworkingSubnetV2#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

