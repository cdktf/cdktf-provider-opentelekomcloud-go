// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwservicegroupmemberv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwServiceGroupMemberV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#dest_port CfwServiceGroupMemberV1#dest_port}.
	DestPort *string `field:"required" json:"destPort" yaml:"destPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#protocol CfwServiceGroupMemberV1#protocol}.
	Protocol *float64 `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#set_id CfwServiceGroupMemberV1#set_id}.
	SetId *string `field:"required" json:"setId" yaml:"setId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#source_port CfwServiceGroupMemberV1#source_port}.
	SourcePort *string `field:"required" json:"sourcePort" yaml:"sourcePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#description CfwServiceGroupMemberV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/cfw_service_group_member_v1#timeouts CfwServiceGroupMemberV1#timeouts}
	Timeouts *CfwServiceGroupMemberV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

