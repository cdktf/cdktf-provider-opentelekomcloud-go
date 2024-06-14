// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudvpceipv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudVpcEipV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#bandwidth_id DataOpentelekomcloudVpcEipV1#bandwidth_id}.
	BandwidthId *string `field:"optional" json:"bandwidthId" yaml:"bandwidthId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#id DataOpentelekomcloudVpcEipV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#name_regex DataOpentelekomcloudVpcEipV1#name_regex}.
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#port_id DataOpentelekomcloudVpcEipV1#port_id}.
	PortId *string `field:"optional" json:"portId" yaml:"portId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#private_ip_address DataOpentelekomcloudVpcEipV1#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#public_ip_address DataOpentelekomcloudVpcEipV1#public_ip_address}.
	PublicIpAddress *string `field:"optional" json:"publicIpAddress" yaml:"publicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#region DataOpentelekomcloudVpcEipV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#status DataOpentelekomcloudVpcEipV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/vpc_eip_v1#tags DataOpentelekomcloudVpcEipV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

