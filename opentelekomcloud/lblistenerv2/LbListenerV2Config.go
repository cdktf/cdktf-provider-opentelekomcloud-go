// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#loadbalancer_id LbListenerV2#loadbalancer_id}.
	LoadbalancerId *string `field:"required" json:"loadbalancerId" yaml:"loadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#protocol LbListenerV2#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#protocol_port LbListenerV2#protocol_port}.
	ProtocolPort *float64 `field:"required" json:"protocolPort" yaml:"protocolPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#admin_state_up LbListenerV2#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#client_ca_tls_container_ref LbListenerV2#client_ca_tls_container_ref}.
	ClientCaTlsContainerRef *string `field:"optional" json:"clientCaTlsContainerRef" yaml:"clientCaTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#default_pool_id LbListenerV2#default_pool_id}.
	DefaultPoolId *string `field:"optional" json:"defaultPoolId" yaml:"defaultPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#default_tls_container_ref LbListenerV2#default_tls_container_ref}.
	DefaultTlsContainerRef *string `field:"optional" json:"defaultTlsContainerRef" yaml:"defaultTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#description LbListenerV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#http2_enable LbListenerV2#http2_enable}.
	Http2Enable interface{} `field:"optional" json:"http2Enable" yaml:"http2Enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#id LbListenerV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#ip_group LbListenerV2#ip_group}
	IpGroup *LbListenerV2IpGroup `field:"optional" json:"ipGroup" yaml:"ipGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#name LbListenerV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#region LbListenerV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#sni_container_refs LbListenerV2#sni_container_refs}.
	SniContainerRefs *[]*string `field:"optional" json:"sniContainerRefs" yaml:"sniContainerRefs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#tags LbListenerV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#tenant_id LbListenerV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#timeouts LbListenerV2#timeouts}
	Timeouts *LbListenerV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#tls_ciphers_policy LbListenerV2#tls_ciphers_policy}.
	TlsCiphersPolicy *string `field:"optional" json:"tlsCiphersPolicy" yaml:"tlsCiphersPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/lb_listener_v2#transparent_client_ip_enable LbListenerV2#transparent_client_ip_enable}.
	TransparentClientIpEnable interface{} `field:"optional" json:"transparentClientIpEnable" yaml:"transparentClientIpEnable"`
}

