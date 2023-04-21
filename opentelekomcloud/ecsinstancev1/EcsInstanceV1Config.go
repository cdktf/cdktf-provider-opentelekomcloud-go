package ecsinstancev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsInstanceV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#availability_zone EcsInstanceV1#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#flavor EcsInstanceV1#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#image_id EcsInstanceV1#image_id}.
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#name EcsInstanceV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// nics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#nics EcsInstanceV1#nics}
	Nics interface{} `field:"required" json:"nics" yaml:"nics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#vpc_id EcsInstanceV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#auto_recovery EcsInstanceV1#auto_recovery}.
	AutoRecovery interface{} `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#data_disks EcsInstanceV1#data_disks}
	DataDisks interface{} `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#delete_disks_on_termination EcsInstanceV1#delete_disks_on_termination}.
	DeleteDisksOnTermination interface{} `field:"optional" json:"deleteDisksOnTermination" yaml:"deleteDisksOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#id EcsInstanceV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#key_name EcsInstanceV1#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#password EcsInstanceV1#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#security_groups EcsInstanceV1#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#system_disk_kms_id EcsInstanceV1#system_disk_kms_id}.
	SystemDiskKmsId *string `field:"optional" json:"systemDiskKmsId" yaml:"systemDiskKmsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#system_disk_size EcsInstanceV1#system_disk_size}.
	SystemDiskSize *float64 `field:"optional" json:"systemDiskSize" yaml:"systemDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#system_disk_type EcsInstanceV1#system_disk_type}.
	SystemDiskType *string `field:"optional" json:"systemDiskType" yaml:"systemDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#tags EcsInstanceV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#timeouts EcsInstanceV1#timeouts}
	Timeouts *EcsInstanceV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/ecs_instance_v1#user_data EcsInstanceV1#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

