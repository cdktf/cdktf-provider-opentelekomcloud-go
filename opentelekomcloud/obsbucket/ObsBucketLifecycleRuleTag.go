// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketLifecycleRuleTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/obs_bucket#key ObsBucket#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/obs_bucket#value ObsBucket#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

