// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLogging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/obs_bucket#target_bucket ObsBucket#target_bucket}.
	TargetBucket *string `field:"required" json:"targetBucket" yaml:"targetBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/obs_bucket#agency ObsBucket#agency}.
	Agency *string `field:"optional" json:"agency" yaml:"agency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/obs_bucket#target_prefix ObsBucket#target_prefix}.
	TargetPrefix *string `field:"optional" json:"targetPrefix" yaml:"targetPrefix"`
}

