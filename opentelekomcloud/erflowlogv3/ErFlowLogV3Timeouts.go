// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erflowlogv3


type ErFlowLogV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/er_flow_log_v3#create ErFlowLogV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/er_flow_log_v3#delete ErFlowLogV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/er_flow_log_v3#update ErFlowLogV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

