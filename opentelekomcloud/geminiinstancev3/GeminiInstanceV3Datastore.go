// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminiinstancev3


type GeminiInstanceV3Datastore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#engine GeminiInstanceV3#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#storage_engine GeminiInstanceV3#storage_engine}.
	StorageEngine *string `field:"required" json:"storageEngine" yaml:"storageEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#version GeminiInstanceV3#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

