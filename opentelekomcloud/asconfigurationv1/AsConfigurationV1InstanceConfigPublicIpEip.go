package asconfigurationv1


type AsConfigurationV1InstanceConfigPublicIpEip struct {
	// bandwidth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/as_configuration_v1#bandwidth AsConfigurationV1#bandwidth}
	Bandwidth *AsConfigurationV1InstanceConfigPublicIpEipBandwidth `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/as_configuration_v1#ip_type AsConfigurationV1#ip_type}.
	IpType *string `field:"required" json:"ipType" yaml:"ipType"`
}

