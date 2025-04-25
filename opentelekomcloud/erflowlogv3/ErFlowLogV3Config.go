// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erflowlogv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ErFlowLogV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#instance_id ErFlowLogV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#log_group_id ErFlowLogV3#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#log_store_type ErFlowLogV3#log_store_type}.
	LogStoreType *string `field:"required" json:"logStoreType" yaml:"logStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#log_stream_id ErFlowLogV3#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#name ErFlowLogV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#resource_id ErFlowLogV3#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#resource_type ErFlowLogV3#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#description ErFlowLogV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#enabled ErFlowLogV3#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#id ErFlowLogV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/er_flow_log_v3#timeouts ErFlowLogV3#timeouts}
	Timeouts *ErFlowLogV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

