// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwthrottlingpolicyv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwThrottlingPolicyV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#instance_id ApigwThrottlingPolicyV2#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#max_api_requests ApigwThrottlingPolicyV2#max_api_requests}.
	MaxApiRequests *float64 `field:"required" json:"maxApiRequests" yaml:"maxApiRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#name ApigwThrottlingPolicyV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#period ApigwThrottlingPolicyV2#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// app_throttles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#app_throttles ApigwThrottlingPolicyV2#app_throttles}
	AppThrottles interface{} `field:"optional" json:"appThrottles" yaml:"appThrottles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#description ApigwThrottlingPolicyV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#id ApigwThrottlingPolicyV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#max_app_requests ApigwThrottlingPolicyV2#max_app_requests}.
	MaxAppRequests *float64 `field:"optional" json:"maxAppRequests" yaml:"maxAppRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#max_ip_requests ApigwThrottlingPolicyV2#max_ip_requests}.
	MaxIpRequests *float64 `field:"optional" json:"maxIpRequests" yaml:"maxIpRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#max_user_requests ApigwThrottlingPolicyV2#max_user_requests}.
	MaxUserRequests *float64 `field:"optional" json:"maxUserRequests" yaml:"maxUserRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#period_unit ApigwThrottlingPolicyV2#period_unit}.
	PeriodUnit *string `field:"optional" json:"periodUnit" yaml:"periodUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#type ApigwThrottlingPolicyV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// user_throttles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/apigw_throttling_policy_v2#user_throttles ApigwThrottlingPolicyV2#user_throttles}
	UserThrottles interface{} `field:"optional" json:"userThrottles" yaml:"userThrottles"`
}

