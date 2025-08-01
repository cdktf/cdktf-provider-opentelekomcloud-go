// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityaclv3


type IdentityAclV3IpCidrs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/identity_acl_v3#cidr IdentityAclV3#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/identity_acl_v3#description IdentityAclV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

