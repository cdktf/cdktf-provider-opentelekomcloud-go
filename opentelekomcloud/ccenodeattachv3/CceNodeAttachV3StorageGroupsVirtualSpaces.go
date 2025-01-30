// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3


type CceNodeAttachV3StorageGroupsVirtualSpaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/cce_node_attach_v3#name CceNodeAttachV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/cce_node_attach_v3#size CceNodeAttachV3#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/cce_node_attach_v3#lvm_lv_type CceNodeAttachV3#lvm_lv_type}.
	LvmLvType *string `field:"optional" json:"lvmLvType" yaml:"lvmLvType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/cce_node_attach_v3#lvm_path CceNodeAttachV3#lvm_path}.
	LvmPath *string `field:"optional" json:"lvmPath" yaml:"lvmPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/cce_node_attach_v3#runtime_lv_type CceNodeAttachV3#runtime_lv_type}.
	RuntimeLvType *string `field:"optional" json:"runtimeLvType" yaml:"runtimeLvType"`
}

