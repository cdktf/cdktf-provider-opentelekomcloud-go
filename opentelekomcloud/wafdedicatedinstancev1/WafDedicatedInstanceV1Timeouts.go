// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedinstancev1


type WafDedicatedInstanceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/waf_dedicated_instance_v1#create WafDedicatedInstanceV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/waf_dedicated_instance_v1#delete WafDedicatedInstanceV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

