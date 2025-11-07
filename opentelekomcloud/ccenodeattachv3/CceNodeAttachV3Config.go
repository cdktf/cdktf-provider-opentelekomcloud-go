// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodeAttachV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#cluster_id CceNodeAttachV3#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#os CceNodeAttachV3#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#server_id CceNodeAttachV3#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#docker_base_size CceNodeAttachV3#docker_base_size}.
	DockerBaseSize *float64 `field:"optional" json:"dockerBaseSize" yaml:"dockerBaseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#id CceNodeAttachV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#k8s_tags CceNodeAttachV3#k8s_tags}.
	K8STags *map[string]*string `field:"optional" json:"k8STags" yaml:"k8STags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#key_pair CceNodeAttachV3#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#lvm_config CceNodeAttachV3#lvm_config}.
	LvmConfig *string `field:"optional" json:"lvmConfig" yaml:"lvmConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#max_pods CceNodeAttachV3#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#name CceNodeAttachV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#password CceNodeAttachV3#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#postinstall CceNodeAttachV3#postinstall}.
	Postinstall *string `field:"optional" json:"postinstall" yaml:"postinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#preinstall CceNodeAttachV3#preinstall}.
	Preinstall *string `field:"optional" json:"preinstall" yaml:"preinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#private_key CceNodeAttachV3#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#runtime CceNodeAttachV3#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#storage CceNodeAttachV3#storage}
	Storage *CceNodeAttachV3Storage `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#system_disk_kms_key_id CceNodeAttachV3#system_disk_kms_key_id}.
	SystemDiskKmsKeyId *string `field:"optional" json:"systemDiskKmsKeyId" yaml:"systemDiskKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#tags CceNodeAttachV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#taints CceNodeAttachV3#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_attach_v3#timeouts CceNodeAttachV3#timeouts}
	Timeouts *CceNodeAttachV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

