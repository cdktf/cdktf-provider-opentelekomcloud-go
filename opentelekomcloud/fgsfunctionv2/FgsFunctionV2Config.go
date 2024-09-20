// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FgsFunctionV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#memory_size FgsFunctionV2#memory_size}.
	MemorySize *float64 `field:"required" json:"memorySize" yaml:"memorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#name FgsFunctionV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#runtime FgsFunctionV2#runtime}.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#timeout FgsFunctionV2#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#agency FgsFunctionV2#agency}.
	Agency *string `field:"optional" json:"agency" yaml:"agency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#app FgsFunctionV2#app}.
	App *string `field:"optional" json:"app" yaml:"app"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#app_agency FgsFunctionV2#app_agency}.
	AppAgency *string `field:"optional" json:"appAgency" yaml:"appAgency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#code_filename FgsFunctionV2#code_filename}.
	CodeFilename *string `field:"optional" json:"codeFilename" yaml:"codeFilename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#code_type FgsFunctionV2#code_type}.
	CodeType *string `field:"optional" json:"codeType" yaml:"codeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#code_url FgsFunctionV2#code_url}.
	CodeUrl *string `field:"optional" json:"codeUrl" yaml:"codeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#concurrency_num FgsFunctionV2#concurrency_num}.
	ConcurrencyNum *float64 `field:"optional" json:"concurrencyNum" yaml:"concurrencyNum"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#custom_image FgsFunctionV2#custom_image}
	CustomImage *FgsFunctionV2CustomImage `field:"optional" json:"customImage" yaml:"customImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#depend_list FgsFunctionV2#depend_list}.
	DependList *[]*string `field:"optional" json:"dependList" yaml:"dependList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#description FgsFunctionV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#encrypted_user_data FgsFunctionV2#encrypted_user_data}.
	EncryptedUserData *string `field:"optional" json:"encryptedUserData" yaml:"encryptedUserData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#func_code FgsFunctionV2#func_code}.
	FuncCode *string `field:"optional" json:"funcCode" yaml:"funcCode"`
	// func_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#func_mounts FgsFunctionV2#func_mounts}
	FuncMounts interface{} `field:"optional" json:"funcMounts" yaml:"funcMounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#functiongraph_version FgsFunctionV2#functiongraph_version}.
	FunctiongraphVersion *string `field:"optional" json:"functiongraphVersion" yaml:"functiongraphVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#gpu_memory FgsFunctionV2#gpu_memory}.
	GpuMemory *float64 `field:"optional" json:"gpuMemory" yaml:"gpuMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#handler FgsFunctionV2#handler}.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#id FgsFunctionV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#initializer_handler FgsFunctionV2#initializer_handler}.
	InitializerHandler *string `field:"optional" json:"initializerHandler" yaml:"initializerHandler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#initializer_timeout FgsFunctionV2#initializer_timeout}.
	InitializerTimeout *float64 `field:"optional" json:"initializerTimeout" yaml:"initializerTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#log_group_id FgsFunctionV2#log_group_id}.
	LogGroupId *string `field:"optional" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#log_group_name FgsFunctionV2#log_group_name}.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#log_topic_id FgsFunctionV2#log_topic_id}.
	LogTopicId *string `field:"optional" json:"logTopicId" yaml:"logTopicId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#log_topic_name FgsFunctionV2#log_topic_name}.
	LogTopicName *string `field:"optional" json:"logTopicName" yaml:"logTopicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#max_instance_num FgsFunctionV2#max_instance_num}.
	MaxInstanceNum *string `field:"optional" json:"maxInstanceNum" yaml:"maxInstanceNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#mount_user_group_id FgsFunctionV2#mount_user_group_id}.
	MountUserGroupId *float64 `field:"optional" json:"mountUserGroupId" yaml:"mountUserGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#mount_user_id FgsFunctionV2#mount_user_id}.
	MountUserId *float64 `field:"optional" json:"mountUserId" yaml:"mountUserId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#network_id FgsFunctionV2#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// reserved_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#reserved_instances FgsFunctionV2#reserved_instances}
	ReservedInstances interface{} `field:"optional" json:"reservedInstances" yaml:"reservedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#tags FgsFunctionV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#timeouts FgsFunctionV2#timeouts}
	Timeouts *FgsFunctionV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#user_data FgsFunctionV2#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#versions FgsFunctionV2#versions}
	Versions interface{} `field:"optional" json:"versions" yaml:"versions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/fgs_function_v2#vpc_id FgsFunctionV2#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

