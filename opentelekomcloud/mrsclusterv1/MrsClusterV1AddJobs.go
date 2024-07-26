// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrsclusterv1


type MrsClusterV1AddJobs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#jar_path MrsClusterV1#jar_path}.
	JarPath *string `field:"required" json:"jarPath" yaml:"jarPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#job_name MrsClusterV1#job_name}.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#job_type MrsClusterV1#job_type}.
	JobType *float64 `field:"required" json:"jobType" yaml:"jobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#submit_job_once_cluster_run MrsClusterV1#submit_job_once_cluster_run}.
	SubmitJobOnceClusterRun interface{} `field:"required" json:"submitJobOnceClusterRun" yaml:"submitJobOnceClusterRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#arguments MrsClusterV1#arguments}.
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#file_action MrsClusterV1#file_action}.
	FileAction *string `field:"optional" json:"fileAction" yaml:"fileAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#hive_script_path MrsClusterV1#hive_script_path}.
	HiveScriptPath *string `field:"optional" json:"hiveScriptPath" yaml:"hiveScriptPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#hql MrsClusterV1#hql}.
	Hql *string `field:"optional" json:"hql" yaml:"hql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#input MrsClusterV1#input}.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#job_log MrsClusterV1#job_log}.
	JobLog *string `field:"optional" json:"jobLog" yaml:"jobLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#output MrsClusterV1#output}.
	Output *string `field:"optional" json:"output" yaml:"output"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/mrs_cluster_v1#shutdown_cluster MrsClusterV1#shutdown_cluster}.
	ShutdownCluster interface{} `field:"optional" json:"shutdownCluster" yaml:"shutdownCluster"`
}

