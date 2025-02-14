// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erroutetablev3


type ErRouteTableV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/er_route_table_v3#create ErRouteTableV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/er_route_table_v3#delete ErRouteTableV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/er_route_table_v3#update ErRouteTableV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

