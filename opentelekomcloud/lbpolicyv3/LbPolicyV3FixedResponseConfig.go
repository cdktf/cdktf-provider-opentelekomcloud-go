// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpolicyv3


type LbPolicyV3FixedResponseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_policy_v3#status_code LbPolicyV3#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_policy_v3#content_type LbPolicyV3#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_policy_v3#message_body LbPolicyV3#message_body}.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

