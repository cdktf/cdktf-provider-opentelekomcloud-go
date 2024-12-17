// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesalarmrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CesAlarmruleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_name CesAlarmrule#alarm_name}.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#condition CesAlarmrule#condition}
	Condition *CesAlarmruleCondition `field:"required" json:"condition" yaml:"condition"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#metric CesAlarmrule#metric}
	Metric *CesAlarmruleMetric `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_action_enabled CesAlarmrule#alarm_action_enabled}.
	AlarmActionEnabled interface{} `field:"optional" json:"alarmActionEnabled" yaml:"alarmActionEnabled"`
	// alarm_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_actions CesAlarmrule#alarm_actions}
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_description CesAlarmrule#alarm_description}.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_enabled CesAlarmrule#alarm_enabled}.
	AlarmEnabled interface{} `field:"optional" json:"alarmEnabled" yaml:"alarmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_level CesAlarmrule#alarm_level}.
	AlarmLevel *float64 `field:"optional" json:"alarmLevel" yaml:"alarmLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#alarm_type CesAlarmrule#alarm_type}.
	AlarmType *string `field:"optional" json:"alarmType" yaml:"alarmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#id CesAlarmrule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ok_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#ok_actions CesAlarmrule#ok_actions}
	OkActions interface{} `field:"optional" json:"okActions" yaml:"okActions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/ces_alarmrule#timeouts CesAlarmrule#timeouts}
	Timeouts *CesAlarmruleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

