package ccenodev3


type CceNodeV3DataVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/cce_node_v3#size CceNodeV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/cce_node_v3#volumetype CceNodeV3#volumetype}.
	Volumetype *string `field:"required" json:"volumetype" yaml:"volumetype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/cce_node_v3#extend_param CceNodeV3#extend_param}.
	ExtendParam *string `field:"optional" json:"extendParam" yaml:"extendParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/cce_node_v3#extend_params CceNodeV3#extend_params}.
	ExtendParams *map[string]*string `field:"optional" json:"extendParams" yaml:"extendParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/cce_node_v3#kms_id CceNodeV3#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}

