// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asmservicemeshv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsmServiceMeshV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#clusters AsmServiceMeshV1#clusters}
	Clusters interface{} `field:"required" json:"clusters" yaml:"clusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#name AsmServiceMeshV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#type AsmServiceMeshV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#version AsmServiceMeshV1#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#id AsmServiceMeshV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#ipv6_enable AsmServiceMeshV1#ipv6_enable}.
	Ipv6Enable interface{} `field:"optional" json:"ipv6Enable" yaml:"ipv6Enable"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#proxy_config AsmServiceMeshV1#proxy_config}
	ProxyConfig *AsmServiceMeshV1ProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// telemetry_config_tracing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#telemetry_config_tracing AsmServiceMeshV1#telemetry_config_tracing}
	TelemetryConfigTracing *AsmServiceMeshV1TelemetryConfigTracing `field:"optional" json:"telemetryConfigTracing" yaml:"telemetryConfigTracing"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/asm_service_mesh_v1#timeouts AsmServiceMeshV1#timeouts}
	Timeouts *AsmServiceMeshV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

