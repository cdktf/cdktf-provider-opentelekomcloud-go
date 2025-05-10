// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaddressgroupmemberv1


type CfwAddressGroupMemberV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_address_group_member_v1#create CfwAddressGroupMemberV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_address_group_member_v1#delete CfwAddressGroupMemberV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_address_group_member_v1#update CfwAddressGroupMemberV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

