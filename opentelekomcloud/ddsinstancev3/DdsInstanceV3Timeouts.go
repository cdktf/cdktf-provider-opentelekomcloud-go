// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsinstancev3


type DdsInstanceV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dds_instance_v3#create DdsInstanceV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dds_instance_v3#delete DdsInstanceV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dds_instance_v3#update DdsInstanceV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

