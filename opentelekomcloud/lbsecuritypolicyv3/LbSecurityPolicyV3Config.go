// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbsecuritypolicyv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbSecurityPolicyV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/lb_security_policy_v3#ciphers LbSecurityPolicyV3#ciphers}.
	Ciphers *[]*string `field:"required" json:"ciphers" yaml:"ciphers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/lb_security_policy_v3#protocols LbSecurityPolicyV3#protocols}.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/lb_security_policy_v3#description LbSecurityPolicyV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/lb_security_policy_v3#id LbSecurityPolicyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/lb_security_policy_v3#name LbSecurityPolicyV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

