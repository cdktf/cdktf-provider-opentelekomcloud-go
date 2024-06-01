// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csssnapshotconfigurationv1


type CssSnapshotConfigurationV1CreationPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/css_snapshot_configuration_v1#enable CssSnapshotConfigurationV1#enable}.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/css_snapshot_configuration_v1#keepday CssSnapshotConfigurationV1#keepday}.
	Keepday *float64 `field:"required" json:"keepday" yaml:"keepday"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/css_snapshot_configuration_v1#period CssSnapshotConfigurationV1#period}.
	Period *string `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/css_snapshot_configuration_v1#prefix CssSnapshotConfigurationV1#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/css_snapshot_configuration_v1#delete_auto CssSnapshotConfigurationV1#delete_auto}.
	DeleteAuto interface{} `field:"optional" json:"deleteAuto" yaml:"deleteAuto"`
}

