// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csbsbackuppolicyv1


type CsbsBackupPolicyV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/csbs_backup_policy_v1#create CsbsBackupPolicyV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/csbs_backup_policy_v1#delete CsbsBackupPolicyV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

