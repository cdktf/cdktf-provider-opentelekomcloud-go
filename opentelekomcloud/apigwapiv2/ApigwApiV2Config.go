// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwApiV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#gateway_id ApigwApiV2#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#group_id ApigwApiV2#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#name ApigwApiV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#request_method ApigwApiV2#request_method}.
	RequestMethod *string `field:"required" json:"requestMethod" yaml:"requestMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#request_protocol ApigwApiV2#request_protocol}.
	RequestProtocol *string `field:"required" json:"requestProtocol" yaml:"requestProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#request_uri ApigwApiV2#request_uri}.
	RequestUri *string `field:"required" json:"requestUri" yaml:"requestUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#type ApigwApiV2#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#authorizer_id ApigwApiV2#authorizer_id}.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// backend_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#backend_params ApigwApiV2#backend_params}
	BackendParams interface{} `field:"optional" json:"backendParams" yaml:"backendParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#body_description ApigwApiV2#body_description}.
	BodyDescription *string `field:"optional" json:"bodyDescription" yaml:"bodyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#cors ApigwApiV2#cors}.
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#description ApigwApiV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#failure_response ApigwApiV2#failure_response}.
	FailureResponse *string `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// func_graph block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#func_graph ApigwApiV2#func_graph}
	FuncGraph *ApigwApiV2FuncGraph `field:"optional" json:"funcGraph" yaml:"funcGraph"`
	// func_graph_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#func_graph_policy ApigwApiV2#func_graph_policy}
	FuncGraphPolicy interface{} `field:"optional" json:"funcGraphPolicy" yaml:"funcGraphPolicy"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#http ApigwApiV2#http}
	Http *ApigwApiV2Http `field:"optional" json:"http" yaml:"http"`
	// http_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#http_policy ApigwApiV2#http_policy}
	HttpPolicy interface{} `field:"optional" json:"httpPolicy" yaml:"httpPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#id ApigwApiV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#match_mode ApigwApiV2#match_mode}.
	MatchMode *string `field:"optional" json:"matchMode" yaml:"matchMode"`
	// mock block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#mock ApigwApiV2#mock}
	Mock *ApigwApiV2Mock `field:"optional" json:"mock" yaml:"mock"`
	// mock_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#mock_policy ApigwApiV2#mock_policy}
	MockPolicy interface{} `field:"optional" json:"mockPolicy" yaml:"mockPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#region ApigwApiV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// request_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#request_params ApigwApiV2#request_params}
	RequestParams interface{} `field:"optional" json:"requestParams" yaml:"requestParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#response_id ApigwApiV2#response_id}.
	ResponseId *string `field:"optional" json:"responseId" yaml:"responseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#security_authentication_enabled ApigwApiV2#security_authentication_enabled}.
	SecurityAuthenticationEnabled interface{} `field:"optional" json:"securityAuthenticationEnabled" yaml:"securityAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#security_authentication_type ApigwApiV2#security_authentication_type}.
	SecurityAuthenticationType *string `field:"optional" json:"securityAuthenticationType" yaml:"securityAuthenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#success_response ApigwApiV2#success_response}.
	SuccessResponse *string `field:"optional" json:"successResponse" yaml:"successResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#tags ApigwApiV2#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/apigw_api_v2#version ApigwApiV2#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

