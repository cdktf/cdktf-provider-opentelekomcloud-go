// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscceaccessv3


type LtsCceAccessV3AccessConfigSingleLogFormat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/lts_cce_access_v3#mode LtsCceAccessV3#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/lts_cce_access_v3#value LtsCceAccessV3#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

