// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcroutetablev1


type VpcRouteTableV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vpc_route_table_v1#create VpcRouteTableV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/vpc_route_table_v1#delete VpcRouteTableV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

