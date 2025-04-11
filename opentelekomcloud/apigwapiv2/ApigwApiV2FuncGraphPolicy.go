// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2


type ApigwApiV2FuncGraphPolicy struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#conditions ApigwApiV2#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#function_urn ApigwApiV2#function_urn}.
	FunctionUrn *string `field:"required" json:"functionUrn" yaml:"functionUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#name ApigwApiV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#authorizer_id ApigwApiV2#authorizer_id}.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// backend_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#backend_params ApigwApiV2#backend_params}
	BackendParams interface{} `field:"optional" json:"backendParams" yaml:"backendParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#effective_mode ApigwApiV2#effective_mode}.
	EffectiveMode *string `field:"optional" json:"effectiveMode" yaml:"effectiveMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#invocation_type ApigwApiV2#invocation_type}.
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#network_type ApigwApiV2#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#timeout ApigwApiV2#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_api_v2#version ApigwApiV2#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

