// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/s3_bucket#days S3Bucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

