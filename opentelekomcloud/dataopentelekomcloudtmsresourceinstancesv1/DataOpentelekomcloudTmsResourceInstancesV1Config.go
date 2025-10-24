// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudtmsresourceinstancesv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudTmsResourceInstancesV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/tms_resource_instances_v1#resource_types DataOpentelekomcloudTmsResourceInstancesV1#resource_types}.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/tms_resource_instances_v1#tags DataOpentelekomcloudTmsResourceInstancesV1#tags}
	Tags interface{} `field:"required" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/tms_resource_instances_v1#id DataOpentelekomcloudTmsResourceInstancesV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/tms_resource_instances_v1#project_id DataOpentelekomcloudTmsResourceInstancesV1#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/data-sources/tms_resource_instances_v1#without_any_tag DataOpentelekomcloudTmsResourceInstancesV1#without_any_tag}.
	WithoutAnyTag interface{} `field:"optional" json:"withoutAnyTag" yaml:"withoutAnyTag"`
}

