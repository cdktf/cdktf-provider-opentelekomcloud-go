// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsinstancev1


type EcsInstanceV1Nics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/ecs_instance_v1#network_id EcsInstanceV1#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/ecs_instance_v1#ip_address EcsInstanceV1#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/ecs_instance_v1#ipv6_enable EcsInstanceV1#ipv6_enable}.
	Ipv6Enable interface{} `field:"optional" json:"ipv6Enable" yaml:"ipv6Enable"`
}

