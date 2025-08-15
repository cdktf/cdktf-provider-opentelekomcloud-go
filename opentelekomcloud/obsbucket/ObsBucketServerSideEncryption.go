// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketServerSideEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/obs_bucket#algorithm ObsBucket#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/obs_bucket#kms_key_id ObsBucket#kms_key_id}.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/obs_bucket#kms_project_id ObsBucket#kms_project_id}.
	KmsProjectId *string `field:"optional" json:"kmsProjectId" yaml:"kmsProjectId"`
}

