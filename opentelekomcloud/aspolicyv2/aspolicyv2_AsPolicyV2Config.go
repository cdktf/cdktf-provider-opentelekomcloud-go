package aspolicyv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsPolicyV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scaling_policy_name AsPolicyV2#scaling_policy_name}.
	ScalingPolicyName *string `field:"required" json:"scalingPolicyName" yaml:"scalingPolicyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scaling_policy_type AsPolicyV2#scaling_policy_type}.
	ScalingPolicyType *string `field:"required" json:"scalingPolicyType" yaml:"scalingPolicyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scaling_resource_id AsPolicyV2#scaling_resource_id}.
	ScalingResourceId *string `field:"required" json:"scalingResourceId" yaml:"scalingResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scaling_resource_type AsPolicyV2#scaling_resource_type}.
	ScalingResourceType *string `field:"required" json:"scalingResourceType" yaml:"scalingResourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#alarm_id AsPolicyV2#alarm_id}.
	AlarmId *string `field:"optional" json:"alarmId" yaml:"alarmId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#cool_down_time AsPolicyV2#cool_down_time}.
	CoolDownTime *float64 `field:"optional" json:"coolDownTime" yaml:"coolDownTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#id AsPolicyV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#region AsPolicyV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scaling_policy_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scaling_policy_action AsPolicyV2#scaling_policy_action}
	ScalingPolicyAction interface{} `field:"optional" json:"scalingPolicyAction" yaml:"scalingPolicyAction"`
	// scheduled_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#scheduled_policy AsPolicyV2#scheduled_policy}
	ScheduledPolicy interface{} `field:"optional" json:"scheduledPolicy" yaml:"scheduledPolicy"`
}

