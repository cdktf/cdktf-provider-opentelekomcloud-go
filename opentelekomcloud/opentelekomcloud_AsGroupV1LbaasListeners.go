// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type AsGroupV1LbaasListeners struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_group_v1#pool_id AsGroupV1#pool_id}.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_group_v1#protocol_port AsGroupV1#protocol_port}.
	ProtocolPort *float64 `field:"required" json:"protocolPort" yaml:"protocolPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_group_v1#weight AsGroupV1#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

