// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5


type EnterpriseVpnConnectionV5Ikepolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#authentication_algorithm EnterpriseVpnConnectionV5#authentication_algorithm}.
	AuthenticationAlgorithm *string `field:"optional" json:"authenticationAlgorithm" yaml:"authenticationAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#authentication_method EnterpriseVpnConnectionV5#authentication_method}.
	AuthenticationMethod *string `field:"optional" json:"authenticationMethod" yaml:"authenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#dh_group EnterpriseVpnConnectionV5#dh_group}.
	DhGroup *string `field:"optional" json:"dhGroup" yaml:"dhGroup"`
	// dpd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#dpd EnterpriseVpnConnectionV5#dpd}
	Dpd *EnterpriseVpnConnectionV5IkepolicyDpd `field:"optional" json:"dpd" yaml:"dpd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#encryption_algorithm EnterpriseVpnConnectionV5#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#ike_version EnterpriseVpnConnectionV5#ike_version}.
	IkeVersion *string `field:"optional" json:"ikeVersion" yaml:"ikeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#lifetime_seconds EnterpriseVpnConnectionV5#lifetime_seconds}.
	LifetimeSeconds *float64 `field:"optional" json:"lifetimeSeconds" yaml:"lifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#local_id EnterpriseVpnConnectionV5#local_id}.
	LocalId *string `field:"optional" json:"localId" yaml:"localId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#local_id_type EnterpriseVpnConnectionV5#local_id_type}.
	LocalIdType *string `field:"optional" json:"localIdType" yaml:"localIdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#peer_id EnterpriseVpnConnectionV5#peer_id}.
	PeerId *string `field:"optional" json:"peerId" yaml:"peerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#peer_id_type EnterpriseVpnConnectionV5#peer_id_type}.
	PeerIdType *string `field:"optional" json:"peerIdType" yaml:"peerIdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/enterprise_vpn_connection_v5#phase_one_negotiation_mode EnterpriseVpnConnectionV5#phase_one_negotiation_mode}.
	PhaseOneNegotiationMode *string `field:"optional" json:"phaseOneNegotiationMode" yaml:"phaseOneNegotiationMode"`
}

