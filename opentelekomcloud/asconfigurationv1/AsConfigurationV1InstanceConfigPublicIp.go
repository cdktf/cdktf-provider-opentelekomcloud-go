// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asconfigurationv1


type AsConfigurationV1InstanceConfigPublicIp struct {
	// eip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/as_configuration_v1#eip AsConfigurationV1#eip}
	Eip *AsConfigurationV1InstanceConfigPublicIpEip `field:"required" json:"eip" yaml:"eip"`
}

