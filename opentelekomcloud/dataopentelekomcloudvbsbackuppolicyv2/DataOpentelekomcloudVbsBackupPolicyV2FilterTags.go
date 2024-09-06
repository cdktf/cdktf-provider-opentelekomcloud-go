// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudvbsbackuppolicyv2


type DataOpentelekomcloudVbsBackupPolicyV2FilterTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/data-sources/vbs_backup_policy_v2#key DataOpentelekomcloudVbsBackupPolicyV2#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/data-sources/vbs_backup_policy_v2#values DataOpentelekomcloudVbsBackupPolicyV2#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

