// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicateddomainv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedDomainV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#domain WafDedicatedDomainV1#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#server WafDedicatedDomainV1#server}
	Server interface{} `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#certificate_id WafDedicatedDomainV1#certificate_id}.
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#cipher WafDedicatedDomainV1#cipher}.
	Cipher *string `field:"optional" json:"cipher" yaml:"cipher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#id WafDedicatedDomainV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#keep_policy WafDedicatedDomainV1#keep_policy}.
	KeepPolicy interface{} `field:"optional" json:"keepPolicy" yaml:"keepPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#pci_3ds WafDedicatedDomainV1#pci_3ds}.
	Pci3Ds interface{} `field:"optional" json:"pci3Ds" yaml:"pci3Ds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#pci_dss WafDedicatedDomainV1#pci_dss}.
	PciDss interface{} `field:"optional" json:"pciDss" yaml:"pciDss"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#policy_id WafDedicatedDomainV1#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#protect_status WafDedicatedDomainV1#protect_status}.
	ProtectStatus *float64 `field:"optional" json:"protectStatus" yaml:"protectStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#proxy WafDedicatedDomainV1#proxy}.
	Proxy interface{} `field:"optional" json:"proxy" yaml:"proxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#region WafDedicatedDomainV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeout_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#timeout_config WafDedicatedDomainV1#timeout_config}
	TimeoutConfig *WafDedicatedDomainV1TimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/waf_dedicated_domain_v1#tls WafDedicatedDomainV1#tls}.
	Tls *string `field:"optional" json:"tls" yaml:"tls"`
}

