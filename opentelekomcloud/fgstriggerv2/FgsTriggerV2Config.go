// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgstriggerv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsTriggerV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#event_data FgsTriggerV2#event_data}.
	EventData *string `field:"required" json:"eventData" yaml:"eventData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#function_urn FgsTriggerV2#function_urn}.
	FunctionUrn *string `field:"required" json:"functionUrn" yaml:"functionUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#type FgsTriggerV2#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#id FgsTriggerV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#status FgsTriggerV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/fgs_trigger_v2#timeouts FgsTriggerV2#timeouts}
	Timeouts *FgsTriggerV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

