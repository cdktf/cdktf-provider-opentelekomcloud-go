package vpcsubnetv1


type VpcSubnetV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_subnet_v1#create VpcSubnetV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_subnet_v1#delete VpcSubnetV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
