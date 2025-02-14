// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmsresourcerecorderv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RmsResourceRecorderV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/rms_resource_recorder_v1#agency_name RmsResourceRecorderV1#agency_name}.
	AgencyName *string `field:"required" json:"agencyName" yaml:"agencyName"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/rms_resource_recorder_v1#selector RmsResourceRecorderV1#selector}
	Selector *RmsResourceRecorderV1Selector `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/rms_resource_recorder_v1#id RmsResourceRecorderV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// obs_channel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/rms_resource_recorder_v1#obs_channel RmsResourceRecorderV1#obs_channel}
	ObsChannel *RmsResourceRecorderV1ObsChannel `field:"optional" json:"obsChannel" yaml:"obsChannel"`
	// smn_channel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/rms_resource_recorder_v1#smn_channel RmsResourceRecorderV1#smn_channel}
	SmnChannel *RmsResourceRecorderV1SmnChannel `field:"optional" json:"smnChannel" yaml:"smnChannel"`
}

