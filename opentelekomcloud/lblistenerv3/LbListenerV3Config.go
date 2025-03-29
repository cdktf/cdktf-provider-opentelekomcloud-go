// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#loadbalancer_id LbListenerV3#loadbalancer_id}.
	LoadbalancerId *string `field:"required" json:"loadbalancerId" yaml:"loadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#protocol LbListenerV3#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#protocol_port LbListenerV3#protocol_port}.
	ProtocolPort *float64 `field:"required" json:"protocolPort" yaml:"protocolPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#admin_state_up LbListenerV3#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#advanced_forwarding LbListenerV3#advanced_forwarding}.
	AdvancedForwarding interface{} `field:"optional" json:"advancedForwarding" yaml:"advancedForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#client_ca_tls_container_ref LbListenerV3#client_ca_tls_container_ref}.
	ClientCaTlsContainerRef *string `field:"optional" json:"clientCaTlsContainerRef" yaml:"clientCaTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#client_timeout LbListenerV3#client_timeout}.
	ClientTimeout *float64 `field:"optional" json:"clientTimeout" yaml:"clientTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#default_pool_id LbListenerV3#default_pool_id}.
	DefaultPoolId *string `field:"optional" json:"defaultPoolId" yaml:"defaultPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#default_tls_container_ref LbListenerV3#default_tls_container_ref}.
	DefaultTlsContainerRef *string `field:"optional" json:"defaultTlsContainerRef" yaml:"defaultTlsContainerRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#description LbListenerV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#http2_enable LbListenerV3#http2_enable}.
	Http2Enable interface{} `field:"optional" json:"http2Enable" yaml:"http2Enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#id LbListenerV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// insert_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#insert_headers LbListenerV3#insert_headers}
	InsertHeaders *LbListenerV3InsertHeaders `field:"optional" json:"insertHeaders" yaml:"insertHeaders"`
	// ip_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#ip_group LbListenerV3#ip_group}
	IpGroup *LbListenerV3IpGroup `field:"optional" json:"ipGroup" yaml:"ipGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#keep_alive_timeout LbListenerV3#keep_alive_timeout}.
	KeepAliveTimeout *float64 `field:"optional" json:"keepAliveTimeout" yaml:"keepAliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#member_retry_enable LbListenerV3#member_retry_enable}.
	MemberRetryEnable interface{} `field:"optional" json:"memberRetryEnable" yaml:"memberRetryEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#member_timeout LbListenerV3#member_timeout}.
	MemberTimeout *float64 `field:"optional" json:"memberTimeout" yaml:"memberTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#name LbListenerV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#security_policy_id LbListenerV3#security_policy_id}.
	SecurityPolicyId *string `field:"optional" json:"securityPolicyId" yaml:"securityPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#sni_container_refs LbListenerV3#sni_container_refs}.
	SniContainerRefs *[]*string `field:"optional" json:"sniContainerRefs" yaml:"sniContainerRefs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#sni_match_algo LbListenerV3#sni_match_algo}.
	SniMatchAlgo *string `field:"optional" json:"sniMatchAlgo" yaml:"sniMatchAlgo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#tags LbListenerV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/lb_listener_v3#tls_ciphers_policy LbListenerV3#tls_ciphers_policy}.
	TlsCiphersPolicy *string `field:"optional" json:"tlsCiphersPolicy" yaml:"tlsCiphersPolicy"`
}

