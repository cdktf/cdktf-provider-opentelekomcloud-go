// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwdomainnamegroupv1


type CfwDomainNameGroupV1DomainNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_domain_name_group_v1#domain_name CfwDomainNameGroupV1#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_domain_name_group_v1#description CfwDomainNameGroupV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

