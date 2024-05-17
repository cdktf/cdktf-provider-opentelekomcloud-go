// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdomainv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDomainV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#hostname WafDomainV1#hostname}.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#proxy WafDomainV1#proxy}.
	Proxy interface{} `field:"required" json:"proxy" yaml:"proxy"`
	// server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#server WafDomainV1#server}
	Server interface{} `field:"required" json:"server" yaml:"server"`
	// block_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#block_page WafDomainV1#block_page}
	BlockPage *WafDomainV1BlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#certificate_id WafDomainV1#certificate_id}.
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#cipher WafDomainV1#cipher}.
	Cipher *string `field:"optional" json:"cipher" yaml:"cipher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#id WafDomainV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#policy_id WafDomainV1#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#sip_header_list WafDomainV1#sip_header_list}.
	SipHeaderList *[]*string `field:"optional" json:"sipHeaderList" yaml:"sipHeaderList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#sip_header_name WafDomainV1#sip_header_name}.
	SipHeaderName *string `field:"optional" json:"sipHeaderName" yaml:"sipHeaderName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#timeouts WafDomainV1#timeouts}
	Timeouts *WafDomainV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/waf_domain_v1#tls WafDomainV1#tls}.
	Tls *string `field:"optional" json:"tls" yaml:"tls"`
}

