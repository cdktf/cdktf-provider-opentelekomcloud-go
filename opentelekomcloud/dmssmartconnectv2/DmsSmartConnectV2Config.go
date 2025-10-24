// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmssmartconnectv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsSmartConnectV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#instance_id DmsSmartConnectV2#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#bandwidth DmsSmartConnectV2#bandwidth}.
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#id DmsSmartConnectV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#node_count DmsSmartConnectV2#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#storage_spec_code DmsSmartConnectV2#storage_spec_code}.
	StorageSpecCode *string `field:"optional" json:"storageSpecCode" yaml:"storageSpecCode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/dms_smart_connect_v2#timeouts DmsSmartConnectV2#timeouts}
	Timeouts *DmsSmartConnectV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

