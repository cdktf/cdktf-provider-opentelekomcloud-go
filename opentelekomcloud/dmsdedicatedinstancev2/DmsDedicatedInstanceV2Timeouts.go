// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsdedicatedinstancev2


type DmsDedicatedInstanceV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/dms_dedicated_instance_v2#create DmsDedicatedInstanceV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/dms_dedicated_instance_v2#delete DmsDedicatedInstanceV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/dms_dedicated_instance_v2#update DmsDedicatedInstanceV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

