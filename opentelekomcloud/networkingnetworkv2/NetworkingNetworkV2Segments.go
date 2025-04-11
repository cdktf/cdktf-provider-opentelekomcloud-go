// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingnetworkv2


type NetworkingNetworkV2Segments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/networking_network_v2#network_type NetworkingNetworkV2#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/networking_network_v2#physical_network NetworkingNetworkV2#physical_network}.
	PhysicalNetwork *string `field:"optional" json:"physicalNetwork" yaml:"physicalNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/networking_network_v2#segmentation_id NetworkingNetworkV2#segmentation_id}.
	SegmentationId *float64 `field:"optional" json:"segmentationId" yaml:"segmentationId"`
}

