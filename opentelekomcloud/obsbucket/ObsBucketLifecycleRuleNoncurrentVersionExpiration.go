// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

