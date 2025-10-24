// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlquotav3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaurusdbMysqlQuotaV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#enterprise_project_id TaurusdbMysqlQuotaV3#enterprise_project_id}.
	EnterpriseProjectId *string `field:"required" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#enterprise_project_name TaurusdbMysqlQuotaV3#enterprise_project_name}.
	EnterpriseProjectName *string `field:"required" json:"enterpriseProjectName" yaml:"enterpriseProjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#id TaurusdbMysqlQuotaV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#instance_quota TaurusdbMysqlQuotaV3#instance_quota}.
	InstanceQuota *float64 `field:"optional" json:"instanceQuota" yaml:"instanceQuota"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#ram_quota TaurusdbMysqlQuotaV3#ram_quota}.
	RamQuota *float64 `field:"optional" json:"ramQuota" yaml:"ramQuota"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_quota_v3#vcpus_quota TaurusdbMysqlQuotaV3#vcpus_quota}.
	VcpusQuota *float64 `field:"optional" json:"vcpusQuota" yaml:"vcpusQuota"`
}

