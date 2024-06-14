// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdomainv1


type WafDomainV1Server struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#address WafDomainV1#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#port WafDomainV1#port}.
	Port *string `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#back_protocol WafDomainV1#back_protocol}.
	BackProtocol *string `field:"optional" json:"backProtocol" yaml:"backProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#client_protocol WafDomainV1#client_protocol}.
	ClientProtocol *string `field:"optional" json:"clientProtocol" yaml:"clientProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#front_protocol WafDomainV1#front_protocol}.
	FrontProtocol *string `field:"optional" json:"frontProtocol" yaml:"frontProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_domain_v1#server_protocol WafDomainV1#server_protocol}.
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
}

