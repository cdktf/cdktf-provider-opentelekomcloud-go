// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityuserv3


type IdentityUserV3LoginProtection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/identity_user_v3#enabled IdentityUserV3#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/identity_user_v3#verification_method IdentityUserV3#verification_method}.
	VerificationMethod *string `field:"required" json:"verificationMethod" yaml:"verificationMethod"`
}

