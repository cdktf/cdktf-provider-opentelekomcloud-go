// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprotocolv3


type IdentityProtocolV3Metadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/identity_protocol_v3#domain_id IdentityProtocolV3#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/identity_protocol_v3#metadata IdentityProtocolV3#metadata}.
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/identity_protocol_v3#xaccount_type IdentityProtocolV3#xaccount_type}.
	XaccountType *string `field:"optional" json:"xaccountType" yaml:"xaccountType"`
}

