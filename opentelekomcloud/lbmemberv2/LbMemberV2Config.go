// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbmemberv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbMemberV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#address LbMemberV2#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#pool_id LbMemberV2#pool_id}.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#protocol_port LbMemberV2#protocol_port}.
	ProtocolPort *float64 `field:"required" json:"protocolPort" yaml:"protocolPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#subnet_id LbMemberV2#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#admin_state_up LbMemberV2#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#id LbMemberV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#name LbMemberV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#region LbMemberV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#tenant_id LbMemberV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#timeouts LbMemberV2#timeouts}
	Timeouts *LbMemberV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_member_v2#weight LbMemberV2#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

