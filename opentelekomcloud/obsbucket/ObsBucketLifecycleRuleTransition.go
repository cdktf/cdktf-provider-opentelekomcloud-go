// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLifecycleRuleTransition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/obs_bucket#storage_class ObsBucket#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}

