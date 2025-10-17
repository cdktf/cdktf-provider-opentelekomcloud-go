// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreassignpartitionsv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReassignPartitionsV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#instance_id DmsReassignPartitionsV2#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// reassignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#reassignments DmsReassignPartitionsV2#reassignments}
	Reassignments interface{} `field:"required" json:"reassignments" yaml:"reassignments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#execute_at DmsReassignPartitionsV2#execute_at}.
	ExecuteAt *float64 `field:"optional" json:"executeAt" yaml:"executeAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#id DmsReassignPartitionsV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#is_schedule DmsReassignPartitionsV2#is_schedule}.
	IsSchedule interface{} `field:"optional" json:"isSchedule" yaml:"isSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#throttle DmsReassignPartitionsV2#throttle}.
	Throttle *float64 `field:"optional" json:"throttle" yaml:"throttle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_reassign_partitions_v2#time_estimate DmsReassignPartitionsV2#time_estimate}.
	TimeEstimate interface{} `field:"optional" json:"timeEstimate" yaml:"timeEstimate"`
}

