// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asgroupv1


type AsGroupV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_group_v1#create AsGroupV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_group_v1#delete AsGroupV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

