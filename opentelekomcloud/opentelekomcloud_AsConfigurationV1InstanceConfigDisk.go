// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type AsConfigurationV1InstanceConfigDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#disk_type AsConfigurationV1#disk_type}.
	DiskType *string `field:"required" json:"diskType" yaml:"diskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#size AsConfigurationV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#volume_type AsConfigurationV1#volume_type}.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#kms_id AsConfigurationV1#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}

