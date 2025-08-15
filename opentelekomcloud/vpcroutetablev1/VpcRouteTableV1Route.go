// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcroutetablev1


type VpcRouteTableV1Route struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_route_table_v1#destination VpcRouteTableV1#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_route_table_v1#nexthop VpcRouteTableV1#nexthop}.
	Nexthop *string `field:"required" json:"nexthop" yaml:"nexthop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_route_table_v1#type VpcRouteTableV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_route_table_v1#description VpcRouteTableV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

