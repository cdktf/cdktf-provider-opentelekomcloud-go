// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csssnapshotconfigurationv1


type CssSnapshotConfigurationV1Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/css_snapshot_configuration_v1#agency CssSnapshotConfigurationV1#agency}.
	Agency *string `field:"required" json:"agency" yaml:"agency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/css_snapshot_configuration_v1#base_path CssSnapshotConfigurationV1#base_path}.
	BasePath *string `field:"required" json:"basePath" yaml:"basePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/css_snapshot_configuration_v1#bucket CssSnapshotConfigurationV1#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/css_snapshot_configuration_v1#kms_id CssSnapshotConfigurationV1#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}

