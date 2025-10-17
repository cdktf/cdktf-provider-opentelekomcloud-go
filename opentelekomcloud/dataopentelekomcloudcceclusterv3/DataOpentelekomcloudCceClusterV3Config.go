// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcceclusterv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCceClusterV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#cluster_type DataOpentelekomcloudCceClusterV3#cluster_type}.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#id DataOpentelekomcloudCceClusterV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#name DataOpentelekomcloudCceClusterV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#region DataOpentelekomcloudCceClusterV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#status DataOpentelekomcloudCceClusterV3#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/cce_cluster_v3#vpc_id DataOpentelekomcloudCceClusterV3#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

