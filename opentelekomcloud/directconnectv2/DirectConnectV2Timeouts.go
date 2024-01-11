// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directconnectv2


type DirectConnectV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/direct_connect_v2#create DirectConnectV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/direct_connect_v2#delete DirectConnectV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/direct_connect_v2#update DirectConnectV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

