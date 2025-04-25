// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscceaccessv3


type LtsCceAccessV3AccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#path_type LtsCceAccessV3#path_type}.
	PathType *string `field:"required" json:"pathType" yaml:"pathType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#black_paths LtsCceAccessV3#black_paths}.
	BlackPaths *[]*string `field:"optional" json:"blackPaths" yaml:"blackPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#container_name_regex LtsCceAccessV3#container_name_regex}.
	ContainerNameRegex *string `field:"optional" json:"containerNameRegex" yaml:"containerNameRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#exclude_envs LtsCceAccessV3#exclude_envs}.
	ExcludeEnvs *map[string]*string `field:"optional" json:"excludeEnvs" yaml:"excludeEnvs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#exclude_k8s_labels LtsCceAccessV3#exclude_k8s_labels}.
	ExcludeK8SLabels *map[string]*string `field:"optional" json:"excludeK8SLabels" yaml:"excludeK8SLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#exclude_labels LtsCceAccessV3#exclude_labels}.
	ExcludeLabels *map[string]*string `field:"optional" json:"excludeLabels" yaml:"excludeLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#include_envs LtsCceAccessV3#include_envs}.
	IncludeEnvs *map[string]*string `field:"optional" json:"includeEnvs" yaml:"includeEnvs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#include_k8s_labels LtsCceAccessV3#include_k8s_labels}.
	IncludeK8SLabels *map[string]*string `field:"optional" json:"includeK8SLabels" yaml:"includeK8SLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#include_labels LtsCceAccessV3#include_labels}.
	IncludeLabels *map[string]*string `field:"optional" json:"includeLabels" yaml:"includeLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#log_envs LtsCceAccessV3#log_envs}.
	LogEnvs *map[string]*string `field:"optional" json:"logEnvs" yaml:"logEnvs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#log_k8s LtsCceAccessV3#log_k8s}.
	LogK8S *map[string]*string `field:"optional" json:"logK8S" yaml:"logK8S"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#log_labels LtsCceAccessV3#log_labels}.
	LogLabels *map[string]*string `field:"optional" json:"logLabels" yaml:"logLabels"`
	// multi_log_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#multi_log_format LtsCceAccessV3#multi_log_format}
	MultiLogFormat *LtsCceAccessV3AccessConfigMultiLogFormat `field:"optional" json:"multiLogFormat" yaml:"multiLogFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#name_space_regex LtsCceAccessV3#name_space_regex}.
	NameSpaceRegex *string `field:"optional" json:"nameSpaceRegex" yaml:"nameSpaceRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#paths LtsCceAccessV3#paths}.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#pod_name_regex LtsCceAccessV3#pod_name_regex}.
	PodNameRegex *string `field:"optional" json:"podNameRegex" yaml:"podNameRegex"`
	// single_log_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#single_log_format LtsCceAccessV3#single_log_format}
	SingleLogFormat *LtsCceAccessV3AccessConfigSingleLogFormat `field:"optional" json:"singleLogFormat" yaml:"singleLogFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#stderr LtsCceAccessV3#stderr}.
	Stderr interface{} `field:"optional" json:"stderr" yaml:"stderr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/lts_cce_access_v3#stdout LtsCceAccessV3#stdout}.
	Stdout interface{} `field:"optional" json:"stdout" yaml:"stdout"`
}

