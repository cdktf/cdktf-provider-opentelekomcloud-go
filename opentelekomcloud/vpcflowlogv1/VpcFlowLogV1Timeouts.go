// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcflowlogv1


type VpcFlowLogV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vpc_flow_log_v1#create VpcFlowLogV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vpc_flow_log_v1#delete VpcFlowLogV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

