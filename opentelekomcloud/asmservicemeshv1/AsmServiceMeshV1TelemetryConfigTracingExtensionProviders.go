// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1


type AsmServiceMeshV1TelemetryConfigTracingExtensionProviders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/asm_service_mesh_v1#name AsmServiceMeshV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/asm_service_mesh_v1#zipkin_service_addr AsmServiceMeshV1#zipkin_service_addr}.
	ZipkinServiceAddr *string `field:"optional" json:"zipkinServiceAddr" yaml:"zipkinServiceAddr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/asm_service_mesh_v1#zipkin_service_port AsmServiceMeshV1#zipkin_service_port}.
	ZipkinServicePort *float64 `field:"optional" json:"zipkinServicePort" yaml:"zipkinServicePort"`
}

