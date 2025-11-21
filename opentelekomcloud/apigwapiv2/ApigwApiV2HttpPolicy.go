// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2


type ApigwApiV2HttpPolicy struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#conditions ApigwApiV2#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#name ApigwApiV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#request_method ApigwApiV2#request_method}.
	RequestMethod *string `field:"required" json:"requestMethod" yaml:"requestMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#request_uri ApigwApiV2#request_uri}.
	RequestUri *string `field:"required" json:"requestUri" yaml:"requestUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#authorizer_id ApigwApiV2#authorizer_id}.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// backend_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#backend_params ApigwApiV2#backend_params}
	BackendParams interface{} `field:"optional" json:"backendParams" yaml:"backendParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#effective_mode ApigwApiV2#effective_mode}.
	EffectiveMode *string `field:"optional" json:"effectiveMode" yaml:"effectiveMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#request_protocol ApigwApiV2#request_protocol}.
	RequestProtocol *string `field:"optional" json:"requestProtocol" yaml:"requestProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#retry_count ApigwApiV2#retry_count}.
	RetryCount *float64 `field:"optional" json:"retryCount" yaml:"retryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#timeout ApigwApiV2#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#url_domain ApigwApiV2#url_domain}.
	UrlDomain *string `field:"optional" json:"urlDomain" yaml:"urlDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#vpc_channel_id ApigwApiV2#vpc_channel_id}.
	VpcChannelId *string `field:"optional" json:"vpcChannelId" yaml:"vpcChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_api_v2#vpc_channel_proxy_host ApigwApiV2#vpc_channel_proxy_host}.
	VpcChannelProxyHost *string `field:"optional" json:"vpcChannelProxyHost" yaml:"vpcChannelProxyHost"`
}

