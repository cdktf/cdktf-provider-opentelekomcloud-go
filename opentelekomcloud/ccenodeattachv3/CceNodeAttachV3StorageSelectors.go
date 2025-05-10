// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3


type CceNodeAttachV3StorageSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#name CceNodeAttachV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#match_label_count CceNodeAttachV3#match_label_count}.
	MatchLabelCount *string `field:"optional" json:"matchLabelCount" yaml:"matchLabelCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#match_label_metadata_cmkid CceNodeAttachV3#match_label_metadata_cmkid}.
	MatchLabelMetadataCmkid *string `field:"optional" json:"matchLabelMetadataCmkid" yaml:"matchLabelMetadataCmkid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#match_label_metadata_encrypted CceNodeAttachV3#match_label_metadata_encrypted}.
	MatchLabelMetadataEncrypted *string `field:"optional" json:"matchLabelMetadataEncrypted" yaml:"matchLabelMetadataEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#match_label_size CceNodeAttachV3#match_label_size}.
	MatchLabelSize *string `field:"optional" json:"matchLabelSize" yaml:"matchLabelSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#match_label_volume_type CceNodeAttachV3#match_label_volume_type}.
	MatchLabelVolumeType *string `field:"optional" json:"matchLabelVolumeType" yaml:"matchLabelVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#type CceNodeAttachV3#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

