// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpolicyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedPolicyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#name WafDedicatedPolicyV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#deep_inspection WafDedicatedPolicyV1#deep_inspection}.
	DeepInspection interface{} `field:"optional" json:"deepInspection" yaml:"deepInspection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#full_detection WafDedicatedPolicyV1#full_detection}.
	FullDetection interface{} `field:"optional" json:"fullDetection" yaml:"fullDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#header_inspection WafDedicatedPolicyV1#header_inspection}.
	HeaderInspection interface{} `field:"optional" json:"headerInspection" yaml:"headerInspection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#id WafDedicatedPolicyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#level WafDedicatedPolicyV1#level}.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#options WafDedicatedPolicyV1#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#protection_mode WafDedicatedPolicyV1#protection_mode}.
	ProtectionMode *string `field:"optional" json:"protectionMode" yaml:"protectionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#region WafDedicatedPolicyV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#shiro_decryption_check WafDedicatedPolicyV1#shiro_decryption_check}.
	ShiroDecryptionCheck interface{} `field:"optional" json:"shiroDecryptionCheck" yaml:"shiroDecryptionCheck"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/waf_dedicated_policy_v1#timeouts WafDedicatedPolicyV1#timeouts}
	Timeouts *WafDedicatedPolicyV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

