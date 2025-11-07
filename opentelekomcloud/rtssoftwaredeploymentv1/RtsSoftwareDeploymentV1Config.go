// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rtssoftwaredeploymentv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RtsSoftwareDeploymentV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#config_id RtsSoftwareDeploymentV1#config_id}.
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#server_id RtsSoftwareDeploymentV1#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#action RtsSoftwareDeploymentV1#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#id RtsSoftwareDeploymentV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#input_values RtsSoftwareDeploymentV1#input_values}.
	InputValues *map[string]*string `field:"optional" json:"inputValues" yaml:"inputValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#output_values RtsSoftwareDeploymentV1#output_values}.
	OutputValues *map[string]*string `field:"optional" json:"outputValues" yaml:"outputValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#region RtsSoftwareDeploymentV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#status RtsSoftwareDeploymentV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#status_reason RtsSoftwareDeploymentV1#status_reason}.
	StatusReason *string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#tenant_id RtsSoftwareDeploymentV1#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/rts_software_deployment_v1#timeouts RtsSoftwareDeploymentV1#timeouts}
	Timeouts *RtsSoftwareDeploymentV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

