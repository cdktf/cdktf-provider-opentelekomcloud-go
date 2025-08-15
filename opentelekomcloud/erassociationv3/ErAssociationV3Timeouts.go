// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erassociationv3


type ErAssociationV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/er_association_v3#create ErAssociationV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/er_association_v3#delete ErAssociationV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

