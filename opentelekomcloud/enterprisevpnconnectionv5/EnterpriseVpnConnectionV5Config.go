// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseVpnConnectionV5Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#customer_gateway_id EnterpriseVpnConnectionV5#customer_gateway_id}.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#gateway_id EnterpriseVpnConnectionV5#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#gateway_ip EnterpriseVpnConnectionV5#gateway_ip}.
	GatewayIp *string `field:"required" json:"gatewayIp" yaml:"gatewayIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#name EnterpriseVpnConnectionV5#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#psk EnterpriseVpnConnectionV5#psk}.
	Psk *string `field:"required" json:"psk" yaml:"psk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#vpn_type EnterpriseVpnConnectionV5#vpn_type}.
	VpnType *string `field:"required" json:"vpnType" yaml:"vpnType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#enable_nqa EnterpriseVpnConnectionV5#enable_nqa}.
	EnableNqa interface{} `field:"optional" json:"enableNqa" yaml:"enableNqa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#ha_role EnterpriseVpnConnectionV5#ha_role}.
	HaRole *string `field:"optional" json:"haRole" yaml:"haRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#id EnterpriseVpnConnectionV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ikepolicy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#ikepolicy EnterpriseVpnConnectionV5#ikepolicy}
	Ikepolicy *EnterpriseVpnConnectionV5Ikepolicy `field:"optional" json:"ikepolicy" yaml:"ikepolicy"`
	// ipsecpolicy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#ipsecpolicy EnterpriseVpnConnectionV5#ipsecpolicy}
	Ipsecpolicy *EnterpriseVpnConnectionV5Ipsecpolicy `field:"optional" json:"ipsecpolicy" yaml:"ipsecpolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#peer_subnets EnterpriseVpnConnectionV5#peer_subnets}.
	PeerSubnets *[]*string `field:"optional" json:"peerSubnets" yaml:"peerSubnets"`
	// policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#policy_rules EnterpriseVpnConnectionV5#policy_rules}
	PolicyRules interface{} `field:"optional" json:"policyRules" yaml:"policyRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#tags EnterpriseVpnConnectionV5#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#timeouts EnterpriseVpnConnectionV5#timeouts}
	Timeouts *EnterpriseVpnConnectionV5Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#tunnel_local_address EnterpriseVpnConnectionV5#tunnel_local_address}.
	TunnelLocalAddress *string `field:"optional" json:"tunnelLocalAddress" yaml:"tunnelLocalAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_connection_v5#tunnel_peer_address EnterpriseVpnConnectionV5#tunnel_peer_address}.
	TunnelPeerAddress *string `field:"optional" json:"tunnelPeerAddress" yaml:"tunnelPeerAddress"`
}

