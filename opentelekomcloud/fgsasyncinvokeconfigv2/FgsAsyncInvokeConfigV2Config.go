// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsasyncinvokeconfigv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsAsyncInvokeConfigV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#function_urn FgsAsyncInvokeConfigV2#function_urn}.
	FunctionUrn *string `field:"required" json:"functionUrn" yaml:"functionUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#max_async_event_age_in_seconds FgsAsyncInvokeConfigV2#max_async_event_age_in_seconds}.
	MaxAsyncEventAgeInSeconds *float64 `field:"required" json:"maxAsyncEventAgeInSeconds" yaml:"maxAsyncEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#max_async_retry_attempts FgsAsyncInvokeConfigV2#max_async_retry_attempts}.
	MaxAsyncRetryAttempts *float64 `field:"required" json:"maxAsyncRetryAttempts" yaml:"maxAsyncRetryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#id FgsAsyncInvokeConfigV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#on_failure FgsAsyncInvokeConfigV2#on_failure}
	OnFailure *FgsAsyncInvokeConfigV2OnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/fgs_async_invoke_config_v2#on_success FgsAsyncInvokeConfigV2#on_success}
	OnSuccess *FgsAsyncInvokeConfigV2OnSuccess `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

