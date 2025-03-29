// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sfsshareaccessrulesv2


type SfsShareAccessRulesV2AccessRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/sfs_share_access_rules_v2#access_level SfsShareAccessRulesV2#access_level}.
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/sfs_share_access_rules_v2#access_to SfsShareAccessRulesV2#access_to}.
	AccessTo *string `field:"required" json:"accessTo" yaml:"accessTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/sfs_share_access_rules_v2#access_type SfsShareAccessRulesV2#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
}

