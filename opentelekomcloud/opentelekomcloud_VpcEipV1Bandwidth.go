// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type VpcEipV1Bandwidth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_eip_v1#name VpcEipV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_eip_v1#share_type VpcEipV1#share_type}.
	ShareType *string `field:"required" json:"shareType" yaml:"shareType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_eip_v1#size VpcEipV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_eip_v1#charge_mode VpcEipV1#charge_mode}.
	ChargeMode *string `field:"optional" json:"chargeMode" yaml:"chargeMode"`
}

