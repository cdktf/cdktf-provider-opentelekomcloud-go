// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erstaticroutev3


type ErStaticRouteV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/er_static_route_v3#create ErStaticRouteV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/er_static_route_v3#delete ErStaticRouteV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/er_static_route_v3#update ErStaticRouteV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

