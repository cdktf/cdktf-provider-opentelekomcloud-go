package asconfigurationv1


type AsConfigurationV1InstanceConfigPublicIp struct {
	// eip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#eip AsConfigurationV1#eip}
	Eip *AsConfigurationV1InstanceConfigPublicIpEip `field:"required" json:"eip" yaml:"eip"`
}

