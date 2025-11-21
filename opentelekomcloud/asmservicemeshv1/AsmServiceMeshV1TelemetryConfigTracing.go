// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1


type AsmServiceMeshV1TelemetryConfigTracing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/asm_service_mesh_v1#default_providers AsmServiceMeshV1#default_providers}.
	DefaultProviders *[]*string `field:"optional" json:"defaultProviders" yaml:"defaultProviders"`
	// extension_providers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/asm_service_mesh_v1#extension_providers AsmServiceMeshV1#extension_providers}
	ExtensionProviders interface{} `field:"optional" json:"extensionProviders" yaml:"extensionProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/asm_service_mesh_v1#random_sampling_percentage AsmServiceMeshV1#random_sampling_percentage}.
	RandomSamplingPercentage *float64 `field:"optional" json:"randomSamplingPercentage" yaml:"randomSamplingPercentage"`
}

