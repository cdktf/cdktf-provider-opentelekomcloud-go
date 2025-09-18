// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1


type AsmServiceMeshV1Clusters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#cluster_id AsmServiceMeshV1#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#installation_nodes AsmServiceMeshV1#installation_nodes}.
	InstallationNodes *[]*string `field:"required" json:"installationNodes" yaml:"installationNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#injection_namespaces AsmServiceMeshV1#injection_namespaces}.
	InjectionNamespaces *[]*string `field:"optional" json:"injectionNamespaces" yaml:"injectionNamespaces"`
}

