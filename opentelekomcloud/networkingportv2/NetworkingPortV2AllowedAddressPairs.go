package networkingportv2


type NetworkingPortV2AllowedAddressPairs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_port_v2#ip_address NetworkingPortV2#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/networking_port_v2#mac_address NetworkingPortV2#mac_address}.
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
}

