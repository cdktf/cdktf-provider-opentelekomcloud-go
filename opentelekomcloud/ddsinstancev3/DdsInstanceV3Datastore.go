// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsinstancev3


type DdsInstanceV3Datastore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/dds_instance_v3#type DdsInstanceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/dds_instance_v3#version DdsInstanceV3#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/dds_instance_v3#storage_engine DdsInstanceV3#storage_engine}.
	StorageEngine *string `field:"optional" json:"storageEngine" yaml:"storageEngine"`
}

