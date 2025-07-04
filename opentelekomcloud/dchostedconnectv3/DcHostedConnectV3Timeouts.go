// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dchostedconnectv3


type DcHostedConnectV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dc_hosted_connect_v3#create DcHostedConnectV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dc_hosted_connect_v3#delete DcHostedConnectV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dc_hosted_connect_v3#update DcHostedConnectV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

