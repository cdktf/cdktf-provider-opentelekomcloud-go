// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbloadbalancerv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbLoadbalancerV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#availability_zones LbLoadbalancerV3#availability_zones}.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#network_ids LbLoadbalancerV3#network_ids}.
	NetworkIds *[]*string `field:"required" json:"networkIds" yaml:"networkIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#admin_state_up LbLoadbalancerV3#admin_state_up}.
	AdminStateUp interface{} `field:"optional" json:"adminStateUp" yaml:"adminStateUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#deletion_protection LbLoadbalancerV3#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#description LbLoadbalancerV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#id LbLoadbalancerV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#ip_target_enable LbLoadbalancerV3#ip_target_enable}.
	IpTargetEnable interface{} `field:"optional" json:"ipTargetEnable" yaml:"ipTargetEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#l4_flavor LbLoadbalancerV3#l4_flavor}.
	L4Flavor *string `field:"optional" json:"l4Flavor" yaml:"l4Flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#l7_flavor LbLoadbalancerV3#l7_flavor}.
	L7Flavor *string `field:"optional" json:"l7Flavor" yaml:"l7Flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#name LbLoadbalancerV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// public_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#public_ip LbLoadbalancerV3#public_ip}
	PublicIp *LbLoadbalancerV3PublicIp `field:"optional" json:"publicIp" yaml:"publicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#router_id LbLoadbalancerV3#router_id}.
	RouterId *string `field:"optional" json:"routerId" yaml:"routerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#subnet_id LbLoadbalancerV3#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#tags LbLoadbalancerV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lb_loadbalancer_v3#vip_address LbLoadbalancerV3#vip_address}.
	VipAddress *string `field:"optional" json:"vipAddress" yaml:"vipAddress"`
}

