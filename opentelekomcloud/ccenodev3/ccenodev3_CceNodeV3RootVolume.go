package ccenodev3


type CceNodeV3RootVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3#size CceNodeV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3#volumetype CceNodeV3#volumetype}.
	Volumetype *string `field:"required" json:"volumetype" yaml:"volumetype"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3#extend_param CceNodeV3#extend_param}.
	ExtendParam *string `field:"optional" json:"extendParam" yaml:"extendParam"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3#kms_id CceNodeV3#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}
