// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type AsConfigurationV1InstanceConfigPublicIpEip struct {
	// bandwidth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#bandwidth AsConfigurationV1#bandwidth}
	Bandwidth *AsConfigurationV1InstanceConfigPublicIpEipBandwidth `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#ip_type AsConfigurationV1#ip_type}.
	IpType *string `field:"required" json:"ipType" yaml:"ipType"`
}
