// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmsresourcerecorderv1


type RmsResourceRecorderV1ObsChannel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/rms_resource_recorder_v1#bucket RmsResourceRecorderV1#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/rms_resource_recorder_v1#region RmsResourceRecorderV1#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/rms_resource_recorder_v1#bucket_prefix RmsResourceRecorderV1#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

