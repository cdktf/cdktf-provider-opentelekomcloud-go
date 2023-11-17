// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomclouddcsmaintainwindowv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudDcsMaintainwindowV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/data-sources/dcs_maintainwindow_v1#begin DataOpentelekomcloudDcsMaintainwindowV1#begin}.
	Begin *string `field:"optional" json:"begin" yaml:"begin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/data-sources/dcs_maintainwindow_v1#default DataOpentelekomcloudDcsMaintainwindowV1#default}.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/data-sources/dcs_maintainwindow_v1#end DataOpentelekomcloudDcsMaintainwindowV1#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/data-sources/dcs_maintainwindow_v1#id DataOpentelekomcloudDcsMaintainwindowV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/data-sources/dcs_maintainwindow_v1#seq DataOpentelekomcloudDcsMaintainwindowV1#seq}.
	Seq *float64 `field:"optional" json:"seq" yaml:"seq"`
}

