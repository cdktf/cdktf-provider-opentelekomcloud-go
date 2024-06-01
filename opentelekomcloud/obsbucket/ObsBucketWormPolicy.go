// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketWormPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/obs_bucket#years ObsBucket#years}.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

