package vpcv1


type VpcV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_v1#create VpcV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpc_v1#delete VpcV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
