// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmssmartconnecttaskv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsSmartConnectTaskV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#instance_id DmsSmartConnectTaskV2#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#task_name DmsSmartConnectTaskV2#task_name}.
	TaskName *string `field:"required" json:"taskName" yaml:"taskName"`
	// destination_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#destination_task DmsSmartConnectTaskV2#destination_task}
	DestinationTask *DmsSmartConnectTaskV2DestinationTask `field:"optional" json:"destinationTask" yaml:"destinationTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#destination_type DmsSmartConnectTaskV2#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#id DmsSmartConnectTaskV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// source_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#source_task DmsSmartConnectTaskV2#source_task}
	SourceTask *DmsSmartConnectTaskV2SourceTask `field:"optional" json:"sourceTask" yaml:"sourceTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#source_type DmsSmartConnectTaskV2#source_type}.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#start_later DmsSmartConnectTaskV2#start_later}.
	StartLater interface{} `field:"optional" json:"startLater" yaml:"startLater"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#timeouts DmsSmartConnectTaskV2#timeouts}
	Timeouts *DmsSmartConnectTaskV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#topics DmsSmartConnectTaskV2#topics}.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dms_smart_connect_task_v2#topics_regex DmsSmartConnectTaskV2#topics_regex}.
	TopicsRegex *string `field:"optional" json:"topicsRegex" yaml:"topicsRegex"`
}

