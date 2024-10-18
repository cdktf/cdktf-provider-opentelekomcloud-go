// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vbsbackupv2


type VbsBackupV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vbs_backup_v2#create VbsBackupV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vbs_backup_v2#delete VbsBackupV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

