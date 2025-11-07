// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminiinstancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GeminiInstanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#availability_zone GeminiInstanceV3#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#flavor GeminiInstanceV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#name GeminiInstanceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#password GeminiInstanceV3#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#subnet_id GeminiInstanceV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#volume_size GeminiInstanceV3#volume_size}.
	VolumeSize *float64 `field:"required" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#vpc_id GeminiInstanceV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// backup_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#backup_strategy GeminiInstanceV3#backup_strategy}
	BackupStrategy *GeminiInstanceV3BackupStrategy `field:"optional" json:"backupStrategy" yaml:"backupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#configuration_id GeminiInstanceV3#configuration_id}.
	ConfigurationId *string `field:"optional" json:"configurationId" yaml:"configurationId"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#datastore GeminiInstanceV3#datastore}
	Datastore *GeminiInstanceV3Datastore `field:"optional" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#id GeminiInstanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#node_num GeminiInstanceV3#node_num}.
	NodeNum *float64 `field:"optional" json:"nodeNum" yaml:"nodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#period GeminiInstanceV3#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#security_group_id GeminiInstanceV3#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#ssl GeminiInstanceV3#ssl}.
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#tags GeminiInstanceV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/gemini_instance_v3#timeouts GeminiInstanceV3#timeouts}
	Timeouts *GeminiInstanceV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

