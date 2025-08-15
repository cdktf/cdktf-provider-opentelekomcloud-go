// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcepapprovalv1


type VpcepApprovalV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpcep_approval_v1#create VpcepApprovalV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpcep_approval_v1#delete VpcepApprovalV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

