package ecsinstancev1


type EcsInstanceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ecs_instance_v1#create EcsInstanceV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ecs_instance_v1#delete EcsInstanceV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/ecs_instance_v1#update EcsInstanceV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
