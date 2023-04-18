package networkingsecgrouprulev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkingSecgroupRuleV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#direction NetworkingSecgroupRuleV2#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#ethertype NetworkingSecgroupRuleV2#ethertype}.
	Ethertype *string `field:"required" json:"ethertype" yaml:"ethertype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#security_group_id NetworkingSecgroupRuleV2#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#description NetworkingSecgroupRuleV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#id NetworkingSecgroupRuleV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#port_range_max NetworkingSecgroupRuleV2#port_range_max}.
	PortRangeMax *float64 `field:"optional" json:"portRangeMax" yaml:"portRangeMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#port_range_min NetworkingSecgroupRuleV2#port_range_min}.
	PortRangeMin *float64 `field:"optional" json:"portRangeMin" yaml:"portRangeMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#protocol NetworkingSecgroupRuleV2#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#region NetworkingSecgroupRuleV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#remote_group_id NetworkingSecgroupRuleV2#remote_group_id}.
	RemoteGroupId *string `field:"optional" json:"remoteGroupId" yaml:"remoteGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#remote_ip_prefix NetworkingSecgroupRuleV2#remote_ip_prefix}.
	RemoteIpPrefix *string `field:"optional" json:"remoteIpPrefix" yaml:"remoteIpPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#tenant_id NetworkingSecgroupRuleV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/networking_secgroup_rule_v2#timeouts NetworkingSecgroupRuleV2#timeouts}
	Timeouts *NetworkingSecgroupRuleV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

