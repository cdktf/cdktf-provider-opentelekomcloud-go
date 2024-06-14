// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicateddomainv1


type WafDedicatedDomainV1TimeoutConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_dedicated_domain_v1#connect_timeout WafDedicatedDomainV1#connect_timeout}.
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_dedicated_domain_v1#read_timeout WafDedicatedDomainV1#read_timeout}.
	ReadTimeout *float64 `field:"optional" json:"readTimeout" yaml:"readTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/waf_dedicated_domain_v1#send_timeout WafDedicatedDomainV1#send_timeout}.
	SendTimeout *float64 `field:"optional" json:"sendTimeout" yaml:"sendTimeout"`
}

