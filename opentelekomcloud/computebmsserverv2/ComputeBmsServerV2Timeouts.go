// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebmsserverv2


type ComputeBmsServerV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_bms_server_v2#create ComputeBmsServerV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_bms_server_v2#delete ComputeBmsServerV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_bms_server_v2#update ComputeBmsServerV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

