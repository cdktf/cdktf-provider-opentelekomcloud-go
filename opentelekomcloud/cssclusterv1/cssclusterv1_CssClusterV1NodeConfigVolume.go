package cssclusterv1


type CssClusterV1NodeConfigVolume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_cluster_v1#size CssClusterV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_cluster_v1#volume_type CssClusterV1#volume_type}.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_cluster_v1#encryption_key CssClusterV1#encryption_key}.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}
