// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafalarmnotificationv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafAlarmNotificationV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#enabled WafAlarmNotificationV1#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#send_frequency WafAlarmNotificationV1#send_frequency}.
	SendFrequency *float64 `field:"required" json:"sendFrequency" yaml:"sendFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#threat WafAlarmNotificationV1#threat}.
	Threat *[]*string `field:"required" json:"threat" yaml:"threat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#times WafAlarmNotificationV1#times}.
	Times *float64 `field:"required" json:"times" yaml:"times"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#topic_urn WafAlarmNotificationV1#topic_urn}.
	TopicUrn *string `field:"required" json:"topicUrn" yaml:"topicUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#id WafAlarmNotificationV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_alarm_notification_v1#locale WafAlarmNotificationV1#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

