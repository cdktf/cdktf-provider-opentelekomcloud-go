// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomclouddehserverv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudDehServerV1Config struct {
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
	// addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/data-sources/deh_server_v1#addresses DataOpentelekomcloudDehServerV1#addresses}
	Addresses interface{} `field:"optional" json:"addresses" yaml:"addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/data-sources/deh_server_v1#dedicated_host_id DataOpentelekomcloudDehServerV1#dedicated_host_id}.
	DedicatedHostId *string `field:"optional" json:"dedicatedHostId" yaml:"dedicatedHostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/data-sources/deh_server_v1#id DataOpentelekomcloudDehServerV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/data-sources/deh_server_v1#region DataOpentelekomcloudDehServerV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/data-sources/deh_server_v1#server_id DataOpentelekomcloudDehServerV1#server_id}.
	ServerId *string `field:"optional" json:"serverId" yaml:"serverId"`
}

