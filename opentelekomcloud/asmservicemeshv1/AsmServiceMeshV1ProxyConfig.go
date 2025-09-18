// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1


type AsmServiceMeshV1ProxyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#exclude_inbound_ports AsmServiceMeshV1#exclude_inbound_ports}.
	ExcludeInboundPorts *string `field:"optional" json:"excludeInboundPorts" yaml:"excludeInboundPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#exclude_ip_ranges AsmServiceMeshV1#exclude_ip_ranges}.
	ExcludeIpRanges *string `field:"optional" json:"excludeIpRanges" yaml:"excludeIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#exclude_outbound_ports AsmServiceMeshV1#exclude_outbound_ports}.
	ExcludeOutboundPorts *string `field:"optional" json:"excludeOutboundPorts" yaml:"excludeOutboundPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#include_inbound_ports AsmServiceMeshV1#include_inbound_ports}.
	IncludeInboundPorts *string `field:"optional" json:"includeInboundPorts" yaml:"includeInboundPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#include_ip_ranges AsmServiceMeshV1#include_ip_ranges}.
	IncludeIpRanges *string `field:"optional" json:"includeIpRanges" yaml:"includeIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#include_outbound_ports AsmServiceMeshV1#include_outbound_ports}.
	IncludeOutboundPorts *string `field:"optional" json:"includeOutboundPorts" yaml:"includeOutboundPorts"`
}

