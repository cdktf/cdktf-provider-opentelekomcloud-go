// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLifecycleRuleExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

