// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aspolicyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsPolicyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#scaling_group_id AsPolicyV1#scaling_group_id}.
	ScalingGroupId *string `field:"required" json:"scalingGroupId" yaml:"scalingGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#scaling_policy_name AsPolicyV1#scaling_policy_name}.
	ScalingPolicyName *string `field:"required" json:"scalingPolicyName" yaml:"scalingPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#scaling_policy_type AsPolicyV1#scaling_policy_type}.
	ScalingPolicyType *string `field:"required" json:"scalingPolicyType" yaml:"scalingPolicyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#alarm_id AsPolicyV1#alarm_id}.
	AlarmId *string `field:"optional" json:"alarmId" yaml:"alarmId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#cool_down_time AsPolicyV1#cool_down_time}.
	CoolDownTime *float64 `field:"optional" json:"coolDownTime" yaml:"coolDownTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#id AsPolicyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#region AsPolicyV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scaling_policy_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#scaling_policy_action AsPolicyV1#scaling_policy_action}
	ScalingPolicyAction *AsPolicyV1ScalingPolicyAction `field:"optional" json:"scalingPolicyAction" yaml:"scalingPolicyAction"`
	// scheduled_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/as_policy_v1#scheduled_policy AsPolicyV1#scheduled_policy}
	ScheduledPolicy *AsPolicyV1ScheduledPolicy `field:"optional" json:"scheduledPolicy" yaml:"scheduledPolicy"`
}

