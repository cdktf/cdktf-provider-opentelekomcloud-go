package ecsinstancev1


type EcsInstanceV1Nics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ecs_instance_v1#network_id EcsInstanceV1#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ecs_instance_v1#ip_address EcsInstanceV1#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}
