// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1


type CfwAclRuleV1SourceRegionListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_acl_rule_v1#region_id CfwAclRuleV1#region_id}.
	RegionId *string `field:"optional" json:"regionId" yaml:"regionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cfw_acl_rule_v1#region_type CfwAclRuleV1#region_type}.
	RegionType *float64 `field:"optional" json:"regionType" yaml:"regionType"`
}

