// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssquotasv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssQuotasV5Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#category DataOpentelekomcloudHssQuotasV5#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#charging_mode DataOpentelekomcloudHssQuotasV5#charging_mode}.
	ChargingMode *string `field:"optional" json:"chargingMode" yaml:"chargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#host_name DataOpentelekomcloudHssQuotasV5#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#id DataOpentelekomcloudHssQuotasV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#resource_id DataOpentelekomcloudHssQuotasV5#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#status DataOpentelekomcloudHssQuotasV5#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#used_status DataOpentelekomcloudHssQuotasV5#used_status}.
	UsedStatus *string `field:"optional" json:"usedStatus" yaml:"usedStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/hss_quotas_v5#version DataOpentelekomcloudHssQuotasV5#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

