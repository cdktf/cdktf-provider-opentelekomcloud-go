// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directconnectv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectConnectV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#provider_name DirectConnectV2#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#admin_state_up DirectConnectV2#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#bandwidth DirectConnectV2#bandwidth}.
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#charge_mode DirectConnectV2#charge_mode}.
	ChargeMode *string `field:"optional" json:"chargeMode" yaml:"chargeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#description DirectConnectV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#device_id DirectConnectV2#device_id}.
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#hosting_id DirectConnectV2#hosting_id}.
	HostingId *string `field:"optional" json:"hostingId" yaml:"hostingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#interface_name DirectConnectV2#interface_name}.
	InterfaceName *string `field:"optional" json:"interfaceName" yaml:"interfaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#location DirectConnectV2#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#name DirectConnectV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#order_id DirectConnectV2#order_id}.
	OrderId *string `field:"optional" json:"orderId" yaml:"orderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#peer_location DirectConnectV2#peer_location}.
	PeerLocation *string `field:"optional" json:"peerLocation" yaml:"peerLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#port_type DirectConnectV2#port_type}.
	PortType *string `field:"optional" json:"portType" yaml:"portType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#product_id DirectConnectV2#product_id}.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#provider_status DirectConnectV2#provider_status}.
	ProviderStatus *string `field:"optional" json:"providerStatus" yaml:"providerStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#redundant_id DirectConnectV2#redundant_id}.
	RedundantId *string `field:"optional" json:"redundantId" yaml:"redundantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#status DirectConnectV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#tenant_id DirectConnectV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#timeouts DirectConnectV2#timeouts}
	Timeouts *DirectConnectV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#type DirectConnectV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/direct_connect_v2#vlan DirectConnectV2#vlan}.
	Vlan *float64 `field:"optional" json:"vlan" yaml:"vlan"`
}

