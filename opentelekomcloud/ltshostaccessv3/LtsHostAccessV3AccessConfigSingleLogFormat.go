// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltshostaccessv3


type LtsHostAccessV3AccessConfigSingleLogFormat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/lts_host_access_v3#mode LtsHostAccessV3#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/lts_host_access_v3#value LtsHostAccessV3#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

