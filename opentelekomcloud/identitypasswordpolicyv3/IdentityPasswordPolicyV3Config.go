// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitypasswordpolicyv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPasswordPolicyV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#id IdentityPasswordPolicyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#maximum_consecutive_identical_chars IdentityPasswordPolicyV3#maximum_consecutive_identical_chars}.
	MaximumConsecutiveIdenticalChars *float64 `field:"optional" json:"maximumConsecutiveIdenticalChars" yaml:"maximumConsecutiveIdenticalChars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#minimum_password_age IdentityPasswordPolicyV3#minimum_password_age}.
	MinimumPasswordAge *float64 `field:"optional" json:"minimumPasswordAge" yaml:"minimumPasswordAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#minimum_password_length IdentityPasswordPolicyV3#minimum_password_length}.
	MinimumPasswordLength *float64 `field:"optional" json:"minimumPasswordLength" yaml:"minimumPasswordLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#number_of_recent_passwords_disallowed IdentityPasswordPolicyV3#number_of_recent_passwords_disallowed}.
	NumberOfRecentPasswordsDisallowed *float64 `field:"optional" json:"numberOfRecentPasswordsDisallowed" yaml:"numberOfRecentPasswordsDisallowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#password_char_combination IdentityPasswordPolicyV3#password_char_combination}.
	PasswordCharCombination *float64 `field:"optional" json:"passwordCharCombination" yaml:"passwordCharCombination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#password_not_username_or_invert IdentityPasswordPolicyV3#password_not_username_or_invert}.
	PasswordNotUsernameOrInvert interface{} `field:"optional" json:"passwordNotUsernameOrInvert" yaml:"passwordNotUsernameOrInvert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/identity_password_policy_v3#password_validity_period IdentityPasswordPolicyV3#password_validity_period}.
	PasswordValidityPeriod *float64 `field:"optional" json:"passwordValidityPeriod" yaml:"passwordValidityPeriod"`
}

