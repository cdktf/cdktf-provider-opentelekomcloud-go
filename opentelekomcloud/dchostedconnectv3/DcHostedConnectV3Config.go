// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dchostedconnectv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcHostedConnectV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#bandwidth DcHostedConnectV3#bandwidth}.
	Bandwidth *float64 `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#hosting_id DcHostedConnectV3#hosting_id}.
	HostingId *string `field:"required" json:"hostingId" yaml:"hostingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#resource_tenant_id DcHostedConnectV3#resource_tenant_id}.
	ResourceTenantId *string `field:"required" json:"resourceTenantId" yaml:"resourceTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#vlan DcHostedConnectV3#vlan}.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#description DcHostedConnectV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#id DcHostedConnectV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#name DcHostedConnectV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#peer_location DcHostedConnectV3#peer_location}.
	PeerLocation *string `field:"optional" json:"peerLocation" yaml:"peerLocation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_hosted_connect_v3#timeouts DcHostedConnectV3#timeouts}
	Timeouts *DcHostedConnectV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

