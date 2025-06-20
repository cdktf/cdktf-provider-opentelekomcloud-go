// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityaclv3


type IdentityAclV3IpRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/identity_acl_v3#range IdentityAclV3#range}.
	Range *string `field:"required" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/identity_acl_v3#description IdentityAclV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

