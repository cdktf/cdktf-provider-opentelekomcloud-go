package dataopentelekomcloudrtsstackresourcev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudRtsStackResourceV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#stack_name DataOpentelekomcloudRtsStackResourceV1#stack_name}.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#id DataOpentelekomcloudRtsStackResourceV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#physical_resource_id DataOpentelekomcloudRtsStackResourceV1#physical_resource_id}.
	PhysicalResourceId *string `field:"optional" json:"physicalResourceId" yaml:"physicalResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#region DataOpentelekomcloudRtsStackResourceV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#resource_name DataOpentelekomcloudRtsStackResourceV1#resource_name}.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/data-sources/rts_stack_resource_v1#resource_type DataOpentelekomcloudRtsStackResourceV1#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

