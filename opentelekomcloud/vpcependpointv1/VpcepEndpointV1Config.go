// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcependpointv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcepEndpointV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#service_id VpcepEndpointV1#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#vpc_id VpcepEndpointV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#description VpcepEndpointV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#enable_dns VpcepEndpointV1#enable_dns}.
	EnableDns interface{} `field:"optional" json:"enableDns" yaml:"enableDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#enable_whitelist VpcepEndpointV1#enable_whitelist}.
	EnableWhitelist interface{} `field:"optional" json:"enableWhitelist" yaml:"enableWhitelist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#id VpcepEndpointV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#policy_statement VpcepEndpointV1#policy_statement}.
	PolicyStatement *string `field:"optional" json:"policyStatement" yaml:"policyStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#port_ip VpcepEndpointV1#port_ip}.
	PortIp *string `field:"optional" json:"portIp" yaml:"portIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#route_tables VpcepEndpointV1#route_tables}.
	RouteTables *[]*string `field:"optional" json:"routeTables" yaml:"routeTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#subnet_id VpcepEndpointV1#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#tags VpcepEndpointV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#timeouts VpcepEndpointV1#timeouts}
	Timeouts *VpcepEndpointV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/vpcep_endpoint_v1#whitelist VpcepEndpointV1#whitelist}.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

