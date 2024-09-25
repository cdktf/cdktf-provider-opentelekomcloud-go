// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerv3


type LbListenerV3InsertHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/lb_listener_v3#forwarded_for_port LbListenerV3#forwarded_for_port}.
	ForwardedForPort interface{} `field:"optional" json:"forwardedForPort" yaml:"forwardedForPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/lb_listener_v3#forwarded_host LbListenerV3#forwarded_host}.
	ForwardedHost interface{} `field:"optional" json:"forwardedHost" yaml:"forwardedHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/lb_listener_v3#forwarded_port LbListenerV3#forwarded_port}.
	ForwardedPort interface{} `field:"optional" json:"forwardedPort" yaml:"forwardedPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/lb_listener_v3#forward_elb_ip LbListenerV3#forward_elb_ip}.
	ForwardElbIp interface{} `field:"optional" json:"forwardElbIp" yaml:"forwardElbIp"`
}

