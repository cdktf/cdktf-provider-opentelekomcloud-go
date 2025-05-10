// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscrossaccountaccessv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsCrossAccountAccessV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#agency_domain_name LtsCrossAccountAccessV2#agency_domain_name}.
	AgencyDomainName *string `field:"required" json:"agencyDomainName" yaml:"agencyDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#agency_name LtsCrossAccountAccessV2#agency_name}.
	AgencyName *string `field:"required" json:"agencyName" yaml:"agencyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#agency_project_id LtsCrossAccountAccessV2#agency_project_id}.
	AgencyProjectId *string `field:"required" json:"agencyProjectId" yaml:"agencyProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_agency_group_id LtsCrossAccountAccessV2#log_agency_group_id}.
	LogAgencyGroupId *string `field:"required" json:"logAgencyGroupId" yaml:"logAgencyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_agency_group_name LtsCrossAccountAccessV2#log_agency_group_name}.
	LogAgencyGroupName *string `field:"required" json:"logAgencyGroupName" yaml:"logAgencyGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_agency_stream_id LtsCrossAccountAccessV2#log_agency_stream_id}.
	LogAgencyStreamId *string `field:"required" json:"logAgencyStreamId" yaml:"logAgencyStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_agency_stream_name LtsCrossAccountAccessV2#log_agency_stream_name}.
	LogAgencyStreamName *string `field:"required" json:"logAgencyStreamName" yaml:"logAgencyStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_group_id LtsCrossAccountAccessV2#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_group_name LtsCrossAccountAccessV2#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_stream_id LtsCrossAccountAccessV2#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#log_stream_name LtsCrossAccountAccessV2#log_stream_name}.
	LogStreamName *string `field:"required" json:"logStreamName" yaml:"logStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#name LtsCrossAccountAccessV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#id LtsCrossAccountAccessV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/lts_cross_account_access_v2#tags LtsCrossAccountAccessV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

