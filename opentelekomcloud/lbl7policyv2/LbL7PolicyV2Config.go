// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbl7policyv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbL7PolicyV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#action LbL7PolicyV2#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#listener_id LbL7PolicyV2#listener_id}.
	ListenerId *string `field:"required" json:"listenerId" yaml:"listenerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#admin_state_up LbL7PolicyV2#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#description LbL7PolicyV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#id LbL7PolicyV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#name LbL7PolicyV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#position LbL7PolicyV2#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#redirect_listener_id LbL7PolicyV2#redirect_listener_id}.
	RedirectListenerId *string `field:"optional" json:"redirectListenerId" yaml:"redirectListenerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#redirect_pool_id LbL7PolicyV2#redirect_pool_id}.
	RedirectPoolId *string `field:"optional" json:"redirectPoolId" yaml:"redirectPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#region LbL7PolicyV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#tenant_id LbL7PolicyV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/lb_l7policy_v2#timeouts LbL7PolicyV2#timeouts}
	Timeouts *LbL7PolicyV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

