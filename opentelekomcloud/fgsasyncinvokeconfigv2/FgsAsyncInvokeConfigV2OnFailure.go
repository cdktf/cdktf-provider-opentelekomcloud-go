// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsasyncinvokeconfigv2


type FgsAsyncInvokeConfigV2OnFailure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/fgs_async_invoke_config_v2#destination FgsAsyncInvokeConfigV2#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/fgs_async_invoke_config_v2#param FgsAsyncInvokeConfigV2#param}.
	Param *string `field:"required" json:"param" yaml:"param"`
}

