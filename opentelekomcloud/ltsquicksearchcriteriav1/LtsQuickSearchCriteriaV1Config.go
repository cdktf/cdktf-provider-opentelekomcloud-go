// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltsquicksearchcriteriav1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsQuickSearchCriteriaV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#criteria LtsQuickSearchCriteriaV1#criteria}.
	Criteria *string `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#log_group_id LtsQuickSearchCriteriaV1#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#log_stream_id LtsQuickSearchCriteriaV1#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#name LtsQuickSearchCriteriaV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#type LtsQuickSearchCriteriaV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_quick_search_criteria_v1#id LtsQuickSearchCriteriaV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

