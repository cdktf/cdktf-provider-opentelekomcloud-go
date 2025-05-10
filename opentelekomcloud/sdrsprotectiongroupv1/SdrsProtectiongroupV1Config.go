// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdrsprotectiongroupv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SdrsProtectiongroupV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#domain_id SdrsProtectiongroupV1#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#name SdrsProtectiongroupV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#source_availability_zone SdrsProtectiongroupV1#source_availability_zone}.
	SourceAvailabilityZone *string `field:"required" json:"sourceAvailabilityZone" yaml:"sourceAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#source_vpc_id SdrsProtectiongroupV1#source_vpc_id}.
	SourceVpcId *string `field:"required" json:"sourceVpcId" yaml:"sourceVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#target_availability_zone SdrsProtectiongroupV1#target_availability_zone}.
	TargetAvailabilityZone *string `field:"required" json:"targetAvailabilityZone" yaml:"targetAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#description SdrsProtectiongroupV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#dr_type SdrsProtectiongroupV1#dr_type}.
	DrType *string `field:"optional" json:"drType" yaml:"drType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#enable SdrsProtectiongroupV1#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#id SdrsProtectiongroupV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/sdrs_protectiongroup_v1#timeouts SdrsProtectiongroupV1#timeouts}
	Timeouts *SdrsProtectiongroupV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

