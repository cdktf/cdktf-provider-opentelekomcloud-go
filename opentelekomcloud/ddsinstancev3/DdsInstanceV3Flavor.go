// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsinstancev3


type DdsInstanceV3Flavor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/dds_instance_v3#num DdsInstanceV3#num}.
	Num *float64 `field:"required" json:"num" yaml:"num"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/dds_instance_v3#spec_code DdsInstanceV3#spec_code}.
	SpecCode *string `field:"required" json:"specCode" yaml:"specCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/dds_instance_v3#type DdsInstanceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/dds_instance_v3#size DdsInstanceV3#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/dds_instance_v3#storage DdsInstanceV3#storage}.
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

