// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type CceNodePoolV3DataVolumes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_pool_v3#size CceNodePoolV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_pool_v3#volumetype CceNodePoolV3#volumetype}.
	Volumetype *string `field:"required" json:"volumetype" yaml:"volumetype"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_pool_v3#extend_param CceNodePoolV3#extend_param}.
	ExtendParam *string `field:"optional" json:"extendParam" yaml:"extendParam"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_pool_v3#kms_id CceNodePoolV3#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}
