// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdrsprotectedinstancev1


type SdrsProtectedInstanceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.0/docs/resources/sdrs_protected_instance_v1#create SdrsProtectedInstanceV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.0/docs/resources/sdrs_protected_instance_v1#delete SdrsProtectedInstanceV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

