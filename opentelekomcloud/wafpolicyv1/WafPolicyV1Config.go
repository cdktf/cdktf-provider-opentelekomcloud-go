// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafpolicyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafPolicyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#name WafPolicyV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#action WafPolicyV1#action}
	Action *WafPolicyV1Action `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#full_detection WafPolicyV1#full_detection}.
	FullDetection interface{} `field:"optional" json:"fullDetection" yaml:"fullDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#hosts WafPolicyV1#hosts}.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#id WafPolicyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#level WafPolicyV1#level}.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#options WafPolicyV1#options}
	Options *WafPolicyV1Options `field:"optional" json:"options" yaml:"options"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#timeouts WafPolicyV1#timeouts}
	Timeouts *WafPolicyV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

