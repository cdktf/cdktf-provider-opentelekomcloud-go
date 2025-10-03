// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingportv2


type NetworkingPortV2AllowedAddressPairs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/networking_port_v2#ip_address NetworkingPortV2#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/networking_port_v2#mac_address NetworkingPortV2#mac_address}.
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
}

