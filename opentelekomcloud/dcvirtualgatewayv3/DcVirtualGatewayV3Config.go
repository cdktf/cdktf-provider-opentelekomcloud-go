// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualgatewayv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcVirtualGatewayV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#local_ep_group DcVirtualGatewayV3#local_ep_group}.
	LocalEpGroup *[]*string `field:"required" json:"localEpGroup" yaml:"localEpGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#name DcVirtualGatewayV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#vpc_id DcVirtualGatewayV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#asn DcVirtualGatewayV3#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#description DcVirtualGatewayV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#id DcVirtualGatewayV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dc_virtual_gateway_v3#local_ep_group_ipv6 DcVirtualGatewayV3#local_ep_group_ipv6}.
	LocalEpGroupIpv6 *[]*string `field:"optional" json:"localEpGroupIpv6" yaml:"localEpGroupIpv6"`
}

