package dataopentelekomclouddmsproductv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudDmsProductV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#engine DataOpentelekomcloudDmsProductV1#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#instance_type DataOpentelekomcloudDmsProductV1#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#bandwidth DataOpentelekomcloudDmsProductV1#bandwidth}.
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#id DataOpentelekomcloudDmsProductV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#io_type DataOpentelekomcloudDmsProductV1#io_type}.
	IoType *string `field:"optional" json:"ioType" yaml:"ioType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#node_num DataOpentelekomcloudDmsProductV1#node_num}.
	NodeNum *string `field:"optional" json:"nodeNum" yaml:"nodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#partition_num DataOpentelekomcloudDmsProductV1#partition_num}.
	PartitionNum *string `field:"optional" json:"partitionNum" yaml:"partitionNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#storage DataOpentelekomcloudDmsProductV1#storage}.
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#storage_spec_code DataOpentelekomcloudDmsProductV1#storage_spec_code}.
	StorageSpecCode *string `field:"optional" json:"storageSpecCode" yaml:"storageSpecCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#version DataOpentelekomcloudDmsProductV1#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/data-sources/dms_product_v1#vm_specification DataOpentelekomcloudDmsProductV1#vm_specification}.
	VmSpecification *string `field:"optional" json:"vmSpecification" yaml:"vmSpecification"`
}

