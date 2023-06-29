package ccenodepoolv3


type CceNodePoolV3Taints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.2/docs/resources/cce_node_pool_v3#effect CceNodePoolV3#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.2/docs/resources/cce_node_pool_v3#key CceNodePoolV3#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.2/docs/resources/cce_node_pool_v3#value CceNodePoolV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

