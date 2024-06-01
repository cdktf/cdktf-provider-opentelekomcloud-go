// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2FuncMounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_function_v2#local_mount_path FgsFunctionV2#local_mount_path}.
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_function_v2#mount_resource FgsFunctionV2#mount_resource}.
	MountResource *string `field:"required" json:"mountResource" yaml:"mountResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_function_v2#mount_share_path FgsFunctionV2#mount_share_path}.
	MountSharePath *string `field:"required" json:"mountSharePath" yaml:"mountSharePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_function_v2#mount_type FgsFunctionV2#mount_type}.
	MountType *string `field:"required" json:"mountType" yaml:"mountType"`
}

