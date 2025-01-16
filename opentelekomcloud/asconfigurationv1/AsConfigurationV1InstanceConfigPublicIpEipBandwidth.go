// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asconfigurationv1


type AsConfigurationV1InstanceConfigPublicIpEipBandwidth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.29/docs/resources/as_configuration_v1#charging_mode AsConfigurationV1#charging_mode}.
	ChargingMode *string `field:"required" json:"chargingMode" yaml:"chargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.29/docs/resources/as_configuration_v1#share_type AsConfigurationV1#share_type}.
	ShareType *string `field:"required" json:"shareType" yaml:"shareType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.29/docs/resources/as_configuration_v1#size AsConfigurationV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

