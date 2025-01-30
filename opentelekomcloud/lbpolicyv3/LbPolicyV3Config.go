// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpolicyv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbPolicyV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#action LbPolicyV3#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#listener_id LbPolicyV3#listener_id}.
	ListenerId *string `field:"required" json:"listenerId" yaml:"listenerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#description LbPolicyV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// fixed_response_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#fixed_response_config LbPolicyV3#fixed_response_config}
	FixedResponseConfig *LbPolicyV3FixedResponseConfig `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#id LbPolicyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#name LbPolicyV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#position LbPolicyV3#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#priority LbPolicyV3#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#project_id LbPolicyV3#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#redirect_listener_id LbPolicyV3#redirect_listener_id}.
	RedirectListenerId *string `field:"optional" json:"redirectListenerId" yaml:"redirectListenerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#redirect_pool_id LbPolicyV3#redirect_pool_id}.
	RedirectPoolId *string `field:"optional" json:"redirectPoolId" yaml:"redirectPoolId"`
	// redirect_pools_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#redirect_pools_config LbPolicyV3#redirect_pools_config}
	RedirectPoolsConfig interface{} `field:"optional" json:"redirectPoolsConfig" yaml:"redirectPoolsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#redirect_url LbPolicyV3#redirect_url}.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// redirect_url_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#redirect_url_config LbPolicyV3#redirect_url_config}
	RedirectUrlConfig *LbPolicyV3RedirectUrlConfig `field:"optional" json:"redirectUrlConfig" yaml:"redirectUrlConfig"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/lb_policy_v3#rules LbPolicyV3#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

