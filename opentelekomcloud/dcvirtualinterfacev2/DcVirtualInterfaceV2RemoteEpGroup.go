// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacev2


type DcVirtualInterfaceV2RemoteEpGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_interface_v2#endpoints DcVirtualInterfaceV2#endpoints}.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_interface_v2#description DcVirtualInterfaceV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_interface_v2#name DcVirtualInterfaceV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_interface_v2#project_id DcVirtualInterfaceV2#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_interface_v2#type DcVirtualInterfaceV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

