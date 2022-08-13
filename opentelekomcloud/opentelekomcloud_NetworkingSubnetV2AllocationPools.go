// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type NetworkingSubnetV2AllocationPools struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_subnet_v2#end NetworkingSubnetV2#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_subnet_v2#start NetworkingSubnetV2#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

