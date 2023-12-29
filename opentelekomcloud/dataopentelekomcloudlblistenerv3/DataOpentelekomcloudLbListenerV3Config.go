// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudlblistenerv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudLbListenerV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#client_ca_tls_container_ref DataOpentelekomcloudLbListenerV3#client_ca_tls_container_ref}.
	ClientCaTlsContainerRef *string `field:"optional" json:"clientCaTlsContainerRef" yaml:"clientCaTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#client_timeout DataOpentelekomcloudLbListenerV3#client_timeout}.
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#default_pool_id DataOpentelekomcloudLbListenerV3#default_pool_id}.
	DefaultPoolId *string `field:"optional" json:"defaultPoolId" yaml:"defaultPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#default_tls_container_ref DataOpentelekomcloudLbListenerV3#default_tls_container_ref}.
	DefaultTlsContainerRef *string `field:"optional" json:"defaultTlsContainerRef" yaml:"defaultTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#description DataOpentelekomcloudLbListenerV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#id DataOpentelekomcloudLbListenerV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#keep_alive_timeout DataOpentelekomcloudLbListenerV3#keep_alive_timeout}.
	KeepAliveTimeout *float64 `field:"optional" json:"keepAliveTimeout" yaml:"keepAliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#loadbalancer_id DataOpentelekomcloudLbListenerV3#loadbalancer_id}.
	LoadbalancerId *string `field:"optional" json:"loadbalancerId" yaml:"loadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#member_address DataOpentelekomcloudLbListenerV3#member_address}.
	MemberAddress *string `field:"optional" json:"memberAddress" yaml:"memberAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#member_device_id DataOpentelekomcloudLbListenerV3#member_device_id}.
	MemberDeviceId *string `field:"optional" json:"memberDeviceId" yaml:"memberDeviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#member_timeout DataOpentelekomcloudLbListenerV3#member_timeout}.
	MemberTimeout *float64 `field:"optional" json:"memberTimeout" yaml:"memberTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#name DataOpentelekomcloudLbListenerV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#protocol DataOpentelekomcloudLbListenerV3#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#protocol_port DataOpentelekomcloudLbListenerV3#protocol_port}.
	ProtocolPort *float64 `field:"optional" json:"protocolPort" yaml:"protocolPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/data-sources/lb_listener_v3#tls_ciphers_policy DataOpentelekomcloudLbListenerV3#tls_ciphers_policy}.
	TlsCiphersPolicy *string `field:"optional" json:"tlsCiphersPolicy" yaml:"tlsCiphersPolicy"`
}

