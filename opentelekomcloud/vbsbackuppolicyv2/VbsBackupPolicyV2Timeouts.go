// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vbsbackuppolicyv2


type VbsBackupPolicyV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/vbs_backup_policy_v2#create VbsBackupPolicyV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/vbs_backup_policy_v2#delete VbsBackupPolicyV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

