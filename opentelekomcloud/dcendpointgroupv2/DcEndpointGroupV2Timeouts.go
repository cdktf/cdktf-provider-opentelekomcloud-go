// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcendpointgroupv2


type DcEndpointGroupV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/dc_endpoint_group_v2#create DcEndpointGroupV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/dc_endpoint_group_v2#delete DcEndpointGroupV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

