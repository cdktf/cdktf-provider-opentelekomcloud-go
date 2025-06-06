// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpceipv1


type VpcEipV1Publicip struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/vpc_eip_v1#type VpcEipV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/vpc_eip_v1#ip_address VpcEipV1#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/vpc_eip_v1#name VpcEipV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/vpc_eip_v1#port_id VpcEipV1#port_id}.
	PortId *string `field:"optional" json:"portId" yaml:"portId"`
}

