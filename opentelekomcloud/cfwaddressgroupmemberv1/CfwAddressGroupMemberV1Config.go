// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaddressgroupmemberv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwAddressGroupMemberV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/cfw_address_group_member_v1#address CfwAddressGroupMemberV1#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/cfw_address_group_member_v1#set_id CfwAddressGroupMemberV1#set_id}.
	SetId *string `field:"required" json:"setId" yaml:"setId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/cfw_address_group_member_v1#address_type CfwAddressGroupMemberV1#address_type}.
	AddressType *float64 `field:"optional" json:"addressType" yaml:"addressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/cfw_address_group_member_v1#description CfwAddressGroupMemberV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/cfw_address_group_member_v1#timeouts CfwAddressGroupMemberV1#timeouts}
	Timeouts *CfwAddressGroupMemberV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

