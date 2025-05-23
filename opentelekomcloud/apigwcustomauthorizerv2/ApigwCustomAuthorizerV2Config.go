// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwcustomauthorizerv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwCustomAuthorizerV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#function_urn ApigwCustomAuthorizerV2#function_urn}.
	FunctionUrn *string `field:"required" json:"functionUrn" yaml:"functionUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#gateway_id ApigwCustomAuthorizerV2#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#name ApigwCustomAuthorizerV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#id ApigwCustomAuthorizerV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#identity ApigwCustomAuthorizerV2#identity}
	Identity interface{} `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#is_body_send ApigwCustomAuthorizerV2#is_body_send}.
	IsBodySend interface{} `field:"optional" json:"isBodySend" yaml:"isBodySend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#ttl ApigwCustomAuthorizerV2#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#type ApigwCustomAuthorizerV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_custom_authorizer_v2#user_data ApigwCustomAuthorizerV2#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

