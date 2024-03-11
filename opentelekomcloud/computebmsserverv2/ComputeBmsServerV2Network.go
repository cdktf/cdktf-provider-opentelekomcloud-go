// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebmsserverv2


type ComputeBmsServerV2Network struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#access_network ComputeBmsServerV2#access_network}.
	AccessNetwork interface{} `field:"optional" json:"accessNetwork" yaml:"accessNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#fixed_ip_v4 ComputeBmsServerV2#fixed_ip_v4}.
	FixedIpV4 *string `field:"optional" json:"fixedIpV4" yaml:"fixedIpV4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#fixed_ip_v6 ComputeBmsServerV2#fixed_ip_v6}.
	FixedIpV6 *string `field:"optional" json:"fixedIpV6" yaml:"fixedIpV6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#name ComputeBmsServerV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#port ComputeBmsServerV2#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.4/docs/resources/compute_bms_server_v2#uuid ComputeBmsServerV2#uuid}.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

