package networkingnetworkv2


type NetworkingNetworkV2Segments struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_network_v2#network_type NetworkingNetworkV2#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_network_v2#physical_network NetworkingNetworkV2#physical_network}.
	PhysicalNetwork *string `field:"optional" json:"physicalNetwork" yaml:"physicalNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_network_v2#segmentation_id NetworkingNetworkV2#segmentation_id}.
	SegmentationId *float64 `field:"optional" json:"segmentationId" yaml:"segmentationId"`
}

