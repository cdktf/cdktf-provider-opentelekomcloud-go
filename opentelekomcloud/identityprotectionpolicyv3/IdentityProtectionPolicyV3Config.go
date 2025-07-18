// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprotectionpolicyv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityProtectionPolicyV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_protection_policy_v3#enable_operation_protection_policy IdentityProtectionPolicyV3#enable_operation_protection_policy}.
	EnableOperationProtectionPolicy interface{} `field:"optional" json:"enableOperationProtectionPolicy" yaml:"enableOperationProtectionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_protection_policy_v3#id IdentityProtectionPolicyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// self_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_protection_policy_v3#self_management IdentityProtectionPolicyV3#self_management}
	SelfManagement *IdentityProtectionPolicyV3SelfManagement `field:"optional" json:"selfManagement" yaml:"selfManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_protection_policy_v3#verification_email IdentityProtectionPolicyV3#verification_email}.
	VerificationEmail *string `field:"optional" json:"verificationEmail" yaml:"verificationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/identity_protection_policy_v3#verification_mobile IdentityProtectionPolicyV3#verification_mobile}.
	VerificationMobile *string `field:"optional" json:"verificationMobile" yaml:"verificationMobile"`
}

