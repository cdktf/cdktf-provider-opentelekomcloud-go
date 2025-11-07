// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpncustomergatewayv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnCustomerGatewayV5Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#name EnterpriseVpnCustomerGatewayV5#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#asn EnterpriseVpnCustomerGatewayV5#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#id EnterpriseVpnCustomerGatewayV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#id_type EnterpriseVpnCustomerGatewayV5#id_type}.
	IdType *string `field:"optional" json:"idType" yaml:"idType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#id_value EnterpriseVpnCustomerGatewayV5#id_value}.
	IdValue *string `field:"optional" json:"idValue" yaml:"idValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/enterprise_vpn_customer_gateway_v5#tags EnterpriseVpnCustomerGatewayV5#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

