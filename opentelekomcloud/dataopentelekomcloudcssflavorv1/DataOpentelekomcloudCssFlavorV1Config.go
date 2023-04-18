package dataopentelekomcloudcssflavorv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCssFlavorV1Config struct {
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
	// disk_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#disk_range DataOpentelekomcloudCssFlavorV1#disk_range}
	DiskRange *DataOpentelekomcloudCssFlavorV1DiskRange `field:"optional" json:"diskRange" yaml:"diskRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#id DataOpentelekomcloudCssFlavorV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#min_cpu DataOpentelekomcloudCssFlavorV1#min_cpu}.
	MinCpu *float64 `field:"optional" json:"minCpu" yaml:"minCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#min_ram DataOpentelekomcloudCssFlavorV1#min_ram}.
	MinRam *float64 `field:"optional" json:"minRam" yaml:"minRam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#name DataOpentelekomcloudCssFlavorV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#type DataOpentelekomcloudCssFlavorV1#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/data-sources/css_flavor_v1#version DataOpentelekomcloudCssFlavorV1#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

