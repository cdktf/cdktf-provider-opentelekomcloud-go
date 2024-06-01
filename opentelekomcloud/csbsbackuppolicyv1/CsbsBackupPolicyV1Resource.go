// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csbsbackuppolicyv1


type CsbsBackupPolicyV1Resource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/csbs_backup_policy_v1#id CsbsBackupPolicyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/csbs_backup_policy_v1#name CsbsBackupPolicyV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/csbs_backup_policy_v1#type CsbsBackupPolicyV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

