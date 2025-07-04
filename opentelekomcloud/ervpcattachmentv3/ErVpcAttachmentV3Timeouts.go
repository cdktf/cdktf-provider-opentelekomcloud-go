// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ervpcattachmentv3


type ErVpcAttachmentV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/er_vpc_attachment_v3#create ErVpcAttachmentV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/er_vpc_attachment_v3#delete ErVpcAttachmentV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/er_vpc_attachment_v3#update ErVpcAttachmentV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

