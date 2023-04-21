package lbipgroupv3


type LbIpgroupV3IpList struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/lb_ipgroup_v3#ip LbIpgroupV3#ip}.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/lb_ipgroup_v3#description LbIpgroupV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

