package computeinstancev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#name ComputeInstanceV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#access_ip_v4 ComputeInstanceV2#access_ip_v4}.
	AccessIpV4 *string `field:"optional" json:"accessIpV4" yaml:"accessIpV4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#access_ip_v6 ComputeInstanceV2#access_ip_v6}.
	AccessIpV6 *string `field:"optional" json:"accessIpV6" yaml:"accessIpV6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#admin_pass ComputeInstanceV2#admin_pass}.
	AdminPass *string `field:"optional" json:"adminPass" yaml:"adminPass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#auto_recovery ComputeInstanceV2#auto_recovery}.
	AutoRecovery interface{} `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#availability_zone ComputeInstanceV2#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#block_device ComputeInstanceV2#block_device}
	BlockDevice interface{} `field:"optional" json:"blockDevice" yaml:"blockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#config_drive ComputeInstanceV2#config_drive}.
	ConfigDrive interface{} `field:"optional" json:"configDrive" yaml:"configDrive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#flavor_id ComputeInstanceV2#flavor_id}.
	FlavorId *string `field:"optional" json:"flavorId" yaml:"flavorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#flavor_name ComputeInstanceV2#flavor_name}.
	FlavorName *string `field:"optional" json:"flavorName" yaml:"flavorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#id ComputeInstanceV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#image_id ComputeInstanceV2#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#image_name ComputeInstanceV2#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#key_pair ComputeInstanceV2#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#metadata ComputeInstanceV2#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#network ComputeInstanceV2#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#power_state ComputeInstanceV2#power_state}.
	PowerState *string `field:"optional" json:"powerState" yaml:"powerState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#region ComputeInstanceV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scheduler_hints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#scheduler_hints ComputeInstanceV2#scheduler_hints}
	SchedulerHints interface{} `field:"optional" json:"schedulerHints" yaml:"schedulerHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#security_groups ComputeInstanceV2#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#ssh_private_key_path ComputeInstanceV2#ssh_private_key_path}.
	SshPrivateKeyPath *string `field:"optional" json:"sshPrivateKeyPath" yaml:"sshPrivateKeyPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#stop_before_destroy ComputeInstanceV2#stop_before_destroy}.
	StopBeforeDestroy interface{} `field:"optional" json:"stopBeforeDestroy" yaml:"stopBeforeDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#tags ComputeInstanceV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#timeouts ComputeInstanceV2#timeouts}
	Timeouts *ComputeInstanceV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/compute_instance_v2#user_data ComputeInstanceV2#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

