// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1


type AsmServiceMeshV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/asm_service_mesh_v1#create AsmServiceMeshV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/asm_service_mesh_v1#delete AsmServiceMeshV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

