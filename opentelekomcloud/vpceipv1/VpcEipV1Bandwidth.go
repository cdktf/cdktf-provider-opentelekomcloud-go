// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpceipv1


type VpcEipV1Bandwidth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpc_eip_v1#name VpcEipV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpc_eip_v1#share_type VpcEipV1#share_type}.
	ShareType *string `field:"required" json:"shareType" yaml:"shareType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpc_eip_v1#size VpcEipV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpc_eip_v1#charge_mode VpcEipV1#charge_mode}.
	ChargeMode *string `field:"optional" json:"chargeMode" yaml:"chargeMode"`
}

