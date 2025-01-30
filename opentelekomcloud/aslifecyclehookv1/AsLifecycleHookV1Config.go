// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aslifecyclehookv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsLifecycleHookV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#notification_topic_urn AsLifecycleHookV1#notification_topic_urn}.
	NotificationTopicUrn *string `field:"required" json:"notificationTopicUrn" yaml:"notificationTopicUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#scaling_group_id AsLifecycleHookV1#scaling_group_id}.
	ScalingGroupId *string `field:"required" json:"scalingGroupId" yaml:"scalingGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#scaling_lifecycle_hook_name AsLifecycleHookV1#scaling_lifecycle_hook_name}.
	ScalingLifecycleHookName *string `field:"required" json:"scalingLifecycleHookName" yaml:"scalingLifecycleHookName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#scaling_lifecycle_hook_type AsLifecycleHookV1#scaling_lifecycle_hook_type}.
	ScalingLifecycleHookType *string `field:"required" json:"scalingLifecycleHookType" yaml:"scalingLifecycleHookType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#default_result AsLifecycleHookV1#default_result}.
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#default_timeout AsLifecycleHookV1#default_timeout}.
	DefaultTimeout *float64 `field:"optional" json:"defaultTimeout" yaml:"defaultTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#id AsLifecycleHookV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/as_lifecycle_hook_v1#notification_metadata AsLifecycleHookV1#notification_metadata}.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
}

