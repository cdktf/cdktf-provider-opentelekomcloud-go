// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltshostaccessv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LtsHostAccessV3Config struct {
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
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#access_config LtsHostAccessV3#access_config}
	AccessConfig *LtsHostAccessV3AccessConfig `field:"required" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#log_group_id LtsHostAccessV3#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#log_stream_id LtsHostAccessV3#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#name LtsHostAccessV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#binary_collect LtsHostAccessV3#binary_collect}.
	BinaryCollect interface{} `field:"optional" json:"binaryCollect" yaml:"binaryCollect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#host_group_ids LtsHostAccessV3#host_group_ids}.
	HostGroupIds *[]*string `field:"optional" json:"hostGroupIds" yaml:"hostGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#id LtsHostAccessV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#log_split LtsHostAccessV3#log_split}.
	LogSplit interface{} `field:"optional" json:"logSplit" yaml:"logSplit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/lts_host_access_v3#tags LtsHostAccessV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

