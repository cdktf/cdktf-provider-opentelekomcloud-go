// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudfgsfunctionsv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudFgsFunctionsV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#enterprise_project_id DataOpentelekomcloudFgsFunctionsV2#enterprise_project_id}.
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#id DataOpentelekomcloudFgsFunctionsV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#name DataOpentelekomcloudFgsFunctionsV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#package_name DataOpentelekomcloudFgsFunctionsV2#package_name}.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#runtime DataOpentelekomcloudFgsFunctionsV2#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/fgs_functions_v2#urn DataOpentelekomcloudFgsFunctionsV2#urn}.
	Urn *string `field:"optional" json:"urn" yaml:"urn"`
}

