// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLifecycleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#enabled ObsBucket#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#name ObsBucket#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#expiration ObsBucket#expiration}
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#noncurrent_version_expiration ObsBucket#noncurrent_version_expiration}
	NoncurrentVersionExpiration interface{} `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// noncurrent_version_transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#noncurrent_version_transition ObsBucket#noncurrent_version_transition}
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#prefix ObsBucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/obs_bucket#transition ObsBucket#transition}
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

