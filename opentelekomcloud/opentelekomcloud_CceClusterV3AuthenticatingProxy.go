// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type CceClusterV3AuthenticatingProxy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_cluster_v3#ca CceClusterV3#ca}.
	Ca *string `field:"required" json:"ca" yaml:"ca"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_cluster_v3#cert CceClusterV3#cert}.
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_cluster_v3#private_key CceClusterV3#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}
