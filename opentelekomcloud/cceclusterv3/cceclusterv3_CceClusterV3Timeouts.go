package cceclusterv3


type CceClusterV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_cluster_v3#create CceClusterV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_cluster_v3#delete CceClusterV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
