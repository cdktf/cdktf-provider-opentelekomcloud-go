// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingrouterroutev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkingRouterRouteV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/networking_router_route_v2#destination_cidr NetworkingRouterRouteV2#destination_cidr}.
	DestinationCidr *string `field:"required" json:"destinationCidr" yaml:"destinationCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/networking_router_route_v2#next_hop NetworkingRouterRouteV2#next_hop}.
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/networking_router_route_v2#router_id NetworkingRouterRouteV2#router_id}.
	RouterId *string `field:"required" json:"routerId" yaml:"routerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/networking_router_route_v2#id NetworkingRouterRouteV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/networking_router_route_v2#region NetworkingRouterRouteV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

