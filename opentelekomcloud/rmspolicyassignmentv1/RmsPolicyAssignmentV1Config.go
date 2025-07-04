// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmspolicyassignmentv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RmsPolicyAssignmentV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#name RmsPolicyAssignmentV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#custom_policy RmsPolicyAssignmentV1#custom_policy}
	CustomPolicy *RmsPolicyAssignmentV1CustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#description RmsPolicyAssignmentV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#id RmsPolicyAssignmentV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#parameters RmsPolicyAssignmentV1#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#period RmsPolicyAssignmentV1#period}.
	Period *string `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#policy_definition_id RmsPolicyAssignmentV1#policy_definition_id}.
	PolicyDefinitionId *string `field:"optional" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// policy_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#policy_filter RmsPolicyAssignmentV1#policy_filter}
	PolicyFilter *RmsPolicyAssignmentV1PolicyFilter `field:"optional" json:"policyFilter" yaml:"policyFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#status RmsPolicyAssignmentV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/rms_policy_assignment_v1#timeouts RmsPolicyAssignmentV1#timeouts}
	Timeouts *RmsPolicyAssignmentV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

