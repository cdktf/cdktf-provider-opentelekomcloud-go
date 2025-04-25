// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsbackupv3


type DdsBackupV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dds_backup_v3#create DdsBackupV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dds_backup_v3#delete DdsBackupV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

