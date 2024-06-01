// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asgroupv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AsGroupV1Config struct {
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
	// Whether to delete instances when they are removed from the AS group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#delete_instances AsGroupV1#delete_instances}
	DeleteInstances *string `field:"required" json:"deleteInstances" yaml:"deleteInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#delete_publicip AsGroupV1#delete_publicip}.
	DeletePublicip interface{} `field:"required" json:"deletePublicip" yaml:"deletePublicip"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#networks AsGroupV1#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#scaling_group_name AsGroupV1#scaling_group_name}.
	ScalingGroupName *string `field:"required" json:"scalingGroupName" yaml:"scalingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#vpc_id AsGroupV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#available_zones AsGroupV1#available_zones}.
	AvailableZones *[]*string `field:"optional" json:"availableZones" yaml:"availableZones"`
	// The cooling duration, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#cool_down_time AsGroupV1#cool_down_time}
	CoolDownTime *float64 `field:"optional" json:"coolDownTime" yaml:"coolDownTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#desire_instance_number AsGroupV1#desire_instance_number}.
	DesireInstanceNumber *float64 `field:"optional" json:"desireInstanceNumber" yaml:"desireInstanceNumber"`
	// The grace period for instance health check, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#health_periodic_audit_grace_period AsGroupV1#health_periodic_audit_grace_period}
	HealthPeriodicAuditGracePeriod *float64 `field:"optional" json:"healthPeriodicAuditGracePeriod" yaml:"healthPeriodicAuditGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#health_periodic_audit_method AsGroupV1#health_periodic_audit_method}.
	HealthPeriodicAuditMethod *string `field:"optional" json:"healthPeriodicAuditMethod" yaml:"healthPeriodicAuditMethod"`
	// The health check period for instances, in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#health_periodic_audit_time AsGroupV1#health_periodic_audit_time}
	HealthPeriodicAuditTime *float64 `field:"optional" json:"healthPeriodicAuditTime" yaml:"healthPeriodicAuditTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#id AsGroupV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#instance_terminate_policy AsGroupV1#instance_terminate_policy}.
	InstanceTerminatePolicy *string `field:"optional" json:"instanceTerminatePolicy" yaml:"instanceTerminatePolicy"`
	// lbaas_listeners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#lbaas_listeners AsGroupV1#lbaas_listeners}
	LbaasListeners interface{} `field:"optional" json:"lbaasListeners" yaml:"lbaasListeners"`
	// The system supports the binding of up to six classic LB listeners, the IDs of which are separated using a comma.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#lb_listener_id AsGroupV1#lb_listener_id}
	LbListenerId *string `field:"optional" json:"lbListenerId" yaml:"lbListenerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#max_instance_number AsGroupV1#max_instance_number}.
	MaxInstanceNumber *float64 `field:"optional" json:"maxInstanceNumber" yaml:"maxInstanceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#min_instance_number AsGroupV1#min_instance_number}.
	MinInstanceNumber *float64 `field:"optional" json:"minInstanceNumber" yaml:"minInstanceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#notifications AsGroupV1#notifications}.
	Notifications *[]*string `field:"optional" json:"notifications" yaml:"notifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#region AsGroupV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#scaling_configuration_id AsGroupV1#scaling_configuration_id}.
	ScalingConfigurationId *string `field:"optional" json:"scalingConfigurationId" yaml:"scalingConfigurationId"`
	// security_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#security_groups AsGroupV1#security_groups}
	SecurityGroups *AsGroupV1SecurityGroups `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#tags AsGroupV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.10/docs/resources/as_group_v1#timeouts AsGroupV1#timeouts}
	Timeouts *AsGroupV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

